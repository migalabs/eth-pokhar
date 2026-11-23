package lido

import (
	"context"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
)

// StakingRouterAddress is Lido's StakingRouter on mainnet. Every staking module
// is registered here and every Lido deposit is sent from it, which is why a
// module pokhar does not scan surfaces as the generic depositor label "lido".
const StakingRouterAddress = "0xFdDf38947aFB03C621C71b06C9C70bce73f12999"

// Only the one view the module check needs; the full StakingRouter ABI is not
// worth a generated binding.
const stakingRouterABI = `[{"inputs":[],"name":"getStakingModuleIds","outputs":[{"internalType":"uint256[]","name":"stakingModuleIds","type":"uint256[]"}],"stateMutability":"view","type":"function"}]`

// GetStakingModuleIds returns the ids of every module registered in the
// StakingRouter. It serves a best-effort check, so it stops once ctx is done
// and after maxRateLimited rate limited answers.
func GetStakingModuleIds(ctx context.Context, nodeEndpoint string, maxRateLimited int) ([]uint64, error) {
	ethClient, err := ethclient.DialContext(ctx, nodeEndpoint)
	if err != nil {
		return nil, err
	}
	defer ethClient.Close()

	parsed, err := abi.JSON(strings.NewReader(stakingRouterABI))
	if err != nil {
		return nil, errors.Wrap(err, "error parsing StakingRouter ABI")
	}
	contract := bind.NewBoundContract(common.HexToAddress(StakingRouterAddress), parsed, ethClient, nil, nil)

	result, err := RetryContractCallContext(ctx, maxRateLimited, func() (interface{}, error) {
		var out []interface{}
		if err := contract.Call(&bind.CallOpts{Context: ctx}, &out, "getStakingModuleIds"); err != nil {
			return nil, err
		}
		if len(out) != 1 {
			return nil, errors.Errorf("unexpected getStakingModuleIds output length %d", len(out))
		}
		return *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int), nil
	})
	if err != nil {
		return nil, err
	}

	raw := result.([]*big.Int)
	ids := make([]uint64, len(raw))
	for i, id := range raw {
		ids[i] = id.Uint64()
	}
	return ids, nil
}

// UnscannedModules returns the ids in onChain that are not in scanned,
// preserving the on-chain order.
func UnscannedModules(onChain []uint64, scanned map[uint64]string) []uint64 {
	var missing []uint64
	for _, id := range onChain {
		if _, ok := scanned[id]; !ok {
			missing = append(missing, id)
		}
	}
	return missing
}
