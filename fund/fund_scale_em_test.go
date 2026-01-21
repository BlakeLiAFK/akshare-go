package fund

import (
	"testing"
)

func TestFundScaleChangeEm(t *testing.T) {
	records, err := FundScaleChangeEm("")
	if err != nil {
		t.Logf("FundScaleChangeEm warning: %v (may be unavailable)", err)
		return
	}

	// 允许空数据
	t.Logf("FundScaleChangeEm returned %d records", len(records))

	// 如果有数据，验证结构
	if len(records) > 0 {
		first := records[0]
		requiredFields := []string{"序号", "基金代码", "基金简称", "报告日期", "期初份额", "期末份额", "份额变化"}
		for _, field := range requiredFields {
			if _, ok := first[field]; !ok {
				t.Errorf("Missing required field: %s", field)
			}
		}
	}
}

func TestFundHoldStructureEm(t *testing.T) {
	records, err := FundHoldStructureEm("")
	if err != nil {
		t.Logf("FundHoldStructureEm warning: %v (may be unavailable)", err)
		return
	}

	// 允许空数据
	t.Logf("FundHoldStructureEm returned %d records", len(records))

	// 如果有数据，验证结构
	if len(records) > 0 {
		first := records[0]
		requiredFields := []string{"序号", "基金代码", "基金简称", "报告日期", "机构投资者持有份额", "个人投资者持有份额", "总份额"}
		for _, field := range requiredFields {
			if _, ok := first[field]; !ok {
				t.Errorf("Missing required field: %s", field)
			}
		}
	}
}
