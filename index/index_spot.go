package index

import (
	"fmt"
	"time"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/tidwall/gjson"
)

// SpotGoodsSymbol 商品现货指数类型
type SpotGoodsSymbol string

const (
	// SpotGoodsBDI 波罗的海干散货指数
	SpotGoodsBDI SpotGoodsSymbol = "波罗的海干散货指数"
	// SpotGoodsGP 钢坯价格指数
	SpotGoodsGP SpotGoodsSymbol = "钢坯价格指数"
	// SpotGoodsPB 澳大利亚粉矿价格
	SpotGoodsPB SpotGoodsSymbol = "澳大利亚粉矿价格"
)

// SpotGoods 新浪财经-商品现货价格指数
//
// 参数:
//   - symbol: 指数类型，可选值: SpotGoodsBDI, SpotGoodsGP, SpotGoodsPB
//
// 返回:
//   - []SpotGoodsQuote: 商品现货价格指数列表
//   - error: 错误信息
//
// 示例:
//
//	// 获取波罗的海干散货指数
//	quotes, err := index.SpotGoods(index.SpotGoodsBDI)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, q := range quotes {
//	    fmt.Printf("%s: %.2f, %.2f%%\n", q.Date.Format("2006-01-02"), q.Price, q.ChangePct)
//	}
func SpotGoods(symbol string) ([]SpotGoodsQuote, error) {
	code, ok := SpotGoodsSymbolMap[symbol]
	if !ok {
		return nil, fmt.Errorf("不支持的指数类型: %s", symbol)
	}

	params := map[string]string{
		"symbol": code,
		"table":  "0",
	}

	headers := map[string]string{
		"Referer": "https://finance.sina.com.cn/futuremarket/spotprice.shtml",
	}

	resp, err := utils.GetWithHeaders(SinaSpotGoodsURL, params, headers)
	if err != nil {
		return nil, fmt.Errorf("请求商品现货指数失败: %w", err)
	}

	result := gjson.Get(resp.String(), "result.data.data")
	if !result.Exists() {
		return nil, fmt.Errorf("未找到商品现货指数数据")
	}

	quotes := make([]SpotGoodsQuote, 0)
	for _, item := range result.Array() {
		dateStr := item.Get("opendate").String()
		date, _ := time.Parse("2006-01-02", dateStr)

		quote := SpotGoodsQuote{
			Date:      date,
			Price:     item.Get("price").Float(),
			Change:    item.Get("zde").Float(),
			ChangePct: item.Get("zdf").Float(),
		}
		quotes = append(quotes, quote)
	}

	return quotes, nil
}

// IndexPriceCFLP 中国公路物流运价指数
//
// 参数:
//   - symbol: 指数类型，可选: "周指数", "月指数", "季度指数", "年度指数"
//
// 返回:
//   - []CFLPIndexQuote: 公路物流运价指数列表
//   - error: 错误信息
type CFLPIndexQuote struct {
	Date       string  `json:"date"`         // 日期
	BaseIndex  float64 `json:"base_index"`   // 定基指数
	ChainIndex float64 `json:"chain_index"`  // 环比指数
	YearOnYear float64 `json:"year_on_year"` // 同比指数
}

func IndexPriceCFLP(symbol string) ([]CFLPIndexQuote, error) {
	indexType, ok := CFLPPriceIndexTypeMap[symbol]
	if !ok {
		return nil, fmt.Errorf("不支持的指数类型: %s", symbol)
	}

	params := map[string]string{
		"marketId":       "1",
		"attribute1":     "5",
		"exponentTypeId": indexType,
		"cateId":         "2",
		"attribute2":     "华北",
		"city":           "",
		"startLine":      "",
		"endLine":        "",
	}

	headers := map[string]string{
		"Origin":     "http://index.0256.cn",
		"Referer":    "http://index.0256.cn/expx.htm",
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
	}

	resp, err := utils.PostWithHeaders(CFLPPriceIndexURL, params, headers)
	if err != nil {
		return nil, fmt.Errorf("请求公路物流运价指数失败: %w", err)
	}

	result := gjson.Parse(resp.String())
	chart1 := result.Get("chart1")

	quotes := make([]CFLPIndexQuote, 0)
	xLabels := chart1.Get("xLebal").Array()
	baseValues := chart1.Get("yLebal").Array()
	chainValues := result.Get("chart2.yLebal").Array()
	yearValues := result.Get("chart3.yLebal").Array()

	for i := 0; i < len(xLabels) && i < len(baseValues) && i < len(chainValues) && i < len(yearValues); i++ {
		quote := CFLPIndexQuote{
			Date:       xLabels[i].String(),
			BaseIndex:  baseValues[i].Float(),
			ChainIndex: chainValues[i].Float(),
			YearOnYear: yearValues[i].Float(),
		}
		quotes = append(quotes, quote)
	}

	return quotes, nil
}

// IndexVolumeCFLP 中国公路物流运量指数
//
// 参数:
//   - symbol: 指数类型，可选: "月指数", "季度指数", "年度指数"
//
// 返回:
//   - []CFLPIndexQuote: 公路物流运量指数列表
//   - error: 错误信息
func IndexVolumeCFLP(symbol string) ([]CFLPIndexQuote, error) {
	indexType, ok := CFLPVolumeIndexTypeMap[symbol]
	if !ok {
		return nil, fmt.Errorf("不支持的指数类型: %s", symbol)
	}

	params := map[string]string{
		"type":       "1",
		"marketId":   "1",
		"expTypeId":  indexType,
		"startDate1": "",
		"endDate1":   "",
		"city":       "",
		"startDate3": "",
		"endDate3":   "",
	}

	headers := map[string]string{
		"Origin":     "http://index.0256.cn",
		"Referer":    "http://index.0256.cn/expx.htm",
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
	}

	resp, err := utils.PostWithHeaders(CFLPVolumeIndexURL, params, headers)
	if err != nil {
		return nil, fmt.Errorf("请求公路物流运量指数失败: %w", err)
	}

	result := gjson.Parse(resp.String())
	chart1 := result.Get("chart1")

	quotes := make([]CFLPIndexQuote, 0)
	xLabels := chart1.Get("xLebal").Array()
	baseValues := chart1.Get("yLebal").Array()
	chainValues := result.Get("chart2.yLebal").Array()
	yearValues := result.Get("chart3.yLebal").Array()

	for i := 0; i < len(xLabels) && i < len(baseValues) && i < len(chainValues) && i < len(yearValues); i++ {
		quote := CFLPIndexQuote{
			Date:       xLabels[i].String(),
			BaseIndex:  baseValues[i].Float(),
			ChainIndex: chainValues[i].Float(),
			YearOnYear: yearValues[i].Float(),
		}
		quotes = append(quotes, quote)
	}

	return quotes, nil
}

// IndexHogSpotPrice 行情宝-生猪市场价格指数
//
// 返回:
//   - []HogIndexQuote: 生猪价格指数列表
//   - error: 错误信息
type HogIndexQuote struct {
	Date      string  `json:"date"`       // 日期
	Index     float64 `json:"index"`      // 指数
	Ma4       float64 `json:"ma4"`        // 4个月均线
	Ma6       float64 `json:"ma6"`        // 6个月均线
	Ma12      float64 `json:"ma12"`       // 12个月均线
	PrePrice  float64 `json:"pre_price"`  // 预售均价
	DealPrice float64 `json:"deal_price"` // 成交均价
	DealWgt   float64 `json:"deal_wgt"`   // 成交均重
}

func IndexHogSpotPrice() ([]HogIndexQuote, error) {
	params := map[string]string{
		"regionId": "0",
	}

	headers := map[string]string{
		"Referer": "https://hqb.nxin.com/pigindex/index.shtml",
	}

	resp, err := utils.GetWithHeaders(HogIndexURL, params, headers)
	if err != nil {
		return nil, fmt.Errorf("请求生猪价格指数失败: %w", err)
	}

	result := gjson.Get(resp.String(), "data")
	if !result.Exists() {
		return nil, fmt.Errorf("未找到生猪价格指数数据")
	}

	quotes := make([]HogIndexQuote, 0)
	for _, item := range result.Array() {
		dateTimestamp := item.Get("日期").Int()
		date := time.Unix(dateTimestamp/1000, 0).Add(8 * time.Hour)

		quote := HogIndexQuote{
			Date:      date.Format("2006-01-02"),
			Index:     item.Get("指数").Float(),
			Ma4:       item.Get("4个月均线").Float(),
			Ma6:       item.Get("6个月均线").Float(),
			Ma12:      item.Get("12个月均线").Float(),
			PrePrice:  item.Get("预售均价").Float(),
			DealPrice: item.Get("成交均价").Float(),
			DealWgt:   item.Get("成交均重").Float(),
		}
		quotes = append(quotes, quote)
	}

	return quotes, nil
}

// IndexERI 浙江省排污权交易指数
//
// 参数:
//   - symbol: 指数类型，可选: "月度", "季度"
//
// 返回:
//   - []ERIIndexQuote: 排污权交易指数列表
//   - error: 错误信息
type ERIIndexQuote struct {
	Date   string  `json:"date"`   // 日期
	Index  float64 `json:"index"`  // 交易指数
	Volume float64 `json:"volume"` // 成交量
	Amount float64 `json:"amount"` // 成交额
}

func IndexERI(symbol string) ([]ERIIndexQuote, error) {
	cycle, ok := ERIndexTypeMap[symbol]
	if !ok {
		return nil, fmt.Errorf("不支持的指数类型: %s", symbol)
	}

	params := map[string]string{
		"cycle":    cycle,
		"regionId": "1",
		"structId": "1",
		"pageSize": "5000",
		"indexId":  "1",
		"orderBy":  "stage.publishTime",
	}

	headers := map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
	}

	resp, err := utils.GetWithHeaders(ERIndexURL, params, headers)
	if err != nil {
		return nil, fmt.Errorf("请求排污权交易指数失败: %w", err)
	}

	data := gjson.Get(resp.String(), "data")
	if !data.Exists() {
		return nil, fmt.Errorf("未找到排污权交易指数数据")
	}

	// 获取索引数据
	quotes := make([]ERIIndexQuote, 0)
	for _, item := range data.Array() {
		publishTime := item.Get("stage.publishTime").String()

		quote := ERIIndexQuote{
			Date:  publishTime[:10], // 简化日期处理
			Index: item.Get("indexValue").Float(),
		}
		quotes = append(quotes, quote)
	}

	// 获取统计数据
	resp2, err := utils.GetWithHeaders("https://zs.zjpwq.net/pwq-index-webapi/dataStatistics", params, headers)
	if err == nil {
		statsData := gjson.Get(resp2.String(), "data")
		if statsData.Exists() {
			for i, item := range statsData.Array() {
				if i < len(quotes) {
					quotes[i].Volume = item.Get("totalQuantity").Float()
					quotes[i].Amount = item.Get("totalCost").Float()
				}
			}
		}
	}

	return quotes, nil
}

// IndexYW 义乌小商品指数
//
// 参数:
//   - symbol: 指数类型，可选: "周价格指数", "月价格指数", "月景气指数"
//
// 返回:
//   - []YWIndexQuote: 义乌小商品指数列表
//   - error: 错误信息
func IndexYW(symbol string) ([]YWIndexQuote, error) {
	var url string
	var isBoomIndex bool

	if symbol == "月景气指数" {
		url = fmt.Sprintf(YWIndexURL, "bi?gcCode=")
		isBoomIndex = true
	} else {
		indexType, ok := YWIndexTypeMap[symbol]
		if !ok {
			return nil, fmt.Errorf("不支持的指数类型: %s", symbol)
		}
		url = fmt.Sprintf(YWIndexURL, indexType+"?gcCode=")
		isBoomIndex = false
	}

	headers := map[string]string{
		"Referer": "https://www.ywindex.com/",
	}

	resp, err := utils.GetWithHeaders(url, nil, headers)
	if err != nil {
		return nil, fmt.Errorf("请求义乌小商品指数失败: %w", err)
	}

	result := gjson.Get(resp.String(), "data")
	if !result.Exists() {
		return nil, fmt.Errorf("未找到义乌小商品指数数据")
	}

	quotes := make([]YWIndexQuote, 0)
	for _, item := range result.Array() {
		var quote YWIndexQuote
		periodNo := item.Get("indextimeno").String()

		if isBoomIndex {
			// 月景气指数
			quote = YWIndexQuote{
				PeriodNo:       periodNo,
				BoomIndex:      item.Get("totalindex").Float(),
				ScopeIndex:     item.Get("scopeindex").Float(),
				ProfitIndex:    item.Get("benifitindex").Float(),
				ConfidentIndex: item.Get("confidentindex").Float(),
			}
		} else {
			// 价格指数
			quote = YWIndexQuote{
				PeriodNo:    periodNo,
				PriceIndex:  item.Get("totalpriceindex").Float(),
				FieldPrice:  item.Get("stockdealpriceindex").Float(),
				NetPrice:    item.Get("netdealpriceindex").Float(),
				OrderPrice:  item.Get("orderdealpriceindex").Float(),
				ExportPrice: item.Get("outdealpriceindex").Float(),
			}
		}
		quotes = append(quotes, quote)
	}

	return quotes, nil
}

// IndexSugarMSweet 沐甜科技数据中心-中国食糖指数
//
// 返回:
//   - []SugarIndex: 中国食糖指数列表
//   - error: 错误信息
func IndexSugarMSweet() ([]SugarIndex, error) {
	params := map[string]string{
		"struts.portlet.action": "/portlet/price!getSTZSJson.action",
		"moduleId":              "cb752447cfe24b44b18c7a7e9abab048",
	}

	headers := map[string]string{
		"Referer": "https://www.msweet.com.cn/",
	}

	resp, err := utils.GetWithHeaders(SugarIndexURL, params, headers)
	if err != nil {
		return nil, fmt.Errorf("请求中国食糖指数失败: %w", err)
	}

	result := gjson.Parse(resp.String())
	category := result.Get("category").Array()
	data := result.Get("data").Array()

	quotes := make([]SugarIndex, 0)
	for i := 0; i < len(category) && i < len(data); i++ {
		quote := SugarIndex{
			Date:           category[i].String(),
			CompositePrice: data[i].Float(),
		}
		if i+1 < len(data) {
			quote.RawSugarPrice = data[i+1].Float()
		}
		if i+2 < len(data) {
			quote.SpotPrice = data[i+2].Float()
		}
		quotes = append(quotes, quote)
	}

	return quotes, nil
}

// DrewryWCIIndex 已移至 index_drewry.go
