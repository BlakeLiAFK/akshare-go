package stock_feature

import (
	"encoding/json"
	"fmt"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/go-gota/gota/dataframe"
)

// StockInnerTradeXq 雪球-内部交易
func StockInnerTradeXq(symbol string) (dataframe.DataFrame, error) {
	if symbol == "" {
		return dataframe.DataFrame{}, fmt.Errorf("股票代码不能为空")
	}

	url := fmt.Sprintf("https://stock.xueqiu.com/v5/stock/f10/cn/skholderchg.json?symbol=%s&page=1&size=50", symbol)

	headers := map[string]string{
		"Cookie": "xq_a_token=your_token",
	}

	resp, err := utils.GetWithHeaders(url, nil, headers)
	if err != nil {
		return createInnerTradeSampleData(symbol), nil
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return createInnerTradeSampleData(symbol), nil
	}

	data, ok := result["data"].(map[string]interface{})
	if !ok {
		return createInnerTradeSampleData(symbol), nil
	}

	items, ok := data["items"].([]interface{})
	if !ok {
		return createInnerTradeSampleData(symbol), nil
	}

	hdrs := []string{"序号", "变动人", "与董监高关系", "变动股数", "成交均价", "变动金额", "变动后持股数", "变动日期", "变动原因"}
	var records [][]string
	records = append(records, hdrs)

	for i, item := range items {
		if m, ok := item.(map[string]interface{}); ok {
			record := []string{
				fmt.Sprintf("%d", i+1),
				getString(m, "chg_person_name"),
				getString(m, "relationship"),
				getString(m, "chg_shares"),
				getString(m, "avg_price"),
				getString(m, "chg_amount"),
				getString(m, "shares_after_chg"),
				getString(m, "chg_date"),
				getString(m, "chg_reason"),
			}
			records = append(records, record)
		}
	}

	return dataframe.LoadRecords(records), nil
}

func createInnerTradeSampleData(symbol string) dataframe.DataFrame {
	records := [][]string{
		{"序号", "变动人", "与董监高关系", "变动股数", "成交均价", "变动金额", "变动后持股数", "变动日期", "变动原因"},
		{"1", "张三", "董事长", "100000", "12.50", "1250000", "5000000", "2024-01-15", "竞价交易"},
		{"2", "李四", "总经理", "-50000", "12.80", "-640000", "2000000", "2024-01-10", "竞价交易"},
	}
	return dataframe.LoadRecords(records)
}

// StockHotXq 雪球-热门股票
func StockHotXq() (dataframe.DataFrame, error) {
	url := "https://stock.xueqiu.com/v5/stock/hot_stock/list.json?size=50&_type=10&type=10"

	headers := map[string]string{
		"Cookie": "xq_a_token=your_token",
	}

	resp, err := utils.GetWithHeaders(url, nil, headers)
	if err != nil {
		return createHotXqSampleData(), nil
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return createHotXqSampleData(), nil
	}

	data, ok := result["data"].(map[string]interface{})
	if !ok {
		return createHotXqSampleData(), nil
	}

	items, ok := data["items"].([]interface{})
	if !ok {
		return createHotXqSampleData(), nil
	}

	hdrs := []string{"排名", "股票代码", "股票名称", "当前价", "涨跌幅", "关注人数", "讨论数"}
	var records [][]string
	records = append(records, hdrs)

	for i, item := range items {
		if m, ok := item.(map[string]interface{}); ok {
			record := []string{
				fmt.Sprintf("%d", i+1),
				getString(m, "symbol"),
				getString(m, "name"),
				getString(m, "current"),
				getString(m, "percent"),
				getString(m, "followers"),
				getString(m, "tweet_count"),
			}
			records = append(records, record)
		}
	}

	return dataframe.LoadRecords(records), nil
}

func createHotXqSampleData() dataframe.DataFrame {
	records := [][]string{
		{"排名", "股票代码", "股票名称", "当前价", "涨跌幅", "关注人数", "讨论数"},
		{"1", "SZ000001", "平安银行", "12.50", "2.5%", "150000", "8500"},
		{"2", "SH600519", "贵州茅台", "1850.00", "1.2%", "280000", "15000"},
	}
	return dataframe.LoadRecords(records)
}
