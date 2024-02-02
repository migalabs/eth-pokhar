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
)

func (p *PostgresDBService) AddNewValidators() (int64, error) {
	count, err := p.psqlPool.Query(p.ctx, addNewValidatorsQuery)
	if err != nil {
		return 0, errors.Wrap(err, "error adding new validators to database")
	}
	return count.CommandTag().RowsAffected(), nil
}
