package fund

import (
	"fmt"
	"strings"
	"time"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/tidwall/gjson"
)

// FundEtfSpotThs 获取同花顺理财-基金数据-每日净值-ETF-实时行情
// https://fund.10jqka.com.cn/datacenter/jz/kfs/etf/
// 参数:
//
//	date: 查询日期，格式 "20240620"，空字符串表示最新
//
// 返回:
//
//	ETF 实时行情数据
func FundEtfSpotThs(date string) ([]map[string]interface{}, error) {
	// 日期格式转换: 20240620 -> 2024-06-20
	innerDate := "0"
	if date != "" && len(date) == 8 {
		innerDate = fmt.Sprintf("%s-%s-%s", date[:4], date[4:6], date[6:])
	}

	url := fmt.Sprintf("https://fund.10jqka.com.cn/data/Net/info/ETF_rate_desc_%s_0_1_9999_0_0_0_jsonp_g.html", innerDate)

	resp, err := utils.Get(url, nil)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	// 去掉 JSONP 包装: g({...}) -> {...}
	text := resp.String()
	text = strings.TrimPrefix(text, "g(")
	text = strings.TrimSuffix(text, ")")

	result := gjson.Parse(text)
	dataMap := result.Get("data.data").Map()

	records := make([]map[string]interface{}, 0, len(dataMap))
	idx := 1
	for _, item := range dataMap {
		record := map[string]interface{}{
			"序号":       idx,
			"基金代码":     item.Get("code").String(),
			"基金名称":     item.Get("name").String(),
			"当前-单位净值":  utils.MustFloat64(item.Get("net").Value()),
			"当前-累计净值":  utils.MustFloat64(item.Get("totalnet").Value()),
			"前一日-单位净值": utils.MustFloat64(item.Get("net1").Value()),
			"前一日-累计净值": utils.MustFloat64(item.Get("totalnet1").Value()),
			"增长值":      utils.MustFloat64(item.Get("ranges").Value()),
			"增长率":      utils.MustFloat64(item.Get("rate").Value()),
			"赎回状态":     item.Get("shstat").String(),
			"申购状态":     item.Get("sgstat").String(),
			"最新-交易日":   item.Get("newdate").String(),
			"最新-单位净值":  utils.MustFloat64(item.Get("newnet").Value()),
			"最新-累计净值":  utils.MustFloat64(item.Get("newtotalnet").Value()),
			"基金类型":     item.Get("typename").String(),
		}

		// 查询日期
		queryDate := innerDate
		if innerDate == "0" && item.Get("newdate").Exists() {
			queryDate = item.Get("newdate").String()
		}
		if queryDate != "0" {
			if t, err := time.Parse("2006-01-02", queryDate); err == nil {
				record["查询日期"] = t.Format("2006-01-02")
			}
		}

		records = append(records, record)
		idx++
	}

	return records, nil
}
