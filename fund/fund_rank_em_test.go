package fund

import (
	"testing"
)

func TestFundOpenFundRankEm(t *testing.T) {
	// 测试全部基金
	records, err := FundOpenFundRankEm("全部")
	if err != nil {
		t.Fatalf("FundOpenFundRankEm(全部) failed: %v", err)
	}

	// 允许空数据
	t.Logf("FundOpenFundRankEm(全部) returned %d records", len(records))

	// 如果有数据，验证字段
	if len(records) > 0 {
		first := records[0]
		requiredFields := []string{"基金代码", "基金简称", "单位净值", "日增长率"}
		for _, field := range requiredFields {
			if _, ok := first[field]; !ok {
				t.Errorf("Missing required field: %s", field)
			}
		}
	}

	// 测试股票型
	records2, err := FundOpenFundRankEm("股票型")
	if err != nil {
		t.Fatalf("FundOpenFundRankEm(股票型) failed: %v", err)
	}

	t.Logf("FundOpenFundRankEm(股票型) returned %d records", len(records2))
}

func TestFundExchangeRankEm(t *testing.T) {
	records, err := FundExchangeRankEm()
	if err != nil {
		t.Fatalf("FundExchangeRankEm failed: %v", err)
	}

	// 允许空数据
	t.Logf("FundExchangeRankEm returned %d records", len(records))
}

func TestFundMoneyRankEm(t *testing.T) {
	records, err := FundMoneyRankEm()
	if err != nil {
		t.Fatalf("FundMoneyRankEm failed: %v", err)
	}

	// 允许空数据
	t.Logf("FundMoneyRankEm returned %d records", len(records))

	// 如果有数据，验证字段
	if len(records) > 0 {
		first := records[0]
		requiredFields := []string{"基金代码", "基金简称", "万份收益", "7日年化收益率"}
		for _, field := range requiredFields {
			if _, ok := first[field]; !ok {
				t.Errorf("Missing required field: %s", field)
			}
		}
	}
}

func TestFundHkRankEm(t *testing.T) {
	records, err := FundHkRankEm()
	if err != nil {
		t.Fatalf("FundHkRankEm failed: %v", err)
	}

	// 允许空数据
	t.Logf("FundHkRankEm returned %d records", len(records))
}
