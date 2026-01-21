package economic

import (
	"testing"
)

func TestMacroAustraliaRetailRateMonthly(t *testing.T) {
	df, err := MacroAustraliaRetailRateMonthly()
	if err != nil {
		t.Fatalf("MacroAustraliaRetailRateMonthly 失败: %v", err)
	}

	t.Logf("零售销售月率数据行数: %d, 列数: %d", df.Nrow(), df.Ncol())
	t.Logf("列名: %v", df.Names())

	if df.Nrow() == 0 {
		t.Error("数据为空")
	}
}

func TestMacroAustraliaTrade(t *testing.T) {
	df, err := MacroAustraliaTrade()
	if err != nil {
		t.Fatalf("MacroAustraliaTrade 失败: %v", err)
	}

	t.Logf("贸易帐数据行数: %d, 列数: %d", df.Nrow(), df.Ncol())
}

func TestMacroAustraliaUnemploymentRate(t *testing.T) {
	df, err := MacroAustraliaUnemploymentRate()
	if err != nil {
		t.Fatalf("MacroAustraliaUnemploymentRate 失败: %v", err)
	}

	t.Logf("失业率数据行数: %d, 列数: %d", df.Nrow(), df.Ncol())
}

func TestMacroAustraliaPPIQuarterly(t *testing.T) {
	df, err := MacroAustraliaPPIQuarterly()
	if err != nil {
		t.Fatalf("MacroAustraliaPPIQuarterly 失败: %v", err)
	}

	t.Logf("PPI季率数据行数: %d, 列数: %d", df.Nrow(), df.Ncol())
}

func TestMacroAustraliaCPIQuarterly(t *testing.T) {
	df, err := MacroAustraliaCPIQuarterly()
	if err != nil {
		t.Fatalf("MacroAustraliaCPIQuarterly 失败: %v", err)
	}

	t.Logf("CPI季率数据行数: %d, 列数: %d", df.Nrow(), df.Ncol())
}

func TestMacroAustraliaCPIYearly(t *testing.T) {
	df, err := MacroAustraliaCPIYearly()
	if err != nil {
		t.Fatalf("MacroAustraliaCPIYearly 失败: %v", err)
	}

	t.Logf("CPI年率数据行数: %d, 列数: %d", df.Nrow(), df.Ncol())
}

func TestMacroAustraliaBankRate(t *testing.T) {
	df, err := MacroAustraliaBankRate()
	if err != nil {
		t.Fatalf("MacroAustraliaBankRate 失败: %v", err)
	}

	t.Logf("央行利率决议数据行数: %d, 列数: %d", df.Nrow(), df.Ncol())
}
