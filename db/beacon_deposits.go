package db

import (
	"time"

	pgx "github.com/jackc/pgx/v5"
	"github.com/migalabs/eth-pokhar/models"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// Postgres intregration variables
var (
	UpsertBeaconDeposit = `
	INSERT INTO t_beacon_deposits (
		f_block_num,
		f_depositor,
		f_tx_hash,
		f_validator_pubkey)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT ON CONSTRAINT t_beacon_deposits_pkey
			DO NOTHING;
	`
	SelectLastDeposit = `
	SELECT *
	FROM t_beacon_deposits
	ORDER BY f_block_num DESC
	LIMIT 1;
	`
)

func insertBeaconDeposit(inputBeaconDeposit models.BeaconDeposit) (string, []interface{}) {
	resultArgs := make([]interface{}, 0)
	resultArgs = append(resultArgs, inputBeaconDeposit.BlockNum)
	resultArgs = append(resultArgs, inputBeaconDeposit.Depositor)
	resultArgs = append(resultArgs, inputBeaconDeposit.TxHash)
	resultArgs = append(resultArgs, inputBeaconDeposit.ValidatorPubkey)
	return UpsertBeaconDeposit, resultArgs
}

func (p *PostgresDBService) ObtainLastDeposit() (models.BeaconDeposit, error) {
	rows, err := p.psqlPool.Query(p.ctx, SelectLastDeposit)
	if err != nil {
		rows.Close()
		return models.BeaconDeposit{}, errors.Wrap(err, "error obtaining last epoch from database")
	}
	deposit := models.BeaconDeposit{}
	rows.Next()
	err = rows.Scan(&deposit.BlockNum, &deposit.Depositor, &deposit.TxHash, &deposit.ValidatorPubkey)
	if err != nil {
		log.Error(err)
	}
	rows.Close()
	return deposit, nil
}

func BeaconDepositOperation(inputBeaconDeposit models.BeaconDeposit) (string, []interface{}) {
	q, args := insertBeaconDeposit(inputBeaconDeposit)
	return q, args
}

func (p *PostgresDBService) CopyBeaconDeposits(rowSrc []models.BeaconDeposit) int64 {

	startTime := time.Now()

	deposits := beaconDepositToCopyData(rowSrc)

	// Create a temporary table with a unique constraint
	_, err := p.psqlPool.Exec(p.ctx, `
        CREATE TEMP TABLE temp_deposits (
            f_block_num bigint,
            f_depositor text,
            f_tx_hash text,
            f_validator_pubkey text,
            UNIQUE (f_tx_hash, f_validator_pubkey)
        )
    `)
	if err != nil {
		wlog.Fatalf("could not create temporary table: %s", err.Error())
	}

	// Add the trigger function to the temporary table
	_, err = p.psqlPool.Exec(p.ctx, `
		CREATE OR REPLACE FUNCTION prevent_duplicates()
		RETURNS TRIGGER AS $$
		BEGIN
			IF NEW.f_tx_hash IS NOT NULL AND NEW.f_validator_pubkey IS NOT NULL THEN
				IF EXISTS (
					SELECT 1 FROM temp_deposits
					WHERE f_tx_hash = NEW.f_tx_hash AND f_validator_pubkey = NEW.f_validator_pubkey
				) THEN
					RETURN NULL; -- Ignore duplicate
				END IF;
			END IF;

			RETURN NEW;
		END;
		$$ LANGUAGE plpgsql;

		CREATE TRIGGER prevent_duplicates_trigger
		BEFORE INSERT OR UPDATE
		ON temp_deposits
		FOR EACH ROW
		EXECUTE FUNCTION prevent_duplicates();
	`)
	if err != nil {
		wlog.Fatalf("could not add trigger to temporary table: %s", err.Error())
	}

	// Copy data into the temporary table, ignoring duplicates
	_, err = p.psqlPool.CopyFrom(
		p.ctx,
		pgx.Identifier{"temp_deposits"},
		[]string{"f_block_num", "f_depositor", "f_tx_hash", "f_validator_pubkey"},
		pgx.CopyFromRows(deposits),
	)
	if err != nil {
		wlog.Fatalf("could not copy rows into temporary table: %s", err.Error())
	}

	// Insert non-duplicate rows from the temporary table into the target table
	count, err := p.psqlPool.Exec(p.ctx, `
        INSERT INTO t_beacon_deposits (f_block_num, f_depositor, f_tx_hash, f_validator_pubkey)
        SELECT f_block_num, f_depositor, f_tx_hash, f_validator_pubkey
        FROM temp_deposits
        ON CONFLICT DO NOTHING
    `)
	if err != nil {
		wlog.Fatalf("could not insert rows into target table: %s", err.Error())
	}

	// Drop the temporary table
	_, err = p.psqlPool.Exec(p.ctx, `DROP TABLE temp_deposits`)
	if err != nil {
		wlog.Fatalf("could not drop temporary table: %s", err.Error())
	}

	wlog.Infof("persisted %d rows in %f seconds", count.RowsAffected(), time.Since(startTime).Seconds())

	return count.RowsAffected()
}

func beaconDepositToCopyData(rows []models.BeaconDeposit) [][]interface{} {
	var copyData [][]interface{}

	for _, row := range rows {
		copyData = append(copyData, []interface{}{row.BlockNum, row.Depositor, row.TxHash, row.ValidatorPubkey})
	}

	return copyData
}
