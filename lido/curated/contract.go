package curated

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/migalabs/eth-pokhar/lido"

	"github.com/ethereum/go-ethereum/ethclient"
)

const NODE_OPS_ADDRESS = "0x55032650b14df07b85bF18A3a3eC8E0Af2e028d5"

var definedOperatorsNames = []string{
	// Wave 0
	"stakingfacilities_lido", // 0
	"certusone_lido",         // 1
	"p2porg_lido",            // 2
	"chorusone_lido",         // 3
	"stakefish_lido",         // 4
	// Wave 1
	"blockscape_lido", // 5
	"dsrv_lido",       // 6
	"everstake_lido",  // 7
	"kiln_lido",       // 8
	// Wave 2
	"rockx_lido",             // 9
	"figment_lido",           // 10
	"allnodes_lido",          // 11
	"anyblockanalytics_lido", // 12
	// Wave 3
	"blockdaemon_lido",     // 13
	"stakin_lido",          // 14
	"chainlayer_lido",      // 15
	"simplystaking_lido",   // 16
	"bridgetower_lido",     // 17
	"stakely_lido",         // 18
	"infstones_lido",       // 19
	"hashquark_lido",       // 20
	"consensyscodefi_lido", // 21
	// Wave 4
	"rocklogicgmbh_lido",    // 22
	"cryptomanufaktur_lido", // 23
	"kukisglobal_lido",      // 24
	"nethermind_lido",       // 25
	"chainsafe_lido",        // 26
	"prysmaticlabs_lido",    // 27
	"sigmaprime_lido",       // 28
	"attestantlimited_lido", // 29
	// Wave 5
	"launchnodes_lido",  // 30
	"senseinode_lido",   // 31
	"a41_lido",          // 32
	"develpgmbh_lido",   // 33
	"ebunker_lido",      // 34
	"gateway.fmas_lido", // 35
	"numic_lido",        // 36
	"parafi_lido",       // 37
	"rockawayx_lido",    // 38
}

type CuratedModuleContract struct {
	contract *NodeOperatorsRegistryCaller
	address  string
}

func NewCuratedModuleContract(nodeEndpoint string) (*CuratedModuleContract, error) {
	ethClient, err := ethclient.Dial(nodeEndpoint)
	if err != nil {
		return nil, err
	}
	defer ethClient.Close()

	contract, err := NewNodeOperatorsRegistryCaller(common.HexToAddress(NODE_OPS_ADDRESS), ethClient)
	if err != nil {
		return nil, err
	}
	return &CuratedModuleContract{contract: contract, address: NODE_OPS_ADDRESS}, nil
}
func (l *CuratedModuleContract) GetNodeOperatorsCount() (int64, error) {
	result, err := lido.RetryContractCall(func() (interface{}, error) {
		return l.contract.GetNodeOperatorsCount(nil)
	})
	if err != nil {
		return 0, err
	}
	return result.(*big.Int).Int64(), nil
}

func (l *CuratedModuleContract) GetOperatorData(index *big.Int) (NodeOperator, error) {
	result, err := lido.RetryContractCall(func() (interface{}, error) {
		return l.contract.GetNodeOperator(nil, index, true)
	})
	if err != nil {
		return NodeOperator{}, err
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

func (l *CuratedModuleContract) GetOperatorKeys(operator NodeOperator, offset uint64, limit uint64) (OperatorKeys, error) {
	result, err := lido.RetryContractCall(func() (interface{}, error) {
		return l.contract.GetSigningKeys(nil, big.NewInt(int64(operator.Index)), big.NewInt(int64(offset)), big.NewInt(int64(limit)))
	})
	if err != nil {
		return OperatorKeys{}, err
	}
	operatorKeys := result.(struct {
		Pubkeys    []byte
		Signatures []byte
		Used       []bool
	})
	return OperatorKeys{
		PubKeys:    operatorKeys.Pubkeys,
		Signatures: operatorKeys.Signatures,
		Used:       operatorKeys.Used,
	}, nil
}

func GetOperatorName(operator NodeOperator) string {
	if operator.Index < uint64(len(definedOperatorsNames)) {
		return definedOperatorsNames[operator.Index]
	}
	return lido.FormatOperatorName(operator.Name)
}
