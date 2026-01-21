package bond

import (
	"fmt"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/go-gota/gota/dataframe"
	"github.com/tidwall/gjson"
)

// BondZHUSRate 中美国债收益率
//
// 参数:
//   - startDate: 开始统计时间，格式: "19901219"
//
// 返回:
//   - dataframe.DataFrame: 包含中美国债收益率数据
//   - error: 错误信息
//
// 数据源: https://data.eastmoney.com/cjsj/zmgzsyl.html
func BondZHUSRate(startDate string) (dataframe.DataFrame, error) {
	url := "https://datacenter.eastmoney.com/api/data/get"

	// 第一次请求获取总页数
	params := map[string]string{
		"type":    "RPTA_WEB_TREASURYYIELD",
		"sty":     "ALL",
		"st":      "SOLAR_DATE",
		"sr":      "-1",
		"token":   "894050c76af8597a853f5b408b759f5d",
		"p":       "1",
		"ps":      "500",
		"pageNo":  "1",
		"pageNum": "1",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	totalPages := int(gjson.Get(resp.String(), "result.pages").Int())
	if totalPages == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到数据")
	}

	var allRecords [][]string
	var headers []string

	// 遍历所有页面
	for page := 1; page <= totalPages; page++ {
		params["p"] = fmt.Sprintf("%d", page)
		params["pageNo"] = fmt.Sprintf("%d", page)
		params["pageNum"] = fmt.Sprintf("%d", page)

		resp, err := utils.Get(url, params)
		if err != nil {
			return dataframe.DataFrame{}, fmt.Errorf("请求第%d页失败: %w", page, err)
		}

		// 解析数据
		dataArray := gjson.Get(resp.String(), "result.data").Array()
		if len(dataArray) == 0 {
			continue
		}

		// 第一页时提取列头
		if page == 1 {
			firstItem := dataArray[0].Map()
			for key := range firstItem {
				headers = append(headers, key)
			}
			allRecords = append(allRecords, headers)
		}

		// 提取数据行
		for _, item := range dataArray {
			var row []string
			itemMap := item.Map()
			for _, header := range headers {
				value := itemMap[header].String()
				row = append(row, value)
			}
			allRecords = append(allRecords, row)
		}
	}

	if len(allRecords) <= 1 {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到有效数据")
	}

	df := dataframe.LoadRecords(allRecords)
	return df, nil
}
