package economic

import (
	"testing"
)

func TestMacroGermanyIFO(t *testing.T) {
	df, err := MacroGermanyIFO()
	if err != nil {
		t.Fatalf("MacroGermanyIFO 失败: %v", err)
	}
	t.Logf("IFO商业景气指数数据行数: %d, 列数: %d", df.Nrow(), df.Ncol())
	if df.Nrow() == 0 {
		t.Error("数据为空")
	}
}

func TestMacroGermanyCPIMonthly(t *testing.T) {
	df, err := MacroGermanyCPIMonthly()
	if err != nil {
		t.Fatalf("MacroGermanyCPIMonthly 失败: %v", err)
	}
	t.Logf("CPI月率数据行数: %d", df.Nrow())
}

func TestMacroGermanyCPIYearly(t *testing.T) {
	df, err := MacroGermanyCPIYearly()
	if err != nil {
		t.Fatalf("MacroGermanyCPIYearly 失败: %v", err)
	}
	t.Logf("CPI年率数据行数: %d", df.Nrow())
}

func TestMacroGermanyTradeAdjusted(t *testing.T) {
	df, err := MacroGermanyTradeAdjusted()
	if err != nil {
		t.Fatalf("MacroGermanyTradeAdjusted 失败: %v", err)
	}
	t.Logf("贸易帐数据行数: %d", df.Nrow())
}

func TestMacroGermanyGDP(t *testing.T) {
	df, err := MacroGermanyGDP()
	if err != nil {
		t.Fatalf("MacroGermanyGDP 失败: %v", err)
	}
	t.Logf("GDP数据行数: %d", df.Nrow())
}

func TestMacroGermanyRetailSaleMonthly(t *testing.T) {
	df, err := MacroGermanyRetailSaleMonthly()
	if err != nil {
		t.Fatalf("MacroGermanyRetailSaleMonthly 失败: %v", err)
	}
	t.Logf("零售销售月率数据行数: %d", df.Nrow())
}

func TestMacroGermanyRetailSaleYearly(t *testing.T) {
	df, err := MacroGermanyRetailSaleYearly()
	if err != nil {
		t.Fatalf("MacroGermanyRetailSaleYearly 失败: %v", err)
	}
	t.Logf("零售销售年率数据行数: %d", df.Nrow())
}

func TestMacroGermanyZEW(t *testing.T) {
	df, err := MacroGermanyZEW()
	if err != nil {
		t.Fatalf("MacroGermanyZEW 失败: %v", err)
	}
	t.Logf("ZEW经济景气指数数据行数: %d", df.Nrow())
}

func TestMacroGermanyCore(t *testing.T) {
	df, err := MacroGermanyCore("EMG00179154")
	if err != nil {
		t.Fatalf("MacroGermanyCore 失败: %v", err)
	}
	t.Logf("核心指标数据行数: %d, 列数: %d", df.Nrow(), df.Ncol())
	if df.Nrow() == 0 {
		t.Error("数据为空")
	}
}
