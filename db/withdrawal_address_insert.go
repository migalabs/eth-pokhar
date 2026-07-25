package db

import "github.com/pkg/errors"

const (
	applyWithdrawalAddressInsertQuery = `
		INSERT INTO t_identified_validators (
			f_validator_pubkey,
			f_pool_name
		)
		SELECT
			v.f_validator_pubkey,
			m.f_pool_name
		FROM t_validator_last_deposit v
		INNER JOIN t_withdrawal_address_insert m
			ON v.f_withdrawal_address = m.f_withdrawal_address
		LEFT JOIN t_identified_validators t
			ON t.f_validator_pubkey = v.f_validator_pubkey
		WHERE v.f_withdrawal_address != ''
			AND (t.f_validator_pubkey IS NULL
				OR t.f_pool_name IS DISTINCT FROM m.f_pool_name)
			-- Skip validators that a later phase claims (see issue #28): it
			-- would overwrite the tag again in the same run.
			AND t.f_pool_name IS DISTINCT FROM 'coinbase'
			AND NOT EXISTS (SELECT 1 FROM t_rocketpool r
				WHERE r.f_validator_pubkey = v.f_validator_pubkey)
			AND NOT EXISTS (SELECT 1 FROM t_lido l
				WHERE l.f_validator_pubkey = v.f_validator_pubkey)
			AND NOT EXISTS (SELECT 1 FROM t_validators_insert vi
				WHERE vi.f_validator_pubkey = v.f_validator_pubkey)
		ON CONFLICT (f_validator_pubkey) DO UPDATE SET
			f_pool_name = EXCLUDED.f_pool_name
		WHERE t_identified_validators.f_pool_name IS DISTINCT FROM EXCLUDED.f_pool_name;
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
