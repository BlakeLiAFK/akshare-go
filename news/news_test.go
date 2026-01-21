package news

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestStockNewsEm 测试东方财富个股新闻接口
func TestStockNewsEm(t *testing.T) {
	// 使用默认股票代码测试
	result, err := StockNewsEm("603777")
	assert.NoError(t, err, "StockNewsEm 不应返回错误")
	assert.NotNil(t, result, "结果不应为空")

	if len(result) > 0 {
		// 验证第一条新闻的字段
		news := result[0]
		assert.Equal(t, "603777", news.Keyword, "关键词应为股票代码")
		assert.NotEmpty(t, news.Title, "标题不应为空")
		assert.NotEmpty(t, news.Url, "链接不应为空")
		t.Logf("获取到 %d 条新闻", len(result))
		t.Logf("第一条新闻: %s", news.Title)
	}
}

// TestStockNewsEmDifferentSymbol 测试不同股票代码
func TestStockNewsEmDifferentSymbol(t *testing.T) {
	// 测试茅台股票
	result, err := StockNewsEm("600519")
	assert.NoError(t, err, "StockNewsEm 不应返回错误")
	assert.NotNil(t, result, "结果不应为空")

	if len(result) > 0 {
		news := result[0]
		assert.Equal(t, "600519", news.Keyword, "关键词应为股票代码")
		t.Logf("茅台新闻数量: %d", len(result))
	}
}

// TestNewsCctv 测试新闻联播文字稿接口
func TestNewsCctv(t *testing.T) {
	// 使用较近的日期测试
	result, err := NewsCctv("20260101")
	assert.NoError(t, err, "NewsCctv 不应返回错误")
	assert.NotNil(t, result, "结果不应为空")

	if len(result) > 0 {
		news := result[0]
		assert.Equal(t, "20260101", news.Date, "日期应匹配")
		assert.NotEmpty(t, news.Title, "标题不应为空")
		t.Logf("获取到 %d 条新闻联播", len(result))
		t.Logf("第一条: %s", news.Title)
	}
}

// TestNewsEconomicBaidu 测试百度股市通经济数据接口
func TestNewsEconomicBaidu(t *testing.T) {
	// 注意：百度接口可能需要 Cookie，这里测试基本功能
	result, err := NewsEconomicBaidu("20260101", "")
	if err != nil {
		t.Logf("百度经济数据接口可能需要Cookie: %v", err)
		t.Skip("跳过需要Cookie的测试")
	}

	assert.NotNil(t, result, "结果不应为空")
	if len(result) > 0 {
		data := result[0]
		t.Logf("获取到 %d 条经济数据", len(result))
		t.Logf("第一条: %s - %s", data.Country, data.Event)
	}
}

// TestNewsTradeNotifySuspendBaidu 测试百度股市通停复牌接口
func TestNewsTradeNotifySuspendBaidu(t *testing.T) {
	result, err := NewsTradeNotifySuspendBaidu("20260101", "")
	if err != nil {
		t.Logf("百度停复牌接口可能需要Cookie: %v", err)
		t.Skip("跳过需要Cookie的测试")
	}

	assert.NotNil(t, result, "结果不应为空")
	if len(result) > 0 {
		data := result[0]
		t.Logf("获取到 %d 条停复牌数据", len(result))
		t.Logf("第一条: %s - %s", data.Code, data.Name)
	}
}

// TestNewsTradeNotifyDividendBaidu 测试百度股市通分红派息接口
func TestNewsTradeNotifyDividendBaidu(t *testing.T) {
	result, err := NewsTradeNotifyDividendBaidu("20260101", "")
	if err != nil {
		t.Logf("百度分红派息接口可能需要Cookie: %v", err)
		t.Skip("跳过需要Cookie的测试")
	}

	assert.NotNil(t, result, "结果不应为空")
	if len(result) > 0 {
		data := result[0]
		t.Logf("获取到 %d 条分红派息数据", len(result))
		t.Logf("第一条: %s - %s", data.Code, data.Name)
	}
}

// TestNewsReportTimeBaidu 测试百度股市通财报发行接口
func TestNewsReportTimeBaidu(t *testing.T) {
	result, err := NewsReportTimeBaidu("20260101", "")
	if err != nil {
		t.Logf("百度财报发行接口可能需要Cookie: %v", err)
		t.Skip("跳过需要Cookie的测试")
	}

	assert.NotNil(t, result, "结果不应为空")
	if len(result) > 0 {
		data := result[0]
		t.Logf("获取到 %d 条财报发行数据", len(result))
		t.Logf("第一条: %s - %s", data.Code, data.Name)
	}
}

// TestNewsCctvDateFormat 测试不同日期格式
func TestNewsCctvDateFormat(t *testing.T) {
	// 测试 2016 年之后的日期
	result, err := NewsCctv("20260101")
	assert.NoError(t, err, "NewsCctv 不应返回错误")
	if len(result) > 0 {
		t.Logf("2026年1月1日新闻数量: %d", len(result))
	}
}
