DROP INDEX IF EXISTS idx_t_lido_group_id;
ALTER TABLE t_lido
    DROP COLUMN IF EXISTS f_group_id,
    DROP COLUMN IF EXISTS f_group_name;
