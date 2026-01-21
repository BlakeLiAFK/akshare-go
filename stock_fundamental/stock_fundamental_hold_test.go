package stock_fundamental

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestStockInstituteHold 测试新浪财经-机构持股一览表接口
func TestStockInstituteHold(t *testing.T) {
	// 测试2020年一季报
	holds, err := StockInstituteHold("20201")
	assert.NoError(t, err)
	assert.True(t, len(holds) > 0, "应该返回数据")
	t.Logf("机构持股一览表记录数: %d", len(holds))

	if len(holds) > 0 {
		t.Logf("第一条记录: 代码=%s, 名称=%s, 机构数=%.0f, 持股比例=%.2f%%",
			holds[0].Code, holds[0].Name, holds[0].InstituteCount, holds[0].HoldingRatio)
	}

	// 测试2020年中报
	holds2, err := StockInstituteHold("20202")
	assert.NoError(t, err)
	t.Logf("2020年中报机构持股一览表记录数: %d", len(holds2))
}

// TestStockInstituteHoldDetail 测试新浪财经-机构持股详情接口
func TestStockInstituteHoldDetail(t *testing.T) {
	// 测试特定股票的机构持股详情（可能有数据）
	details, err := StockInstituteHoldDetail("600433", "20201")
	// 某些股票可能没有机构持股数据，这是正常的
	if err == nil {
		t.Logf("600433机构持股详情记录数: %d", len(details))
		if len(details) > 0 {
			t.Logf("第一条记录: 类型=%s, 名称=%s, 持股=%.2f万股, 比例=%.2f%%",
				details[0].InstituteType, details[0].InstituteName, details[0].Holdings, details[0].HoldingRatio)
		}
	}

	// 测试另一只股票（通常有数据）
	details2, err := StockInstituteHoldDetail("300003", "20201")
	assert.NoError(t, err)
	assert.True(t, len(details2) > 0, "应该返回数据")
	t.Logf("300003机构持股详情记录数: %d", len(details2))

	if len(details2) > 0 {
		t.Logf("第一条记录: 类型=%s, 名称=%s, 持股=%.2f万股, 比例=%.2f%%",
			details2[0].InstituteType, details2[0].InstituteName, details2[0].Holdings, details2[0].HoldingRatio)
	}
}

// TestStockZhAGbjgEm 测试东方财富-股本结构接口
func TestStockZhAGbjgEm(t *testing.T) {
	// 测试特定股票的股本结构
	gbjg, err := StockZhAGbjgEm("603392.SH")
	assert.NoError(t, err)
	assert.True(t, len(gbjg) > 0, "应该返回数据")
	t.Logf("股本结构记录数: %d", len(gbjg))

	if len(gbjg) > 0 {
		t.Logf("第一条记录: 日期=%s, 总股本=%.2f万股, 流通A股=%.2f万股, 原因=%s",
			gbjg[0].ChangeDate, gbjg[0].TotalShares, gbjg[0].ListedAShares, gbjg[0].ChangeReason)
	}

	// 测试另一只股票
	gbjg2, err := StockZhAGbjgEm("600000.SH")
	assert.NoError(t, err)
	t.Logf("600000.SH股本结构记录数: %d", len(gbjg2))
}

// TestStockZygcEm 测试东方财富-主营构成接口
func TestStockZygcEm(t *testing.T) {
	// 测试特定股票的主营构成
	zygc, err := StockZygcEm("SH688041")
	assert.NoError(t, err)
	assert.True(t, len(zygc) > 0, "应该返回数据")
	t.Logf("主营构成记录数: %d", len(zygc))

	if len(zygc) > 0 {
		t.Logf("第一条记录: 代码=%s, 日期=%s, 分类=%s, 构成=%s, 收入=%.2f万元, 比例=%.2f%%",
			zygc[0].Code, zygc[0].ReportDate, zygc[0].CategoryType,
			zygc[0].MainComposition, zygc[0].MainIncome, zygc[0].IncomeRatio)
	}

	// 统计不同分类类型
	categoryTypes := make(map[string]int)
	for _, item := range zygc {
		categoryTypes[item.CategoryType]++
	}
	t.Logf("分类类型统计: %v", categoryTypes)
}

// TestStockInstituteHoldInvalidParams 测试无效参数
func TestStockInstituteHoldInvalidParams(t *testing.T) {
	// 测试空参数
	_, err := StockInstituteHold("")
	assert.Error(t, err)

	// 测试格式错误
	_, err = StockInstituteHold("2020")
	assert.Error(t, err)
}

// TestStockInstituteHoldDetailInvalidParams 测试无效参数
func TestStockInstituteHoldDetailInvalidParams(t *testing.T) {
	// 测试空股票代码
	_, err := StockInstituteHoldDetail("", "20201")
	assert.Error(t, err)

	// 测试格式错误的报告期
	_, err = StockInstituteHoldDetail("600433", "2020")
	assert.Error(t, err)
}

// TestStockZhAGbjgEmInvalidParams 测试无效参数
func TestStockZhAGbjgEmInvalidParams(t *testing.T) {
	// 测试空股票代码
	_, err := StockZhAGbjgEm("")
	assert.Error(t, err)
}

// TestStockZygcEmInvalidParams 测试无效参数
func TestStockZygcEmInvalidParams(t *testing.T) {
	// 测试空股票代码
	_, err := StockZygcEm("")
	assert.Error(t, err)
}

// TestTranslateInstituteType 测试机构类型翻译
func TestTranslateInstituteType(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"fund", "基金"},
		{"socialSecurity", "全国社保"},
		{"qfii", "QFII"},
		{"insurance", "保险"},
		{"unknown", "unknown"},
	}

	for _, tt := range tests {
		result := translateInstituteType(tt.input)
		assert.Equal(t, tt.expected, result)
	}
}

// TestTranslateCategoryType 测试分类类型翻译
func TestTranslateCategoryType(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"2", "按产品分类"},
		{"3", "按地区分类"},
		{"1", "1"},
		{"unknown", "unknown"},
	}

	for _, tt := range tests {
		result := translateCategoryType(tt.input)
		assert.Equal(t, tt.expected, result)
	}
}

// TestParseJSONP 测试JSONP解析
func TestParseJSONP(t *testing.T) {
	// 测试正常的JSONP响应
	jsonp := `var details={"data":{"fund_001":[["001","基金A",100,200]]}};`
	jsonStr, err := parseJSONP(jsonp)
	assert.NoError(t, err)
	assert.Contains(t, jsonStr, "{")
	assert.Contains(t, jsonStr, "}")

	// 测试无效的JSONP响应
	invalid := "invalid jsonp"
	_, err = parseJSONP(invalid)
	assert.Error(t, err)
}
