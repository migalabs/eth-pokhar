package models

type Transaction struct {
	BlockNum  uint64  `json:"block_num"`
	Value     float64 `json:"value"`
	From      string  `json:"from"`
	To        string  `json:"to"`
	TxHash    string  `json:"tx_hash"`
	Depositor string  `json:"depositor"`
}

func (t Transaction) Type() ModelType {
	return TransactionModel
}
