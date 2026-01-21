package economic

import (
	"testing"
)

func TestMacroCanadaNewHouseRate(t *testing.T) {
	df, err := MacroCanadaNewHouseRate()
	if err != nil {
		t.Fatalf("MacroCanadaNewHouseRate 失败: %v", err)
	}

	t.Logf("新屋开工数据行数: %d, 列数: %d", df.Nrow(), df.Ncol())
	t.Logf("列名: %v", df.Names())

	if df.Nrow() == 0 {
		t.Error("数据为空")
	}
}

func TestMacroCanadaUnemploymentRate(t *testing.T) {
	df, err := MacroCanadaUnemploymentRate()
	if err != nil {
		t.Fatalf("MacroCanadaUnemploymentRate 失败: %v", err)
	}

	t.Logf("失业率数据行数: %d, 列数: %d", df.Nrow(), df.Ncol())
}

func TestMacroCanadaTrade(t *testing.T) {
	df, err := MacroCanadaTrade()
	if err != nil {
		t.Fatalf("MacroCanadaTrade 失败: %v", err)
	}

	t.Logf("贸易帐数据行数: %d, 列数: %d", df.Nrow(), df.Ncol())
}

func TestMacroCanadaRetailRateMonthly(t *testing.T) {
	df, err := MacroCanadaRetailRateMonthly()
	if err != nil {
		t.Fatalf("MacroCanadaRetailRateMonthly 失败: %v", err)
	}

	t.Logf("零售销售月率数据行数: %d, 列数: %d", df.Nrow(), df.Ncol())
}

func TestMacroCanadaBankRate(t *testing.T) {
	df, err := MacroCanadaBankRate()
	if err != nil {
		t.Fatalf("MacroCanadaBankRate 失败: %v", err)
	}

	t.Logf("央行利率决议数据行数: %d, 列数: %d", df.Nrow(), df.Ncol())
}

func TestMacroCanadaCoreCPIYearly(t *testing.T) {
	df, err := MacroCanadaCoreCPIYearly()
	if err != nil {
		t.Fatalf("MacroCanadaCoreCPIYearly 失败: %v", err)
	}

	t.Logf("核心CPI年率数据行数: %d, 列数: %d", df.Nrow(), df.Ncol())
}

func TestMacroCanadaCoreCPIMonthly(t *testing.T) {
	df, err := MacroCanadaCoreCPIMonthly()
	if err != nil {
		t.Fatalf("MacroCanadaCoreCPIMonthly 失败: %v", err)
	}

	t.Logf("核心CPI月率数据行数: %d, 列数: %d", df.Nrow(), df.Ncol())
}

func TestMacroCanadaCPIYearly(t *testing.T) {
	df, err := MacroCanadaCPIYearly()
	if err != nil {
		t.Fatalf("MacroCanadaCPIYearly 失败: %v", err)
	}

	t.Logf("CPI年率数据行数: %d, 列数: %d", df.Nrow(), df.Ncol())
}

func TestMacroCanadaCPIMonthly(t *testing.T) {
	df, err := MacroCanadaCPIMonthly()
	if err != nil {
		t.Fatalf("MacroCanadaCPIMonthly 失败: %v", err)
	}

	t.Logf("CPI月率数据行数: %d, 列数: %d", df.Nrow(), df.Ncol())
}

func TestMacroCanadaGDPMonthly(t *testing.T) {
	df, err := MacroCanadaGDPMonthly()
	if err != nil {
		t.Fatalf("MacroCanadaGDPMonthly 失败: %v", err)
	}

	t.Logf("GDP月率数据行数: %d, 列数: %d", df.Nrow(), df.Ncol())
}
