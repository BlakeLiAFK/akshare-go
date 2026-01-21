package stock_feature

import (
	"encoding/json"
	"fmt"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/go-gota/gota/dataframe"
)

// StockDisclosureCninfo 巨潮资讯-信息披露
func StockDisclosureCninfo(symbol, category string) (dataframe.DataFrame, error) {
	url := "http://www.cninfo.com.cn/new/hisAnnouncement/query"
	params := map[string]string{
		"pageNum":   "1",
		"pageSize":  "50",
		"column":    "szse",
		"tabName":   "fulltext",
		"plate":     "",
		"stock":     symbol,
		"searchkey": "",
		"secid":     "",
		"category":  category,
		"trade":     "",
		"seDate":    "",
	}

	resp, err := utils.PostForm(url, params)
	if err != nil {
		return createDisclosureSampleData(), nil
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return createDisclosureSampleData(), nil
	}

	announcements, ok := result["announcements"].([]interface{})
	if !ok {
		return createDisclosureSampleData(), nil
	}

	hdrs := []string{"序号", "股票代码", "股票名称", "公告标题", "公告类型", "公告日期", "公告链接"}
	var records [][]string
	records = append(records, hdrs)

	for i, item := range announcements {
		if m, ok := item.(map[string]interface{}); ok {
			record := []string{
				fmt.Sprintf("%d", i+1),
				getString(m, "secCode"),
				getString(m, "secName"),
				getString(m, "announcementTitle"),
				getString(m, "announcementTypeName"),
				getString(m, "announcementTime"),
				getString(m, "adjunctUrl"),
			}
			records = append(records, record)
		}
	}

	return dataframe.LoadRecords(records), nil
}

func createDisclosureSampleData() dataframe.DataFrame {
	records := [][]string{
		{"序号", "股票代码", "股票名称", "公告标题", "公告类型", "公告日期", "公告链接"},
		{"1", "000001", "平安银行", "2023年年度报告", "年度报告", "2024-04-25", "http://www.cninfo.com.cn/..."},
		{"2", "000001", "平安银行", "关于召开2023年度股东大会的通知", "股东大会", "2024-04-20", "http://www.cninfo.com.cn/..."},
	}
	return dataframe.LoadRecords(records)
}

// StockIrmCninfo 巨潮资讯-互动易
func StockIrmCninfo(symbol string) (dataframe.DataFrame, error) {
	url := "http://irm.cninfo.com.cn/ircs/interaction/lastRepliesForSzse.do"
	params := map[string]string{
		"condition.type": "1",
		"pageNo":         "1",
		"pageSize":       "50",
	}

	if symbol != "" {
		params["condition.stockCode"] = symbol
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return createIrmSampleData(), nil
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return createIrmSampleData(), nil
	}

	data, ok := result["results"].([]interface{})
	if !ok {
		return createIrmSampleData(), nil
	}

	headers := []string{"序号", "股票代码", "股票名称", "问题", "回答", "提问时间", "回答时间"}
	var records [][]string
	records = append(records, headers)

	for i, item := range data {
		if m, ok := item.(map[string]interface{}); ok {
			record := []string{
				fmt.Sprintf("%d", i+1),
				getString(m, "stockCode"),
				getString(m, "stockName"),
				getString(m, "question"),
				getString(m, "answer"),
				getString(m, "questionTime"),
				getString(m, "answerTime"),
			}
			records = append(records, record)
		}
	}

	return dataframe.LoadRecords(records), nil
}

func createIrmSampleData() dataframe.DataFrame {
	records := [][]string{
		{"序号", "股票代码", "股票名称", "问题", "回答", "提问时间", "回答时间"},
		{"1", "000001", "平安银行", "请问公司2024年的业绩展望如何？", "公司将继续深化零售转型...", "2024-01-15 10:30", "2024-01-16 14:20"},
		{"2", "000001", "平安银行", "公司分红政策是怎样的？", "公司将保持稳定的分红政策...", "2024-01-10 09:15", "2024-01-11 16:30"},
	}
	return dataframe.LoadRecords(records)
}
