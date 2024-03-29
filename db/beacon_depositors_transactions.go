package db

/*

This file together with the model, has all the needed methods to interact with the epoch_metrics table of the database

*/

import (
	"strings"
	"time"

	"github.com/google/uuid"
	pgx "github.com/jackc/pgx/v5"
	"github.com/migalabs/eth-pokhar/models"
	"github.com/pkg/errors"
)

// Postgres intregration variables
var (
	selectCheckpointPerDepositor = `
	WITH max_block_per_depositor AS (
		select f_depositor, MAX(f_block_num) f_max_block_num
		from t_beacon_depositors_transactions
		group by f_depositor)
	
	SELECT d.f_depositor, COALESCE(MAX(t.f_max_block_num), 0) f_max_block_num
		FROM t_beacon_deposits d
		LEFT JOIN max_block_per_depositor t ON d.f_depositor = t.f_depositor
		GROUP BY d.f_depositor
		ORDER BY f_max_block_num ASC;
	`
)

func (p *PostgresDBService) ObtainCheckpointPerDepositor() ([]models.DepositorCheckpoint, error) {
	conn, err := p.psqlPool.Acquire(p.ctx)
	if err != nil {
		return nil, errors.Wrap(err, "error acquiring database connection")
	}
	defer conn.Release()

	rows, err := conn.Query(p.ctx, selectCheckpointPerDepositor)
	if err != nil {
		rows.Close()
		return nil, err
	}
	defer rows.Close()

	var checkpoints []models.DepositorCheckpoint
	for rows.Next() {
		var checkpoint models.DepositorCheckpoint
		err = rows.Scan(&checkpoint.Depositor, &checkpoint.Checkpoint)
		if err != nil {
			return nil, err
		}
		checkpoints = append(checkpoints, checkpoint)
	}
	return checkpoints, nil
}

func (p *PostgresDBService) CopyTransactions(rowSrc []models.Transaction) int64 {
	if len(rowSrc) == 0 {
		return 0
	}
	startTime := time.Now()

	// Generate a random text to append to the table name
	randomText := uuid.New().String()

	// Create the temporary table name with the random text
	tempTableName := "temp_transactions_" + strings.ReplaceAll(randomText, "-", "_")

	deposits := transactionsToCopyData(rowSrc)

	// Acquire a connection from the pool
	conn, err := p.psqlPool.Acquire(p.ctx)
	if err != nil {
		wlog.Fatalf("Unable to acquire connection from pool: %s", err.Error())
	}
	defer conn.Release()

	// Create a temporary table with a unique constraint
	_, err = conn.Exec(p.ctx, `
        CREATE TEMP TABLE IF NOT EXISTS `+tempTableName+` (
            f_block_num bigint,
            f_value numeric,
            f_from text,
            f_to text,
            f_tx_hash text,
            f_depositor text,
            UNIQUE (f_tx_hash)
        )
    `)
	if err != nil {
		wlog.Fatalf("could not create temporary table: %s", err.Error())
	}

	// Copy data into the temporary table, ignoring duplicates
	_, err = conn.CopyFrom(
		p.ctx,
		pgx.Identifier{tempTableName},
		[]string{"f_block_num", "f_value", "f_from", "f_to", "f_tx_hash", "f_depositor"},
		pgx.CopyFromRows(deposits),
	)
	if err != nil {
		wlog.Fatalf("could not copy rows into temporary table: %s", err.Error())
	}

	// Insert non-duplicate rows from the temporary table into the target table
	count, err := conn.Exec(p.ctx, `
        INSERT INTO t_beacon_depositors_transactions (f_block_num, f_value, f_from, f_to, f_tx_hash, f_depositor)
        SELECT f_block_num, f_value, f_from, f_to, f_tx_hash, f_depositor
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

func transactionsToCopyData(rows []models.Transaction) [][]interface{} {
	var copyData [][]interface{}

	for _, row := range rows {
		copyData = append(copyData, []interface{}{row.BlockNum, row.Value, row.From, row.To, row.TxHash, row.Depositor})
	}

	return copyData
}
