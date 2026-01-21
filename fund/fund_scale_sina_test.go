package fund

import (
	"testing"
)

func TestFundScaleOpenSina(t *testing.T) {
	symbols := []string{"股票型基金", "混合型基金", "债券型基金", "货币型基金", "QDII基金"}

	for _, symbol := range symbols {
		t.Run(symbol, func(t *testing.T) {
			records, err := FundScaleOpenSina(symbol)
			if err != nil {
				t.Logf("FundScaleOpenSina(%s) warning: %v (may be unavailable)", symbol, err)
				return
			}

			// 允许空数据
			t.Logf("FundScaleOpenSina(%s) returned %d records", symbol, len(records))

			// 如果有数据，验证结构
			if len(records) > 0 {
				first := records[0]
				requiredFields := []string{"序号", "基金代码", "基金简称", "单位净值", "总募集规模", "最近总份额", "成立日期", "基金经理", "更新日期"}
				for _, field := range requiredFields {
					if _, ok := first[field]; !ok {
						t.Errorf("Missing required field: %s", field)
					}
				}
			}
		})
	}
}

func TestFundScaleCloseSina(t *testing.T) {
	records, err := FundScaleCloseSina()
	if err != nil {
		t.Logf("FundScaleCloseSina warning: %v (may be unavailable)", err)
		return
	}

	// 允许空数据
	t.Logf("FundScaleCloseSina returned %d records", len(records))

	// 如果有数据，验证结构
	if len(records) > 0 {
		first := records[0]
		requiredFields := []string{"序号", "基金代码", "基金简称", "单位净值", "总募集规模", "最近总份额", "成立日期", "基金经理", "更新日期"}
		for _, field := range requiredFields {
			if _, ok := first[field]; !ok {
				t.Errorf("Missing required field: %s", field)
			}
		}
	}
}

func TestFundScaleStructuredSina(t *testing.T) {
	records, err := FundScaleStructuredSina()
	if err != nil {
		t.Logf("FundScaleStructuredSina warning: %v (may be unavailable)", err)
		return
	}

	// 允许空数据
	t.Logf("FundScaleStructuredSina returned %d records", len(records))

	// 如果有数据，验证结构
	if len(records) > 0 {
		first := records[0]
		requiredFields := []string{"序号", "基金代码", "基金简称", "单位净值", "总募集规模", "最近总份额", "成立日期", "基金经理", "更新日期"}
		for _, field := range requiredFields {
			if _, ok := first[field]; !ok {
				t.Errorf("Missing required field: %s", field)
			}
		}
	}
}
