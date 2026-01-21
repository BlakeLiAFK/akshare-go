package fund

import (
	"fmt"
	"strings"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/tidwall/gjson"
)

// FundNewFoundEm 获取基金数据-新发基金-新成立基金
// https://fund.eastmoney.com/data/xinfound.html
// 返回:
//
//	新成立基金数据
func FundNewFoundEm() ([]map[string]interface{}, error) {
	url := "https://fund.eastmoney.com/data/FundNewIssue.aspx"
	params := map[string]string{
		"t":     "xcln",
		"sort":  "jzrgq,desc",
		"y":     "",
		"page":  "1,50000",
		"isbuy": "1",
		"v":     "0.4069919776543214",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	// 处理 demjson 格式: var newfunddata={...}
	text := resp.String()
	text = strings.TrimPrefix(text, "var newfunddata=")
	text = strings.TrimSpace(text)

	result := gjson.Parse(text)
	datas := result.Get("datas").Array()

	records := make([]map[string]interface{}, 0, len(datas))
	for _, item := range datas {
		arr := strings.Split(item.String(), ",")
		if len(arr) < 19 {
			continue
		}

		// 处理费率（去除%符号）
		feeRate := strings.TrimSuffix(strings.TrimSpace(arr[18]), "%")

		record := map[string]interface{}{
			"基金代码":  arr[0],
			"基金简称":  arr[1],
			"发行公司":  arr[2],
			"基金类型":  arr[4],
			"集中认购期": arr[10],
			"募集份额":  utils.MustFloat64(arr[5]),
			"成立日期":  arr[6],
			"成立来涨幅": utils.MustFloat64(strings.ReplaceAll(arr[7], ",", "")),
			"基金经理":  arr[8],
			"申购状态":  arr[9],
			"优惠费率":  utils.MustFloat64(feeRate),
		}
		records = append(records, record)
	}

	return records, nil
}
