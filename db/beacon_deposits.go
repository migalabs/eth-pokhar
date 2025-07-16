package db

import (
	"strings"
	"time"

	"github.com/google/uuid"
	pgx "github.com/jackc/pgx/v5"
	"github.com/migalabs/eth-pokhar/models"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

// Postgres intregration variables
var (
	selectLastDeposit = `
	SELECT *
	FROM t_beacon_deposits
	ORDER BY f_block_num DESC
	LIMIT 1;
	`
)

func (p *PostgresDBService) ObtainLastDeposit() (models.BeaconDeposit, error) {
	conn, err := p.psqlPool.Acquire(p.ctx)
	if err != nil {
		return models.BeaconDeposit{}, errors.Wrap(err, "error acquiring connection")
	}
	defer conn.Release()
	rows, err := conn.Query(p.ctx, selectLastDeposit)
	if err != nil {
		rows.Close()
		return models.BeaconDeposit{}, errors.Wrap(err, "error obtaining last epoch from database")
	}
	deposit := models.BeaconDeposit{}
	rows.Next()
	err = rows.Scan(&deposit.BlockNum, &deposit.Depositor, &deposit.TxHash, &deposit.ValidatorPubkey, &deposit.WithdrawalAddress)
	if err != nil {
		log.Errorf("error scanning last deposit: %s", err.Error())
	}
	rows.Close()
	return deposit, nil
}

func (p *PostgresDBService) CopyBeaconDeposits(rowSrc []models.BeaconDeposit) int64 {
	if len(rowSrc) == 0 {
		return 0
	}
	p.writerThreadsWG.Add(1)
	defer p.writerThreadsWG.Done()
	startTime := time.Now()

	// Generate a random text to append to the table name
	randomText := uuid.New().String()

	// Create the temporary table name with the random text
	tempTableName := "temp_deposits_" + strings.ReplaceAll(randomText, "-", "_")

	deposits := beaconDepositToCopyData(rowSrc)

	conn, err := p.psqlPool.Acquire(p.ctx)
	if err != nil {
		wlog.Fatalf("could not acquire connection: %s", err.Error())
	}
	defer conn.Release()

	// Create a temporary table with a unique constraint
	_, err = conn.Exec(p.ctx, `
        CREATE TEMP TABLE IF NOT EXISTS `+tempTableName+` (
            f_block_num bigint,
            f_depositor text,
            f_tx_hash text,
            f_validator_pubkey text,
			f_withdrawal_address text NOT NULL DEFAULT ''
        )
    `)
	if err != nil {
		wlog.Fatalf("could not create temporary table: %s", err.Error())
	}

	// Copy data into the temporary table, ignoring duplicates
	_, err = conn.CopyFrom(
		p.ctx,
		pgx.Identifier{tempTableName},
		[]string{"f_block_num", "f_depositor", "f_tx_hash", "f_validator_pubkey", "f_withdrawal_address"},
		pgx.CopyFromRows(deposits),
	)
	if err != nil {
		wlog.Fatalf("could not copy rows into temporary table: %s", err.Error())
	}

	// Insert non-duplicate rows from the temporary table into the target table
	count, err := conn.Exec(p.ctx, `
        INSERT INTO t_beacon_deposits (f_block_num, f_depositor, f_tx_hash, f_validator_pubkey, f_withdrawal_address)
        SELECT f_block_num, f_depositor, f_tx_hash, f_validator_pubkey, f_withdrawal_address
        FROM `+tempTableName+`
        ON CONFLICT DO NOTHING
    `)
	if err != nil {
		wlog.Fatalf("could not insert rows into target table: %s", err.Error())
	}

	// Drop the temporary table
	_, err = conn.Exec(p.ctx, `DROP TABLE IF EXISTS `+tempTableName)
	if err != nil {
		wlog.Fatalf("could not drop temporary table: %s", err.Error())
	}
	if count.RowsAffected() > 0 {
		wlog.Debugf("persisted %d rows in %f seconds", count.RowsAffected(), time.Since(startTime).Seconds())
	}

	return count.RowsAffected()
}

func beaconDepositToCopyData(rows []models.BeaconDeposit) [][]interface{} {
	var copyData [][]interface{}

	for _, row := range rows {
		copyData = append(copyData, []interface{}{row.BlockNum, row.Depositor, row.TxHash, row.ValidatorPubkey, row.WithdrawalAddress})
	}

	return copyData
}
