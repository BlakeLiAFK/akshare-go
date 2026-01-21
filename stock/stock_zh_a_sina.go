package stock

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/go-gota/gota/dataframe"
)

// StockZhASpotSina 新浪A股实时行情
func StockZhASpotSina() (dataframe.DataFrame, error) {
	// 获取股票总数
	countURL := ZhSinaAStockCountURL
	resp, err := utils.Get(countURL, nil)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("获取股票数量失败: %w", err)
	}

	countStr := strings.TrimSpace(resp.String())
	totalCount, _ := strconv.Atoi(countStr)
	if totalCount == 0 {
		totalCount = 5000
	}

	// 分页获取数据
	pageSize := 80
	totalPages := (totalCount + pageSize - 1) / pageSize

	headers := []string{"代码", "名称", "最新价", "涨跌额", "涨跌幅", "买入", "卖出", "昨收", "今开", "最高", "最低", "成交量", "成交额"}
	var records [][]string
	records = append(records, headers)

	for page := 1; page <= totalPages && page <= 10; page++ { // 限制最多10页
		payload := map[string]string{
			"page":   strconv.Itoa(page),
			"num":    strconv.Itoa(pageSize),
			"sort":   "symbol",
			"asc":    "1",
			"node":   "hs_a",
			"symbol": "",
			"_s_r_a": "page",
		}

		resp, err := utils.Get(ZhSinaAStockURL, payload)
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

// StockZhAHistSina 新浪A股历史数据
func StockZhAHistSina(symbol, period, adjust string) (dataframe.DataFrame, error) {
	if symbol == "" {
		return dataframe.DataFrame{}, fmt.Errorf("股票代码不能为空")
	}

	if period == "" {
		period = "daily"
	}
	if adjust == "" {
		adjust = ""
	}

	code, market := ParseStockCode(symbol)
	fullSymbol := market + code

	var url string
	switch adjust {
	case "qfq":
		url = fmt.Sprintf(ZhSinaAStockQfqURL, fullSymbol)
	case "hfq":
		url = fmt.Sprintf(ZhSinaAStockHfqURL, fullSymbol)
	default:
		url = fmt.Sprintf(ZhSinaAStockHistURL, fullSymbol)
	}

	resp, err := utils.Get(url, nil)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	text := resp.String()

	// 解析JavaScript格式的数据
	records := parseHistData(text)
	if len(records) <= 1 {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到历史数据")
	}

	return dataframe.LoadRecords(records), nil
}

func parseHistData(text string) [][]string {
	headers := []string{"日期", "开盘", "收盘", "最高", "最低", "成交量"}
	records := [][]string{headers}

	// 简化解析，提取日期和价格数据
	re := regexp.MustCompile(`\[["'](\d{4}-\d{2}-\d{2})["'],\s*([\d.]+),\s*([\d.]+),\s*([\d.]+),\s*([\d.]+),\s*([\d.]+)`)
	matches := re.FindAllStringSubmatch(text, -1)

	for _, match := range matches {
		if len(match) >= 7 {
			records = append(records, match[1:7])
		}
	}

	return records
}

// StockZhAMinuteSina 新浪A股分时数据
func StockZhAMinuteSina(symbol string) (dataframe.DataFrame, error) {
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

	// 提取JSON数据
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

// StockZhATickSina 新浪A股逐笔数据
func StockZhATickSina(symbol, date string) (dataframe.DataFrame, error) {
	if symbol == "" {
		return dataframe.DataFrame{}, fmt.Errorf("股票代码不能为空")
	}

	code, market := ParseStockCode(symbol)
	fullSymbol := market + code

	url := fmt.Sprintf("http://market.finance.sina.com.cn/downxls.php?date=%s&symbol=%s", date, fullSymbol)

	resp, err := utils.Get(url, nil)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	text := resp.String()
	lines := strings.Split(text, "\n")

	headers := []string{"时间", "价格", "涨跌", "成交量", "成交额", "性质"}
	var records [][]string
	records = append(records, headers)

	for i, line := range lines {
		if i == 0 || strings.TrimSpace(line) == "" {
			continue
		}
		fields := strings.Split(line, "\t")
		if len(fields) >= 6 {
			records = append(records, fields[:6])
		}
	}

	return dataframe.LoadRecords(records), nil
}
