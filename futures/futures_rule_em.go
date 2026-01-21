package futures

import (
	"encoding/json"
	"fmt"

	"github.com/BlakeLiAFK/akshare/utils"
)

// FuturesRuleEM 东方财富期货交易规则
type FuturesRuleEM struct {
	Exchange        string `json:"exchange"`          // 交易所
	Symbol          string `json:"symbol"`            // 品种代码
	Name            string `json:"name"`              // 品种名称
	Unit            string `json:"unit"`              // 交易单位
	QuoteUnit       string `json:"quote_unit"`        // 报价单位
	MinPriceChange  string `json:"min_price_change"`  // 最小变动价位
	PriceLimitRatio string `json:"price_limit_ratio"` // 涨跌停板幅度
	DeliveryMonth   string `json:"delivery_month"`    // 合约交割月份
	TradingHours    string `json:"trading_hours"`     // 交易时间
	LastTradingDay  string `json:"last_trading_day"`  // 最后交易日
	LastDeliveryDay string `json:"last_delivery_day"` // 最后交割日
	DeliveryGrade   string `json:"delivery_grade"`    // 交割品级
	MarginRatio     string `json:"margin_ratio"`      // 最初交易保证金
	DeliveryMethod  string `json:"delivery_method"`   // 交割方式
}

// futuresRuleEMResponse API响应
type futuresRuleEMResponse struct {
	Data []struct {
		JYS     string `json:"jys"`     // 交易所
		PZDM    string `json:"pzdm"`    // 品种代码
		PZMC    string `json:"pzmc"`    // 品种名称
		JYDW    string `json:"jydw"`    // 交易单位
		BJDW    string `json:"bjdw"`    // 报价单位
		ZXBDJW  string `json:"zxbdjw"`  // 最小变动价位
		ZDTBFD  string `json:"zdtbfd"`  // 涨跌停板幅度
		HYJGYF  string `json:"hyjgyf"`  // 合约交割月份
		JYSJ    string `json:"jysj"`    // 交易时间
		ZHJYR   string `json:"zhjyr"`   // 最后交易日
		ZHJGR   string `json:"zhjgr"`   // 最后交割日
		JGPJ    string `json:"jgpj"`    // 交割品级
		ZCJYBZJ string `json:"zcjybzj"` // 最初交易保证金
		JGFS    string `json:"jgfs"`    // 交割方式
	} `json:"Data"`
}

// FuturesRuleEMFunc 东方财富网-期货行情-品种及交易规则
//
// 数据源: https://portal.eastmoneyfutures.com/pages/service/jyts.html#jyrl
//
// 返回:
//   - []FuturesRuleEM: 品种及交易规则
//   - error: 错误信息
func FuturesRuleEMFunc() ([]FuturesRuleEM, error) {
	url := "https://eastmoneyfutures.com/api/ComManage/GetPZJYInfo"

	resp, err := utils.GetWithHeaders(url, nil, map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
		"Referer":    "https://portal.eastmoneyfutures.com/",
	})
	if err != nil {
		return nil, fmt.Errorf("请求期货交易规则失败: %w", err)
	}

	var apiResp futuresRuleEMResponse
	if err := json.Unmarshal(resp.Body(), &apiResp); err != nil {
		return nil, fmt.Errorf("解析期货交易规则响应失败: %w", err)
	}

	result := make([]FuturesRuleEM, 0, len(apiResp.Data))
	for _, item := range apiResp.Data {
		result = append(result, FuturesRuleEM{
			Exchange:        item.JYS,
			Symbol:          item.PZDM,
			Name:            item.PZMC,
			Unit:            item.JYDW,
			QuoteUnit:       item.BJDW,
			MinPriceChange:  item.ZXBDJW,
			PriceLimitRatio: item.ZDTBFD,
			DeliveryMonth:   item.HYJGYF,
			TradingHours:    item.JYSJ,
			LastTradingDay:  item.ZHJYR,
			LastDeliveryDay: item.ZHJGR,
			DeliveryGrade:   item.JGPJ,
			MarginRatio:     item.ZCJYBZJ,
			DeliveryMethod:  item.JGFS,
		})
	}

	return result, nil
}
