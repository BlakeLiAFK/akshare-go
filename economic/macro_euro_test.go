package economic

import (
	"testing"
)

func TestMacroEuroGDPYoY(t *testing.T) {
	df, err := MacroEuroGDPYoY()
	if err != nil {
		t.Fatalf("MacroEuroGDPYoY 失败: %v", err)
	}
	t.Logf("欧元区季度GDP年率数据行数: %d, 列数: %d", df.Nrow(), df.Ncol())
	if df.Nrow() == 0 {
		t.Error("数据为空")
	}
}

func TestMacroEuroCPIMoM(t *testing.T) {
	df, err := MacroEuroCPIMoM()
	if err != nil {
		t.Fatalf("MacroEuroCPIMoM 失败: %v", err)
	}
	t.Logf("欧元区CPI月率数据行数: %d", df.Nrow())
}

func TestMacroEuroCPIYoY(t *testing.T) {
	df, err := MacroEuroCPIYoY()
	if err != nil {
		t.Fatalf("MacroEuroCPIYoY 失败: %v", err)
	}
	t.Logf("欧元区CPI年率数据行数: %d", df.Nrow())
}

func TestMacroEuroPPIMoM(t *testing.T) {
	df, err := MacroEuroPPIMoM()
	if err != nil {
		t.Fatalf("MacroEuroPPIMoM 失败: %v", err)
	}
	t.Logf("欧元区PPI月率数据行数: %d", df.Nrow())
}

func TestMacroEuroRetailSalesMoM(t *testing.T) {
	df, err := MacroEuroRetailSalesMoM()
	if err != nil {
		t.Fatalf("MacroEuroRetailSalesMoM 失败: %v", err)
	}
	t.Logf("欧元区零售销售月率数据行数: %d", df.Nrow())
}

func TestMacroEuroEmploymentChangeQoQ(t *testing.T) {
	df, err := MacroEuroEmploymentChangeQoQ()
	if err != nil {
		t.Fatalf("MacroEuroEmploymentChangeQoQ 失败: %v", err)
	}
	t.Logf("欧元区季调后就业人数季率数据行数: %d", df.Nrow())
}

func TestMacroEuroUnemploymentRateMoM(t *testing.T) {
	df, err := MacroEuroUnemploymentRateMoM()
	if err != nil {
		t.Fatalf("MacroEuroUnemploymentRateMoM 失败: %v", err)
	}
	t.Logf("欧元区失业率数据行数: %d", df.Nrow())
}

func TestMacroEuroTradeBalance(t *testing.T) {
	df, err := MacroEuroTradeBalance()
	if err != nil {
		t.Fatalf("MacroEuroTradeBalance 失败: %v", err)
	}
	t.Logf("欧元区未季调贸易帐数据行数: %d", df.Nrow())
}

func TestMacroEuroCurrentAccountMoM(t *testing.T) {
	df, err := MacroEuroCurrentAccountMoM()
	if err != nil {
		t.Fatalf("MacroEuroCurrentAccountMoM 失败: %v", err)
	}
	t.Logf("欧元区经常帐数据行数: %d", df.Nrow())
}

func TestMacroEuroIndustrialProductionMoM(t *testing.T) {
	df, err := MacroEuroIndustrialProductionMoM()
	if err != nil {
		t.Fatalf("MacroEuroIndustrialProductionMoM 失败: %v", err)
	}
	t.Logf("欧元区工业产出月率数据行数: %d", df.Nrow())
}

func TestMacroEuroManufacturingPMI(t *testing.T) {
	df, err := MacroEuroManufacturingPMI()
	if err != nil {
		t.Fatalf("MacroEuroManufacturingPMI 失败: %v", err)
	}
	t.Logf("欧元区制造业PMI初值数据行数: %d", df.Nrow())
}

func TestMacroEuroServicesPMI(t *testing.T) {
	df, err := MacroEuroServicesPMI()
	if err != nil {
		t.Fatalf("MacroEuroServicesPMI 失败: %v", err)
	}
	t.Logf("欧元区服务业PMI终值数据行数: %d", df.Nrow())
}

func TestMacroEuroZEWEconomicSentiment(t *testing.T) {
	df, err := MacroEuroZEWEconomicSentiment()
	if err != nil {
		t.Fatalf("MacroEuroZEWEconomicSentiment 失败: %v", err)
	}
	t.Logf("欧元区ZEW经济景气指数数据行数: %d", df.Nrow())
}

func TestMacroEuroSentixInvestorConfidence(t *testing.T) {
	df, err := MacroEuroSentixInvestorConfidence()
	if err != nil {
		t.Fatalf("MacroEuroSentixInvestorConfidence 失败: %v", err)
	}
	t.Logf("欧元区Sentix投资者信心指数数据行数: %d", df.Nrow())
}

func TestMacroEuroLMEHolding(t *testing.T) {
	df, err := MacroEuroLMEHolding()
	if err != nil {
		t.Fatalf("MacroEuroLMEHolding 失败: %v", err)
	}
	t.Logf("LME持仓报告数据行数: %d, 列数: %d", df.Nrow(), df.Ncol())
}

func TestMacroEuroLMEStock(t *testing.T) {
	df, err := MacroEuroLMEStock()
	if err != nil {
		t.Fatalf("MacroEuroLMEStock 失败: %v", err)
	}
	t.Logf("LME库存报告数据行数: %d, 列数: %d", df.Nrow(), df.Ncol())
}
