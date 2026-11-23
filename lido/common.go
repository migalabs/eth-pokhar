package lido

import (
	"context"
	"strings"
	"time"

	"github.com/migalabs/eth-pokhar/utils"
	"github.com/pkg/errors"
)

const PublicKeyLength = 48

// CheckPubkeysLength reports whether a getSigningKeys answer holds exactly
// count concatenated public keys. A short answer would otherwise be sliced past
// its end, a panic inside the per-operator goroutines that stops the process.
func CheckPubkeysLength(pubkeys []byte, count uint64) error {
	if want := count * PublicKeyLength; uint64(len(pubkeys)) != want {
		return errors.Errorf("signing keys answer has %d bytes, want %d for %d keys", len(pubkeys), want, count)
	}
	return nil
}

// maxFailedAttempts is how many times a call may fail for a reason other than
// rate limiting before RetryContractCall gives up.
const maxFailedAttempts = 5

// RetryContractCall retries call until it succeeds. Rate limited answers (429)
// are retried without limit, since the key registries must be read in full,
// and any other error up to maxFailedAttempts times.
func RetryContractCall(call func() (interface{}, error)) (interface{}, error) {
	return retryContractCall(context.Background(), -1, call)
}

// RetryContractCallContext is RetryContractCall for best-effort checks that
// must not hold up a run: it also gives up once ctx is done, and after
// maxRateLimited rate limited answers instead of retrying them forever.
func RetryContractCallContext(ctx context.Context, maxRateLimited int, call func() (interface{}, error)) (interface{}, error) {
	return retryContractCall(ctx, maxRateLimited, call)
}

// retryContractCall implements both retry policies. A negative maxRateLimited
// retries rate limited answers without limit.
func retryContractCall(ctx context.Context, maxRateLimited int, call func() (interface{}, error)) (interface{}, error) {
	failed, rateLimited := 0, 0
	for {
		if err := ctx.Err(); err != nil {
			return nil, errors.Wrap(err, "error making contract call")
		}
		result, err := call()
		if err == nil {
			return result, nil
		}
		if strings.Contains(err.Error(), utils.ErrorCode429) {
			rateLimited++
			if maxRateLimited >= 0 && rateLimited > maxRateLimited {
				return nil, errors.Wrapf(err, "error making contract call: still rate limited after %d retries", maxRateLimited)
			}
		} else {
			failed++
			if failed > maxFailedAttempts {
				return nil, errors.Wrap(err, "error making contract call")
			}
		}
		select {
		case <-ctx.Done():
			return nil, errors.Wrap(ctx.Err(), "error making contract call")
		case <-time.After(utils.GetRandomTimeout()):
		}
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
