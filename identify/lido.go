package identify

import (
	"encoding/hex"
	"math/big"
	"strings"
	"sync"

	"github.com/migalabs/eth-pokhar/lido"
	log "github.com/sirupsen/logrus"
)

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

const maxBatchSize = 100

func (i *Identify) IdentifyLidoValidators() error {

	operatorsValidatorCount, err := i.dbClient.ObtainLidoOperatorsValidatorCount()
	if err != nil {
		return err
	}

	log.Debug("Creating a new instance of LidoContract")
	lidoContract, err := lido.NewLidoContract(i.iConfig.ElEndpoint)
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
		go func(operator lido.NodeOperator, operatorValidatorCount uint64) {
			defer wg.Done()
			i.processOperatorKeys(operator, operatorValidatorCount)
			<-workerSemaphore
		}(operator, operatorValidatorCount)
	}
	log.Debug("Finished getting keys for each operator")

	wg.Wait()

	log.Debug("Identifying lido validators")
	err = i.dbClient.IdentifyLidoValidators()
	if err != nil {
		return err
	}
	log.Debug("Identified lido validators")

	return nil
}

func (i *Identify) processOperatorKeys(operator lido.NodeOperator, operatorValidatorCount uint64) error {
	log.Debug("Creating a new instance of LidoContract")
	lidoContract, err := lido.NewLidoContract(i.iConfig.ElEndpoint)
	// Check if there was an error
	if err != nil {
		return err
	}
	log.Debug("Created a new instance of LidoContract")

	operatorName := getOperatorName(operator)
	log.Infof("Getting new keys for operator %v", operatorName)
	if operator.TotalSigningKeys-operatorValidatorCount == 0 {
		log.Infof("No new keys for operator %v", operatorName)
		return nil
	}
	remainingKeys := operator.TotalSigningKeys - operatorValidatorCount
	var validatorPubkeys []string
	var batchSize uint64
	var batchIndex uint64
	for keyIndex := operatorValidatorCount; keyIndex < operator.TotalSigningKeys; keyIndex++ {
		if validatorPubkeys == nil {
			batchIndex = 0
			batchSize = min(remainingKeys, maxBatchSize)
			validatorPubkeys = make([]string, batchSize)
		}

		key, err := lidoContract.GetOperatorKey(operator, keyIndex)
		if err != nil {
			return err
		}
		validatorPubkeys[batchIndex] = hex.EncodeToString(key.Key)
		isLastKey := keyIndex == operator.TotalSigningKeys-1
		remainingKeys--
		if batchIndex == batchSize-1 || isLastKey {
			log.Debugf("Inserting %v keys for operator %v into the database", batchSize, operatorName)
			count := i.dbClient.CopyLidoOperatorValidators(operatorName, operator.Index, validatorPubkeys)
			log.Debugf("Inserted %v validators for operator %v. %v remaining", count, operatorName, remainingKeys)
			validatorPubkeys = nil
		} else {
			batchIndex++
		}
	}
	log.Infof("Got %v new keys for operator %v", operator.TotalSigningKeys-operatorValidatorCount, operatorName)

	return nil
}

func getOperatorName(operator lido.NodeOperator) string {
	if operator.Index < uint64(len(definedOperatorsNames)) {
		return definedOperatorsNames[operator.Index]
	}
	return formatOperatorName(operator.Name)
}

func formatOperatorName(name string) string {
	lower := strings.ToLower(name)
	return strings.ReplaceAll(lower, " ", "") + "_lido"
}
