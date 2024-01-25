package beacondepositorstransactions

import (
	"context"
	"time"

	"github.com/migalabs/eth-pokhar/config"
	log "github.com/sirupsen/logrus"
)

type BeaconDepositorsTransactions struct {
	ctx           context.Context
	cancel        context.CancelFunc
	config        *config.BeaconDepositorsTransactionsConfig
	routineClosed chan struct{} // signal that everything was closed succesfully

	stop bool
}

func NewBeaconDepositorsTransactions(pCtx context.Context, config *config.BeaconDepositorsTransactionsConfig) (*BeaconDepositorsTransactions, error) {

	ctx, cancel := context.WithCancel(pCtx)

	return &BeaconDepositorsTransactions{
		ctx:    ctx,
		cancel: cancel,
		config: config,
	}, nil
}

func (b *BeaconDepositorsTransactions) Run() {
	defer b.cancel()
	initTime := time.Now()
	log.Info("Starting BeaconDepositorsTransactions")
	analysisDuration := time.Since(initTime).Seconds()
	log.Info("BeaconDepositorsTransactions finished in ", analysisDuration)
	b.routineClosed <- struct{}{}

}

func (b *BeaconDepositorsTransactions) Close() {
	log.Info("Sudden closed detected, closing BeaconDepositorsTransactions")
	b.stop = true
	<-b.routineClosed // Wait for services to stop before returning
}
