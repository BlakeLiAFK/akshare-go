package stock

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/go-gota/gota/dataframe"
)

// StockHkSpotSina 新浪港股实时行情
func StockHkSpotSina() (dataframe.DataFrame, error) {
	url := "http://vip.stock.finance.sina.com.cn/quotes_service/api/json_v2.php/Market_Center.getHKStockData"
	params := map[string]string{
		"page":   "1",
		"num":    "1000",
		"sort":   "symbol",
		"asc":    "1",
		"node":   "hk_main",
		"_s_r_a": "auto",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	var data []map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &data); err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("解析JSON失败: %w", err)
	}

	headers := []string{"代码", "名称", "英文名", "最新价", "涨跌额", "涨跌幅", "昨收", "今开", "最高", "最低", "成交量", "成交额"}
	var records [][]string
	records = append(records, headers)

	for _, item := range data {
		record := []string{
			getString(item, "symbol"),
			getString(item, "name"),
			getString(item, "engname"),
			getString(item, "lasttrade"),
			getString(item, "pricechange"),
			getString(item, "changepercent"),
			getString(item, "prevclose"),
			getString(item, "open"),
			getString(item, "high"),
			getString(item, "low"),
			getString(item, "volume"),
			getString(item, "amount"),
		}
		records = append(records, record)
	}

	return dataframe.LoadRecords(records), nil
}

// StockHkHistSina 新浪港股历史数据
func StockHkHistSina(symbol, period, adjust string) (dataframe.DataFrame, error) {
	if symbol == "" {
		return dataframe.DataFrame{}, fmt.Errorf("股票代码不能为空")
	}

	if period == "" {
		period = "daily"
	}

	url := fmt.Sprintf("https://finance.sina.com.cn/stock/hkstock/%s/hisdata_klc2/klc_kl.js", symbol)

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

// StockHkMinuteSina 新浪港股分时数据
func StockHkMinuteSina(symbol string) (dataframe.DataFrame, error) {
	if symbol == "" {
		return dataframe.DataFrame{}, fmt.Errorf("股票代码不能为空")
	}

	url := fmt.Sprintf("https://quotes.sina.cn/hk/api/jsonp.php/var%%20_%s=/HK_MinKLineService.getHKMinKLineData?symbol=%s&datalen=240", symbol, symbol)

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

	headers := []string{"时间", "价格", "均价", "成交量"}
	var records [][]string
	records = append(records, headers)

	for _, item := range data {
		record := []string{
			getString(item, "d"),
			getString(item, "c"),
			getString(item, "a"),
			getString(item, "v"),
		}
		records = append(records, record)
	}

	return dataframe.LoadRecords(records), nil
}

// StockHkFamousSina 新浪港股知名股票
func StockHkFamousSina(symbol string) (dataframe.DataFrame, error) {
	if symbol == "" {
		symbol = "blue_chip" // 蓝筹股
	}

	nodeMap := map[string]string{
		"blue_chip": "hk_blue_chip",
		"red_chip":  "hk_red_chip",
		"state":     "hk_state_owned",
		"hs_index":  "hk_hs_index",
		"main":      "hk_main",
	}

	node := nodeMap[symbol]
	if node == "" {
		node = "hk_main"
	}

	url := "http://vip.stock.finance.sina.com.cn/quotes_service/api/json_v2.php/Market_Center.getHKStockData"
	params := map[string]string{
		"page":   "1",
		"num":    "100",
		"sort":   "symbol",
		"asc":    "1",
		"node":   node,
		"_s_r_a": "auto",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	var data []map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &data); err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("解析JSON失败: %w", err)
	}

	headers := []string{"代码", "名称", "最新价", "涨跌幅", "成交量", "成交额"}
	var records [][]string
	records = append(records, headers)

	for _, item := range data {
		record := []string{
			getString(item, "symbol"),
			getString(item, "name"),
			getString(item, "lasttrade"),
			getString(item, "changepercent"),
			getString(item, "volume"),
			getString(item, "amount"),
		}
		records = append(records, record)
	}

	return dataframe.LoadRecords(records), nil
}
