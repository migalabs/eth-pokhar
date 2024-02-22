ALTER TABLE t_identified_validators 
ALTER COLUMN f_pool_name
SET DEFAULT 'others';

UPDATE t_identified_validators
SET f_pool_name = 'others'
WHERE f_pool_name = 'solo_stakers';