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
)

func (p *PostgresDBService) ApplyValidatorsInsert() error {
	p.writerThreadsWG.Add(1)
	defer p.writerThreadsWG.Done()
	conn, err := p.psqlPool.Acquire(p.ctx)
	if err != nil {
		return errors.Wrap(err, "error acquiring connection")
	}
	defer conn.Release()

	_, err = conn.Query(p.ctx, applyValidatorsInsertQuery)
	if err != nil {
		return errors.Wrap(err, "error applying validators insert")
	}
	return nil
}
