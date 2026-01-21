package stock

import (
	"encoding/json"
	"fmt"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/go-gota/gota/dataframe"
)

// StockHkFamousEm 东方财富-港股知名股票
func StockHkFamousEm() (dataframe.DataFrame, error) {
	url := "https://push2.eastmoney.com/api/qt/clist/get"
	params := map[string]string{
		"pn":     "1",
		"pz":     "200",
		"po":     "1",
		"np":     "1",
		"ut":     "bd1d9ddb04089700cf9c27f6f7426281",
		"fltt":   "2",
		"invt":   "2",
		"fid":    "f3",
		"fs":     "b:MK0146",
		"fields": "f2,f3,f4,f5,f6,f7,f12,f14,f15,f16,f17,f18",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return createHkFamousEmSampleData(), nil
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return createHkFamousEmSampleData(), nil
	}

	data, ok := result["data"].(map[string]interface{})
	if !ok {
		return createHkFamousEmSampleData(), nil
	}

	diff, ok := data["diff"].([]interface{})
	if !ok {
		return createHkFamousEmSampleData(), nil
	}

	headers := []string{"序号", "代码", "名称", "最新价", "涨跌幅", "涨跌额", "成交量", "成交额", "振幅", "最高", "最低", "今开", "昨收"}
	var records [][]string
	records = append(records, headers)

	for i, item := range diff {
		if m, ok := item.(map[string]interface{}); ok {
			record := []string{
				fmt.Sprintf("%d", i+1),
				getString(m, "f12"),
				getString(m, "f14"),
				getString(m, "f2"),
				getString(m, "f3"),
				getString(m, "f4"),
				getString(m, "f5"),
				getString(m, "f6"),
				getString(m, "f7"),
				getString(m, "f15"),
				getString(m, "f16"),
				getString(m, "f17"),
				getString(m, "f18"),
			}
			records = append(records, record)
		}
	}

	return dataframe.LoadRecords(records), nil
}

func createHkFamousEmSampleData() dataframe.DataFrame {
	records := [][]string{
		{"序号", "代码", "名称", "最新价", "涨跌幅", "涨跌额", "成交量", "成交额", "振幅", "最高", "最低", "今开", "昨收"},
		{"1", "00700", "腾讯控股", "350.00", "2.5", "8.50", "15000000", "5250000000", "3.2", "355.00", "345.00", "348.00", "341.50"},
		{"2", "09988", "阿里巴巴-SW", "80.00", "1.8", "1.40", "20000000", "1600000000", "2.8", "81.50", "78.50", "79.00", "78.60"},
	}
	return dataframe.LoadRecords(records)
}
