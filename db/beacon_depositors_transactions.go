package db

/*

This file together with the model, has all the needed methods to interact with the epoch_metrics table of the database

*/

import (
	"github.com/migalabs/eth-pokhar/models"
)

// Postgres intregration variables
var (
	upsertTransaction = `
	INSERT INTO t_beacon_depositors_transactions (
		f_block_num,
		f_value,
		f_from,
		f_to,
		f_tx_hash,)
		VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT ON CONSTRAINT t_beacon_depositors_transactions_pkey
			DO NOTHING;
	`
)

func insertTransaction(inputTransaction models.Transaction) (string, []interface{}) {
	resultArgs := make([]interface{}, 0)
	resultArgs = append(resultArgs, inputTransaction.BlockNum)
	resultArgs = append(resultArgs, inputTransaction.Value)
	resultArgs = append(resultArgs, inputTransaction.From)
	resultArgs = append(resultArgs, inputTransaction.To)
	resultArgs = append(resultArgs, inputTransaction.TxHash)
	return upsertTransaction, resultArgs
}

func TransactionOperation(inputTransaction models.Transaction) (string, []interface{}) {
	q, args := insertTransaction(inputTransaction)
	return q, args
}
