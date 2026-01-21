package economic

import (
	"testing"
)

func TestMacroSwissSVME(t *testing.T) {
	df, err := MacroSwissSVME()
	if err != nil {
		t.Fatalf("MacroSwissSVME 失败: %v", err)
	}
	t.Logf("SVME采购经理人指数数据行数: %d, 列数: %d", df.Nrow(), df.Ncol())
	if df.Nrow() == 0 {
		t.Error("数据为空")
	}
}

func TestMacroSwissTrade(t *testing.T) {
	df, err := MacroSwissTrade()
	if err != nil {
		t.Fatalf("MacroSwissTrade 失败: %v", err)
	}
	t.Logf("贸易帐数据行数: %d", df.Nrow())
}

func TestMacroSwissCPIYearly(t *testing.T) {
	df, err := MacroSwissCPIYearly()
	if err != nil {
		t.Fatalf("MacroSwissCPIYearly 失败: %v", err)
	}
	t.Logf("CPI年率数据行数: %d", df.Nrow())
}

func TestMacroSwissGDPQuarterly(t *testing.T) {
	df, err := MacroSwissGDPQuarterly()
	if err != nil {
		t.Fatalf("MacroSwissGDPQuarterly 失败: %v", err)
	}
	t.Logf("GDP季率数据行数: %d", df.Nrow())
}

func TestMacroSwissGDPYearly(t *testing.T) {
	df, err := MacroSwissGDPYearly()
	if err != nil {
		t.Fatalf("MacroSwissGDPYearly 失败: %v", err)
	}
	t.Logf("GDP年率数据行数: %d", df.Nrow())
}

func TestMacroSwissBankRate(t *testing.T) {
	df, err := MacroSwissBankRate()
	if err != nil {
		t.Fatalf("MacroSwissBankRate 失败: %v", err)
	}
	t.Logf("央行利率决议数据行数: %d", df.Nrow())
}

func TestMacroSwissCore(t *testing.T) {
	df, err := MacroSwissCore("EMG00341602")
	if err != nil {
		t.Fatalf("MacroSwissCore 失败: %v", err)
	}
	t.Logf("核心指标数据行数: %d, 列数: %d", df.Nrow(), df.Ncol())
	if df.Nrow() == 0 {
		t.Error("数据为空")
	}
}
