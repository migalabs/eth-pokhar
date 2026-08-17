package db

import "github.com/pkg/errors"

const (
	// Rows with a declared dimension assert f_operator or f_custodian only
	// (see migration 000013 and ApplyDimensionColumns); they must not touch
	// the legacy f_pool_name label.
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
	WHERE f_dimension IS NULL
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
