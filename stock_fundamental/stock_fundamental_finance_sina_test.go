package stock_fundamental

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestStockFinancialReportSina 测试新浪财务报表接口
func TestStockFinancialReportSina(t *testing.T) {
	// 测试资产负债表
	df, err := StockFinancialReportSina("sh600600", "资产负债表")
	assert.NoError(t, err)
	assert.True(t, df.Nrow() > 0, "应该返回数据")
	t.Logf("资产负债表行数: %d, 列数: %d", df.Nrow(), df.Ncol())

	// 测试利润表
	df2, err := StockFinancialReportSina("sh600600", "利润表")
	assert.NoError(t, err)
	assert.True(t, df2.Nrow() > 0, "应该返回数据")
	t.Logf("利润表行数: %d, 列数: %d", df2.Nrow(), df2.Ncol())

	// 测试现金流量表
	df3, err := StockFinancialReportSina("sh600600", "现金流量表")
	assert.NoError(t, err)
	assert.True(t, df3.Nrow() > 0, "应该返回数据")
	t.Logf("现金流量表行数: %d, 列数: %d", df3.Nrow(), df3.Ncol())

	// 测试错误的报表类型
	_, err = StockFinancialReportSina("sh600600", "错误的报表")
	assert.Error(t, err)
}

// TestStockFinancialAbstract 测试新浪财务摘要接口
func TestStockFinancialAbstract(t *testing.T) {
	df, err := StockFinancialAbstract("600004")
	assert.NoError(t, err)
	assert.True(t, df.Nrow() > 0, "应该返回数据")
	t.Logf("财务摘要行数: %d, 列数: %d", df.Nrow(), df.Ncol())

	// 打印前几行数据
	if df.Nrow() > 0 {
		t.Logf("列名: %v", df.Names())
	}
}

// TestStockFinancialAnalysisIndicatorEm 测试东方财富财务分析指标接口
func TestStockFinancialAnalysisIndicatorEm(t *testing.T) {
	// 测试按报告期
	indicators, err := StockFinancialAnalysisIndicatorEm("301389.SZ", "按报告期")
	assert.NoError(t, err)
	assert.True(t, len(indicators) > 0, "应该返回数据")
	t.Logf("按报告期指标数量: %d", len(indicators))

	if len(indicators) > 0 {
		t.Logf("第一个指标: 报告期=%s, ROE=%.2f", indicators[0].ReportDate, indicators[0].ROE)
	}

	// 测试按单季度
	indicators2, err := StockFinancialAnalysisIndicatorEm("301389.SZ", "按单季度")
	assert.NoError(t, err)
	assert.True(t, len(indicators2) > 0, "应该返回数据")
	t.Logf("按单季度指标数量: %d", len(indicators2))
}

// TestStockHistoryDividend 测试历史分红接口
func TestStockHistoryDividend(t *testing.T) {
	dividends, err := StockHistoryDividend()
	assert.NoError(t, err)
	assert.True(t, len(dividends) > 0, "应该返回数据")
	t.Logf("历史分红记录数: %d", len(dividends))

	if len(dividends) > 0 {
		t.Logf("第一条记录: 代码=%s, 名称=%s, 累计股息=%.2f, 分红次数=%d",
			dividends[0].Code, dividends[0].Name, dividends[0].TotalDividend, dividends[0].DividendCount)
	}
}

// TestStockHistoryDividendDetail 测试分红详情接口
func TestStockHistoryDividendDetail(t *testing.T) {
	// 测试分红详情
	details, err := StockHistoryDividendDetail("000002", "分红", "")
	assert.NoError(t, err)
	t.Logf("分红详情记录数: %d", len(details))

	if len(details) > 0 {
		t.Logf("第一条记录: 公告日期=%s, 派息=%.2f, 送股=%.2f",
			details[0].AnnounceDate, details[0].Dividend, details[0].BonusShare)
	}
}

// TestStockIpoInfo 测试新股发行信息接口
func TestStockIpoInfo(t *testing.T) {
	info, err := StockIpoInfo("600004")
	assert.NoError(t, err)
	assert.True(t, len(info) > 0, "应该返回数据")
	t.Logf("新股发行信息项数: %d", len(info))

	if len(info) > 0 {
		t.Logf("第一条信息: %s = %s", info[0].Item, info[0].Value)
	}
}

// TestStockMainStockHolder 测试主要股东接口
func TestStockMainStockHolder(t *testing.T) {
	holders, err := StockMainStockHolder("600004")
	assert.NoError(t, err)
	t.Logf("主要股东记录数: %d", len(holders))

	if len(holders) > 0 {
		t.Logf("第一条记录: 股东=%s, 持股=%.2f, 比例=%.2f%%",
			holders[0].HolderName, holders[0].Holdings, holders[0].HoldingRatio)
	}
}

// TestStockFundStockHolder 测试基金持股接口
func TestStockFundStockHolder(t *testing.T) {
	funds, err := StockFundStockHolder("601318")
	assert.NoError(t, err)
	t.Logf("基金持股记录数: %d", len(funds))

	if len(funds) > 0 {
		t.Logf("第一条记录: 基金=%s(%s), 持仓=%.2f, 比例=%.2f%%",
			funds[0].FundName, funds[0].FundCode, funds[0].Holdings, funds[0].HoldingRatio)
	}
}

// TestStockFinancialReportSinaInvalidParams 测试无效参数
func TestStockFinancialReportSinaInvalidParams(t *testing.T) {
	// 测试空股票代码
	_, err := StockFinancialReportSina("", "资产负债表")
	assert.Error(t, err)
}

// TestStockFinancialAbstractInvalidParams 测试无效参数
func TestStockFinancialAbstractInvalidParams(t *testing.T) {
	// 测试空股票代码
	_, err := StockFinancialAbstract("")
	assert.Error(t, err)
}

// TestStockFinancialAnalysisIndicatorEmInvalidParams 测试无效参数
func TestStockFinancialAnalysisIndicatorEmInvalidParams(t *testing.T) {
	// 测试空股票代码
	_, err := StockFinancialAnalysisIndicatorEm("", "按报告期")
	assert.Error(t, err)
}

// TestStockHistoryDividendDetailInvalidParams 测试无效参数
func TestStockHistoryDividendDetailInvalidParams(t *testing.T) {
	// 测试空股票代码
	_, err := StockHistoryDividendDetail("", "分红", "")
	assert.Error(t, err)
}

// TestStockIpoInfoInvalidParams 测试无效参数
func TestStockIpoInfoInvalidParams(t *testing.T) {
	// 测试空股票代码
	_, err := StockIpoInfo("")
	assert.Error(t, err)
}

// TestStockMainStockHolderInvalidParams 测试无效参数
func TestStockMainStockHolderInvalidParams(t *testing.T) {
	// 测试空股票代码
	_, err := StockMainStockHolder("")
	assert.Error(t, err)
}

// TestStockFundStockHolderInvalidParams 测试无效参数
func TestStockFundStockHolderInvalidParams(t *testing.T) {
	// 测试空股票代码
	_, err := StockFundStockHolder("")
	assert.Error(t, err)
}
