-- Curated Module v2 introduces a group level above operators.
-- These two columns let us tag a key with both its operator (existing
-- f_operator column) and its group (new columns).
--
-- Both columns are nullable: rows produced by the legacy modules (curated,
-- csm, sdvt) leave them NULL, and CM v2 externalOperator references will
-- populate them on existing v1 Curated rows without touching f_operator.
ALTER TABLE t_lido
    ADD COLUMN IF NOT EXISTS f_group_id integer,
    ADD COLUMN IF NOT EXISTS f_group_name text;

-- Partial index for the most common lookup pattern: "all validators in group X".
-- Restricted to rows where the group is set so it stays small.
CREATE INDEX IF NOT EXISTS idx_t_lido_group_id
    ON t_lido (f_group_id)
    WHERE f_group_id IS NOT NULL;
