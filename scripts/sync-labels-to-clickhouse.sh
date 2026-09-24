#!/bin/bash
#
# sync-labels-to-clickhouse.sh
#
# Distributes the validator labels produced by eth-pokhar (Postgres,
# t_identified_validators) to a goteth ClickHouse database (t_eth2_pubkeys),
# implementing the design of issue #31. Up to three values travel per
# validator: the display label f_pool_name, and the two pure dimensions
# f_operator and f_custodian, so consumers can tell who runs a validator from
# who holds its withdrawal credentials without re-deriving either.
#
# The dimension columns only exist in destinations that ran goteth migration
# 000041. See SYNC_DIMENSIONS below for how a destination without them is
# handled.
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
#   SYNC_DIMENSIONS            auto (default), true or false: whether to carry
#                              f_operator and f_custodian (see below)
#   SYNC_MIN_DIMENSION_RATIO   minimum new/old ratio for the number of rows
#                              carrying each dimension (default: 0, ratio gate
#                              disabled; a drop to zero is always rejected)
#   TEXTFILE_DIR               optional node_exporter textfile directory
#   CRON_NAME                  metric label (default: labels_sync)
#   SYNC_LOCK_FILE             lock path (default: /tmp/labels-sync-<db>-<port>.lock)
#
# Credentials: the ClickHouse password travels via the CLICKHOUSE_PASSWORD
# environment variable and queries via stdin, never on a command line (which
# any local user can read in /proc/<pid>/cmdline). Docker-based CH_CLIENT
# values must forward the variable, e.g.:
#   CH_CLIENT="docker exec -i -e CLICKHOUSE_PASSWORD my-clickhouse clickhouse-client"
# The Postgres credentials are embedded in the postgresql() query text by
# necessity (table function syntax); stdin keeps them off the process list
# and ClickHouse masks table function secrets in its query_log. A server-side
# named collection removes them from the query text entirely if preferred.
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
SYNC_DIMENSIONS="${SYNC_DIMENSIONS:-auto}"
SYNC_MIN_DIMENSION_RATIO="${SYNC_MIN_DIMENSION_RATIO:-0}"
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

# One run at a time per target database: overlapping runs would race on the
# same staging table and a double EXCHANGE could silently re-apply a stale
# mapping with both runs reporting success.
SYNC_LOCK_FILE="${SYNC_LOCK_FILE:-/tmp/labels-sync-${CH_DB}-${CH_PORT}.lock}"
exec 9>"$SYNC_LOCK_FILE"
if ! flock -n 9; then
    echo "$LOG_PREFIX ERROR: another sync targeting ${CH_DB}:${CH_PORT} is still running (lock: $SYNC_LOCK_FILE), aborting"
    exit 1
fi

export CLICKHOUSE_PASSWORD="$CH_PASS"
ch() {
    printf '%s' "$1" | $CH_CLIENT --port "$CH_PORT" --user "$CH_USER" --database "$CH_DB"
}

echo "$LOG_PREFIX $(date -u +%FT%TZ) Starting labels sync..."

# ---------------------------------------------------------------- dimensions mode
# f_operator and f_custodian exist only in destinations that ran goteth
# migration 000041, and deployments are not upgraded in lockstep. Refusing to
# run against an older destination would also stop distributing pool labels
# there, which is the silent staleness this script exists to prevent, so the
# default adapts to the destination and says so. Resolved before any DDL, so
# a destination that cannot take the dimensions is never left half written.
#
#   auto (default)  carry the dimensions when the destination has them,
#                   distribute pool labels alone (with a warning) when it does not
#   true            require them, fail before writing anything otherwise
#   false           never carry them
DIM_COLS=$(ch "
    SELECT count()
    FROM system.columns
    WHERE database = currentDatabase()
      AND table = 't_eth2_pubkeys'
      AND name IN ('f_operator', 'f_custodian')
") || exit 1

case "$SYNC_DIMENSIONS" in
    true)
        if [ "$DIM_COLS" -ne 2 ]; then
            echo "$LOG_PREFIX ERROR: SYNC_DIMENSIONS=true but ${CH_DB}.t_eth2_pubkeys has $DIM_COLS of the 2 dimension columns."
            echo "$LOG_PREFIX Apply goteth migration 000041 to the destination first. Nothing was written."
            exit 1
        fi
        CARRY_DIMENSIONS=true
        ;;
    false)
        CARRY_DIMENSIONS=false
        ;;
    auto)
        if [ "$DIM_COLS" -eq 2 ]; then
            CARRY_DIMENSIONS=true
        else
            CARRY_DIMENSIONS=false
            echo "$LOG_PREFIX WARNING: ${CH_DB}.t_eth2_pubkeys has $DIM_COLS of the 2 dimension columns (goteth migration 000041), syncing pool labels only"
        fi
        ;;
    *)
        echo "$LOG_PREFIX ERROR: SYNC_DIMENSIONS must be auto, true or false (got '$SYNC_DIMENSIONS')"
        exit 1
        ;;
esac

if [ "$CARRY_DIMENSIONS" = true ]; then
    echo "$LOG_PREFIX Carrying entity dimensions f_operator and f_custodian"
    SNAPSHOT_DIM_DDL=",
        f_operator String,
        f_custodian String"
    SNAPSHOT_DIM_COLS=", f_operator, f_custodian"
    # NULL in Postgres means "no declared signal for this dimension"; the
    # destination columns are plain String defaulting to '', so the absence
    # keeps exactly one spelling all the way through.
    SNAPSHOT_DIM_SELECT=",
        coalesce(f_operator, ''),
        coalesce(f_custodian, '')"
    STAGING_DIM_COLS=", f_operator, f_custodian"
    STAGING_DIM_SELECT=",
        iv.f_operator,
        iv.f_custodian"
else
    SNAPSHOT_DIM_DDL=""
    SNAPSHOT_DIM_COLS=""
    SNAPSHOT_DIM_SELECT=""
    STAGING_DIM_COLS=""
    STAGING_DIM_SELECT=""
fi

# ---------------------------------------------------------------- previous state
PREV_COUNT=$(ch "SELECT count() FROM t_eth2_pubkeys") || exit 1
echo "$LOG_PREFIX Previous mapping: $PREV_COUNT entries"

if [ "$PREV_COUNT" -gt 0 ] && [ "$SYNC_MAX_POOL_DROP" -gt 0 ]; then
    ch "SELECT f_pool_name, count() FROM t_eth2_pubkeys FINAL GROUP BY f_pool_name FORMAT TSV" \
        > "$TMPDIR_SYNC/pools_prev.tsv" || exit 1
fi

# ---------------------------------------------------------------- import snapshot
# Postgres is read exactly ONCE per run into t_labels_snapshot; the main
# mapping and the optional t_pubkey_pool mirror are both derived from it, so
# every distributed table reflects the same instant of the source and the
# gates below vouch for all of them. A second postgresql() read for the
# mirror would reintroduce the cross-table drift issue #31 exists to prevent.
# Recreated rather than CREATE IF NOT EXISTS plus TRUNCATE: the table is
# ephemeral, and deployments that ran an earlier version of this script still
# have it with the pre-dimension schema, which IF NOT EXISTS would keep
# forever. Unqualified names resolve in CH_DB, so each destination database
# owns its own snapshot and concurrent runs against different ones do not
# collide.
ch "DROP TABLE IF EXISTS t_labels_snapshot" || exit 1
ch "CREATE TABLE t_labels_snapshot (
        f_validator_pubkey String,
        f_pool_name String${SNAPSHOT_DIM_DDL}
    ) ENGINE = MergeTree ORDER BY tuple()" || exit 1

echo "$LOG_PREFIX Importing labels snapshot from Postgres..."
ch "
    INSERT INTO t_labels_snapshot (f_validator_pubkey, f_pool_name${SNAPSHOT_DIM_COLS})
    SELECT
        f_validator_pubkey,
        f_pool_name${SNAPSHOT_DIM_SELECT}
    FROM postgresql('${PG_HOST}:${PG_PORT}', '${PG_DB}', 't_identified_validators', '${PG_USER}', '${PG_PASS}')
" || { echo "$LOG_PREFIX ERROR: snapshot import failed, live mapping untouched"; exit 1; }

# Recreated, not reused: EXCHANGE TABLES swaps names without comparing
# schemas, so a staging table left over from before a column was added to
# t_eth2_pubkeys swaps that column straight back out of the live table, with
# no error and no way to tell from the log. Recreating costs nothing here,
# because the rollback copy this table is meant to hold is the one the swap
# below puts there, not the one from the previous run, which the old code
# truncated at exactly this point anyway.
ch "DROP TABLE IF EXISTS t_eth2_pubkeys_staging" || exit 1
ch "CREATE TABLE t_eth2_pubkeys_staging AS t_eth2_pubkeys" || exit 1

ch "
    INSERT INTO t_eth2_pubkeys_staging (f_val_idx, f_public_key, f_pool_name, f_pool${STAGING_DIM_COLS})
    SELECT
        vls.f_val_idx,
        vls.f_public_key,
        iv.f_pool_name,
        iv.f_pool_name AS f_pool${STAGING_DIM_SELECT}
    FROM t_labels_snapshot AS iv
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

# A pool that stops being populated upstream shows up in the pool-drop gate
# above; a dimension that does not would quietly blank the column for every
# validator instead. A drop to zero from a non-zero previous state is
# unambiguous breakage and is always rejected; smaller drops need an explicit
# ratio, because the first run against a freshly migrated destination
# legitimately goes from no coverage at all to full coverage.
if [ "$CARRY_DIMENSIONS" = true ] && [ "$PREV_COUNT" -gt 0 ]; then
    PREV_DIMS=$(ch "
        SELECT countIf(f_operator != ''), countIf(f_custodian != '')
        FROM t_eth2_pubkeys FINAL FORMAT TSV") || exit 1
    NEW_DIMS=$(ch "
        SELECT countIf(f_operator != ''), countIf(f_custodian != '')
        FROM t_eth2_pubkeys_staging FINAL FORMAT TSV") || exit 1
    read -r PREV_OP PREV_CUST <<< "$PREV_DIMS"
    read -r NEW_OP NEW_CUST <<< "$NEW_DIMS"
    DIM_VIOLATIONS=$(awk -v po="$PREV_OP" -v pc="$PREV_CUST" -v no="$NEW_OP" -v nc="$NEW_CUST" \
        -v r="$SYNC_MIN_DIMENSION_RATIO" 'BEGIN {
            if (po > 0 && no == 0) printf "f_operator lost every row (%d -> 0)\n", po
            else if (r > 0 && po > 0 && no < po * r) printf "f_operator dropped below ratio %s (%d -> %d)\n", r, po, no
            if (pc > 0 && nc == 0) printf "f_custodian lost every row (%d -> 0)\n", pc
            else if (r > 0 && pc > 0 && nc < pc * r) printf "f_custodian dropped below ratio %s (%d -> %d)\n", r, pc, nc
        }')
    if [ -n "$DIM_VIOLATIONS" ]; then
        echo "$LOG_PREFIX GATE FAILED: entity dimension coverage collapsed, live mapping untouched."
        echo "$DIM_VIOLATIONS" | while read -r line; do echo "$LOG_PREFIX   $line"; done
        exit 1
    fi
fi

# ---------------------------------------------------------------- atomic apply
echo "$LOG_PREFIX Gates passed. Applying snapshot atomically..."
ch "EXCHANGE TABLES t_eth2_pubkeys AND t_eth2_pubkeys_staging" || exit 1

# From this point on the new mapping is LIVE: any failure below makes the run
# exit non-zero for alerting purposes, but must not be read as "the sync did
# not happen". Rolling back requires an explicit EXCHANGE TABLES, not a rerun.
post_swap_fail() {
    echo "$LOG_PREFIX WARNING: $1 failed AFTER the swap: the new t_eth2_pubkeys mapping IS already live. Do not roll back based on this exit code alone."
    exit 1
}

# ------------------------------------------------- optional t_pubkey_pool mirror
# Some deployments also serve a pubkey -> pool mirror without requiring a
# val_idx assignment (labels for deposits still in the pending queue). Same
# staging + EXCHANGE pattern, built from t_labels_snapshot: the exact same
# gated read of Postgres the main mapping was built from, never a second one.
if [ "${SYNC_PUBKEY_POOL:-false}" = "true" ]; then
    echo "$LOG_PREFIX Syncing t_pubkey_pool mirror..."
    ch "CREATE TABLE IF NOT EXISTS t_pubkey_pool_staging AS t_pubkey_pool" || post_swap_fail "t_pubkey_pool staging setup"
    ch "TRUNCATE TABLE t_pubkey_pool_staging" || post_swap_fail "t_pubkey_pool staging truncate"
    ch "
        INSERT INTO t_pubkey_pool_staging (f_public_key, f_pool_name)
        SELECT
            concat('0x', iv.f_validator_pubkey) AS f_public_key,
            iv.f_pool_name
        FROM t_labels_snapshot AS iv
    " || post_swap_fail "t_pubkey_pool build (live mirror untouched)"
    PP_COUNT=$(ch "SELECT count() FROM t_pubkey_pool_staging") || post_swap_fail "t_pubkey_pool count"
    if [ "$PP_COUNT" -eq 0 ]; then
        post_swap_fail "t_pubkey_pool gate: snapshot is empty (live mirror untouched)"
    fi
    ch "EXCHANGE TABLES t_pubkey_pool AND t_pubkey_pool_staging" || post_swap_fail "t_pubkey_pool swap (live mirror untouched)"
    echo "$LOG_PREFIX t_pubkey_pool updated: $PP_COUNT entries"
fi

ch "TRUNCATE TABLE t_labels_snapshot" || true

# ------------------------------------------------- version stamp
VERSION=$(date +%s)
ch "CREATE TABLE IF NOT EXISTS t_eth2_pubkeys_version (
        f_version UInt64,
        f_synced_at DateTime,
        f_entries UInt64
    ) ENGINE = ReplacingMergeTree(f_version) ORDER BY tuple()" || post_swap_fail "version table setup"
ch "INSERT INTO t_eth2_pubkeys_version VALUES ($VERSION, now(), $NEW_COUNT)" || post_swap_fail "version stamp"

echo "$LOG_PREFIX Applied version $VERSION: $PREV_COUNT -> $NEW_COUNT entries"
echo "$LOG_PREFIX Previous mapping kept in t_eth2_pubkeys_staging (rollback: EXCHANGE TABLES again)"
status=0
