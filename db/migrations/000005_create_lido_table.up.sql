CREATE TABLE IF NOT EXISTS public.t_lido
(
    f_validator_pubkey text NOT NULL,
    f_operator text NOT NULL,
    f_operator_index integer NOT NULL,
    CONSTRAINT t_lido_pkey PRIMARY KEY (f_validator_pubkey)
);
