package economic

import (
	"testing"
)

func TestMacroChinaHKCPI(t *testing.T) {
	df, err := MacroChinaHKCPI()
	if err != nil {
		t.Fatalf("MacroChinaHKCPI 失败: %v", err)
	}
	t.Logf("消费者物价指数数据行数: %d, 列数: %d", df.Nrow(), df.Ncol())
	if df.Nrow() == 0 {
		t.Error("数据为空")
	}
}

func TestMacroChinaHKCPIRatio(t *testing.T) {
	df, err := MacroChinaHKCPIRatio()
	if err != nil {
		t.Fatalf("MacroChinaHKCPIRatio 失败: %v", err)
	}
	t.Logf("消费者物价指数年率数据行数: %d", df.Nrow())
}

func TestMacroChinaHKRateOfUnemployment(t *testing.T) {
	df, err := MacroChinaHKRateOfUnemployment()
	if err != nil {
		t.Fatalf("MacroChinaHKRateOfUnemployment 失败: %v", err)
	}
	t.Logf("失业率数据行数: %d", df.Nrow())
}

func TestMacroChinaHKGBP(t *testing.T) {
	df, err := MacroChinaHKGBP()
	if err != nil {
		t.Fatalf("MacroChinaHKGBP 失败: %v", err)
	}
	t.Logf("香港GDP数据行数: %d", df.Nrow())
}

func TestMacroChinaHKGBPRatio(t *testing.T) {
	df, err := MacroChinaHKGBPRatio()
	if err != nil {
		t.Fatalf("MacroChinaHKGBPRatio 失败: %v", err)
	}
	t.Logf("香港GDP同比数据行数: %d", df.Nrow())
}

func TestMacroChinaHKBuildingVolume(t *testing.T) {
	df, err := MacroChinaHKBuildingVolume()
	if err != nil {
		t.Fatalf("MacroChinaHKBuildingVolume 失败: %v", err)
	}
	t.Logf("香港楼宇买卖合约数量数据行数: %d", df.Nrow())
}

func TestMacroChinaHKBuildingAmount(t *testing.T) {
	df, err := MacroChinaHKBuildingAmount()
	if err != nil {
		t.Fatalf("MacroChinaHKBuildingAmount 失败: %v", err)
	}
	t.Logf("香港楼宇买卖合约成交金额数据行数: %d", df.Nrow())
}

func TestMacroChinaHKTradeDiffRatio(t *testing.T) {
	df, err := MacroChinaHKTradeDiffRatio()
	if err != nil {
		t.Fatalf("MacroChinaHKTradeDiffRatio 失败: %v", err)
	}
	t.Logf("香港商品贸易差额年率数据行数: %d", df.Nrow())
}

func TestMacroChinaHKPPI(t *testing.T) {
	df, err := MacroChinaHKPPI()
	if err != nil {
		t.Fatalf("MacroChinaHKPPI 失败: %v", err)
	}
	t.Logf("香港制造业PPI年率数据行数: %d", df.Nrow())
}

func TestMacroChinaHKCore(t *testing.T) {
	df, err := MacroChinaHKCore("EMG01336996")
	if err != nil {
		t.Fatalf("MacroChinaHKCore 失败: %v", err)
	}
	t.Logf("核心指标数据行数: %d, 列数: %d", df.Nrow(), df.Ncol())
	if df.Nrow() == 0 {
		t.Error("数据为空")
	}
}
