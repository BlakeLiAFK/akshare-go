package option

import (
	"fmt"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/tidwall/gjson"
)

// OptionCurrentDaySse 上海证券交易所-产品-股票期权-信息披露-当日合约
// http://www.sse.com.cn/assortment/options/disclo/preinfo/
func OptionCurrentDaySse() ([]OptionCurrentDaySseItem, error) {
	url := "http://query.sse.com.cn/commonQuery.do"

	params := map[string]string{
		"isPagination": "false",
		"expireDate":   "",
		"securityId":   "",
		"sqlId":        "SSE_ZQPZ_YSP_GGQQZSXT_XXPL_DRHY_SEARCH_L",
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

	var result []OptionCurrentDaySseItem
	resultArr.ForEach(func(key, value gjson.Result) bool {
		result = append(result, OptionCurrentDaySseItem{
			SecurityId:     value.Get("SECURITY_ID").String(),
			ContractId:     value.Get("CONTRACT_ID").String(),
			ContractSymbol: value.Get("CONTRACT_SYMBOL").String(),
			SecurityName:   value.Get("SECURITYNAMEBYID").String(),
			CallOrPut:      value.Get("CALL_OR_PUT").String(),
			ExercisePrice:  value.Get("EXERCISE_PRICE").String(),
			ContractUnit:   value.Get("CONTRACT_UNIT").String(),
			EndDate:        value.Get("END_DATE").String(),
			DeliveryDate:   value.Get("DELIVERY_DATE").String(),
			ExpireDate:     value.Get("EXPIRE_DATE").String(),
			StartDate:      value.Get("START_DATE").String(),
		})
		return true
	})

	return result, nil
}
