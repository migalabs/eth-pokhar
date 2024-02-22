package db

import "github.com/pkg/errors"

const (
	addNewValidatorsQuery = `
		INSERT INTO t_identified_validators (f_validator_pubkey, f_pool_name)
		SELECT DISTINCT F_VALIDATOR_PUBKEY, 'solo_stakers'::text
		FROM T_BEACON_DEPOSITS
		WHERE F_VALIDATOR_PUBKEY != ''
		ON CONFLICT (f_validator_pubkey) DO NOTHING;
	`
	truncateIdentifiedValidatorsQuery = `
		TRUNCATE TABLE t_identified_validators;
	`

	identifyWhalesQuery = `
		UPDATE t_identified_validators 
		SET f_pool_name = subquery2.whale_id
		FROM (
			SELECT DISTINCT 
				t1.f_validator_pubkey, 
				whale_id	
			FROM 
				t_beacon_deposits t1
			RIGHT JOIN (
				SELECT 
					f_depositor, 
					CONCAT('whale_0x', LEFT(f_depositor, 4)) AS whale_id
				FROM (
					SELECT
						COUNT(*) AS COUNT,
						F_DEPOSITOR
					FROM
						(
							SELECT DISTINCT
								F_VALIDATOR_PUBKEY,
								F_DEPOSITOR
							FROM
								T_BEACON_DEPOSITS
							WHERE
								F_VALIDATOR_PUBKEY != ''
						) aux
					GROUP BY
						F_DEPOSITOR
				) AS count_subquery
				WHERE 
					count >= $1
			) AS subquery
			ON 
				t1.f_depositor = subquery.f_depositor
			WHERE 
				F_VALIDATOR_PUBKEY != ''
		) AS subquery2
		WHERE 
			t_identified_validators.f_validator_pubkey = subquery2.f_validator_pubkey;
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

	_, err = conn.Exec(p.ctx, truncateIdentifiedValidatorsQuery)
	if err != nil {
		return errors.Wrap(err, "error truncating identified validators")
	}
	return nil
}

func (p *PostgresDBService) IdentifyWhales(threshold int) error {
	conn, err := p.psqlPool.Acquire(p.ctx)
	if err != nil {
		return errors.Wrap(err, "error acquiring database connection")
	}
	defer conn.Release()

	_, err = conn.Exec(p.ctx, identifyWhalesQuery, threshold)
	if err != nil {
		return errors.Wrap(err, "error identifying whales")
	}
	return nil
}
