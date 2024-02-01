package beacondepositorstransactions

import (
	"context"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/migalabs/eth-pokhar/alchemy"
	"github.com/migalabs/eth-pokhar/config"
	"github.com/migalabs/eth-pokhar/db"
	"github.com/migalabs/eth-pokhar/utils"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type BeaconDepositorsTransactions struct {
	ctx           context.Context
	cancel        context.CancelFunc
	iConfig       *config.BeaconDepositorsTransactionsConfig
	dbClient      *db.PostgresDBService // client to communicate with psql
	ethClient     *ethclient.Client
	routineClosed chan struct{} // signal that everything was closed succesfully
	alchemyClient *alchemy.AlchemyClient
	stop          bool
	contractABI   abi.ABI
	wgMainRoutine *sync.WaitGroup
	wgDownload    *sync.WaitGroup
	wgUpdateTx    *sync.WaitGroup
}

func NewBeaconDepositorsTransactions(pCtx context.Context, iConfig *config.BeaconDepositorsTransactionsConfig) (*BeaconDepositorsTransactions, error) {

	ctx, cancel := context.WithCancel(pCtx)

	alchemyClient, err := alchemy.NewAlchemyClient(ctx, iConfig.AlchemyURL)
	if err != nil {
		return &BeaconDepositorsTransactions{
			ctx:    ctx,
			cancel: cancel,
		}, errors.Wrap(err, "Error creating alchemy client")
	}

	idbClient, err := db.New(ctx, iConfig.DBUrl, db.WithWorkers(iConfig.DBWorkers))
	if err != nil {
		return &BeaconDepositorsTransactions{
			ctx:    ctx,
			cancel: cancel,
		}, errors.Wrap(err, "unable to generate DB Client.")
	}

	idbClient.Connect()

	elClient, err := ethclient.DialContext(ctx, iConfig.ElEndpoint)
	if err != nil {
		return &BeaconDepositorsTransactions{
			ctx:    ctx,
			cancel: cancel,
		}, errors.Wrap(err, "unable to generate Eth Client.")
	}

	contractABI, err := abi.JSON(strings.NewReader(string(utils.BeaconchainABI)))
	if err != nil {
		log.Fatal(err)
	}

	return &BeaconDepositorsTransactions{
		ctx:           ctx,
		cancel:        cancel,
		iConfig:       iConfig,
		dbClient:      idbClient,
		ethClient:     elClient,
		contractABI:   contractABI,
		alchemyClient: alchemyClient,
		wgMainRoutine: &sync.WaitGroup{},
		wgDownload:    &sync.WaitGroup{},
		wgUpdateTx:    &sync.WaitGroup{},
	}, nil
}

func (b *BeaconDepositorsTransactions) Run() {
	defer b.cancel()
	initTime := time.Now()
	log.Info("Starting BeaconDepositorsTransactions")
	b.wgDownload.Add(1)
	go b.downloadBeaconDeposits()

	b.wgDownload.Wait()
	analysisDuration := time.Since(initTime).Seconds()
	log.Info("BeaconDepositorsTransactions finished in ", analysisDuration)

	initTime = time.Now()
	b.wgUpdateTx.Add(1)
	log.Info("Starting updateDepositorsTransactions")
	go b.updateDepositorsTransactions()
	b.wgUpdateTx.Wait()
	analysisDuration = time.Since(initTime).Seconds()
	log.Info("updateDepositorsTransactions finished in ", analysisDuration)

	b.stop = true
	b.routineClosed <- struct{}{}

}

func (b *BeaconDepositorsTransactions) Close() {
	log.Info("Sudden closed detected, closing BeaconDepositorsTransactions")
	b.stop = true
	<-b.routineClosed // Wait for services to stop before returning
}
