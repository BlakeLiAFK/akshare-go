package movie

import (
	"fmt"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/tidwall/gjson"
)

// VideoTvItem 电视剧集数据结构
type VideoTvItem struct {
	Rank       int     `json:"rank"`        // 排序
	Name       string  `json:"name"`        // 名称
	Genre      string  `json:"genre"`       // 类型
	PlayIndex  float64 `json:"play_index"`  // 播映指数
	MediaHeat  float64 `json:"media_heat"`  // 媒体热度
	UserHeat   float64 `json:"user_heat"`   // 用户热度
	Praise     float64 `json:"praise"`      // 好评度
	ViewDegree float64 `json:"view_degree"` // 观看度
	StatDate   string  `json:"stat_date"`   // 统计日期
}

// VideoVarietyShowItem 综艺节目数据结构
type VideoVarietyShowItem struct {
	Rank       int     `json:"rank"`        // 排序
	Name       string  `json:"name"`        // 名称
	Genre      string  `json:"genre"`       // 类型
	PlayIndex  float64 `json:"play_index"`  // 播映指数
	MediaHeat  float64 `json:"media_heat"`  // 媒体热度
	UserHeat   float64 `json:"user_heat"`   // 用户热度
	Praise     float64 `json:"praise"`      // 好评度
	ViewDegree float64 `json:"view_degree"` // 观看度
	StatDate   string  `json:"stat_date"`   // 统计日期
}

// VideoTv 获取艺恩-视频放映-电视剧集
//
// 目标地址: https://www.endata.com.cn/Video/index.html
//
// 返回:
//   - []VideoTvItem: 电视剧集数据列表
//   - error: 错误信息
//
// 注意: 此接口需要解密
func VideoTv() ([]VideoTvItem, error) {
	params := map[string]string{
		"tvType":     "2",
		"MethodName": "BoxOffice_GetTvData_PlayIndexRank",
	}

	headers := map[string]string{
		"User-Agent":   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
		"Content-Type": "application/x-www-form-urlencoded",
	}

	resp, err := utils.PostFormWithHeaders(yienAPIURL, params, headers)
	if err != nil {
		return nil, fmt.Errorf("获取电视剧集数据失败: %w", err)
	}

	text := resp.String()
	result := gjson.Parse(text)
	data := result.Get("Data.Table")
	reportDate := result.Get("Data.Table1.0.MaxDate").String()

	var items []VideoTvItem
	if data.Exists() && data.IsArray() {
		data.ForEach(func(_, value gjson.Result) bool {
			item := VideoTvItem{
				Rank:       int(value.Get("Irank").Int()),
				Name:       value.Get("TvName").String(),
				Genre:      value.Get("TvType").String(),
				PlayIndex:  value.Get("PlayIndex").Float(),
				MediaHeat:  value.Get("MediaHeat").Float(),
				UserHeat:   value.Get("UserHeat").Float(),
				Praise:     value.Get("Praise").Float(),
				ViewDegree: value.Get("ViewDegree").Float(),
				StatDate:   reportDate,
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

// VideoVarietyShow 获取艺恩-视频放映-综艺节目
//
// 目标地址: https://www.endata.com.cn/Video/index.html
//
// 返回:
//   - []VideoVarietyShowItem: 综艺节目数据列表
//   - error: 错误信息
//
// 注意: 此接口需要解密
func VideoVarietyShow() ([]VideoVarietyShowItem, error) {
	params := map[string]string{
		"tvType":     "8",
		"MethodName": "BoxOffice_GetTvData_PlayIndexRank",
	}

	headers := map[string]string{
		"User-Agent":   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
		"Content-Type": "application/x-www-form-urlencoded",
	}

	resp, err := utils.PostFormWithHeaders(yienAPIURL, params, headers)
	if err != nil {
		return nil, fmt.Errorf("获取综艺节目数据失败: %w", err)
	}

	text := resp.String()
	result := gjson.Parse(text)
	data := result.Get("Data.Table")
	reportDate := result.Get("Data.Table1.0.MaxDate").String()

	var items []VideoVarietyShowItem
	if data.Exists() && data.IsArray() {
		data.ForEach(func(_, value gjson.Result) bool {
			item := VideoVarietyShowItem{
				Rank:       int(value.Get("Irank").Int()),
				Name:       value.Get("TvName").String(),
				Genre:      value.Get("TvType").String(),
				PlayIndex:  value.Get("PlayIndex").Float(),
				MediaHeat:  value.Get("MediaHeat").Float(),
				UserHeat:   value.Get("UserHeat").Float(),
				Praise:     value.Get("Praise").Float(),
				ViewDegree: value.Get("ViewDegree").Float(),
				StatDate:   reportDate,
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
