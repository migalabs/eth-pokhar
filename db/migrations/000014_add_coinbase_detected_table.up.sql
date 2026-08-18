-- Materializes the coinbase detection so the pure dimension columns do not
-- depend on the mutable display label. Before this table, the dimension pass
-- inferred "coinbase-detected" from f_pool_name = 'coinbase', which breaks as
-- soon as a later phase or a display-claiming pin overwrites the label: the
-- validator's coinbase custody and operation fallbacks silently disappear
-- (review finding on PR #43). IdentifyCoinbaseValidators now persists its
-- full detection result here (idempotent inserts, the set only grows) and
-- both the display update and the dimension pass read from this table.
CREATE TABLE IF NOT EXISTS public.t_coinbase_detected (
    f_validator_pubkey TEXT NOT NULL,
    CONSTRAINT t_coinbase_detected_pkey PRIMARY KEY (f_validator_pubkey)
);

-- Seed with the currently coinbase-labeled set: exactly the validators whose
-- label came from the detection heuristic (or from a coinbase depositor
-- mapping, whose custody is coinbase's in either case).
INSERT INTO t_coinbase_detected (f_validator_pubkey)
SELECT f_validator_pubkey FROM t_identified_validators
WHERE f_pool_name = 'coinbase'
ON CONFLICT (f_validator_pubkey) DO NOTHING;
