package curated

import (
	"testing"

	"github.com/migalabs/eth-pokhar/lido"
)

// A curated operator added to the display names without its operator label
// (or the other way round) would label its validators inconsistently.
func TestCanonicalOperatorPerDisplayName(t *testing.T) {
	if got, want := len(lido.CuratedCanonicalOperators), len(definedOperatorsNames); got != want {
		t.Fatalf("lido.CuratedCanonicalOperators has %d operators and definedOperatorsNames %d: extend both together", got, want)
	}
}
