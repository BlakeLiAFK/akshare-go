package fund

import (
	"testing"
)

func TestFundPortfolioHoldEm(t *testing.T) {
	// 测试默认参数
	records, err := FundPortfolioHoldEm("", "")
	if err != nil {
		t.Fatalf("FundPortfolioHoldEm failed: %v", err)
	}

	// 允许空数据
	t.Logf("FundPortfolioHoldEm returned %d records", len(records))

	// 如果有数据，验证字段
	if len(records) > 0 {
		first := records[0]
		requiredFields := []string{"股票代码", "股票名称", "持仓市值"}
		for _, field := range requiredFields {
			if _, ok := first[field]; !ok {
				t.Errorf("Missing required field: %s", field)
			}
		}
	}
}

func TestFundPortfolioBondHoldEm(t *testing.T) {
	records, err := FundPortfolioBondHoldEm("000001", "2023")
	if err != nil {
		t.Logf("FundPortfolioBondHoldEm warning: %v (may be unavailable)", err)
		return
	}

	// 允许空数据
	t.Logf("FundPortfolioBondHoldEm returned %d records", len(records))

	// 如果有数据，验证结构
	if len(records) > 0 {
		first := records[0]
		requiredFields := []string{"序号", "债券代码", "债券名称", "占净值比例", "持仓市值", "季度"}
		for _, field := range requiredFields {
			if _, ok := first[field]; !ok {
				t.Errorf("Missing required field: %s", field)
			}
		}
	}
}

func TestFundPortfolioIndustryAllocationEm(t *testing.T) {
	// 测试默认参数
	records, err := FundPortfolioIndustryAllocationEm("", "")
	if err != nil {
		t.Fatalf("FundPortfolioIndustryAllocationEm failed: %v", err)
	}

	// 允许空数据
	t.Logf("FundPortfolioIndustryAllocationEm returned %d records", len(records))

	// 如果有数据，验证字段
	if len(records) > 0 {
		first := records[0]
		requiredFields := []string{"季度", "行业名称", "占净值比", "市值"}
		for _, field := range requiredFields {
			if _, ok := first[field]; !ok {
				t.Errorf("Missing required field: %s", field)
			}
		}
	}
}

func TestFundPortfolioChangeEm(t *testing.T) {
	indicators := []string{"累计买入", "累计卖出"}

	for _, indicator := range indicators {
		t.Run(indicator, func(t *testing.T) {
			records, err := FundPortfolioChangeEm("003567", indicator, "2023")
			if err != nil {
				t.Logf("FundPortfolioChangeEm(%s) warning: %v (may be unavailable)", indicator, err)
				return
			}

			// 允许空数据
			t.Logf("FundPortfolioChangeEm(%s) returned %d records", indicator, len(records))

			// 如果有数据，验证结构
			if len(records) > 0 {
				first := records[0]
				requiredFields := []string{"序号", "股票代码", "股票名称", "本期累计买入金额", "占期初基金资产净值比例", "季度"}
				for _, field := range requiredFields {
					if _, ok := first[field]; !ok {
						t.Errorf("Missing required field: %s", field)
					}
				}
			}
		})
	}
}
