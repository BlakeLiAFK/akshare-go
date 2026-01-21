package option

import (
	"fmt"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/tidwall/gjson"
)

// OptionLhbEm 东方财富网-数据中心-期货期权-期权龙虎榜单
// https://data.eastmoney.com/other/qqlhb.html
// symbol: 期权代码，如 "510050", "510300", "159919"
// indicator: 指标，如 "期权交易情况-认沽交易量", "期权持仓情况-认沽持仓量", "期权交易情况-认购交易量", "期权持仓情况-认购持仓量"
// tradeDate: 交易日期，格式 "20220121"
func OptionLhbEm(symbol, indicator, tradeDate string) ([]OptionLhbEmItem, error) {
	url := "https://datacenter-web.eastmoney.com/api/data/get"

	dateFormatted := fmt.Sprintf("%s-%s-%s", tradeDate[:4], tradeDate[4:6], tradeDate[6:])

	params := map[string]string{
		"type":   "RPT_IF_BILLBOARD_TD",
		"sty":    "ALL",
		"filter": fmt.Sprintf(`(SECURITY_CODE="%s")(TRADE_DATE='%s')`, symbol, dateFormatted),
		"p":      "1",
		"pss":    "200",
		"source": "IFBILLBOARD",
		"client": "WEB",
		"ut":     "b2884a393a59ad64002292a3e90d46a5",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	text := resp.String()
	dataArr := gjson.Get(text, "result.data")

	var result []OptionLhbEmItem

	// 根据 indicator 选择不同的数据范围
	startIdx := 0
	endIdx := 7

	switch indicator {
	case "期权交易情况-认沽交易量":
		startIdx, endIdx = 0, 7
	case "期权持仓情况-认沽持仓量":
		startIdx, endIdx = 7, 14
	case "期权交易情况-认购交易量":
		startIdx, endIdx = 14, 21
	case "期权持仓情况-认购持仓量":
		startIdx, endIdx = 21, 28
	}

	items := dataArr.Array()
	for i := startIdx; i < endIdx && i < len(items); i++ {
		value := items[i]

		var volume, change, netVolume int64
		var ratio float64

		switch indicator {
		case "期权交易情况-认沽交易量":
			volume = value.Get("SELL_VOLUME").Int()
			change = value.Get("SELL_VOLUME_CHANGE").Int()
			netVolume = value.Get("NET_SELL_VOLUME").Int()
			ratio = value.Get("SELL_VOLUME_RATIO").Float()
		case "期权持仓情况-认沽持仓量":
			volume = value.Get("SELL_POSITION").Int()
			change = value.Get("SELL_POSITION_CHANGE").Int()
			netVolume = value.Get("NET_SELL_POSITION").Int()
			ratio = value.Get("SELL_POSITION_RATIO").Float()
		case "期权交易情况-认购交易量":
			volume = value.Get("BUY_VOLUME").Int()
			change = value.Get("BUY_VOLUME_CHANGE").Int()
			netVolume = value.Get("NET_BUY_VOLUME").Int()
			ratio = value.Get("BUY_VOLUME_RATIO").Float()
		case "期权持仓情况-认购持仓量":
			volume = value.Get("BUY_POSITION").Int()
			change = value.Get("BUY_POSITION_CHANGE").Int()
			netVolume = value.Get("NET_BUY_POSITION").Int()
			ratio = value.Get("BUY_POSITION_RATIO").Float()
		}

		result = append(result, OptionLhbEmItem{
			TradeType:    value.Get("TRADE_TYPE").String(),
			TradeDate:    value.Get("TRADE_DATE").String(),
			SecurityCode: value.Get("SECURITY_CODE").String(),
			TargetName:   value.Get("TARGET_NAME").String(),
			Rank:         int(value.Get("MEMBER_RANK").Int()),
			Member:       value.Get("MEMBER_NAME_ABBR").String(),
			Volume:       volume,
			Change:       change,
			NetVolume:    netVolume,
			Ratio:        ratio,
		})
	}

	return result, nil
}
