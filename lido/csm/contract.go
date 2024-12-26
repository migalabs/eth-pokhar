package csm

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	lido "github.com/migalabs/eth-pokhar/lido"
)

const csmContractAddress = "0xdA7dE2ECdDfccC6c3AF10108Db212ACBBf9EA83F"

type CSMContract struct {
	contract *CsmCaller
	address  string
}

func NewCSMContract(nodeEndpoint string) (*CSMContract, error) {
	ethClient, err := ethclient.Dial(nodeEndpoint)
	if err != nil {
		return nil, err
	}
	defer ethClient.Close()

	contract, err := NewCsmCaller(common.HexToAddress(csmContractAddress), ethClient)
	if err != nil {
		return nil, err
	}
	return &CSMContract{contract: contract, address: csmContractAddress}, nil
}
func (l *CSMContract) GetNodeOperatorsCount() (int64, error) {
	result, err := lido.RetryContractCall(func() (interface{}, error) {
		return l.contract.GetNodeOperatorsCount(nil)
	})
	if err != nil {
		return 0, err
	}
	return result.(*big.Int).Int64(), nil
}

func (l *CSMContract) GetOperatorData(index *big.Int) (NodeOperatorCustom, error) {
	result, err := lido.RetryContractCall(func() (interface{}, error) {
		return l.contract.GetNodeOperator(nil, index)
	})
	if err != nil {
		return NodeOperatorCustom{}, err
	}
	operator := result.(NodeOperator)
	return NodeOperatorCustom{
		Index:    index.Uint64(),
		Operator: operator,
	}, nil
}

func (l *CSMContract) GetOperatorKeys(operator NodeOperatorCustom, startIndex uint64, keyCount uint64) ([][]byte, error) {
	result, err := lido.RetryContractCall(func() (interface{}, error) {
		return l.contract.GetSigningKeys(nil, big.NewInt(int64(operator.Index)), big.NewInt(0), big.NewInt(int64(keyCount)))
	})
	if err != nil {
		return [][]byte{}, err
	}
	keysConcatenated := result.([]byte)
	keys := make([][]byte, keyCount)
	for i := uint64(0); i < keyCount; i++ {
		keys[i] = keysConcatenated[i*lido.PublicKeyLength : (i+1)*lido.PublicKeyLength]
	}

	return keys, nil
}

func GetOperatorName(operator NodeOperatorCustom) string {
	return fmt.Sprintf("csm_operator%d_lido", operator.Index)
}
