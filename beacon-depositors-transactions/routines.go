package beacondepositorstransactions

import (
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/migalabs/eth-pokhar/alchemy"
	"github.com/migalabs/eth-pokhar/utils"
)

// func (b *BeaconDepositorsTransactions) initialDownload() {

// }

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
			log.Fatal(err)
		}
		num, err := strconv.ParseUint(strings.TrimPrefix(newTransfers[0].BlockNum, "0x"), 16, 64)
		if err != nil {
			log.Fatal(err)
		}
		log.Debugf("Downloaded 1000 more deposits on block %d", num)
		params.PageKey = newPageKey
		firstCall = false
		err = b.processDepositTransfers(newTransfers, 10)
		if err != nil {
			log.Fatal(err)
		}

	}
}
