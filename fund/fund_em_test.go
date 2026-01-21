package fund

import (
	"strings"
	"testing"
)

func TestFundPurchaseEm(t *testing.T) {
	records, err := FundPurchaseEm()
	if err != nil {
		t.Fatalf("FundPurchaseEm failed: %v", err)
	}

	// 允许空数据
	t.Logf("FundPurchaseEm returned %d records", len(records))

	// 如果有数据，验证字段
	if len(records) > 0 {
		first := records[0]
		requiredFields := []string{"基金代码", "基金简称", "申购状态", "赎回状态"}
		for _, field := range requiredFields {
			if _, ok := first[field]; !ok {
				t.Errorf("Missing required field: %s", field)
			}
		}
	}
}

func TestFundNameEm(t *testing.T) {
	records, err := FundNameEm()
	if err != nil {
		t.Fatalf("FundNameEm failed: %v", err)
	}

	// 允许空数据
	t.Logf("FundNameEm returned %d records", len(records))

	// 如果有数据，验证字段
	if len(records) > 0 {
		first := records[0]
		requiredFields := []string{"基金代码", "基金简称", "基金类型"}
		for _, field := range requiredFields {
			if _, ok := first[field]; !ok {
				t.Errorf("Missing required field: %s", field)
			}
		}
	}
}

func TestFundInfoIndexEm(t *testing.T) {
	// 测试沪深指数-被动指数型
	records, err := FundInfoIndexEm("沪深指数", "被动指数型")
	if err != nil {
		t.Fatalf("FundInfoIndexEm failed: %v", err)
	}

	// 允许空数据
	t.Logf("FundInfoIndexEm returned %d records", len(records))

	// 如果有数据，验证字段
	if len(records) > 0 {
		first := records[0]
		requiredFields := []string{"基金代码", "基金名称", "单位净值", "日增长率"}
		for _, field := range requiredFields {
			if _, ok := first[field]; !ok {
				t.Errorf("Missing required field: %s", field)
			}
		}
	}
}

func TestFundOpenFundDailyEm(t *testing.T) {
	records, err := FundOpenFundDailyEm()
	if err != nil {
		t.Fatalf("FundOpenFundDailyEm failed: %v", err)
	}

	// 允许空数据
	t.Logf("FundOpenFundDailyEm returned %d records", len(records))

	// 如果有数据，验证字段
	if len(records) > 0 {
		first := records[0]
		requiredFields := []string{"基金代码", "基金简称"}
		for _, field := range requiredFields {
			if _, ok := first[field]; !ok {
				t.Errorf("Missing required field: %s", field)
			}
		}
	}
}

func TestFundFinancialFundDailyEm(t *testing.T) {
	records, err := FundFinancialFundDailyEm()
	if err != nil {
		t.Fatalf("FundFinancialFundDailyEm failed: %v", err)
	}

	// 允许空数据
	t.Logf("FundFinancialFundDailyEm returned %d records", len(records))

	// 如果有数据，验证字段
	if len(records) > 0 {
		first := records[0]
		requiredFields := []string{"基金代码", "基金简称", "万份收益", "7日年化"}
		for _, field := range requiredFields {
			if _, ok := first[field]; !ok {
				t.Errorf("Missing required field: %s", field)
			}
		}
	}
}

func TestFundMoneyFundInfoEm(t *testing.T) {
	// 测试默认基金代码
	records, err := FundMoneyFundInfoEm("")
	if err != nil {
		t.Fatalf("FundMoneyFundInfoEm failed: %v", err)
	}

	// 允许空数据
	t.Logf("FundMoneyFundInfoEm returned %d records", len(records))

	// 如果有数据，验证字段
	if len(records) > 0 {
		first := records[0]
		requiredFields := []string{"净值日期", "每万份收益", "7日年化收益率"}
		for _, field := range requiredFields {
			if _, ok := first[field]; !ok {
				t.Errorf("Missing required field: %s", field)
			}
		}
	}
}

func TestFundValueEstimationEm(t *testing.T) {
	// 测试全部基金
	records, err := FundValueEstimationEm("")
	if err != nil {
		t.Fatalf("FundValueEstimationEm failed: %v", err)
	}

	// 允许空数据
	t.Logf("FundValueEstimationEm returned %d records", len(records))

	// 如果有数据，验证字段
	if len(records) > 0 {
		first := records[0]
		requiredFields := []string{"基金代码", "基金简称", "单位净值"}
		for _, field := range requiredFields {
			if _, ok := first[field]; !ok {
				t.Errorf("Missing required field: %s", field)
			}
		}
	}
}

func TestFundMoneyFundDailyEm(t *testing.T) {
	records, err := FundMoneyFundDailyEm()
	if err != nil {
		t.Logf("FundMoneyFundDailyEm warning: %v (may be unavailable)", err)
		return
	}

	// 允许空数据
	t.Logf("FundMoneyFundDailyEm returned %d records", len(records))

	// 如果有数据，验证字段
	if len(records) > 0 {
		first := records[0]
		requiredFields := []string{"基金代码", "基金简称", "日涨幅", "成立日期", "基金经理"}
		for _, field := range requiredFields {
			if _, ok := first[field]; !ok {
				t.Errorf("Missing required field: %s", field)
			}
		}

		// 验证日期字段存在（动态生成的字段名）
		hasDateField := false
		for key := range first {
			if strings.Contains(key, "万份收益") || strings.Contains(key, "7日年化") || strings.Contains(key, "单位净值") {
				hasDateField = true
				break
			}
		}
		if !hasDateField {
			t.Errorf("Missing date-related fields (万份收益/7日年化/单位净值)")
		}
	}
}
