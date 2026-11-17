package identify

import (
	"math/big"
	"os"
	"reflect"
	"strconv"
	"testing"

	db "github.com/migalabs/eth-pokhar/db"
	"github.com/migalabs/eth-pokhar/lido"
	"github.com/migalabs/eth-pokhar/lido/csm"
)

func TestScannedStakingModules(t *testing.T) {
	want := map[uint64]string{1: "curated", 2: "SDVT", 3: "CSM", 4: "CMv2", 5: "CSM 0x02"}
	if got := scannedStakingModules(); !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestLidoEntityTags(t *testing.T) {
	protocols, tags := lidoEntityTags()
	wantProtocols := []string{db.LidoProtocolCMv2, db.LidoProtocolCSM0x02}
	wantTags := []string{"lido_cmv2", "lido_csm0x02"}
	if !reflect.DeepEqual(protocols, wantProtocols) || !reflect.DeepEqual(tags, wantTags) {
		t.Errorf("got %v -> %v, want %v -> %v", protocols, tags, wantProtocols, wantTags)
	}
}

func TestCSMFamilyProtocolsAreUnique(t *testing.T) {
	seen := map[string]bool{db.LidoProtocolCurated: true, db.LidoProtocolSDVT: true}
	for _, module := range csmFamilyModules {
		if seen[module.protocol] {
			t.Errorf("protocol %q is used twice: t_lido key counts would mix modules", module.protocol)
		}
		seen[module.protocol] = true
	}
}

// TestCSMFamilyOnMainnet reads every CSM-family module and the StakingRouter
// through the real bindings. It needs an execution endpoint, so it only runs
// when POKHAR_TEST_EL_ENDPOINT is set.
func TestCSMFamilyOnMainnet(t *testing.T) {
	endpoint := os.Getenv("POKHAR_TEST_EL_ENDPOINT")
	if endpoint == "" {
		t.Skip("POKHAR_TEST_EL_ENDPOINT not set")
	}

	ids, err := lido.GetStakingModuleIds(endpoint)
	if err != nil {
		t.Fatalf("listing staking modules: %v", err)
	}
	if missing := lido.UnscannedModules(ids, scannedStakingModules()); len(missing) > 0 {
		t.Errorf("StakingRouter modules not scanned by pokhar: %v", missing)
	}

	for _, module := range csmFamilyModules {
		contract, err := csm.NewContractAt(endpoint, module.address)
		if err != nil {
			t.Fatalf("%s: %v", module.name, err)
		}
		count, err := contract.GetNodeOperatorsCount()
		if err != nil || count == 0 {
			t.Fatalf("%s: operators=%d err=%v", module.name, count, err)
		}
		var deposited uint64
		for op := int64(0); op < count; op++ {
			operator, err := contract.GetOperatorData(big.NewInt(op))
			if err != nil {
				t.Fatalf("%s operator %d: %v", module.name, op, err)
			}
			deposited += uint64(operator.Operator.TotalDepositedKeys)
		}
		t.Logf("%s (module %d): operators=%d deposited=%d", module.name, module.id, count, deposited)
	}
}

func TestLidoDimensions(t *testing.T) {
	d := lidoDimensions()
	custody := map[string]string{}
	for k, protocol := range d.CustodyProtocols {
		custody[protocol] = d.CustodyTags[k]
	}
	wantCustody := map[string]string{
		db.LidoProtocolCurated: "lido_cmv1", db.LidoProtocolSDVT: "lido_sdvt", db.LidoProtocolCSM: "lido_csm",
		db.LidoProtocolCMv2: "lido_cmv2", db.LidoProtocolCSM0x02: "lido_csm0x02",
	}
	if !reflect.DeepEqual(custody, wantCustody) {
		t.Errorf("custody: got %v, want %v", custody, wantCustody)
	}

	operators := map[string]string{}
	for k, protocol := range d.OperatorProtocols {
		operators[protocol+"/"+strconv.Itoa(int(d.OperatorIndexes[k]))] = d.OperatorTags[k]
	}
	if got, want := len(operators), len(lido.CuratedCanonicalOperators)+len(lido.CMv2CuratedOperator); got != want {
		t.Errorf("got %d operator entries, want %d", got, want)
	}
	for key, want := range map[string]string{
		"curated/10": "figment", "curated/25": "twinstake",
		"cmv2/44": "figment", "cmv2/8": "twinstake", "cmv2/3": "rockx", "cmv2/6": "consensys",
	} {
		if operators[key] != want {
			t.Errorf("%s: got %q, want %q", key, operators[key], want)
		}
	}
}
