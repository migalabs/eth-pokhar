package curated

import "github.com/ethereum/go-ethereum/common"

type NodeOperator struct {
	Index             uint64
	Active            bool
	Name              string
	RewardAddress     common.Address
	StakingLimit      uint64
	StoppedValidators uint64
	TotalSigningKeys  uint64
	UsedSigningKeys   uint64
}

type OperatorKeys struct {
	PubKeys    []byte
	Signatures []byte
	Used       []bool
}
