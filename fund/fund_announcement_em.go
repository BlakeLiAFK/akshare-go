package fund

import (
	"fmt"
	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/tidwall/gjson"
)

// FundAnnouncementDividendEm 获取基金分红配送公告
// 参数: symbol 基金代码，如 "000001"
func FundAnnouncementDividendEm(symbol string) ([]map[string]interface{}, error) {
	if symbol == "" {
		symbol = "000001"
	}

	url := "https://api.fund.eastmoney.com/f10/JJGG"
	params := map[string]string{
		"fundcode":  symbol,
		"pageIndex": "1",
		"pageSize":  "200",
		"type":      "1",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	result := gjson.ParseBytes(resp.Body())
	items := result.Get("Data.LSJJGGList").Array()

	records := make([]map[string]interface{}, 0, len(items))
	for _, item := range items {
		record := map[string]interface{}{
			"公告标题": item.Get("TITLE").String(),
			"公告日期": item.Get("PUBDATE").String(),
			"公告类型": "分红配送",
		}
		records = append(records, record)
	}

	return records, nil
}

// FundAnnouncementReportEm 获取基金定期报告公告
// 参数: symbol 基金代码，如 "000001"
func FundAnnouncementReportEm(symbol string) ([]map[string]interface{}, error) {
	if symbol == "" {
		symbol = "000001"
	}

	url := "https://api.fund.eastmoney.com/f10/JJGG"
	params := map[string]string{
		"fundcode":  symbol,
		"pageIndex": "1",
		"pageSize":  "200",
		"type":      "2",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	result := gjson.ParseBytes(resp.Body())
	items := result.Get("Data.LSJJGGList").Array()

	records := make([]map[string]interface{}, 0, len(items))
	for _, item := range items {
		record := map[string]interface{}{
			"公告标题": item.Get("TITLE").String(),
			"公告日期": item.Get("PUBDATE").String(),
			"公告类型": "定期报告",
		}
		records = append(records, record)
	}

	return records, nil
}

// FundAnnouncementPersonnelEm 获取基金人事调整公告
// 参数: symbol 基金代码，如 "000001"
func FundAnnouncementPersonnelEm(symbol string) ([]map[string]interface{}, error) {
	if symbol == "" {
		symbol = "000001"
	}

	url := "https://api.fund.eastmoney.com/f10/JJGG"
	params := map[string]string{
		"fundcode":  symbol,
		"pageIndex": "1",
		"pageSize":  "200",
		"type":      "3",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	result := gjson.ParseBytes(resp.Body())
	items := result.Get("Data.LSJJGGList").Array()

	records := make([]map[string]interface{}, 0, len(items))
	for _, item := range items {
		record := map[string]interface{}{
			"公告标题": item.Get("TITLE").String(),
			"公告日期": item.Get("PUBDATE").String(),
			"公告类型": "人事调整",
		}
		records = append(records, record)
	}

	return records, nil
}
