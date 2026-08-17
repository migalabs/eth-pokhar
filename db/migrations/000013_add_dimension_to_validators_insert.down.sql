ALTER TABLE t_validators_insert
    DROP CONSTRAINT IF EXISTS t_validators_insert_dimension_check;

ALTER TABLE t_validators_insert
    DROP COLUMN IF EXISTS f_dimension;
