package beacondepositorstransactions

import (
	"encoding/hex"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"

	"github.com/migalabs/eth-pokhar/alchemy"
	"github.com/migalabs/eth-pokhar/models"
	"github.com/migalabs/eth-pokhar/utils"
)

// func (b *BeaconDepositorsTransactions) initialDownload() {

// }

func (b *BeaconDepositorsTransactions) downloadBeaconDeposits() {
	defer b.wgDownload.Done()
	firstCall := true

	lastBlocknumProcessed := b.getDepositsCheckpoint()
	log.Infof("Last block processed: %d", lastBlocknumProcessed)
	lastBlocknumProcessedHex := "0x" + strconv.FormatUint(lastBlocknumProcessed, 16)
	params := alchemy.NewGetAssetTransfersArgs(
		alchemy.SetToAddress(utils.BeaconContractAddress),
		alchemy.SetFromBlock(lastBlocknumProcessedHex),
	)
	for params.PageKey != "" || firstCall {
		newTransfers, newPageKey, err := b.alchemyClient.GetAssetTransfers(b.ctx, params)
		if err != nil {
			log.Fatal(err)
		}
		params.PageKey = newPageKey
		firstCall = false
		err = b.processTransfers(newTransfers)
		if err != nil {
			log.Fatal(err)
		}

	}
}

func (b *BeaconDepositorsTransactions) processTransfers(transfers []alchemy.AssetTransfer) error {
	for _, transfer := range transfers {
		txHash := common.HexToHash(transfer.Hash)
		receipt, err := b.ethClient.TransactionReceipt(b.ctx, txHash)
		if err != nil {
			return err
		}
		if receipt.Status == 1 {
			for _, log := range receipt.Logs {

				event := make(map[string]interface{})
				err := b.contractABI.UnpackIntoMap(event, "DepositEvent", log.Data)
				if err != nil {
					return err
				}
				pubkey := hex.EncodeToString(event["pubkey"].([]byte))
				blockNum, err := strconv.ParseUint(strings.TrimPrefix(transfer.BlockNum, "0x"), 16, 64)
				if err != nil {
					return err
				}
				deposit := models.BeaconDeposit{
					BlockNum:        blockNum,
					Depositor:       strings.TrimPrefix(transfer.From, "0x"),
					TxHash:          strings.TrimPrefix(txHash.String(), "0x"),
					ValidatorPubkey: pubkey,
				}
				b.dbClient.Persist(deposit)
			}
		}
	}
	return nil
}
