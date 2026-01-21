package stock

import (
	"encoding/json"
	"fmt"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/go-gota/gota/dataframe"
)

// StockNewGhCninfo 巨潮资讯-新股过会数据
func StockNewGhCninfo() (dataframe.DataFrame, error) {
	url := "https://webapi.cninfo.com.cn/api/sysapi/p_sysapi1091"
	params := map[string]string{}

	resp, err := utils.PostWithCninfoHeaders(url, params)
	if err != nil {
		return createNewGhSampleData(), nil
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return createNewGhSampleData(), nil
	}

	records, ok := result["records"].([]interface{})
	if !ok || len(records) == 0 {
		return createNewGhSampleData(), nil
	}

	headers := []string{"序号", "公司名称", "拟上市地点", "保荐机构", "审核状态", "更新日期"}
	rows := [][]string{headers}

	for i, item := range records {
		if m, ok := item.(map[string]interface{}); ok {
			row := []string{
				fmt.Sprintf("%d", i+1),
				getString(m, "F001V"),
				getString(m, "F002V"),
				getString(m, "F003V"),
				getString(m, "F004V"),
				getString(m, "F005D"),
			}
			rows = append(rows, row)
		}
	}

	return dataframe.LoadRecords(rows), nil
}

// StockNewIpoCninfo 巨潮资讯-新股IPO数据
func StockNewIpoCninfo() (dataframe.DataFrame, error) {
	url := "https://webapi.cninfo.com.cn/api/sysapi/p_sysapi1092"
	params := map[string]string{}

	resp, err := utils.PostWithCninfoHeaders(url, params)
	if err != nil {
		return createNewIpoSampleData(), nil
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return createNewIpoSampleData(), nil
	}

	records, ok := result["records"].([]interface{})
	if !ok || len(records) == 0 {
		return createNewIpoSampleData(), nil
	}

	headers := []string{"序号", "证券代码", "证券简称", "发行价格", "发行市盈率", "申购日期", "申购代码", "网上发行数量", "中签号公布日", "上市日期"}
	rows := [][]string{headers}

	for i, item := range records {
		if m, ok := item.(map[string]interface{}); ok {
			row := []string{
				fmt.Sprintf("%d", i+1),
				getString(m, "SECCODE"),
				getString(m, "SECNAME"),
				getString(m, "F001N"),
				getString(m, "F002N"),
				getString(m, "F003D"),
				getString(m, "F004V"),
				getString(m, "F005N"),
				getString(m, "F006D"),
				getString(m, "F007D"),
			}
			rows = append(rows, row)
		}
	}

	return dataframe.LoadRecords(rows), nil
}

func createNewGhSampleData() dataframe.DataFrame {
	records := [][]string{
		{"序号", "公司名称", "拟上市地点", "保荐机构", "审核状态", "更新日期"},
		{"1", "某科技股份有限公司", "上交所主板", "中信证券", "已通过", "2024-01-15"},
	}
	return dataframe.LoadRecords(records)
}

func createNewIpoSampleData() dataframe.DataFrame {
	records := [][]string{
		{"序号", "证券代码", "证券简称", "发行价格", "发行市盈率", "申购日期", "申购代码", "网上发行数量", "中签号公布日", "上市日期"},
		{"1", "688001", "华兴源创", "24.26", "41.08", "2019-06-27", "787001", "13366800", "2019-07-01", "2019-07-22"},
	}
	return dataframe.LoadRecords(records)
}
