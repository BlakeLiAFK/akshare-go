package fx

import (
	"fmt"
	"time"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/tidwall/gjson"
)

// FxSpotQuoteItem 人民币外汇即期报价数据结构
type FxSpotQuoteItem struct {
	CcyPair  string  `json:"ccy_pair"`  // 货币对
	BidPrice float64 `json:"bid_price"` // 买报价
	AskPrice float64 `json:"ask_price"` // 卖报价
}

// FxSwapQuoteItem 人民币外汇远掉报价数据结构
type FxSwapQuoteItem struct {
	CcyPair string `json:"ccy_pair"` // 货币对
	Week1   string `json:"week1"`    // 1周
	Month1  string `json:"month1"`   // 1月
	Month3  string `json:"month3"`   // 3月
	Month6  string `json:"month6"`   // 6月
	Month9  string `json:"month9"`   // 9月
	Year1   string `json:"year1"`    // 1年
}

// FxPairQuoteItem 外币对即期报价数据结构
type FxPairQuoteItem struct {
	CcyPair  string  `json:"ccy_pair"`  // 货币对
	BidPrice float64 `json:"bid_price"` // 买报价
	AskPrice float64 `json:"ask_price"` // 卖报价
}

// FxSpotQuote 获取人民币外汇即期报价
//
// 目标地址: http://www.chinamoney.com.cn/chinese/mkdatapfx/
//
// 返回:
//   - []FxSpotQuoteItem: 人民币外汇即期报价数据列表
//   - error: 错误信息
//
// 示例:
//
//	data, err := fx.FxSpotQuote()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, item := range data {
//	    fmt.Printf("%s 买: %.4f 卖: %.4f\n", item.CcyPair, item.BidPrice, item.AskPrice)
//	}
func FxSpotQuote() ([]FxSpotQuoteItem, error) {
	params := map[string]string{
		"t": fmt.Sprintf("%d", time.Now().UnixMilli()),
	}

	headers := map[string]string{
		"User-Agent": ShortUserAgent,
	}

	resp, err := utils.PostFormWithHeaders(FxSpotURL, params, headers)
	if err != nil {
		return nil, fmt.Errorf("获取人民币外汇即期报价失败: %w", err)
	}

	text := resp.String()
	records := gjson.Get(text, "records")
	if !records.Exists() {
		return nil, fmt.Errorf("解析数据失败")
	}

	var items []FxSpotQuoteItem
	records.ForEach(func(_, value gjson.Result) bool {
		item := FxSpotQuoteItem{
			CcyPair:  value.Get("ccyPair").String(),
			BidPrice: value.Get("bidPrc").Float(),
			AskPrice: value.Get("askPrc").Float(),
		}
		items = append(items, item)
		return true
	})

	return items, nil
}

// FxSwapQuote 获取人民币外汇远掉报价
//
// 目标地址: https://www.chinamoney.com.cn/chinese/index.html
//
// 返回:
//   - []FxSwapQuoteItem: 人民币外汇远掉报价数据列表
//   - error: 错误信息
//
// 示例:
//
//	data, err := fx.FxSwapQuote()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, item := range data {
//	    fmt.Printf("%s 1周: %s 1年: %s\n", item.CcyPair, item.Week1, item.Year1)
//	}
func FxSwapQuote() ([]FxSwapQuoteItem, error) {
	params := map[string]string{
		"t": fmt.Sprintf("%d", time.Now().UnixMilli()),
	}

	headers := map[string]string{
		"User-Agent": ShortUserAgent,
	}

	resp, err := utils.PostFormWithHeaders(FxSwapURL, params, headers)
	if err != nil {
		return nil, fmt.Errorf("获取人民币外汇远掉报价失败: %w", err)
	}

	text := resp.String()
	records := gjson.Get(text, "records")
	if !records.Exists() {
		return nil, fmt.Errorf("解析数据失败")
	}

	var items []FxSwapQuoteItem
	records.ForEach(func(_, value gjson.Result) bool {
		item := FxSwapQuoteItem{
			CcyPair: value.Get("ccyPair").String(),
			Week1:   value.Get("label_1W").String(),
			Month1:  value.Get("label_1M").String(),
			Month3:  value.Get("label_3M").String(),
			Month6:  value.Get("label_6M").String(),
			Month9:  value.Get("label_9M").String(),
			Year1:   value.Get("label_1Y").String(),
		}
		items = append(items, item)
		return true
	})

	return items, nil
}

// FxPairQuote 获取外币对即期报价
//
// 目标地址: http://www.chinamoney.com.cn/chinese/mkdatapfx/
//
// 返回:
//   - []FxPairQuoteItem: 外币对即期报价数据列表
//   - error: 错误信息
//
// 示例:
//
//	data, err := fx.FxPairQuote()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, item := range data {
//	    fmt.Printf("%s 买: %.4f 卖: %.4f\n", item.CcyPair, item.BidPrice, item.AskPrice)
//	}
func FxPairQuote() ([]FxPairQuoteItem, error) {
	params := map[string]string{
		"t": fmt.Sprintf("%d", time.Now().UnixMilli()),
	}

	headers := map[string]string{
		"User-Agent": ShortUserAgent,
	}

	resp, err := utils.PostFormWithHeaders(FxPairURL, params, headers)
	if err != nil {
		return nil, fmt.Errorf("获取外币对即期报价失败: %w", err)
	}

	text := resp.String()
	records := gjson.Get(text, "records")
	if !records.Exists() {
		return nil, fmt.Errorf("解析数据失败")
	}

	var items []FxPairQuoteItem
	records.ForEach(func(_, value gjson.Result) bool {
		item := FxPairQuoteItem{
			CcyPair:  value.Get("ccyPair").String(),
			BidPrice: value.Get("bidPrc").Float(),
			AskPrice: value.Get("askPrc").Float(),
		}
		items = append(items, item)
		return true
	})

	return items, nil
}
