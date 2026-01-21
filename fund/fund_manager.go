package fund

import (
	"fmt"
	"strings"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/tidwall/gjson"
)

// FundManagerEm 获取天天基金网-基金数据-基金经理大全
// https://fund.eastmoney.com/manager/default.html
// 返回:
//
//	基金经理大全数据
func FundManagerEm() ([]map[string]interface{}, error) {
	url := "https://fund.eastmoney.com/Data/FundDataPortfolio_Interface.aspx"
	params := map[string]string{
		"dt": "14",
		"mc": "returnjson",
		"ft": "all",
		"pn": "500",
		"pi": "1",
		"sc": "abbname",
		"st": "asc",
	}

	// 第一次请求获取总页数
	resp, err := utils.Get(url, params)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	text := resp.String()
	text = strings.TrimPrefix(text, "var returnjson= ")
	text = strings.TrimSpace(text)

	result := gjson.Parse(text)
	totalPages := result.Get("pages").Int()

	allRecords := make([]map[string]interface{}, 0)

	// 遍历所有页面
	for page := int64(1); page <= totalPages; page++ {
		params["pi"] = fmt.Sprintf("%d", page)

		resp, err := utils.Get(url, params)
		if err != nil {
			continue
		}

		text := resp.String()
		text = strings.TrimPrefix(text, "var returnjson= ")
		text = strings.TrimSpace(text)

		pageResult := gjson.Parse(text)
		datas := pageResult.Get("data").Array()

		for _, item := range datas {
			// 提取基金代码和基金名称列表
			fundCodes := strings.Split(item.Get("5").String(), ",")
			fundNames := strings.Split(item.Get("6").String(), ",")

			// 展开每个基金
			for i := 0; i < len(fundCodes) && i < len(fundNames); i++ {
				// 处理最佳回报（去除%）
				bestReturn := strings.TrimSuffix(item.Get("8").String(), "%")
				// 处理资产规模（去除"亿元"）
				assetScale := strings.TrimSuffix(item.Get("11").String(), "亿元")

				record := map[string]interface{}{
					"序号":        len(allRecords) + 1,
					"姓名":        item.Get("2").String(),
					"所属公司":      item.Get("4").String(),
					"现任基金代码":    fundCodes[i],
					"现任基金":      fundNames[i],
					"累计从业时间":    utils.MustFloat64(item.Get("7").Value()),
					"现任基金资产总规模": utils.MustFloat64(assetScale),
					"现任基金最佳回报":  utils.MustFloat64(bestReturn),
				}
				allRecords = append(allRecords, record)
			}
		}
	}

	return allRecords, nil
}
