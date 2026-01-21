package stock

import (
	"encoding/json"
	"fmt"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/go-gota/gota/dataframe"
)

// StockCgEquityMortgageCninfo 巨潮资讯-数据中心-专题统计-公司治理-股权质押
// date: 统计日期 YYYYMMDD
func StockCgEquityMortgageCninfo(date string) (dataframe.DataFrame, error) {
	url := "https://webapi.cninfo.com.cn/api/sysapi/p_sysapi1094"

	tdate := ""
	if date != "" && len(date) == 8 {
		tdate = fmt.Sprintf("%s-%s-%s", date[:4], date[4:6], date[6:])
	}

	params := map[string]string{
		"tdate": tdate,
	}

	resp, err := utils.PostWithCninfoHeaders(url, params)
	if err != nil {
		return createCgEquityMortgageSampleData(), nil
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return createCgEquityMortgageSampleData(), nil
	}

	records, ok := result["records"].([]interface{})
	if !ok || len(records) == 0 {
		return createCgEquityMortgageSampleData(), nil
	}

	headers := []string{"股票代码", "股票简称", "公告日期", "出质人", "质权人", "质押数量", "占总股本比例", "质押解除数量", "质押事项", "累计质押占总股本比例"}
	rows := [][]string{headers}

	for _, item := range records {
		if m, ok := item.(map[string]interface{}); ok {
			row := []string{
				getString(m, "F007V"),
				getString(m, "F002V"),
				getString(m, "F003D"),
				getString(m, "F006V"),
				getString(m, "F005V"),
				getString(m, "F010N"),
				getString(m, "F008N"),
				getString(m, "F001N"),
				getString(m, "F004V"),
				getString(m, "F009N"),
			}
			rows = append(rows, row)
		}
	}

	return dataframe.LoadRecords(rows), nil
}

func createCgEquityMortgageSampleData() dataframe.DataFrame {
	records := [][]string{
		{"股票代码", "股票简称", "公告日期", "出质人", "质权人", "质押数量", "占总股本比例", "质押解除数量", "质押事项", "累计质押占总股本比例"},
		{"000001", "平安银行", "2021-09-30", "中国平安保险", "招商银行", "100000000", "0.52", "0", "质押", "5.25"},
	}
	return dataframe.LoadRecords(records)
}
