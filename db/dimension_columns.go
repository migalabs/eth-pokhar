package db

import (
	"time"

	"github.com/pkg/errors"
)

const (
	// One declarative pass computes both pure dimensions from their sources,
	// with no precedence BETWEEN dimensions (that arbitration is exactly what
	// the legacy f_pool_name does and what these columns exist to avoid).
	// Within each dimension, stronger evidence wins:
	//
	//   f_operator:  Lido registry > Rocket Pool registry > depositor mapping
	//                > coinbase detection (fallback: detected validators are
	//                operated by the exchange unless a better signal exists;
	//                the depositor mapping beats it because e.g. blockdaemon
	//                operates validators whose custody is coinbase's).
	//   f_custodian: withdrawal address mapping > coinbase detection (same
	//                fallback reasoning on the custody side).
	//
	// Heuristic labels (whales, solo_stakers) and manual pins only exist in
	// f_pool_name: their dimension is undeclared, so they never feed the pure
	// columns. NULL here means "no declared signal", not "unknown entity".
	//
	// The change guard keeps steady-state runs down to the real daily delta;
	// the first run after the migration rewrites every row once (write volume
	// comparable to a weekly full rebuild, which production absorbs weekly).
	applyDimensionColumnsQuery = `
		UPDATE t_identified_validators t
		SET f_operator = src.op,
		    f_custodian = src.cust
		FROM (
			SELECT
				v.f_validator_pubkey,
				COALESCE(
					l.f_operator,
					CASE WHEN r.f_validator_pubkey IS NOT NULL THEN 'rocketpool' END,
					md.f_pool_name,
					CASE WHEN cur.f_pool_name = 'coinbase' THEN 'coinbase' END
				) AS op,
				COALESCE(
					mw.f_pool_name,
					CASE WHEN cur.f_pool_name = 'coinbase' THEN 'coinbase' END
				) AS cust
			FROM t_validator_last_deposit v
			JOIN t_identified_validators cur
				ON cur.f_validator_pubkey = v.f_validator_pubkey
			LEFT JOIN t_lido l
				ON l.f_validator_pubkey = v.f_validator_pubkey
			LEFT JOIN t_rocketpool r
				ON r.f_validator_pubkey = v.f_validator_pubkey
			LEFT JOIN t_depositors_insert md
				ON md.f_depositor = v.f_depositor
			LEFT JOIN t_withdrawal_address_insert mw
				ON mw.f_withdrawal_address = v.f_withdrawal_address
				AND v.f_withdrawal_address != ''
		) src
		WHERE t.f_validator_pubkey = src.f_validator_pubkey
		  AND (t.f_operator IS DISTINCT FROM src.op
		       OR t.f_custodian IS DISTINCT FROM src.cust);
	`

	analyzeAfterDimensionRows = 100000
)

// ApplyDimensionColumns recomputes f_operator and f_custodian for every
// validator whose derived values changed. Runs at the end of the identify
// pipeline, after all phases and mappings have settled, and is deterministic
// from the mapping tables: full rebuilds and incremental runs converge to the
// same values. Returns the number of rows written.
func (p *PostgresDBService) ApplyDimensionColumns() (int64, error) {
	p.writerThreadsWG.Add(1)
	defer p.writerThreadsWG.Done()
	startTime := time.Now()

	conn, err := p.psqlPool.Acquire(p.ctx)
	if err != nil {
		return 0, errors.Wrap(err, "error acquiring connection")
	}
	defer conn.Release()

	cmd, err := conn.Exec(p.ctx, applyDimensionColumnsQuery)
	if err != nil {
		return 0, errors.Wrap(err, "error applying dimension columns")
	}
	rows := cmd.RowsAffected()

	// A large write (first fill, weekly rebuild) skews planner stats; refresh
	// them so the hourly queries keep good plans.
	if rows > analyzeAfterDimensionRows {
		if _, err := conn.Exec(p.ctx, `ANALYZE t_identified_validators;`); err != nil {
			return rows, errors.Wrap(err, "error analyzing after dimension update")
		}
	}

	wlog.Debugf("dimension columns updated: %d rows in %.2fs", rows, time.Since(startTime).Seconds())
	return rows, nil
}
