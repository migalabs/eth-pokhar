-- Dual tagging (custodian / operator), phase 1 of issue #28's structural fix.
--
-- Custody (who controls the withdrawal credentials) and operation (who runs
-- the validator) are two orthogonal dimensions that f_pool_name squeezes into
-- a single value, so one signal always shadows the other and phase order acts
-- as the arbiter. These columns give each dimension its own slot:
--
--   f_operator   pure operation label. Sources, strongest first: Lido registry
--                operator name, Rocket Pool registry, depositor mapping
--                (t_depositors_insert), coinbase detection as fallback.
--   f_custodian  pure custody label. Sources: withdrawal address mapping
--                (t_withdrawal_address_insert), coinbase detection as fallback.
--
-- NULL means "no declared signal for this dimension", not "unknown entity".
-- f_pool_name stays untouched as the legacy/display label with the historical
-- precedence, including heuristics (whales, solo_stakers) and manual pins,
-- which intentionally do NOT feed the pure columns: their dimension is
-- undeclared. Existing consumers keep working unchanged.
ALTER TABLE t_identified_validators
    ADD COLUMN IF NOT EXISTS f_operator TEXT,
    ADD COLUMN IF NOT EXISTS f_custodian TEXT;

COMMENT ON COLUMN t_identified_validators.f_operator IS
    'Pure operation dimension (who runs the validator). NULL = no declared operator signal. See migration 000012.';
COMMENT ON COLUMN t_identified_validators.f_custodian IS
    'Pure custody dimension (who controls the withdrawal credentials). NULL = no declared custodian signal. See migration 000012.';
