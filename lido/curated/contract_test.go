package curated

import "testing"

func TestGetOperatorName(t *testing.T) {
	tests := []struct {
		name     string
		operator NodeOperator
		want     string
	}{
		{
			name:     "a defined index takes its tag from the list, not from the chain",
			operator: NodeOperator{Index: 0, Name: "Whatever The Chain Says"},
			want:     "stakingfacilities_lido",
		},
		{
			// Galaxy acquired CryptoManufaktur's assets in July 2024; the
			// Lido registry still returns the old name for this index.
			name:     "index 23 reports Galaxy even though the registry says CryptoManufaktur",
			operator: NodeOperator{Index: 23, Name: "CryptoManufaktur"},
			want:     "galaxy_lido",
		},
		{
			name:     "an index past the list falls back to the on-chain name",
			operator: NodeOperator{Index: uint64(len(definedOperatorsNames)), Name: "New Operator"},
			want:     "newoperator_lido",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetOperatorName(tt.operator); got != tt.want {
				t.Errorf("GetOperatorName(%+v) = %q, want %q", tt.operator, got, tt.want)
			}
		})
	}
}

// Two indices sharing a tag would silently merge two operators into one
// entity, which is exactly the kind of mistake a hand-edited list invites.
func TestDefinedOperatorNamesAreUnique(t *testing.T) {
	seen := make(map[string]int, len(definedOperatorsNames))
	for i, name := range definedOperatorsNames {
		if first, dup := seen[name]; dup {
			t.Errorf("tag %q is used by both index %d and index %d", name, first, i)
			continue
		}
		seen[name] = i
	}
}
