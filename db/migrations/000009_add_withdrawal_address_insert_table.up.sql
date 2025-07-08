CREATE TABLE IF NOT EXISTS public.t_withdrawal_address_insert
(
    f_withdrawal_address text NOT NULL,
    f_pool_name text NOT NULL,
    CONSTRAINT t_withdrawal_address_insert_pkey PRIMARY KEY (f_withdrawal_address)
);