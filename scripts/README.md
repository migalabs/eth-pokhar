# Scripts

## sync-labels-to-clickhouse.sh

Distributes the labels produced by the identify pipeline (Postgres `t_identified_validators`) to a goteth ClickHouse database (`t_eth2_pubkeys`). Replaces the legacy per-machine TRUNCATE + INSERT scripts with the design from issue #31:

- **Atomic**: imports into `t_eth2_pubkeys_staging`, applies with `EXCHANGE TABLES`. Consumers never observe an empty or partial mapping. After a successful run the staging table holds the previous mapping, so rollback is one more `EXCHANGE TABLES`.
- **Versioned**: each applied snapshot inserts a row into `t_eth2_pubkeys_version` (query with `FINAL`, it keeps the highest `f_version`). Consumers such as the goteth pool summary reroll can log exactly which labels version they operate on.
- **Gated**: the snapshot is rejected, keeping the previous mapping, if it is empty, if it shrinks below `SYNC_MIN_COUNT_RATIO` of the previous count, or if any existing pool loses more than `SYNC_MAX_POOL_DROP` validators. Legitimate renames or splits are declared in the allowlist file (`SYNC_POOL_RENAME_ALLOWLIST`, one pool name per line) for the run where they happen.
- **Monitored**: with `TEXTFILE_DIR` set, exports `cron_job_last_run_timestamp_seconds`, `cron_job_last_run_time_taken_milliseconds` and `cron_job_last_run_exit_status` for node_exporter's textfile collector.
- **Optional mirror**: with `SYNC_PUBKEY_POOL=true`, also refreshes the `t_pubkey_pool` mirror (pubkey to pool without a val_idx, used to label deposits still in the pending queue) with the same staging plus `EXCHANGE TABLES` pattern. Enable it only on deployments that have that table.

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

Note: the import runs server-side via the `postgresql()` table function, so `PG_HOST:PG_PORT` must be reachable from the ClickHouse **server**, not from the shell running the script. If the Postgres is remote, open the tunnel before invoking the script and bind it on an address the server can reach.

### Recommended alerts

- `time() - cron_job_last_run_timestamp_seconds{cron_name="..."} > 2 * <sync period>`: the sync stopped running (this once went unnoticed for two months).
- `cron_job_last_run_exit_status{cron_name="..."} != 0`: the last snapshot was rejected by a gate or the import failed.
