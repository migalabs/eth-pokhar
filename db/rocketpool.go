package db

import (
	"strings"
	"time"

	"github.com/google/uuid"
	pgx "github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
)

const (
	selectValidatorCount = `
	SELECT count(*)
	FROM t_rocketpool;
	`
	identifyRocketpoolValidators = `
	UPDATE t_identified_validators 
	SET f_pool_name = 'rocketpool'
	FROM t_rocketpool
	WHERE t_identified_validators.f_validator_pubkey = t_rocketpool.f_validator_pubkey;
	`
)

// IdentifyRocketpoolValidators identifies the rocketpool validators and adds them to the identified validators table
func (p *PostgresDBService) IdentifyRocketpoolValidators() error {
	conn, err := p.psqlPool.Acquire(p.ctx)
	if err != nil {
		return errors.Wrap(err, "error acquiring database connection")
	}
	defer conn.Release()

	_, err = conn.Query(p.ctx, identifyRocketpoolValidators)
	if err != nil {
		return errors.Wrap(err, "error identifying rocketpool validators")
	}
	return nil
}

// ObtainRocketpoolValidatorCount returns the number of validators in the database
func (p *PostgresDBService) ObtainRocketpoolValidatorCount() (int, error) {
	conn, err := p.psqlPool.Acquire(p.ctx)
	if err != nil {
		return 0, errors.Wrap(err, "error acquiring database connection")
	}
	defer conn.Release()

	var count int
	err = conn.QueryRow(p.ctx, selectValidatorCount).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "error obtaining validator count from database")
	}
	return count, nil
}

// CopyRocketpoolValidators copies the validators to the database
func (p *PostgresDBService) CopyRocketpoolValidators(rowSrc []string) int64 {
	if len(rowSrc) == 0 {
		return 0
	}
	p.writerThreadsWG.Add(1)
	defer p.writerThreadsWG.Done()
	startTime := time.Now()

	// Generate a random text to append to the table name
	randomText := uuid.New().String()

	// Create the temporary table name with the random text
	tempTableName := "temp_rocketpool_" + strings.ReplaceAll(randomText, "-", "_")

	var validators [][]interface{}
	for _, row := range rowSrc {
		validators = append(validators, []interface{}{row})
	}

	// Acquire a database connection
	conn, err := p.psqlPool.Acquire(p.ctx)
	if err != nil {
		return 0
	}
	defer conn.Release()

	// Create a temporary table with a unique constraint
	_, err = conn.Exec(p.ctx, `
		CREATE TEMP TABLE IF NOT EXISTS `+tempTableName+` (
			f_validator_pubkey text
		);
	`)
	if err != nil {
		return 0
	}

	// Copy the data to the temporary table
	_, err = conn.CopyFrom(p.ctx, pgx.Identifier{tempTableName}, []string{"f_validator_pubkey"}, pgx.CopyFromRows(validators))
	if err != nil {
		return 0
	}

	// Insert the data from the temporary table to the main table
	count, err := conn.Exec(p.ctx, `
		INSERT INTO t_rocketpool (f_validator_pubkey)
		SELECT f_validator_pubkey
		FROM `+tempTableName+`
		ON CONFLICT (f_validator_pubkey) DO NOTHING;
	`)
	if err != nil {
		return 0
	}

	// Drop the temporary table
	_, err = conn.Exec(p.ctx, `DROP TABLE `+tempTableName+`;`)
	if err != nil {
		return 0
	}
	if count.RowsAffected() > 0 {
		wlog.Debugf("persisted %d rows in %f seconds", count.RowsAffected(), time.Since(startTime).Seconds())
	}

	return count.RowsAffected()
}
