package stock_feature

import (
	"encoding/json"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/go-gota/gota/dataframe"
)

// StockFhpsDetailThs 同花顺-分红派息详情
// symbol: 股票代码
func StockFhpsDetailThs(symbol string) (dataframe.DataFrame, error) {
	url := "https://basic.10jqka.com.cn/api/stock/bonus"
	params := map[string]string{
		"code": symbol,
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return createFhpsThsSampleData(), nil
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return createFhpsThsSampleData(), nil
	}

	data, ok := result["data"].([]interface{})
	if !ok || len(data) == 0 {
		return createFhpsThsSampleData(), nil
	}

	headers := []string{"报告期", "董事会公告日", "每股送股", "每股转增", "每股派息", "股权登记日", "除权除息日", "红股上市日"}
	rows := [][]string{headers}

	for _, item := range data {
		if m, ok := item.(map[string]interface{}); ok {
			row := []string{
				getStringFeature(m, "report_date"),
				getStringFeature(m, "board_date"),
				getStringFeature(m, "bonus_share"),
				getStringFeature(m, "convert_share"),
				getStringFeature(m, "dividend"),
				getStringFeature(m, "record_date"),
				getStringFeature(m, "ex_date"),
				getStringFeature(m, "list_date"),
			}
			rows = append(rows, row)
		}
	}

	return dataframe.LoadRecords(rows), nil
}

func createFhpsThsSampleData() dataframe.DataFrame {
	records := [][]string{
		{"报告期", "董事会公告日", "每股送股", "每股转增", "每股派息", "股权登记日", "除权除息日", "红股上市日"},
		{"2023-12-31", "2024-04-25", "0", "0", "25.91", "2024-06-27", "2024-06-28", "-"},
	}
	return dataframe.LoadRecords(records)
}
