package beacondepositorstransactions

import (
	"context"
	"strings"
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

	idbClient, err := db.New(ctx, iConfig.DBUrl)
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
		routineClosed: make(chan struct{}),
	}, nil
}

func (b *BeaconDepositorsTransactions) Run() {
	defer b.cancel()

	if !b.stop {
		initTime := time.Now()
		log.Info("Starting beacon_depositors_transactions...")
		b.downloadBeaconDeposits()
		analysisDuration := time.Since(initTime).Seconds()
		log.Info("Finished beacon_depositors_transactions in ", analysisDuration)
	}
	b.CloseConnections()
	log.Debug("Sending signal that beacon_depositors_transactions finished")
	b.routineClosed <- struct{}{}
}

func (b *BeaconDepositorsTransactions) Stop() {
	log.Info("Closing beacon_depositors_transactions...")
	b.stop = true
	// Wait until the routine is closed
	<-b.routineClosed
	log.Info("Closed beacon_depositors_transactions")
}

func (b *BeaconDepositorsTransactions) CloseConnections() {
	log.Debug("Closing Eth connection")
	b.ethClient.Close()
	log.Debug("Eth connection closed")
	log.Debug("Closing Alchemy connection")
	b.alchemyClient.Close()
	log.Debug("Alchemy connection closed")
	log.Debug("Closing DB connection")
	b.dbClient.Finish()
	log.Debug("DB connection closed")
}
