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
	//   f_operator:  declared pin (f_dimension = 'operator') > canonical Lido
	//                operator ($3..$5, by module and registry index) > Lido
	//                registry tag > Rocket Pool registry > depositor mapping >
	//                coinbase detection (fallback: detected validators are
	//                operated by the exchange unless a better signal exists;
	//                the depositor mapping beats it because e.g. blockdaemon
	//                operates validators whose custody is coinbase's).
	//   f_custodian: declared pin (f_dimension = 'custodian') > Lido module
	//                ($1/$2, by module) > withdrawal address mapping > coinbase
	//                detection (same fallback reasoning on the custody side).
	//
	// Lido keys get one custodian per staking module (lido_cmv1, lido_sdvt,
	// ...), the custody entities the dashboard groups them by, and their
	// company as operator (figment, not figment_lido). Both come from the key
	// registry, a stronger fact than any address mapping: should the shared
	// Lido withdrawal vault ever be mapped, the modules must not collapse into
	// one custodian. The canonical operator is resolved here, at the end of the
	// run, rather than stored in t_lido, so a mapping change reaches every key
	// on the next run instead of waiting for the operator's keys to change.
	//
	// Declared pins (migration 000013) are curated per-pubkey facts, so they
	// rank first in their dimension and double as the manual override; the
	// depositor mapping in particular mislabels delegated operation (the
	// depositor of e.g. a liquid-restaking validator is the platform, not
	// whoever runs it), which is exactly what an operator pin corrects.
	//
	// Heuristic labels (whales, solo_stakers) and legacy pins (f_dimension
	// IS NULL) only exist in f_pool_name: their dimension is undeclared, so
	// they never feed the pure columns. NULL here means "no declared signal",
	// not "unknown entity".
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
				cur.f_validator_pubkey,
				COALESCE(
					pinop.f_pool_name,
					lo.tag,
					l.f_operator,
					CASE WHEN r.f_validator_pubkey IS NOT NULL THEN 'rocketpool' END,
					md.f_pool_name,
					CASE WHEN cb.f_validator_pubkey IS NOT NULL THEN 'coinbase' END
				) AS op,
				COALESCE(
					pincust.f_pool_name,
					lc.tag,
					mw.f_pool_name,
					CASE WHEN cb.f_validator_pubkey IS NOT NULL THEN 'coinbase' END
				) AS cust
			FROM t_identified_validators cur
			LEFT JOIN t_validator_last_deposit v
				ON v.f_validator_pubkey = cur.f_validator_pubkey
			LEFT JOIN t_coinbase_detected cb
				ON cb.f_validator_pubkey = cur.f_validator_pubkey
			LEFT JOIN t_validators_insert pinop
				ON pinop.f_validator_pubkey = cur.f_validator_pubkey
				AND pinop.f_dimension = 'operator'
			LEFT JOIN t_validators_insert pincust
				ON pincust.f_validator_pubkey = cur.f_validator_pubkey
				AND pincust.f_dimension = 'custodian'
			LEFT JOIN t_lido l
				ON l.f_validator_pubkey = cur.f_validator_pubkey
			LEFT JOIN unnest($1::text[], $2::text[]) AS lc(protocol, tag)
				ON lc.protocol = l.f_protocol
			LEFT JOIN unnest($3::text[], $4::int[], $5::text[]) AS lo(protocol, idx, tag)
				ON lo.protocol = l.f_protocol
				AND lo.idx = l.f_operator_index
			LEFT JOIN t_rocketpool r
				ON r.f_validator_pubkey = cur.f_validator_pubkey
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

// LidoDimensions carries the Lido-specific dimension sources as parallel
// slices: CustodyProtocols[k] keys get CustodyTags[k] as custodian, and the
// key of OperatorProtocols[k] / OperatorIndexes[k] gets OperatorTags[k] as
// operator.
type LidoDimensions struct {
	CustodyProtocols  []string
	CustodyTags       []string
	OperatorProtocols []string
	OperatorIndexes   []int32
	OperatorTags      []string
}

func (d LidoDimensions) validate() error {
	if len(d.CustodyProtocols) != len(d.CustodyTags) {
		return errors.Errorf("custody protocols and tags differ in length: %d vs %d", len(d.CustodyProtocols), len(d.CustodyTags))
	}
	if len(d.OperatorProtocols) != len(d.OperatorIndexes) || len(d.OperatorProtocols) != len(d.OperatorTags) {
		return errors.Errorf("operator protocols, indexes and tags differ in length: %d, %d, %d",
			len(d.OperatorProtocols), len(d.OperatorIndexes), len(d.OperatorTags))
	}
	return nil
}

// ApplyDimensionColumns recomputes f_operator and f_custodian for every
// validator whose derived values changed. Runs at the end of the identify
// pipeline, after all phases and mappings have settled, and is deterministic
// from the mapping tables: full rebuilds and incremental runs converge to the
// same values. Returns the number of rows written.
func (p *PostgresDBService) ApplyDimensionColumns(lido LidoDimensions) (int64, error) {
	if err := lido.validate(); err != nil {
		return 0, err
	}
	p.writerThreadsWG.Add(1)
	defer p.writerThreadsWG.Done()
	startTime := time.Now()

	conn, err := p.psqlPool.Acquire(p.ctx)
	if err != nil {
		return 0, errors.Wrap(err, "error acquiring connection")
	}
	defer conn.Release()

	cmd, err := conn.Exec(p.ctx, applyDimensionColumnsQuery,
		lido.CustodyProtocols, lido.CustodyTags,
		lido.OperatorProtocols, lido.OperatorIndexes, lido.OperatorTags)
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
