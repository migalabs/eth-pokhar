package db

import (
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
