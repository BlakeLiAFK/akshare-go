package stock_feature

import (
	"fmt"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/tidwall/gjson"
)

// StockZtPoolItem 涨停股池项
type StockZtPoolItem struct {
	Index         int     `json:"index"`           // 序号
	Code          string  `json:"code"`            // 代码
	Name          string  `json:"name"`            // 名称
	ChangeRate    float64 `json:"change_rate"`     // 涨跌幅
	LatestPrice   float64 `json:"latest_price"`    // 最新价
	Turnover      float64 `json:"turnover"`        // 成交额
	CirculateMV   float64 `json:"circulate_mv"`    // 流通市值
	TotalMV       float64 `json:"total_mv"`        // 总市值
	TurnoverRate  float64 `json:"turnover_rate"`   // 换手率
	SealAmount    float64 `json:"seal_amount"`     // 封板资金
	FirstSealTime string  `json:"first_seal_time"` // 首次封板时间
	LastSealTime  string  `json:"last_seal_time"`  // 最后封板时间
	FailCount     int     `json:"fail_count"`      // 炸板次数
	ZtStatistics  string  `json:"zt_statistics"`   // 涨停统计
	ContinuousNum int     `json:"continuous_num"`  // 连板数
	Industry      string  `json:"industry"`        // 所属行业
}

// StockZtPoolPreviousItem 昨日涨停股池项
type StockZtPoolPreviousItem struct {
	Index        int     `json:"index"`         // 序号
	Code         string  `json:"code"`          // 代码
	Name         string  `json:"name"`          // 名称
	LatestPrice  float64 `json:"latest_price"`  // 最新价
	ZtPrice      float64 `json:"zt_price"`      // 涨停价
	ChangeRate   float64 `json:"change_rate"`   // 涨跌幅
	Turnover     float64 `json:"turnover"`      // 成交额
	CirculateMV  float64 `json:"circulate_mv"`  // 流通市值
	TotalMV      float64 `json:"total_mv"`      // 总市值
	TurnoverRate float64 `json:"turnover_rate"` // 换手率
	Amplitude    float64 `json:"amplitude"`     // 振幅
	Industry     string  `json:"industry"`      // 所属行业
}

// StockDtPoolItem 跌停股池项
type StockDtPoolItem struct {
	Index         int     `json:"index"`           // 序号
	Code          string  `json:"code"`            // 代码
	Name          string  `json:"name"`            // 名称
	ChangeRate    float64 `json:"change_rate"`     // 涨跌幅
	LatestPrice   float64 `json:"latest_price"`    // 最新价
	Turnover      float64 `json:"turnover"`        // 成交额
	CirculateMV   float64 `json:"circulate_mv"`    // 流通市值
	TotalMV       float64 `json:"total_mv"`        // 总市值
	TurnoverRate  float64 `json:"turnover_rate"`   // 换手率
	SealAmount    float64 `json:"seal_amount"`     // 封板资金
	FirstSealTime string  `json:"first_seal_time"` // 首次封板时间
	LastSealTime  string  `json:"last_seal_time"`  // 最后封板时间
	OpenCount     int     `json:"open_count"`      // 开板次数
	ContinuousNum int     `json:"continuous_num"`  // 连板数
	Industry      string  `json:"industry"`        // 所属行业
}

// StockZtPoolEm 东方财富网-行情中心-涨停板行情-涨停股池
// date: 交易日，格式 "20241008"
func StockZtPoolEm(date string) ([]StockZtPoolItem, error) {
	url := "https://push2ex.eastmoney.com/getTopicZTPool"
	params := map[string]string{
		"ut":        "7eea3edcaed734bea9cbfc24409ed989",
		"dpt":       "wz.ztzt",
		"Pageindex": "0",
		"pagesize":  "10000",
		"sort":      "fbt:asc",
		"date":      date,
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return nil, fmt.Errorf("获取涨停股池失败: %w", err)
	}

	result := gjson.ParseBytes(resp.Body())
	if !result.Get("data").Exists() || result.Get("data").String() == "null" {
		return nil, nil
	}

	pool := result.Get("data.pool").Array()
	if len(pool) == 0 {
		return nil, nil
	}

	var items []StockZtPoolItem
	for i, item := range pool {
		ztStat := item.Get("zttj")
		days := ztStat.Get("days").Int()
		ct := ztStat.Get("ct").Int()

		firstTime := fmt.Sprintf("%06d", item.Get("fbt").Int())
		lastTime := fmt.Sprintf("%06d", item.Get("lbt").Int())

		items = append(items, StockZtPoolItem{
			Index:         i + 1,
			Code:          item.Get("c").String(),
			Name:          item.Get("n").String(),
			ChangeRate:    item.Get("zdp").Float(),
			LatestPrice:   item.Get("p").Float() / 1000,
			Turnover:      item.Get("amount").Float(),
			CirculateMV:   item.Get("ltsz").Float(),
			TotalMV:       item.Get("tshare").Float(),
			TurnoverRate:  item.Get("hs").Float(),
			SealAmount:    item.Get("fund").Float(),
			FirstSealTime: firstTime,
			LastSealTime:  lastTime,
			FailCount:     int(item.Get("zbc").Int()),
			ZtStatistics:  fmt.Sprintf("%d/%d", days, ct),
			ContinuousNum: int(item.Get("lbc").Int()),
			Industry:      item.Get("hybk").String(),
		})
	}

	return items, nil
}

// StockZtPoolPreviousEm 东方财富网-行情中心-涨停板行情-昨日涨停股池
// date: 交易日，格式 "20240415"
func StockZtPoolPreviousEm(date string) ([]StockZtPoolPreviousItem, error) {
	url := "https://push2ex.eastmoney.com/getYesterdayZTPool"
	params := map[string]string{
		"ut":        "7eea3edcaed734bea9cbfc24409ed989",
		"dpt":       "wz.ztzt",
		"Pageindex": "0",
		"pagesize":  "5000",
		"sort":      "zs:desc",
		"date":      date,
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return nil, fmt.Errorf("获取昨日涨停股池失败: %w", err)
	}

	result := gjson.ParseBytes(resp.Body())
	if !result.Get("data").Exists() || result.Get("data").String() == "null" {
		return nil, nil
	}

	pool := result.Get("data.pool").Array()
	if len(pool) == 0 {
		return nil, nil
	}

	var items []StockZtPoolPreviousItem
	for i, item := range pool {
		items = append(items, StockZtPoolPreviousItem{
			Index:        i + 1,
			Code:         item.Get("c").String(),
			Name:         item.Get("n").String(),
			LatestPrice:  item.Get("p").Float() / 1000,
			ZtPrice:      item.Get("ztp").Float() / 1000,
			ChangeRate:   item.Get("zdp").Float(),
			Turnover:     item.Get("amount").Float(),
			CirculateMV:  item.Get("ltsz").Float(),
			TotalMV:      item.Get("tshare").Float(),
			TurnoverRate: item.Get("hs").Float(),
			Amplitude:    item.Get("zf").Float(),
			Industry:     item.Get("hybk").String(),
		})
	}

	return items, nil
}

// StockDtPoolEm 东方财富网-行情中心-涨停板行情-跌停股池
// date: 交易日，格式 "20240415"
func StockDtPoolEm(date string) ([]StockDtPoolItem, error) {
	url := "https://push2ex.eastmoney.com/getTopicDTPool"
	params := map[string]string{
		"ut":        "7eea3edcaed734bea9cbfc24409ed989",
		"dpt":       "wz.ztzt",
		"Pageindex": "0",
		"pagesize":  "10000",
		"sort":      "fund:asc",
		"date":      date,
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return nil, fmt.Errorf("获取跌停股池失败: %w", err)
	}

	result := gjson.ParseBytes(resp.Body())
	if !result.Get("data").Exists() || result.Get("data").String() == "null" {
		return nil, nil
	}

	pool := result.Get("data.pool").Array()
	if len(pool) == 0 {
		return nil, nil
	}

	var items []StockDtPoolItem
	for i, item := range pool {
		firstTime := fmt.Sprintf("%06d", item.Get("fbt").Int())
		lastTime := fmt.Sprintf("%06d", item.Get("lbt").Int())

		items = append(items, StockDtPoolItem{
			Index:         i + 1,
			Code:          item.Get("c").String(),
			Name:          item.Get("n").String(),
			ChangeRate:    item.Get("zdp").Float(),
			LatestPrice:   item.Get("p").Float() / 1000,
			Turnover:      item.Get("amount").Float(),
			CirculateMV:   item.Get("ltsz").Float(),
			TotalMV:       item.Get("tshare").Float(),
			TurnoverRate:  item.Get("hs").Float(),
			SealAmount:    item.Get("fund").Float(),
			FirstSealTime: firstTime,
			LastSealTime:  lastTime,
			OpenCount:     int(item.Get("obc").Int()),
			ContinuousNum: int(item.Get("lbc").Int()),
			Industry:      item.Get("hybk").String(),
		})
	}

	return items, nil
}

// StockZbPoolEm 东方财富网-行情中心-涨停板行情-炸板股池
// date: 交易日，格式 "20240415"
func StockZbPoolEm(date string) ([]map[string]interface{}, error) {
	url := "https://push2ex.eastmoney.com/getTopicZBPool"
	params := map[string]string{
		"ut":        "7eea3edcaed734bea9cbfc24409ed989",
		"dpt":       "wz.ztzt",
		"Pageindex": "0",
		"pagesize":  "10000",
		"sort":      "fbt:asc",
		"date":      date,
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return nil, fmt.Errorf("获取炸板股池失败: %w", err)
	}

	result := gjson.ParseBytes(resp.Body())
	if !result.Get("data").Exists() || result.Get("data").String() == "null" {
		return nil, nil
	}

	pool := result.Get("data.pool").Array()
	var items []map[string]interface{}
	for i, item := range pool {
		m := item.Value().(map[string]interface{})
		m["index"] = i + 1
		items = append(items, m)
	}

	return items, nil
}
