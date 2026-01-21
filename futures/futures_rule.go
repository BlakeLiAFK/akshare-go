package futures

import (
	"fmt"
	"strings"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/PuerkitoBio/goquery"
)

// FuturesRule 期货交易规则
type FuturesRule struct {
	Exchange        string  `json:"exchange"`          // 交易所
	Symbol          string  `json:"symbol"`            // 合约代码
	Name            string  `json:"name"`              // 合约名称
	Multiplier      float64 `json:"multiplier"`        // 合约乘数
	MinPriceChange  float64 `json:"min_price_change"`  // 最小变动价位
	MarginRatio     float64 `json:"margin_ratio"`      // 交易保证金比例
	PriceLimitRatio float64 `json:"price_limit_ratio"` // 涨跌停板幅度
	MaxOrderSize    int64   `json:"max_order_size"`    // 限价单每笔最大下单手数
	DeliveryMonth   string  `json:"delivery_month"`    // 交割月份
	LastTradingDay  string  `json:"last_trading_day"`  // 最后交易日
	LastDeliveryDay string  `json:"last_delivery_day"` // 最后交割日
	TradingHours    string  `json:"trading_hours"`     // 交易时间
}

// FuturesRuleFunc 国泰君安期货-交易日历数据表
//
// 数据源: https://www.gtjaqh.com/pc/calendar.html
//
// 参数:
//   - date: 需要指定为交易日，且是近期的日期，格式 "20231205"
//
// 返回:
//   - []FuturesRule: 交易日历数据
//   - error: 错误信息
func FuturesRuleFunc(date string) ([]FuturesRule, error) {
	url := "https://www.gtjaqh.com/pc/calendar"
	params := map[string]string{
		"date": date,
	}

	resp, err := utils.GetWithHeaders(url, params, map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
	})
	if err != nil {
		return nil, fmt.Errorf("请求交易日历数据失败: %w", err)
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(resp.Body())))
	if err != nil {
		return nil, fmt.Errorf("解析HTML失败: %w", err)
	}

	var result []FuturesRule
	// 查找表格，跳过表头
	doc.Find("table").First().Find("tr").Each(func(i int, s *goquery.Selection) {
		if i < 2 { // 跳过表头行
			return
		}
		var row []string
		s.Find("td").Each(func(_ int, td *goquery.Selection) {
			row = append(row, strings.TrimSpace(td.Text()))
		})
		if len(row) >= 9 {
			rule := FuturesRule{
				Exchange:        row[0],
				Symbol:          row[1],
				Name:            row[2],
				Multiplier:      utils.MustParseFloat(row[3]),
				MinPriceChange:  utils.MustParseFloat(row[4]),
				MarginRatio:     utils.MustParseFloat(strings.TrimSuffix(row[5], "%")),
				PriceLimitRatio: utils.MustParseFloat(strings.TrimSuffix(row[6], "%")),
				MaxOrderSize:    int64(utils.MustParseFloat(row[7])),
			}
			if len(row) > 8 {
				rule.DeliveryMonth = row[8]
			}
			result = append(result, rule)
		}
	})

	return result, nil
}
