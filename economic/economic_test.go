package economic

import (
	"testing"
)

// TestMacroBankUsaInterestRate 测试美联储利率决议报告
func TestMacroBankUsaInterestRate(t *testing.T) {
	data, err := MacroBankUsaInterestRate()
	if err != nil {
		t.Fatalf("获取美联储利率决议报告失败: %v", err)
	}

	if len(data) == 0 {
		t.Fatal("美联储利率决议报告数据为空")
	}

	t.Logf("美联储利率决议报告数据行数: %d", len(data))
	t.Logf("第一条数据: 日期=%v, 今值=%.2f%%, 预测值=%.2f%%, 前值=%.2f%%",
		data[0].Date, data[0].Current, data[0].Forecast, data[0].Previous)
}

// TestMacroBankEuroInterestRate 测试欧洲央行利率决议报告
func TestMacroBankEuroInterestRate(t *testing.T) {
	data, err := MacroBankEuroInterestRate()
	if err != nil {
		t.Fatalf("获取欧洲央行利率决议报告失败: %v", err)
	}

	if len(data) == 0 {
		t.Fatal("欧洲央行利率决议报告数据为空")
	}

	t.Logf("欧洲央行利率决议报告数据行数: %d", len(data))
	t.Logf("第一条数据: 日期=%v, 今值=%.2f%%", data[0].Date, data[0].Current)
}

// TestMacroBankNewzealandInterestRate 测试新西兰联储利率决议报告
func TestMacroBankNewzealandInterestRate(t *testing.T) {
	data, err := MacroBankNewzealandInterestRate()
	if err != nil {
		t.Fatalf("获取新西兰联储利率决议报告失败: %v", err)
	}

	if len(data) == 0 {
		t.Fatal("新西兰联储利率决议报告数据为空")
	}

	t.Logf("新西兰联储利率决议报告数据行数: %d", len(data))
	t.Logf("第一条数据: 日期=%v, 今值=%.2f%%", data[0].Date, data[0].Current)
}

// TestMacroBankChinaInterestRate 测试中国央行利率决议报告
func TestMacroBankChinaInterestRate(t *testing.T) {
	data, err := MacroBankChinaInterestRate()
	if err != nil {
		t.Fatalf("获取中国央行利率决议报告失败: %v", err)
	}

	if len(data) == 0 {
		t.Fatal("中国央行利率决议报告数据为空")
	}

	t.Logf("中国央行利率决议报告数据行数: %d", len(data))
	t.Logf("第一条数据: 日期=%v, 今值=%.2f%%", data[0].Date, data[0].Current)
}

// TestMacroBankSwitzerlandInterestRate 测试瑞士央行利率决议报告
func TestMacroBankSwitzerlandInterestRate(t *testing.T) {
	data, err := MacroBankSwitzerlandInterestRate()
	if err != nil {
		t.Fatalf("获取瑞士央行利率决议报告失败: %v", err)
	}

	if len(data) == 0 {
		t.Fatal("瑞士央行利率决议报告数据为空")
	}

	t.Logf("瑞士央行利率决议报告数据行数: %d", len(data))
	t.Logf("第一条数据: 日期=%v, 今值=%.2f%%", data[0].Date, data[0].Current)
}

// TestMacroBankEnglishInterestRate 测试英国央行利率决议报告
func TestMacroBankEnglishInterestRate(t *testing.T) {
	data, err := MacroBankEnglishInterestRate()
	if err != nil {
		t.Fatalf("获取英国央行利率决议报告失败: %v", err)
	}

	if len(data) == 0 {
		t.Fatal("英国央行利率决议报告数据为空")
	}

	t.Logf("英国央行利率决议报告数据行数: %d", len(data))
	t.Logf("第一条数据: 日期=%v, 今值=%.2f%%", data[0].Date, data[0].Current)
}

// TestMacroBankAustraliaInterestRate 测试澳洲联储利率决议报告
func TestMacroBankAustraliaInterestRate(t *testing.T) {
	data, err := MacroBankAustraliaInterestRate()
	if err != nil {
		t.Fatalf("获取澳洲联储利率决议报告失败: %v", err)
	}

	if len(data) == 0 {
		t.Fatal("澳洲联储利率决议报告数据为空")
	}

	t.Logf("澳洲联储利率决议报告数据行数: %d", len(data))
	t.Logf("第一条数据: 日期=%v, 今值=%.2f%%", data[0].Date, data[0].Current)
}

// TestMacroBankJapanInterestRate 测试日本央行利率决议报告
func TestMacroBankJapanInterestRate(t *testing.T) {
	data, err := MacroBankJapanInterestRate()
	if err != nil {
		t.Fatalf("获取日本央行利率决议报告失败: %v", err)
	}

	if len(data) == 0 {
		t.Fatal("日本央行利率决议报告数据为空")
	}

	t.Logf("日本央行利率决议报告数据行数: %d", len(data))
	t.Logf("第一条数据: 日期=%v, 今值=%.2f%%", data[0].Date, data[0].Current)
}

// TestMacroBankRussiaInterestRate 测试俄罗斯央行利率决议报告
func TestMacroBankRussiaInterestRate(t *testing.T) {
	data, err := MacroBankRussiaInterestRate()
	if err != nil {
		t.Fatalf("获取俄罗斯央行利率决议报告失败: %v", err)
	}

	if len(data) == 0 {
		t.Fatal("俄罗斯央行利率决议报告数据为空")
	}

	t.Logf("俄罗斯央行利率决议报告数据行数: %d", len(data))
	t.Logf("第一条数据: 日期=%v, 今值=%.2f%%", data[0].Date, data[0].Current)
}

// TestMacroBankIndiaInterestRate 测试印度央行利率决议报告
func TestMacroBankIndiaInterestRate(t *testing.T) {
	data, err := MacroBankIndiaInterestRate()
	if err != nil {
		t.Fatalf("获取印度央行利率决议报告失败: %v", err)
	}

	if len(data) == 0 {
		t.Fatal("印度央行利率决议报告数据为空")
	}

	t.Logf("印度央行利率决议报告数据行数: %d", len(data))
	t.Logf("第一条数据: 日期=%v, 今值=%.2f%%", data[0].Date, data[0].Current)
}

// TestMacroBankBrazilInterestRate 测试巴西央行利率决议报告
func TestMacroBankBrazilInterestRate(t *testing.T) {
	data, err := MacroBankBrazilInterestRate()
	if err != nil {
		t.Fatalf("获取巴西央行利率决议报告失败: %v", err)
	}

	if len(data) == 0 {
		t.Fatal("巴西央行利率决议报告数据为空")
	}

	t.Logf("巴西央行利率决议报告数据行数: %d", len(data))
	t.Logf("第一条数据: 日期=%v, 今值=%.2f%%", data[0].Date, data[0].Current)
}

// ===== 美国宏观数据测试 =====

// TestMacroUsaGdpMonthly 测试美国GDP报告
func TestMacroUsaGdpMonthly(t *testing.T) {
	data, err := MacroUsaGdpMonthly()
	if err != nil {
		t.Fatalf("获取美国GDP报告失败: %v", err)
	}

	if len(data) == 0 {
		t.Fatal("美国GDP报告数据为空")
	}

	t.Logf("美国GDP报告数据行数: %d", len(data))
	t.Logf("第一条数据: 商品=%s, 日期=%v, 今值=%.2f", data[0].Commodity, data[0].Date, data[0].Current)
}

// TestMacroUsaCpiMonthly 测试美国CPI月率报告
func TestMacroUsaCpiMonthly(t *testing.T) {
	data, err := MacroUsaCpiMonthly()
	if err != nil {
		t.Fatalf("获取美国CPI月率报告失败: %v", err)
	}

	if len(data) == 0 {
		t.Fatal("美国CPI月率报告数据为空")
	}

	t.Logf("美国CPI月率报告数据行数: %d", len(data))
	t.Logf("第一条数据: 商品=%s, 日期=%v, 今值=%.2f%%", data[0].Commodity, data[0].Date, data[0].Current)
}

// TestMacroUsaNonFarm 测试美国非农就业人数报告
func TestMacroUsaNonFarm(t *testing.T) {
	data, err := MacroUsaNonFarm()
	if err != nil {
		t.Fatalf("获取美国非农就业人数报告失败: %v", err)
	}

	if len(data) == 0 {
		t.Fatal("美国非农就业人数报告数据为空")
	}

	t.Logf("美国非农就业人数报告数据行数: %d", len(data))
	t.Logf("第一条数据: 商品=%s, 日期=%v, 今值=%.0f万人", data[0].Commodity, data[0].Date, data[0].Current)
}

// TestMacroUsaRigCount 测试贝克休斯钻井平台报告
func TestMacroUsaRigCount(t *testing.T) {
	data, err := MacroUsaRigCount()
	if err != nil {
		t.Fatalf("获取贝克休斯钻井平台报告失败: %v", err)
	}

	if len(data) == 0 {
		t.Fatal("贝克休斯钻井平台报告数据为空")
	}

	t.Logf("贝克休斯钻井平台报告数据行数: %d", len(data))
	t.Logf("第一条数据: 日期=%s, 钻井总数=%.0f, 变化=%.0f", data[0].Date, data[0].TotalRigCount, data[0].TotalRigChange)
}

// TestMacroUsaPhs 测试美国未决房屋销售月率（东方财富）
func TestMacroUsaPhs(t *testing.T) {
	data, err := MacroUsaPhs()
	if err != nil {
		t.Fatalf("获取美国未决房屋销售月率失败: %v", err)
	}

	if len(data) == 0 {
		t.Fatal("美国未决房屋销售月率数据为空")
	}

	t.Logf("美国未决房屋销售月率数据行数: %d", len(data))
	t.Logf("第一条数据: 时间=%s, 现值=%.2f, 前值=%.2f", data[0].Time, data[0].Current, data[0].Previous)
}

// TestMacroUsaCpiYoy 测试美国CPI年率（东方财富）
func TestMacroUsaCpiYoy(t *testing.T) {
	data, err := MacroUsaCpiYoy()
	if err != nil {
		t.Fatalf("获取美国CPI年率失败: %v", err)
	}

	if len(data) == 0 {
		t.Fatal("美国CPI年率数据为空")
	}

	t.Logf("美国CPI年率数据行数: %d", len(data))
	t.Logf("第一条数据: 时间=%s, 现值=%.2f%%", data[0].Time, data[0].Current)
}

// TestMacroUsaCrudeInner 测试美国原油产量报告
func TestMacroUsaCrudeInner(t *testing.T) {
	data, err := MacroUsaCrudeInner()
	if err != nil {
		t.Fatalf("获取美国原油产量报告失败: %v", err)
	}

	if len(data) == 0 {
		t.Fatal("美国原油产量报告数据为空")
	}

	t.Logf("美国原油产量报告数据行数: %d", len(data))
	t.Logf("第一条数据: 日期=%s, 总产量=%.0f, 变化=%.0f", data[0].Date, data[0].TotalProduction, data[0].TotalChange)
}

// TestMacroUsaCftcNcHolding 测试CFTC外汇类非商业持仓
func TestMacroUsaCftcNcHolding(t *testing.T) {
	data, err := MacroUsaCftcNcHolding()
	if err != nil {
		t.Fatalf("获取CFTC外汇类非商业持仓失败: %v", err)
	}

	if len(data) == 0 {
		t.Fatal("CFTC外汇类非商业持仓数据为空")
	}

	t.Logf("CFTC外汇类非商业持仓数据行数: %d", len(data))
	t.Logf("第一条数据: 日期=%s, 数据字段数=%d", data[0].Date, len(data[0].Data))
}

// TestMacroUsaCmeMerchantGoodsHolding 测试CME贵金属报告
func TestMacroUsaCmeMerchantGoodsHolding(t *testing.T) {
	data, err := MacroUsaCmeMerchantGoodsHolding()
	if err != nil {
		t.Fatalf("获取CME贵金属报告失败: %v", err)
	}

	if len(data) == 0 {
		t.Fatal("CME贵金属报告数据为空")
	}

	t.Logf("CME贵金属报告数据行数: %d", len(data))
	t.Logf("第一条数据: 日期=%s, 品种=%s, 成交量=%.0f", data[0].Date, data[0].Variety, data[0].Volume)
}
