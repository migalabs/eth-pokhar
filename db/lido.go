package db

import (
	"strings"
	"time"

	"github.com/google/uuid"
	pgx "github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
)

const (
	selectLidoOperatorsValidatorCount = `
	SELECT count(*)
	FROM t_lido
	GROUP BY f_operator_index
	ORDER BY f_operator_index ASC;
	`
)

// Obtain LidoOperatorsValidatorCount returns the number of validators in the Lido table for each operator
func (p *PostgresDBService) ObtainLidoOperatorsValidatorCount() ([]uint64, error) {
	conn, err := p.psqlPool.Acquire(p.ctx)
	if err != nil {
		return nil, errors.Wrap(err, "error acquiring database connection")
	}
	defer conn.Release()

	rows, err := conn.Query(p.ctx, selectLidoOperatorsValidatorCount)
	if err != nil {
		return nil, errors.Wrap(err, "error obtaining validator count from database")
	}
	defer rows.Close()

	var operatorsValidatorCount []uint64
	for rows.Next() {
		var count uint64
		err = rows.Scan(&count)
		if err != nil {
			return nil, errors.Wrap(err, "error scanning validator count from database")
		}
		operatorsValidatorCount = append(operatorsValidatorCount, count)
	}
	return operatorsValidatorCount, nil
}

// CopyLidoOperatorValidators copies the validators to the database for a given operator
func (p *PostgresDBService) CopyLidoOperatorValidators(operator string, operatorIndex uint64, rowSrc []string) int64 {
	if len(rowSrc) == 0 {
		return 0
	}
	p.writerThreadsWG.Add(1)
	defer p.writerThreadsWG.Done()
	startTime := time.Now()

	// Generate a random text to append to the table name
	randomText := uuid.New().String()

	// Create the temporary table name with the random text
	tempTableName := "temp_lido_" + strings.ReplaceAll(randomText, "-", "_")

	var validators [][]interface{}
	for _, row := range rowSrc {
		validators = append(validators, []interface{}{row, operator, operatorIndex})
	}

	// Acquire a database connection
	conn, err := p.psqlPool.Acquire(p.ctx)
	if err != nil {
		wlog.Fatalf("error acquiring database connection: %v", err)
	}
	defer conn.Release()

	// Create a temporary table with a unique constraint
	_, err = conn.Exec(p.ctx, `
		CREATE TEMP TABLE IF NOT EXISTS `+tempTableName+` (
			f_validator_pubkey text,
			f_operator text,
			f_operator_index integer
		);
	`)
	if err != nil {
		wlog.Fatalf("error creating temporary table: %v", err)
	}

	// Copy the data to the temporary table
	_, err = conn.CopyFrom(p.ctx, pgx.Identifier{tempTableName}, []string{"f_validator_pubkey", "f_operator", "f_operator_index"}, pgx.CopyFromRows(validators))
	if err != nil {
		wlog.Fatalf("error copying data to temporary table: %v", err)
	}

	// Insert the data from the temporary table to the main table
	count, err := conn.Exec(p.ctx, `
		INSERT INTO t_lido (f_validator_pubkey, f_operator, f_operator_index)
		SELECT f_validator_pubkey, f_operator, f_operator_index
		FROM `+tempTableName+`
		ON CONFLICT (f_validator_pubkey) DO NOTHING;
	`)
	if err != nil {
		wlog.Fatalf("error inserting data from temporary table to main table: %v", err)
	}

	// Drop the temporary table
	_, err = conn.Exec(p.ctx, `DROP TABLE `+tempTableName+`;`)
	if err != nil {
		wlog.Fatalf("error dropping temporary table: %v", err)
	}
	if count.RowsAffected() > 0 {
		wlog.Debugf("persisted %d rows in %f seconds", count.RowsAffected(), time.Since(startTime).Seconds())
	}

	return count.RowsAffected()
}
