package interest_rate

import (
	"fmt"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/tidwall/gjson"
)

const (
	// 东方财富银行间拆借利率URL
	emInterbankURL = "https://datacenter-web.eastmoney.com/api/data/v1/get"
)

// 市场代码映射
var marketMap = map[string]string{
	"上海银行同业拆借市场":  "001",
	"中国银行同业拆借市场":  "002",
	"伦敦银行同业拆借市场":  "003",
	"欧洲银行同业拆借市场":  "004",
	"香港银行同业拆借市场":  "005",
	"新加坡银行同业拆借市场": "006",
}

// 货币代码映射
var symbolMap = map[string]string{
	"Shibor人民币": "CNY",
	"Chibor人民币": "CNY",
	"Libor英镑":   "GBP",
	"Libor欧元":   "EUR",
	"Libor美元":   "USD",
	"Libor日元":   "JPY",
	"Euribor欧元": "EUR",
	"Hibor美元":   "USD",
	"Hibor人民币":  "CNH",
	"Hibor港币":   "HKD",
	"Sibor星元":   "SGD",
	"Sibor美元":   "USD",
}

// 期限代码映射
var indicatorMap = map[string]string{
	"隔夜":  "001",
	"1周":  "101",
	"2周":  "102",
	"3周":  "103",
	"1月":  "201",
	"2月":  "202",
	"3月":  "203",
	"4月":  "204",
	"5月":  "205",
	"6月":  "206",
	"7月":  "207",
	"8月":  "208",
	"9月":  "209",
	"10月": "210",
	"11月": "211",
	"1年":  "301",
}

// RateInterbankItem 银行间拆借利率数据结构
type RateInterbankItem struct {
	ReportDate string  `json:"report_date"` // 报告日
	Rate       float64 `json:"rate"`        // 利率
	Change     float64 `json:"change"`      // 涨跌
}

// RateInterbank 获取东方财富-银行间拆借利率数据
//
// 目标地址: https://data.eastmoney.com/shibor/shibor.aspx
//
// 参数:
//   - market: 市场，可选 "上海银行同业拆借市场", "中国银行同业拆借市场", "伦敦银行同业拆借市场", "欧洲银行同业拆借市场", "香港银行同业拆借市场", "新加坡银行同业拆借市场"
//   - symbol: 品种，可选 "Shibor人民币", "Chibor人民币", "Libor英镑", "Libor欧元", "Libor美元", "Libor日元", "Euribor欧元", "Hibor美元", "Hibor人民币", "Hibor港币", "Sibor星元", "Sibor美元"
//   - indicator: 期限，可选 "隔夜", "1周", "2周", "3周", "1月", "2月", "3月", "4月", "5月", "6月", "7月", "8月", "9月", "10月", "11月", "1年"
//
// 返回:
//   - []RateInterbankItem: 拆借利率数据列表
//   - error: 错误信息
//
// 示例:
//
//	data, err := interest_rate.RateInterbank("上海银行同业拆借市场", "Shibor人民币", "3月")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, item := range data[:10] {
//	    fmt.Printf("%s %.4f%% %.4f\n", item.ReportDate, item.Rate, item.Change)
//	}
func RateInterbank(market, symbol, indicator string) ([]RateInterbankItem, error) {
	marketCode, ok := marketMap[market]
	if !ok {
		return nil, fmt.Errorf("无效的市场参数: %s", market)
	}

	symbolCode, ok := symbolMap[symbol]
	if !ok {
		return nil, fmt.Errorf("无效的品种参数: %s", symbol)
	}

	indicatorCode, ok := indicatorMap[indicator]
	if !ok {
		return nil, fmt.Errorf("无效的期限参数: %s", indicator)
	}

	headers := map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
	}

	var allItems []RateInterbankItem
	page := 1

	for {
		params := map[string]string{
			"reportName":   "RPT_IMP_INTRESTRATEN",
			"columns":      "REPORT_DATE,REPORT_PERIOD,IR_RATE,CHANGE_RATE,INDICATOR_ID,LATEST_RECORD,MARKET,MARKET_CODE,CURRENCY,CURRENCY_CODE",
			"quoteColumns": "",
			"filter":       fmt.Sprintf(`(MARKET_CODE="%s")(CURRENCY_CODE="%s")(INDICATOR_ID="%s")`, marketCode, symbolCode, indicatorCode),
			"pageNumber":   fmt.Sprintf("%d", page),
			"pageSize":     "500",
			"sortTypes":    "-1",
			"sortColumns":  "REPORT_DATE",
			"source":       "WEB",
			"client":       "WEB",
		}

		resp, err := utils.GetWithHeaders(emInterbankURL, params, headers)
		if err != nil {
			break
		}

		text := resp.String()
		result := gjson.Get(text, "result")
		if !result.Exists() {
			break
		}

		data := result.Get("data")
		if !data.Exists() || !data.IsArray() || len(data.Array()) == 0 {
			break
		}

		data.ForEach(func(_, value gjson.Result) bool {
			item := RateInterbankItem{
				ReportDate: value.Get("REPORT_DATE").String()[:10],
				Rate:       value.Get("IR_RATE").Float(),
				Change:     value.Get("CHANGE_RATE").Float(),
			}
			allItems = append(allItems, item)
			return true
		})

		totalPages := result.Get("pages").Int()
		if int64(page) >= totalPages {
			break
		}
		page++
	}

	// 按日期升序排序
	for i := 0; i < len(allItems)/2; i++ {
		j := len(allItems) - 1 - i
		allItems[i], allItems[j] = allItems[j], allItems[i]
	}

	return allItems, nil
}
