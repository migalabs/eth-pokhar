package beacondepositorstransactions

import (
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/migalabs/eth-pokhar/alchemy"
	"github.com/migalabs/eth-pokhar/utils"
)

func (b *BeaconDepositorsTransactions) updateDepositorsTransactions() {
	defer b.wgUpdateTx.Done()
	log.Info("Getting checkpoints")
	checkpoints, err := b.dbClient.ObtainCheckpointPerDepositor()
	if err != nil {
		log.Fatal(err)
	}
	log.Info("Got checkpoints")

	log.Info("Fetching new transactions")
	for _, checkpoint := range checkpoints {
		newTransactions, err := b.fetchNewTransactions(checkpoint)
		if err != nil {
			log.Fatal(err)
		}
		b.dbClient.CopyTransactions(newTransactions)
	}

}

func (b *BeaconDepositorsTransactions) downloadBeaconDeposits() {
	defer b.wgDownload.Done()
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
		newTransfers, newPageKey, err := b.alchemyClient.GetAssetTransfers(b.ctx, params)
		if err != nil {
			log.Fatal(err)
		}
		num, err := strconv.ParseUint(strings.TrimPrefix(newTransfers[0].BlockNum, "0x"), 16, 64)
		if err != nil {
			log.Fatal(err)
		}
		log.Debugf("Downloaded 1000 more deposits on block %d", num)
		params.PageKey = newPageKey
		firstCall = false
		err = b.processDepositTransfers(newTransfers, 15)
		if err != nil {
			log.Fatal(err)
		}

	}
}
