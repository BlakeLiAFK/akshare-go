package economic

import (
	"testing"
)

func TestMacroStockFinance(t *testing.T) {
	df, err := MacroStockFinance()
	if err != nil {
		t.Fatalf("MacroStockFinance 失败: %v", err)
	}
	t.Logf("股票筹资数据行数: %d, 列数: %d", df.Nrow(), df.Ncol())
	if df.Nrow() == 0 {
		t.Error("数据为空")
	}
}

func TestMacroRMBLoan(t *testing.T) {
	df, err := MacroRMBLoan()
	if err != nil {
		t.Fatalf("MacroRMBLoan 失败: %v", err)
	}
	t.Logf("新增人民币贷款数据行数: %d, 列数: %d", df.Nrow(), df.Ncol())
	if df.Nrow() == 0 {
		t.Error("数据为空")
	}
}

func TestMacroRMBDeposit(t *testing.T) {
	df, err := MacroRMBDeposit()
	if err != nil {
		t.Fatalf("MacroRMBDeposit 失败: %v", err)
	}
	t.Logf("人民币存款余额数据行数: %d, 列数: %d", df.Nrow(), df.Ncol())
	if df.Nrow() == 0 {
		t.Error("数据为空")
	}
}
