package db

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	pgx "github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
)

const (
	selectLidoOperatorsValidatorCount = `
	WITH operators_sequence AS (
		SELECT generate_series(0, (SELECT MAX(f_operator_index) FROM t_lido WHERE f_protocol=$1)) AS sequence_number
	)
	SELECT 
		COUNT(l.f_operator_index) as count
	FROM operators_sequence os
	LEFT JOIN t_lido l ON os.sequence_number = l.f_operator_index 
		AND l.f_protocol = $1
	GROUP BY os.sequence_number
	ORDER BY os.sequence_number ASC;
	`

	// The display label is the operator tag, except for protocols listed in
	// ($1, $2) as (protocol, entity tag) pairs: those modules show as a single
	// entity and keep the per-operator tag only in t_lido.f_operator, which
	// feeds the f_operator dimension.
	identifyLidoValidators = `
	UPDATE t_identified_validators t
	SET f_pool_name = src.f_pool_name
	FROM (
		SELECT l.f_validator_pubkey, COALESCE(m.tag, l.f_operator) AS f_pool_name
		FROM t_lido l
		LEFT JOIN unnest($1::text[], $2::text[]) AS m(protocol, tag)
			ON m.protocol = l.f_protocol
	) src
	WHERE t.f_validator_pubkey = src.f_validator_pubkey
		AND t.f_pool_name IS DISTINCT FROM src.f_pool_name;
	`
	LidoProtocolCurated = "curated"
	LidoProtocolCSM     = "csm"
	LidoProtocolSDVT    = "sdvt"
	// Curated Module v2 (StakingRouter module 4).
	LidoProtocolCMv2 = "cmv2"
	// Community Staking 0x02 (StakingRouter module 5).
	LidoProtocolCSM0x02 = "csm0x02"
)

// IdentifyLidoValidators labels every validator found in t_lido. entityProtocols
// and entityTags are parallel: validators of entityProtocols[k] get
// entityTags[k] as their label, the rest get their operator tag.
func (p *PostgresDBService) IdentifyLidoValidators(entityProtocols []string, entityTags []string) error {
	if len(entityProtocols) != len(entityTags) {
		return errors.Errorf("entity protocols and tags differ in length: %d vs %d", len(entityProtocols), len(entityTags))
	}
	p.writerThreadsWG.Add(1)
	defer p.writerThreadsWG.Done()

	conn, err := p.psqlPool.Acquire(p.ctx)
	if err != nil {
		return errors.Wrap(err, "error acquiring database connection")
	}
	defer conn.Release()

	// Exec, not Query: the rows of a Query were never closed, so a failing
	// UPDATE went unreported.
	_, err = conn.Exec(p.ctx, identifyLidoValidators, entityProtocols, entityTags)
	if err != nil {
		return errors.Wrap(err, "error identifying lido validators")
	}
	return nil
}

// Obtain LidoOperatorsValidatorCount returns the number of validators in the Lido table for each operator
func (p *PostgresDBService) ObtainLidoOperatorsValidatorCount(protocol string) ([]uint64, error) {
	conn, err := p.psqlPool.Acquire(p.ctx)
	if err != nil {
		return nil, errors.Wrap(err, "error acquiring database connection")
	}
	defer conn.Release()

	rows, err := conn.Query(p.ctx, selectLidoOperatorsValidatorCount, protocol)
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

// ReconcileLidoOperatorValidators makes t_lido reflect exactly the on-chain key
// set for a given operator/protocol. It removes fossil keys (in DB but no longer
// on-chain) and inserts missing keys; the UPSERT also refreshes f_operator,
// f_operator_index and f_protocol for any row whose values diverged from
// on-chain. All work happens inside a single transaction so consumers never
// observe a partially-rebuilt state for the operator.
func (p *PostgresDBService) ReconcileLidoOperatorValidators(operator string, operatorIndex uint64, onChainKeys []string, protocol string) (int64, int64, error) {
	p.writerThreadsWG.Add(1)
	defer p.writerThreadsWG.Done()
	startTime := time.Now()

	conn, err := p.psqlPool.Acquire(p.ctx)
	if err != nil {
		return 0, 0, errors.Wrap(err, "error acquiring database connection")
	}
	defer conn.Release()

	tx, err := conn.Begin(p.ctx)
	if err != nil {
		return 0, 0, errors.Wrap(err, "error beginning transaction")
	}
	defer tx.Rollback(p.ctx)

	tempTableName := "tmp_lido_reconcile_" + strings.ReplaceAll(uuid.New().String(), "-", "_")
	_, err = tx.Exec(p.ctx, `
		CREATE TEMP TABLE `+tempTableName+` (
			f_validator_pubkey text PRIMARY KEY
		) ON COMMIT DROP;
	`)
	if err != nil {
		return 0, 0, errors.Wrap(err, "error creating temp table")
	}

	if len(onChainKeys) > 0 {
		rows := make([][]interface{}, len(onChainKeys))
		for i, k := range onChainKeys {
			rows[i] = []interface{}{k}
		}
		_, err = tx.CopyFrom(p.ctx, pgx.Identifier{tempTableName}, []string{"f_validator_pubkey"}, pgx.CopyFromRows(rows))
		if err != nil {
			return 0, 0, errors.Wrap(err, "error copying on-chain keys into temp table")
		}
	}

	delCmd, err := tx.Exec(p.ctx, `
		DELETE FROM t_lido
		WHERE f_operator_index = $1
		  AND f_protocol = $2
		  AND f_validator_pubkey NOT IN (SELECT f_validator_pubkey FROM `+tempTableName+`);
	`, operatorIndex, protocol)
	if err != nil {
		return 0, 0, errors.Wrap(err, "error deleting fossil keys")
	}
	removed := delCmd.RowsAffected()

	insCmd, err := tx.Exec(p.ctx, `
		INSERT INTO t_lido (f_validator_pubkey, f_operator, f_operator_index, f_protocol)
		SELECT t.f_validator_pubkey, $1, $2, $3
		FROM `+tempTableName+` t
		ON CONFLICT (f_validator_pubkey) DO UPDATE
			SET f_operator = EXCLUDED.f_operator,
			    f_operator_index = EXCLUDED.f_operator_index,
			    f_protocol = EXCLUDED.f_protocol
		WHERE t_lido.f_operator       IS DISTINCT FROM EXCLUDED.f_operator
		   OR t_lido.f_operator_index IS DISTINCT FROM EXCLUDED.f_operator_index
		   OR t_lido.f_protocol       IS DISTINCT FROM EXCLUDED.f_protocol;
	`, operator, operatorIndex, protocol)
	if err != nil {
		return 0, 0, errors.Wrap(err, "error inserting fresh keys")
	}
	upserted := insCmd.RowsAffected()

	if err := tx.Commit(p.ctx); err != nil {
		return 0, 0, errors.Wrap(err, "error committing reconcile transaction")
	}

	if upserted > 0 || removed > 0 {
		wlog.Debugf("reconciled operator %v (proto=%v): upserted=%d removed=%d in %.2fs",
			operator, protocol, upserted, removed, time.Since(startTime).Seconds())
	}
	return upserted, removed, nil
}

// LidoOperatorCountKey builds the lookup key used by the map returned from
// ObtainLidoOperatorKeyCounts.
func LidoOperatorCountKey(protocol string, operatorIndex uint64) string {
	return fmt.Sprintf("%s/%d", protocol, operatorIndex)
}

// ObtainLidoOperatorKeyCounts returns how many keys t_lido currently holds per
// (protocol, operator index). Used as the checkpoint to decide whether an
// operator needs no work, an incremental key fetch or a full reconcile.
func (p *PostgresDBService) ObtainLidoOperatorKeyCounts() (map[string]uint64, error) {
	conn, err := p.psqlPool.Acquire(p.ctx)
	if err != nil {
		return nil, errors.Wrap(err, "error acquiring database connection")
	}
	defer conn.Release()

	rows, err := conn.Query(p.ctx, `
		SELECT f_protocol, f_operator_index, count(*)
		FROM t_lido
		GROUP BY f_protocol, f_operator_index;
	`)
	if err != nil {
		return nil, errors.Wrap(err, "error obtaining lido operator key counts")
	}
	defer rows.Close()

	counts := make(map[string]uint64)
	for rows.Next() {
		var protocol string
		var operatorIndex int64
		var count int64
		if err := rows.Scan(&protocol, &operatorIndex, &count); err != nil {
			return nil, errors.Wrap(err, "error scanning lido operator key count")
		}
		counts[LidoOperatorCountKey(protocol, uint64(operatorIndex))] = uint64(count)
	}
	return counts, nil
}

// AppendLidoOperatorValidators upserts a batch of keys for an operator WITHOUT
// removing anything. Meant for the incremental path, where only the keys added
// on-chain since the last run are fetched: passing such a partial set to
// ReconcileLidoOperatorValidators would delete every previously stored key of
// the operator, so removals are handled exclusively by the full reconcile.
func (p *PostgresDBService) AppendLidoOperatorValidators(operator string, operatorIndex uint64, newKeys []string, protocol string) (int64, error) {
	if len(newKeys) == 0 {
		return 0, nil
	}
	p.writerThreadsWG.Add(1)
	defer p.writerThreadsWG.Done()
	startTime := time.Now()

	conn, err := p.psqlPool.Acquire(p.ctx)
	if err != nil {
		return 0, errors.Wrap(err, "error acquiring database connection")
	}
	defer conn.Release()

	tx, err := conn.Begin(p.ctx)
	if err != nil {
		return 0, errors.Wrap(err, "error beginning transaction")
	}
	defer tx.Rollback(p.ctx)

	tempTableName := "tmp_lido_append_" + strings.ReplaceAll(uuid.New().String(), "-", "_")
	_, err = tx.Exec(p.ctx, `
		CREATE TEMP TABLE `+tempTableName+` (
			f_validator_pubkey text PRIMARY KEY
		) ON COMMIT DROP;
	`)
	if err != nil {
		return 0, errors.Wrap(err, "error creating temp table")
	}

	rows := make([][]interface{}, len(newKeys))
	for i, k := range newKeys {
		rows[i] = []interface{}{k}
	}
	_, err = tx.CopyFrom(p.ctx, pgx.Identifier{tempTableName}, []string{"f_validator_pubkey"}, pgx.CopyFromRows(rows))
	if err != nil {
		return 0, errors.Wrap(err, "error copying new keys into temp table")
	}

	insCmd, err := tx.Exec(p.ctx, `
		INSERT INTO t_lido (f_validator_pubkey, f_operator, f_operator_index, f_protocol)
		SELECT t.f_validator_pubkey, $1, $2, $3
		FROM `+tempTableName+` t
		ON CONFLICT (f_validator_pubkey) DO UPDATE
			SET f_operator = EXCLUDED.f_operator,
			    f_operator_index = EXCLUDED.f_operator_index,
			    f_protocol = EXCLUDED.f_protocol
		WHERE t_lido.f_operator       IS DISTINCT FROM EXCLUDED.f_operator
		   OR t_lido.f_operator_index IS DISTINCT FROM EXCLUDED.f_operator_index
		   OR t_lido.f_protocol       IS DISTINCT FROM EXCLUDED.f_protocol;
	`, operator, operatorIndex, protocol)
	if err != nil {
		return 0, errors.Wrap(err, "error appending new keys")
	}
	upserted := insCmd.RowsAffected()

	if err := tx.Commit(p.ctx); err != nil {
		return 0, errors.Wrap(err, "error committing append transaction")
	}

	if upserted > 0 {
		wlog.Debugf("appended %d keys for operator %v (proto=%v) in %.2fs",
			upserted, operator, protocol, time.Since(startTime).Seconds())
	}
	return upserted, nil
}

// CopyLidoOperatorValidators copies the validators to the database for a given operator
func (p *PostgresDBService) CopyLidoOperatorValidators(operator string, operatorIndex uint64, rowSrc []string, protocol string) int64 {
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
		validators = append(validators, []interface{}{row, operator, operatorIndex, protocol})
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
			f_operator_index integer,
			f_protocol text
		);
	`)
	if err != nil {
		wlog.Fatalf("error creating temporary table: %v", err)
	}

	// Copy the data to the temporary table
	_, err = conn.CopyFrom(p.ctx, pgx.Identifier{tempTableName}, []string{"f_validator_pubkey", "f_operator", "f_operator_index", "f_protocol"}, pgx.CopyFromRows(validators))
	if err != nil {
		wlog.Fatalf("error copying data to temporary table: %v", err)
	}

	// Insert the data from the temporary table to the main table
	count, err := conn.Exec(p.ctx, `
		INSERT INTO t_lido (f_validator_pubkey, f_operator, f_operator_index, f_protocol)
		SELECT f_validator_pubkey, f_operator, f_operator_index, f_protocol
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
