package stock_feature

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestStockAHighLowStatistics 测试创新高新低统计
func TestStockAHighLowStatistics(t *testing.T) {
	result, err := StockAHighLowStatistics("all")
	assert.NoError(t, err)
	t.Logf("创新高新低统计: %d 条记录", result.Nrow()-1) // -1 因为包含表头
}

// TestStockAccountStatisticsEm 测试股票账户统计
func TestStockAccountStatisticsEm(t *testing.T) {
	result, err := StockAccountStatisticsEm()
	assert.NoError(t, err)
	t.Logf("股票账户统计: %d 条记录", len(result))
}

// TestStockZtPoolEm 测试涨停股池
func TestStockZtPoolEm(t *testing.T) {
	result, err := StockZtPoolEm("20240415")
	assert.NoError(t, err)
	t.Logf("涨停股池: %d 条记录", len(result))
}

// TestStockZtPoolPreviousEm 测试昨日涨停股池
func TestStockZtPoolPreviousEm(t *testing.T) {
	result, err := StockZtPoolPreviousEm("20240415")
	assert.NoError(t, err)
	t.Logf("昨日涨停股池: %d 条记录", len(result))
}

// TestStockDtPoolEm 测试跌停股池
func TestStockDtPoolEm(t *testing.T) {
	result, err := StockDtPoolEm("20240415")
	assert.NoError(t, err)
	t.Logf("跌停股池: %d 条记录", len(result))
}

// TestStockFhpsEm 测试分红送配
func TestStockFhpsEm(t *testing.T) {
	result, err := StockFhpsEm("20231231")
	assert.NoError(t, err)
	t.Logf("分红送配: %d 条记录", result.Nrow())
}

// TestStockFhpsDetailEm 测试分红送配详情
func TestStockFhpsDetailEm(t *testing.T) {
	result, err := StockFhpsDetailEm("300073")
	assert.NoError(t, err)
	t.Logf("分红送配详情: %d 条记录", result.Nrow())
}

// TestStockHkIndicatorEniu 测试港股指标
func TestStockHkIndicatorEniu(t *testing.T) {
	result, err := StockHkIndicatorEniu("hk01093", "市盈率")
	assert.NoError(t, err)
	t.Logf("港股指标: %d 条记录", len(result))
}

// TestStockZhASpotEm 测试沪深京A股实时行情
func TestStockZhASpotEm(t *testing.T) {
	result, err := StockZhASpotEm()
	assert.NoError(t, err)
	t.Logf("沪深京A股实时行情: %d 条记录", len(result))
	if len(result) > 0 {
		t.Logf("示例: %s %s 最新价:%.2f", result[0].Code, result[0].Name, result[0].LatestPrice)
	}
}

// TestStockZhAGdhs 测试股东户数
func TestStockZhAGdhs(t *testing.T) {
	result, err := StockZhAGdhs("最新")
	assert.NoError(t, err)
	t.Logf("股东户数: %d 条记录", result.Nrow()-1) // -1 因为包含表头
}

// TestStockTfpEm 测试停复牌信息
func TestStockTfpEm(t *testing.T) {
	result, err := StockTfpEm("20240426")
	assert.NoError(t, err)
	t.Logf("停复牌信息: %d 条记录", len(result))
}

// TestStockInfoGlobalEm 测试全球财经快讯
func TestStockInfoGlobalEm(t *testing.T) {
	result, err := StockInfoGlobalEm()
	assert.NoError(t, err)
	t.Logf("全球财经快讯: %d 条记录", result.Nrow()-1) // -1 因为包含表头
}

// TestStockLhbDetailEm 测试龙虎榜详情
func TestStockLhbDetailEm(t *testing.T) {
	result, err := StockLhbDetailEm("20240401", "20240410")
	assert.NoError(t, err)
	t.Logf("龙虎榜详情: %d 条记录", len(result))
}

// TestStockHsgtFundFlowSummaryEm 测试沪深港通资金流向
func TestStockHsgtFundFlowSummaryEm(t *testing.T) {
	result, err := StockHsgtFundFlowSummaryEm()
	assert.NoError(t, err)
	t.Logf("沪深港通资金流向: %d 条记录", len(result))
}

// TestStockGpzyProfileEm 测试股权质押市场概况
func TestStockGpzyProfileEm(t *testing.T) {
	result, err := StockGpzyProfileEm()
	assert.NoError(t, err)
	t.Logf("股权质押市场概况: %d 条记录", len(result))
}

// TestStockDxsylEm 测试打新收益率
func TestStockDxsylEm(t *testing.T) {
	result, err := StockDxsylEm("20240426")
	assert.NoError(t, err)
	t.Logf("打新收益率: %d 条记录", result.Nrow()-1) // -1 因为包含表头
}

// TestStockCommentEm 测试千股千评
func TestStockCommentEm(t *testing.T) {
	result, err := StockCommentEm()
	assert.NoError(t, err)
	t.Logf("千股千评: %d 条记录", result.Nrow()-1) // -1 因为包含表头
}
