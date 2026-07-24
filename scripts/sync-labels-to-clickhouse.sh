#!/bin/bash
#
# sync-labels-to-clickhouse.sh
#
# Distributes the validator -> pool labels produced by eth-pokhar (Postgres,
# t_identified_validators) to a goteth ClickHouse database (t_eth2_pubkeys),
# implementing the design of issue #31:
#
#   1. Atomic apply: the snapshot is imported into t_eth2_pubkeys_staging and
#      applied with EXCHANGE TABLES, so consumers always see either the old or
#      the new complete mapping, never an empty or partial one.
#   2. Versioned: every applied snapshot writes a row to t_eth2_pubkeys_version
#      (ReplacingMergeTree keyed on a single row, query it with FINAL).
#   3. Gated: the snapshot is rejected (keeping the previous mapping) if it
#      shrinks more than SYNC_MIN_COUNT_RATIO allows, or if any existing pool
#      loses more than SYNC_MAX_POOL_DROP validators and is not listed in the
#      rename allowlist.
#   4. Monitored: optionally exports cron_job_* metrics for node_exporter's
#      textfile collector. Alert on cron_job_last_run_timestamp_seconds age.
#
# After a successful run, t_eth2_pubkeys_staging holds the PREVIOUS mapping.
# Rollback = running EXCHANGE TABLES again by hand.
#
# Configuration (environment variables, e.g. from a protected .env file):
#
#   PG_HOST / PG_PORT / PG_DB / PG_USER / PG_PASS
#       Labels Postgres, as reachable FROM the ClickHouse server process
#       (the import uses the postgresql() table function server-side).
#   CH_CLIENT      command used to reach ClickHouse (default: clickhouse-client)
#   CH_PORT        native port (default: 9000)
#   CH_USER / CH_PASS / CH_DB
#   SYNC_MIN_COUNT_RATIO       minimum new/old row ratio (default: 0.95)
#   SYNC_MAX_POOL_DROP         max validators an existing pool may lose
#                              (default: 5000, 0 disables the gate)
#   SYNC_POOL_RENAME_ALLOWLIST optional file with one pool name per line,
#                              exempt from the pool-drop gate (renames/splits)
#   TEXTFILE_DIR               optional node_exporter textfile directory
#   CRON_NAME                  metric label (default: labels_sync)
#
set -u

LOG_PREFIX="[labels-sync]"
PG_HOST="${PG_HOST:?PG_HOST is required}"
PG_PORT="${PG_PORT:?PG_PORT is required}"
PG_DB="${PG_DB:?PG_DB is required}"
PG_USER="${PG_USER:?PG_USER is required}"
PG_PASS="${PG_PASS:?PG_PASS is required}"
CH_CLIENT="${CH_CLIENT:-clickhouse-client}"
CH_PORT="${CH_PORT:-9000}"
CH_USER="${CH_USER:-default}"
CH_PASS="${CH_PASS:-}"
CH_DB="${CH_DB:?CH_DB is required}"
SYNC_MIN_COUNT_RATIO="${SYNC_MIN_COUNT_RATIO:-0.95}"
SYNC_MAX_POOL_DROP="${SYNC_MAX_POOL_DROP:-5000}"
SYNC_POOL_RENAME_ALLOWLIST="${SYNC_POOL_RENAME_ALLOWLIST:-}"
TEXTFILE_DIR="${TEXTFILE_DIR:-}"
CRON_NAME="${CRON_NAME:-labels_sync}"

status=1
start=$(date +%s%3N)
TMPDIR_SYNC=$(mktemp -d)

cleanup() {
    end=$(date +%s%3N)
    runtime=$((end - start))
    rm -rf "$TMPDIR_SYNC"
    if [ -n "$TEXTFILE_DIR" ]; then
        prom="$TEXTFILE_DIR/${CRON_NAME}.prom"
        {
            echo "cron_job_last_run_timestamp_seconds{cron_name=\"${CRON_NAME}\"} $(date +%s)"
            echo "cron_job_last_run_time_taken_milliseconds{cron_name=\"${CRON_NAME}\"} ${runtime}"
            echo "cron_job_last_run_exit_status{cron_name=\"${CRON_NAME}\"} ${status}"
        } > "${prom}.$$" && mv "${prom}.$$" "${prom}"
    fi
    echo "$LOG_PREFIX $(date -u +%FT%TZ) Finished with status=$status in ${runtime}ms"
}
trap cleanup EXIT

ch() {
    $CH_CLIENT --port "$CH_PORT" --user "$CH_USER" --password "$CH_PASS" --database "$CH_DB" --query "$1"
}

echo "$LOG_PREFIX $(date -u +%FT%TZ) Starting labels sync..."

# ---------------------------------------------------------------- previous state
PREV_COUNT=$(ch "SELECT count() FROM t_eth2_pubkeys") || exit 1
echo "$LOG_PREFIX Previous mapping: $PREV_COUNT entries"

if [ "$PREV_COUNT" -gt 0 ] && [ "$SYNC_MAX_POOL_DROP" -gt 0 ]; then
    ch "SELECT f_pool_name, count() FROM t_eth2_pubkeys FINAL GROUP BY f_pool_name FORMAT TSV" \
        > "$TMPDIR_SYNC/pools_prev.tsv" || exit 1
fi

# ---------------------------------------------------------------- import to staging
ch "CREATE TABLE IF NOT EXISTS t_eth2_pubkeys_staging AS t_eth2_pubkeys" || exit 1
ch "TRUNCATE TABLE t_eth2_pubkeys_staging" || exit 1

echo "$LOG_PREFIX Importing snapshot into t_eth2_pubkeys_staging..."
ch "
    INSERT INTO t_eth2_pubkeys_staging (f_val_idx, f_public_key, f_pool_name, f_pool)
    SELECT
        vls.f_val_idx,
        vls.f_public_key,
        iv.f_pool_name,
        iv.f_pool_name AS f_pool
    FROM postgresql('${PG_HOST}:${PG_PORT}', '${PG_DB}', 't_identified_validators', '${PG_USER}', '${PG_PASS}') AS iv
    INNER JOIN t_validator_last_status AS vls
        ON concat('0x', iv.f_validator_pubkey) = vls.f_public_key
" || { echo "$LOG_PREFIX ERROR: import into staging failed, live mapping untouched"; exit 1; }

NEW_COUNT=$(ch "SELECT count() FROM t_eth2_pubkeys_staging") || exit 1
echo "$LOG_PREFIX New snapshot: $NEW_COUNT entries"

# ---------------------------------------------------------------- sanity gates
if [ "$NEW_COUNT" -eq 0 ]; then
    echo "$LOG_PREFIX GATE FAILED: new snapshot is empty, live mapping untouched"
    exit 1
fi

if [ "$PREV_COUNT" -gt 0 ]; then
    ok=$(awk -v new="$NEW_COUNT" -v prev="$PREV_COUNT" -v ratio="$SYNC_MIN_COUNT_RATIO" \
        'BEGIN { print (new >= prev * ratio) ? 1 : 0 }')
    if [ "$ok" -ne 1 ]; then
        echo "$LOG_PREFIX GATE FAILED: snapshot shrank from $PREV_COUNT to $NEW_COUNT entries (min ratio $SYNC_MIN_COUNT_RATIO), live mapping untouched"
        exit 1
    fi
fi

if [ "$PREV_COUNT" -gt 0 ] && [ "$SYNC_MAX_POOL_DROP" -gt 0 ]; then
    ch "SELECT f_pool_name, count() FROM t_eth2_pubkeys_staging FINAL GROUP BY f_pool_name FORMAT TSV" \
        > "$TMPDIR_SYNC/pools_new.tsv" || exit 1
    ALLOWLIST="${SYNC_POOL_RENAME_ALLOWLIST}"
    VIOLATIONS=$(awk -F'\t' -v maxdrop="$SYNC_MAX_POOL_DROP" -v allowfile="$ALLOWLIST" '
        BEGIN {
            if (allowfile != "") {
                while ((getline line < allowfile) > 0) { allowed[line] = 1 }
                close(allowfile)
            }
        }
        NR == FNR { newc[$1] = $2; next }
        {
            drop = $2 - newc[$1]
            if (drop > maxdrop && !($1 in allowed)) {
                printf "%s lost %d validators (%d -> %d)\n", $1, drop, $2, newc[$1]
            }
        }
    ' "$TMPDIR_SYNC/pools_new.tsv" "$TMPDIR_SYNC/pools_prev.tsv")
    if [ -n "$VIOLATIONS" ]; then
        echo "$LOG_PREFIX GATE FAILED: pool drops above threshold ($SYNC_MAX_POOL_DROP), live mapping untouched."
        echo "$LOG_PREFIX If these are legitimate renames or splits, add the pool names to the rename allowlist and rerun:"
        echo "$VIOLATIONS" | while read -r line; do echo "$LOG_PREFIX   $line"; done
        exit 1
    fi
fi

# ---------------------------------------------------------------- atomic apply
echo "$LOG_PREFIX Gates passed. Applying snapshot atomically..."
ch "EXCHANGE TABLES t_eth2_pubkeys AND t_eth2_pubkeys_staging" || exit 1

# ------------------------------------------------- optional t_pubkey_pool mirror
# Some deployments also serve a pubkey -> pool mirror without requiring a
# val_idx assignment (labels for deposits still in the pending queue). Same
# staging + EXCHANGE pattern; the source snapshot already passed the gates.
if [ "${SYNC_PUBKEY_POOL:-false}" = "true" ]; then
    echo "$LOG_PREFIX Syncing t_pubkey_pool mirror..."
    ch "CREATE TABLE IF NOT EXISTS t_pubkey_pool_staging AS t_pubkey_pool" || exit 1
    ch "TRUNCATE TABLE t_pubkey_pool_staging" || exit 1
    ch "
        INSERT INTO t_pubkey_pool_staging (f_public_key, f_pool_name)
        SELECT
            concat('0x', iv.f_validator_pubkey) AS f_public_key,
            iv.f_pool_name
        FROM postgresql('${PG_HOST}:${PG_PORT}', '${PG_DB}', 't_identified_validators', '${PG_USER}', '${PG_PASS}') AS iv
    " || { echo "$LOG_PREFIX ERROR: t_pubkey_pool import failed, live mirror untouched"; exit 1; }
    PP_COUNT=$(ch "SELECT count() FROM t_pubkey_pool_staging") || exit 1
    if [ "$PP_COUNT" -eq 0 ]; then
        echo "$LOG_PREFIX GATE FAILED: t_pubkey_pool snapshot is empty, live mirror untouched"
        exit 1
    fi
    ch "EXCHANGE TABLES t_pubkey_pool AND t_pubkey_pool_staging" || exit 1
    echo "$LOG_PREFIX t_pubkey_pool updated: $PP_COUNT entries"
fi

# ------------------------------------------------- version stamp
VERSION=$(date +%s)
ch "CREATE TABLE IF NOT EXISTS t_eth2_pubkeys_version (
        f_version UInt64,
        f_synced_at DateTime,
        f_entries UInt64
    ) ENGINE = ReplacingMergeTree(f_version) ORDER BY tuple()" || exit 1
ch "INSERT INTO t_eth2_pubkeys_version VALUES ($VERSION, now(), $NEW_COUNT)" || exit 1

echo "$LOG_PREFIX Applied version $VERSION: $PREV_COUNT -> $NEW_COUNT entries"
echo "$LOG_PREFIX Previous mapping kept in t_eth2_pubkeys_staging (rollback: EXCHANGE TABLES again)"
status=0
