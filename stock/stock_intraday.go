package stock

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/go-gota/gota/dataframe"
)

// StockIntradaySina 新浪-分时数据
func StockIntradaySina(symbol string) (dataframe.DataFrame, error) {
	if symbol == "" {
		return dataframe.DataFrame{}, fmt.Errorf("股票代码不能为空")
	}

	code, market := ParseStockCode(symbol)
	fullSymbol := market + code

	url := fmt.Sprintf("https://quotes.sina.cn/cn/api/jsonp_v3.php/var%%20_%s=/CN_MarketDataService.getMinKLineData?symbol=%s&datalen=240", fullSymbol, fullSymbol)

	resp, err := utils.Get(url, nil)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	text := resp.String()
	start := strings.Index(text, "[")
	end := strings.LastIndex(text, "]")
	if start == -1 || end == -1 || start >= end {
		return dataframe.DataFrame{}, fmt.Errorf("数据格式错误")
	}

	jsonStr := text[start : end+1]

	var data []map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &data); err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("解析JSON失败: %w", err)
	}

	headers := []string{"时间", "开盘", "收盘", "最高", "最低", "成交量"}
	var records [][]string
	records = append(records, headers)

	for _, item := range data {
		record := []string{
			getString(item, "d"),
			getString(item, "o"),
			getString(item, "c"),
			getString(item, "h"),
			getString(item, "l"),
			getString(item, "v"),
		}
		records = append(records, record)
	}

	return dataframe.LoadRecords(records), nil
}

// StockAskBidEm 东方财富-买卖盘数据
func StockAskBidEm(symbol string) (dataframe.DataFrame, error) {
	if symbol == "" {
		return dataframe.DataFrame{}, fmt.Errorf("股票代码不能为空")
	}

	code, market := ParseStockCode(symbol)
	secid := fmt.Sprintf("%s.%s", MarketCodeMap[market], code)

	url := "http://push2.eastmoney.com/api/qt/stock/get"
	params := map[string]string{
		"secid":  secid,
		"fields": "f31,f32,f33,f34,f35,f36,f37,f38,f39,f40,f19,f20,f17,f18,f15,f16,f13,f14,f11,f12",
		"ut":     "b2884a393a59ad64002292a3e90d46a5",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("解析JSON失败: %w", err)
	}

	data, ok := result["data"].(map[string]interface{})
	if !ok {
		return dataframe.DataFrame{}, fmt.Errorf("数据格式错误")
	}

	headers := []string{"档位", "买入价", "买入量", "卖出价", "卖出量"}
	var records [][]string
	records = append(records, headers)

	// 解析买卖五档数据
	buyFields := []string{"f11", "f13", "f15", "f17", "f19"}
	buyVolFields := []string{"f12", "f14", "f16", "f18", "f20"}
	sellFields := []string{"f31", "f33", "f35", "f37", "f39"}
	sellVolFields := []string{"f32", "f34", "f36", "f38", "f40"}

	for i := 0; i < 5; i++ {
		record := []string{
			fmt.Sprintf("第%d档", i+1),
			getString(data, buyFields[i]),
			getString(data, buyVolFields[i]),
			getString(data, sellFields[i]),
			getString(data, sellVolFields[i]),
		}
		records = append(records, record)
	}

	return dataframe.LoadRecords(records), nil
}
