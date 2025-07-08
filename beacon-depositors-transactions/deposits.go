package beacondepositorstransactions

import (
	"encoding/hex"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/migalabs/eth-pokhar/alchemy"
	"github.com/migalabs/eth-pokhar/models"
	"github.com/migalabs/eth-pokhar/utils"
	logger "github.com/sirupsen/logrus"
)

func (b *BeaconDepositorsTransactions) getDepositsCheckpoint() uint64 {
	lastDeposit, err := b.dbClient.ObtainLastDeposit()
	if err != nil {
		logger.Fatalf("Error obtaining last deposit: %s", err.Error())
	}

	return lastDeposit.BlockNum
}

func (b *BeaconDepositorsTransactions) processDepositTransfers(transfers []alchemy.AssetTransfer, numWorkers int) error {
	var wg sync.WaitGroup

	errCh := make(chan error, len(transfers))
	transferCh := make(chan alchemy.AssetTransfer, len(transfers))
	depositsCh := make(chan models.BeaconDeposit, len(transfers))
	// Create worker pool
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			alreadyProcessed := make(map[string]bool)
			for transfer := range transferCh {
				if b.stop {
					return
				}
				// Your existing processing logic here...
				txHash := common.HexToHash(transfer.Hash)
				if _, ok := alreadyProcessed[transfer.Hash]; ok {
					continue
				}
				alreadyProcessed[transfer.Hash] = true
				var attempt int
				for {
					receipt, err := b.ethClient.TransactionReceipt(b.ctx, txHash)

					if err != nil {
						if strings.Contains(err.Error(), utils.ErrorCode429) {
							if attempt > 10 {
								logger.Debugf("Retrying after 429 error (Attempt %d)\n", attempt)
							}
							time.Sleep(250 * time.Millisecond)
							attempt++
							continue
						}
						logger.Errorf("Error getting receipt for tx %s: %s", txHash.String(), err.Error())
						errCh <- err
						return
					}
					if receipt.Status == 1 {
						for _, log := range receipt.Logs {
							event := make(map[string]interface{})
							err := b.contractABI.UnpackIntoMap(event, "DepositEvent", log.Data)
							if err != nil {
								continue
							}
							pubkey := hex.EncodeToString(event["pubkey"].([]byte))
							blockNum, err := strconv.ParseUint(strings.TrimPrefix(transfer.BlockNum, utils.AddressPrefix), 16, 64)
							if err != nil {
								logger.Errorf("Error parsing block number: %s", err.Error())
								errCh <- err
								return
							}
							withdrawalCredentials, ok := event["withdrawal_credentials"].([]byte)
							if !ok {
								logger.Errorf("Error parsing withdrawal credentials for tx %s", txHash.String())
								errCh <- err
								return
							}
							deposit := models.BeaconDeposit{
								BlockNum:          blockNum,
								Depositor:         strings.TrimPrefix(transfer.From, utils.AddressPrefix),
								TxHash:            strings.TrimPrefix(txHash.String(), utils.AddressPrefix),
								ValidatorPubkey:   pubkey,
								WithdrawalAddress: hex.EncodeToString(withdrawalCredentials[utils.WithdrawalPrefixLen:]),
							}
							depositsCh <- deposit
						}
					}
					break
				}
			}
		}()
	}

	// Feed transfers to the worker pool
	go func() {
		for _, transfer := range transfers {
			transferCh <- transfer
		}
		close(transferCh)
	}()

	go func() {
		var deposits []models.BeaconDeposit
		for deposit := range depositsCh {
			deposits = append(deposits, deposit)
		}
		if len(deposits) > 0 && !b.stop {
			b.dbClient.CopyBeaconDeposits(deposits)
		}
	}()

	// Wait for all workers to finish
	wg.Wait()

	// Close error channel after all workers are done
	close(errCh)
	// Collect errors from goroutines
	for err := range errCh {
		if err != nil {
			return err
		}
	}
	close(depositsCh)

	return nil
}
