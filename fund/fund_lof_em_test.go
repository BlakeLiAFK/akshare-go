package fund

import (
	"testing"
)

func TestFundLofSpotEm(t *testing.T) {
	records, err := FundLofSpotEm()
	if err != nil {
		t.Fatalf("FundLofSpotEm failed: %v", err)
	}

	// 允许空数据
	t.Logf("FundLofSpotEm returned %d records", len(records))
}

func TestFundLofHistEm(t *testing.T) {
	records, err := FundLofHistEm("161725", "daily", "", "", "")
	if err != nil {
		t.Fatalf("FundLofHistEm failed: %v", err)
	}

	// 允许空数据
	t.Logf("FundLofHistEm returned %d records", len(records))
}

func TestFundLofHistMinEm(t *testing.T) {
	records, err := FundLofHistMinEm("161725", "", "", "1", "")
	if err != nil {
		t.Fatalf("FundLofHistMinEm failed: %v", err)
	}

	// 允许空数据
	t.Logf("FundLofHistMinEm returned %d records", len(records))
}
