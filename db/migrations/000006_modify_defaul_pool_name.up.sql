ALTER TABLE t_identified_validators 
ALTER COLUMN f_pool_name
SET DEFAULT 'solo_stakers';

UPDATE t_identified_validators
SET f_pool_name = 'solo_stakers'
WHERE f_pool_name = 'others';