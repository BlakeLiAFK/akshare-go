package stock

import (
	"encoding/json"
	"fmt"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/go-gota/gota/dataframe"
)

// StockHsgtNorthNetFlowIn 沪股通/深股通资金流向-北向资金净流入
func StockHsgtNorthNetFlowIn() (dataframe.DataFrame, error) {
	url := "http://push2his.eastmoney.com/api/qt/kamt.kline/get"
	params := map[string]string{
		"fields1": "f1,f3,f5",
		"fields2": "f51,f52,f53,f54,f55,f56",
		"klt":     "101",
		"lmt":     "500",
		"ut":      "b2884a393a59ad64002292a3e90d46a5",
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

	headers := []string{"日期", "沪股通净流入", "深股通净流入", "北向资金净流入", "沪股通余额", "深股通余额"}
	var records [][]string
	records = append(records, headers)

	if hk2sh, ok := data["hk2sh"].([]interface{}); ok {
		for _, item := range hk2sh {
			if str, ok := item.(string); ok {
				fields := splitFields(str)
				if len(fields) >= 6 {
					records = append(records, fields[:6])
				}
			}
		}
	}

	return dataframe.LoadRecords(records), nil
}

func splitFields(s string) []string {
	var fields []string
	start := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ',' {
			fields = append(fields, s[start:i])
			start = i + 1
		}
	}
	if start < len(s) {
		fields = append(fields, s[start:])
	}
	return fields
}

// StockHsgtHoldStock 沪深港通持股
func StockHsgtHoldStock(market, indicator string) (dataframe.DataFrame, error) {
	if market == "" {
		market = "北向"
	}
	if indicator == "" {
		indicator = "今日排行"
	}

	marketMap := map[string]string{
		"北向":  "001",
		"沪股通": "002",
		"深股通": "003",
	}

	mktCode := marketMap[market]
	if mktCode == "" {
		return dataframe.DataFrame{}, fmt.Errorf("不支持的市场: %s", market)
	}

	url := "http://datacenter-web.eastmoney.com/api/data/v1/get"
	params := map[string]string{
		"sortColumns": "ADD_MARKET_CAP",
		"sortTypes":   "-1",
		"pageSize":    "500",
		"pageNumber":  "1",
		"reportName":  "RPT_MUTUAL_HOLD_DET",
		"columns":     "ALL",
		"filter":      fmt.Sprintf("(MARKET_CODE=\"%s\")", mktCode),
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("解析JSON失败: %w", err)
	}

	resultData, ok := result["result"].(map[string]interface{})
	if !ok {
		return dataframe.DataFrame{}, fmt.Errorf("数据格式错误")
	}

	data, ok := resultData["data"].([]interface{})
	if !ok {
		return dataframe.DataFrame{}, fmt.Errorf("数据格式错误")
	}

	headers := []string{"日期", "代码", "名称", "持股数量", "持股市值", "持股数量变化", "持股市值变化", "持股占比"}
	var records [][]string
	records = append(records, headers)

	for _, item := range data {
		if m, ok := item.(map[string]interface{}); ok {
			record := []string{
				getString(m, "TRADE_DATE"),
				getString(m, "SECURITY_CODE"),
				getString(m, "SECURITY_NAME"),
				getString(m, "HOLD_SHARES"),
				getString(m, "HOLD_MARKET_CAP"),
				getString(m, "SHARES_CHANGE"),
				getString(m, "MARKET_CAP_CHANGE"),
				getString(m, "HOLD_RATIO"),
			}
			records = append(records, record)
		}
	}

	return dataframe.LoadRecords(records), nil
}

// StockHsgtBoardRank 沪深港通板块排行
func StockHsgtBoardRank() (dataframe.DataFrame, error) {
	url := "http://datacenter-web.eastmoney.com/api/data/v1/get"
	params := map[string]string{
		"sortColumns": "BOARD_NET_AMT",
		"sortTypes":   "-1",
		"pageSize":    "100",
		"pageNumber":  "1",
		"reportName":  "RPT_MUTUAL_BOARD_STATISTICS",
		"columns":     "ALL",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("解析JSON失败: %w", err)
	}

	resultData, ok := result["result"].(map[string]interface{})
	if !ok {
		return dataframe.DataFrame{}, fmt.Errorf("数据格式错误")
	}

	data, ok := resultData["data"].([]interface{})
	if !ok {
		return dataframe.DataFrame{}, fmt.Errorf("数据格式错误")
	}

	headers := []string{"日期", "板块名称", "板块涨跌幅", "北向净买入", "北向持股市值", "北向持股变化"}
	var records [][]string
	records = append(records, headers)

	for _, item := range data {
		if m, ok := item.(map[string]interface{}); ok {
			record := []string{
				getString(m, "TRADE_DATE"),
				getString(m, "BOARD_NAME"),
				getString(m, "BOARD_PCT_CHG"),
				getString(m, "BOARD_NET_AMT"),
				getString(m, "BOARD_HOLD_MV"),
				getString(m, "BOARD_MV_CHANGE"),
			}
			records = append(records, record)
		}
	}

	return dataframe.LoadRecords(records), nil
}
