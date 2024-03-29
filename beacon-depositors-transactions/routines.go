package beacondepositorstransactions

import (
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"

	"sync"

	"time"

	"github.com/migalabs/eth-pokhar/alchemy"
	"github.com/migalabs/eth-pokhar/models"
	"github.com/migalabs/eth-pokhar/utils"
)

func (b *BeaconDepositorsTransactions) updateDepositorsTransactions() {
	log.Info("Getting checkpoints")
	checkpoints, err := b.dbClient.ObtainCheckpointPerDepositor()
	if err != nil {
		log.Fatalf("Error obtaining checkpoints: %s", err.Error())
	}
	log.Info("Got checkpoints")

	log.Info("Fetching new transactions")

	var wg sync.WaitGroup

	checkpointsCh := make(chan models.DepositorCheckpoint)

	for i := 0; i < b.iConfig.Workers; i++ {
		wg.Add(1)
		go b.workerFetchTransactions(&wg, checkpointsCh)
	}
	totalCheckpoints := len(checkpoints)

	checkpointsProcessed := 0

	// Log progress and ETA every 1 minute
	go func() {
		// Start a ticker to log progress every 1 minute
		ticker := time.NewTicker(1 * time.Minute)
		defer ticker.Stop()
		startTime := time.Now()
		for range ticker.C {
			if b.stop {
				return
			}
			elapsedTime := time.Since(startTime)
			processedPercentage := float64(checkpointsProcessed) / float64(totalCheckpoints) * 100
			remainingPercentage := 100 - processedPercentage
			if processedPercentage == 0 {
				continue
			}
			eta := time.Duration(float64(elapsedTime) / processedPercentage * remainingPercentage)
			log.Infof("Progress: %.2f%% | ETA: %s", processedPercentage, eta.Round(time.Minute))
		}
	}()

	for _, checkpoint := range checkpoints {
		if b.stop {
			return
		}
		checkpointsCh <- checkpoint
		checkpointsProcessed++
	}

	close(checkpointsCh)

	wg.Wait()

	log.Info("All transactions fetched")
}

func (b *BeaconDepositorsTransactions) downloadBeaconDeposits() {
	firstCall := true

	lastBlocknumProcessed := b.getDepositsCheckpoint()
	log.Infof("Last block processed: %d", lastBlocknumProcessed)
	startBlock := uint64(0)
	if lastBlocknumProcessed > 1000 {
		startBlock = lastBlocknumProcessed - 1000
	}
	lastBlocknumProcessedHex := utils.AddressPrefix + strconv.FormatUint(startBlock, 16)
	params := alchemy.NewGetAssetTransfersArgs(
		alchemy.SetToAddress(utils.BeaconContractAddress),
		alchemy.SetFromBlock(lastBlocknumProcessedHex),
		alchemy.SetCategory([]string{"external", "internal"}),
	)
	for params.PageKey != "" || firstCall {
		if b.stop {
			return
		}
		newTransfers, newPageKey, err := b.alchemyClient.GetAssetTransfers(b.ctx, params)
		if err != nil {
			log.Fatalf("Error getting asset transfers: %s", err.Error())
		}
		num, err := strconv.ParseUint(strings.TrimPrefix(newTransfers[0].BlockNum, utils.AddressPrefix), 16, 64)
		if err != nil {
			log.Fatalf("Error parsing block number: %s", err.Error())
		}
		log.Infof("Downloaded 1000 more deposits on block %d", num)
		params.PageKey = newPageKey
		firstCall = false
		err = b.processDepositTransfers(newTransfers, b.iConfig.Workers)
		if err != nil {
			log.Fatalf("Error processing deposit transfers: %s", err.Error())
		}

	}
}
