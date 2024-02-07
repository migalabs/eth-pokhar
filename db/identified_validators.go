package db

import "github.com/pkg/errors"

const (
	addNewValidatorsQuery = `
		INSERT INTO t_identified_validators (f_validator_pubkey, f_pool_name)
		SELECT DISTINCT F_VALIDATOR_PUBKEY, 'others'::text
		FROM T_BEACON_DEPOSITS
		WHERE F_VALIDATOR_PUBKEY != ''
		ON CONFLICT (f_validator_pubkey) DO NOTHING;
	`
	tuncateIdentifiedValidatorsQuery = `
		TRUNCATE TABLE t_identified_validators;
	`
)

func (p *PostgresDBService) AddNewValidators() error {
	p.writerThreadsWG.Add(1)
	defer p.writerThreadsWG.Done()
	conn, err := p.psqlPool.Acquire(p.ctx)
	if err != nil {
		return errors.Wrap(err, "error acquiring connection from pool")
	}
	defer conn.Release()

	_, err = conn.Exec(p.ctx, addNewValidatorsQuery)
	if err != nil {
		return errors.Wrap(err, "error adding new validators to database")
	}
	return nil
}

func (p *PostgresDBService) TruncateIdentifiedValidators() error {
	p.writerThreadsWG.Add(1)
	defer p.writerThreadsWG.Done()
	conn, err := p.psqlPool.Acquire(p.ctx)
	if err != nil {
		return errors.Wrap(err, "error acquiring connection from pool")
	}
	defer conn.Release()

	_, err = conn.Exec(p.ctx, tuncateIdentifiedValidatorsQuery)
	if err != nil {
		return errors.Wrap(err, "error truncating identified validators")
	}
	return nil
}
