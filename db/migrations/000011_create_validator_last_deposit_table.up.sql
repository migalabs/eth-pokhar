-- Materialized index of the latest deposit facts per validator, kept up to
-- date incrementally by CopyBeaconDeposits. Depositor and withdrawal address
-- are tracked separately: rows downloaded before migration 000008 have
-- f_withdrawal_address = '' forever, so the latest row overall and the latest
-- row with a non-empty withdrawal address can differ for the same validator.
CREATE TABLE IF NOT EXISTS public.t_validator_last_deposit (
    f_validator_pubkey TEXT NOT NULL,
    f_depositor TEXT NOT NULL,
    f_depositor_block BIGINT NOT NULL,
    f_withdrawal_address TEXT NOT NULL DEFAULT '',
    f_withdrawal_block BIGINT NOT NULL DEFAULT 0,
    CONSTRAINT t_validator_last_deposit_pkey PRIMARY KEY (f_validator_pubkey)
);

CREATE INDEX IF NOT EXISTS idx_t_validator_last_deposit_f_depositor
    ON public.t_validator_last_deposit (f_depositor);
CREATE INDEX IF NOT EXISTS idx_t_validator_last_deposit_f_withdrawal_address
    ON public.t_validator_last_deposit (f_withdrawal_address);

-- Backfill from the current contents of t_beacon_deposits. Mirrors
-- rebuildValidatorLastDepositQuery in db/validator_last_deposit.go.
INSERT INTO t_validator_last_deposit (
    f_validator_pubkey,
    f_depositor,
    f_depositor_block,
    f_withdrawal_address,
    f_withdrawal_block
)
SELECT
    ld.f_validator_pubkey,
    ld.f_depositor,
    ld.f_block_num,
    COALESCE(lw.f_withdrawal_address, ''),
    COALESCE(lw.f_block_num, 0)
FROM (
    SELECT DISTINCT ON (f_validator_pubkey)
        f_validator_pubkey,
        f_depositor,
        f_block_num
    FROM t_beacon_deposits
    ORDER BY f_validator_pubkey, f_block_num DESC
) ld
LEFT JOIN (
    SELECT DISTINCT ON (f_validator_pubkey)
        f_validator_pubkey,
        f_withdrawal_address,
        f_block_num
    FROM t_beacon_deposits
    WHERE f_withdrawal_address != ''
    ORDER BY f_validator_pubkey, f_block_num DESC
) lw ON ld.f_validator_pubkey = lw.f_validator_pubkey
ON CONFLICT (f_validator_pubkey) DO NOTHING;

ANALYZE t_validator_last_deposit;
