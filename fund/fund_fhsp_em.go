package fund

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/tidwall/gjson"
)

// FundFhEm 获取天天基金网-基金数据-分红送配-基金分红
// https://fund.eastmoney.com/data/fundfenhong.html
func FundFhEm(year string) ([]map[string]interface{}, error) {
	if year == "" {
		year = "2025"
	}

	url := "https://fund.eastmoney.com/Data/funddataIndex_Interface.aspx"
	params := map[string]string{
		"dt":    "8",
		"page":  "1",
		"rank":  "BZDM",
		"sort":  "asc",
		"gs":    "",
		"ftype": "",
		"year":  year,
	}

	// 获取总页数
	resp, err := utils.Get(url, params)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	text := resp.String()
	// 提取总页数: var pages,datas=[总页数,...]
	startIdx := strings.Index(text, "=")
	endIdx := strings.Index(text, ";")
	if startIdx == -1 || endIdx == -1 {
		return nil, fmt.Errorf("解析响应失败")
	}

	jsonText := text[startIdx+1 : endIdx]
	result := gjson.Parse(jsonText)
	totalPage := result.Array()[0].Int()

	records := make([]map[string]interface{}, 0)

	// 遍历所有分页
	for page := int64(1); page <= totalPage; page++ {
		params["page"] = strconv.FormatInt(page, 10)
		resp, err = utils.Get(url, params)
		if err != nil {
			continue
		}

		text = resp.String()
		// 提取数据数组: [[...],[...],...]
		startIdx = strings.Index(text, "[[")
		endIdx = strings.Index(text, ";var jjfh_jjgs")
		if startIdx == -1 || endIdx == -1 {
			continue
		}

		jsonText = text[startIdx:endIdx]
		dataList := gjson.Parse(jsonText).Array()

		for i, item := range dataList {
			values := item.Array()
			if len(values) < 7 {
				continue
			}

			record := map[string]interface{}{
				"序号":    i + 1 + int(page-1)*50,
				"基金代码":  values[0].String(),
				"基金简称":  values[1].String(),
				"权益登记日": values[2].String(),
				"除息日期":  values[3].String(),
				"分红":    utils.MustFloat64(values[4].String()),
				"分红发放日": values[5].String(),
			}
			records = append(records, record)
		}
	}

	return records, nil
}

// FundCfEm 获取天天基金网-基金数据-分红送配-基金拆分
// https://fund.eastmoney.com/data/fundchaifen.html
func FundCfEm(year string) ([]map[string]interface{}, error) {
	if year == "" {
		year = "2025"
	}

	url := "https://fund.eastmoney.com/Data/funddataIndex_Interface.aspx"
	params := map[string]string{
		"dt":    "9",
		"page":  "1",
		"rank":  "FSRQ",
		"sort":  "desc",
		"gs":    "",
		"ftype": "",
		"year":  year,
	}

	// 获取总页数
	resp, err := utils.Get(url, params)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	text := resp.String()
	startIdx := strings.Index(text, "=")
	endIdx := strings.Index(text, ";")
	if startIdx == -1 || endIdx == -1 {
		return nil, fmt.Errorf("解析响应失败")
	}

	jsonText := text[startIdx+1 : endIdx]
	result := gjson.Parse(jsonText)
	totalPage := result.Array()[0].Int()

	records := make([]map[string]interface{}, 0)

	// 遍历所有分页
	for page := int64(1); page <= totalPage; page++ {
		params["page"] = strconv.FormatInt(page, 10)
		resp, err = utils.Get(url, params)
		if err != nil {
			continue
		}

		text = resp.String()
		startIdx = strings.Index(text, "[[")
		endIdx = strings.Index(text, ";var jjcf_jjgs")
		if startIdx == -1 || endIdx == -1 {
			continue
		}

		jsonText = text[startIdx:endIdx]
		dataList := gjson.Parse(jsonText).Array()

		for i, item := range dataList {
			values := item.Array()
			if len(values) < 6 {
				continue
			}

			record := map[string]interface{}{
				"序号":   i + 1 + int(page-1)*50,
				"基金代码": values[0].String(),
				"基金简称": values[1].String(),
				"拆分日":  values[2].String(),
				"拆分折算": utils.MustFloat64(values[3].String()),
				"拆分类型": values[4].String(),
			}
			records = append(records, record)
		}
	}

	return records, nil
}

// FundFhRankEm 获取天天基金网-基金数据-分红送配-基金分红排行
// https://fund.eastmoney.com/data/fundleijifenhong.html
func FundFhRankEm() ([]map[string]interface{}, error) {
	url := "https://fund.eastmoney.com/Data/funddataIndex_Interface.aspx"
	params := map[string]string{
		"dt":    "10",
		"page":  "1",
		"rank":  "FHFCZ",
		"sort":  "desc",
		"gs":    "",
		"ftype": "",
	}

	// 获取总页数
	resp, err := utils.Get(url, params)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	text := resp.String()
	startIdx := strings.Index(text, "=")
	endIdx := strings.Index(text, ";")
	if startIdx == -1 || endIdx == -1 {
		return nil, fmt.Errorf("解析响应失败")
	}

	jsonText := text[startIdx+1 : endIdx]
	result := gjson.Parse(jsonText)
	totalPage := result.Array()[0].Int()

	records := make([]map[string]interface{}, 0)

	// 遍历所有分页
	for page := int64(1); page <= totalPage; page++ {
		params["page"] = strconv.FormatInt(page, 10)
		resp, err = utils.Get(url, params)
		if err != nil {
			continue
		}

		text = resp.String()
		startIdx = strings.Index(text, "[[")
		endIdx = strings.Index(text, ";var fhph_jjgs")
		if startIdx == -1 || endIdx == -1 {
			continue
		}

		jsonText = text[startIdx:endIdx]
		dataList := gjson.Parse(jsonText).Array()

		for i, item := range dataList {
			values := item.Array()
			if len(values) < 6 {
				continue
			}

			record := map[string]interface{}{
				"序号":   i + 1 + int(page-1)*50,
				"基金代码": values[0].String(),
				"基金简称": values[1].String(),
				"累计分红": utils.MustFloat64(values[2].String()),
				"累计次数": utils.MustInt64(values[3].String()),
				"成立日期": values[4].String(),
			}
			records = append(records, record)
		}
	}

	return records, nil
}
