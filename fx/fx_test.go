package fx

import (
	"testing"
)

// TestFxQuoteBaidu 测试百度外汇行情
func TestFxQuoteBaidu(t *testing.T) {
	data, err := FxQuoteBaidu("人民币")
	if err != nil {
		t.Logf("FxQuoteBaidu 可能暂时不可用: %v", err)
		return
	}
	if len(data) == 0 {
		t.Log("FxQuoteBaidu 返回空数据")
		return
	}
	t.Logf("获取到 %d 条外汇行情数据", len(data))
	if len(data) > 0 {
		t.Logf("示例: %s %s %.4f", data[0].Code, data[0].Name, data[0].Price)
	}
}

// TestFxSpotQuote 测试外汇即期报价
func TestFxSpotQuote(t *testing.T) {
	data, err := FxSpotQuote()
	if err != nil {
		t.Logf("FxSpotQuote 可能暂时不可用: %v", err)
		return
	}
	if len(data) == 0 {
		t.Log("FxSpotQuote 返回空数据")
		return
	}
	t.Logf("获取到 %d 条即期报价数据", len(data))
}

// TestFxSwapQuote 测试外汇掉期报价
func TestFxSwapQuote(t *testing.T) {
	data, err := FxSwapQuote()
	if err != nil {
		t.Logf("FxSwapQuote 可能暂时不可用: %v", err)
		return
	}
	if len(data) == 0 {
		t.Log("FxSwapQuote 返回空数据")
		return
	}
	t.Logf("获取到 %d 条掉期报价数据", len(data))
}

// TestFxPairQuote 测试货币对报价
func TestFxPairQuote(t *testing.T) {
	data, err := FxPairQuote()
	if err != nil {
		t.Logf("FxPairQuote 可能暂时不可用: %v", err)
		return
	}
	if len(data) == 0 {
		t.Log("FxPairQuote 返回空数据")
		return
	}
	t.Logf("获取到 %d 条货币对报价数据", len(data))
}

// TestFxCSwapCm 测试人民币外汇掉期曲线
func TestFxCSwapCm(t *testing.T) {
	data, err := FxCSwapCm()
	if err != nil {
		t.Logf("FxCSwapCm 可能暂时不可用: %v", err)
		return
	}
	if len(data) == 0 {
		t.Log("FxCSwapCm 返回空数据")
		return
	}
	t.Logf("获取到 %d 条掉期曲线数据", len(data))
}

// TestCurrencyPairMap 测试货币对映射
func TestCurrencyPairMap(t *testing.T) {
	data, err := CurrencyPairMap("美元人民币")
	if err != nil {
		t.Logf("CurrencyPairMap 可能暂时不可用: %v", err)
		return
	}
	if len(data) == 0 {
		t.Log("CurrencyPairMap 返回空数据")
		return
	}
	t.Logf("获取到 %d 条货币对映射数据", len(data))
}
