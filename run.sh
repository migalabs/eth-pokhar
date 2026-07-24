#!/bin/bash

# If ONLY_DEPOSITS is true then use --only-deposits flag, else use the default
if [ "${ONLY_DEPOSITS}" = "true" ]; then
    ./eth_pokhar beacon_depositors_transactions --log-level=${LOG_LEVEL} --el-endpoint=${EL_ENDPOINT} --db-url=${DB_URL} --workers-num=${WORKER_NUM} --alchemy-url=${ALCHEMY_URL} --only-deposits || exit 1
else
    ./eth_pokhar beacon_depositors_transactions --log-level=${LOG_LEVEL} --el-endpoint=${EL_ENDPOINT} --db-url=${DB_URL} --workers-num=${WORKER_NUM} --alchemy-url=${ALCHEMY_URL} || exit 1
fi

# RECREATE_TABLE=true truncates t_identified_validators and rebuilds it from
# scratch. Meant for methodology changes or a periodic (e.g. weekly) full
# rebuild; daily runs should be incremental (default).
RECREATE_FLAG=""
if [ "${RECREATE_TABLE:-false}" = "true" ]; then
    RECREATE_FLAG="--recreate-table"
fi

./eth_pokhar identify --log-level=${LOG_LEVEL} --el-endpoint=${EL_ENDPOINT} --db-url=${DB_URL} --workers-num=${WORKER_NUM} --alchemy-url=${ALCHEMY_URL} ${RECREATE_FLAG} --whale-threshold=${WHALE_THRESHOLD} || exit 1
