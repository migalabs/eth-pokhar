package lido

import (
	"reflect"
	"testing"
)

func TestUnscannedModules(t *testing.T) {
	scanned := map[uint64]string{1: "curated", 2: "SDVT", 3: "CSM"}
	tests := []struct {
		name    string
		onChain []uint64
		want    []uint64
	}{
		{"all scanned", []uint64{1, 2, 3}, nil},
		{"new modules reported in on-chain order", []uint64{1, 2, 3, 5, 4}, []uint64{5, 4}},
		{"empty router", nil, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := UnscannedModules(tt.onChain, scanned)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
