package cmv2

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/migalabs/eth-pokhar/lido"
	log "github.com/sirupsen/logrus"
)

// Lido Curated Module v2 addresses.
// TODO: replace with the mainnet addresses once Lido confirms them.
// Hoodi testnet values, taken from https://vgorkavenko.github.io/cm-groups-viewer/
// (chainId 560048) and verified by direct RPC calls against the deployed proxies.
const (
	MetaRegistryAddressHoodi  = "0x857289cCBFBc4C134Cc312022a104CD9b38d8AAE"
	CuratedModuleAddressHoodi = "0x87EB69Ae51317405FD285efD2326a4a11f6173b9"
)

// Client groups the two reads we need for the CM v2 pipeline: group metadata
// from the MetaRegistry and signing keys from the Curated module itself.
//
// The contracts are independent proxies; an instance of Client holds a caller
// against each of them so that the identify step can interleave group
// enumeration and per-operator key fetches without re-dialing the EL.
type Client struct {
	endpoint      string
	ethClient     *ethclient.Client
	metaRegistry  *MetaRegistryCaller
	curatedModule *CuratedModuleCaller
	metaAddress   common.Address
	curAddress    common.Address
}

// Close releases the underlying EL connection. Safe to call multiple times.
func (c *Client) Close() {
	if c.ethClient != nil {
		c.ethClient.Close()
		c.ethClient = nil
	}
}

// NewClient dials the EL endpoint and binds callers for the MetaRegistry and
// CuratedModule v2 proxies at the supplied addresses. Pass empty strings to
// fall back to the Hoodi testnet defaults — useful during early development
// while the mainnet deployment is pending.
func NewClient(endpoint, metaRegistryAddr, curatedModuleAddr string) (*Client, error) {
	if endpoint == "" {
		return nil, fmt.Errorf("empty EL endpoint")
	}
	if metaRegistryAddr == "" {
		log.Warnf("CM v2: no MetaRegistry address supplied, falling back to Hoodi testnet default %s", MetaRegistryAddressHoodi)
		metaRegistryAddr = MetaRegistryAddressHoodi
	}
	if curatedModuleAddr == "" {
		log.Warnf("CM v2: no CuratedModule address supplied, falling back to Hoodi testnet default %s", CuratedModuleAddressHoodi)
		curatedModuleAddr = CuratedModuleAddressHoodi
	}
	if !common.IsHexAddress(metaRegistryAddr) {
		return nil, fmt.Errorf("invalid MetaRegistry address: %q", metaRegistryAddr)
	}
	if !common.IsHexAddress(curatedModuleAddr) {
		return nil, fmt.Errorf("invalid CuratedModule address: %q", curatedModuleAddr)
	}
	ethClient, err := ethclient.Dial(endpoint)
	if err != nil {
		return nil, fmt.Errorf("dialing EL endpoint: %w", err)
	}
	// Note: the ethClient stays alive for the lifetime of the Client. Callers
	// that want to release the underlying http connections can invoke Close()
	// once they are done with the contract.
	metaAddr := common.HexToAddress(metaRegistryAddr)
	curAddr := common.HexToAddress(curatedModuleAddr)

	meta, err := NewMetaRegistryCaller(metaAddr, ethClient)
	if err != nil {
		return nil, fmt.Errorf("binding MetaRegistry caller: %w", err)
	}
	cur, err := NewCuratedModuleCaller(curAddr, ethClient)
	if err != nil {
		return nil, fmt.Errorf("binding CuratedModule caller: %w", err)
	}
	return &Client{
		endpoint:      endpoint,
		ethClient:     ethClient,
		metaRegistry:  meta,
		curatedModule: cur,
		metaAddress:   metaAddr,
		curAddress:    curAddr,
	}, nil
}

// GetOperatorGroupsCount returns the total number of operator group slots in
// the MetaRegistry, including the slot 0 sentinel reserved for operators
// without a group (the contract exposes this sentinel through NoGroupId).
func (c *Client) GetOperatorGroupsCount() (int64, error) {
	result, err := lido.RetryContractCall(func() (interface{}, error) {
		return c.metaRegistry.GetOperatorGroupsCount(nil)
	})
	if err != nil {
		return 0, err
	}
	return result.(*big.Int).Int64(), nil
}

// NoGroupId returns the sentinel groupId used by the MetaRegistry to mark
// operators that have not been assigned to any group.
func (c *Client) NoGroupId() (*big.Int, error) {
	result, err := lido.RetryContractCall(func() (interface{}, error) {
		return c.metaRegistry.NOGROUPID(nil)
	})
	if err != nil {
		return nil, err
	}
	return result.(*big.Int), nil
}

// GetOperatorGroup returns the OperatorGroup struct for the given groupId,
// containing the list of native sub-operators (with their share) and the
// opaque external-operator references that point back to v1 NodeOperators
// Registry entries.
func (c *Client) GetOperatorGroup(groupId int64) (IMetaRegistryOperatorGroup, error) {
	result, err := lido.RetryContractCall(func() (interface{}, error) {
		return c.metaRegistry.GetOperatorGroup(nil, big.NewInt(groupId))
	})
	if err != nil {
		return IMetaRegistryOperatorGroup{}, err
	}
	return result.(IMetaRegistryOperatorGroup), nil
}

// GetOperatorMetadata returns the name/description metadata stored in the
// MetaRegistry for a given node operator. The name is the on-chain human-
// readable name (e.g. "Attestant (BVI) Limited"). Description is currently
// unused by the identify pipeline but kept around for future surfacing.
func (c *Client) GetOperatorMetadata(nodeOperatorId uint64) (OperatorMetadata, error) {
	result, err := lido.RetryContractCall(func() (interface{}, error) {
		return c.metaRegistry.GetOperatorMetadata(nil, new(big.Int).SetUint64(nodeOperatorId))
	})
	if err != nil {
		return OperatorMetadata{}, err
	}
	return result.(OperatorMetadata), nil
}

// GetNodeOperatorsCount returns the total number of operators registered in
// the Curated module v2 contract.
func (c *Client) GetNodeOperatorsCount() (int64, error) {
	result, err := lido.RetryContractCall(func() (interface{}, error) {
		return c.curatedModule.GetNodeOperatorsCount(nil)
	})
	if err != nil {
		return 0, err
	}
	return result.(*big.Int).Int64(), nil
}

// GetNodeOperator returns the on-chain NodeOperator struct (counts of added,
// deposited, vetted, exited keys, plus addresses) for a single operator.
func (c *Client) GetNodeOperator(nodeOperatorId uint64) (NodeOperator, error) {
	result, err := lido.RetryContractCall(func() (interface{}, error) {
		return c.curatedModule.GetNodeOperator(nil, new(big.Int).SetUint64(nodeOperatorId))
	})
	if err != nil {
		return NodeOperator{}, err
	}
	return result.(NodeOperator), nil
}

// GetSigningKeys reads `count` consecutive signing keys starting from
// `offset` for the given operator. The contract returns the keys as a packed
// byte array of fixed-size BLS public keys.
func (c *Client) GetSigningKeys(nodeOperatorId uint64, offset uint64, count uint64) ([]byte, error) {
	result, err := lido.RetryContractCall(func() (interface{}, error) {
		return c.curatedModule.GetSigningKeys(
			nil,
			new(big.Int).SetUint64(nodeOperatorId),
			new(big.Int).SetUint64(offset),
			new(big.Int).SetUint64(count),
		)
	})
	if err != nil {
		return nil, err
	}
	return result.([]byte), nil
}

// GetNonce returns the Curated module's nonce, which advances on any state
// change relevant to operators or keys. The identify pipeline can use it to
// skip a re-fetch when nothing has changed since the previous successful run.
func (c *Client) GetNonce() (*big.Int, error) {
	result, err := lido.RetryContractCall(func() (interface{}, error) {
		return c.curatedModule.GetNonce(nil)
	})
	if err != nil {
		return nil, err
	}
	return result.(*big.Int), nil
}

// CuratedV1StakingModuleID is the staking-router module id of the legacy
// Curated NodeOperatorsRegistry (v1). MetaRegistry external-operator
// references with this module id point back to entries already tracked by
// lido/curated in t_lido (f_protocol = 'curated').
const CuratedV1StakingModuleID uint16 = 1

// DecodeExternalOperator parses the opaque `data` field of an
// IMetaRegistry.ExternalOperator into (moduleId, operatorId). The on-chain
// layout used by the MetaRegistry is a 10-byte big-endian record:
//
//	bytes[0:2]  → uint16 moduleId
//	bytes[2:10] → uint64 operatorId
//
// Verified empirically against Hoodi (chainId 560048) groups #1..#4 in the
// cm-groups-viewer (module 1 / no. {1,18,28,13}). If Lido ever extends the
// encoding to carry extra fields we'll need to widen this decoder.
func DecodeExternalOperator(data []byte) (moduleId uint16, operatorId uint64, ok bool) {
	if len(data) < 10 {
		return 0, 0, false
	}
	moduleId = uint16(data[0])<<8 | uint16(data[1])
	for i := 2; i < 10; i++ {
		operatorId = operatorId<<8 | uint64(data[i])
	}
	return moduleId, operatorId, true
}

// FormatGroupName produces the canonical group label persisted in t_lido. The
// MetaRegistry does not expose human-readable names for groups yet, so we use
// a stable numeric label that downstream consumers can rely on. The leading
// zero-pad keeps the label sortable.
func FormatGroupName(groupId int64) string {
	return fmt.Sprintf("cmv2_group_%03d", groupId)
}

// FormatOperatorName mirrors lido.FormatOperatorName for CM v2 operators: the
// name comes from the on-chain MetaRegistry metadata and is normalized to a
// URL-safe slug. If the registry returns an empty name (early-deployment
// operators sometimes have no name set), we fall back to a numeric label.
func FormatOperatorName(opId uint64, raw string) string {
	trimmed := strings.TrimSpace(raw)
	if trimmed == "" {
		return fmt.Sprintf("cmv2_operator_%d", opId)
	}
	return lido.FormatOperatorName(trimmed)
}
