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
func FormatOperatorName(name string) string {
	lower := strings.ToLower(name)
	return strings.ReplaceAll(lower, " ", "") + "_lido"
}
