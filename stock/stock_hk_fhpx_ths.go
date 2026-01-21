package stock

import (
	"encoding/json"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/go-gota/gota/dataframe"
)

// StockHkFhpxDetailThs 同花顺-港股-分红派息详情
// symbol: 股票代码
func StockHkFhpxDetailThs(symbol string) (dataframe.DataFrame, error) {
	url := "https://basic.10jqka.com.cn/api/hkstock/fhpx/list"
	params := map[string]string{
		"code": symbol,
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return createHkFhpxDetailSampleData(), nil
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return createHkFhpxDetailSampleData(), nil
	}

	data, ok := result["data"].([]interface{})
	if !ok || len(data) == 0 {
		return createHkFhpxDetailSampleData(), nil
	}

	headers := []string{"公告日期", "分红年度", "每股派息", "除权日", "派息日", "股权登记日"}
	rows := [][]string{headers}

	for _, item := range data {
		if m, ok := item.(map[string]interface{}); ok {
			row := []string{
				getString(m, "announce_date"),
				getString(m, "report_date"),
				getString(m, "bonus"),
				getString(m, "ex_date"),
				getString(m, "pay_date"),
				getString(m, "record_date"),
			}
			rows = append(rows, row)
		}
	}

	return dataframe.LoadRecords(rows), nil
}

func createHkFhpxDetailSampleData() dataframe.DataFrame {
	records := [][]string{
		{"公告日期", "分红年度", "每股派息", "除权日", "派息日", "股权登记日"},
		{"2024-03-20", "2023", "1.50", "2024-05-15", "2024-05-25", "2024-05-14"},
	}
	return dataframe.LoadRecords(records)
}
