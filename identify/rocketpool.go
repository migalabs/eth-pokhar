package identify

import (
	"encoding/hex"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/migalabs/eth-pokhar/utils"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"

	"github.com/rocket-pool/rocketpool-go/minipool"
	"github.com/rocket-pool/rocketpool-go/rocketpool"
	"github.com/rocket-pool/rocketpool-go/utils/client"
)

// https://github.com/alrevuelta/eth-metrics/blob/master/pools/rocketpool.go

type RocketpoolMinipool struct {
	Address     []byte
	Pubkey      []byte
	NodeAddress []byte
}

const MaxRetries = 10
const TimeoutSeconds = 1

// Mainnet rocket pool contract
const RocketStorage = "0x1d8f8f00cfa6758d7bE78336684788Fb0ee0Fa46"

func (i *Identify) GetRocketPoolKeys() ([]string, error) {
	var rocketpoolKeys []string

	numValidators, err := i.dbClient.ObtainRocketpoolValidatorCount()
	if err != nil {
		return nil, errors.Wrap(err, "could not obtain rocketpool validator count")
	}

	lastAddressSavedIndex := numValidators - 1

	proxy := client.NewEth1ClientProxy(60*time.Second, i.iConfig.ElEndpoint)

	rp, err := rocketpool.NewRocketPool(proxy, common.HexToAddress(RocketStorage))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("bad contract address: %s", RocketStorage))
	}

	// Get all minipools
	log.Info("Getting minipool addresses")
	minipools, err := getMinipoolAddressesWithErrorHandling(rp, uint64(lastAddressSavedIndex+1), nil)
	if err != nil {
		return nil, errors.Wrap(err, "error getting minipool addresses")
	}
	log.Info("Got minipools addresses")

	log.WithFields(log.Fields{
		"LastAddressSavedIndex": lastAddressSavedIndex,
	}).Debug("Last address saved index")
	// Get the validator pubkey for each minipool
	var wg sync.WaitGroup
	minipoolKeysCh := make(chan string, len(minipools))
	workerSemaphore := make(chan struct{}, i.iConfig.Workers)

	for i := 0; i < len(minipools); i++ {
		wg.Add(1)
		workerSemaphore <- struct{}{}
		go func(minipoolAddress common.Address) {
			defer wg.Done()

			retry := 0
			for {
				pubkey, err := getMiniPoolPubkey(rp, minipoolAddress)
				if err != nil {
					if !strings.Contains(err.Error(), "429") {
						retry++
					}
					if retry > 5 {
						log.WithFields(log.Fields{
							"MinipoolAddress": minipoolAddress.Hex(),
						}).Error("Could not get minipool info")
						return
					}
					waitTime := utils.GetRandomTimeout()
					time.Sleep(waitTime)
					continue
				}
				log.WithFields(log.Fields{
					"Pubkey": hex.EncodeToString(pubkey),
				}).Debug("Got minipool pubkey")

				minipoolKeysCh <- hex.EncodeToString(pubkey)
				break
			}
			<-workerSemaphore
		}(minipools[i])
	}

	go func() {
		wg.Wait()
		close(minipoolKeysCh)
	}()

	for key := range minipoolKeysCh {
		rocketpoolKeys = append(rocketpoolKeys, key)
	}

	return rocketpoolKeys, nil
}

// Get all minipool addresses
func getMinipoolAddressesWithErrorHandling(rp *rocketpool.RocketPool, startFrom uint64, opts *bind.CallOpts) ([]common.Address, error) {

	// Get minipool count
	minipoolCount, err := minipool.GetMinipoolCount(rp, opts)
	if err != nil {
		return []common.Address{}, err
	}
	// Load minipool addresses in batches
	addresses := make([]common.Address, minipoolCount-startFrom)
	for bsi := startFrom; bsi < minipoolCount; bsi += minipool.MinipoolAddressBatchSize {

		// Get batch start & end index
		msi := bsi
		mei := bsi + minipool.MinipoolAddressBatchSize
		if mei > minipoolCount {
			mei = minipoolCount
		}

		// Load addresses
		var wg errgroup.Group
		for mi := msi; mi < mei; mi++ {
			mi := mi
			wg.Go(func() error {
				var err error
				var address common.Address
				for i := 0; i < MaxRetries; i++ {
					address, err = minipool.GetMinipoolAt(rp, mi, opts)
					if err == nil {
						addresses[mi-startFrom] = address
						break
					}

					log.WithFields(log.Fields{
						"Error": err,
					}).Debug("Error getting minipool address. Retrying...")

					time.Sleep(TimeoutSeconds * time.Second)

				}
				return err
			})
		}
		if err := wg.Wait(); err != nil {
			return []common.Address{}, err
		}

	}

	// Return
	return addresses, nil

}

func getMiniPoolPubkey(
	rp *rocketpool.RocketPool,
	address common.Address) ([]byte, error) {

	pubkey, err := minipool.GetMinipoolPubkey(
		rp,
		common.BytesToAddress(address.Bytes()),
		nil)

	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("could not get minipool key: %s", address))
	}

	return pubkey.Bytes(), nil
}
