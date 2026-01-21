package stock

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/go-gota/gota/dataframe"
)

// StockZhKcbSpotSina 新浪科创板实时行情
func StockZhKcbSpotSina() (dataframe.DataFrame, error) {
	countURL := ZhSinaKcbStockCountURL
	resp, err := utils.Get(countURL, nil)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("获取股票数量失败: %w", err)
	}

	countStr := resp.String()
	totalCount, _ := strconv.Atoi(countStr)
	if totalCount == 0 {
		totalCount = 600
	}

	pageSize := 80
	totalPages := (totalCount + pageSize - 1) / pageSize

	headers := []string{"代码", "名称", "最新价", "涨跌额", "涨跌幅", "买入", "卖出", "昨收", "今开", "最高", "最低", "成交量", "成交额"}
	var records [][]string
	records = append(records, headers)

	for page := 1; page <= totalPages; page++ {
		payload := map[string]string{
			"page":   strconv.Itoa(page),
			"num":    strconv.Itoa(pageSize),
			"sort":   "symbol",
			"asc":    "1",
			"node":   "kcb",
			"symbol": "",
			"_s_r_a": "auto",
		}

		resp, err := utils.Get(ZhSinaKcbStockURL, payload)
		if err != nil {
			continue
		}

		var data []map[string]interface{}
		if err := json.Unmarshal(resp.Body(), &data); err != nil {
			continue
		}

		for _, item := range data {
			record := []string{
				getString(item, "symbol"),
				getString(item, "name"),
				getString(item, "trade"),
				getString(item, "pricechange"),
				getString(item, "changepercent"),
				getString(item, "buy"),
				getString(item, "sell"),
				getString(item, "settlement"),
				getString(item, "open"),
				getString(item, "high"),
				getString(item, "low"),
				getString(item, "volume"),
				getString(item, "amount"),
			}
			records = append(records, record)
		}
	}

	return dataframe.LoadRecords(records), nil
}

// StockZhKcbHistSina 新浪科创板历史数据
func StockZhKcbHistSina(symbol, period, adjust string) (dataframe.DataFrame, error) {
	if symbol == "" {
		return dataframe.DataFrame{}, fmt.Errorf("股票代码不能为空")
	}

	code, market := ParseStockCode(symbol)
	fullSymbol := market + code

	var url string
	switch adjust {
	case "qfq":
		url = fmt.Sprintf(ZhSinaKcbStockQfqURL, fullSymbol)
	case "hfq":
		url = fmt.Sprintf(ZhSinaKcbStockHfqURL, fullSymbol)
	default:
		url = fmt.Sprintf(ZhSinaKcbStockHistURL, fullSymbol, period, fullSymbol)
	}

	resp, err := utils.Get(url, nil)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	text := resp.String()
	records := parseHistData(text)
	if len(records) <= 1 {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到历史数据")
	}

	return dataframe.LoadRecords(records), nil
}
