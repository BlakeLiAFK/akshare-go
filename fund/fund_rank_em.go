package fund

import (
	"fmt"
	"strings"
	"time"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/tidwall/gjson"
)

// FundOpenFundRankEm 获取开放式基金排名
// 参数: symbol 基金类型 "全部"、"股票型"、"混合型"、"债券型"、"指数型"、"QDII"、"LOF"、"FOF"
func FundOpenFundRankEm(symbol string) ([]map[string]interface{}, error) {
	if symbol == "" {
		symbol = "全部"
	}

	symbolMap := map[string]string{
		"全部":   "all",
		"股票型":  "gp",
		"混合型":  "hh",
		"债券型":  "zq",
		"指数型":  "zs",
		"QDII": "qdii",
		"LOF":  "lof",
		"FOF":  "fof",
	}

	ft, ok := symbolMap[symbol]
	if !ok {
		return nil, fmt.Errorf("不支持的基金类型: %s", symbol)
	}

	url := "https://fund.eastmoney.com/data/rankhandler.aspx"
	params := map[string]string{
		"op":         "ph",
		"dt":         "kf",
		"ft":         ft,
		"rs":         "",
		"gs":         "0",
		"sc":         "zzf",
		"st":         "desc",
		"sd":         time.Now().AddDate(-1, 0, 0).Format("2006-01-02"),
		"ed":         time.Now().Format("2006-01-02"),
		"qdii":       "1",
		"tabSubtype": ",,,,,",
		"pi":         "1",
		"pn":         "50000",
		"dx":         "1",
	}

	headers := map[string]string{
		"Referer": "https://fund.eastmoney.com/data/fundranking.html",
	}

	resp, err := utils.GetWithHeaders(url, params, headers)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	text := resp.String()
	// 去除 var rankData = {...} 格式
	text = strings.TrimPrefix(text, "var rankData = ")
	text = strings.TrimSuffix(text, ";")

	result := gjson.Parse(text)
	datas := result.Get("datas").Array()

	records := make([]map[string]interface{}, 0, len(datas))
	for _, item := range datas {
		parts := strings.Split(item.String(), ",")
		if len(parts) < 18 {
			continue
		}

		record := map[string]interface{}{
			"基金代码": parts[0],
			"基金简称": parts[1],
			"日期":   parts[3],
			"单位净值": utils.MustFloat64(parts[4]),
			"累计净值": utils.MustFloat64(parts[5]),
			"日增长率": utils.MustFloat64(parts[6]),
			"近1周":  utils.MustFloat64(parts[7]),
			"近1月":  utils.MustFloat64(parts[8]),
			"近3月":  utils.MustFloat64(parts[9]),
			"近6月":  utils.MustFloat64(parts[10]),
			"近1年":  utils.MustFloat64(parts[11]),
			"近2年":  utils.MustFloat64(parts[12]),
			"近3年":  utils.MustFloat64(parts[13]),
			"今年来":  utils.MustFloat64(parts[14]),
			"成立来":  utils.MustFloat64(parts[15]),
			"手续费":  utils.MustFloat64(parts[16]),
		}
		records = append(records, record)
	}

	return records, nil
}

// FundExchangeRankEm 获取场内交易基金排名
func FundExchangeRankEm() ([]map[string]interface{}, error) {
	url := "https://fund.eastmoney.com/data/rankhandler.aspx"
	params := map[string]string{
		"op":         "ph",
		"dt":         "fb",
		"ft":         "ct",
		"rs":         "",
		"gs":         "0",
		"sc":         "zzf",
		"st":         "desc",
		"sd":         time.Now().AddDate(-1, 0, 0).Format("2006-01-02"),
		"ed":         time.Now().Format("2006-01-02"),
		"qdii":       "1",
		"tabSubtype": ",,,,,",
		"pi":         "1",
		"pn":         "30000",
		"dx":         "1",
	}

	headers := map[string]string{
		"Referer": "https://fund.eastmoney.com/data/fundblist.html",
	}

	resp, err := utils.GetWithHeaders(url, params, headers)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	text := resp.String()
	text = strings.TrimPrefix(text, "var rankData = ")
	text = strings.TrimSuffix(text, ";")

	result := gjson.Parse(text)
	datas := result.Get("datas").Array()

	records := make([]map[string]interface{}, 0, len(datas))
	for _, item := range datas {
		parts := strings.Split(item.String(), ",")
		if len(parts) < 18 {
			continue
		}

		record := map[string]interface{}{
			"基金代码": parts[0],
			"基金简称": parts[1],
			"日期":   parts[3],
			"单位净值": utils.MustFloat64(parts[4]),
			"累计净值": utils.MustFloat64(parts[5]),
			"日增长率": utils.MustFloat64(parts[6]),
			"近1周":  utils.MustFloat64(parts[7]),
			"近1月":  utils.MustFloat64(parts[8]),
			"近3月":  utils.MustFloat64(parts[9]),
			"近6月":  utils.MustFloat64(parts[10]),
			"近1年":  utils.MustFloat64(parts[11]),
			"近2年":  utils.MustFloat64(parts[12]),
			"近3年":  utils.MustFloat64(parts[13]),
			"今年来":  utils.MustFloat64(parts[14]),
			"成立来":  utils.MustFloat64(parts[15]),
			"手续费":  utils.MustFloat64(parts[16]),
			"成立日期": parts[17],
		}
		records = append(records, record)
	}

	return records, nil
}

// FundMoneyRankEm 获取货币型基金排名
func FundMoneyRankEm() ([]map[string]interface{}, error) {
	url := "https://fund.eastmoney.com/data/rankhandler.aspx"
	params := map[string]string{
		"op":         "ph",
		"dt":         "kf",
		"ft":         "hb",
		"rs":         "",
		"gs":         "0",
		"sc":         "w",
		"st":         "desc",
		"sd":         time.Now().AddDate(0, -3, 0).Format("2006-01-02"),
		"ed":         time.Now().Format("2006-01-02"),
		"qdii":       "1",
		"tabSubtype": ",,,,,",
		"pi":         "1",
		"pn":         "10000",
		"dx":         "1",
	}

	headers := map[string]string{
		"Referer": "https://fund.eastmoney.com/data/fbsrankings.html",
	}

	resp, err := utils.GetWithHeaders(url, params, headers)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	text := resp.String()
	text = strings.TrimPrefix(text, "var rankData = ")
	text = strings.TrimSuffix(text, ";")

	result := gjson.Parse(text)
	datas := result.Get("datas").Array()

	records := make([]map[string]interface{}, 0, len(datas))
	for _, item := range datas {
		parts := strings.Split(item.String(), ",")
		if len(parts) < 17 {
			continue
		}

		record := map[string]interface{}{
			"基金代码":    parts[0],
			"基金简称":    parts[1],
			"日期":      parts[3],
			"万份收益":    utils.MustFloat64(parts[4]),
			"7日年化收益率": utils.MustFloat64(parts[5]),
			"近1月年化":   utils.MustFloat64(parts[6]),
			"近3月年化":   utils.MustFloat64(parts[7]),
			"近6月年化":   utils.MustFloat64(parts[8]),
			"近1年年化":   utils.MustFloat64(parts[9]),
			"近2年年化":   utils.MustFloat64(parts[10]),
			"近3年年化":   utils.MustFloat64(parts[11]),
			"今年来年化":   utils.MustFloat64(parts[12]),
			"成立来年化":   utils.MustFloat64(parts[13]),
			"手续费":     utils.MustFloat64(parts[14]),
		}
		records = append(records, record)
	}

	return records, nil
}

// FundLcxRankEm 获取理财型基金排名
func FundLcxRankEm() ([]map[string]interface{}, error) {
	url := "https://fund.eastmoney.com/data/rankhandler.aspx"
	params := map[string]string{
		"op":         "ph",
		"dt":         "rq",
		"ft":         "all",
		"rs":         "",
		"gs":         "0",
		"sc":         "w",
		"st":         "desc",
		"sd":         time.Now().AddDate(0, -3, 0).Format("2006-01-02"),
		"ed":         time.Now().Format("2006-01-02"),
		"qdii":       "1",
		"tabSubtype": ",,,,,",
		"pi":         "1",
		"pn":         "10000",
		"dx":         "1",
	}

	headers := map[string]string{
		"Referer": "https://fund.eastmoney.com/data/lcxjjranking.html",
	}

	resp, err := utils.GetWithHeaders(url, params, headers)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	text := resp.String()
	text = strings.TrimPrefix(text, "var rankData = ")
	text = strings.TrimSuffix(text, ";")

	result := gjson.Parse(text)
	datas := result.Get("datas").Array()

	records := make([]map[string]interface{}, 0, len(datas))
	for _, item := range datas {
		parts := strings.Split(item.String(), ",")
		if len(parts) < 15 {
			continue
		}

		record := map[string]interface{}{
			"基金代码": parts[0],
			"基金简称": parts[1],
			"日期":   parts[3],
			"万份收益": utils.MustFloat64(parts[4]),
			"7日年化": utils.MustFloat64(parts[5]),
			"近1月":  utils.MustFloat64(parts[6]),
			"近3月":  utils.MustFloat64(parts[7]),
			"近6月":  utils.MustFloat64(parts[8]),
			"近1年":  utils.MustFloat64(parts[9]),
			"今年来":  utils.MustFloat64(parts[10]),
			"成立来":  utils.MustFloat64(parts[11]),
			"手续费":  utils.MustFloat64(parts[12]),
		}
		records = append(records, record)
	}

	return records, nil
}

// FundHkRankEm 获取香港基金排名
func FundHkRankEm() ([]map[string]interface{}, error) {
	url := "https://overseas.1234567.com.cn/f10DataApi/api/HKFDJZ"
	params := map[string]string{
		"pageIndex": "1",
		"pageSize":  "10000",
		"sortField": "SYL",
		"sortType":  "-1",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	result := gjson.ParseBytes(resp.Body())
	items := result.Get("Data").Array()

	records := make([]map[string]interface{}, 0, len(items))
	for _, item := range items {
		record := map[string]interface{}{
			"基金代码": item.Get("FCODE").String(),
			"基金简称": item.Get("SHORTNAME").String(),
			"单位净值": utils.MustFloat64(item.Get("NAV").String()),
			"日期":   item.Get("FSRQ").String(),
			"日增长率": utils.MustFloat64(item.Get("RZDF").String()),
			"近1周":  utils.MustFloat64(item.Get("SYL_Z").String()),
			"近1月":  utils.MustFloat64(item.Get("SYL_Y").String()),
			"近3月":  utils.MustFloat64(item.Get("SYL_3Y").String()),
			"近6月":  utils.MustFloat64(item.Get("SYL_6Y").String()),
			"近1年":  utils.MustFloat64(item.Get("SYL_1N").String()),
			"近3年":  utils.MustFloat64(item.Get("SYL_3N").String()),
			"今年来":  utils.MustFloat64(item.Get("SYL_JN").String()),
			"成立来":  utils.MustFloat64(item.Get("SYL").String()),
		}
		records = append(records, record)
	}

	return records, nil
}
