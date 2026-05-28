package lido

import (
	"strings"
	"time"

	"github.com/migalabs/eth-pokhar/utils"
	"github.com/pkg/errors"
)

const PublicKeyLength = 48

func RetryContractCall(call func() (interface{}, error)) (interface{}, error) {
	retry := 0
	for {
		result, err := call()
		if err != nil {
			if !strings.Contains(err.Error(), utils.ErrorCode429) {
				retry++
			}
			if retry > 5 {
				return nil, errors.Wrap(err, "error making contract call")
			}
			waitTime := utils.GetRandomTimeout()
			time.Sleep(waitTime)
			continue
		}
		return result, nil
	}
}
// FormatOperatorName normalizes a raw on-chain operator name into a URL-safe
// pool tag. Spaces are dropped and ':' (used by SDVT cluster names like
// "LidoxSSV: AgileAntelope") is replaced with '_' so frontends don't need to
// URL-decode the tag. The "_lido" suffix is only appended when the name does
// not already self-identify as Lido — SDVT operators start with "lidox..." so
// the suffix would just be redundant. Operators whose normalized name does not
// contain "lido" (e.g. future Curated additions) still get the suffix so
// downstream consumers that filter on "%lido%" keep working.
func FormatOperatorName(name string) string {
	lower := strings.ToLower(name)
	lower = strings.ReplaceAll(lower, " ", "")
	lower = strings.ReplaceAll(lower, ":", "_")
	if strings.Contains(lower, "lido") {
		return lower
	}
	return lower + "_lido"
}
