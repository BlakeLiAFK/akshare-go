// Package stock 提供股票行情数据接口
//
// 本模块是 akshare-go 的核心模块，提供 A 股、港股、美股的实时行情、
// 历史 K 线、资金流向、板块数据等功能。
//
// 主要功能:
//   - A股实时行情: StockZhASpotEm, StockZhASpot
//   - A股历史K线: StockZhAHist, StockZhADaily
//   - 港股数据: StockHkSpot, StockHkDaily
//   - 美股数据: StockUsSpot, StockUsDaily
//   - 概念板块: StockBoardConceptNameEm, StockBoardConceptSpotEm
//   - 行业板块: StockBoardIndustryNameEm, StockBoardIndustrySpotEm
//   - 资金流向: StockIndividualFundFlow, StockMarketFundFlow
//   - 股票信息: StockInfoACodeName, StockIndividualInfoEm
//
// 数据来源:
//   - 东方财富 (eastmoney.com)
//   - 新浪财经 (sina.com.cn)
//   - 巨潮资讯 (cninfo.com.cn)
//   - 上交所/深交所/北交所
//
// 使用示例:
//
//	// 获取 A 股实时行情
//	quotes, err := stock.StockZhASpotEm()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, q := range quotes[:10] {
//	    fmt.Printf("%s(%s): %.2f\n", q.Name, q.Code, q.Price)
//	}
//
//	// 获取历史 K 线
//	klines, err := stock.StockZhAHist("000001", "daily", "20240101", "20240115", "qfq")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, k := range klines {
//	    fmt.Printf("%s: O=%.2f H=%.2f L=%.2f C=%.2f\n",
//	        k.Date.Format("2006-01-02"), k.Open, k.High, k.Low, k.Close)
//	}
package stock
