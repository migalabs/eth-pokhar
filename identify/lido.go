package identify

import (
	"encoding/hex"
	"strings"

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

func (i *Identify) IdentifyLidoValidators() error {
	log.Debug("Creating a new instance of LidoContract")

	operatorsValidatorCount, err := i.dbClient.ObtainLidoOperatorsValidatorCount()
	if err != nil {
		return err
	}

	lidoContract, err := lido.NewLidoContract(i.iConfig.ElEndpoint)
	// Check if there was an error
	if err != nil {
		return err
	}
	log.Debug("Created a new instance of LidoContract")

	log.Debug("Calling the GetOperatorsIndexes function")
	operatorsIndexes, err := lidoContract.GetOperatorsIndexes()
	if err != nil {
		return err
	}
	log.Debugf("Got indexes: %v", operatorsIndexes)

	log.Debug("Calling the GetOperatorsData function")
	operatorsData, err := lidoContract.GetOperatorsData(operatorsIndexes)
	if err != nil {
		return err
	}
	log.Debug("Got operatorsData")

	log.Debug("Getting keys for each operator")
	for _, operator := range operatorsData {
		operatorValidatorCount := uint64(0)
		if operator.Index < uint64(len(operatorsValidatorCount)) {
			operatorValidatorCount = operatorsValidatorCount[operator.Index]
		}
		var operatorName string
		if operator.Index < uint64(len(definedOperatorsNames)) {
			operatorName = definedOperatorsNames[operator.Index]
		} else {
			operatorName = formatOperatorName(operator.Name)
		}
		log.Debugf("Getting new keys for operator %v", operatorName)
		validatorPubkeys := make([]string, operator.TotalSigningKeys-operatorValidatorCount)
		for keyIndex := operatorValidatorCount; keyIndex < operator.TotalSigningKeys; keyIndex++ {
			key, err := lidoContract.GetOperatorKey(operator, keyIndex)
			if err != nil {
				return err
			}
			validatorPubkeys[keyIndex-operatorValidatorCount] = hex.EncodeToString(key.Key)
		}
		if len(validatorPubkeys) == 0 {
			log.Debugf("No new keys for operator %v", operatorName)
			continue
		}
		log.Debugf("Got %v new keys for operator %v", len(validatorPubkeys), operatorName)
		log.Debugf("Inserting new keys for operator %v into the database", operatorName)
		count := i.dbClient.CopyLidoOperatorValidators(operatorName, operator.Index, validatorPubkeys)
		log.Debugf("Inserted %v validators for operator %v", count, operatorName)
	}
	log.Debug("Finished getting keys for each operator")

	return nil
}

func formatOperatorName(name string) string {
	lower := strings.ToLower(name)
	return strings.ReplaceAll(lower, " ", "") + "_lido"
}
