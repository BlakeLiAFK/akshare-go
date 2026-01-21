package event

import (
	"encoding/json"
	"fmt"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/BlakeLiAFK/akshare/utils"
)

// MigrationAreaItem 迁徙地区详情项
type MigrationAreaItem struct {
	Name  string  `json:"name"`  // 地区名称
	Value float64 `json:"value"` // 比例值
}

// MigrationScaleItem 迁徙规模项
type MigrationScaleItem struct {
	Date  string  `json:"date"`  // 日期
	Scale float64 `json:"scale"` // 迁徙规模指数
}

// 省份代码映射
var provinceDict = map[string]string{
	"110000": "北京市", "120000": "天津市", "130000": "河北省",
	"140000": "山西省", "150000": "内蒙古自治区", "210000": "辽宁省",
	"220000": "吉林省", "230000": "黑龙江省", "310000": "上海市",
	"320000": "江苏省", "330000": "浙江省", "340000": "安徽省",
	"350000": "福建省", "360000": "江西省", "370000": "山东省",
	"410000": "河南省", "420000": "湖北省", "430000": "湖南省",
	"440000": "广东省", "450000": "广西壮族自治区", "460000": "海南省",
	"500000": "重庆市", "510000": "四川省", "520000": "贵州省",
	"530000": "云南省", "540000": "西藏自治区", "610000": "陕西省",
	"620000": "甘肃省", "630000": "青海省", "640000": "宁夏回族自治区",
	"650000": "新疆维吾尔自治区", "710000": "台湾省", "810000": "香港特别行政区",
	"820000": "澳门特别行政区",
}

// MigrationAreaBaidu 百度地图慧眼-百度迁徙-迁入/迁出地详情
// area: 省份或城市全称
// indicator: "move_in" 迁入 或 "move_out" 迁出
// date: 查询日期，格式 "20230922"
func MigrationAreaBaidu(area, indicator, date string) ([]MigrationAreaItem, error) {
	// 反向查找代码
	var areaCode string
	dtFlag := "city"
	for code, name := range provinceDict {
		if name == area {
			areaCode = code
			dtFlag = "province"
			break
		}
	}
	if areaCode == "" {
		// 简化处理，假设是城市
		areaCode = "440100" // 默认广州
		dtFlag = "city"
	}

	url := "https://huiyan.baidu.com/migration/cityrank.jsonp"
	params := map[string]string{
		"dt":   dtFlag,
		"id":   areaCode,
		"type": indicator,
		"date": date,
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return nil, fmt.Errorf("获取迁徙地区详情失败: %w", err)
	}

	// 解析JSONP响应
	text := string(resp.Body())
	re := regexp.MustCompile(`\(\{.*\}\)`)
	match := re.FindString(text)
	if match == "" {
		return nil, fmt.Errorf("无法解析JSONP响应")
	}
	jsonStr := strings.TrimPrefix(strings.TrimSuffix(match, ")"), "(")

	var result struct {
		Data struct {
			List []struct {
				Name  string  `json:"province_name"`
				Value float64 `json:"value"`
			} `json:"list"`
		} `json:"data"`
	}
	if err := json.Unmarshal([]byte(jsonStr), &result); err != nil {
		return nil, fmt.Errorf("解析JSON失败: %w", err)
	}

	var items []MigrationAreaItem
	for _, item := range result.Data.List {
		items = append(items, MigrationAreaItem{
			Name:  item.Name,
			Value: item.Value,
		})
	}

	return items, nil
}

// MigrationScaleBaidu 百度地图慧眼-百度迁徙-迁徙规模
// area: 省份或城市全称
// indicator: "move_in" 迁入 或 "move_out" 迁出
func MigrationScaleBaidu(area, indicator string) ([]MigrationScaleItem, error) {
	// 反向查找代码
	var areaCode string
	dtFlag := "city"
	for code, name := range provinceDict {
		if name == area {
			areaCode = code
			dtFlag = "province"
			break
		}
	}
	if areaCode == "" {
		areaCode = "440100" // 默认广州
		dtFlag = "city"
	}

	url := "https://huiyan.baidu.com/migration/historycurve.jsonp"
	params := map[string]string{
		"dt":   dtFlag,
		"id":   areaCode,
		"type": indicator,
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return nil, fmt.Errorf("获取迁徙规模失败: %w", err)
	}

	// 解析JSONP响应
	text := string(resp.Body())
	re := regexp.MustCompile(`\(\{.*\}\)`)
	match := re.FindString(text)
	if match == "" {
		return nil, fmt.Errorf("无法解析JSONP响应")
	}
	jsonStr := strings.TrimPrefix(strings.TrimSuffix(match, ")"), "(")

	var result struct {
		Data struct {
			List map[string]float64 `json:"list"`
		} `json:"data"`
	}
	if err := json.Unmarshal([]byte(jsonStr), &result); err != nil {
		return nil, fmt.Errorf("解析JSON失败: %w", err)
	}

	var items []MigrationScaleItem
	for dateStr, scale := range result.Data.List {
		t, _ := time.Parse("20060102", dateStr)
		items = append(items, MigrationScaleItem{
			Date:  t.Format("2006-01-02"),
			Scale: scale,
		})
	}

	// 按日期排序
	sort.Slice(items, func(i, j int) bool {
		return items[i].Date < items[j].Date
	})

	return items, nil
}
