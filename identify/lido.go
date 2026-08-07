package identify

import (
	"encoding/hex"
	"math/big"
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

func (i *Identify) IdentifyLidoValidators() error {
	// One snapshot of stored key counts shared by the three modules. Operators
	// are visited once per run, so intra-run staleness is not a concern.
	keyCounts, err := i.dbClient.ObtainLidoOperatorKeyCounts()
	if err != nil {
		return err
	}

	log.Debug("Identifying lido curated module validators")
	err = i.identifyCuratedModule(keyCounts)
	if err != nil {
		return err
	}
	log.Debug("Identified lido curated module validators")

	log.Debug("Identifying lido SDVT module validators")
	err = i.identifySDVT(keyCounts)
	if err != nil {
		return err
	}
	log.Debug("Identified lido SDVT module validators")

	log.Debug("Identifying lido csm validators")
	err = i.identifyCSM(keyCounts)
	if err != nil {
		return err
	}
	log.Debug("Identified lido csm validators")

	// One single pass over t_lido after the three modules updated it: the
	// UPDATE joins half a million keys against t_identified_validators, so
	// running it per module tripled the cost for the same final state.
	startTime := time.Now()
	err = i.dbClient.IdentifyLidoValidators()
	if err != nil {
		return err
	}
	log.Infof("Applied lido pool names in %v", time.Since(startTime))

	return nil
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

///// CSM /////

func (i *Identify) identifyCSM(keyCounts map[string]uint64) error {
	log.Debug("Creating a new instance of LidoContract")
	csmContract, err := csm.NewCSMContract(i.iConfig.ElEndpoint)
	// Check if there was an error
	if err != nil {
		return err
	}

	log.Debug("Calling the GetOperatorsIndexes function")
	operatorsCount, err := csmContract.GetNodeOperatorsCount()
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
			if err := i.processCSMOperatorKeys(operatorIndex, keyCounts, stats); err != nil {
				log.Errorf("Error reconciling CSM operator %v: %v", operatorIndex, err)
			}
			<-workerSemaphore
		}(operatorIndex)
	}
	log.Debug("Finished getting keys for each operator")

	wg.Wait()
	if i.stop {
		return nil
	}
	log.Infof("CSM module: operators=%d skipped=%d incremental=%d full=%d",
		operatorsCount, stats.skipped.Load(), stats.incremental.Load(), stats.full.Load())

	return nil
}

func (i *Identify) processCSMOperatorKeys(operatorIndex int64, keyCounts map[string]uint64, stats *lidoModuleStats) error {
	csmContract, err := csm.NewCSMContract(i.iConfig.ElEndpoint)
	if err != nil {
		return err
	}

	operator, err := csmContract.GetOperatorData(big.NewInt(operatorIndex))
	if err != nil {
		return err
	}

	operatorName := csm.GetOperatorName(operator)
	totalKeys := uint64(operator.Operator.TotalDepositedKeys)

	action, offset := i.lidoOperatorAction(db.LidoProtocolCSM, operator.Index, totalKeys, keyCounts, stats)
	if action == "skip" {
		log.Debugf("CSM operator %v: up to date (on-chain=%d)", operatorName, totalKeys)
		return nil
	}
	log.Infof("Reconciling keys for CSM operator %v (action=%s on-chain total=%v)", operatorName, action, totalKeys)

	keysString := make([]string, 0, totalKeys-offset)
	if totalKeys > offset {
		keys, err := csmContract.GetOperatorKeys(operator, offset, totalKeys-offset)
		if err != nil {
			return err
		}
		for _, key := range keys {
			keysString = append(keysString, hex.EncodeToString(key))
		}
	}

	if action == "incremental" {
		upserted, err := i.dbClient.AppendLidoOperatorValidators(operatorName, operator.Index, keysString, db.LidoProtocolCSM)
		if err != nil {
			return err
		}
		log.Infof("CSM operator %v appended: on-chain=%d new=%d", operatorName, totalKeys, upserted)
		return nil
	}

	upserted, removed, err := i.dbClient.ReconcileLidoOperatorValidators(operatorName, operator.Index, keysString, db.LidoProtocolCSM)
	if err != nil {
		return err
	}
	log.Infof("CSM operator %v reconciled: on-chain=%d upserted=%d removed=%d", operatorName, totalKeys, upserted, removed)
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
