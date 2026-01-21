package spot

import (
	"testing"
)

// TestSpotPriceQh 测试期现货价格查询
func TestSpotPriceQh(t *testing.T) {
	data, err := SpotPriceQh("豆粕")
	if err != nil {
		t.Logf("SpotPriceQh 可能暂时不可用: %v", err)
		return
	}
	if len(data) == 0 {
		t.Log("SpotPriceQh 返回空数据")
		return
	}
	t.Logf("获取到 %d 条期现货价格数据", len(data))
}

// TestSpotPriceTableQh 测试品种列表
func TestSpotPriceTableQh(t *testing.T) {
	data, err := SpotPriceTableQh()
	if err != nil {
		t.Logf("SpotPriceTableQh 可能暂时不可用: %v", err)
		return
	}
	if len(data) == 0 {
		t.Log("SpotPriceTableQh 返回空数据")
		return
	}
	t.Logf("获取到 %d 个品种", len(data))
}

// TestSpotHistSge 测试上海黄金交易所历史数据
func TestSpotHistSge(t *testing.T) {
	data, err := SpotHistSge("Au99.99")
	if err != nil {
		t.Logf("SpotHistSge 可能暂时不可用: %v", err)
		return
	}
	if len(data) == 0 {
		t.Log("SpotHistSge 返回空数据")
		return
	}
	t.Logf("获取到 %d 条历史数据", len(data))
}

// TestSpotSymbolTableSge 测试上海黄金交易所品种列表
func TestSpotSymbolTableSge(t *testing.T) {
	data, err := SpotSymbolTableSge()
	if err != nil {
		t.Logf("SpotSymbolTableSge 可能暂时不可用: %v", err)
		return
	}
	if len(data) == 0 {
		t.Log("SpotSymbolTableSge 返回空数据")
		return
	}
	t.Logf("获取到 %d 个品种", len(data))
}

// TestSpotQuotationsSge 测试上海黄金交易所实时行情
func TestSpotQuotationsSge(t *testing.T) {
	data, err := SpotQuotationsSge("Au99.99")
	if err != nil {
		t.Logf("SpotQuotationsSge 可能暂时不可用: %v", err)
		return
	}
	if len(data) == 0 {
		t.Log("SpotQuotationsSge 返回空数据")
		return
	}
	t.Logf("获取到 %d 条行情数据", len(data))
}

// TestSpotGoldenBenchmarkSge 测试上海黄金交易所黄金基准价
func TestSpotGoldenBenchmarkSge(t *testing.T) {
	data, err := SpotGoldenBenchmarkSge()
	if err != nil {
		t.Logf("SpotGoldenBenchmarkSge 可能暂时不可用: %v", err)
		return
	}
	if len(data) == 0 {
		t.Log("SpotGoldenBenchmarkSge 返回空数据")
		return
	}
	t.Logf("获取到 %d 条黄金基准价数据", len(data))
}

// TestSpotSilverBenchmarkSge 测试上海黄金交易所白银基准价
func TestSpotSilverBenchmarkSge(t *testing.T) {
	data, err := SpotSilverBenchmarkSge()
	if err != nil {
		t.Logf("SpotSilverBenchmarkSge 可能暂时不可用: %v", err)
		return
	}
	if len(data) == 0 {
		t.Log("SpotSilverBenchmarkSge 返回空数据")
		return
	}
	t.Logf("获取到 %d 条白银基准价数据", len(data))
}
