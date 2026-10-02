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
			// Galaxy acquired CryptoManufaktur's assets in July 2024 and the
			// registry reports "Galaxy" for this index today. The tag had
			// been left behind because the list is consulted first.
			name:     "index 23 reports Galaxy, agreeing with the registry",
			operator: NodeOperator{Index: 23, Name: "Galaxy"},
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

func TestPinnedTagDrift(t *testing.T) {
	tests := []struct {
		name     string
		operator NodeOperator
		want     string
	}{
		{
			name:     "silent when the registry agrees with the pinned tag",
			operator: NodeOperator{Index: 0, Name: "Staking Facilities"},
			want:     "",
		},
		{
			name:     "silent when the operator is not pinned at all",
			operator: NodeOperator{Index: uint64(len(definedOperatorsNames)), Name: "Anything"},
			want:     "",
		},
		{
			name:     "reports the tag a renamed operator would get",
			operator: NodeOperator{Index: 0, Name: "Renamed Co"},
			want:     "renamedco_lido",
		},
		{
			// The case this exists for: index 23 kept reporting
			// CryptoManufaktur long after the registry said Galaxy.
			name:     "would have caught the CryptoManufaktur drift",
			operator: NodeOperator{Index: 23, Name: "CryptoManufaktur"},
			want:     "cryptomanufaktur_lido",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PinnedTagDrift(tt.operator); got != tt.want {
				t.Errorf("PinnedTagDrift(%+v) = %q, want %q", tt.operator, got, tt.want)
			}
		})
	}
}
