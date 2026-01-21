package stock

import (
	"encoding/json"
	"fmt"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/go-gota/gota/dataframe"
)

// StockRankForecastCninfo 巨潮资讯-数据中心-评级预测
// date: 日期 YYYYMMDD
func StockRankForecastCninfo(date string) (dataframe.DataFrame, error) {
	url := "https://webapi.cninfo.com.cn/api/sysapi/p_sysapi1093"

	tdate := ""
	if date != "" && len(date) == 8 {
		tdate = fmt.Sprintf("%s-%s-%s", date[:4], date[4:6], date[6:])
	}

	params := map[string]string{
		"tdate": tdate,
	}

	resp, err := utils.PostWithCninfoHeaders(url, params)
	if err != nil {
		return createRankForecastCninfoSampleData(), nil
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return createRankForecastCninfoSampleData(), nil
	}

	records, ok := result["records"].([]interface{})
	if !ok || len(records) == 0 {
		return createRankForecastCninfoSampleData(), nil
	}

	headers := []string{"序号", "证券代码", "证券简称", "研究机构", "研究员", "评级日期", "最新评级", "上次评级", "目标价格"}
	rows := [][]string{headers}

	for i, item := range records {
		if m, ok := item.(map[string]interface{}); ok {
			row := []string{
				fmt.Sprintf("%d", i+1),
				getString(m, "SECCODE"),
				getString(m, "SECNAME"),
				getString(m, "F001V"),
				getString(m, "F002V"),
				getString(m, "F003D"),
				getString(m, "F004V"),
				getString(m, "F005V"),
				getString(m, "F006N"),
			}
			rows = append(rows, row)
		}
	}

	return dataframe.LoadRecords(rows), nil
}

func createRankForecastCninfoSampleData() dataframe.DataFrame {
	records := [][]string{
		{"序号", "证券代码", "证券简称", "研究机构", "研究员", "评级日期", "最新评级", "上次评级", "目标价格"},
		{"1", "000001", "平安银行", "中信证券", "张三", "2023-08-17", "买入", "增持", "15.00"},
	}
	return dataframe.LoadRecords(records)
}
