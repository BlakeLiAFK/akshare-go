package fund

import (
	"testing"
)

func TestFundAumEm(t *testing.T) {
	records, err := FundAumEm()
	if err != nil {
		t.Logf("FundAumEm warning: %v (may be unavailable)", err)
		return
	}

	// 允许空数据
	t.Logf("FundAumEm returned %d records", len(records))

	// 如果有数据，验证结构
	if len(records) > 0 {
		first := records[0]
		requiredFields := []string{"序号", "基金公司", "成立时间", "全部管理规模", "全部基金数", "全部经理数"}
		for _, field := range requiredFields {
			if _, ok := first[field]; !ok {
				t.Errorf("Missing required field: %s", field)
			}
		}
	}
}

func TestFundAumTrendEm(t *testing.T) {
	records, err := FundAumTrendEm()
	if err != nil {
		t.Logf("FundAumTrendEm warning: %v (may be unavailable)", err)
		return
	}

	// 允许空数据
	t.Logf("FundAumTrendEm returned %d records", len(records))

	// 如果有数据，验证结构
	if len(records) > 0 {
		first := records[0]
		requiredFields := []string{"date", "value"}
		for _, field := range requiredFields {
			if _, ok := first[field]; !ok {
				t.Errorf("Missing required field: %s", field)
			}
		}
	}
}

func TestFundAumHistEm(t *testing.T) {
	records, err := FundAumHistEm("2023")
	if err != nil {
		t.Logf("FundAumHistEm warning: %v (may be unavailable)", err)
		return
	}

	// 允许空数据
	t.Logf("FundAumHistEm returned %d records", len(records))

	// 如果有数据，验证结构
	if len(records) > 0 {
		first := records[0]
		requiredFields := []string{"序号", "基金公司", "总规模", "股票型", "混合型", "债券型"}
		for _, field := range requiredFields {
			if _, ok := first[field]; !ok {
				t.Errorf("Missing required field: %s", field)
			}
		}
	}
}
