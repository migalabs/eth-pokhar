package identify

import (
	"encoding/hex"
	"math/big"
	"sync"

	db "github.com/migalabs/eth-pokhar/db"
	"github.com/migalabs/eth-pokhar/lido"
	"github.com/migalabs/eth-pokhar/lido/csm"
	"github.com/migalabs/eth-pokhar/lido/curated"
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

	log.Debug("Identifying lido csm validators")
	err = i.identifyCSM()
	if err != nil {
		return err
	}
	log.Debug("Identified lido csm validators")

	return nil
}

///// CSM /////

func (i *Identify) identifyCSM() error {
	savedOperatorsValidatorsCount, err := i.dbClient.ObtainLidoOperatorsValidatorCount(db.LidoProtocolCSM)
	if err != nil {
		return err
	}

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

		savedOperatorValidatorCount := uint64(0)
		if operator.Index < uint64(len(savedOperatorsValidatorsCount)) {
			savedOperatorValidatorCount = savedOperatorsValidatorsCount[operator.Index]
		}
		go func(operator csm.NodeOperatorCustom, savedOperatorValidatorCount uint64) {
			defer wg.Done()
			err := i.processCSMOperatorKeys(operator, savedOperatorValidatorCount)
			if err != nil {
				log.Fatalf("Error processing operator keys: %v", err)
			}
			<-workerSemaphore
		}(operator, savedOperatorValidatorCount)
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

func (i *Identify) processCSMOperatorKeys(operator csm.NodeOperatorCustom, savedOperatorValidatorCount uint64) error {
	savedOperatorValidatorCount32 := uint32(savedOperatorValidatorCount)
	operatorName := csm.GetOperatorName(operator)
	log.Infof("Getting new keys for operator %v", operatorName)
	remainingKeys := operator.Operator.TotalDepositedKeys - savedOperatorValidatorCount32
	if remainingKeys == 0 {
		log.Infof("No new keys for operator %v", operatorName)
		return nil
	}

	log.Debug("Creating a new instance of LidoContract")
	lidoContract, err := csm.NewCSMContract(i.iConfig.ElEndpoint)
	// Check if there was an error
	if err != nil {
		return err
	}
	log.Debug("Created a new instance of LidoContract")
	keys, err := lidoContract.GetOperatorKeys(operator, uint64(savedOperatorValidatorCount32), uint64(remainingKeys))
	if err != nil {
		return err
	}
	keysString := make([]string, len(keys))
	for i, key := range keys {
		keysString[i] = hex.EncodeToString(key)
	}
	count := i.dbClient.CopyLidoOperatorValidators(operatorName, operator.Index, keysString, db.LidoProtocolCSM)
	log.Debugf("Inserted %v validators for operator %v", count, operatorName)
	log.Infof("Got %v new keys for operator %v", remainingKeys, operatorName)

	return nil
}

////// Curated Module //////

// identifyCuratedModule identifies the validators for the curated module and adds them to the lido table
func (i *Identify) identifyCuratedModule() error {

	operatorsValidatorCount, err := i.dbClient.ObtainLidoOperatorsValidatorCount(db.LidoProtocolCurated)
	if err != nil {
		return err
	}

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

		operatorValidatorCount := uint64(0)
		if operator.Index < uint64(len(operatorsValidatorCount)) {
			operatorValidatorCount = operatorsValidatorCount[operator.Index]
		}
		go func(operator curated.NodeOperator, operatorValidatorCount uint64) {
			defer wg.Done()
			i.processCuratedOperatorKeys(operator, operatorValidatorCount)
			<-workerSemaphore
		}(operator, operatorValidatorCount)
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

func (i *Identify) processCuratedOperatorKeys(operator curated.NodeOperator, operatorValidatorCount uint64) error {
	log.Debug("Creating a new instance of LidoContract")
	lidoContract, err := curated.NewCuratedModuleContract(i.iConfig.ElEndpoint)
	// Check if there was an error
	if err != nil {
		return err
	}
	log.Debug("Created a new instance of LidoContract")

	operatorName := curated.GetOperatorName(operator)
	log.Infof("Getting new keys for operator %v", operatorName)
	remainingKeys := operator.TotalSigningKeys - operatorValidatorCount
	if remainingKeys == 0 {
		log.Infof("No new keys for operator %v", operatorName)
		return nil
	}
	savedKeys := int64(0)
	var validatorPubkeys []string

	offset := operatorValidatorCount
	for {
		if i.stop {
			break
		}
		limit := min(remainingKeys-offset, maxBatchSize)
		validatorPubkeys = make([]string, limit)

		operatorKeys, err := lidoContract.GetOperatorKeys(operator, offset, limit)
		if err != nil {
			return err
		}
		for i := uint64(0); i < limit; i++ {
			key := operatorKeys.PubKeys[i*lido.PublicKeyLength : (i+1)*lido.PublicKeyLength]
			validatorPubkeys[i] = hex.EncodeToString(key)
		}

		log.Debugf("Inserting %v keys for operator %v into the database", limit, operatorName)
		count := i.dbClient.CopyLidoOperatorValidators(operatorName, operator.Index, validatorPubkeys, db.LidoProtocolCurated)
		log.Debugf("Inserted %v validators for operator %v. %v remaining", count, operatorName, remainingKeys)
		savedKeys += count
		done := limit < maxBatchSize
		if done {
			break
		}
		offset += limit
	}
	log.Infof("Got %v new keys for operator %v", savedKeys, operatorName)

	return nil
}
