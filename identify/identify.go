package identify

import (
	"context"
	"sync"
	"sync/atomic"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/migalabs/eth-pokhar/alchemy"
	"github.com/migalabs/eth-pokhar/config"
	"github.com/migalabs/eth-pokhar/db"
)

type Identify struct {
	ctx                  context.Context
	cancel               context.CancelFunc
	iConfig              *config.IdentifyConfig
	dbClient             *db.PostgresDBService // client to communicate with psql
	ethClient            *ethclient.Client
	routineClosed        chan struct{} // signal that everything was closed succesfully
	alchemyClient        *alchemy.AlchemyClient
	stop                 bool
	contractABI          abi.ABI
	wgMainRoutine        *sync.WaitGroup
	wgDownload           *sync.WaitGroup
	wgUpdateTx           *sync.WaitGroup
	checkpointsProcessed *atomic.Uint64
}

func NewIdentify(pCtx context.Context, iConfig *config.IdentifyConfig) (*Identify, error) {

	return &Identify{}, nil
}

func (i *Identify) Run() {
}

func (i *Identify) Close() {
}
