CREATE TABLE IF NOT EXISTS public.t_beacon_depositors_transactions
(
    f_block_num bigint DEFAULT 0,
    f_value double precision DEFAULT 0,
    f_to bytea,
    f_from bytea,
    f_tx_hash bytea NOT NULL,
    CONSTRAINT t_beacon_depositors_transactions_pkey PRIMARY KEY (f_tx_hash)
)
