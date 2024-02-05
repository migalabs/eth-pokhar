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

	// Create a wait group to wait for all workers to finish
	var wg sync.WaitGroup

	// Create a channel to communicate with workers
	checkpointsCh := make(chan models.DepositorCheckpoint)

	// Define the number of workers in the thread pool
	numWorkers := 15

	// Start the workers
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go b.workerFetchTransactions(&wg, checkpointsCh)
	}

	// Calculate total number of checkpoints
	totalCheckpoints := len(checkpoints)

	// Create a counter to keep track of processed checkpoints
	checkpointsProcessed := 0

	// Log progress and ETA every 1 minute
	go func() {
		// Start a ticker to log progress every 1 minute
		ticker := time.NewTicker(1 * time.Minute)
		defer ticker.Stop()
		startTime := time.Now()
		for range ticker.C {
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

	// Send checkpoints to workers for processing
	for _, checkpoint := range checkpoints {
		checkpointsCh <- checkpoint
		checkpointsProcessed++
	}

	// Close the channel to signal that no more transactions will be sent
	close(checkpointsCh)

	// Wait for all workers to finish
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
	lastBlocknumProcessedHex := "0x" + strconv.FormatUint(startBlock, 16)
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
		num, err := strconv.ParseUint(strings.TrimPrefix(newTransfers[0].BlockNum, "0x"), 16, 64)
		if err != nil {
			log.Fatalf("Error parsing block number: %s", err.Error())
		}
		log.Debugf("Downloaded 1000 more deposits on block %d", num)
		params.PageKey = newPageKey
		firstCall = false
		err = b.processDepositTransfers(newTransfers, 15)
		if err != nil {
			log.Fatalf("Error processing deposit transfers: %s", err.Error())
		}

	}
}
