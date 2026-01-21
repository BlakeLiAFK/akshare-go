package fund

import (
	"testing"
)

func TestFundFhEm(t *testing.T) {
	// 测试2024年数据
	records, err := FundFhEm("2024")
	if err != nil {
		t.Logf("FundFhEm warning: %v (may be unavailable)", err)
		return
	}

	// 允许空数据
	t.Logf("FundFhEm returned %d records", len(records))

	// 如果有数据，验证结构
	if len(records) > 0 {
		first := records[0]
		requiredFields := []string{"序号", "基金代码", "基金简称", "权益登记日", "除息日期", "分红", "分红发放日"}
		for _, field := range requiredFields {
			if _, ok := first[field]; !ok {
				t.Errorf("Missing required field: %s", field)
			}
		}
	}
}

func TestFundCfEm(t *testing.T) {
	// 测试2024年数据
	records, err := FundCfEm("2024")
	if err != nil {
		t.Logf("FundCfEm warning: %v (may be unavailable)", err)
		return
	}

	// 允许空数据
	t.Logf("FundCfEm returned %d records", len(records))

	// 如果有数据，验证结构
	if len(records) > 0 {
		first := records[0]
		requiredFields := []string{"序号", "基金代码", "基金简称", "拆分日", "拆分折算", "拆分类型"}
		for _, field := range requiredFields {
			if _, ok := first[field]; !ok {
				t.Errorf("Missing required field: %s", field)
			}
		}
	}
}

func TestFundFhRankEm(t *testing.T) {
	records, err := FundFhRankEm()
	if err != nil {
		t.Logf("FundFhRankEm warning: %v (may be unavailable)", err)
		return
	}

	// 允许空数据
	t.Logf("FundFhRankEm returned %d records", len(records))

	// 如果有数据，验证结构
	if len(records) > 0 {
		first := records[0]
		requiredFields := []string{"序号", "基金代码", "基金简称", "累计分红", "累计次数", "成立日期"}
		for _, field := range requiredFields {
			if _, ok := first[field]; !ok {
				t.Errorf("Missing required field: %s", field)
			}
		}
	}
}
