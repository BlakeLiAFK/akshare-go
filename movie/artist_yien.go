package movie

import (
	"fmt"
	"time"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/tidwall/gjson"
)

// BusinessValueArtistItem 艺人商业价值数据结构
type BusinessValueArtistItem struct {
	Rank          int     `json:"rank"`           // 排名
	Artist        string  `json:"artist"`         // 艺人
	BusinessValue float64 `json:"business_value"` // 商业价值
	ProHeat       float64 `json:"pro_heat"`       // 专业热度
	FocusHeat     float64 `json:"focus_heat"`     // 关注热度
	PredictHeat   float64 `json:"predict_heat"`   // 预测热度
	Reputation    float64 `json:"reputation"`     // 美誉度
	StatDate      string  `json:"stat_date"`      // 统计日期
}

// OnlineValueArtistItem 艺人流量价值数据结构
type OnlineValueArtistItem struct {
	Rank        int     `json:"rank"`         // 排名
	Artist      string  `json:"artist"`       // 艺人
	FlowValue   float64 `json:"flow_value"`   // 流量价值
	ProHeat     float64 `json:"pro_heat"`     // 专业热度
	FocusHeat   float64 `json:"focus_heat"`   // 关注热度
	PredictHeat float64 `json:"predict_heat"` // 预测热度
	SalesPower  float64 `json:"sales_power"`  // 带货力
	StatDate    string  `json:"stat_date"`    // 统计日期
}

// BusinessValueArtist 获取艺恩-艺人-艺人商业价值
//
// 目标地址: https://www.endata.com.cn/Marketing/Artist/business.html
//
// 返回:
//   - []BusinessValueArtistItem: 艺人商业价值数据列表
//   - error: 错误信息
//
// 注意: 此接口需要解密
func BusinessValueArtist() ([]BusinessValueArtistItem, error) {
	params := map[string]string{
		"Order":      "BusinessValueIndex_L1",
		"OrderType":  "DESC",
		"PageIndex":  "1",
		"PageSize":   "100",
		"MethodName": "Data_GetList_Star",
	}

	headers := map[string]string{
		"User-Agent":   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
		"Content-Type": "application/x-www-form-urlencoded",
	}

	resp, err := utils.PostFormWithHeaders(yienAPIURL, params, headers)
	if err != nil {
		return nil, fmt.Errorf("获取艺人商业价值失败: %w", err)
	}

	text := resp.String()
	result := gjson.Parse(text)
	data := result.Get("Data.Table")

	var items []BusinessValueArtistItem
	statDate := time.Now().Format("2006-01-02")

	if data.Exists() && data.IsArray() {
		data.ForEach(func(_, value gjson.Result) bool {
			item := BusinessValueArtistItem{
				Rank:          int(value.Get("Irank").Int()),
				Artist:        value.Get("StarName").String(),
				BusinessValue: value.Get("BusinessValueIndex_L1").Float(),
				ProHeat:       value.Get("ProHeatIndex").Float(),
				FocusHeat:     value.Get("FocusHeatIndex").Float(),
				PredictHeat:   value.Get("PredictHeatIndex").Float(),
				Reputation:    value.Get("ReputationIndex").Float(),
				StatDate:      statDate,
			}
			items = append(items, item)
			return true
		})
	}

	if len(items) == 0 {
		return nil, fmt.Errorf("数据需要解密，请参考Python版本使用jm.js解密")
	}

	return items, nil
}

// OnlineValueArtist 获取艺恩-艺人-艺人流量价值
//
// 目标地址: https://www.endata.com.cn/Marketing/Artist/business.html
//
// 返回:
//   - []OnlineValueArtistItem: 艺人流量价值数据列表
//   - error: 错误信息
//
// 注意: 此接口需要解密
func OnlineValueArtist() ([]OnlineValueArtistItem, error) {
	params := map[string]string{
		"Order":      "FlowValueIndex_L1",
		"OrderType":  "DESC",
		"PageIndex":  "1",
		"PageSize":   "100",
		"MethodName": "Data_GetList_Star",
	}

	headers := map[string]string{
		"User-Agent":   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
		"Content-Type": "application/x-www-form-urlencoded",
	}

	resp, err := utils.PostFormWithHeaders(yienAPIURL, params, headers)
	if err != nil {
		return nil, fmt.Errorf("获取艺人流量价值失败: %w", err)
	}

	text := resp.String()
	result := gjson.Parse(text)
	data := result.Get("Data.Table")

	var items []OnlineValueArtistItem
	statDate := time.Now().Format("2006-01-02")

	if data.Exists() && data.IsArray() {
		data.ForEach(func(_, value gjson.Result) bool {
			item := OnlineValueArtistItem{
				Rank:        int(value.Get("Irank").Int()),
				Artist:      value.Get("StarName").String(),
				FlowValue:   value.Get("FlowValueIndex_L1").Float(),
				ProHeat:     value.Get("ProHeatIndex").Float(),
				FocusHeat:   value.Get("FocusHeatIndex").Float(),
				PredictHeat: value.Get("PredictHeatIndex").Float(),
				SalesPower:  value.Get("SalesPowerIndex").Float(),
				StatDate:    statDate,
			}
			items = append(items, item)
			return true
		})
	}

	if len(items) == 0 {
		return nil, fmt.Errorf("数据需要解密，请参考Python版本使用jm.js解密")
	}

	return items, nil
}
