package db

import "github.com/pkg/errors"

const (
	applyDepositorsInsertQuery = `
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
				f_depositor,
				ROW_NUMBER() OVER (
					PARTITION BY f_validator_pubkey 
					ORDER BY f_block_num DESC
				) as rn
			FROM t_beacon_deposits
		) t1
		INNER JOIN t_depositors_insert t2 
			ON t1.f_depositor = t2.f_depositor
		WHERE t1.rn = 1
		ON CONFLICT (f_validator_pubkey) DO UPDATE SET 
			f_pool_name = EXCLUDED.f_pool_name;
	`
)

func (p *PostgresDBService) ApplyDepositorsInsert() error {
	p.writerThreadsWG.Add(1)
	defer p.writerThreadsWG.Done()
	conn, err := p.psqlPool.Acquire(p.ctx)
	if err != nil {
		return errors.Wrap(err, "error acquiring connection")
	}
	defer conn.Release()

	_, err = conn.Exec(p.ctx, applyDepositorsInsertQuery)
	if err != nil {
		return errors.Wrap(err, "error applying validators insert")
	}
	return nil
}
