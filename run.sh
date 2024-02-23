#!/bin/bash

# If ONLY_DEPOSITS is true then use --only-deposits flag, else use the default
if [ "${ONLY_DEPOSITS}" = "true" ]; then
    ./eth_pokhar beacon_depositors_transactions --log-level=${LOG_LEVEL} --el-endpoint=${EL_ENDPOINT} --db-url=${DB_URL} --workers-num=${WORKER_NUM} --alchemy-url=${ALCHEMY_URL} --only-deposits || exit 1
else
    ./eth_pokhar beacon_depositors_transactions --log-level=${LOG_LEVEL} --el-endpoint=${EL_ENDPOINT} --db-url=${DB_URL} --workers-num=${WORKER_NUM} --alchemy-url=${ALCHEMY_URL} || exit 1
fi

./eth_pokhar identify --log-level=${LOG_LEVEL} --el-endpoint=${EL_ENDPOINT} --db-url=${DB_URL} --workers-num=${WORKER_NUM} --alchemy-url=${ALCHEMY_URL} --recreate-table --whale-threshold=${WHALE_THRESHOLD} || exit 1
