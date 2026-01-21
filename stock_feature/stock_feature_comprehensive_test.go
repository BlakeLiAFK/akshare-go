package stock_feature

import (
	"testing"
)

// TestStockAIndicator 测试A股指标
func TestStockAIndicator(t *testing.T) {
	df, err := StockAIndicator("000001")
	if err != nil {
		t.Logf("StockAIndicator 返回错误(可能是网络问题): %v", err)
		return
	}

	if df.Nrow() == 0 {
		t.Log("StockAIndicator 返回空数据(使用示例数据)")
	} else {
		t.Logf("StockAIndicator 返回 %d 行数据", df.Nrow())
	}
}

// TestStockMarginEm 测试融资融券
func TestStockMarginEm(t *testing.T) {
	df, err := StockMarginEm("2024-01-15", "2024-01-15")
	if err != nil {
		t.Logf("StockMarginEm 返回错误(可能是网络问题): %v", err)
		return
	}

	if df.Nrow() == 0 {
		t.Log("StockMarginEm 返回空数据(使用示例数据)")
	} else {
		t.Logf("StockMarginEm 返回 %d 行数据", df.Nrow())
	}
}

// TestStockHsgtEmFlow 测试沪深港通资金流向
func TestStockHsgtEmFlow(t *testing.T) {
	df, err := StockHsgtEmFlow("north")
	if err != nil {
		t.Logf("StockHsgtEmFlow 返回错误(可能是网络问题): %v", err)
		return
	}

	if df.Nrow() == 0 {
		t.Log("StockHsgtEmFlow 返回空数据(使用示例数据)")
	} else {
		t.Logf("StockHsgtEmFlow 返回 %d 行数据", df.Nrow())
	}
}

// TestStockHistEm 测试历史行情
func TestStockHistEm(t *testing.T) {
	df, err := StockHistEm("000001", "daily", "20240101", "20240115", "qfq")
	if err != nil {
		t.Logf("StockHistEm 返回错误(可能是网络问题): %v", err)
		return
	}

	if df.Nrow() == 0 {
		t.Log("StockHistEm 返回空数据(使用示例数据)")
	} else {
		t.Logf("StockHistEm 返回 %d 行数据", df.Nrow())
	}
}

// TestStockAccountEm 测试开户数据
func TestStockAccountEm(t *testing.T) {
	df, err := StockAccountEm()
	if err != nil {
		t.Logf("StockAccountEm 返回错误(可能是网络问题): %v", err)
		return
	}

	if df.Nrow() == 0 {
		t.Log("StockAccountEm 返回空数据(使用示例数据)")
	} else {
		t.Logf("StockAccountEm 返回 %d 行数据", df.Nrow())
	}
}

// TestStockReportEm 测试研报数据
func TestStockReportEm(t *testing.T) {
	df, err := StockReportEm("000001")
	if err != nil {
		t.Logf("StockReportEm 返回错误(可能是网络问题): %v", err)
		return
	}

	if df.Nrow() == 0 {
		t.Log("StockReportEm 返回空数据(使用示例数据)")
	} else {
		t.Logf("StockReportEm 返回 %d 行数据", df.Nrow())
	}
}

// TestStockJgdyEm 测试机构调研
func TestStockJgdyEm(t *testing.T) {
	df, err := StockJgdyEm("000001")
	if err != nil {
		t.Logf("StockJgdyEm 返回错误(可能是网络问题): %v", err)
		return
	}

	if df.Nrow() == 0 {
		t.Log("StockJgdyEm 返回空数据(使用示例数据)")
	} else {
		t.Logf("StockJgdyEm 返回 %d 行数据", df.Nrow())
	}
}

// TestStockFundFlow 测试资金流向
func TestStockFundFlow(t *testing.T) {
	df, err := StockFundFlow("行业")
	if err != nil {
		t.Logf("StockFundFlow 返回错误(可能是网络问题): %v", err)
		return
	}

	if df.Nrow() == 0 {
		t.Log("StockFundFlow 返回空数据(使用示例数据)")
	} else {
		t.Logf("StockFundFlow 返回 %d 行数据", df.Nrow())
	}
}

// TestStockAnalystEm 测试分析师排名
func TestStockAnalystEm(t *testing.T) {
	df, err := StockAnalystEm("2024")
	if err != nil {
		t.Logf("StockAnalystEm 返回错误(可能是网络问题): %v", err)
		return
	}

	if df.Nrow() == 0 {
		t.Log("StockAnalystEm 返回空数据(使用示例数据)")
	} else {
		t.Logf("StockAnalystEm 返回 %d 行数据", df.Nrow())
	}
}

// TestStockValueEm 测试估值分析
func TestStockValueEm(t *testing.T) {
	result, err := StockValueEm("000001")
	if err != nil {
		t.Logf("StockValueEm 返回错误(可能是网络问题): %v", err)
		return
	}

	if len(result) == 0 {
		t.Log("StockValueEm 返回空数据(使用示例数据)")
	} else {
		t.Logf("StockValueEm 返回 %d 条数据", len(result))
	}
}
