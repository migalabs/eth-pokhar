CREATE TABLE IF NOT EXISTS public.t_identified_validators
(
    f_validator_pubkey text NOT NULL,
    f_pool_name text NOT NULL DEFAULT 'others'::text,
    CONSTRAINT t_identified_validators_pkey PRIMARY KEY (f_validator_pubkey)
);

CREATE TABLE IF NOT EXISTS public.t_validators_insert
(
    f_validator_pubkey text NOT NULL,
    f_pool_name text NOT NULL,
    CONSTRAINT t_validators_insert_pkey PRIMARY KEY (f_validator_pubkey)
);

CREATE TABLE IF NOT EXISTS public.t_depositors_insert
(
    f_depositor text NOT NULL,
    f_pool_name text NOT NULL,
    CONSTRAINT t_depositors_insert_pkey PRIMARY KEY (f_depositor)
);