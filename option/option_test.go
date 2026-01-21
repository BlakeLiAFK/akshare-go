package option

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestOptionCommSymbol 测试商品期权品种列表
func TestOptionCommSymbol(t *testing.T) {
	t.Skip("商品期权网站结构可能已变化，跳过测试")

	result, err := OptionCommSymbol()
	assert.NoError(t, err, "OptionCommSymbol 不应返回错误")
	assert.NotEmpty(t, result, "品种列表不应为空")

	if len(result) > 0 {
		t.Logf("获取到 %d 个品种", len(result))
		t.Logf("第一个品种: %s (%s)", result[0].Name, result[0].Code)
	}
}

// TestOptionCurrentEm 测试东方财富期权行情
func TestOptionCurrentEm(t *testing.T) {
	result, err := OptionCurrentEm()
	assert.NoError(t, err, "OptionCurrentEm 不应返回错误")
	assert.NotEmpty(t, result, "期权行情不应为空")

	if len(result) > 0 {
		t.Logf("获取到 %d 条期权数据", len(result))
		t.Logf("第一条: %s %s 最新价:%.4f", result[0].Code, result[0].Name, result[0].LatestPrice)
	}
}

// TestOptionCurrentCffexEm 测试中金所期权行情
func TestOptionCurrentCffexEm(t *testing.T) {
	result, err := OptionCurrentCffexEm()
	assert.NoError(t, err, "OptionCurrentCffexEm 不应返回错误")

	if len(result) > 0 {
		t.Logf("获取到 %d 条中金所期权数据", len(result))
		t.Logf("第一条: %s %s", result[0].Code, result[0].Name)
	}
}

// TestOptionCurrentDaySse 测试上交所当日合约
func TestOptionCurrentDaySse(t *testing.T) {
	result, err := OptionCurrentDaySse()
	assert.NoError(t, err, "OptionCurrentDaySse 不应返回错误")

	if len(result) > 0 {
		t.Logf("获取到 %d 条上交所期权合约", len(result))
		t.Logf("第一条: %s %s", result[0].ContractId, result[0].ContractSymbol)
	}
}

// TestOptionDailyStatsSse 测试上交所期权每日统计
func TestOptionDailyStatsSse(t *testing.T) {
	result, err := OptionDailyStatsSse("20240626")
	assert.NoError(t, err, "OptionDailyStatsSse 不应返回错误")

	if len(result) > 0 {
		t.Logf("获取到 %d 条统计数据", len(result))
		t.Logf("第一条: %s %s 总成交量:%d", result[0].SecurityCode, result[0].SecurityName, result[0].TotalVolume)
	}
}

// TestOptionDailyStatsSzse 测试深交所期权每日统计
func TestOptionDailyStatsSzse(t *testing.T) {
	result, err := OptionDailyStatsSzse("20240626")
	assert.NoError(t, err, "OptionDailyStatsSzse 不应返回错误")

	if len(result) > 0 {
		t.Logf("获取到 %d 条统计数据", len(result))
		t.Logf("第一条: %s %s", result[0].SecurityCode, result[0].SecurityName)
	}
}

// TestOptionRiskIndicatorSse 测试上交所期权风险指标
func TestOptionRiskIndicatorSse(t *testing.T) {
	result, err := OptionRiskIndicatorSse("20240626")
	assert.NoError(t, err, "OptionRiskIndicatorSse 不应返回错误")

	if len(result) > 0 {
		t.Logf("获取到 %d 条风险指标数据", len(result))
		t.Logf("第一条: %s Delta:%.4f", result[0].ContractSymbol, result[0].Delta)
	}
}

// TestOptionPremiumAnalysisEm 测试期权折溢价分析
func TestOptionPremiumAnalysisEm(t *testing.T) {
	result, err := OptionPremiumAnalysisEm()
	assert.NoError(t, err, "OptionPremiumAnalysisEm 不应返回错误")

	if len(result) > 0 {
		t.Logf("获取到 %d 条折溢价数据", len(result))
		t.Logf("第一条: %s %s 折溢价率:%.2f%%", result[0].OptionCode, result[0].OptionName, result[0].PremiumRate)
	}
}

// TestOptionRiskAnalysisEm 测试期权风险分析
func TestOptionRiskAnalysisEm(t *testing.T) {
	result, err := OptionRiskAnalysisEm()
	assert.NoError(t, err, "OptionRiskAnalysisEm 不应返回错误")

	if len(result) > 0 {
		t.Logf("获取到 %d 条风险分析数据", len(result))
		t.Logf("第一条: %s Delta:%.4f Gamma:%.4f", result[0].OptionName, result[0].Delta, result[0].Gamma)
	}
}

// TestOptionValueAnalysisEm 测试期权价值分析
func TestOptionValueAnalysisEm(t *testing.T) {
	result, err := OptionValueAnalysisEm()
	assert.NoError(t, err, "OptionValueAnalysisEm 不应返回错误")

	if len(result) > 0 {
		t.Logf("获取到 %d 条价值分析数据", len(result))
		t.Logf("第一条: %s 时间价值:%.4f 内在价值:%.4f", result[0].OptionName, result[0].TimeValue, result[0].IntrinsicValue)
	}
}
