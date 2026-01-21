package index

import (
	"fmt"
	"strings"
	"time"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/tidwall/gjson"
)

// IndexGlobalSpotEM 东方财富网-全球指数实时行情
//
// 返回:
//   - []IndexQuote: 全球指数实时行情列表
//   - error: 错误信息
//
// 示例:
//
//	quotes, err := index.IndexGlobalSpotEM()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, q := range quotes {
//	    fmt.Printf("%s(%s): %.2f, %.2f%%\n", q.Name, q.Code, q.Price, q.ChangePct)
//	}
func IndexGlobalSpotEM() ([]IndexQuote, error) {
	params := map[string]string{
		"np":   "2",
		"fltt": "1",
		"invt": "2",
		"fs": "i:1.000001,i:0.399001,i:0.399005,i:0.399006,i:1.000300,i:100.HSI,i:100.HSCEI,i:124.HSCCI," +
			"i:100.TWII,i:100.N225,i:100.KOSPI200,i:100.KS11,i:100.STI,i:100.SENSEX,i:100.KLSE,i:100.SET," +
			"i:100.PSI,i:100.KSE100,i:100.VNINDEX,i:100.JKSE,i:100.CSEALL,i:100.SX5E,i:100.FTSE,i:100.MCX," +
			"i:100.AXX,i:100.FCHI,i:100.GDAXI,i:100.RTS,i:100.IBEX,i:100.PSI20,i:100.OMXC20,i:100.BFX," +
			"i:100.AEX,i:100.WIG,i:100.OMXSPI,i:100.SSMI,i:100.HEX,i:100.OSEBX,i:100.ATX,i:100.MIB," +
			"i:100.ASE,i:100.ICEXI,i:100.PX,i:100.ISEQ,i:100.DJIA,i:100.SPX,i:100.NDX,i:100.TSX," +
			"i:100.BVSP,i:100.MXX,i:100.AS51,i:100.AORD,i:100.NZ50,i:100.UDI,i:100.BDI,i:100.CRB",
		"fields": "f12,f13,f14,f292,f1,f2,f4,f3,f152,f17,f18,f15,f16,f7,f124",
		"fid":    "f3",
		"pn":     "1",
		"pz":     "200",
		"po":     "1",
		"dect":   "1",
		"wbp2u":  "|0|0|0|web",
	}

	headers := map[string]string{
		"Referer": "https://quote.eastmoney.com/center/gridlist.html#global_qtzs",
	}
	resp, err := utils.GetWithHeaders(EmIndexListURL, params, headers)
	if err != nil {
		return nil, fmt.Errorf("请求全球指数行情失败: %w", err)
	}

	return parseGlobalIndexSpot(resp.String())
}

// parseGlobalIndexSpot 解析全球指数行情响应
func parseGlobalIndexSpot(text string) ([]IndexQuote, error) {
	data := gjson.Get(text, "data.diff")
	if !data.Exists() || data.Type != gjson.JSON {
		return nil, fmt.Errorf("未找到全球指数数据")
	}

	quotes := make([]IndexQuote, 0)
	// diff 是一个对象，需要遍历其值
	data.ForEach(func(key, value gjson.Result) bool {
		quote := IndexQuote{
			Seq:       int(value.Get("f0").Int()) + 1,
			Code:      value.Get("f12").String(),
			Name:      value.Get("f14").String(),
			Open:      value.Get("f17").Float() / 100,
			Change:    value.Get("f4").Float() / 100,
			ChangePct: value.Get("f3").Float() / 100,
			Price:     value.Get("f2").Float() / 100,
			High:      value.Get("f15").Float() / 100,
			Low:       value.Get("f16").Float() / 100,
			PreClose:  value.Get("f18").Float() / 100,
			Amplitude: value.Get("f7").Float() / 100,
		}

		// 解析更新时间 (Unix时间戳，单位秒)
		if ts := value.Get("f124").Int(); ts > 0 {
			t := time.Unix(ts, 0)
			quote.UpdateTime = t.Format("2006-01-02 15:04:05")
		}

		if quote.Code != "" {
			quotes = append(quotes, quote)
		}
		return true
	})

	return quotes, nil
}

// IndexGlobalHistEM 东方财富网-全球指数历史行情
//
// 参数:
//   - symbol: 指数名称，如 "美元指数"、"标普500"、"纳斯达克"等
//   - 可以通过 IndexGlobalSpotEM() 获取所有可用指数
//
// 返回:
//   - []GlobalIndexKLine: 全球指数K线数据
//   - error: 错误信息
//
// 示例:
//
//	// 获取美元指数历史行情
//	klines, err := index.IndexGlobalHistEM("美元指数")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, k := range klines {
//	    fmt.Printf("%s: 开盘=%.2f, 收盘=%.2f\n", k.Date.Format("2006-01-02"), k.Open, k.Close)
//	}
func IndexGlobalHistEM(symbol string) ([]GlobalIndexKLine, error) {
	symbolInfo, ok := GlobalEMSymbolMap[symbol]
	if !ok {
		return nil, fmt.Errorf("未找到指数: %s，请先调用 IndexGlobalSpotEM() 查看可用指数", symbol)
	}

	secid := fmt.Sprintf("%s.%s", symbolInfo.MarketID, symbolInfo.Code)

	params := map[string]string{
		"secid":   secid,
		"klt":     "101",
		"fqt":     "1",
		"lmt":     "50000",
		"end":     "20500000",
		"iscca":   "1",
		"fields1": "f1,f2,f3,f4,f5,f6,f7,f8",
		"fields2": "f51,f52,f53,f54,f55,f56,f57,f58,f59,f60,f61,f62,f63,f64",
		"ut":      "f057cbcbce2a86e2866ab8877db1d059",
		"forcect": "1",
	}

	headers := map[string]string{
		"Referer": "https://quote.eastmoney.com/gb/zsUDI.html",
	}
	resp, err := utils.GetWithHeaders(EmIndexKLineURL, params, headers)
	if err != nil {
		return nil, fmt.Errorf("请求全球指数历史数据失败: %w", err)
	}

	klinesData := gjson.Get(resp.String(), "data.klines")
	if !klinesData.Exists() {
		return nil, fmt.Errorf("未找到历史数据")
	}

	klines := make([]GlobalIndexKLine, 0)
	for _, item := range klinesData.Array() {
		// 数据格式: "日期,今开,最新价,最高,最低,-,-,振幅,-,-,-,-,-,-"
		parts := strings.Split(item.String(), ",")
		if len(parts) < 5 {
			continue
		}

		date, _ := time.Parse("2006-01-02", parts[0])

		kline := GlobalIndexKLine{
			Date:  date,
			Open:  utils.MustFloat64(parts[1]),
			Close: utils.MustFloat64(parts[2]),
			High:  utils.MustFloat64(parts[3]),
			Low:   utils.MustFloat64(parts[4]),
		}
		klines = append(klines, kline)
	}

	return klines, nil
}
