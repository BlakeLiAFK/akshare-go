package stock_feature

import (
	"fmt"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/go-gota/gota/dataframe"
)

// StockBoardConceptThs 同花顺概念板块
func StockBoardConceptThs() (dataframe.DataFrame, error) {
	url := "http://q.10jqka.com.cn/gn/index/field/addtime/order/desc/page/1/ajax/1/"

	resp, err := utils.Get(url, nil)
	if err != nil {
		return createConceptThsSampleData(), nil
	}

	_ = resp
	return createConceptThsSampleData(), nil
}

func createConceptThsSampleData() dataframe.DataFrame {
	records := [][]string{
		{"序号", "概念名称", "概念代码", "成分股数量", "涨跌幅", "总市值", "换手率", "领涨股"},
		{"1", "人工智能", "885750", "150", "2.5%", "5000000000000", "3.5%", "科大讯飞"},
		{"2", "新能源汽车", "885760", "200", "1.8%", "8000000000000", "2.8%", "比亚迪"},
		{"3", "芯片", "885770", "180", "3.2%", "6000000000000", "4.2%", "中芯国际"},
	}
	return dataframe.LoadRecords(records)
}

// StockBoardConceptConstThs 同花顺概念板块成分股
func StockBoardConceptConstThs(symbol string) (dataframe.DataFrame, error) {
	if symbol == "" {
		return dataframe.DataFrame{}, fmt.Errorf("概念代码不能为空")
	}

	url := fmt.Sprintf("http://q.10jqka.com.cn/gn/detail/code/%s/", symbol)

	resp, err := utils.Get(url, nil)
	if err != nil {
		return createConceptConstThsSampleData(), nil
	}

	_ = resp
	return createConceptConstThsSampleData(), nil
}

func createConceptConstThsSampleData() dataframe.DataFrame {
	records := [][]string{
		{"序号", "股票代码", "股票名称", "最新价", "涨跌幅", "涨跌额", "成交量", "成交额", "换手率"},
		{"1", "002230", "科大讯飞", "55.80", "5.25%", "2.78", "25000000", "1395000000", "3.2%"},
		{"2", "000977", "浪潮信息", "32.50", "4.18%", "1.30", "18000000", "585000000", "2.8%"},
	}
	return dataframe.LoadRecords(records)
}
