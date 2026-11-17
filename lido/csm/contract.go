package csm

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	lido "github.com/migalabs/eth-pokhar/lido"
)

// Lido staking modules that expose the CSM operator and key registry interface
// (getNodeOperatorsCount, getNodeOperator, getSigningKeys), so the CSM binding
// reads all of them by address. Their getNodeOperator tuple has one more
// trailing field (usedPriorityQueue) than the generated NodeOperator; the
// decoder ignores it and every field we read lines up. Ids are the
// StakingRouter module ids.
const (
	// CSMContractAddress is the Community Staking Module (module 3, type
	// community-onchain-v1).
	CSMContractAddress = "0xdA7dE2ECdDfccC6c3AF10108Db212ACBBf9EA83F"
	// CMv2ContractAddress is the Curated Module v2 (module 4, type
	// curated-onchain-v2): curated operators running 0x02 validators.
	CMv2ContractAddress = "0xDa5F930cE326EB5205085D66c72A4E79d60cB8C1"
	// CSM0x02ContractAddress is Community Staking 0x02 (module 5, type
	// community-onchain-v1).
	CSM0x02ContractAddress = "0x792Cd25e4aE3578375031FB55e048E163A804F7B"
)

// Entity labels for the modules the dashboard shows as a single entity ("Lido
// CMv2" and "Lido CSM0x02", mapped to those display names downstream, the same
// way csm_operator<N>_lido becomes "Lido CSM"). Two constraints from consumers:
// they keep "lido" in the name because consumers classify pooled stake with
// LIKE '%lido%', and they must not start with "csm" because CSM metrics select
// LIKE 'csm_%' ('_' matches any character) and would absorb them.
const (
	CMv2EntityTag    = "lido_cmv2"
	CSM0x02EntityTag = "lido_csm0x02"
)

type CSMContract struct {
	contract *CsmCaller
	address  string
}

// NewCSMContract returns a reader for the Community Staking Module.
func NewCSMContract(nodeEndpoint string) (*CSMContract, error) {
	return NewContractAt(nodeEndpoint, CSMContractAddress)
}

// NewContractAt returns a reader for any module that shares the CSM interface,
// e.g. CMv2ContractAddress or CSM0x02ContractAddress.
func NewContractAt(nodeEndpoint string, address string) (*CSMContract, error) {
	ethClient, err := ethclient.Dial(nodeEndpoint)
	if err != nil {
		return nil, err
	}
	defer ethClient.Close()

	contract, err := NewCsmCaller(common.HexToAddress(address), ethClient)
	if err != nil {
		return nil, err
	}
	return &CSMContract{contract: contract, address: address}, nil
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
		return l.contract.GetSigningKeys(nil, big.NewInt(int64(operator.Index)), big.NewInt(int64(startIndex)), big.NewInt(int64(keyCount)))
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

// CMv2OperatorName is the per-operator tag of a Curated Module v2 operator. It
// lands in t_lido.f_operator and so in the f_operator dimension, while the
// entity label is CMv2EntityTag.
func CMv2OperatorName(operator NodeOperatorCustom) string {
	return fmt.Sprintf("%s_operator%d", CMv2EntityTag, operator.Index)
}

// CSM0x02OperatorName is the per-operator tag of a Community Staking 0x02
// operator, with CSM0x02EntityTag as its entity label.
func CSM0x02OperatorName(operator NodeOperatorCustom) string {
	return fmt.Sprintf("%s_operator%d", CSM0x02EntityTag, operator.Index)
}
