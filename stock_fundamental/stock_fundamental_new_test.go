package stock_fundamental

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestStockFinancialHkAnalysisIndicatorEm 测试港股财务分析主要指标
func TestStockFinancialHkAnalysisIndicatorEm(t *testing.T) {
	result, err := StockFinancialHkAnalysisIndicatorEm("00700", "年度")
	assert.NoError(t, err)
	t.Logf("港股主要指标(年度): %d 条记录", len(result))
}

// TestStockFinancialHkReportEm 测试港股财务报表
func TestStockFinancialHkReportEm(t *testing.T) {
	result, err := StockFinancialHkReportEm("00700", "资产负债表", "年度")
	assert.NoError(t, err)
	t.Logf("港股资产负债表: %d 条记录", len(result))
}

// TestStockFinancialAbstractNewThs 测试同花顺重要指标
func TestStockFinancialAbstractNewThs(t *testing.T) {
	result, err := StockFinancialAbstractNewThs("000063", "按报告期")
	assert.NoError(t, err)
	t.Logf("同花顺重要指标: %d 条记录", len(result))
}

// TestStockFinancialDebtNewThs 测试同花顺资产负债表
func TestStockFinancialDebtNewThs(t *testing.T) {
	result, err := StockFinancialDebtNewThs("000063", "按报告期")
	assert.NoError(t, err)
	t.Logf("同花顺资产负债表: %d 条记录", len(result))
}

// TestStockFinancialBenefitNewThs 测试同花顺利润表
func TestStockFinancialBenefitNewThs(t *testing.T) {
	result, err := StockFinancialBenefitNewThs("000063", "按报告期")
	assert.NoError(t, err)
	t.Logf("同花顺利润表: %d 条记录", len(result))
}

// TestStockFinancialCashNewThs 测试同花顺现金流量表
func TestStockFinancialCashNewThs(t *testing.T) {
	result, err := StockFinancialCashNewThs("000063", "按报告期")
	assert.NoError(t, err)
	t.Logf("同花顺现金流量表: %d 条记录", len(result))
}

// TestStockFinancialUsAnalysisIndicatorEm 测试美股财务分析主要指标
func TestStockFinancialUsAnalysisIndicatorEm(t *testing.T) {
	result, err := StockFinancialUsAnalysisIndicatorEm("TSLA", "年报")
	assert.NoError(t, err)
	t.Logf("美股主要指标: %d 条记录", len(result))
}

// TestStockFinancialUsReportEm 测试美股财务报表
func TestStockFinancialUsReportEm(t *testing.T) {
	result, err := StockFinancialUsReportEm("TSLA", "资产负债表", "年报")
	assert.NoError(t, err)
	t.Logf("美股资产负债表: %d 条记录", len(result))
}

// TestStockNoticeReport 测试股票公告
func TestStockNoticeReport(t *testing.T) {
	result, err := StockNoticeReport("全部", "20240101")
	assert.NoError(t, err)
	t.Logf("股票公告: %d 条记录", len(result))
}

// TestStockInstituteRecommend 测试机构推荐
func TestStockInstituteRecommend(t *testing.T) {
	result, err := StockInstituteRecommend("最新投资评级")
	assert.NoError(t, err)
	t.Logf("机构推荐: %d 条记录", len(result))
}

// TestStockInstituteRecommendDetail 测试股票评级记录
func TestStockInstituteRecommendDetail(t *testing.T) {
	result, err := StockInstituteRecommendDetail("000001")
	assert.NoError(t, err)
	t.Logf("股票评级记录: %d 条记录", len(result))
}
