package currency

import (
	"os"
	"testing"
)

// TestCurrencyBocSina 测试中国银行外汇牌价查询
func TestCurrencyBocSina(t *testing.T) {
	rates, err := CurrencyBocSina("美元", "20230801", "20230810")
	if err != nil {
		t.Fatalf("CurrencyBocSina 失败: %v", err)
	}

	if len(rates) == 0 {
		t.Fatal("未获取到任何数据")
	}

	t.Logf("中国银行外汇牌价数据行数: %d", len(rates))

	// 验证第一条数据
	first := rates[0]
	t.Logf("第一条数据: 日期=%s, 汇买价=%.4f, 中间价=%.4f",
		first.Date.Format("2006-01-02"), first.BuyingRate, first.MiddleRate)

	if first.Date.IsZero() {
		t.Error("日期解析失败")
	}
}

// TestCurrencyBocSafe 测试人民币汇率中间价查询
func TestCurrencyBocSafe(t *testing.T) {
	rates, err := CurrencyBocSafe()
	if err != nil {
		t.Fatalf("CurrencyBocSafe 失败: %v", err)
	}

	if len(rates) == 0 {
		t.Fatal("未获取到任何数据")
	}

	t.Logf("人民币汇率中间价数据行数: %d", len(rates))

	// 验证第一条数据
	if len(rates) > 0 {
		first := rates[0]
		t.Logf("第一条数据字段数: %d", len(first))
		for k, v := range first {
			t.Logf("  %s: %v", k, v)
			break // 只打印一个字段示例
		}
	}
}

// TestCurrencyLatest 测试最新汇率查询（需要API Key）
func TestCurrencyLatest(t *testing.T) {
	apiKey := os.Getenv("CURRENCYSCOOP_API_KEY")
	if apiKey == "" {
		t.Skip("跳过测试: 需要设置 CURRENCYSCOOP_API_KEY 环境变量")
	}

	rates, err := CurrencyLatest("USD", "CNY,EUR", apiKey)
	if err != nil {
		t.Fatalf("CurrencyLatest 失败: %v", err)
	}

	if len(rates) == 0 {
		t.Fatal("未获取到任何数据")
	}

	t.Logf("最新汇率数据行数: %d", len(rates))

	for _, rate := range rates {
		t.Logf("%s: %.4f (日期=%s)", rate.Currency, rate.Value, rate.Date.Format("2006-01-02"))
	}
}

// TestCurrencyHistory 测试历史汇率查询（需要API Key）
func TestCurrencyHistory(t *testing.T) {
	apiKey := os.Getenv("CURRENCYSCOOP_API_KEY")
	if apiKey == "" {
		t.Skip("跳过测试: 需要设置 CURRENCYSCOOP_API_KEY 环境变量")
	}

	rates, err := CurrencyHistory("USD", "2023-02-03", "CNY,EUR", apiKey)
	if err != nil {
		t.Fatalf("CurrencyHistory 失败: %v", err)
	}

	if len(rates) == 0 {
		t.Fatal("未获取到任何数据")
	}

	t.Logf("历史汇率数据行数: %d", len(rates))

	for _, rate := range rates {
		t.Logf("%s: %.4f (日期=%s)", rate.Currency, rate.Value, rate.Date.Format("2006-01-02"))
	}
}

// TestCurrencyTimeSeries 测试汇率时间序列查询（需要API Key和特殊权限）
func TestCurrencyTimeSeries(t *testing.T) {
	apiKey := os.Getenv("CURRENCYSCOOP_API_KEY")
	if apiKey == "" {
		t.Skip("跳过测试: 需要设置 CURRENCYSCOOP_API_KEY 环境变量")
	}

	timeSeries, err := CurrencyTimeSeries("USD", "2023-02-03", "2023-02-05", "CNY", apiKey)
	if err != nil {
		// 时间序列需要特殊权限，失败是预期的
		t.Logf("CurrencyTimeSeries 失败（可能需要特殊权限）: %v", err)
		return
	}

	t.Logf("时间序列数据天数: %d", len(timeSeries))

	for date, rates := range timeSeries {
		t.Logf("日期 %s: %d 个汇率", date, len(rates))
	}
}

// TestCurrencyCurrencies 测试货币列表查询（需要API Key）
func TestCurrencyCurrencies(t *testing.T) {
	apiKey := os.Getenv("CURRENCYSCOOP_API_KEY")
	if apiKey == "" {
		t.Skip("跳过测试: 需要设置 CURRENCYSCOOP_API_KEY 环境变量")
	}

	currencies, err := CurrencyCurrencies("fiat", apiKey)
	if err != nil {
		t.Fatalf("CurrencyCurrencies 失败: %v", err)
	}

	if len(currencies) == 0 {
		t.Fatal("未获取到任何货币数据")
	}

	t.Logf("货币列表数量: %d", len(currencies))

	// 打印前3个货币
	for i := 0; i < 3 && i < len(currencies); i++ {
		curr := currencies[i]
		t.Logf("%s (%s): %s - %s",
			curr.CurrencyCode, curr.Symbol, curr.CurrencyName, curr.Countries)
	}
}

// TestCurrencyConvert 测试货币转换（需要API Key）
func TestCurrencyConvert(t *testing.T) {
	apiKey := os.Getenv("CURRENCYSCOOP_API_KEY")
	if apiKey == "" {
		t.Skip("跳过测试: 需要设置 CURRENCYSCOOP_API_KEY 环境变量")
	}

	result, err := CurrencyConvert("USD", "CNY", 10000, apiKey)
	if err != nil {
		t.Fatalf("CurrencyConvert 失败: %v", err)
	}

	t.Logf("%.2f %s = %.2f %s (时间=%s)",
		result.Amount, result.From, result.Value, result.To,
		result.Timestamp.Format("2006-01-02 15:04:05"))

	if result.From != "USD" || result.To != "CNY" {
		t.Error("货币代码不匹配")
	}

	if result.Value <= 0 {
		t.Error("转换结果应该大于0")
	}
}

// TestCurrencyBocSina_UnsupportedCurrency 测试不支持的货币
func TestCurrencyBocSina_UnsupportedCurrency(t *testing.T) {
	_, err := CurrencyBocSina("火星币", "20230801", "20230810")
	if err == nil {
		t.Fatal("应该返回错误：不支持的货币")
	}

	t.Logf("正确处理不支持的货币: %v", err)
}
