package db

import "github.com/pkg/errors"

const (
	applyWithdrawalAddressInsertQuery = `
		INSERT INTO t_identified_validators (
			f_validator_pubkey,
			f_pool_name
		)
		SELECT DISTINCT 
			t1.f_validator_pubkey, 
			t2.f_pool_name 
		FROM (
			SELECT 
				f_validator_pubkey,
				f_withdrawal_address,
				ROW_NUMBER() OVER (
					PARTITION BY f_validator_pubkey 
					ORDER BY f_block_num DESC
				) as rn
			FROM t_beacon_deposits
			WHERE f_withdrawal_address!=''
		) t1
		INNER JOIN t_withdrawal_address_insert t2
			ON t1.f_withdrawal_address = t2.f_withdrawal_address
		WHERE t1.rn = 1
		ON CONFLICT (f_validator_pubkey) DO UPDATE SET 
			f_pool_name = EXCLUDED.f_pool_name;
	`
)

func (p *PostgresDBService) ApplyWithdrawalAddressInsert() error {
	p.writerThreadsWG.Add(1)
	defer p.writerThreadsWG.Done()
	conn, err := p.psqlPool.Acquire(p.ctx)
	if err != nil {
		return errors.Wrap(err, "error acquiring connection")
	}
	defer conn.Release()

	_, err = conn.Exec(p.ctx, applyWithdrawalAddressInsertQuery)
	if err != nil {
		return errors.Wrap(err, "error applying withdrawal address insert")
	}
	return nil
}
