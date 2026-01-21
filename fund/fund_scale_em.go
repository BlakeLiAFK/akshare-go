package fund

import (
	"fmt"
	"strings"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/tidwall/gjson"
)

// FundScaleChangeEm 获取基金规模变动数据
// https://fund.eastmoney.com/data/gmbdlist.html
func FundScaleChangeEm(symbol string) ([]map[string]interface{}, error) {
	url := "https://fund.eastmoney.com/data/FundDataPortfolio_Interface.aspx"
	params := map[string]string{
		"dt": "9",
		"pi": "1",
		"pn": "50",
		"mc": "hypzDetail",
		"st": "desc",
		"sc": "reportdate",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	// 提取JSON部分（从第一个{到倒数第二个字符）
	text := resp.String()
	startIdx := strings.Index(text, "{")
	if startIdx == -1 {
		return nil, fmt.Errorf("解析响应失败")
	}
	jsonText := text[startIdx : len(text)-1]

	result := gjson.Parse(jsonText)
	totalPages := result.Get("pages").Int()

	records := make([]map[string]interface{}, 0)

	// 遍历所有分页
	for page := int64(1); page <= totalPages; page++ {
		if page > 1 {
			params["pi"] = fmt.Sprintf("%d", page)
			resp, err = utils.Get(url, params)
			if err != nil {
				continue
			}

			text = resp.String()
			startIdx = strings.Index(text, "{")
			if startIdx == -1 {
				continue
			}
			jsonText = text[startIdx : len(text)-1]
			result = gjson.Parse(jsonText)
		}

		dataList := result.Get("data").Array()
		for i, item := range dataList {
			record := map[string]interface{}{
				"序号":       i + 1 + int(page-1)*50,
				"基金代码":     item.Get("FCODE").String(),
				"基金简称":     item.Get("SHORTNAME").String(),
				"报告日期":     item.Get("REPORTDATE").String(),
				"期初份额":     utils.MustFloat64(item.Get("SHARESBEGIN").String()),
				"期末份额":     utils.MustFloat64(item.Get("SHARESEND").String()),
				"份额变化":     utils.MustFloat64(item.Get("SHARESCHANGE").String()),
				"期初基金资产净值": utils.MustFloat64(item.Get("NVBEGIN").String()),
				"期末基金资产净值": utils.MustFloat64(item.Get("NVEND").String()),
				"净值变化":     utils.MustFloat64(item.Get("NVCHANGE").String()),
			}
			records = append(records, record)
		}
	}

	return records, nil
}

// FundHoldStructureEm 获取基金持有人结构数据
// https://fund.eastmoney.com/data/cyrjglist.html
func FundHoldStructureEm(symbol string) ([]map[string]interface{}, error) {
	url := "https://fund.eastmoney.com/data/FundDataPortfolio_Interface.aspx"
	params := map[string]string{
		"dt": "11",
		"pi": "1",
		"pn": "50",
		"mc": "hypzDetail",
		"st": "desc",
		"sc": "reportdate",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	// 提取JSON部分
	text := resp.String()
	startIdx := strings.Index(text, "{")
	if startIdx == -1 {
		return nil, fmt.Errorf("解析响应失败")
	}
	jsonText := text[startIdx : len(text)-1]

	result := gjson.Parse(jsonText)
	totalPages := result.Get("pages").Int()

	records := make([]map[string]interface{}, 0)

	// 遍历所有分页
	for page := int64(1); page <= totalPages; page++ {
		if page > 1 {
			params["pi"] = fmt.Sprintf("%d", page)
			resp, err = utils.Get(url, params)
			if err != nil {
				continue
			}

			text = resp.String()
			startIdx = strings.Index(text, "{")
			if startIdx == -1 {
				continue
			}
			jsonText = text[startIdx : len(text)-1]
			result = gjson.Parse(jsonText)
		}

		dataList := result.Get("data").Array()
		for i, item := range dataList {
			record := map[string]interface{}{
				"序号":        i + 1 + int(page-1)*50,
				"基金代码":      item.Get("FCODE").String(),
				"基金简称":      item.Get("SHORTNAME").String(),
				"报告日期":      item.Get("REPORTDATE").String(),
				"机构投资者持有份额": utils.MustFloat64(item.Get("JGCGBL").String()),
				"个人投资者持有份额": utils.MustFloat64(item.Get("GRCGBL").String()),
				"内部持有份额":    utils.MustFloat64(item.Get("NBCGBL").String()),
				"总份额":       utils.MustFloat64(item.Get("ENDSHARES").String()),
			}
			records = append(records, record)
		}
	}

	return records, nil
}
