package db

import "github.com/pkg/errors"

const (
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
	ON CONFLICT (f_validator_pubkey) DO UPDATE SET f_pool_name = EXCLUDED.f_pool_name;
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
