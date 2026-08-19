-- Declared-dimension pins, phase 2 of issue #28's structural fix.
--
-- Phase 1 (migration 000012) computes the pure dimension columns from
-- address-level mappings and on-chain registries. Those sources cannot express
-- pubkey-level operator knowledge: the depositor of a delegated validator is
-- its platform or custodian, not whoever runs the machines (e.g. validators
-- deposited by a liquid-restaking protocol and operated by a third party).
-- This column lets a t_validators_insert row declare WHICH dimension it
-- asserts:
--
--   NULL          legacy pin: writes f_pool_name, never the pure columns.
--   'operator'    writes f_operator and also the display label: entities are
--                 the operator when known (Lido validators show kiln, figment
--                 and so on, not "lido"), so an operator pin claims both.
--   'custodian'   writes f_custodian only; f_pool_name is left untouched.
--
-- Declared rows are the strongest evidence within their dimension: they are
-- curated per-pubkey facts (DVT registries, operator-identifying graffiti,
-- confirmed off-chain data), so they beat every inferred source and double as
-- the manual override.
--
-- The primary key stays f_validator_pubkey: one row per pubkey, one assertion.
-- A pubkey cannot carry a legacy pin and a declared pin at the same time.
ALTER TABLE t_validators_insert
    ADD COLUMN IF NOT EXISTS f_dimension TEXT;

ALTER TABLE t_validators_insert
    ADD CONSTRAINT t_validators_insert_dimension_check
    CHECK (f_dimension IS NULL OR f_dimension IN ('operator', 'custodian'));

COMMENT ON COLUMN t_validators_insert.f_dimension IS
    'Dimension this pin asserts: NULL = legacy f_pool_name pin, operator = f_operator plus the display label, custodian = f_custodian only. See migration 000013.';
