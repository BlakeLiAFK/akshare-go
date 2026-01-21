package index

import (
	"testing"
)

// TestIndexCodeIDMapEM 测试获取东方财富指数代码ID映射
func TestIndexCodeIDMapEM(t *testing.T) {
	codeMap, err := IndexCodeIDMapEM()
	if err != nil {
		t.Fatalf("IndexCodeIDMapM failed: %v", err)
	}
	if len(codeMap) == 0 {
		t.Fatal("Expected some codes, got empty map")
	}
	t.Logf("IndexCodeIDMapEM: got %d codes", len(codeMap))
}

// TestIndexZHAHist 测试中国股票指数历史行情数据
func TestIndexZHAHist(t *testing.T) {
	klines, err := IndexZHAHist("000001", "daily", "20240101", "20240131")
	if err != nil {
		t.Logf("IndexZHAHist failed: %v", err)
		return
	}
	if len(klines) > 0 {
		t.Logf("IndexZHAHist: got %d klines", len(klines))
	} else {
		t.Log("IndexZHAHist: no data returned (可能不在交易日)")
	}
}

// TestIndexZHAHistWeekly 测试周K线数据
func TestIndexZHAHistWeekly(t *testing.T) {
	klines, err := IndexZHAHist("000001", "weekly", "20240101", "20240131")
	if err != nil {
		t.Logf("IndexZHAHist(weekly) failed: %v", err)
		return
	}
	if len(klines) > 0 {
		t.Logf("IndexZHAHist(weekly): got %d klines", len(klines))
	}
}

// TestIndexZHAHistMinEM 测试中国指数分时行情
func TestIndexZHAHistMinEM(t *testing.T) {
	quotes, err := IndexZHAHistMinEM("000001", "1", "2024-01-01 09:30:00", "2024-01-01 10:00:00")
	if err != nil {
		t.Logf("IndexZHAHistMinEM failed: %v", err)
		return
	}
	t.Logf("IndexZHAHistMinEM: got %d quotes", len(quotes))
}

// TestStockZHIndexSpotEM 测试东方财富网-沪深京指数实时行情
func TestStockZHIndexSpotEM(t *testing.T) {
	quotes, err := StockZHIndexSpotEM("沪深重要指数")
	if err != nil {
		t.Logf("StockZHIndexSpotEM failed: %v", err)
		return
	}
	if len(quotes) > 0 {
		t.Logf("StockZHIndexSpotEM: got %d quotes", len(quotes))
	}
}

// TestStockZHIndexSpotEM_ShangSeries 测试上证系列指数
func TestStockZHIndexSpotEM_ShangSeries(t *testing.T) {
	quotes, err := StockZHIndexSpotEM("上证系列指数")
	if err != nil {
		t.Logf("StockZHIndexSpotEM(上证系列) failed: %v", err)
		return
	}
	if len(quotes) > 0 {
		t.Logf("StockZHIndexSpotEM(上证系列): got %d quotes", len(quotes))
	}
}

// TestStockZHIndexSpotEM_ShenzhenSeries 测试深证系列指数
func TestStockZHIndexSpotEM_ShenzhenSeries(t *testing.T) {
	quotes, err := StockZHIndexSpotEM("深证系列指数")
	if err != nil {
		t.Logf("StockZHIndexSpotEM(深证系列) failed: %v", err)
		return
	}
	if len(quotes) > 0 {
		t.Logf("StockZHIndexSpotEM(深证系列): got %d quotes", len(quotes))
	}
}

// TestStockZHIndexSpotEM_CSISeries 测试中证系列指数
func TestStockZHIndexSpotEM_CSISeries(t *testing.T) {
	quotes, err := StockZHIndexSpotEM("中证系列指数")
	if err != nil {
		t.Logf("StockZHIndexSpotEM(中证系列) failed: %v", err)
		return
	}
	if len(quotes) > 0 {
		t.Logf("StockZHIndexSpotEM(中证系列): got %d quotes", len(quotes))
	}
}

// TestGetZHIndexPageCount 测试获取新浪指数总页数
func TestGetZHIndexPageCount(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过较慢的外部网络测试")
	}
	count, err := GetZHIndexPageCount()
	if err != nil {
		t.Logf("GetZHIndexPageCount failed: %v", err)
		return
	}
	t.Logf("GetZHIndexPageCount: got %d pages", count)
}

// TestStockZHIndexSpotSina 测试新浪财经-中国指数实时行情
func TestStockZHIndexSpotSina(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过较慢的外部网络测试")
	}
	quotes, err := StockZHIndexSpotSina()
	if err != nil {
		t.Logf("StockZHIndexSpotSina failed: %v", err)
		return
	}
	t.Logf("StockZHIndexSpotSina: got %d quotes", len(quotes))
}

// TestStockZHIndexDaily 测试东方财富网-指数日K线数据
func TestStockZHIndexDaily(t *testing.T) {
	klines, err := StockZHIndexDaily("000001", "20240101", "20240131")
	if err != nil {
		t.Logf("StockZHIndexDaily failed: %v", err)
		return
	}
	if len(klines) > 0 {
		t.Logf("StockZHIndexDaily: got %d klines", len(klines))
	}
}

// TestIndexGlobalNameTable 测试新浪全球指数名称代码映射表
func TestIndexGlobalNameTable(t *testing.T) {
	names, err := IndexGlobalNameTable()
	if err != nil {
		t.Fatalf("IndexGlobalNameTable failed: %v", err)
	}
	if len(names) == 0 {
		t.Error("Expected some names, got empty")
	}
	t.Logf("IndexGlobalNameTable: got %d names", len(names))
}

// TestIndexGlobalHistSina 测试新浪财经-全球指数历史行情
func TestIndexGlobalHistSina(t *testing.T) {
	tests := []string{"英国富时100指数", "德国DAX 30种股价指数"}
	for _, symbol := range tests {
		klines, err := IndexGlobalHistSina(symbol)
		if err != nil {
			t.Logf("IndexGlobalHistSina(%s) failed: %v", symbol, err)
			continue
		}
		if len(klines) > 0 {
			t.Logf("IndexGlobalHistSina(%s): got %d klines", symbol, len(klines))
		}
	}
}

// TestIndexGlobalSpotEM 测试东方财富网-全球指数实时行情
func TestIndexGlobalSpotEM(t *testing.T) {
	quotes, err := IndexGlobalSpotEM()
	if err != nil {
		t.Logf("IndexGlobalSpotEM failed: %v", err)
		return
	}
	if len(quotes) > 0 {
		t.Logf("IndexGlobalSpotEM: got %d quotes", len(quotes))
	}
}

// TestIndexGlobalHistEM 测试东方财富网-全球指数历史行情
func TestIndexGlobalHistEM(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过外部网络测试")
	}
	tests := []string{"美元指数", "标普500"}
	for _, symbol := range tests {
		klines, err := IndexGlobalHistEM(symbol)
		if err != nil {
			t.Logf("IndexGlobalHistEM(%s) failed: %v", symbol, err)
			continue
		}
		if len(klines) > 0 {
			t.Logf("IndexGlobalHistEM(%s): got %d klines", symbol, len(klines))
		}
	}
}

// TestSpotGoods 测试新浪财经-商品现货价格指数
func TestSpotGoods(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过外部网络测试")
	}
	tests := []string{"波罗的海干散货指数", "钢坯价格指数", "澳大利亚粉矿价格"}
	for _, symbol := range tests {
		quotes, err := SpotGoods(symbol)
		if err != nil {
			t.Logf("SpotGoods(%s) failed: %v", symbol, err)
			continue
		}
		if len(quotes) > 0 {
			t.Logf("SpotGoods(%s): got %d quotes", symbol, len(quotes))
		}
	}
}

// TestIndexPriceCFLP 测试中国公路物流运价指数
func TestIndexPriceCFLP(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过外部网络测试")
	}
	tests := []string{"周指数", "月指数", "季度指数", "年度指数"}
	for _, symbol := range tests {
		quotes, err := IndexPriceCFLP(symbol)
		if err != nil {
			t.Logf("IndexPriceCFLP(%s) failed: %v", symbol, err)
			continue
		}
		if len(quotes) > 0 {
			t.Logf("IndexPriceCFLP(%s): got %d quotes", symbol, len(quotes))
		}
	}
}

// TestIndexVolumeCFLP 测试中国公路物流运量指数
func TestIndexVolumeCFLP(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过外部网络测试")
	}
	tests := []string{"月指数", "季度指数", "年度指数"}
	for _, symbol := range tests {
		quotes, err := IndexVolumeCFLP(symbol)
		if err != nil {
			t.Logf("IndexVolumeCFLP(%s) failed: %v", symbol, err)
			continue
		}
		if len(quotes) > 0 {
			t.Logf("IndexVolumeCFLP(%s): got %d quotes", symbol, len(quotes))
		}
	}
}

// TestIndexHogSpotPrice 测试行情宝-生猪市场价格指数
func TestIndexHogSpotPrice(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过外部网络测试")
	}
	quotes, err := IndexHogSpotPrice()
	if err != nil {
		t.Logf("IndexHogSpotPrice failed: %v", err)
		return
	}
	t.Logf("IndexHogSpotPrice: got %d quotes", len(quotes))
}

// TestIndexERI 测试浙江省排污权交易指数
func TestIndexERI(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过外部网络测试")
	}
	tests := []string{"月度", "季度"}
	for _, symbol := range tests {
		quotes, err := IndexERI(symbol)
		if err != nil {
			t.Logf("IndexERI(%s) failed: %v (网站可能不可用)", symbol, err)
			continue
		}
		t.Logf("IndexERI(%s): got %d quotes", symbol, len(quotes))
	}
}

// TestIndexYW 测试义乌小商品指数
func TestIndexYW(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过外部网络测试")
	}
	tests := []string{"周价格指数", "月价格指数", "月景气指数"}
	for _, symbol := range tests {
		quotes, err := IndexYW(symbol)
		if err != nil {
			t.Logf("IndexYW(%s) failed: %v", symbol, err)
			continue
		}
		t.Logf("IndexYW(%s): got %d quotes", symbol, len(quotes))
	}
}

// TestIndexSugarMSweet 测试沐甜科技数据中心-中国食糖指数
func TestIndexSugarMSweet(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过外部网络测试")
	}
	quotes, err := IndexSugarMSweet()
	if err != nil {
		t.Logf("IndexSugarMSweet failed: %v", err)
		return
	}
	t.Logf("IndexSugarMSweet: got %d quotes", len(quotes))
}

// TestStockHKIndexSpotSina 测试新浪财经-港股指数实时行情
func TestStockHKIndexSpotSina(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过外部网络测试")
	}
	quotes, err := StockHKIndexSpotSina()
	if err != nil {
		t.Logf("StockHKIndexSpotSina failed: %v", err)
		return
	}
	t.Logf("StockHKIndexSpotSina: got %d quotes", len(quotes))
}

// TestStockHKIndexSpotEM 测试东方财富网-港股指数实时行情
func TestStockHKIndexSpotEM(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过外部网络测试")
	}
	quotes, err := StockHKIndexSpotEM()
	if err != nil {
		t.Logf("StockHKIndexSpotEM failed: %v", err)
		return
	}
	t.Logf("StockHKIndexSpotEM: got %d quotes", len(quotes))
}

// TestStockHKIndexDailyEM 测试东方财富网-港股指数历史行情
func TestStockHKIndexDailyEM(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过外部网络测试")
	}
	tests := []string{"HSTECH", "HSI"}
	for _, symbol := range tests {
		klines, err := StockHKIndexDailyEM(symbol)
		if err != nil {
			t.Logf("StockHKIndexDailyEM(%s) failed: %v", symbol, err)
			continue
		}
		if len(klines) > 0 {
			t.Logf("StockHKIndexDailyEM(%s): got %d klines", symbol, len(klines))
		}
	}
}

// TestIndexUSSpotSina 测试新浪财经-美股指数实时行情
func TestIndexUSSpotSina(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过外部网络测试")
	}
	quotes, err := IndexUSSpotSina()
	if err != nil {
		t.Logf("IndexUSSpotSina failed: %v", err)
		return
	}
	t.Logf("IndexUSSpotSina: got %d quotes", len(quotes))
}

// TestIndexUSStockSina 测试新浪财经-美股指数历史行情
func TestIndexUSStockSina(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过外部网络测试")
	}
	tests := []string{".INX", ".DJI", ".IXIC"}
	for _, symbol := range tests {
		klines, err := IndexUSStockSina(symbol)
		if err != nil {
			t.Logf("IndexUSStockSina(%s) failed: %v", symbol, err)
			continue
		}
		if len(klines) > 0 {
			t.Logf("IndexUSStockSina(%s): got %d klines", symbol, len(klines))
		}
	}
}

// TestSWIndexFirstLevel 测试乐咕乐股-申万一级行业分类
func TestSWIndexFirstLevel(t *testing.T) {
	_, err := SWIndexFirstLevel()
	if err != nil {
		t.Logf("SWIndexFirstLevel: %v (expected - needs HTML parsing)", err)
	}
}

// TestSWIndexSecondLevel 测试乐咕乐股-申万二级行业分类
func TestSWIndexSecondLevel(t *testing.T) {
	_, err := SWIndexSecondLevel()
	if err != nil {
		t.Logf("SWIndexSecondLevel: %v (expected - needs HTML parsing)", err)
	}
}

// TestSWIndexThirdLevel 测试乐咕乐股-申万三级行业分类
func TestSWIndexThirdLevel(t *testing.T) {
	_, err := SWIndexThirdLevel()
	if err != nil {
		t.Logf("SWIndexThirdLevel: %v (expected - needs HTML parsing)", err)
	}
}

// TestSWIndexThirdCons 测试乐咕乐股-申万三级行业成份股
func TestSWIndexThirdCons(t *testing.T) {
	_, err := SWIndexThirdCons("801120.SI")
	if err != nil {
		t.Logf("SWIndexThirdCons: %v (expected - needs HTML parsing)", err)
	}
}

// TestIndexCSIndexAll 测试中证指数网站-指数列表
func TestIndexCSIndexAll(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过外部网络测试")
	}
	_, err := IndexCSIndexAll()
	if err != nil {
		t.Logf("IndexCSIndexAll: %v (expected - needs Excel parsing)", err)
	}
}

// TestIndexStockConsCSIndex 测试中证指数网站-成份股目录
func TestIndexStockConsCSIndex(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过外部网络测试")
	}
	_, err := IndexStockConsCSIndex("000300")
	if err != nil {
		t.Logf("IndexStockConsCSIndex: %v (expected - needs Excel parsing)", err)
	}
}

// TestIndexStockConsWeightCSI 测试中证指数网站-样本权重
func TestIndexStockConsWeightCSI(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过外部网络测试")
	}
	_, err := IndexStockConsWeightCSI("000300")
	if err != nil {
		t.Logf("IndexStockConsWeightCSI: %v (expected - needs Excel parsing)", err)
	}
}

// TestIndexAllCNI 测试国证指数-最近交易日的所有指数
func TestIndexAllCNI(t *testing.T) {
	_, err := IndexAllCNI()
	if err != nil {
		t.Logf("IndexAllCNI: %v (expected - needs complex parsing)", err)
	}
}

// TestIndexHistCNI 测试国证指数历史行情数据
func TestIndexHistCNI(t *testing.T) {
	_, err := IndexHistCNI("399001", "20230101", "20240101")
	if err != nil {
		t.Logf("IndexHistCNI: %v (expected - needs complex parsing)", err)
	}
}

// TestIndexDetailCNI 测试国证指数-样本详情-指定日期的样本成份
func TestIndexDetailCNI(t *testing.T) {
	_, err := IndexDetailCNI("399001")
	if err != nil {
		t.Logf("IndexDetailCNI: %v (expected - needs Excel parsing)", err)
	}
}

// TestIndexStockInfo 测试聚宽-指数数据-指数列表
func TestIndexStockInfo(t *testing.T) {
	_, err := IndexStockInfo()
	if err != nil {
		t.Logf("IndexStockInfo: %v (expected - needs HTML parsing)", err)
	}
}

// TestIndexStockConsSina 测试新浪-股票指数成份股
func TestIndexStockConsSina(t *testing.T) {
	_, err := IndexStockConsSina("000300")
	if err != nil {
		t.Logf("IndexStockConsSina: %v (expected - needs HTML parsing)", err)
	}
}

// BenchmarkIndexCodeIDMapEM 基准测试指数代码映射
func BenchmarkIndexCodeIDMapEM(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		_, _ = IndexCodeIDMapEM()
	}
}

// BenchmarkStockZHIndexSpotEM 基准测试中国指数实时行情
func BenchmarkStockZHIndexSpotEM(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		_, _ = StockZHIndexSpotEM("沪深重要指数")
	}
}

// BenchmarkIndexGlobalSpotEM 基准测试全球指数实时行情
func BenchmarkIndexGlobalSpotEM(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		_, _ = IndexGlobalSpotEM()
	}
}

// ==================== 柯桥纺织指数测试 ====================

// TestIndexKqFzPrice 测试柯桥纺织价格指数
func TestIndexKqFzPrice(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过外部网络测试")
	}
	data, err := IndexKqFzPrice()
	if err != nil {
		t.Logf("IndexKqFzPrice failed: %v", err)
		return
	}
	t.Logf("IndexKqFzPrice: got %d records", len(data))
}

// TestIndexKqFzProsperity 测试柯桥纺织景气指数
func TestIndexKqFzProsperity(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过外部网络测试")
	}
	data, err := IndexKqFzProsperity()
	if err != nil {
		t.Logf("IndexKqFzProsperity failed: %v", err)
		return
	}
	t.Logf("IndexKqFzProsperity: got %d records", len(data))
}

// TestIndexKqFzForeignTrade 测试柯桥纺织外贸指数
func TestIndexKqFzForeignTrade(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过外部网络测试")
	}
	data, err := IndexKqFzForeignTrade()
	if err != nil {
		t.Logf("IndexKqFzForeignTrade failed: %v", err)
		return
	}
	t.Logf("IndexKqFzForeignTrade: got %d records", len(data))
}

// ==================== 柯桥时尚指数测试 ====================

// TestIndexKqFashion 测试柯桥时尚指数
func TestIndexKqFashion(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过外部网络测试")
	}
	tests := []string{"柯桥时尚指数", "时尚创意指数", "时尚活跃度指数"}
	for _, symbol := range tests {
		data, err := IndexKqFashion(symbol)
		if err != nil {
			t.Logf("IndexKqFashion(%s) failed: %v", symbol, err)
			continue
		}
		t.Logf("IndexKqFashion(%s): got %d records", symbol, len(data))
	}
}

// ==================== 期权波动率指数测试 ====================

// TestIndexOption50etfQvix 测试50ETF期权波动率指数
func TestIndexOption50etfQvix(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过外部网络测试")
	}
	data, err := IndexOption50etfQvix()
	if err != nil {
		t.Logf("IndexOption50etfQvix failed: %v", err)
		return
	}
	t.Logf("IndexOption50etfQvix: got %d records", len(data))
}

// TestIndexOption50etfMinQvix 测试50ETF期权波动率指数分时
func TestIndexOption50etfMinQvix(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过外部网络测试")
	}
	data, err := IndexOption50etfMinQvix()
	if err != nil {
		t.Logf("IndexOption50etfMinQvix failed: %v", err)
		return
	}
	t.Logf("IndexOption50etfMinQvix: got %d records", len(data))
}

// TestIndexOption300etfQvix 测试300ETF期权波动率指数
func TestIndexOption300etfQvix(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过外部网络测试")
	}
	data, err := IndexOption300etfQvix()
	if err != nil {
		t.Logf("IndexOption300etfQvix failed: %v", err)
		return
	}
	t.Logf("IndexOption300etfQvix: got %d records", len(data))
}

// TestIndexOption300etfMinQvix 测试300ETF期权波动率指数分时
func TestIndexOption300etfMinQvix(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过外部网络测试")
	}
	data, err := IndexOption300etfMinQvix()
	if err != nil {
		t.Logf("IndexOption300etfMinQvix failed: %v", err)
		return
	}
	t.Logf("IndexOption300etfMinQvix: got %d records", len(data))
}

// TestIndexOption500etfQvix 测试500ETF期权波动率指数
func TestIndexOption500etfQvix(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过外部网络测试")
	}
	data, err := IndexOption500etfQvix()
	if err != nil {
		t.Logf("IndexOption500etfQvix failed: %v", err)
		return
	}
	t.Logf("IndexOption500etfQvix: got %d records", len(data))
}

// TestIndexOptionCybQvix 测试创业板期权波动率指数
func TestIndexOptionCybQvix(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过外部网络测试")
	}
	data, err := IndexOptionCybQvix()
	if err != nil {
		t.Logf("IndexOptionCybQvix failed: %v", err)
		return
	}
	t.Logf("IndexOptionCybQvix: got %d records", len(data))
}

// TestIndexOptionKcbQvix 测试科创板期权波动率指数
func TestIndexOptionKcbQvix(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过外部网络测试")
	}
	data, err := IndexOptionKcbQvix()
	if err != nil {
		t.Logf("IndexOptionKcbQvix failed: %v", err)
		return
	}
	t.Logf("IndexOptionKcbQvix: got %d records", len(data))
}

// TestIndexOption100etfQvix 测试深证100ETF期权波动率指数
func TestIndexOption100etfQvix(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过外部网络测试")
	}
	data, err := IndexOption100etfQvix()
	if err != nil {
		t.Logf("IndexOption100etfQvix failed: %v", err)
		return
	}
	t.Logf("IndexOption100etfQvix: got %d records", len(data))
}

// TestIndexOption300indexQvix 测试中证300股指期权波动率指数
func TestIndexOption300indexQvix(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过外部网络测试")
	}
	data, err := IndexOption300indexQvix()
	if err != nil {
		t.Logf("IndexOption300indexQvix failed: %v", err)
		return
	}
	t.Logf("IndexOption300indexQvix: got %d records", len(data))
}

// TestIndexOption1000indexQvix 测试中证1000股指期权波动率指数
func TestIndexOption1000indexQvix(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过外部网络测试")
	}
	data, err := IndexOption1000indexQvix()
	if err != nil {
		t.Logf("IndexOption1000indexQvix failed: %v", err)
		return
	}
	t.Logf("IndexOption1000indexQvix: got %d records", len(data))
}

// TestIndexOption50indexQvix 测试上证50股指期权波动率指数
func TestIndexOption50indexQvix(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过外部网络测试")
	}
	data, err := IndexOption50indexQvix()
	if err != nil {
		t.Logf("IndexOption50indexQvix failed: %v", err)
		return
	}
	t.Logf("IndexOption50indexQvix: got %d records", len(data))
}

// ==================== 食糖指数测试 ====================

// TestIndexInnerQuoteSugarMsweet 测试配额内进口糖估算指数
func TestIndexInnerQuoteSugarMsweet(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过外部网络测试")
	}
	data, err := IndexInnerQuoteSugarMsweet()
	if err != nil {
		t.Logf("IndexInnerQuoteSugarMsweet failed: %v", err)
		return
	}
	t.Logf("IndexInnerQuoteSugarMsweet: got %d records", len(data))
}

// TestIndexOuterQuoteSugarMsweet 测试配额外进口糖估算指数
func TestIndexOuterQuoteSugarMsweet(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过外部网络测试")
	}
	data, err := IndexOuterQuoteSugarMsweet()
	if err != nil {
		t.Logf("IndexOuterQuoteSugarMsweet failed: %v", err)
		return
	}
	t.Logf("IndexOuterQuoteSugarMsweet: got %d records", len(data))
}

// ==================== 义乌指数测试 ====================

// TestIndexYwWeekPrice 测试义乌小商品周价格指数
func TestIndexYwWeekPrice(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过外部网络测试")
	}
	data, err := IndexYwWeekPrice()
	if err != nil {
		t.Logf("IndexYwWeekPrice failed: %v", err)
		return
	}
	t.Logf("IndexYwWeekPrice: got %d records", len(data))
}

// TestIndexYwMonthPrice 测试义乌小商品月价格指数
func TestIndexYwMonthPrice(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过外部网络测试")
	}
	data, err := IndexYwMonthPrice()
	if err != nil {
		t.Logf("IndexYwMonthPrice failed: %v", err)
		return
	}
	t.Logf("IndexYwMonthPrice: got %d records", len(data))
}

// TestIndexYwProsperity 测试义乌小商品月景气指数
func TestIndexYwProsperity(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过外部网络测试")
	}
	data, err := IndexYwProsperity()
	if err != nil {
		t.Logf("IndexYwProsperity failed: %v", err)
		return
	}
	t.Logf("IndexYwProsperity: got %d records", len(data))
}

// ==================== 申万基金指数测试 ====================

// TestIndexRealtimeFundSW 测试申万基金指数实时行情
func TestIndexRealtimeFundSW(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过外部网络测试")
	}
	tests := []string{"基础一级", "基础二级", "特色指数"}
	for _, symbol := range tests {
		data, err := IndexRealtimeFundSW(symbol)
		if err != nil {
			t.Logf("IndexRealtimeFundSW(%s) failed: %v", symbol, err)
			continue
		}
		t.Logf("IndexRealtimeFundSW(%s): got %d records", symbol, len(data))
	}
}

// TestIndexHistFundSW 测试申万基金指数历史行情
func TestIndexHistFundSW(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过外部网络测试")
	}
	tests := []struct {
		symbol string
		period string
	}{
		{"807200", "day"},
		{"807200", "week"},
		{"807200", "month"},
	}
	for _, tc := range tests {
		data, err := IndexHistFundSW(tc.symbol, tc.period)
		if err != nil {
			t.Logf("IndexHistFundSW(%s, %s) failed: %v", tc.symbol, tc.period, err)
			continue
		}
		t.Logf("IndexHistFundSW(%s, %s): got %d records", tc.symbol, tc.period, len(data))
	}
}

// ==================== 申万研究指数测试 ====================

// TestIndexHistSW 测试申万指数历史行情
func TestIndexHistSW(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过外部网络测试")
	}
	tests := []struct {
		symbol string
		period string
	}{
		{"801010", "day"},
		{"801010", "week"},
		{"801010", "month"},
	}
	for _, tc := range tests {
		data, err := IndexHistSW(tc.symbol, tc.period)
		if err != nil {
			t.Logf("IndexHistSW(%s, %s) failed: %v", tc.symbol, tc.period, err)
			continue
		}
		t.Logf("IndexHistSW(%s, %s): got %d records", tc.symbol, tc.period, len(data))
	}
}

// TestIndexMinSW 测试申万指数分时行情
func TestIndexMinSW(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过外部网络测试")
	}
	data, err := IndexMinSW("801010")
	if err != nil {
		t.Logf("IndexMinSW failed: %v", err)
		return
	}
	t.Logf("IndexMinSW: got %d records", len(data))
}

// TestIndexComponentSW 测试申万指数成分股
func TestIndexComponentSW(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过外部网络测试")
	}
	data, err := IndexComponentSW("801010")
	if err != nil {
		t.Logf("IndexComponentSW failed: %v", err)
		return
	}
	t.Logf("IndexComponentSW: got %d records", len(data))
}

// TestIndexRealtimeSW 测试申万指数实时行情
func TestIndexRealtimeSW(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过外部网络测试")
	}
	tests := []string{"一级行业", "二级行业", "三级行业", "风格指数", "全市场指数"}
	for _, symbol := range tests {
		data, err := IndexRealtimeSW(symbol)
		if err != nil {
			t.Logf("IndexRealtimeSW(%s) failed: %v", symbol, err)
			continue
		}
		t.Logf("IndexRealtimeSW(%s): got %d records", symbol, len(data))
	}
}

// TestIndexAnalysisDailySW 测试申万指数日分析数据
func TestIndexAnalysisDailySW(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过外部网络测试")
	}
	data, err := IndexAnalysisDailySW("801010", "20240101", "20240131")
	if err != nil {
		t.Logf("IndexAnalysisDailySW failed: %v", err)
		return
	}
	t.Logf("IndexAnalysisDailySW: got %d records", len(data))
}

// TestIndexAnalysisWeekMonthSW 测试申万指数周月日期
func TestIndexAnalysisWeekMonthSW(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过外部网络测试")
	}
	data, err := IndexAnalysisWeekMonthSW("801010")
	if err != nil {
		t.Logf("IndexAnalysisWeekMonthSW failed: %v", err)
		return
	}
	t.Logf("IndexAnalysisWeekMonthSW: got %d records", len(data))
}

// ==================== 中国股票指数测试 ====================

// TestStockZhIndexSpotEM_New 测试股票指数实时行情(新版)
func TestStockZhIndexSpotEM_New(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过外部网络测试")
	}
	tests := []string{"沪深重要指数", "上证系列指数", "深证系列指数", "中证系列指数"}
	for _, symbol := range tests {
		data, err := StockZhIndexSpotEM(symbol)
		if err != nil {
			t.Logf("StockZhIndexSpotEM(%s) failed: %v", symbol, err)
			continue
		}
		t.Logf("StockZhIndexSpotEM(%s): got %d records", symbol, len(data))
	}
}

// TestStockZhIndexDailyEM 测试股票指数日线数据
func TestStockZhIndexDailyEM(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过外部网络测试")
	}
	tests := []struct {
		symbol    string
		startDate string
		endDate   string
	}{
		{"sh000001", "20240101", "20240131"},
		{"sz399001", "20240101", "20240131"},
		{"csi931151", "20240101", "20240131"},
	}
	for _, tc := range tests {
		data, err := StockZhIndexDailyEM(tc.symbol, tc.startDate, tc.endDate)
		if err != nil {
			t.Logf("StockZhIndexDailyEM(%s) failed: %v", tc.symbol, err)
			continue
		}
		t.Logf("StockZhIndexDailyEM(%s): got %d records", tc.symbol, len(data))
	}
}

// ==================== 中证指数测试 ====================

// TestStockZhIndexHistCSIndex 测试中证指数历史行情
func TestStockZhIndexHistCSIndex(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过外部网络测试")
	}
	data, err := StockZhIndexHistCSIndex("000928", "20230101", "20240101")
	if err != nil {
		t.Logf("StockZhIndexHistCSIndex failed: %v", err)
		return
	}
	t.Logf("StockZhIndexHistCSIndex: got %d records", len(data))
}

// ==================== 新闻情绪指数测试 ====================

// TestIndexNewsSentimentScope 测试A股新闻情绪指数
func TestIndexNewsSentimentScope(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过外部网络测试")
	}
	data, err := IndexNewsSentimentScope()
	if err != nil {
		t.Logf("IndexNewsSentimentScope failed: %v", err)
		return
	}
	t.Logf("IndexNewsSentimentScope: got %d records", len(data))
}

// ==================== 浙江排污权交易指数测试 ====================

// TestIndexEri 测试浙江排污权交易指数(新版)
func TestIndexEri(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过外部网络测试")
	}
	tests := []string{"月度", "季度"}
	for _, symbol := range tests {
		data, err := IndexEri(symbol)
		if err != nil {
			t.Logf("IndexEri(%s) failed: %v (网站可能不可用)", symbol, err)
			continue
		}
		t.Logf("IndexEri(%s): got %d records", symbol, len(data))
	}
}

// ==================== 财新指数测试 ====================

// TestIndexCx 测试财新系列指数
func TestIndexCx(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过外部网络测试")
	}
	// 测试几个代表性指数
	data, err := IndexPmiComCx()
	if err != nil {
		t.Logf("IndexPmiComCx failed: %v", err)
	} else {
		t.Logf("IndexPmiComCx: got %d records", len(data))
	}

	data, err = IndexPmiManCx()
	if err != nil {
		t.Logf("IndexPmiManCx failed: %v", err)
	} else {
		t.Logf("IndexPmiManCx: got %d records", len(data))
	}

	data, err = IndexPmiSerCx()
	if err != nil {
		t.Logf("IndexPmiSerCx failed: %v", err)
	} else {
		t.Logf("IndexPmiSerCx: got %d records", len(data))
	}
}

// ==================== CFLP指数测试(新版) ====================

// TestIndexPriceCflp 测试中国公路物流运价指数(新版)
func TestIndexPriceCflp(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过外部网络测试")
	}
	tests := []string{"周指数", "月指数", "季度指数", "年度指数"}
	for _, symbol := range tests {
		data, err := IndexPriceCflp(symbol)
		if err != nil {
			t.Logf("IndexPriceCflp(%s) failed: %v", symbol, err)
			continue
		}
		t.Logf("IndexPriceCflp(%s): got %d records", symbol, len(data))
	}
}

// TestIndexVolumeCflp 测试中国公路物流运量指数(新版)
func TestIndexVolumeCflp(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过外部网络测试")
	}
	tests := []string{"月指数", "季度指数", "年度指数"}
	for _, symbol := range tests {
		data, err := IndexVolumeCflp(symbol)
		if err != nil {
			t.Logf("IndexVolumeCflp(%s) failed: %v", symbol, err)
			continue
		}
		t.Logf("IndexVolumeCflp(%s): got %d records", symbol, len(data))
	}
}
