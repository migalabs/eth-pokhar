package beacondepositorstransactions

import (
	"strconv"
	"sync"

	"github.com/migalabs/eth-pokhar/alchemy"
	"github.com/migalabs/eth-pokhar/models"
	log "github.com/sirupsen/logrus"
)

func (b *BeaconDepositorsTransactions) workerFetchTransactions(wg *sync.WaitGroup, checkpointsCh <-chan models.DepositorCheckpoint) {
	defer wg.Done()

	for checkpoint := range checkpointsCh {
		if b.stop {
			return
		}
		newTransactions, err := b.fetchNewTransactions(checkpoint)
		if err != nil {
			log.Fatalf("Error fetching new transactions: %s", err.Error())
		}
		go func() {
			b.dbClient.CopyTransactions(newTransactions)
			b.checkpointsProcessed.Add(1)
		}()
	}
}

func (b *BeaconDepositorsTransactions) fetchNewTransactions(depositorCheckpoint models.DepositorCheckpoint) ([]models.Transaction, error) {
	fromBlock := "0x" + strconv.FormatUint(depositorCheckpoint.Checkpoint+1, 16)
	newTransfers := make([]alchemy.AssetTransfer, 0)
	params := alchemy.NewGetAssetTransfersArgs(
		alchemy.SetToAddress("0x"+depositorCheckpoint.Depositor),
		alchemy.SetFromBlock(fromBlock),
		alchemy.SetCategory([]string{"external"}),
		alchemy.SetExcludeZeroValue(false),
	)

	toTransfers, err := b.alchemyClient.GetAllAssetTransfers(b.ctx, params)
	if err != nil {
		return nil, err
	}
	newTransfers = append(newTransfers, toTransfers...)

	params = alchemy.NewGetAssetTransfersArgs(
		alchemy.SetFromAddress("0x"+depositorCheckpoint.Depositor),
		alchemy.SetFromBlock(fromBlock),
		alchemy.SetCategory([]string{"external"}),
		alchemy.SetExcludeZeroValue(false),
	)
	fromTransfers, err := b.alchemyClient.GetAllAssetTransfers(b.ctx, params)
	if err != nil {
		return nil, err
	}
	newTransfers = append(newTransfers, fromTransfers...)

	return transfersToTransactions(newTransfers, depositorCheckpoint.Depositor)
}

func transfersToTransactions(transfers []alchemy.AssetTransfer, depositor string) ([]models.Transaction, error) {
	var transactions []models.Transaction
	var transactionsMap = make(map[string]bool)
	for _, transfer := range transfers {
		//Ignore duplicates (edgy case when a transfer is both from and to the depositor, it will be duplicated in the results)
		if _, ok := transactionsMap[transfer.Hash]; ok {
			continue
		}
		transactionsMap[transfer.Hash] = true
		from := transfer.From
		if len(from) > 2 {
			from = transfer.From[2:]
		}
		to := transfer.To
		if len(to) > 2 {
			to = transfer.To[2:]
		}
		txHash := transfer.Hash
		if len(txHash) > 2 {
			txHash = transfer.Hash[2:]
		}
		blockNum, err := strconv.ParseUint(transfer.BlockNum, 0, 64)
		if err != nil {
			return nil, err
		}

		transactions = append(transactions, models.Transaction{
			BlockNum:  blockNum,
			Value:     transfer.Value,
			From:      from,
			To:        to,
			TxHash:    txHash,
			Depositor: depositor,
		})
	}
	return transactions, nil
}
