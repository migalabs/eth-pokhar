package lido

import (
	"encoding/binary"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
)

// Custody labels of the Lido staking modules: the f_custodian dimension of
// every key found in that module's registry. Lido is presented as one custody
// entity per module ("Lido CMv1", "Lido SDVT", ...), and each module's
// validators split into their node operators through f_operator. The CMv2 and
// CSM 0x02 labels are also those modules' display label (f_pool_name).
const (
	CuratedCustodyTag = "lido_cmv1"
	SDVTCustodyTag    = "lido_sdvt"
	CSMCustodyTag     = "lido_csm"
)

// CuratedCanonicalOperators is the operator label (f_operator) of every
// Curated Module v1 node operator, by registry index. It is the company, not
// its Lido slice: an operator that also runs validators outside Lido keeps
// the label it has there (figment, kiln, p2porg...), so the operators view
// shows one Figment and not "figment" plus "figment_lido". The others take
// the base of their display tag, with the registry's current name where the
// company renamed (Jump Crypto, Solstice, HashKey Cloud, Twinstake, MAVAN).
//
// The display label (f_pool_name, definedOperatorsNames in lido/curated) is
// a separate list and is not affected.
var CuratedCanonicalOperators = []string{
	"stakingfacilities", // 0  Staking Facilities
	"jumpcrypto",        // 1  Jump Crypto (displayed as certusone_lido)
	"p2porg",            // 2  P2P.org
	"chorusone",         // 3  Chorus One
	"stakefish",         // 4  stakefish
	"blockscape",        // 5  Blockscape
	"dsrv",              // 6  DSRV
	"everstake",         // 7  Everstake
	"kiln",              // 8  Kiln
	"rockx",             // 9  RockX
	"figment",           // 10 Figment
	"allnodes",          // 11 Allnodes
	"anyblockanalytics", // 12 Anyblock Analytics
	"blockdaemon",       // 13 Blockdaemon
	"stakin",            // 14 Stakin by The Tie
	"chainlayer",        // 15 ChainLayer
	"simplystaking",     // 16 Simply Staking
	"solstice",          // 17 Solstice (displayed as bridgetower_lido)
	"stakely",           // 18 Stakely
	"infstones",         // 19 InfStones
	"hashkeycloud",      // 20 HashKey Cloud (displayed as hashquark_lido)
	"consensys",         // 21 Consensys
	"rocklogicgmbh",     // 22 RockLogic GmbH
	"galaxy_digital",    // 23 Galaxy (displayed as cryptomanufaktur_lido)
	"kukisglobal",       // 24 Kukis Global
	"twinstake",         // 25 Twinstake (displayed as nethermind_lido)
	"chainsafe",         // 26 ChainSafe
	"prysmteam",         // 27 Prysm Team at Offchain Labs
	"sigmaprime",        // 28 Sigma Prime
	"attestantlimited",  // 29 Attestant (BVI) Limited
	"launchnodes",       // 30 Launchnodes
	"senseinode",        // 31 SenseiNode
	"a41",               // 32 A41
	"develpgmbh",        // 33 Develp GmbH
	"ebunker",           // 34 Ebunker
	"gateway.fmas",      // 35 Gateway.fm AS
	"mavan",             // 36 MAVAN (displayed as numic_lido)
	"parafi",            // 37 ParaFi Technologies LLC
	"rockawayx",         // 38 RockawayX Infra
}

// CMv2CuratedOperator links each Curated Module v2 node operator to the
// Curated Module v1 operator of the same company, so both share one operator
// label. Taken from the CMv2 MetaRegistry: getNodeOperatorGroupId(op) ->
// getOperatorGroup(group).externalOperators, which reference module 1 (the
// curated NodeOperatorsRegistry). Operators 3 and 6 belong to no group; their
// link comes from where their stake was consolidated from (rockx_lido and
// consensyscodefi_lido). CheckCMv2OperatorLinks reports drift against the
// chain; an operator missing here keeps its index tag lido_cmv2_operator<N>.
var CMv2CuratedOperator = map[uint64]uint64{
	0: 29, 1: 22, 2: 16, 3: 9, 4: 33, 5: 5, 6: 21, 7: 5, 8: 25, 9: 24,
	10: 34, 11: 34, 12: 18, 13: 35, 14: 31, 15: 7, 16: 13, 17: 15, 18: 14, 19: 4,
	20: 2, 21: 22, 22: 22, 23: 2, 24: 28, 25: 37, 26: 36, 27: 30, 28: 30, 29: 20,
	30: 4, 31: 6, 32: 6, 33: 7, 34: 38, 35: 17, 36: 26, 37: 0, 38: 0, 39: 11,
	40: 19, 41: 19, 42: 23, 43: 13, 44: 10, 45: 27,
}

// CMv2MetaRegistryAddress is the MetaRegistry of the Curated Module v2.
const CMv2MetaRegistryAddress = "0xA64b339eebD3dC3De848298B6a140955932901d8"

const cmv2MetaRegistryABI = `[
 {"inputs":[{"name":"nodeOperatorId","type":"uint256"}],"name":"getNodeOperatorGroupId","outputs":[{"name":"","type":"uint256"}],"stateMutability":"view","type":"function"},
 {"inputs":[{"name":"groupId","type":"uint256"}],"name":"getOperatorGroup","outputs":[{"name":"","type":"tuple","components":[
   {"name":"name","type":"string"},
   {"name":"subNodeOperators","type":"tuple[]","components":[{"name":"nodeOperatorId","type":"uint64"},{"name":"share","type":"uint16"}]},
   {"name":"externalOperators","type":"tuple[]","components":[{"name":"data","type":"bytes"}]}]}],"stateMutability":"view","type":"function"}
]`

type cmv2OperatorGroup struct {
	Name             string
	SubNodeOperators []struct {
		NodeOperatorId uint64
		Share          uint16
	}
	ExternalOperators []struct {
		Data []byte
	}
}

// DecodeExternalOperator splits a MetaRegistry external operator reference:
// a big-endian uint16 staking module id followed by a big-endian uint64 node
// operator id.
func DecodeExternalOperator(data []byte) (moduleID uint16, operatorID uint64, err error) {
	if len(data) != 10 {
		return 0, 0, errors.Errorf("unexpected external operator length %d", len(data))
	}
	return binary.BigEndian.Uint16(data[:2]), binary.BigEndian.Uint64(data[2:]), nil
}

// CMv2CuratedLinksOnChain returns, for CMv2 node operators 0..count-1, the
// curated v1 operator the MetaRegistry groups each one with. Operators in no
// group, or whose group references no curated v1 operator, are absent.
func CMv2CuratedLinksOnChain(nodeEndpoint string, count uint64) (map[uint64]uint64, error) {
	ethClient, err := ethclient.Dial(nodeEndpoint)
	if err != nil {
		return nil, err
	}
	defer ethClient.Close()
	parsed, err := abi.JSON(strings.NewReader(cmv2MetaRegistryABI))
	if err != nil {
		return nil, errors.Wrap(err, "error parsing MetaRegistry ABI")
	}
	contract := bind.NewBoundContract(common.HexToAddress(CMv2MetaRegistryAddress), parsed, ethClient, nil, nil)

	links := make(map[uint64]uint64)
	for op := uint64(0); op < count; op++ {
		groupRes, err := RetryContractCall(func() (interface{}, error) {
			var out []interface{}
			if err := contract.Call(&bind.CallOpts{}, &out, "getNodeOperatorGroupId", new(big.Int).SetUint64(op)); err != nil {
				return nil, err
			}
			return *abi.ConvertType(out[0], new(*big.Int)).(**big.Int), nil
		})
		if err != nil {
			return nil, errors.Wrapf(err, "group of CMv2 operator %d", op)
		}
		groupID := groupRes.(*big.Int)
		if groupID.Sign() == 0 {
			continue
		}
		groupRaw, err := RetryContractCall(func() (interface{}, error) {
			var out []interface{}
			if err := contract.Call(&bind.CallOpts{}, &out, "getOperatorGroup", groupID); err != nil {
				return nil, err
			}
			return *abi.ConvertType(out[0], new(cmv2OperatorGroup)).(*cmv2OperatorGroup), nil
		})
		if err != nil {
			return nil, errors.Wrapf(err, "group %s of CMv2 operator %d", groupID, op)
		}
		for _, ext := range groupRaw.(cmv2OperatorGroup).ExternalOperators {
			moduleID, operatorID, err := DecodeExternalOperator(ext.Data)
			if err != nil {
				return nil, errors.Wrapf(err, "group %s of CMv2 operator %d", groupID, op)
			}
			if moduleID == 1 {
				links[op] = operatorID
				break
			}
		}
	}
	return links, nil
}
