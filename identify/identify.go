package identify

import (
	"context"
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

	idbClient, err := db.New(ctx, iConfig.DBUrl)
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
		routineClosed: make(chan struct{}, 1),
		alchemyClient: alchemyClient,
		stop:          false,
	}, nil
}

func (i *Identify) Run() {
	defer i.cancel()
	initTime := time.Now()

	log.Info("Starting Identify routine")

	if i.iConfig.RecreateTable && !i.stop {
		log.Info("Truncating identified validators table")
		err := i.dbClient.TruncateIdentifiedValidators()
		if err != nil {
			log.Fatalf("Error truncating identified validators table: %v", err)
		}
		log.Info("Truncated identified validators")
	}

	if !i.stop {
		log.Info("Adding new validators to database")
		err := i.dbClient.AddNewValidators()
		if err != nil {
			log.Fatalf("Error adding new validators to database: %v", err)
		}
		log.Info("Added new validators to database")
	}

	if !i.stop {
		log.Info("Applying validators insert")
		err := i.dbClient.ApplyValidatorsInsert()
		if err != nil {
			log.Fatalf("Error applying validators insert: %v", err)
		}
		log.Info("Applied validators insert")
	}
	if !i.stop {
		log.Info("Applying depositors insert")
		err := i.dbClient.ApplyDepositorsInsert()
		if err != nil {
			log.Fatalf("Error applying depositors insert: %v", err)
		}
		log.Info("Applied depositors insert")
	}

	if !i.stop {
		log.Info("Identifying coinbase validators")
		err := i.dbClient.IdentifyCoinbaseValidators()
		if err != nil {
			log.Fatalf("Error identifying coinbase validators: %v", err)
		}
		log.Info("Identified coinbase validators")
	}

	if !i.stop {
		log.Info("Identifying rocketpool validators")
		startTime := time.Now()

		newRocketpoolKeys, err := i.GetRocketPoolKeys()
		if err != nil {
			log.Fatalf("Error identifying rocketpool validators: %v", err)
		}
		log.WithFields(log.Fields{
			"NewDetectedKeys": len(newRocketpoolKeys),
			"Duration":        time.Since(startTime),
		}).Info("RocketPool Keys:")
		i.dbClient.CopyRocketpoolValidators(newRocketpoolKeys)
		err = i.dbClient.IdentifyRocketpoolValidators()
		if err != nil {
			log.Fatalf("Error identifying rocketpool validators: %v", err)
		}
		log.Info("Identified rocketpool validators")
	}

	endTime := time.Now()
	log.Infof("Identify routine finished in %v", endTime.Sub(initTime))
	i.CloseConnections()
	log.Debug("Sending signal that routine is closed")
	i.routineClosed <- struct{}{}
}

func (i *Identify) Stop() {
	log.Info("Closing identify...")
	i.stop = true
	// Wait until the routine is closed
	<-i.routineClosed
	log.Info("Closed identify")
}

func (i *Identify) CloseConnections() {
	log.Debug("Closing Eth connection")
	i.ethClient.Close()
	log.Debug("Eth connection closed")
	log.Debug("Closing Alchemy connection")
	i.alchemyClient.Close()
	log.Debug("Alchemy connection closed")
	log.Debug("Closing DB connection")
	i.dbClient.Finish()
	log.Debug("DB connection closed")
}
