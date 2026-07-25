package db

import (
	"strings"

	"github.com/pkg/errors"
)

const (
	addNewValidatorsQuery = `
		INSERT INTO t_identified_validators (f_validator_pubkey, f_pool_name)
		SELECT v.f_validator_pubkey, 'solo_stakers'::text
		FROM t_validator_last_deposit v
		LEFT JOIN t_identified_validators t
			ON t.f_validator_pubkey = v.f_validator_pubkey
		WHERE t.f_validator_pubkey IS NULL
			AND v.f_validator_pubkey != ''
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
					%s, 
					CONCAT('whale_0x', LEFT(%s, 4)) AS whale_id
				FROM (
					SELECT
						COUNT(*) AS COUNT,
						%s
					FROM
						(
							SELECT DISTINCT
								F_VALIDATOR_PUBKEY,
								%s
							FROM
								T_BEACON_DEPOSITS
							WHERE
								F_VALIDATOR_PUBKEY != ''
						) aux
					GROUP BY
						%s
				) AS count_subquery
				WHERE 
					count >= $1
			) AS subquery
			ON 
				t1.%s = subquery.%s
			WHERE 
				F_VALIDATOR_PUBKEY != ''
		) AS subquery2
		WHERE 
			t_identified_validators.f_validator_pubkey = subquery2.f_validator_pubkey
			AND t_identified_validators.f_pool_name IS DISTINCT FROM subquery2.whale_id
			AND (
				t_identified_validators.f_pool_name IS NULL
				OR t_identified_validators.f_pool_name = ''
				OR t_identified_validators.f_pool_name = 'solo_stakers'
				OR t_identified_validators.f_pool_name LIKE 'whale\_%'
			);
		`

	depositorsColumnName        = "f_depositor"
	withdrawalAddressColumnName = "f_withdrawal_address"
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

	withdrawalAddressQuery := strings.ReplaceAll(identifyWhalesQuery, "%s", withdrawalAddressColumnName)

	_, err = conn.Exec(p.ctx, withdrawalAddressQuery, threshold)
	if err != nil {
		return errors.Wrap(err, "error identifying withdrawal address whales")
	}

	depositorsQuery := strings.ReplaceAll(identifyWhalesQuery, "%s", depositorsColumnName)
	_, err = conn.Exec(p.ctx, depositorsQuery, threshold)
	if err != nil {
		return errors.Wrap(err, "error identifying depositor whales")
	}
	return nil
}
