package identify

import (
	"context"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/migalabs/eth-pokhar/alchemy"
	"github.com/migalabs/eth-pokhar/config"
	"github.com/migalabs/eth-pokhar/db"
	log "github.com/sirupsen/logrus"
)

type Identify struct {
	ctx           context.Context
	cancel        context.CancelFunc
	iConfig       *config.IdentifyConfig
	dbClient      *db.PostgresDBService // client to communicate with psql
	ethClient     *ethclient.Client
	routineClosed chan struct{} // signal that everything was closed succesfully
	alchemyClient *alchemy.AlchemyClient
	stop          bool
	wgMainRoutine *sync.WaitGroup
}

func NewIdentify(pCtx context.Context, iConfig *config.IdentifyConfig) (*Identify, error) {
	ctx, cancel := context.WithCancel(pCtx)

	alchemyClient, err := alchemy.NewAlchemyClient(ctx, iConfig.AlchemyURL)
	if err != nil {
		return &Identify{
			ctx:    ctx,
			cancel: cancel,
		}, err
	}

	idbClient, err := db.New(ctx, iConfig.DBUrl, db.WithWorkers(iConfig.DBWorkers))
	if err != nil {
		return &Identify{
			ctx:    ctx,
			cancel: cancel,
		}, err
	}

	idbClient.Connect()

	elClient, err := ethclient.DialContext(ctx, iConfig.ElEndpoint)
	if err != nil {
		return &Identify{
			ctx:    ctx,
			cancel: cancel,
		}, err
	}

	return &Identify{
		ctx:           ctx,
		cancel:        cancel,
		iConfig:       iConfig,
		dbClient:      idbClient,
		ethClient:     elClient,
		routineClosed: make(chan struct{}),
		alchemyClient: alchemyClient,
		wgMainRoutine: &sync.WaitGroup{},
		stop:          false,
	}, nil
}

func (i *Identify) Run() {
	defer i.cancel()
	initTime := time.Now()

	log.Info("Starting Identify routine")

	if i.iConfig.RecreateTable {
		log.Info("Truncating identified validators table")
		err := i.dbClient.TruncateIdentifiedValidators()
		if err != nil {
			log.Fatalf("Error truncating identified validators table: %v", err)
		}
		log.Info("Truncated identified validators")
	}

	log.Info("Adding new validators to database")
	err := i.dbClient.AddNewValidators()
	if err != nil {
		log.Fatalf("Error adding new validators to database: %v", err)
	}
	log.Info("Added new validators to database")

	log.Info("Identifying coinbase validators")
	err = i.dbClient.IdentifyCoinbaseValidators()
	if err != nil {
		log.Fatalf("Error identifying coinbase validators: %v", err)
	}
	log.Info("Identified coinbase validators")

	endTime := time.Now()
	log.Infof("Identify routine finished in %v", endTime.Sub(initTime))

	i.stop = true
	i.routineClosed <- struct{}{}
}

func (i *Identify) Close() {
	i.cancel()
	i.ethClient.Close()
	i.stop = true
	log.Info("Identify closed")
	<-i.routineClosed
}
