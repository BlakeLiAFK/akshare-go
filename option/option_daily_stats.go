package option

import (
	"fmt"
	"strings"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/tidwall/gjson"
)

// OptionDailyStatsSse 上海证券交易所-产品-股票期权-每日统计
// https://www.sse.com.cn/assortment/options/date/
// date: 交易日，格式 "20240626"
func OptionDailyStatsSse(date string) ([]OptionDailyStatsSseItem, error) {
	url := "http://query.sse.com.cn/commonQuery.do"

	params := map[string]string{
		"isPagination": "false",
		"sqlId":        "COMMON_SSE_ZQPZ_YSP_QQ_SJTJ_MRTJ_CX",
		"tradeDate":    date,
	}

	headers := map[string]string{
		"Accept":          "*/*",
		"Accept-Encoding": "gzip, deflate",
		"Accept-Language": "zh-CN,zh;q=0.9,en;q=0.8",
		"Cache-Control":   "no-cache",
		"Connection":      "keep-alive",
		"Host":            "query.sse.com.cn",
		"Pragma":          "no-cache",
		"Referer":         "https://www.sse.com.cn/",
	}

	resp, err := utils.GetWithHeaders(url, params, headers)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	text := resp.String()
	resultArr := gjson.Get(text, "result")

	var result []OptionDailyStatsSseItem
	resultArr.ForEach(func(key, value gjson.Result) bool {
		result = append(result, OptionDailyStatsSseItem{
			SecurityCode:     value.Get("SECURITY_CODE").String(),
			SecurityName:     value.Get("SECURITY_ABBR").String(),
			ContractCount:    int(utils.MustParseInt(strings.ReplaceAll(value.Get("CONTRACT_VOLUME").String(), ",", ""))),
			TotalAmount:      utils.MustParseFloat(strings.ReplaceAll(value.Get("TOTAL_MONEY").String(), ",", "")),
			TotalVolume:      utils.MustParseInt(strings.ReplaceAll(value.Get("TOTAL_VOLUME").String(), ",", "")),
			CallVolume:       utils.MustParseInt(strings.ReplaceAll(value.Get("CALL_VOLUME").String(), ",", "")),
			PutVolume:        utils.MustParseInt(strings.ReplaceAll(value.Get("PUT_VOLUME").String(), ",", "")),
			PutCallRatio:     utils.MustParseFloat(strings.ReplaceAll(value.Get("CP_RATE").String(), ",", "")),
			OpenInterest:     utils.MustParseInt(strings.ReplaceAll(value.Get("LEAVES_QTY").String(), ",", "")),
			OpenCallInterest: utils.MustParseInt(strings.ReplaceAll(value.Get("LEAVES_CALL_QTY").String(), ",", "")),
			OpenPutInterest:  utils.MustParseInt(strings.ReplaceAll(value.Get("LEAVES_PUT_QTY").String(), ",", "")),
			TradeDate:        value.Get("TRADE_DATE").String(),
		})
		return true
	})

	return result, nil
}

// OptionDailyStatsSzse 深圳证券交易所-市场数据-期权数据-日度概况
// https://investor.szse.cn/market/option/day/index.html
// date: 交易日，格式 "20240626"
func OptionDailyStatsSzse(date string) ([]OptionDailyStatsSzseItem, error) {
	url := "https://investor.szse.cn/api/report/ShowReport/data"

	dateFormatted := fmt.Sprintf("%s-%s-%s", date[:4], date[4:6], date[6:])

	params := map[string]string{
		"SHOWTYPE":     "JSON",
		"CATALOGID":    "ysprdzb",
		"TABKEY":       "tab1",
		"txtQueryDate": dateFormatted,
		"random":       "0.0652692406565949",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	text := resp.String()
	dataArr := gjson.Get(text, "0.data")

	var result []OptionDailyStatsSzseItem
	dataArr.ForEach(func(key, value gjson.Result) bool {
		result = append(result, OptionDailyStatsSzseItem{
			SecurityCode:     value.Get("bddm").String(),
			SecurityName:     value.Get("bdmc").String(),
			Volume:           utils.MustParseInt(strings.ReplaceAll(value.Get("cjl").String(), ",", "")),
			CallVolume:       utils.MustParseInt(strings.ReplaceAll(value.Get("rccjl").String(), ",", "")),
			PutVolume:        utils.MustParseInt(strings.ReplaceAll(value.Get("rpcjl").String(), ",", "")),
			PutCallRatio:     utils.MustParseFloat(strings.ReplaceAll(value.Get("rcrpccb").String(), ",", "")),
			OpenInterest:     utils.MustParseInt(strings.ReplaceAll(value.Get("wpchyzs").String(), ",", "")),
			OpenCallInterest: utils.MustParseInt(strings.ReplaceAll(value.Get("wpcrchys").String(), ",", "")),
			OpenPutInterest:  utils.MustParseInt(strings.ReplaceAll(value.Get("wpcrphys").String(), ",", "")),
			TradeDate:        dateFormatted,
		})
		return true
	})

	return result, nil
}
