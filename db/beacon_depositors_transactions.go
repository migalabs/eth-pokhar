package db

/*

This file together with the model, has all the needed methods to interact with the epoch_metrics table of the database

*/

import (
	"time"

	pgx "github.com/jackc/pgx/v5"
	"github.com/migalabs/eth-pokhar/models"
)

// Postgres intregration variables
var (
	UpsertTransaction = `
	INSERT INTO t_beacon_depositors_transactions (
		f_block_num,
		f_value,
		f_from,
		f_to,
		f_tx_hash,
		f_depositor)
		VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT ON CONSTRAINT t_beacon_depositors_transactions_pkey
			DO NOTHING;
	`

	SelectCheckpointPerDepositor = `
	SELECT d.f_depositor, COALESCE(MAX(t.f_block_num), 0) f_max_block_num
	FROM t_beacon_deposits d
	LEFT JOIN t_beacon_depositors_transactions t ON d.f_depositor = t.f_depositor
	GROUP BY d.f_depositor;
	`
)

func insertTransaction(inputTransaction models.Transaction) (string, []interface{}) {
	resultArgs := make([]interface{}, 0)
	resultArgs = append(resultArgs, inputTransaction.BlockNum)
	resultArgs = append(resultArgs, inputTransaction.Value)
	resultArgs = append(resultArgs, inputTransaction.From)
	resultArgs = append(resultArgs, inputTransaction.To)
	resultArgs = append(resultArgs, inputTransaction.TxHash)
	resultArgs = append(resultArgs, inputTransaction.Depositor)
	return UpsertTransaction, resultArgs
}

func (p *PostgresDBService) ObtainCheckpointPerDepositor() ([]models.DepositorCheckpoint, error) {
	rows, err := p.psqlPool.Query(p.ctx, SelectCheckpointPerDepositor)
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

func TransactionOperation(inputTransaction models.Transaction) (string, []interface{}) {
	q, args := insertTransaction(inputTransaction)
	return q, args
}

func (p *PostgresDBService) CopyTransactions(rowSrc []models.Transaction) int64 {
	if len(rowSrc) == 0 {
		return 0
	}
	startTime := time.Now()

	transactions := transactionsToCopyData(rowSrc)

	// Copy data into the target table, ignoring duplicates
	_, err := p.psqlPool.CopyFrom(
		p.ctx,
		pgx.Identifier{"t_beacon_depositors_transactions"},
		[]string{"f_block_num", "f_value", "f_from", "f_to", "f_tx_hash", "f_depositor"},
		pgx.CopyFromRows(transactions),
	)

	if err != nil {
		wlog.Fatalf("could not copy rows into target table: %s", err.Error())
	}

	wlog.Infof("Persisted %d rows in %f seconds", len(transactions), time.Since(startTime).Seconds())

	return int64(len(transactions))
}

func transactionsToCopyData(rows []models.Transaction) [][]interface{} {
	var copyData [][]interface{}

	for _, row := range rows {
		copyData = append(copyData, []interface{}{row.BlockNum, row.Value, row.From, row.To, row.TxHash, row.Depositor})
	}

	return copyData
}
