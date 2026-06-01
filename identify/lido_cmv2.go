package identify

import (
	"encoding/hex"
	"fmt"
	"sync"

	db "github.com/migalabs/eth-pokhar/db"
	"github.com/migalabs/eth-pokhar/lido"
	"github.com/migalabs/eth-pokhar/lido/cmv2"
	log "github.com/sirupsen/logrus"
)

// cmv2MetaRegistryAddress and cmv2CuratedModuleAddress are placeholders. Until
// the mainnet deployment is announced by Lido, we run against Hoodi, whose
// addresses are baked into the cmv2 package. When mainnet addresses are
// available, wire them through the IdentifyConfig.
//
// Set both to empty strings to use the Hoodi defaults from lido/cmv2.
const (
	cmv2MetaRegistryAddress  = ""
	cmv2CuratedModuleAddress = ""
)

// IdentifyCMv2Validators reads the new Curated Module v2 layout (groups,
// sub-operators inside each group, external operator references back to v1)
// and persists the resulting tags into t_lido.
//
// Pipeline shape mirrors the existing CSM/SDVT/Curated paths: per-group
// concurrent fan-out, each task fetches the keys and inserts them. The work
// expected at production scale is small (single-digit groups, a few tens of
// sub-operators) so we keep the worker pool conservative.
//
// This function is intentionally tolerant of an empty contract (groups count
// == 0): in early mainnet deployment phases the registry may be live but
// unpopulated, in which case the function returns cleanly without writes.
func (i *Identify) IdentifyCMv2Validators() error {
	log.Debug("Starting Curated Module v2 identification")
	client, err := cmv2.NewClient(i.iConfig.ElEndpoint, cmv2MetaRegistryAddress, cmv2CuratedModuleAddress)
	if err != nil {
		return fmt.Errorf("creating cmv2 client: %w", err)
	}
	defer client.Close()

	groupsCount, err := client.GetOperatorGroupsCount()
	if err != nil {
		return fmt.Errorf("reading getOperatorGroupsCount: %w", err)
	}
	noGroupBN, err := client.NoGroupId()
	if err != nil {
		return fmt.Errorf("reading NO_GROUP_ID: %w", err)
	}
	noGroupID := noGroupBN.Int64()
	log.Infof("CM v2 groups discovered: total=%d (sentinel NO_GROUP_ID=%d)", groupsCount, noGroupID)

	if groupsCount <= 1 {
		log.Info("CM v2 has no populated groups yet; skipping")
		return nil
	}

	workerSemaphore := make(chan struct{}, 4)
	var wg sync.WaitGroup

	for groupID := int64(0); groupID < groupsCount; groupID++ {
		if i.stop {
			break
		}
		if groupID == noGroupID {
			continue
		}
		wg.Add(1)
		workerSemaphore <- struct{}{}
		go func(gid int64) {
			defer wg.Done()
			defer func() { <-workerSemaphore }()
			if err := i.processCMv2Group(client, gid); err != nil {
				log.Errorf("CM v2 group %d failed: %v", gid, err)
			}
		}(groupID)
	}
	wg.Wait()
	if i.stop {
		return nil
	}

	log.Debug("Propagating CM v2 group tags onto t_identified_validators via t_lido")
	if err := i.dbClient.IdentifyLidoValidators(); err != nil {
		return fmt.Errorf("propagating CM v2 tags: %w", err)
	}
	log.Info("CM v2 identification finished")
	return nil
}

// processCMv2Group handles a single group: enumerates its sub-operators, pulls
// each one's signing keys from the Curated v2 contract, writes them into
// t_lido with f_group_id/f_group_name set, and tags any v1 NOR rows pointed
// at by externalOperators[] with the same group.
func (i *Identify) processCMv2Group(client *cmv2.Client, groupID int64) error {
	group, err := client.GetOperatorGroup(groupID)
	if err != nil {
		return fmt.Errorf("getOperatorGroup(%d): %w", groupID, err)
	}
	groupName := cmv2.FormatGroupName(groupID)
	log.Infof("CM v2 group %d (%s): sub_operators=%d external_operators=%d",
		groupID, groupName, len(group.SubNodeOperators), len(group.ExternalOperators))

	// 1. Sub-operators native to CM v2 → fetch keys + persist with new protocol tag.
	for _, sub := range group.SubNodeOperators {
		if i.stop {
			return nil
		}
		if err := i.processCMv2SubOperator(client, groupID, groupName, sub.NodeOperatorId); err != nil {
			log.Errorf("CM v2 sub-operator %d (group %d): %v", sub.NodeOperatorId, groupID, err)
			// Continue with the rest of the group; failures are per-operator.
		}
	}

	// 2. External operators → propagate the group tag onto existing v1 Curated rows.
	v1Indices := make([]uint64, 0, len(group.ExternalOperators))
	for _, ext := range group.ExternalOperators {
		modID, opID, ok := cmv2.DecodeExternalOperator(ext.Data)
		if !ok {
			log.Warnf("CM v2 group %d: external operator data too short (%d bytes), skipping", groupID, len(ext.Data))
			continue
		}
		if modID != cmv2.CuratedV1StakingModuleID {
			log.Warnf("CM v2 group %d: external operator references unsupported moduleId=%d (operatorId=%d), skipping until we add support",
				groupID, modID, opID)
			continue
		}
		v1Indices = append(v1Indices, opID)
	}
	if len(v1Indices) > 0 {
		updated, err := i.dbClient.TagCuratedV1OperatorsWithGroup(v1Indices, groupID, groupName)
		if err != nil {
			return fmt.Errorf("tagging v1 operators with group %d: %w", groupID, err)
		}
		log.Infof("CM v2 group %d: tagged %d existing v1 Curated rows (operators=%v)",
			groupID, updated, v1Indices)
	}
	return nil
}

// processCMv2SubOperator reads the on-chain key set for a single sub-operator
// inside a group and writes the keys to t_lido. We page through the keys with
// the same maxBatchSize ceiling we use for the v1 Curated module.
func (i *Identify) processCMv2SubOperator(client *cmv2.Client, groupID int64, groupName string, nodeOperatorID uint64) error {
	op, err := client.GetNodeOperator(nodeOperatorID)
	if err != nil {
		return fmt.Errorf("getNodeOperator(%d): %w", nodeOperatorID, err)
	}
	meta, err := client.GetOperatorMetadata(nodeOperatorID)
	if err != nil {
		return fmt.Errorf("getOperatorMetadata(%d): %w", nodeOperatorID, err)
	}
	operatorTag := cmv2.FormatOperatorName(nodeOperatorID, meta.Name)

	totalKeys := uint64(op.TotalAddedKeys)
	log.Infof("CM v2 group %d operator %d (%q → tag=%s): on-chain total keys=%d",
		groupID, nodeOperatorID, meta.Name, operatorTag, totalKeys)
	if totalKeys == 0 {
		return nil
	}

	keys := make([]string, 0, totalKeys)
	offset := uint64(0)
	for offset < totalKeys {
		if i.stop {
			return nil
		}
		batch := totalKeys - offset
		if batch > maxBatchSize {
			batch = maxBatchSize
		}
		raw, err := client.GetSigningKeys(nodeOperatorID, offset, batch)
		if err != nil {
			return fmt.Errorf("getSigningKeys(op=%d, off=%d, lim=%d): %w", nodeOperatorID, offset, batch, err)
		}
		for k := uint64(0); k < batch; k++ {
			keys = append(keys, hex.EncodeToString(raw[k*lido.PublicKeyLength:(k+1)*lido.PublicKeyLength]))
		}
		offset += batch
	}

	inserted, err := i.dbClient.CopyLidoCMv2OperatorValidators(operatorTag, nodeOperatorID, groupID, groupName, keys)
	if err != nil {
		return fmt.Errorf("persisting CM v2 keys (op=%d, group=%d): %w", nodeOperatorID, groupID, err)
	}
	log.Infof("CM v2 group %d operator %d: persisted %d keys (tag=%s)", groupID, nodeOperatorID, inserted, operatorTag)
	return nil
}

// Avoid unused-import lint when db is only used through dbClient.
var _ = db.LidoProtocolCMv2
