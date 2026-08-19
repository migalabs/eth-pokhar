package db

import "github.com/pkg/errors"

const (
	// Legacy pins (f_dimension NULL) and operator pins both write the display
	// label: entities have always been the operator when known (Lido
	// validators show their operator, not "lido"). Custodian pins assert
	// custody only (see migration 000013 and ApplyDimensionColumns) and must
	// not touch f_pool_name.
	applyValidatorsInsertQuery = `
	INSERT INTO t_identified_validators (
		f_validator_pubkey,
		f_pool_name
	)
	SELECT
		f_validator_pubkey,
		f_pool_name
	FROM
		t_validators_insert
	WHERE f_dimension IS DISTINCT FROM 'custodian'
	ON CONFLICT (f_validator_pubkey) DO UPDATE SET f_pool_name = EXCLUDED.f_pool_name
	WHERE t_identified_validators.f_pool_name IS DISTINCT FROM EXCLUDED.f_pool_name;
	`

	// A custodian pin must not depend on some other phase creating the row:
	// without this, a pin declared before the validator is tracked is silently
	// dropped until an unrelated write creates the row (review finding on
	// PR #43). Creates missing rows with the pipeline default and never
	// touches existing ones; the dimension pass then applies the pin.
	ensureCustodianPinRowsQuery = `
	INSERT INTO t_identified_validators (
		f_validator_pubkey,
		f_pool_name
	)
	SELECT
		f_validator_pubkey,
		'solo_stakers'
	FROM
		t_validators_insert
	WHERE f_dimension = 'custodian'
	ON CONFLICT (f_validator_pubkey) DO NOTHING;
	`
)

func (p *PostgresDBService) ApplyValidatorsInsert() error {
	p.writerThreadsWG.Add(1)
	defer p.writerThreadsWG.Done()
	conn, err := p.psqlPool.Acquire(p.ctx)
	if err != nil {
		return errors.Wrap(err, "error acquiring connection")
	}
	defer conn.Release()

	_, err = conn.Exec(p.ctx, applyValidatorsInsertQuery)
	if err != nil {
		return errors.Wrap(err, "error applying validators insert")
	}

	_, err = conn.Exec(p.ctx, ensureCustodianPinRowsQuery)
	if err != nil {
		return errors.Wrap(err, "error ensuring custodian pin rows")
	}
	return nil
}
