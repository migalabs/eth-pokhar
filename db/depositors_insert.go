package db

import "github.com/pkg/errors"

const (
	applyDepositorsInsertQuery = `
		INSERT INTO t_identified_validators (
			f_validator_pubkey,
			f_pool_name
		)
		SELECT
			v.f_validator_pubkey,
			m.f_pool_name
		FROM t_validator_last_deposit v
		INNER JOIN t_depositors_insert m
			ON v.f_depositor = m.f_depositor
		LEFT JOIN t_identified_validators t
			ON t.f_validator_pubkey = v.f_validator_pubkey
		WHERE (t.f_validator_pubkey IS NULL
			OR t.f_pool_name IS DISTINCT FROM m.f_pool_name)
			-- Skip validators that a later phase claims: it would overwrite the
			-- tag again in the same run, so writing it here is pure churn that
			-- the phase order (issue #28) then has to undo.
			AND t.f_pool_name IS DISTINCT FROM 'coinbase'
			AND NOT (v.f_withdrawal_address != '' AND EXISTS (
				SELECT 1 FROM t_withdrawal_address_insert w
				WHERE w.f_withdrawal_address = v.f_withdrawal_address))
			AND NOT EXISTS (SELECT 1 FROM t_rocketpool r
				WHERE r.f_validator_pubkey = v.f_validator_pubkey)
			AND NOT EXISTS (SELECT 1 FROM t_lido l
				WHERE l.f_validator_pubkey = v.f_validator_pubkey)
			-- Custodian pins (migration 000013) never write f_pool_name, so
			-- only pins that claim the display label suppress this phase.
			AND NOT EXISTS (SELECT 1 FROM t_validators_insert vi
				WHERE vi.f_validator_pubkey = v.f_validator_pubkey
				AND vi.f_dimension IS DISTINCT FROM 'custodian')
		ON CONFLICT (f_validator_pubkey) DO UPDATE SET
			f_pool_name = EXCLUDED.f_pool_name
		WHERE t_identified_validators.f_pool_name IS DISTINCT FROM EXCLUDED.f_pool_name;
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
