package db

import (
	"fmt"
	"time"

	pgx "github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
)

const (
	// Both statements deduplicate the batch with DISTINCT ON before upserting:
	// a batch can carry several deposits for the same validator and Postgres
	// rejects ON CONFLICT DO UPDATE statements that touch a row twice. The
	// <= guards keep the upserts idempotent when the downloader re-scans the
	// last blocks and replays deposits already applied.
	upsertLastDepositorQuery = `
		INSERT INTO t_validator_last_deposit (
			f_validator_pubkey,
			f_depositor,
			f_depositor_block,
			f_withdrawal_address,
			f_withdrawal_block
		)
		SELECT DISTINCT ON (f_validator_pubkey)
			f_validator_pubkey,
			f_depositor,
			f_block_num,
			'',
			0
		FROM %s
		ORDER BY f_validator_pubkey, f_block_num DESC
		ON CONFLICT (f_validator_pubkey) DO UPDATE SET
			f_depositor = EXCLUDED.f_depositor,
			f_depositor_block = EXCLUDED.f_depositor_block
		WHERE t_validator_last_deposit.f_depositor_block <= EXCLUDED.f_depositor_block;
	`

	upsertLastWithdrawalQuery = `
		INSERT INTO t_validator_last_deposit (
			f_validator_pubkey,
			f_depositor,
			f_depositor_block,
			f_withdrawal_address,
			f_withdrawal_block
		)
		SELECT DISTINCT ON (f_validator_pubkey)
			f_validator_pubkey,
			f_depositor,
			f_block_num,
			f_withdrawal_address,
			f_block_num
		FROM %s
		WHERE f_withdrawal_address != ''
		ORDER BY f_validator_pubkey, f_block_num DESC
		ON CONFLICT (f_validator_pubkey) DO UPDATE SET
			f_withdrawal_address = EXCLUDED.f_withdrawal_address,
			f_withdrawal_block = EXCLUDED.f_withdrawal_block
		WHERE t_validator_last_deposit.f_withdrawal_block <= EXCLUDED.f_withdrawal_block;
	`

	// Mirrors the backfill in migration 000011.
	rebuildValidatorLastDepositQuery = `
		INSERT INTO t_validator_last_deposit (
			f_validator_pubkey,
			f_depositor,
			f_depositor_block,
			f_withdrawal_address,
			f_withdrawal_block
		)
		SELECT
			ld.f_validator_pubkey,
			ld.f_depositor,
			ld.f_block_num,
			COALESCE(lw.f_withdrawal_address, ''),
			COALESCE(lw.f_block_num, 0)
		FROM (
			SELECT DISTINCT ON (f_validator_pubkey)
				f_validator_pubkey,
				f_depositor,
				f_block_num
			FROM t_beacon_deposits
			ORDER BY f_validator_pubkey, f_block_num DESC
		) ld
		LEFT JOIN (
			SELECT DISTINCT ON (f_validator_pubkey)
				f_validator_pubkey,
				f_withdrawal_address,
				f_block_num
			FROM t_beacon_deposits
			WHERE f_withdrawal_address != ''
			ORDER BY f_validator_pubkey, f_block_num DESC
		) lw ON ld.f_validator_pubkey = lw.f_validator_pubkey;
	`
)

// upsertValidatorLastDeposit refreshes t_validator_last_deposit from a batch
// of freshly downloaded deposits sitting in tempTableName. It must run inside
// the same transaction that inserts the batch into t_beacon_deposits so the
// two tables never drift apart.
func (p *PostgresDBService) upsertValidatorLastDeposit(tx pgx.Tx, tempTableName string) error {
	_, err := tx.Exec(p.ctx, fmt.Sprintf(upsertLastDepositorQuery, tempTableName))
	if err != nil {
		return errors.Wrap(err, "error upserting last depositor per validator")
	}

	_, err = tx.Exec(p.ctx, fmt.Sprintf(upsertLastWithdrawalQuery, tempTableName))
	if err != nil {
		return errors.Wrap(err, "error upserting last withdrawal address per validator")
	}

	return nil
}

// RebuildValidatorLastDeposit rebuilds t_validator_last_deposit from scratch
// out of t_beacon_deposits. Meant for the periodic full rebuild (identify with
// --recreate-table), as a safety valve against any drift.
func (p *PostgresDBService) RebuildValidatorLastDeposit() error {
	p.writerThreadsWG.Add(1)
	defer p.writerThreadsWG.Done()
	startTime := time.Now()

	conn, err := p.psqlPool.Acquire(p.ctx)
	if err != nil {
		return errors.Wrap(err, "error acquiring connection")
	}
	defer conn.Release()

	tx, err := conn.Begin(p.ctx)
	if err != nil {
		return errors.Wrap(err, "error beginning rebuild transaction")
	}
	defer tx.Rollback(p.ctx)

	_, err = tx.Exec(p.ctx, `TRUNCATE TABLE t_validator_last_deposit;`)
	if err != nil {
		return errors.Wrap(err, "error truncating t_validator_last_deposit")
	}

	_, err = tx.Exec(p.ctx, rebuildValidatorLastDepositQuery)
	if err != nil {
		return errors.Wrap(err, "error rebuilding t_validator_last_deposit")
	}

	if err := tx.Commit(p.ctx); err != nil {
		return errors.Wrap(err, "error committing rebuild transaction")
	}

	_, err = conn.Exec(p.ctx, `ANALYZE t_validator_last_deposit;`)
	if err != nil {
		return errors.Wrap(err, "error analyzing t_validator_last_deposit")
	}

	wlog.Infof("rebuilt t_validator_last_deposit in %.2f seconds", time.Since(startTime).Seconds())
	return nil
}
