package identify

import (
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"sort"
	"sync"
	"sync/atomic"
	"time"

	db "github.com/migalabs/eth-pokhar/db"
	"github.com/migalabs/eth-pokhar/lido"
	"github.com/migalabs/eth-pokhar/lido/csm"
	"github.com/migalabs/eth-pokhar/lido/curated"
	"github.com/migalabs/eth-pokhar/lido/sdvt"
	log "github.com/sirupsen/logrus"
)

const maxBatchSize = 500

// lidoModuleStats aggregates per-operator outcomes of a module pass.
type lidoModuleStats struct {
	skipped     atomic.Int64
	incremental atomic.Int64
	full        atomic.Int64
}

// lidoOperatorAction decides how to treat an operator by comparing the number
// of keys already stored in t_lido against the on-chain total (which comes for
// free with the operator data). Key registries are append-mostly, so most
// operators need no work at all; fewer on-chain keys than stored means keys
// were removed (the curated module compacts indexes on removal), which forces
// a full refetch so the reconcile can drop fossils. Full runs (recreate-table)
// force the full path for every operator: that weekly pass also absorbs
// operator renames and any drift the incremental path cannot see.
func (i *Identify) lidoOperatorAction(protocol string, operatorIndex uint64, onChainTotal uint64, keyCounts map[string]uint64, stats *lidoModuleStats) (string, uint64) {
	if !i.iConfig.RecreateTable {
		dbCount := keyCounts[db.LidoOperatorCountKey(protocol, operatorIndex)]
		if dbCount == onChainTotal {
			stats.skipped.Add(1)
			return "skip", 0
		}
		if dbCount < onChainTotal {
			stats.incremental.Add(1)
			return "incremental", dbCount
		}
	}
	stats.full.Add(1)
	return "full", 0
}

// csmFamilyModule is a Lido staking module read through the CSM binding: CSM
// itself and the modules built on its codebase.
type csmFamilyModule struct {
	id           uint64 // StakingRouter module id
	name         string
	address      string
	protocol     string // t_lido.f_protocol
	operatorName func(csm.NodeOperatorCustom) string
	// entityTag, when set, is the f_pool_name of every validator of the
	// module; the per-operator tag still feeds f_operator. Empty means the
	// per-operator tag is also the entity label.
	entityTag string
}

var csmFamilyModules = []csmFamilyModule{
	{id: 3, name: "CSM", address: csm.CSMContractAddress, protocol: db.LidoProtocolCSM, operatorName: csm.GetOperatorName},
	{id: 4, name: "CMv2", address: csm.CMv2ContractAddress, protocol: db.LidoProtocolCMv2, operatorName: csm.CMv2OperatorName, entityTag: csm.CMv2EntityTag},
	{id: 5, name: "CSM 0x02", address: csm.CSM0x02ContractAddress, protocol: db.LidoProtocolCSM0x02, operatorName: csm.CSM0x02OperatorName, entityTag: csm.CSM0x02EntityTag},
}

// scannedStakingModules maps every StakingRouter module id pokhar reads to its
// name: curated and SDVT, plus the CSM family.
func scannedStakingModules() map[uint64]string {
	scanned := map[uint64]string{1: "curated", 2: "SDVT"}
	for _, module := range csmFamilyModules {
		scanned[module.id] = module.name
	}
	return scanned
}

// lidoEntityTags returns, as parallel slices, the protocols whose validators
// are labelled with a module-wide entity tag and that tag.
func lidoEntityTags() (protocols []string, tags []string) {
	for _, module := range csmFamilyModules {
		if module.entityTag != "" {
			protocols = append(protocols, module.protocol)
			tags = append(tags, module.entityTag)
		}
	}
	return protocols, tags
}

// lidoDimensions returns the Lido sources of the f_custodian and f_operator
// dimensions: the custody entity of every module (Lido is shown as one custody
// entity per module) and the canonical operator of every curated v1 and CMv2
// node operator. SDVT clusters and CSM operators keep their registry tag as
// operator. Static on purpose, so the dimension pass is deterministic and
// never depends on an RPC answer; warnCMv2OperatorDrift keeps it honest.
func lidoDimensions() db.LidoDimensions {
	d := db.LidoDimensions{
		CustodyProtocols: []string{db.LidoProtocolCurated, db.LidoProtocolSDVT, db.LidoProtocolCSM, db.LidoProtocolCMv2, db.LidoProtocolCSM0x02},
		CustodyTags:      []string{lido.CuratedCustodyTag, lido.SDVTCustodyTag, lido.CSMCustodyTag, csm.CMv2EntityTag, csm.CSM0x02EntityTag},
	}
	for index, operator := range lido.CuratedCanonicalOperators {
		d.OperatorProtocols = append(d.OperatorProtocols, db.LidoProtocolCurated)
		d.OperatorIndexes = append(d.OperatorIndexes, int32(index))
		d.OperatorTags = append(d.OperatorTags, operator)
	}
	cmv2Operators := make([]uint64, 0, len(lido.CMv2CuratedOperator))
	for op := range lido.CMv2CuratedOperator {
		cmv2Operators = append(cmv2Operators, op)
	}
	sort.Slice(cmv2Operators, func(a, b int) bool { return cmv2Operators[a] < cmv2Operators[b] })
	for _, op := range cmv2Operators {
		curated := lido.CMv2CuratedOperator[op]
		if curated >= uint64(len(lido.CuratedCanonicalOperators)) {
			continue
		}
		d.OperatorProtocols = append(d.OperatorProtocols, db.LidoProtocolCMv2)
		d.OperatorIndexes = append(d.OperatorIndexes, int32(op))
		d.OperatorTags = append(d.OperatorTags, lido.CuratedCanonicalOperators[curated])
	}
	return d
}

// warnCMv2OperatorDrift compares the static CMv2 -> curated v1 operator links
// with the MetaRegistry groups and warns about new operators and changed
// links, which would otherwise leave a company split across two operator
// labels without anyone noticing. Operators the MetaRegistry puts in no group
// are expected to be linked by hand (see lido.CMv2CuratedOperator). Best
// effort: a failed check never blocks the run.
func (i *Identify) warnCMv2OperatorDrift() {
	module, err := csm.NewContractAt(i.iConfig.ElEndpoint, csm.CMv2ContractAddress)
	if err != nil {
		log.Warnf("Could not check CMv2 operator links: %v", err)
		return
	}
	count, err := module.GetNodeOperatorsCount()
	if err != nil {
		log.Warnf("Could not check CMv2 operator links: %v", err)
		return
	}
	onChain, err := lido.CMv2CuratedLinksOnChain(i.iConfig.ElEndpoint, uint64(count))
	if err != nil {
		log.Warnf("Could not check CMv2 operator links: %v", err)
		return
	}
	for op := uint64(0); op < uint64(count); op++ {
		static, known := lido.CMv2CuratedOperator[op]
		chain, grouped := onChain[op]
		switch {
		case !known && grouped:
			log.Warnf("CMv2 operator %d is not in lido.CMv2CuratedOperator; the MetaRegistry links it to curated operator %d. It keeps the tag lido_cmv2_operator%d until the table is updated", op, chain, op)
		case !known:
			log.Warnf("CMv2 operator %d is not in lido.CMv2CuratedOperator and the MetaRegistry links it to no curated operator: it keeps the tag lido_cmv2_operator%d", op, op)
		case grouped && chain != static:
			log.Warnf("CMv2 operator %d is linked to curated operator %d in lido.CMv2CuratedOperator but the MetaRegistry says %d: confirm and update the table", op, static, chain)
		}
	}
}

// warnUnscannedStakingModules flags StakingRouter modules pokhar does not read.
// Their keys never reach t_lido, so their validators fall back to the
// StakingRouter depositor label "lido" (issue #47); this is the tripwire for
// the next module Lido adds. Best effort: a failed check never blocks the run.
func (i *Identify) warnUnscannedStakingModules() {
	ids, err := lido.GetStakingModuleIds(i.iConfig.ElEndpoint)
	if err != nil {
		log.Warnf("Could not list Lido staking modules: %v", err)
		return
	}
	for _, id := range lido.UnscannedModules(ids, scannedStakingModules()) {
		log.Warnf("Lido StakingRouter has module %d, which pokhar does not scan: its validators fall back to the generic \"lido\" label", id)
	}
}

func (i *Identify) IdentifyLidoValidators() error {
	i.warnUnscannedStakingModules()
	i.warnCMv2OperatorDrift()

	// One snapshot of stored key counts shared by every module. Operators are
	// visited once per run, so intra-run staleness is not a concern.
	keyCounts, err := i.dbClient.ObtainLidoOperatorKeyCounts()
	if err != nil {
		return err
	}

	// A module that fails must not cost the others their labels. The keys of
	// earlier runs stay in t_lido, so the UPDATE below is still correct for a
	// module that could not be refreshed. Skipping it would be much worse: in
	// the weekly RECREATE_TABLE run t_identified_validators starts truncated
	// and the depositor phases leave t_lido keys alone, so every Lido
	// validator would stay solo_stakers until the next run.
	var moduleErrs []error
	log.Debug("Identifying lido curated module validators")
	if err := i.identifyCuratedModule(keyCounts); err != nil {
		moduleErrs = append(moduleErrs, fmt.Errorf("curated: %w", err))
	}

	log.Debug("Identifying lido SDVT module validators")
	if err := i.identifySDVT(keyCounts); err != nil {
		moduleErrs = append(moduleErrs, fmt.Errorf("SDVT: %w", err))
	}

	for _, module := range csmFamilyModules {
		log.Debugf("Identifying lido %s validators", module.name)
		if err := i.identifyCSMFamily(module, keyCounts); err != nil {
			moduleErrs = append(moduleErrs, fmt.Errorf("%s: %w", module.name, err))
		}
	}

	// One single pass over t_lido after every module updated it: the UPDATE
	// joins half a million keys against t_identified_validators, so running it
	// per module multiplied the cost for the same final state.
	startTime := time.Now()
	entityProtocols, entityTags := lidoEntityTags()
	if err := i.dbClient.IdentifyLidoValidators(entityProtocols, entityTags); err != nil {
		return errors.Join(append(moduleErrs, err)...)
	}
	log.Infof("Applied lido pool names in %v", time.Since(startTime))

	return errors.Join(moduleErrs...)
}

///// Simple DVT /////

// identifySDVT identifies the validators for the Simple DVT module. Same key
// registry semantics as the Curated module: only the contract address and
// operator name resolution differ. Operator names come from the on-chain
// contract (no pre-defined override list).
func (i *Identify) identifySDVT(keyCounts map[string]uint64) error {
	log.Debug("Creating a new instance of SDVT contract")
	sdvtContract, err := sdvt.NewSDVTContract(i.iConfig.ElEndpoint)
	if err != nil {
		return err
	}

	operatorsCount, err := sdvtContract.GetNodeOperatorsCount()
	if err != nil {
		return err
	}
	log.Debugf("Found %v SDVT operators", operatorsCount)

	stats := &lidoModuleStats{}
	workerSemaphore := make(chan struct{}, 10)
	var wg sync.WaitGroup

	for operatorIndex := int64(0); operatorIndex < operatorsCount; operatorIndex++ {
		if i.stop {
			break
		}
		wg.Add(1)
		workerSemaphore <- struct{}{}

		go func(operatorIndex int64) {
			defer wg.Done()
			if err := i.processSDVTOperatorKeys(operatorIndex, keyCounts, stats); err != nil {
				log.Errorf("Error reconciling SDVT operator %v: %v", operatorIndex, err)
			}
			<-workerSemaphore
		}(operatorIndex)
	}

	wg.Wait()
	if i.stop {
		return nil
	}
	log.Infof("SDVT module: operators=%d skipped=%d incremental=%d full=%d",
		operatorsCount, stats.skipped.Load(), stats.incremental.Load(), stats.full.Load())

	return nil
}

func (i *Identify) processSDVTOperatorKeys(operatorIndex int64, keyCounts map[string]uint64, stats *lidoModuleStats) error {
	sdvtContract, err := sdvt.NewSDVTContract(i.iConfig.ElEndpoint)
	if err != nil {
		return err
	}

	operator, err := sdvtContract.GetOperatorData(big.NewInt(operatorIndex))
	if err != nil {
		return err
	}

	operatorName := sdvt.GetOperatorName(operator)
	totalKeys := operator.TotalSigningKeys

	action, offset := i.lidoOperatorAction(db.LidoProtocolSDVT, operator.Index, totalKeys, keyCounts, stats)
	if action == "skip" {
		log.Debugf("SDVT operator %v: up to date (on-chain=%d)", operatorName, totalKeys)
		return nil
	}
	log.Infof("Reconciling keys for SDVT operator %v (action=%s on-chain total=%v)", operatorName, action, totalKeys)

	validatorPubkeys := make([]string, 0, totalKeys-offset)
	for offset < totalKeys {
		if i.stop {
			return nil
		}
		limit := totalKeys - offset
		if limit > maxBatchSize {
			limit = maxBatchSize
		}

		operatorKeys, err := sdvtContract.GetOperatorKeys(operator, offset, limit)
		if err != nil {
			return err
		}
		for k := uint64(0); k < limit; k++ {
			key := operatorKeys.PubKeys[k*lido.PublicKeyLength : (k+1)*lido.PublicKeyLength]
			validatorPubkeys = append(validatorPubkeys, hex.EncodeToString(key))
		}
		offset += limit
	}

	if action == "incremental" {
		upserted, err := i.dbClient.AppendLidoOperatorValidators(operatorName, operator.Index, validatorPubkeys, db.LidoProtocolSDVT)
		if err != nil {
			return err
		}
		log.Infof("SDVT operator %v appended: on-chain=%d new=%d", operatorName, totalKeys, upserted)
		return nil
	}

	upserted, removed, err := i.dbClient.ReconcileLidoOperatorValidators(operatorName, operator.Index, validatorPubkeys, db.LidoProtocolSDVT)
	if err != nil {
		return err
	}
	log.Infof("SDVT operator %v reconciled: on-chain=%d upserted=%d removed=%d", operatorName, totalKeys, upserted, removed)
	return nil
}

///// CSM family (CSM, Curated Module v2, CSM 0x02) /////

// identifyCSMFamily reconciles the deposited keys of every operator of a module
// that shares the CSM interface.
func (i *Identify) identifyCSMFamily(module csmFamilyModule, keyCounts map[string]uint64) error {
	log.Debugf("Creating a new instance of the %s contract", module.name)
	moduleContract, err := csm.NewContractAt(i.iConfig.ElEndpoint, module.address)
	if err != nil {
		return err
	}

	operatorsCount, err := moduleContract.GetNodeOperatorsCount()
	if err != nil {
		return err
	}
	log.Debugf("Found %v %s operators", operatorsCount, module.name)

	stats := &lidoModuleStats{}
	workerSemaphore := make(chan struct{}, 10)
	var wg sync.WaitGroup

	for operatorIndex := int64(0); operatorIndex < operatorsCount; operatorIndex++ {
		if i.stop {
			break
		}
		wg.Add(1)
		workerSemaphore <- struct{}{}

		go func(operatorIndex int64) {
			defer wg.Done()
			if err := i.processCSMFamilyOperatorKeys(module, operatorIndex, keyCounts, stats); err != nil {
				log.Errorf("Error reconciling %s operator %v: %v", module.name, operatorIndex, err)
			}
			<-workerSemaphore
		}(operatorIndex)
	}

	wg.Wait()
	if i.stop {
		return nil
	}
	log.Infof("%s module: operators=%d skipped=%d incremental=%d full=%d",
		module.name, operatorsCount, stats.skipped.Load(), stats.incremental.Load(), stats.full.Load())

	return nil
}

func (i *Identify) processCSMFamilyOperatorKeys(module csmFamilyModule, operatorIndex int64, keyCounts map[string]uint64, stats *lidoModuleStats) error {
	moduleContract, err := csm.NewContractAt(i.iConfig.ElEndpoint, module.address)
	if err != nil {
		return err
	}

	operator, err := moduleContract.GetOperatorData(big.NewInt(operatorIndex))
	if err != nil {
		return err
	}

	operatorName := module.operatorName(operator)
	totalKeys := uint64(operator.Operator.TotalDepositedKeys)

	action, offset := i.lidoOperatorAction(module.protocol, operator.Index, totalKeys, keyCounts, stats)
	if action == "skip" {
		log.Debugf("%s operator %v: up to date (on-chain=%d)", module.name, operatorName, totalKeys)
		return nil
	}
	log.Infof("Reconciling keys for %s operator %v (action=%s on-chain total=%v)", module.name, operatorName, action, totalKeys)

	// Batched like the curated path: CMv2 operators absorb whole curated
	// fleets through consolidations, so one call per operator would not scale.
	keysString := make([]string, 0, totalKeys-offset)
	for offset < totalKeys {
		if i.stop {
			return nil
		}
		limit := totalKeys - offset
		if limit > maxBatchSize {
			limit = maxBatchSize
		}
		keys, err := moduleContract.GetOperatorKeys(operator, offset, limit)
		if err != nil {
			return err
		}
		for _, key := range keys {
			keysString = append(keysString, hex.EncodeToString(key))
		}
		offset += limit
	}

	if action == "incremental" {
		upserted, err := i.dbClient.AppendLidoOperatorValidators(operatorName, operator.Index, keysString, module.protocol)
		if err != nil {
			return err
		}
		log.Infof("%s operator %v appended: on-chain=%d new=%d", module.name, operatorName, totalKeys, upserted)
		return nil
	}

	upserted, removed, err := i.dbClient.ReconcileLidoOperatorValidators(operatorName, operator.Index, keysString, module.protocol)
	if err != nil {
		return err
	}
	log.Infof("%s operator %v reconciled: on-chain=%d upserted=%d removed=%d", module.name, operatorName, totalKeys, upserted, removed)
	return nil
}

////// Curated Module //////

// identifyCuratedModule identifies the validators for the curated module and adds them to the lido table
func (i *Identify) identifyCuratedModule(keyCounts map[string]uint64) error {
	log.Debug("Creating a new instance of LidoContract")
	lidoContract, err := curated.NewCuratedModuleContract(i.iConfig.ElEndpoint)
	// Check if there was an error
	if err != nil {
		return err
	}
	log.Debug("Created a new instance of LidoContract")

	log.Debug("Calling the GetOperatorsIndexes function")
	operatorsCount, err := lidoContract.GetNodeOperatorsCount()
	if err != nil {
		return err
	}
	log.Debugf("Found %v operators", operatorsCount)

	log.Debug("Getting keys for each operator")
	stats := &lidoModuleStats{}
	workerSemaphore := make(chan struct{}, 10)
	var wg sync.WaitGroup

	for operatorIndex := int64(0); operatorIndex < operatorsCount; operatorIndex++ {
		if i.stop {
			break
		}
		wg.Add(1)
		workerSemaphore <- struct{}{}

		go func(operatorIndex int64) {
			defer wg.Done()
			if err := i.processCuratedOperatorKeys(operatorIndex, keyCounts, stats); err != nil {
				log.Errorf("Error reconciling curated operator %v: %v", operatorIndex, err)
			}
			<-workerSemaphore
		}(operatorIndex)
	}
	log.Debug("Finished getting keys for each operator")

	wg.Wait()
	if i.stop {
		return nil
	}
	log.Infof("Curated module: operators=%d skipped=%d incremental=%d full=%d",
		operatorsCount, stats.skipped.Load(), stats.incremental.Load(), stats.full.Load())

	return nil
}

func (i *Identify) processCuratedOperatorKeys(operatorIndex int64, keyCounts map[string]uint64, stats *lidoModuleStats) error {
	lidoContract, err := curated.NewCuratedModuleContract(i.iConfig.ElEndpoint)
	if err != nil {
		return err
	}

	operator, err := lidoContract.GetOperatorData(big.NewInt(operatorIndex))
	if err != nil {
		return err
	}

	operatorName := curated.GetOperatorName(operator)
	totalKeys := operator.TotalSigningKeys

	action, offset := i.lidoOperatorAction(db.LidoProtocolCurated, operator.Index, totalKeys, keyCounts, stats)
	if action == "skip" {
		log.Debugf("Curated operator %v: up to date (on-chain=%d)", operatorName, totalKeys)
		return nil
	}
	log.Infof("Reconciling keys for curated operator %v (action=%s on-chain total=%v)", operatorName, action, totalKeys)

	validatorPubkeys := make([]string, 0, totalKeys-offset)
	for offset < totalKeys {
		if i.stop {
			return nil
		}
		limit := totalKeys - offset
		if limit > maxBatchSize {
			limit = maxBatchSize
		}

		operatorKeys, err := lidoContract.GetOperatorKeys(operator, offset, limit)
		if err != nil {
			return err
		}
		for k := uint64(0); k < limit; k++ {
			key := operatorKeys.PubKeys[k*lido.PublicKeyLength : (k+1)*lido.PublicKeyLength]
			validatorPubkeys = append(validatorPubkeys, hex.EncodeToString(key))
		}
		offset += limit
	}

	if action == "incremental" {
		upserted, err := i.dbClient.AppendLidoOperatorValidators(operatorName, operator.Index, validatorPubkeys, db.LidoProtocolCurated)
		if err != nil {
			return err
		}
		log.Infof("Curated operator %v appended: on-chain=%d new=%d", operatorName, totalKeys, upserted)
		return nil
	}

	upserted, removed, err := i.dbClient.ReconcileLidoOperatorValidators(operatorName, operator.Index, validatorPubkeys, db.LidoProtocolCurated)
	if err != nil {
		return err
	}
	log.Infof("Curated operator %v reconciled: on-chain=%d upserted=%d removed=%d", operatorName, totalKeys, upserted, removed)
	return nil
}
