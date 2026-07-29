#!/bin/bash
# Launcher for sync-labels-to-clickhouse.sh: loads a local env file, optionally
# opens an SSH tunnel to the labels Postgres, and runs the sync engine.
#
# Usage:
#   run-labels-sync.sh [/path/to/labels-sync.env]
#
# The env file defaults to labels-sync.env next to this script. It must define
# the engine variables (PG_*, CH_*, SYNC_*, TEXTFILE_DIR, CRON_NAME; see
# README.md) and may define the tunnel variables:
#
#   SYNC_TUNNEL_SSH_HOST   ssh host (alias or user@host) that can reach the
#                          labels Postgres. Empty or unset: no tunnel is
#                          opened and PG_HOST/PG_PORT are used as-is.
#   SYNC_TUNNEL_BIND       local address to bind (default 127.0.0.1)
#   SYNC_TUNNEL_PORT       local port to bind (default 15440)
#   SYNC_TUNNEL_TARGET     host:port of Postgres as seen from the ssh host
#                          (default localhost:5432)
#   SYNC_TUNNEL_WAIT_SECS  max seconds to wait for the tunnel (default 20)
set -u

DIR="$(cd "$(dirname "$0")" && pwd)"
ENV_FILE="${1:-$DIR/labels-sync.env}"

if [ ! -f "$ENV_FILE" ]; then
    echo "[labels-sync-wrapper] ERROR: env file not found: $ENV_FILE"
    exit 1
fi

set -a
. "$ENV_FILE"
set +a

SYNC_TUNNEL_SSH_HOST="${SYNC_TUNNEL_SSH_HOST:-}"
SYNC_TUNNEL_BIND="${SYNC_TUNNEL_BIND:-127.0.0.1}"
SYNC_TUNNEL_PORT="${SYNC_TUNNEL_PORT:-15440}"
SYNC_TUNNEL_TARGET="${SYNC_TUNNEL_TARGET:-localhost:5432}"
SYNC_TUNNEL_WAIT_SECS="${SYNC_TUNNEL_WAIT_SECS:-20}"

if [ -n "$SYNC_TUNNEL_SSH_HOST" ]; then
    ssh -N -L "${SYNC_TUNNEL_BIND}:${SYNC_TUNNEL_PORT}:${SYNC_TUNNEL_TARGET}" "$SYNC_TUNNEL_SSH_HOST" &
    TUNNEL_PID=$!
    trap 'kill $TUNNEL_PID 2>/dev/null || true; wait $TUNNEL_PID 2>/dev/null || true' EXIT

    # 0.0.0.0 is not connectable as a destination, probe via loopback instead
    CHECK_HOST="$SYNC_TUNNEL_BIND"
    [ "$CHECK_HOST" = "0.0.0.0" ] && CHECK_HOST=127.0.0.1

    waited=0
    until timeout 2 bash -c "echo > /dev/tcp/${CHECK_HOST}/${SYNC_TUNNEL_PORT}" 2>/dev/null; do
        if ! kill -0 "$TUNNEL_PID" 2>/dev/null; then
            echo "[labels-sync-wrapper] ERROR: ssh tunnel process died"
            exit 1
        fi
        waited=$((waited + 1))
        if [ "$waited" -ge "$SYNC_TUNNEL_WAIT_SECS" ]; then
            echo "[labels-sync-wrapper] ERROR: tunnel not reachable at ${CHECK_HOST}:${SYNC_TUNNEL_PORT} after ${SYNC_TUNNEL_WAIT_SECS}s"
            exit 1
        fi
        sleep 1
    done
fi

exec_status=0
bash "${SYNC_ENGINE_PATH:-$DIR/sync-labels-to-clickhouse.sh}" || exec_status=$?
exit $exec_status
