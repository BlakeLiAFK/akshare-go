package futures_derivative

import (
	"encoding/json"
	"fmt"

	"github.com/BlakeLiAFK/akshare/utils"
)

// FuturesHogCore 玄田数据-核心数据
// 参数: symbol 品种类型，可选 "外三元", "内三元", "土杂猪"
// 返回: 生猪价格数据
func FuturesHogCore(symbol string) ([]map[string]interface{}, error) {
	if symbol == "" {
		symbol = "外三元"
	}

	url := "https://xt.yangzhu.vip/data/getzhujiahitsdata"
	var ptype string

	switch symbol {
	case "外三元":
		ptype = "1"
	case "内三元":
		ptype = "2"
	case "土杂猪":
		ptype = "3"
	default:
		return nil, fmt.Errorf("请输入正确的 symbol 参数: 外三元, 内三元, 土杂猪")
	}

	params := map[string]string{
		"ptype":    ptype,
		"areano":   "-1",
		"datetype": "0",
	}

	resp, err := utils.Post(url, params)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	var result struct {
		Data [][]interface{} `json:"data"`
	}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return nil, fmt.Errorf("解析JSON失败: %w", err)
	}

	records := make([]map[string]interface{}, 0)
	for _, item := range result.Data {
		if len(item) < 2 {
			continue
		}

		record := map[string]interface{}{
			"date":  fmt.Sprintf("%v", item[1]),
			"value": utils.MustFloat64(fmt.Sprintf("%v", item[0])),
		}
		records = append(records, record)
	}

	return records, nil
}

// FuturesHogCost 玄田数据-成本维度
// 参数: symbol 品种类型，可选 "玉米", "豆粕", "二元母猪价格", "仔猪价格"
// 返回: 生猪成本相关数据
func FuturesHogCost(symbol string) ([]map[string]interface{}, error) {
	if symbol == "" {
		symbol = "玉米"
	}

	var url string
	var ptype string

	switch symbol {
	case "玉米":
		url = "https://xt.yangzhu.vip/data/getzhujiahitsdata"
		ptype = "4"
	case "豆粕":
		url = "https://xt.yangzhu.vip/data/getzhujiahitsdata"
		ptype = "5"
	case "二元母猪价格":
		url = "https://xt.yangzhu.vip/data/getmapdata"
		ptype = "1"
	case "仔猪价格":
		url = "https://xt.yangzhu.vip/data/getmapdata"
		ptype = "2"
	default:
		return nil, fmt.Errorf("请输入正确的 symbol 参数: 玉米, 豆粕, 二元母猪价格, 仔猪价格")
	}

	params := map[string]string{
		"ptype":  ptype,
		"areano": "-1",
	}

	if symbol == "玉米" || symbol == "豆粕" {
		params["datetype"] = "0"
	}

	resp, err := utils.Post(url, params)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	var result struct {
		Data [][]interface{} `json:"data"`
	}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return nil, fmt.Errorf("解析JSON失败: %w", err)
	}

	records := make([]map[string]interface{}, 0)
	for _, item := range result.Data {
		if len(item) < 2 {
			continue
		}

		var record map[string]interface{}
		if symbol == "玉米" || symbol == "豆粕" {
			// getzhujiahitsdata 返回 [value, date]
			record = map[string]interface{}{
				"date":  fmt.Sprintf("%v", item[1]),
				"value": utils.MustFloat64(fmt.Sprintf("%v", item[0])),
			}
		} else {
			// getmapdata 返回 [date, value]
			record = map[string]interface{}{
				"date":  fmt.Sprintf("%v", item[0]),
				"value": utils.MustFloat64(fmt.Sprintf("%v", item[1])),
			}
		}
		records = append(records, record)
	}

	return records, nil
}

// FuturesHogSupply 玄田数据-供应维度
// 参数: symbol 品种类型，可选 "猪肉批发价", "储备冻猪肉", "饲料原料数据", "白条肉",
//
//	"生猪产能", "育肥猪", "肉类价格指数", "猪粮比价"
//
// 返回: 生猪供应相关数据
func FuturesHogSupply(symbol string) ([]map[string]interface{}, error) {
	if symbol == "" {
		symbol = "猪肉批发价"
	}

	url := "https://xt.yangzhu.vip/data/getmapdata"
	var ptype string

	switch symbol {
	case "猪肉批发价":
		ptype = "3"
	case "储备冻猪肉":
		ptype = "4"
	case "饲料原料数据":
		ptype = "5"
	case "白条肉":
		ptype = "6"
	case "生猪产能":
		ptype = "7"
	case "育肥猪":
		ptype = "9"
	case "肉类价格指数":
		ptype = "10"
	case "猪粮比价":
		ptype = "11"
	default:
		return nil, fmt.Errorf("请输入正确的 symbol 参数")
	}

	params := map[string]string{
		"ptype":  ptype,
		"areano": "-1",
	}

	resp, err := utils.Post(url, params)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	var result struct {
		Data [][]interface{} `json:"data"`
	}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return nil, fmt.Errorf("解析JSON失败: %w", err)
	}

	records := make([]map[string]interface{}, 0)

	switch symbol {
	case "猪肉批发价", "肉类价格指数":
		// 返回 [date, item, value]，删除 item 列
		for _, item := range result.Data {
			if len(item) < 3 {
				continue
			}
			record := map[string]interface{}{
				"date":  fmt.Sprintf("%v", item[0]),
				"value": utils.MustFloat64(fmt.Sprintf("%v", item[2])),
			}
			records = append(records, record)
		}
	case "储备冻猪肉", "猪粮比价":
		// 返回 [date, value]
		for _, item := range result.Data {
			if len(item) < 2 {
				continue
			}
			record := map[string]interface{}{
				"date":  fmt.Sprintf("%v", item[0]),
				"value": utils.MustFloat64(fmt.Sprintf("%v", item[1])),
			}
			records = append(records, record)
		}
	case "饲料原料数据":
		// 返回 [周期, 大豆进口金额, 大豆播种面积, 玉米进口金额, 玉米播种面积]
		for _, item := range result.Data {
			if len(item) < 5 {
				continue
			}
			record := map[string]interface{}{
				"周期":     fmt.Sprintf("%v", item[0]),
				"大豆进口金额": utils.MustFloat64(fmt.Sprintf("%v", item[1])),
				"大豆播种面积": utils.MustFloat64(fmt.Sprintf("%v", item[2])),
				"玉米进口金额": utils.MustFloat64(fmt.Sprintf("%v", item[3])),
				"玉米播种面积": utils.MustFloat64(fmt.Sprintf("%v", item[4])),
			}
			records = append(records, record)
		}
	case "白条肉":
		// 返回 [周期, 白条肉平均出厂价格, 环比, 同比]
		for _, item := range result.Data {
			if len(item) < 4 {
				continue
			}
			record := map[string]interface{}{
				"周期":        fmt.Sprintf("%v", item[0]),
				"白条肉平均出厂价格": utils.MustFloat64(fmt.Sprintf("%v", item[1])),
				"环比":        utils.MustFloat64(fmt.Sprintf("%v", item[2])),
				"同比":        utils.MustFloat64(fmt.Sprintf("%v", item[3])),
			}
			records = append(records, record)
		}
	case "生猪产能":
		// 返回 [周期, 能繁母猪存栏, 猪肉产量, 生猪存栏, 生猪出栏]
		for _, item := range result.Data {
			if len(item) < 5 {
				continue
			}
			record := map[string]interface{}{
				"周期":     fmt.Sprintf("%v", item[0]),
				"能繁母猪存栏": utils.MustFloat64(fmt.Sprintf("%v", item[1])),
				"猪肉产量":   utils.MustFloat64(fmt.Sprintf("%v", item[2])),
				"生猪存栏":   utils.MustFloat64(fmt.Sprintf("%v", item[3])),
				"生猪出栏":   utils.MustFloat64(fmt.Sprintf("%v", item[4])),
			}
			records = append(records, record)
		}
	case "育肥猪":
		// 返回 JSON 对象数组，提取 date 和 benzhou 字段
		// 这里需要特殊处理，因为返回格式不同
		respBody := resp.Body()
		var objResult struct {
			Data []map[string]interface{} `json:"data"`
		}
		if err := json.Unmarshal(respBody, &objResult); err != nil {
			return nil, fmt.Errorf("解析JSON失败: %w", err)
		}
		for _, item := range objResult.Data {
			date, _ := item["date"]
			benzhou, _ := item["benzhou"]
			record := map[string]interface{}{
				"date":    fmt.Sprintf("%v", date),
				"benzhou": utils.MustFloat64(fmt.Sprintf("%v", benzhou)),
			}
			records = append(records, record)
		}
	}

	return records, nil
}
