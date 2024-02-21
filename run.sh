#!/bin/bash

./eth_pokhar beacon_depositors_transactions --log-level=${LOG_LEVEL} --el-endpoint=${EL_ENDPOINT} --db-url=${DB_URL} --workers-num=${WORKER_NUM} --alchemy-url=${ALCHEMY_URL}

./eth_pokhar identify --log-level=${LOG_LEVEL} --el-endpoint=${EL_ENDPOINT} --db-url=${DB_URL} --workers-num=${WORKER_NUM} --alchemy-url=${ALCHEMY_URL} --recreate-table --whale-threshold=${WHALE_THRESHOLD}
