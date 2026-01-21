package fund

import (
	"fmt"

	"github.com/BlakeLiAFK/akshare/stock_feature"
	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/tidwall/gjson"
)

// FundStockPositionLg 获取乐咕乐股基金股票型基金仓位数据
// https://legulegu.com/stockdata/fund-position/pos-stock
func FundStockPositionLg(date string) ([]map[string]interface{}, error) {
	url := "https://legulegu.com/api/stockdata/fund-position"
	token := stock_feature.GetTokenLg()

	params := map[string]string{
		"token":    token,
		"type":     "pos_stock",
		"category": "总仓位",
		"marketId": "5",
	}

	csrfHeaders, err := stock_feature.GetCookieCsrf("https://legulegu.com/stockdata/fund-position/pos-stock")
	if err != nil {
		return nil, fmt.Errorf("获取CSRF失败: %w", err)
	}

	resp, err := utils.GetWithHeaders(url, params, csrfHeaders)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	result := gjson.ParseBytes(resp.Body())
	items := result.Array()

	records := make([]map[string]interface{}, 0, len(items))
	for _, item := range items {
		record := map[string]interface{}{
			"date":     item.Get("date").String(),
			"close":    utils.MustFloat64(item.Get("close").String()),
			"position": utils.MustFloat64(item.Get("position").String()),
		}
		records = append(records, record)
	}

	return records, nil
}

// FundBalancePositionLg 获取乐咕乐股平衡混合型基金仓位数据
// https://legulegu.com/stockdata/fund-position/pos-pingheng
func FundBalancePositionLg() ([]map[string]interface{}, error) {
	url := "https://legulegu.com/api/stockdata/fund-position"
	token := stock_feature.GetTokenLg()

	params := map[string]string{
		"token":    token,
		"type":     "pos_pingheng",
		"category": "总仓位",
		"marketId": "5",
	}

	csrfHeaders, err := stock_feature.GetCookieCsrf("https://legulegu.com/stockdata/fund-position/pos-pingheng")
	if err != nil {
		return nil, fmt.Errorf("获取CSRF失败: %w", err)
	}

	resp, err := utils.GetWithHeaders(url, params, csrfHeaders)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	result := gjson.ParseBytes(resp.Body())
	items := result.Array()

	records := make([]map[string]interface{}, 0, len(items))
	for _, item := range items {
		record := map[string]interface{}{
			"date":     item.Get("date").String(),
			"close":    utils.MustFloat64(item.Get("close").String()),
			"position": utils.MustFloat64(item.Get("position").String()),
		}
		records = append(records, record)
	}

	return records, nil
}

// FundLinghuoPositionLg 获取乐咕乐股灵活配置型基金仓位数据
// https://legulegu.com/stockdata/fund-position/pos-linghuo
func FundLinghuoPositionLg() ([]map[string]interface{}, error) {
	url := "https://legulegu.com/api/stockdata/fund-position"
	token := stock_feature.GetTokenLg()

	params := map[string]string{
		"token":    token,
		"type":     "pos_linghuo",
		"category": "总仓位",
		"marketId": "5",
	}

	csrfHeaders, err := stock_feature.GetCookieCsrf("https://legulegu.com/stockdata/fund-position/pos-linghuo")
	if err != nil {
		return nil, fmt.Errorf("获取CSRF失败: %w", err)
	}

	resp, err := utils.GetWithHeaders(url, params, csrfHeaders)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	result := gjson.ParseBytes(resp.Body())
	items := result.Array()

	records := make([]map[string]interface{}, 0, len(items))
	for _, item := range items {
		record := map[string]interface{}{
			"date":     item.Get("date").String(),
			"close":    utils.MustFloat64(item.Get("close").String()),
			"position": utils.MustFloat64(item.Get("position").String()),
		}
		records = append(records, record)
	}

	return records, nil
}
