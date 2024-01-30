CREATE TABLE IF NOT EXISTS public.t_beacon_depositors_transactions
(
    f_block_num bigint DEFAULT 0,
    f_value double precision DEFAULT 0,
    f_to text,
    f_from text,
    f_tx_hash text NOT NULL,
    CONSTRAINT t_beacon_depositors_transactions_pkey PRIMARY KEY (f_tx_hash)
);

CREATE TABLE IF NOT EXISTS public.t_beacon_deposits
(
    f_block_num bigint DEFAULT 0,
    f_depositor text,
    f_tx_hash text NOT NULL,
    f_validator_pubkey text,
    CONSTRAINT t_beacon_deposits_pkey PRIMARY KEY (f_tx_hash,f_validator_pubkey)
);