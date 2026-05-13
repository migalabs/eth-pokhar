package identify

import (
	"encoding/hex"
	"math/big"
	"sync"

	db "github.com/migalabs/eth-pokhar/db"
	"github.com/migalabs/eth-pokhar/lido"
	"github.com/migalabs/eth-pokhar/lido/csm"
	"github.com/migalabs/eth-pokhar/lido/curated"
	"github.com/migalabs/eth-pokhar/lido/sdvt"
	log "github.com/sirupsen/logrus"
)

const maxBatchSize = 500

func (i *Identify) IdentifyLidoValidators() error {
	log.Debug("Identifying lido curated module validators")
	err := i.identifyCuratedModule()
	if err != nil {
		return err
	}
	log.Debug("Identified lido curated module validators")

	log.Debug("Identifying lido SDVT module validators")
	err = i.identifySDVT()
	if err != nil {
		return err
	}
	log.Debug("Identified lido SDVT module validators")

	log.Debug("Identifying lido csm validators")
	err = i.identifyCSM()
	if err != nil {
		return err
	}
	log.Debug("Identified lido csm validators")

	return nil
}

///// Simple DVT /////

// identifySDVT identifies the validators for the Simple DVT module. Same key
// registry semantics as the Curated module — only the contract address and
// operator name resolution differ. Operator names come from the on-chain
// contract (no pre-defined override list).
func (i *Identify) identifySDVT() error {
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

	workerSemaphore := make(chan struct{}, 10)
	var wg sync.WaitGroup

	for operatorIndex := int64(0); operatorIndex < operatorsCount; operatorIndex++ {
		if i.stop {
			break
		}
		wg.Add(1)
		workerSemaphore <- struct{}{}
		operator, err := sdvtContract.GetOperatorData(big.NewInt(operatorIndex))
		if err != nil {
			return err
		}

		go func(operator curated.NodeOperator) {
			defer wg.Done()
			if err := i.processSDVTOperatorKeys(operator); err != nil {
				log.Errorf("Error reconciling SDVT operator %v: %v", operator.Index, err)
			}
			<-workerSemaphore
		}(operator)
	}

	wg.Wait()
	if i.stop {
		return nil
	}

	err = i.dbClient.IdentifyLidoValidators()
	if err != nil {
		return err
	}
	return nil
}

func (i *Identify) processSDVTOperatorKeys(operator curated.NodeOperator) error {
	sdvtContract, err := sdvt.NewSDVTContract(i.iConfig.ElEndpoint)
	if err != nil {
		return err
	}

	operatorName := sdvt.GetOperatorName(operator)
	totalKeys := operator.TotalSigningKeys
	log.Infof("Reconciling keys for SDVT operator %v (on-chain total=%v)", operatorName, totalKeys)

	validatorPubkeys := make([]string, 0, totalKeys)
	offset := uint64(0)
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

	upserted, removed, renamed, err := i.dbClient.ReconcileLidoOperatorValidators(operatorName, operator.Index, validatorPubkeys, db.LidoProtocolSDVT)
	if err != nil {
		return err
	}
	log.Infof("SDVT operator %v reconciled: on-chain=%d upserted=%d removed=%d renamed=%d", operatorName, totalKeys, upserted, removed, renamed)
	return nil
}

///// CSM /////

func (i *Identify) identifyCSM() error {
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
	workerSemaphore := make(chan struct{}, 10)
	var wg sync.WaitGroup

	for operatorIndex := int64(0); operatorIndex < operatorsCount; operatorIndex++ {
		if i.stop {
			break
		}
		wg.Add(1)
		workerSemaphore <- struct{}{}
		operator, err := csmContract.GetOperatorData(big.NewInt(operatorIndex))
		if err != nil {
			return err
		}

		go func(operator csm.NodeOperatorCustom) {
			defer wg.Done()
			err := i.processCSMOperatorKeys(operator)
			if err != nil {
				log.Fatalf("Error processing operator keys: %v", err)
			}
			<-workerSemaphore
		}(operator)
	}
	log.Debug("Finished getting keys for each operator")

	wg.Wait()
	if i.stop {
		return nil
	}

	log.Debug("Identifying lido curated validators")
	err = i.dbClient.IdentifyLidoValidators()
	if err != nil {
		return err
	}
	log.Debug("Identified lido validators")

	return nil
}

func (i *Identify) processCSMOperatorKeys(operator csm.NodeOperatorCustom) error {
	operatorName := csm.GetOperatorName(operator)
	totalKeys := uint64(operator.Operator.TotalDepositedKeys)
	log.Infof("Reconciling keys for CSM operator %v (on-chain total=%v)", operatorName, totalKeys)

	keysString := make([]string, 0, totalKeys)
	if totalKeys > 0 {
		lidoContract, err := csm.NewCSMContract(i.iConfig.ElEndpoint)
		if err != nil {
			return err
		}
		keys, err := lidoContract.GetOperatorKeys(operator, 0, totalKeys)
		if err != nil {
			return err
		}
		for _, key := range keys {
			keysString = append(keysString, hex.EncodeToString(key))
		}
	}

	upserted, removed, renamed, err := i.dbClient.ReconcileLidoOperatorValidators(operatorName, operator.Index, keysString, db.LidoProtocolCSM)
	if err != nil {
		return err
	}
	log.Infof("CSM operator %v reconciled: on-chain=%d upserted=%d removed=%d renamed=%d", operatorName, totalKeys, upserted, removed, renamed)
	return nil
}

////// Curated Module //////

// identifyCuratedModule identifies the validators for the curated module and adds them to the lido table
func (i *Identify) identifyCuratedModule() error {
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
	workerSemaphore := make(chan struct{}, 10)
	var wg sync.WaitGroup

	for operatorIndex := int64(0); operatorIndex < operatorsCount; operatorIndex++ {
		if i.stop {
			break
		}
		wg.Add(1)
		workerSemaphore <- struct{}{}
		operator, err := lidoContract.GetOperatorData(big.NewInt(operatorIndex))
		if err != nil {
			return err
		}

		go func(operator curated.NodeOperator) {
			defer wg.Done()
			if err := i.processCuratedOperatorKeys(operator); err != nil {
				log.Errorf("Error reconciling curated operator %v: %v", operator.Index, err)
			}
			<-workerSemaphore
		}(operator)
	}
	log.Debug("Finished getting keys for each operator")

	wg.Wait()
	if i.stop {
		return nil
	}

	log.Debug("Identifying lido curated validators")
	err = i.dbClient.IdentifyLidoValidators()
	if err != nil {
		return err
	}
	log.Debug("Identified lido validators")

	return nil
}

func (i *Identify) processCuratedOperatorKeys(operator curated.NodeOperator) error {
	lidoContract, err := curated.NewCuratedModuleContract(i.iConfig.ElEndpoint)
	if err != nil {
		return err
	}

	operatorName := curated.GetOperatorName(operator)
	totalKeys := operator.TotalSigningKeys
	log.Infof("Reconciling keys for curated operator %v (on-chain total=%v)", operatorName, totalKeys)

	validatorPubkeys := make([]string, 0, totalKeys)
	offset := uint64(0)
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

	upserted, removed, renamed, err := i.dbClient.ReconcileLidoOperatorValidators(operatorName, operator.Index, validatorPubkeys, db.LidoProtocolCurated)
	if err != nil {
		return err
	}
	log.Infof("Curated operator %v reconciled: on-chain=%d upserted=%d removed=%d renamed=%d", operatorName, totalKeys, upserted, removed, renamed)
	return nil
}
