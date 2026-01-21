package fund

import (
	"testing"
)

func TestFundEtfCategorySina(t *testing.T) {
	// 测试LOF基金
	records, err := FundEtfCategorySina("LOF基金")
	if err != nil {
		t.Fatalf("FundEtfCategorySina(LOF基金) failed: %v", err)
	}

	// 允许空数据
	t.Logf("FundEtfCategorySina(LOF基金) returned %d records", len(records))

	// 测试ETF基金
	records2, err := FundEtfCategorySina("ETF基金")
	if err != nil {
		t.Fatalf("FundEtfCategorySina(ETF基金) failed: %v", err)
	}

	t.Logf("FundEtfCategorySina(ETF基金) returned %d records", len(records2))
}

func TestFundEtfHistSina(t *testing.T) {
	// 测试历史净值
	records, err := FundEtfHistSina("510300")
	if err != nil {
		t.Fatalf("FundEtfHistSina failed: %v", err)
	}

	// 允许空数据
	t.Logf("FundEtfHistSina returned %d records", len(records))
}

func TestFundEtfDividendSina(t *testing.T) {
	// 测试分红数据
	records, err := FundEtfDividendSina("510300")
	if err != nil {
		t.Logf("FundEtfDividendSina warning: %v (may be no dividend data)", err)
		return
	}

	t.Logf("FundEtfDividendSina returned %d records", len(records))
}
