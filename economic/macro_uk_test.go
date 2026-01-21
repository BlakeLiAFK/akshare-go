package economic

import (
	"testing"
)

func TestMacroUKHalifaxMonthly(t *testing.T) {
	df, err := MacroUKHalifaxMonthly()
	if err != nil {
		t.Fatalf("MacroUKHalifaxMonthly 失败: %v", err)
	}
	t.Logf("Halifax房价指数月率数据行数: %d, 列数: %d", df.Nrow(), df.Ncol())
	if df.Nrow() == 0 {
		t.Error("数据为空")
	}
}

func TestMacroUKHalifaxYearly(t *testing.T) {
	df, err := MacroUKHalifaxYearly()
	if err != nil {
		t.Fatalf("MacroUKHalifaxYearly 失败: %v", err)
	}
	t.Logf("Halifax房价指数年率数据行数: %d", df.Nrow())
}

func TestMacroUKTrade(t *testing.T) {
	df, err := MacroUKTrade()
	if err != nil {
		t.Fatalf("MacroUKTrade 失败: %v", err)
	}
	t.Logf("贸易帐数据行数: %d", df.Nrow())
}

func TestMacroUKBankRate(t *testing.T) {
	df, err := MacroUKBankRate()
	if err != nil {
		t.Fatalf("MacroUKBankRate 失败: %v", err)
	}
	t.Logf("央行利率决议数据行数: %d", df.Nrow())
}

func TestMacroUKCoreCPIYearly(t *testing.T) {
	df, err := MacroUKCoreCPIYearly()
	if err != nil {
		t.Fatalf("MacroUKCoreCPIYearly 失败: %v", err)
	}
	t.Logf("核心CPI年率数据行数: %d", df.Nrow())
}

func TestMacroUKCoreCPIMonthly(t *testing.T) {
	df, err := MacroUKCoreCPIMonthly()
	if err != nil {
		t.Fatalf("MacroUKCoreCPIMonthly 失败: %v", err)
	}
	t.Logf("核心CPI月率数据行数: %d", df.Nrow())
}

func TestMacroUKCPIYearly(t *testing.T) {
	df, err := MacroUKCPIYearly()
	if err != nil {
		t.Fatalf("MacroUKCPIYearly 失败: %v", err)
	}
	t.Logf("CPI年率数据行数: %d", df.Nrow())
}

func TestMacroUKCPIMonthly(t *testing.T) {
	df, err := MacroUKCPIMonthly()
	if err != nil {
		t.Fatalf("MacroUKCPIMonthly 失败: %v", err)
	}
	t.Logf("CPI月率数据行数: %d", df.Nrow())
}

func TestMacroUKRetailMonthly(t *testing.T) {
	df, err := MacroUKRetailMonthly()
	if err != nil {
		t.Fatalf("MacroUKRetailMonthly 失败: %v", err)
	}
	t.Logf("零售销售月率数据行数: %d", df.Nrow())
}

func TestMacroUKRetailYearly(t *testing.T) {
	df, err := MacroUKRetailYearly()
	if err != nil {
		t.Fatalf("MacroUKRetailYearly 失败: %v", err)
	}
	t.Logf("零售销售年率数据行数: %d", df.Nrow())
}

func TestMacroUKRightmoveYearly(t *testing.T) {
	df, err := MacroUKRightmoveYearly()
	if err != nil {
		t.Fatalf("MacroUKRightmoveYearly 失败: %v", err)
	}
	t.Logf("Rightmove房价指数年率数据行数: %d", df.Nrow())
}

func TestMacroUKRightmoveMonthly(t *testing.T) {
	df, err := MacroUKRightmoveMonthly()
	if err != nil {
		t.Fatalf("MacroUKRightmoveMonthly 失败: %v", err)
	}
	t.Logf("Rightmove房价指数月率数据行数: %d", df.Nrow())
}

func TestMacroUKGDPQuarterly(t *testing.T) {
	df, err := MacroUKGDPQuarterly()
	if err != nil {
		t.Fatalf("MacroUKGDPQuarterly 失败: %v", err)
	}
	t.Logf("GDP季率数据行数: %d", df.Nrow())
}

func TestMacroUKGDPYearly(t *testing.T) {
	df, err := MacroUKGDPYearly()
	if err != nil {
		t.Fatalf("MacroUKGDPYearly 失败: %v", err)
	}
	t.Logf("GDP年率数据行数: %d", df.Nrow())
}

func TestMacroUKUnemploymentRate(t *testing.T) {
	df, err := MacroUKUnemploymentRate()
	if err != nil {
		t.Fatalf("MacroUKUnemploymentRate 失败: %v", err)
	}
	t.Logf("失业率数据行数: %d", df.Nrow())
}

func TestMacroUKCore(t *testing.T) {
	df, err := MacroUKCore("EMG00010348")
	if err != nil {
		t.Fatalf("MacroUKCore 失败: %v", err)
	}
	t.Logf("核心指标数据行数: %d, 列数: %d", df.Nrow(), df.Ncol())
	if df.Nrow() == 0 {
		t.Error("数据为空")
	}
}
