package db

import (
	"github.com/migalabs/eth-pokhar/models"
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
		ON CONFLICT ON CONSTRAINT t_beacon_deposits_pkey_pkey
			DO NOTHING;
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

func BeaconDepositOperation(inputBeaconDeposit models.BeaconDeposit) (string, []interface{}) {
	q, args := insertBeaconDeposit(inputBeaconDeposit)
	return q, args
}
