package option

import (
	"fmt"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/tidwall/gjson"
)

// OptionRiskIndicatorSse 上海证券交易所-产品-股票期权-期权风险指标
// http://www.sse.com.cn/assortment/options/risk/
// date: 日期，格式 "20240626"，20150209 开始
func OptionRiskIndicatorSse(date string) ([]OptionRiskIndicatorSseItem, error) {
	url := "http://query.sse.com.cn/commonQuery.do"

	params := map[string]string{
		"isPagination":   "false",
		"trade_date":     date,
		"sqlId":          "SSE_ZQPZ_YSP_GGQQZSXT_YSHQ_QQFXZB_DATE_L",
		"contractSymbol": "",
	}

	headers := map[string]string{
		"Accept":          "*/*",
		"Accept-Encoding": "gzip, deflate",
		"Accept-Language": "zh-CN,zh;q=0.9,en;q=0.8",
		"Cache-Control":   "no-cache",
		"Connection":      "keep-alive",
		"Host":            "query.sse.com.cn",
		"Pragma":          "no-cache",
		"Referer":         "http://www.sse.com.cn/",
	}

	resp, err := utils.GetWithHeaders(url, params, headers)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	text := resp.String()
	resultArr := gjson.Get(text, "result")

	var result []OptionRiskIndicatorSseItem
	resultArr.ForEach(func(key, value gjson.Result) bool {
		result = append(result, OptionRiskIndicatorSseItem{
			TradeDate:      value.Get("TRADE_DATE").String(),
			SecurityId:     value.Get("SECURITY_ID").String(),
			ContractId:     value.Get("CONTRACT_ID").String(),
			ContractSymbol: value.Get("CONTRACT_SYMBOL").String(),
			Delta:          value.Get("DELTA_VALUE").Float(),
			Theta:          value.Get("THETA_VALUE").Float(),
			Gamma:          value.Get("GAMMA_VALUE").Float(),
			Vega:           value.Get("VEGA_VALUE").Float(),
			Rho:            value.Get("RHO_VALUE").Float(),
			ImpliedVol:     value.Get("IMPLC_VOLATLTY").Float(),
		})
		return true
	})

	return result, nil
}
