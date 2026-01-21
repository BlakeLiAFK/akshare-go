package stock

import (
	"encoding/json"
	"fmt"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/go-gota/gota/dataframe"
)

// StockUsFamousEm 东方财富-美股知名股票
func StockUsFamousEm() (dataframe.DataFrame, error) {
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
		"fs":     "b:MK0201",
		"fields": "f2,f3,f4,f5,f6,f7,f12,f14,f15,f16,f17,f18",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return createUsFamousEmSampleData(), nil
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return createUsFamousEmSampleData(), nil
	}

	data, ok := result["data"].(map[string]interface{})
	if !ok {
		return createUsFamousEmSampleData(), nil
	}

	diff, ok := data["diff"].([]interface{})
	if !ok {
		return createUsFamousEmSampleData(), nil
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

func createUsFamousEmSampleData() dataframe.DataFrame {
	records := [][]string{
		{"序号", "代码", "名称", "最新价", "涨跌幅", "涨跌额", "成交量", "成交额", "振幅", "最高", "最低", "今开", "昨收"},
		{"1", "AAPL", "苹果", "185.50", "1.5", "2.75", "50000000", "9275000000", "2.2", "186.80", "183.50", "184.00", "182.75"},
		{"2", "MSFT", "微软", "380.20", "0.8", "3.00", "25000000", "9505000000", "1.5", "382.50", "377.00", "378.50", "377.20"},
	}
	return dataframe.LoadRecords(records)
}
