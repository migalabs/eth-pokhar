package lido

import (
	"context"
	"errors"
	"testing"
)

func TestCheckPubkeysLength(t *testing.T) {
	if err := CheckPubkeysLength(make([]byte, 3*PublicKeyLength), 3); err != nil {
		t.Errorf("full answer rejected: %v", err)
	}
	for _, size := range []int{0, 3*PublicKeyLength - 1, 3*PublicKeyLength + 1} {
		if err := CheckPubkeysLength(make([]byte, size), 3); err == nil {
			t.Errorf("answer of %d bytes for 3 keys accepted", size)
		}
	}
}

func TestRetryContractCallContextCapsRateLimiting(t *testing.T) {
	calls := 0
	_, err := RetryContractCallContext(context.Background(), 1, func() (interface{}, error) {
		calls++
		return nil, errors.New("429 Too Many Requests")
	})
	if err == nil || calls != 2 {
		t.Fatalf("calls=%d err=%v, want 2 calls and an error", calls, err)
	}
}

func TestRetryContractCallContextStopsWithContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	calls := 0
	if _, err := RetryContractCallContext(ctx, -1, func() (interface{}, error) {
		calls++
		return nil, errors.New("429 Too Many Requests")
	}); err == nil || calls != 0 {
		t.Fatalf("calls=%d err=%v, want no call and an error", calls, err)
	}
}
