package curated

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/migalabs/eth-pokhar/lido"

	"github.com/ethereum/go-ethereum/ethclient"
)

const NODE_OPS_ADDRESS = "0x55032650b14df07b85bF18A3a3eC8E0Af2e028d5"

// definedOperatorsNames pins the tag of every Curated operator by registry
// index, because the on-chain name is a free-form string the operator set
// itself and is not stable enough to build a URL with. GetOperatorName only
// falls back to lido.FormatOperatorName past the end of this list.
//
// Last reconciled against mainnet on 2026-10-02 by reading
// getNodeOperator(id, true) for every id. Entries that differ from the
// registry on purpose are the ones whose registry name carries a legal
// suffix or a character that has no business in a URL.
//
// The consequence, and it has already bitten: an operator that renames itself
// on chain keeps the tag written here until someone edits it by hand. The
// name is fetched on every run and then dropped on the floor for these
// indices, so the drift is invisible. Check the entry against
// getNodeOperator(index, true) before trusting it.
var definedOperatorsNames = []string{
	// Wave 0
	"stakingfacilities_lido", // 0
	"jumpcrypto_lido",        // 1 (registry: Jump Crypto)
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
	"blockdaemon_lido",   // 13
	"stakin_lido",        // 14
	"chainlayer_lido",    // 15
	"simplystaking_lido", // 16
	"solstice_lido",      // 17 (registry: Solstice)
	"stakely_lido",       // 18
	"infstones_lido",     // 19
	"hashkeycloud_lido",  // 20 (registry: HashKey Cloud)
	"consensys_lido",     // 21 (registry: Consensys)
	// Wave 4
	"rocklogicgmbh_lido", // 22
	// Galaxy Digital acquired substantially all assets of CryptoManufaktur
	// LLC, including its engineering team, on 2024-07-19. The registry was
	// updated to match: getNodeOperator(23, true) returns "Galaxy" today.
	// This entry was simply never updated with it, and the list wins over
	// the chain, so the stale tag survived. Custody is unaffected either
	// way: the withdrawal credentials stay with the Lido vault, which is
	// what the "_lido" suffix records.
	// https://www.galaxy.com/newsroom/galaxy-expands-blockchain-infrastructure-capabilities-asset-acquisition-crypto-manufaktur
	"galaxy_lido",           // 23 (registered as CryptoManufaktur)
	"kukisglobal_lido",      // 24
	"twinstake_lido",        // 25 (registry: Twinstake)
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
	"mavan_lido",        // 36 (registry: MAVAN)
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

// PinnedTagDrift returns the tag the registry name would produce when it
// disagrees with the pinned one, and "" when they agree or the operator is
// not pinned.
//
// Overriding the registry is the point of definedOperatorsNames, so a
// disagreement is not an error by itself: several registry names carry legal
// suffixes or characters that have no business in a URL. What it must never
// be is invisible. The name is read from the chain on every run and then
// dropped for pinned indices, so an operator that renames itself keeps the
// old tag until a human happens to notice, which for index 23 took two years.
// Surfacing the disagreement turns that into one log line per run.
func PinnedTagDrift(operator NodeOperator) string {
	if operator.Index >= uint64(len(definedOperatorsNames)) {
		return ""
	}
	fromRegistry := lido.FormatOperatorName(operator.Name)
	if fromRegistry == definedOperatorsNames[operator.Index] {
		return ""
	}
	return fromRegistry
}
