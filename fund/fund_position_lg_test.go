package fund

import (
	"testing"
)

func TestFundStockPositionLg(t *testing.T) {
	records, err := FundStockPositionLg("")
	if err != nil {
		t.Logf("FundStockPositionLg warning: %v (may be unavailable)", err)
		return
	}

	// 允许空数据
	t.Logf("FundStockPositionLg returned %d records", len(records))

	// 如果有数据，验证结构
	if len(records) > 0 {
		first := records[0]
		requiredFields := []string{"date", "close", "position"}
		for _, field := range requiredFields {
			if _, ok := first[field]; !ok {
				t.Errorf("Missing required field: %s", field)
			}
		}
	}
}

func TestFundBalancePositionLg(t *testing.T) {
	records, err := FundBalancePositionLg()
	if err != nil {
		t.Logf("FundBalancePositionLg warning: %v (may be unavailable)", err)
		return
	}

	// 允许空数据
	t.Logf("FundBalancePositionLg returned %d records", len(records))

	// 如果有数据，验证结构
	if len(records) > 0 {
		first := records[0]
		requiredFields := []string{"date", "close", "position"}
		for _, field := range requiredFields {
			if _, ok := first[field]; !ok {
				t.Errorf("Missing required field: %s", field)
			}
		}
	}
}

func TestFundLinghuoPositionLg(t *testing.T) {
	records, err := FundLinghuoPositionLg()
	if err != nil {
		t.Logf("FundLinghuoPositionLg warning: %v (may be unavailable)", err)
		return
	}

	// 允许空数据
	t.Logf("FundLinghuoPositionLg returned %d records", len(records))

	// 如果有数据，验证结构
	if len(records) > 0 {
		first := records[0]
		requiredFields := []string{"date", "close", "position"}
		for _, field := range requiredFields {
			if _, ok := first[field]; !ok {
				t.Errorf("Missing required field: %s", field)
			}
		}
	}
}
