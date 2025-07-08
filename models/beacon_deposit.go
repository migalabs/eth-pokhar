package models

type BeaconDeposit struct {
	BlockNum          uint64 `json:"block_num"`
	Depositor         string `json:"depositor"`
	TxHash            string `json:"tx_hash"`
	ValidatorPubkey   string `json:"validator_pubkey"`
	WithdrawalAddress string `json:"withdrawal_address"`
}
