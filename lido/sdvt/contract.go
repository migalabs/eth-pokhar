package sdvt

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/migalabs/eth-pokhar/lido"
	"github.com/migalabs/eth-pokhar/lido/curated"
)

// Lido Simple DVT NodeOperatorsRegistry proxy on mainnet.
// Shares the same ABI as the Curated module; only the address differs.
const NODE_OPS_ADDRESS = "0xae7B191a31F627b4eB1d4DaC64eaB9976995b433"

type SDVTContract struct {
	contract *curated.NodeOperatorsRegistryCaller
	address  string
}

func NewSDVTContract(nodeEndpoint string) (*SDVTContract, error) {
	ethClient, err := ethclient.Dial(nodeEndpoint)
	if err != nil {
		return nil, err
	}
	defer ethClient.Close()

	contract, err := curated.NewNodeOperatorsRegistryCaller(common.HexToAddress(NODE_OPS_ADDRESS), ethClient)
	if err != nil {
		return nil, err
	}
	return &SDVTContract{contract: contract, address: NODE_OPS_ADDRESS}, nil
}

func (s *SDVTContract) GetNodeOperatorsCount() (int64, error) {
	result, err := lido.RetryContractCall(func() (interface{}, error) {
		return s.contract.GetNodeOperatorsCount(nil)
	})
	if err != nil {
		return 0, err
	}
	return result.(*big.Int).Int64(), nil
}

func (s *SDVTContract) GetOperatorData(index *big.Int) (curated.NodeOperator, error) {
	result, err := lido.RetryContractCall(func() (interface{}, error) {
		return s.contract.GetNodeOperator(nil, index, true)
	})
	if err != nil {
		return curated.NodeOperator{}, err
	}
	operator := result.(struct {
		Active            bool
		Name              string
		RewardAddress     common.Address
		StakingLimit      uint64
		StoppedValidators uint64
		TotalSigningKeys  uint64
		UsedSigningKeys   uint64
	})
	return curated.NodeOperator{
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

func (s *SDVTContract) GetOperatorKeys(operator curated.NodeOperator, offset uint64, limit uint64) (curated.OperatorKeys, error) {
	result, err := lido.RetryContractCall(func() (interface{}, error) {
		return s.contract.GetSigningKeys(nil, big.NewInt(int64(operator.Index)), big.NewInt(int64(offset)), big.NewInt(int64(limit)))
	})
	if err != nil {
		return curated.OperatorKeys{}, err
	}
	operatorKeys := result.(struct {
		Pubkeys    []byte
		Signatures []byte
		Used       []bool
	})
	return curated.OperatorKeys{
		PubKeys:    operatorKeys.Pubkeys,
		Signatures: operatorKeys.Signatures,
		Used:       operatorKeys.Used,
	}, nil
}

// GetOperatorName returns the tag for an SDVT operator. Unlike Curated, SDVT
// has no canonical pre-defined operator-name list — names come straight from
// the on-chain contract and are normalized through lido.FormatOperatorName.
func GetOperatorName(operator curated.NodeOperator) string {
	return lido.FormatOperatorName(operator.Name)
}
