package beacondepositorstransactions

import (
	log "github.com/sirupsen/logrus"
)

func (b *BeaconDepositorsTransactions) getDepositsCheckpoint() uint64 {
	lastDeposit, err := b.dbClient.ObtainLastDeposit()
	if err != nil {
		log.Fatal(err)
	}

	return lastDeposit.BlockNum
}
