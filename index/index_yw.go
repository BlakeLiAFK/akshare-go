package index

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
)

// YwPriceIndex 义乌小商品价格指数
type YwPriceIndex struct {
	Period         time.Time `json:"period"`           // 期数
	PriceIndex     float64   `json:"price_index"`      // 价格指数
	StockDealIndex float64   `json:"stock_deal_index"` // 场内价格指数
	NetDealIndex   float64   `json:"net_deal_index"`   // 网上价格指数
	OrderDealIndex float64   `json:"order_deal_index"` // 订单价格指数
	OutDealIndex   float64   `json:"out_deal_index"`   // 出口价格指数
}

// YwProsperityIndex 义乌小商品景气指数
type YwProsperityIndex struct {
	Period          time.Time `json:"period"`           // 期数
	ProsperityIndex float64   `json:"prosperity_index"` // 景气指数
	ScopeIndex      float64   `json:"scope_index"`      // 规模指数
	BenefitIndex    float64   `json:"benefit_index"`    // 效益指数
	ConfidentIndex  float64   `json:"confident_index"`  // 市场信心指数
}

// ywPriceResponse 义乌价格指数API响应
type ywPriceResponse struct {
	Data []struct {
		IndexTimeNo         string  `json:"indextimeno"`
		TotalPriceIndex     float64 `json:"totalpriceindex"`
		StockDealPriceIndex float64 `json:"stockdealpriceindex"`
		NetDealPriceIndex   float64 `json:"netdealpriceindex"`
		OrderDealPriceIndex float64 `json:"orderdealpriceindex"`
		OutDealPriceIndex   float64 `json:"outdealpriceindex"`
	} `json:"data"`
}

// ywProsperityResponse 义乌景气指数API响应
type ywProsperityResponse struct {
	Data []struct {
		IndexTimeNo    string  `json:"indextimeno"`
		TotalIndex     float64 `json:"totalindex"`
		ScopeIndex     float64 `json:"scopeindex"`
		BenefitIndex   float64 `json:"benifitindex"`
		ConfidentIndex float64 `json:"confidentindex"`
	} `json:"data"`
}

// IndexYwWeekPrice 义乌小商品周价格指数
//
// 数据源: https://www.ywindex.com/Home/Product/index/
//
// 返回:
//   - []YwPriceIndex: 周价格指数数据
//   - error: 错误信息
func IndexYwWeekPrice() ([]YwPriceIndex, error) {
	return fetchYwPriceIndex("piweek")
}

// IndexYwMonthPrice 义乌小商品月价格指数
//
// 数据源: https://www.ywindex.com/Home/Product/index/
//
// 返回:
//   - []YwPriceIndex: 月价格指数数据
//   - error: 错误信息
func IndexYwMonthPrice() ([]YwPriceIndex, error) {
	return fetchYwPriceIndex("month")
}

// IndexYwProsperity 义乌小商品月景气指数
//
// 数据源: https://www.ywindex.com/Home/Product/index/
//
// 返回:
//   - []YwProsperityIndex: 月景气指数数据
//   - error: 错误信息
func IndexYwProsperity() ([]YwProsperityIndex, error) {
	apiURL := "https://apiserver.chinagoods.com/yiwuindex/v1/active/industry/class/history/bi?gcCode="

	// 跳过SSL验证
	client := resty.New().
		SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetTransport(&http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		})

	resp, err := client.R().Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("请求义乌小商品景气指数失败: %w", err)
	}

	var apiResp ywProsperityResponse
	if err := json.Unmarshal(resp.Body(), &apiResp); err != nil {
		return nil, fmt.Errorf("解析义乌小商品景气指数响应失败: %w", err)
	}

	result := make([]YwProsperityIndex, 0, len(apiResp.Data))
	for _, item := range apiResp.Data {
		period, err := time.Parse("2006-01-02", item.IndexTimeNo[:10])
		if err != nil {
			continue
		}

		result = append(result, YwProsperityIndex{
			Period:          period,
			ProsperityIndex: item.TotalIndex,
			ScopeIndex:      item.ScopeIndex,
			BenefitIndex:    item.BenefitIndex,
			ConfidentIndex:  item.ConfidentIndex,
		})
	}

	return result, nil
}

// fetchYwPriceIndex 获取义乌价格指数
func fetchYwPriceIndex(period string) ([]YwPriceIndex, error) {
	apiURL := fmt.Sprintf("https://apiserver.chinagoods.com/yiwuindex/v1/active/industry/class/history/%s?gcCode=", period)

	// 跳过SSL验证
	client := resty.New().
		SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetTransport(&http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		})

	resp, err := client.R().Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("请求义乌小商品价格指数失败: %w", err)
	}

	var apiResp ywPriceResponse
	if err := json.Unmarshal(resp.Body(), &apiResp); err != nil {
		return nil, fmt.Errorf("解析义乌小商品价格指数响应失败: %w", err)
	}

	result := make([]YwPriceIndex, 0, len(apiResp.Data))
	for _, item := range apiResp.Data {
		periodTime, err := time.Parse("2006-01-02", item.IndexTimeNo[:10])
		if err != nil {
			continue
		}

		result = append(result, YwPriceIndex{
			Period:         periodTime,
			PriceIndex:     item.TotalPriceIndex,
			StockDealIndex: item.StockDealPriceIndex,
			NetDealIndex:   item.NetDealPriceIndex,
			OrderDealIndex: item.OrderDealPriceIndex,
			OutDealIndex:   item.OutDealPriceIndex,
		})
	}

	return result, nil
}
