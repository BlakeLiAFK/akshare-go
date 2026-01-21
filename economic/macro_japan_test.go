package economic

import (
	"testing"
)

func TestMacroJapanBankRate(t *testing.T) {
	df, err := MacroJapanBankRate()
	if err != nil {
		t.Fatalf("MacroJapanBankRate 失败: %v", err)
	}

	t.Logf("央行利率决议数据行数: %d, 列数: %d", df.Nrow(), df.Ncol())
	t.Logf("列名: %v", df.Names())

	if df.Nrow() == 0 {
		t.Error("数据为空")
	}
}

func TestMacroJapanCPIYearly(t *testing.T) {
	df, err := MacroJapanCPIYearly()
	if err != nil {
		t.Fatalf("MacroJapanCPIYearly 失败: %v", err)
	}

	t.Logf("CPI年率数据行数: %d, 列数: %d", df.Nrow(), df.Ncol())
}

func TestMacroJapanCoreCPIYearly(t *testing.T) {
	df, err := MacroJapanCoreCPIYearly()
	if err != nil {
		t.Fatalf("MacroJapanCoreCPIYearly 失败: %v", err)
	}

	t.Logf("核心CPI年率数据行数: %d, 列数: %d", df.Nrow(), df.Ncol())
}

func TestMacroJapanUnemploymentRate(t *testing.T) {
	df, err := MacroJapanUnemploymentRate()
	if err != nil {
		t.Fatalf("MacroJapanUnemploymentRate 失败: %v", err)
	}

	t.Logf("失业率数据行数: %d, 列数: %d", df.Nrow(), df.Ncol())
}

func TestMacroJapanHeadIndicator(t *testing.T) {
	df, err := MacroJapanHeadIndicator()
	if err != nil {
		t.Fatalf("MacroJapanHeadIndicator 失败: %v", err)
	}

	t.Logf("领先指标终值数据行数: %d, 列数: %d", df.Nrow(), df.Ncol())
}

func TestMacroJapanCore(t *testing.T) {
	df, err := MacroJapanCore("EMG00342252")
	if err != nil {
		t.Fatalf("MacroJapanCore 失败: %v", err)
	}
	t.Logf("核心指标数据行数: %d, 列数: %d", df.Nrow(), df.Ncol())
	if df.Nrow() == 0 {
		t.Error("数据为空")
	}
}
