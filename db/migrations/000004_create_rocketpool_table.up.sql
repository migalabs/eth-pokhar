CREATE TABLE IF NOT EXISTS public.t_rocketpool
(
    f_validator_pubkey text NOT NULL,
    CONSTRAINT t_rocketpool_pkey PRIMARY KEY (f_validator_pubkey)
);
