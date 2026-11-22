package lido

import (
	"context"
	"os"
	"regexp"
	"testing"
)

func TestCuratedCanonicalOperators(t *testing.T) {
	// The registry only grows. curated.TestCanonicalOperatorPerDisplayName
	// keeps this list in step with the display names; the run warns when the
	// chain has more operators than either.
	if got := len(CuratedCanonicalOperators); got < 39 {
		t.Fatalf("expected one operator per curated registry index (at least 39), got %d", got)
	}
	urlSafe := regexp.MustCompile(`^[a-z0-9._-]+$`)
	seen := map[string]int{}
	for index, name := range CuratedCanonicalOperators {
		if !urlSafe.MatchString(name) {
			t.Errorf("index %d: %q is not a lowercase URL-safe tag", index, name)
		}
		if prev, dup := seen[name]; dup {
			t.Errorf("indexes %d and %d share %q: one company per registry entry", prev, index, name)
		}
		seen[name] = index
	}
	// Companies that also operate outside Lido must keep that label.
	for index, want := range map[int]string{2: "p2porg", 10: "figment", 13: "blockdaemon", 23: "galaxy_digital", 25: "twinstake"} {
		if CuratedCanonicalOperators[index] != want {
			t.Errorf("index %d: got %q, want %q", index, CuratedCanonicalOperators[index], want)
		}
	}
}

func TestCMv2CuratedOperatorInRange(t *testing.T) {
	for op, curated := range CMv2CuratedOperator {
		if curated >= uint64(len(CuratedCanonicalOperators)) {
			t.Errorf("CMv2 operator %d links to curated %d, out of range", op, curated)
		}
	}
}

func TestDecodeExternalOperator(t *testing.T) {
	module, operator, err := DecodeExternalOperator([]byte{0x00, 0x01, 0, 0, 0, 0, 0, 0, 0, 0x19})
	if err != nil || module != 1 || operator != 25 {
		t.Fatalf("got module=%d operator=%d err=%v, want 1, 25, nil", module, operator, err)
	}
	if _, _, err := DecodeExternalOperator([]byte{0x00, 0x01}); err == nil {
		t.Fatal("expected an error for a short payload")
	}
}

// TestCMv2CuratedOperatorMatchesChain checks the static links against the
// MetaRegistry. It needs an execution endpoint, so it only runs when
// POKHAR_TEST_EL_ENDPOINT is set.
func TestCMv2CuratedOperatorMatchesChain(t *testing.T) {
	endpoint := os.Getenv("POKHAR_TEST_EL_ENDPOINT")
	if endpoint == "" {
		t.Skip("POKHAR_TEST_EL_ENDPOINT not set")
	}
	links, unchecked, err := CMv2CuratedLinksOnChain(context.Background(), endpoint, uint64(len(CMv2CuratedOperator)), 10)
	if err != nil {
		t.Fatal(err)
	}
	for op, err := range unchecked {
		t.Errorf("CMv2 operator %d: group not readable: %v", op, err)
	}
	for op, static := range CMv2CuratedOperator {
		chain, grouped := links[op]
		if !grouped {
			t.Logf("CMv2 operator %d: in no MetaRegistry group, linked by hand to curated %d", op, static)
			continue
		}
		if chain != static {
			t.Errorf("CMv2 operator %d: table says curated %d, MetaRegistry says %d", op, static, chain)
		}
	}
}
