package models

type Transaction struct {
	BlockNum uint64  `json:"block_num"`
	Value    float64 `json:"value"`
	From     []byte  `json:"from"`
	To       []byte  `json:"to"`
	TxHash   []byte  `json:"tx_hash"`
}

func (t Transaction) Type() ModelType {
	return TransactionModel
}
