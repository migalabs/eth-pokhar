CREATE INDEX IF NOT EXISTS idx_t_beacon_deposits_f_withdrawal_address ON public.t_beacon_deposits (f_withdrawal_address);
CREATE INDEX IF NOT EXISTS idx_t_beacon_deposits_f_depositor ON public.t_beacon_deposits (f_depositor);
CREATE INDEX IF NOT EXISTS idx_t_beacon_deposits_f_validator_pubkey ON public.t_beacon_deposits (f_validator_pubkey);
