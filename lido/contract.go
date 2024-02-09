package lido

import (
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/migalabs/eth-pokhar/utils"
	"github.com/pkg/errors"

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

func (l *LidoContract) GetNodeOperatorsCount() (int64, error) {
	retry := 0
	var operatorsCount *big.Int
	var err error
	for {
		operatorsCount, err = l.contract.GetNodeOperatorsCount(nil)
		if err != nil {
			if !strings.Contains(err.Error(), "429") {
				retry++
			}
			if retry > 5 {
				return 0, errors.Wrap(err, "error getting operators count")
			}
			waitTime := utils.GetRandomTimeout()
			time.Sleep(waitTime)
			continue
		}
		break
	}

	return operatorsCount.Int64(), nil
}

func (l *LidoContract) GetOperatorData(index *big.Int) (NodeOperator, error) {
	retry := 0
	for {
		operator, err := l.contract.GetNodeOperator(nil, index, true)
		if err != nil {
			if !strings.Contains(err.Error(), "429") {
				retry++
			}
			if retry > 5 {
				return NodeOperator{}, errors.Wrap(err, "error getting operators count")
			}
			waitTime := utils.GetRandomTimeout()
			time.Sleep(waitTime)
			continue
		}
		return NodeOperator{
			Index:             index.Uint64(),
			Active:            operator.Active,
			Name:              operator.Name,
			RewardAddress:     operator.RewardAddress,
			StakingLimit:      operator.StakingLimit,
			StoppedValidators: operator.StoppedValidators,
			TotalSigningKeys:  operator.TotalSigningKeys,
			UsedSigningKeys:   operator.UsedSigningKeys,
		}, nil
	}

}

func (l *LidoContract) GetOperatorKey(operator NodeOperator, keyIndex uint64) (OperatorKey, error) {
	retry := 0
	for {
		key, err := l.contract.GetSigningKey(nil, big.NewInt(int64(operator.Index)), big.NewInt(int64(keyIndex)))
		if err != nil {
			if !strings.Contains(err.Error(), "429") {
				retry++
			}
			if retry > 5 {
				return OperatorKey{}, errors.Wrap(err, "error getting operators count")
			}
			waitTime := utils.GetRandomTimeout()
			time.Sleep(waitTime)
			continue
		}
		return OperatorKey{
			Key:              key.Key,
			DepositSignature: key.DepositSignature,
			Used:             key.Used,
		}, nil
	}
}
