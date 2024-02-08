package lido

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/ethclient"
)

const NODE_OPS_ADDRESS = "0x55032650b14df07b85bF18A3a3eC8E0Af2e028d5"

type LidoContract struct {
	contract *NodeOperatorsRegistryCaller
	address  string
}

func NewLidoContract(nodeEndpoint string) (*LidoContract, error) {
	ethClient, err := ethclient.Dial(nodeEndpoint)
	if err != nil {
		return nil, err
	}
	defer ethClient.Close()

	contract, err := NewNodeOperatorsRegistryCaller(common.HexToAddress(NODE_OPS_ADDRESS), ethClient)
	if err != nil {
		return nil, err
	}
	return &LidoContract{contract: contract, address: NODE_OPS_ADDRESS}, nil
}

func (l *LidoContract) GetOperatorsIndexes() ([]*big.Int, error) {
	operatorsCount, err := l.contract.GetNodeOperatorsCount(nil)
	if err != nil {
		return nil, err
	}
	operatorsIndexes := make([]*big.Int, operatorsCount.Int64())
	for i := 0; i < int(operatorsCount.Int64()); i++ {
		operatorsIndexes[i] = big.NewInt(int64(i))
	}
	return operatorsIndexes, nil
}

func (l *LidoContract) GetOperatorsData(indexes []*big.Int) ([]NodeOperator, error) {
	operators := make([]NodeOperator, len(indexes))
	for i, index := range indexes {
		operator, err := l.contract.GetNodeOperator(nil, index, true)
		if err != nil {
			return nil, err
		}
		operators[i] = NodeOperator{
			Index:             index.Uint64(),
			Active:            operator.Active,
			Name:              operator.Name,
			RewardAddress:     operator.RewardAddress,
			StakingLimit:      operator.StakingLimit,
			StoppedValidators: operator.StoppedValidators,
			TotalSigningKeys:  operator.TotalSigningKeys,
			UsedSigningKeys:   operator.UsedSigningKeys,
		}
	}
	return operators, nil
}

func (l *LidoContract) GetOperatorKey(operator NodeOperator, keyIndex uint64) (OperatorKey, error) {
	key, err := l.contract.GetSigningKey(nil, big.NewInt(int64(operator.Index)), big.NewInt(int64(keyIndex)))
	if err != nil {
		return OperatorKey{}, err
	}
	return OperatorKey{
		Key:              key.Key,
		DepositSignature: key.DepositSignature,
		Used:             key.Used,
	}, nil

}
