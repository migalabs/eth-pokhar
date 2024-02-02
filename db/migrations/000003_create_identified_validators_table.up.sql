CREATE TABLE IF NOT EXISTS public.t_identified_validators
(
    f_validator_pubkey text NOT NULL,
    f_pool_name text NOT NULL DEFAULT 'others'::text,
    CONSTRAINT t_identified_validators_pkey PRIMARY KEY (f_validator_pubkey)
);