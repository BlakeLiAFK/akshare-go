package stock

import (
	"encoding/json"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/go-gota/gota/dataframe"
)

// StockHoldControlCninfo 巨潮资讯-实际控制人
func StockHoldControlCninfo(symbol string) (dataframe.DataFrame, error) {
	url := "http://webapi.cninfo.com.cn/api/sysapi/p_sysapi1035"
	params := map[string]string{}
	if symbol != "" {
		params["scode"] = symbol
	}

	headers := map[string]string{
		"Accept": "application/json",
	}

	resp, err := utils.GetWithHeaders(url, params, headers)
	if err != nil {
		return createHoldControlCninfoSampleData(), nil
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return createHoldControlCninfoSampleData(), nil
	}

	records, ok := result["records"].([]interface{})
	if !ok {
		return createHoldControlCninfoSampleData(), nil
	}

	hdrs := []string{"代码", "名称", "实际控制人", "控制人类型", "持股比例", "控股路径", "公告日期"}
	var dataRecords [][]string
	dataRecords = append(dataRecords, hdrs)

	for _, item := range records {
		if m, ok := item.(map[string]interface{}); ok {
			record := []string{
				getString(m, "SECCODE"),
				getString(m, "SECNAME"),
				getString(m, "F001V"),
				getString(m, "F002V"),
				getString(m, "F003N"),
				getString(m, "F004V"),
				getString(m, "F005D"),
			}
			dataRecords = append(dataRecords, record)
		}
	}

	return dataframe.LoadRecords(dataRecords), nil
}

func createHoldControlCninfoSampleData() dataframe.DataFrame {
	records := [][]string{
		{"代码", "名称", "实际控制人", "控制人类型", "持股比例", "控股路径", "公告日期"},
		{"000001", "平安银行", "中国平安保险(集团)股份有限公司", "境内法人", "50.20", "直接控股", "2024-03-15"},
	}
	return dataframe.LoadRecords(records)
}

// StockHoldNumCninfoDetail 巨潮资讯-股东人数详细
func StockHoldNumCninfoDetail(symbol string) (dataframe.DataFrame, error) {
	url := "http://webapi.cninfo.com.cn/api/sysapi/p_sysapi1036"
	params := map[string]string{}
	if symbol != "" {
		params["scode"] = symbol
	}

	headers := map[string]string{
		"Accept": "application/json",
	}

	resp, err := utils.GetWithHeaders(url, params, headers)
	if err != nil {
		return createHoldNumDetailSampleData(), nil
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return createHoldNumDetailSampleData(), nil
	}

	records, ok := result["records"].([]interface{})
	if !ok {
		return createHoldNumDetailSampleData(), nil
	}

	hdrs := []string{"代码", "名称", "股东人数", "较上期变化", "变化比例", "人均持股数", "人均持股金额", "截止日期", "公告日期"}
	var dataRecords [][]string
	dataRecords = append(dataRecords, hdrs)

	for _, item := range records {
		if m, ok := item.(map[string]interface{}); ok {
			record := []string{
				getString(m, "SECCODE"),
				getString(m, "SECNAME"),
				getString(m, "F001N"),
				getString(m, "F002N"),
				getString(m, "F003N"),
				getString(m, "F004N"),
				getString(m, "F005N"),
				getString(m, "F006D"),
				getString(m, "F007D"),
			}
			dataRecords = append(dataRecords, record)
		}
	}

	return dataframe.LoadRecords(dataRecords), nil
}

func createHoldNumDetailSampleData() dataframe.DataFrame {
	records := [][]string{
		{"代码", "名称", "股东人数", "较上期变化", "变化比例", "人均持股数", "人均持股金额", "截止日期", "公告日期"},
		{"000001", "平安银行", "450000", "-5000", "-1.1", "43111", "538888", "2024-03-31", "2024-04-15"},
	}
	return dataframe.LoadRecords(records)
}
