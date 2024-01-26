package beacondepositorstransactions

import (
	"context"
	"sync"
	"time"

	"github.com/migalabs/eth-pokhar/alchemy"
	"github.com/migalabs/eth-pokhar/config"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type BeaconDepositorsTransactions struct {
	ctx           context.Context
	cancel        context.CancelFunc
	iconfig       *config.BeaconDepositorsTransactionsConfig
	routineClosed chan struct{} // signal that everything was closed succesfully
	alchemyClient *alchemy.AlchemyClient
	stop          bool
	wgMainRoutine *sync.WaitGroup
	wgDownload    *sync.WaitGroup
}

func NewBeaconDepositorsTransactions(pCtx context.Context, iconfig *config.BeaconDepositorsTransactionsConfig) (*BeaconDepositorsTransactions, error) {

	ctx, cancel := context.WithCancel(pCtx)

	alchemyClient, err := alchemy.NewAlchemyClient(ctx, iconfig.AlchemyURL)
	if err != nil {
		return &BeaconDepositorsTransactions{
			ctx:    ctx,
			cancel: cancel,
		}, errors.Wrap(err, "Error creating alchemy client")
	}

	return &BeaconDepositorsTransactions{
		ctx:           ctx,
		cancel:        cancel,
		iconfig:       iconfig,
		alchemyClient: alchemyClient,
		wgMainRoutine: &sync.WaitGroup{},
		wgDownload:    &sync.WaitGroup{},
	}, nil
}

func (b *BeaconDepositorsTransactions) Run() {
	defer b.cancel()
	initTime := time.Now()
	log.Info("Starting BeaconDepositorsTransactions")
	b.wgDownload.Add(1)
	go b.initialDownload()

	b.wgDownload.Wait()
	analysisDuration := time.Since(initTime).Seconds()
	log.Info("BeaconDepositorsTransactions finished in ", analysisDuration)
	b.stop = true
	b.routineClosed <- struct{}{}

}

func (b *BeaconDepositorsTransactions) Close() {
	log.Info("Sudden closed detected, closing BeaconDepositorsTransactions")
	b.stop = true
	<-b.routineClosed // Wait for services to stop before returning
}
