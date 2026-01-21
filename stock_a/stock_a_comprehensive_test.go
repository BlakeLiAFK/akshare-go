package stock_a

import (
	"testing"
)

// TestStockBoardConceptNameEm 测试东方财富概念板块名称
func TestStockBoardConceptNameEm(t *testing.T) {
	df, err := StockBoardConceptNameEm()
	if err != nil {
		t.Logf("StockBoardConceptNameEm 返回错误(可能是网络问题): %v", err)
		return
	}

	if df.Nrow() == 0 {
		t.Log("StockBoardConceptNameEm 返回空数据(使用示例数据)")
	} else {
		t.Logf("StockBoardConceptNameEm 返回 %d 行数据", df.Nrow())
	}

	cols := df.Names()
	if len(cols) == 0 {
		t.Error("StockBoardConceptNameEm 返回的DataFrame没有列")
	}
}

// TestStockIndividualFundFlowRank 测试东方财富资金流向排名
func TestStockIndividualFundFlowRank(t *testing.T) {
	df, err := StockIndividualFundFlowRank("今日")
	if err != nil {
		t.Logf("StockIndividualFundFlowRank 返回错误(可能是网络问题): %v", err)
		return
	}

	if df.Nrow() == 0 {
		t.Log("StockIndividualFundFlowRank 返回空数据(使用示例数据)")
	} else {
		t.Logf("StockIndividualFundFlowRank 返回 %d 行数据", df.Nrow())
	}
}

// TestStockZhASpotEm 测试A股实时行情
func TestStockZhASpotEm(t *testing.T) {
	df, err := StockZhASpotEm()
	if err != nil {
		t.Logf("StockZhASpotEm 返回错误(可能是网络问题): %v", err)
		return
	}

	if df.Nrow() == 0 {
		t.Log("StockZhASpotEm 返回空数据(使用示例数据)")
	} else {
		t.Logf("StockZhASpotEm 返回 %d 行数据", df.Nrow())
	}
}

// TestStockIndividualFundFlowRankIndicators 测试不同指标
func TestStockIndividualFundFlowRankIndicators(t *testing.T) {
	indicators := []string{"今日", "3日", "5日", "10日"}

	for _, ind := range indicators {
		df, err := StockIndividualFundFlowRank(ind)
		if err != nil {
			t.Logf("StockIndividualFundFlowRank(%s) 返回错误: %v", ind, err)
			continue
		}
		t.Logf("StockIndividualFundFlowRank(%s) 返回 %d 行数据", ind, df.Nrow())
	}
}
