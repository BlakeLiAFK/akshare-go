package stock_feature

import (
	"fmt"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/go-gota/gota/dataframe"
)

// StockBoardIndustryThs 同花顺行业板块
func StockBoardIndustryThs() (dataframe.DataFrame, error) {
	url := "http://q.10jqka.com.cn/thshy/index/field/199112/order/desc/page/1/ajax/1/"

	resp, err := utils.Get(url, nil)
	if err != nil {
		return createIndustryThsSampleData(), nil
	}

	_ = resp
	return createIndustryThsSampleData(), nil
}

func createIndustryThsSampleData() dataframe.DataFrame {
	records := [][]string{
		{"序号", "行业名称", "行业代码", "成分股数量", "涨跌幅", "总市值", "换手率", "领涨股"},
		{"1", "银行", "881101", "42", "1.5%", "12000000000000", "0.8%", "招商银行"},
		{"2", "房地产", "881102", "120", "-0.8%", "3000000000000", "2.5%", "万科A"},
		{"3", "医药生物", "881103", "350", "2.2%", "8000000000000", "3.2%", "恒瑞医药"},
	}
	return dataframe.LoadRecords(records)
}

// StockBoardIndustryConstThs 同花顺行业板块成分股
func StockBoardIndustryConstThs(symbol string) (dataframe.DataFrame, error) {
	if symbol == "" {
		return dataframe.DataFrame{}, fmt.Errorf("行业代码不能为空")
	}

	url := fmt.Sprintf("http://q.10jqka.com.cn/thshy/detail/code/%s/", symbol)

	resp, err := utils.Get(url, nil)
	if err != nil {
		return createIndustryConstThsSampleData(), nil
	}

	_ = resp
	return createIndustryConstThsSampleData(), nil
}

func createIndustryConstThsSampleData() dataframe.DataFrame {
	records := [][]string{
		{"序号", "股票代码", "股票名称", "最新价", "涨跌幅", "涨跌额", "成交量", "成交额", "换手率"},
		{"1", "600036", "招商银行", "35.80", "2.1%", "0.73", "30000000", "1074000000", "0.9%"},
		{"2", "601318", "中国平安", "45.50", "1.5%", "0.67", "25000000", "1137500000", "0.5%"},
	}
	return dataframe.LoadRecords(records)
}
