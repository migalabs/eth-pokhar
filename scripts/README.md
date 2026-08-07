# Scripts

## sync-labels-to-clickhouse.sh

Distributes the labels produced by the identify pipeline (Postgres `t_identified_validators`) to a goteth ClickHouse database (`t_eth2_pubkeys`). Replaces the legacy per-machine TRUNCATE + INSERT scripts with the design from issue #31:

- **Atomic**: imports into `t_eth2_pubkeys_staging`, applies with `EXCHANGE TABLES`. Consumers never observe an empty or partial mapping. After a successful run the staging table holds the previous mapping, so rollback is one more `EXCHANGE TABLES`.
- **Versioned**: each applied snapshot inserts a row into `t_eth2_pubkeys_version` (query with `FINAL`, it keeps the highest `f_version`). Consumers such as the goteth pool summary reroll can log exactly which labels version they operate on.
- **Gated**: the snapshot is rejected, keeping the previous mapping, if it is empty, if it shrinks below `SYNC_MIN_COUNT_RATIO` of the previous count, or if any existing pool loses more than `SYNC_MAX_POOL_DROP` validators. Legitimate renames or splits are declared in the allowlist file (`SYNC_POOL_RENAME_ALLOWLIST`, one pool name per line) for the run where they happen.
- **Monitored**: with `TEXTFILE_DIR` set, exports `cron_job_last_run_timestamp_seconds`, `cron_job_last_run_time_taken_milliseconds` and `cron_job_last_run_exit_status` for node_exporter's textfile collector.
- **Optional mirror**: with `SYNC_PUBKEY_POOL=true`, also refreshes the `t_pubkey_pool` mirror (pubkey to pool without a val_idx, used to label deposits still in the pending queue) with the same staging plus `EXCHANGE TABLES` pattern. Enable it only on deployments that have that table.
- **Single source read**: Postgres is read exactly once per run into `t_labels_snapshot`, and both the main mapping and the mirror are derived from it. All distributed tables therefore reflect the same instant of the source and are covered by the same gates; there is no window for cross-table drift between them.

### Example

```bash
# /etc/cron.d style, every 6 hours. Keep the env file readable only by the cron user.
0 */6 * * *  . /path/to/labels-sync.env && /path/to/scripts/sync-labels-to-clickhouse.sh >> /var/log/labels-sync.log 2>&1
```

```bash
# labels-sync.env
export PG_HOST=<postgres host as reachable from the ClickHouse SERVER process>
export PG_PORT=5439
export PG_DB=eth_pokhar
export PG_USER=...
export PG_PASS=...
export CH_CLIENT="docker exec -i my-clickhouse-container clickhouse-client"
export CH_PORT=9000
export CH_USER=...
export CH_PASS=...
export CH_DB=goteth
export TEXTFILE_DIR=/path/to/node_exporter/textfiles
export CRON_NAME=labels_sync_mainnet
```

Note: the import runs server-side via the `postgresql()` table function, so `PG_HOST:PG_PORT` must be reachable from the ClickHouse **server**, not from the shell running the script. If the Postgres is remote, open the tunnel before invoking the script and bind it on an address the server can reach, or use the launcher below which handles that.

## run-labels-sync.sh

Launcher meant to be the cron entry point. It sources a local env file (argument, or `labels-sync.env` next to the script by default), optionally opens an SSH tunnel to the labels Postgres, waits until the tunnel actually accepts connections, and runs `sync-labels-to-clickhouse.sh` with a cleanup trap that closes the tunnel on exit. Keeps every deployment on the same reviewed logic: the only thing that lives outside the repo is the env file with credentials and endpoints.

```bash
# crontab, every 6 hours
0 */6 * * * /path/to/scripts/run-labels-sync.sh /path/to/labels-sync.env >> /var/log/labels-sync.log 2>&1
```

Tunnel variables, all optional, set in the env file:

```bash
export SYNC_TUNNEL_SSH_HOST=labels-db-host   # unset: no tunnel, PG_HOST used as-is
export SYNC_TUNNEL_BIND=0.0.0.0              # bind on an address the ClickHouse server reaches
export SYNC_TUNNEL_PORT=15440
export SYNC_TUNNEL_TARGET=localhost:5439     # Postgres as seen from the ssh host
export SYNC_TUNNEL_WAIT_SECS=20
```

### Credentials handling

Secrets never travel on a command line, where any local user can read them in `/proc/<pid>/cmdline` for the duration of the call:

- The ClickHouse password is passed through the `CLICKHOUSE_PASSWORD` environment variable (honored by clickhouse-client). Docker-based `CH_CLIENT` values must forward it without a value so it never hits a command line either: `CH_CLIENT="docker exec -i -e CLICKHOUSE_PASSWORD my-clickhouse clickhouse-client"`.
- Queries go through stdin instead of `--query`, which also keeps the Postgres credentials embedded in the `postgresql()` table function calls off the process list. ClickHouse masks table function secrets in `system.query_log`; defining a server-side named collection for the labels Postgres removes them from the query text entirely if you want the extra belt.
- The env file remains the at-rest copy of the secrets: keep it readable only by the cron user.

### Concurrency

The script takes a non-blocking `flock` on `SYNC_LOCK_FILE` (default `/tmp/labels-sync-<db>-<port>.lock`) and aborts with a non-zero status if another run against the same target is still going. Overlapping runs would race on the same staging table, and a double `EXCHANGE TABLES` could silently re-apply a stale mapping with both runs reporting success. A skipped run raises the exit-status metric, which is the correct signal: the interesting event is the previous run still holding the lock.

### On-call note: failures after the swap

Every failure before `EXCHANGE TABLES` leaves the live mapping untouched: rerunning is always safe. Failures AFTER the swap (mirror sync, version stamping) exit non-zero for alerting but the new mapping is already live; the log prints an explicit warning in that case. Do not "roll back" on the basis of the exit code alone: rollback is only ever the manual `EXCHANGE TABLES`, and only if the applied snapshot itself is the problem.

### Recommended alerts

- `time() - cron_job_last_run_timestamp_seconds{cron_name="..."} > 2 * <sync period>`: the sync stopped running (this once went unnoticed for two months).
- `cron_job_last_run_exit_status{cron_name="..."} != 0`: the last snapshot was rejected by a gate, the import failed, a concurrent run was skipped, or a post-swap step failed (see the on-call note).
