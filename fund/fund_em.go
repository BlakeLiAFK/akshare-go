package fund

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/PuerkitoBio/goquery"
	"github.com/dop251/goja"
	"github.com/tidwall/gjson"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// FundPurchaseEm 获取基金申购状态
// https://fund.eastmoney.com/Fund_sgzt_bzdm.html
func FundPurchaseEm() ([]map[string]interface{}, error) {
	url := "https://fund.eastmoney.com/Data/Fund_JJJZ_Data.aspx"
	params := map[string]string{
		"t":    "8",
		"page": "1,50000",
		"js":   "reData",
		"sort": "fcode,asc",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	text := strings.TrimPrefix(resp.String(), "var reData=")
	result := gjson.Parse(text)
	datas := result.Get("datas").Array()

	records := make([]map[string]interface{}, 0, len(datas))
	for i, item := range datas {
		arr := strings.Split(item.String(), ",")
		if len(arr) < 14 {
			continue
		}

		record := map[string]interface{}{
			"序号":             i + 1,
			"基金代码":           arr[0],
			"基金简称":           arr[1],
			"基金类型":           arr[2],
			"最新净值/万份收益":      utils.MustFloat64(arr[3]),
			"最新净值/万份收益-报告时间": arr[4],
			"申购状态":           arr[5],
			"赎回状态":           arr[6],
			"下一开放日":          arr[7],
			"购买起点":           utils.MustFloat64(arr[8]),
			"日累计限定金额":        utils.MustFloat64(arr[9]),
			"手续费":            utils.MustFloat64(strings.TrimSuffix(arr[12], "%")),
		}
		records = append(records, record)
	}

	return records, nil
}

// FundNameEm 获取所有基金名称和类型
// https://fund.eastmoney.com/manager/default.html
func FundNameEm() ([]map[string]interface{}, error) {
	url := "https://fund.eastmoney.com/js/fundcode_search.js"

	resp, err := utils.Get(url, nil)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	text := resp.String()
	text = strings.TrimPrefix(text, "var r = ")
	text = strings.TrimSuffix(text, ";")

	result := gjson.Parse(text)
	items := result.Array()

	records := make([]map[string]interface{}, 0, len(items))
	for _, item := range items {
		arr := item.Array()
		if len(arr) < 5 {
			continue
		}

		record := map[string]interface{}{
			"基金代码": arr[0].String(),
			"拼音缩写": arr[1].String(),
			"基金简称": arr[2].String(),
			"基金类型": arr[3].String(),
			"拼音全称": arr[4].String(),
		}
		records = append(records, record)
	}

	return records, nil
}

// FundInfoIndexEm 获取指数型基金信息
// https://fund.eastmoney.com/trade/zs.html
// 参数: symbol (沪深指数/行业主题等), indicator (被动指数型/增强指数型)
func FundInfoIndexEm(symbol, indicator string) ([]map[string]interface{}, error) {
	if symbol == "" {
		symbol = "沪深指数"
	}
	if indicator == "" {
		indicator = "被动指数型"
	}

	symbolMap := map[string]string{
		"全部": "", "沪深指数": "053", "行业主题": "054",
		"大盘指数": "01", "中盘指数": "02", "小盘指数": "03",
		"股票指数": "050|001", "债券指数": "050|003",
	}

	indicatorMap := map[string]string{
		"全部": "", "被动指数型": "051", "增强指数型": "052",
	}

	url := "https://api.fund.eastmoney.com/FundTradeRank/GetRankList"
	params := map[string]string{
		"ft":   "zs",
		"sc":   "1n",
		"st":   "desc",
		"pi":   "1",
		"pn":   "10000",
		"fr":   symbolMap[symbol],
		"fr1":  indicatorMap[indicator],
		"fl":   "0",
		"isab": "1",
	}

	// 处理股票指数和债券指数
	if strings.Contains(symbolMap[symbol], "|") {
		parts := strings.Split(symbolMap[symbol], "|")
		params["fr"] = parts[0]
		params["ftype"] = parts[1]
	}

	headers := map[string]string{
		"Referer": "https://fund.eastmoney.com/",
	}

	resp, err := utils.GetWithHeaders(url, params, headers)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	result := gjson.ParseBytes(resp.Body())
	dataStr := result.Get("Data").String()
	dataResult := gjson.Parse(dataStr)
	datas := dataResult.Get("datas").Array()

	records := make([]map[string]interface{}, 0, len(datas))
	for _, item := range datas {
		parts := strings.Split(item.String(), "|")
		if len(parts) < 25 {
			continue
		}

		record := map[string]interface{}{
			"基金代码": parts[0],
			"基金名称": parts[1],
			"单位净值": utils.MustFloat64(parts[4]),
			"日期":   parts[3],
			"日增长率": utils.MustFloat64(parts[5]),
			"近1周":  utils.MustFloat64(parts[6]),
			"近1月":  utils.MustFloat64(parts[7]),
			"近3月":  utils.MustFloat64(parts[8]),
			"近6月":  utils.MustFloat64(parts[9]),
			"近1年":  utils.MustFloat64(parts[10]),
			"近2年":  utils.MustFloat64(parts[11]),
			"近3年":  utils.MustFloat64(parts[12]),
			"今年来":  utils.MustFloat64(parts[13]),
			"成立来":  utils.MustFloat64(parts[14]),
			"手续费":  utils.MustFloat64(parts[18]),
			"起购金额": parts[24],
			"跟踪标的": symbol,
			"跟踪方式": indicator,
		}
		records = append(records, record)
	}

	return records, nil
}

// FundOpenFundDailyEm 获取当前交易日所有开放式基金净值
// https://fund.eastmoney.com/fund.html
func FundOpenFundDailyEm() ([]map[string]interface{}, error) {
	url := "https://fund.eastmoney.com/Data/Fund_JJJZ_Data.aspx"
	params := map[string]string{
		"t":        "1",
		"lx":       "1",
		"sort":     "zdf,desc",
		"page":     "1,50000",
		"onlySale": "0",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	text := strings.TrimPrefix(resp.String(), "var db=")
	result := gjson.Parse(text)

	showday := result.Get("showday").Array()
	day0 := ""
	day1 := ""
	if len(showday) >= 2 {
		day0 = showday[0].String()
		day1 = showday[1].String()
	}

	datas := result.Get("datas").Array()
	records := make([]map[string]interface{}, 0, len(datas))

	for _, item := range datas {
		arr := strings.Split(item.String(), ",")
		if len(arr) < 18 {
			continue
		}

		record := map[string]interface{}{
			"基金代码":         arr[0],
			"基金简称":         arr[1],
			day0 + "-单位净值": utils.MustFloat64(arr[3]),
			day0 + "-累计净值": utils.MustFloat64(arr[4]),
			day1 + "-单位净值": utils.MustFloat64(arr[5]),
			day1 + "-累计净值": utils.MustFloat64(arr[6]),
			"日增长值":         utils.MustFloat64(arr[7]),
			"日增长率":         utils.MustFloat64(arr[8]),
			"申购状态":         arr[9],
			"赎回状态":         arr[10],
			"手续费":          utils.MustFloat64(arr[17]),
		}
		records = append(records, record)
	}

	return records, nil
}

// FundOpenFundInfoEm 东方财富网-天天基金网-基金数据-开放式基金净值
// https://fund.eastmoney.com/fund.html
// 参数:
//
//	symbol: 基金代码
//	indicator: 指标类型，可选值：
//	  "单位净值走势", "累计净值走势", "累计收益率走势",
//	  "同类排名走势", "同类排名百分比", "分红送配详情", "拆分详情"
//	period: 时间周期，可选值（仅"累计收益率走势"使用）：
//	  "1月", "3月", "6月", "1年", "3年", "5年", "今年来", "成立来"
func FundOpenFundInfoEm(symbol, indicator, period string) ([]map[string]interface{}, error) {
	if symbol == "" {
		symbol = "710001"
	}
	if indicator == "" {
		indicator = "单位净值走势"
	}
	if period == "" {
		period = "成立来"
	}

	// 单位净值走势
	if indicator == "单位净值走势" {
		url := fmt.Sprintf("https://fund.eastmoney.com/pingzhongdata/%s.js", symbol)
		resp, err := utils.Get(url, nil)
		if err != nil {
			return nil, fmt.Errorf("请求失败: %w", err)
		}

		// 使用goja执行JS
		vm := goja.New()
		_, err = vm.RunString(resp.String())
		if err != nil {
			return nil, fmt.Errorf("执行JS失败: %w", err)
		}

		// 获取Data_netWorthTrend变量
		dataValue := vm.Get("Data_netWorthTrend")
		if dataValue == nil || goja.IsUndefined(dataValue) {
			return []map[string]interface{}{}, nil
		}

		dataExport := dataValue.Export()
		dataArray, ok := dataExport.([]interface{})
		if !ok {
			return []map[string]interface{}{}, nil
		}

		records := make([]map[string]interface{}, 0, len(dataArray))
		for _, item := range dataArray {
			itemMap, ok := item.(map[string]interface{})
			if !ok {
				continue
			}

			// x是毫秒时间戳，转换为日期
			var dateStr string
			if xVal, ok := itemMap["x"]; ok {
				if xFloat, ok := xVal.(float64); ok {
					t := time.UnixMilli(int64(xFloat)).UTC()
					// 转换为上海时区
					shanghaiLoc, _ := time.LoadLocation("Asia/Shanghai")
					t = t.In(shanghaiLoc)
					dateStr = t.Format("2006-01-02")
				}
			}

			record := map[string]interface{}{
				"净值日期": dateStr,
				"单位净值": itemMap["y"],
				"日增长率": itemMap["equityReturn"],
			}
			records = append(records, record)
		}

		return records, nil
	}

	// 累计净值走势
	if indicator == "累计净值走势" {
		url := fmt.Sprintf("https://fund.eastmoney.com/pingzhongdata/%s.js", symbol)
		resp, err := utils.Get(url, nil)
		if err != nil {
			return nil, fmt.Errorf("请求失败: %w", err)
		}

		vm := goja.New()
		_, err = vm.RunString(resp.String())
		if err != nil {
			return nil, fmt.Errorf("执行JS失败: %w", err)
		}

		dataValue := vm.Get("Data_ACWorthTrend")
		if dataValue == nil || goja.IsUndefined(dataValue) {
			return []map[string]interface{}{}, nil
		}

		dataExport := dataValue.Export()
		dataArray, ok := dataExport.([]interface{})
		if !ok {
			return []map[string]interface{}{}, nil
		}

		records := make([]map[string]interface{}, 0, len(dataArray))
		for _, item := range dataArray {
			itemArray, ok := item.([]interface{})
			if !ok || len(itemArray) < 2 {
				continue
			}

			var dateStr string
			if xFloat, ok := itemArray[0].(float64); ok {
				t := time.UnixMilli(int64(xFloat)).UTC()
				shanghaiLoc, _ := time.LoadLocation("Asia/Shanghai")
				t = t.In(shanghaiLoc)
				dateStr = t.Format("2006-01-02")
			}

			record := map[string]interface{}{
				"净值日期": dateStr,
				"累计净值": itemArray[1],
			}
			records = append(records, record)
		}

		return records, nil
	}

	// 累计收益率走势（使用API，不需要JS）
	if indicator == "累计收益率走势" {
		url := "https://api.fund.eastmoney.com/pinzhong/LJSYLZS"
		periodMap := map[string]string{
			"1月":  "m",
			"3月":  "q",
			"6月":  "hy",
			"1年":  "y",
			"3年":  "try",
			"5年":  "fiy",
			"今年来": "sy",
			"成立来": "se",
		}

		params := map[string]string{
			"fundCode":  symbol,
			"indexcode": "000300",
			"type":      periodMap[period],
		}

		headers := map[string]string{
			"Referer": "https://fund.eastmoney.com/",
		}

		resp, err := utils.GetWithHeaders(url, params, headers)
		if err != nil {
			return nil, fmt.Errorf("请求失败: %w", err)
		}

		result := gjson.ParseBytes(resp.Body())
		items := result.Get("Data.0.data").Array()

		records := make([]map[string]interface{}, 0, len(items))
		for _, item := range items {
			dateVal := item.Get("0").Int()
			t := time.UnixMilli(dateVal).UTC()
			shanghaiLoc, _ := time.LoadLocation("Asia/Shanghai")
			t = t.In(shanghaiLoc)
			dateStr := t.Format("2006-01-02")

			record := map[string]interface{}{
				"日期":    dateStr,
				"累计收益率": item.Get("1").Float(),
			}
			records = append(records, record)
		}

		return records, nil
	}

	// 同类排名走势
	if indicator == "同类排名走势" {
		url := fmt.Sprintf("https://fund.eastmoney.com/pingzhongdata/%s.js", symbol)
		resp, err := utils.Get(url, nil)
		if err != nil {
			return nil, fmt.Errorf("请求失败: %w", err)
		}

		vm := goja.New()
		_, err = vm.RunString(resp.String())
		if err != nil {
			return nil, fmt.Errorf("执行JS失败: %w", err)
		}

		dataValue := vm.Get("Data_rateInSimilarType")
		if dataValue == nil || goja.IsUndefined(dataValue) {
			return []map[string]interface{}{}, nil
		}

		dataExport := dataValue.Export()
		dataArray, ok := dataExport.([]interface{})
		if !ok {
			return []map[string]interface{}{}, nil
		}

		records := make([]map[string]interface{}, 0, len(dataArray))
		for _, item := range dataArray {
			itemMap, ok := item.(map[string]interface{})
			if !ok {
				continue
			}

			var dateStr string
			if xVal, ok := itemMap["x"]; ok {
				if xFloat, ok := xVal.(float64); ok {
					t := time.UnixMilli(int64(xFloat)).UTC()
					shanghaiLoc, _ := time.LoadLocation("Asia/Shanghai")
					t = t.In(shanghaiLoc)
					dateStr = t.Format("2006-01-02")
				}
			}

			record := map[string]interface{}{
				"报告日期":          dateStr,
				"同类型排名-每日近三月排名": itemMap["sc"],
				"总排名-每日近三月排名":   itemMap["rank"],
			}
			records = append(records, record)
		}

		return records, nil
	}

	// 同类排名百分比
	if indicator == "同类排名百分比" {
		url := fmt.Sprintf("https://fund.eastmoney.com/pingzhongdata/%s.js", symbol)
		resp, err := utils.Get(url, nil)
		if err != nil {
			return nil, fmt.Errorf("请求失败: %w", err)
		}

		vm := goja.New()
		_, err = vm.RunString(resp.String())
		if err != nil {
			return nil, fmt.Errorf("执行JS失败: %w", err)
		}

		dataValue := vm.Get("Data_rateInSimilarPersent")
		if dataValue == nil || goja.IsUndefined(dataValue) {
			return []map[string]interface{}{}, nil
		}

		dataExport := dataValue.Export()
		dataArray, ok := dataExport.([]interface{})
		if !ok {
			return []map[string]interface{}{}, nil
		}

		records := make([]map[string]interface{}, 0, len(dataArray))
		for _, item := range dataArray {
			itemArray, ok := item.([]interface{})
			if !ok || len(itemArray) < 2 {
				continue
			}

			var dateStr string
			if xFloat, ok := itemArray[0].(float64); ok {
				t := time.UnixMilli(int64(xFloat)).UTC()
				shanghaiLoc, _ := time.LoadLocation("Asia/Shanghai")
				t = t.In(shanghaiLoc)
				dateStr = t.Format("2006-01-02")
			}

			record := map[string]interface{}{
				"报告日期": dateStr,
				"同类型排名-每日近3月收益排名百分比": itemArray[1],
			}
			records = append(records, record)
		}

		return records, nil
	}

	// 分红送配详情（使用HTML解析，不需要JS）
	if indicator == "分红送配详情" {
		url := fmt.Sprintf("https://fundf10.eastmoney.com/fhsp_%s.html", symbol)
		resp, err := utils.Get(url, nil)
		if err != nil {
			return nil, fmt.Errorf("请求失败: %w", err)
		}

		doc, err := goquery.NewDocumentFromReader(strings.NewReader(resp.String()))
		if err != nil {
			return nil, fmt.Errorf("解析HTML失败: %w", err)
		}

		// 查找分红表格
		tables := doc.Find("table.w782.comm.cfxq")
		if tables.Length() == 0 {
			return []map[string]interface{}{}, nil
		}

		var targetTable *goquery.Selection
		if tables.Length() >= 2 {
			targetTable = tables.Eq(1)
		} else {
			targetTable = tables.Eq(0)
		}

		// 检查是否有"暂无分红信息!"
		firstRow := targetTable.Find("tbody tr").First()
		if firstRow.Length() > 0 {
			// 检查整行文本
			rowText := firstRow.Text()
			if strings.Contains(rowText, "暂无分红信息") {
				return []map[string]interface{}{}, nil
			}
		}

		// 解析表头
		var headers []string
		targetTable.Find("thead tr th").Each(func(i int, th *goquery.Selection) {
			headers = append(headers, strings.TrimSpace(th.Text()))
		})

		// 解析数据行
		records := make([]map[string]interface{}, 0)
		targetTable.Find("tbody tr").Each(func(i int, tr *goquery.Selection) {
			record := make(map[string]interface{})
			tr.Find("td").Each(func(j int, td *goquery.Selection) {
				if j < len(headers) {
					record[headers[j]] = strings.TrimSpace(td.Text())
				}
			})
			if len(record) > 0 {
				// 再次检查是否包含"暂无"文本
				hasEmpty := false
				for _, val := range record {
					if str, ok := val.(string); ok && strings.Contains(str, "暂无分红信息") {
						hasEmpty = true
						break
					}
				}
				if !hasEmpty {
					records = append(records, record)
				}
			}
		})

		return records, nil
	}

	// 拆分详情（使用HTML解析，不需要JS）
	if indicator == "拆分详情" {
		url := fmt.Sprintf("https://fundf10.eastmoney.com/fhsp_%s.html", symbol)
		resp, err := utils.Get(url, nil)
		if err != nil {
			return nil, fmt.Errorf("请求失败: %w", err)
		}

		doc, err := goquery.NewDocumentFromReader(strings.NewReader(resp.String()))
		if err != nil {
			return nil, fmt.Errorf("解析HTML失败: %w", err)
		}

		// 查找拆分表格
		tables := doc.Find("table.w782.comm.cfxq")
		if tables.Length() == 0 {
			return []map[string]interface{}{}, nil
		}

		var targetTable *goquery.Selection
		if tables.Length() >= 3 {
			targetTable = tables.Eq(2)
		} else if tables.Length() >= 2 {
			targetTable = tables.Eq(1)
		} else {
			return []map[string]interface{}{}, nil
		}

		// 检查是否有"暂无拆分信息!"
		firstRow := targetTable.Find("tbody tr").First()
		if firstRow.Length() > 0 {
			rowText := firstRow.Text()
			if strings.Contains(rowText, "暂无拆分信息") {
				return []map[string]interface{}{}, nil
			}
		}

		// 解析表头
		var headers []string
		targetTable.Find("thead tr th").Each(func(i int, th *goquery.Selection) {
			headers = append(headers, strings.TrimSpace(th.Text()))
		})

		// 解析数据行
		records := make([]map[string]interface{}, 0)
		targetTable.Find("tbody tr").Each(func(i int, tr *goquery.Selection) {
			record := make(map[string]interface{})
			tr.Find("td").Each(func(j int, td *goquery.Selection) {
				if j < len(headers) {
					record[headers[j]] = strings.TrimSpace(td.Text())
				}
			})
			if len(record) > 0 {
				// 再次检查是否包含"暂无"文本
				hasEmpty := false
				for _, val := range record {
					if str, ok := val.(string); ok && strings.Contains(str, "暂无拆分信息") {
						hasEmpty = true
						break
					}
				}
				if !hasEmpty {
					records = append(records, record)
				}
			}
		})

		return records, nil
	}

	return []map[string]interface{}{}, nil
}

// FundIndividualDetailInfoEm Python源码中不存在此函数，可能是命名错误或已废弃
// 注：雪球有类似函数 fund_individual_detail_info_xq，应在fund_xq.go中实现
func FundIndividualDetailInfoEm(symbol, indicator string) (map[string]interface{}, error) {
	return nil, fmt.Errorf("此函数在Python akshare源码中不存在，可能是命名错误或已废弃")
}

func FundFinancialFundDailyEm() ([]map[string]interface{}, error) {
	url := "https://api.fund.eastmoney.com/FundNetValue/GetLCJJJZ"
	params := map[string]string{
		"letter":     "",
		"jjgsid":     "0",
		"searchtext": "",
		"sort":       "ljjz,desc",
		"page":       "1,10000",
	}

	headers := map[string]string{
		"Referer": "https://fund.eastmoney.com/lcjj.html",
	}

	resp, err := utils.GetWithHeaders(url, params, headers)
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
			"日期":   item.Get("RQ").String(),
			"万份收益": utils.MustFloat64(item.Get("DWSY").String()),
			"7日年化": utils.MustFloat64(item.Get("QRZNH").String()),
		}
		records = append(records, record)
	}

	return records, nil
}

// FundMoneyFundDailyEm 东方财富网-天天基金网-基金数据-货币型基金收益
// https://fund.eastmoney.com/HBJJ_pjsyl.html
func FundMoneyFundDailyEm() ([]map[string]interface{}, error) {
	url := "https://fund.eastmoney.com/HBJJ_pjsyl.html"

	resp, err := utils.GetWithHeaders(url, nil, map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
	})
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	// 使用gb2312解码
	reader := transform.NewReader(bytes.NewReader(resp.Body()), simplifiedchinese.GB18030.NewDecoder())
	bodyBytes, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("解码失败: %w", err)
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("解析HTML失败: %w", err)
	}

	// 找到第二个表格（索引1）
	tables := doc.Find("table")
	if tables.Length() < 2 {
		return nil, fmt.Errorf("未找到数据表格")
	}

	table := tables.Eq(1)

	// 提取显示日期（第一行的第6-11列）
	showDay := make([]string, 0)
	table.Find("tr").First().Find("td").Each(func(i int, s *goquery.Selection) {
		if i >= 5 && i < 11 {
			showDay = append(showDay, strings.TrimSpace(s.Text()))
		}
	})

	if len(showDay) < 6 {
		return nil, fmt.Errorf("未能提取日期信息")
	}

	records := make([]map[string]interface{}, 0)

	// 从第三行开始解析数据（跳过标题行）
	table.Find("tr").Each(func(rowIdx int, row *goquery.Selection) {
		if rowIdx < 2 { // 跳过前两行
			return
		}

		cells := row.Find("td")
		if cells.Length() < 14 {
			return
		}

		// 提取基金简称，去掉"基金吧档案"后缀
		fundName := strings.TrimSpace(cells.Eq(3).Text())
		fundName = strings.TrimSuffix(fundName, "基金吧档案")
		fundName = strings.TrimSuffix(fundName, "基金吧")
		fundName = strings.TrimSuffix(fundName, "档案")
		fundName = strings.TrimSpace(fundName)

		record := map[string]interface{}{
			"基金代码": strings.TrimSpace(cells.Eq(2).Text()),
			"基金简称": fundName,
			fmt.Sprintf("%s-万份收益", showDay[0]):   utils.MustFloat64(cells.Eq(4).Text()),
			fmt.Sprintf("%s-7日年化%%", showDay[1]): utils.MustFloat64(cells.Eq(5).Text()),
			fmt.Sprintf("%s-单位净值", showDay[2]):   utils.MustFloat64(cells.Eq(6).Text()),
			fmt.Sprintf("%s-万份收益", showDay[3]):   utils.MustFloat64(cells.Eq(7).Text()),
			fmt.Sprintf("%s-7日年化%%", showDay[4]): utils.MustFloat64(cells.Eq(8).Text()),
			fmt.Sprintf("%s-单位净值", showDay[5]):   utils.MustFloat64(cells.Eq(9).Text()),
			"日涨幅":  utils.MustFloat64(cells.Eq(10).Text()),
			"成立日期": strings.TrimSpace(cells.Eq(11).Text()),
			"基金经理": strings.TrimSpace(cells.Eq(12).Text()),
			"手续费":  strings.TrimSpace(cells.Eq(13).Text()),
			"可购全部": strings.TrimSpace(cells.Eq(14).Text()),
		}
		records = append(records, record)
	})

	return records, nil
}

func FundMoneyFundInfoEm(symbol string) ([]map[string]interface{}, error) {
	if symbol == "" {
		symbol = "000009"
	}

	url := "https://api.fund.eastmoney.com/f10/lsjz"
	params := map[string]string{
		"fundCode":  symbol,
		"pageIndex": "1",
		"pageSize":  "10000",
		"startDate": "",
		"endDate":   "",
	}

	headers := map[string]string{
		"Referer": fmt.Sprintf("https://fundf10.eastmoney.com/jjjz_%s.html", symbol),
	}

	resp, err := utils.GetWithHeaders(url, params, headers)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	result := gjson.ParseBytes(resp.Body())
	items := result.Get("Data.LSJZList").Array()

	records := make([]map[string]interface{}, 0, len(items))
	for _, item := range items {
		record := map[string]interface{}{
			"净值日期":    item.Get("FSRQ").String(),
			"每万份收益":   utils.MustFloat64(item.Get("DWJZ").String()),
			"7日年化收益率": utils.MustFloat64(item.Get("LJJZ").String()),
			"申购状态":    item.Get("SGZT").String(),
			"赎回状态":    item.Get("SHZT").String(),
		}
		records = append(records, record)
	}

	return records, nil
}

func FundValueEstimationEm(symbol string) ([]map[string]interface{}, error) {
	if symbol == "" {
		symbol = "全部"
	}

	symbolMap := map[string]string{
		"全部":     "1",
		"股票型":    "2",
		"混合型":    "3",
		"债券型":    "4",
		"指数型":    "5",
		"QDII":   "6",
		"ETF联接":  "7",
		"LOF":    "8",
		"场内交易基金": "9",
	}

	url := "https://api.fund.eastmoney.com/FundGuZhi/GetFundGZList"
	params := map[string]string{
		"type":      symbolMap[symbol],
		"sort":      "3",
		"orderType": "desc",
		"canbuy":    "0",
		"pageIndex": "1",
		"pageSize":  "20000",
	}

	headers := map[string]string{
		"Referer": "https://fund.eastmoney.com/",
	}

	resp, err := utils.GetWithHeaders(url, params, headers)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	result := gjson.ParseBytes(resp.Body())
	items := result.Get("Data.list").Array()

	records := make([]map[string]interface{}, 0, len(items))
	for _, item := range items {
		record := map[string]interface{}{
			"基金代码": item.Get("FCODE").String(),
			"基金简称": item.Get("SHORTNAME").String(),
			"估算时间": item.Get("GZTIME").String(),
			"单位净值": utils.MustFloat64(item.Get("DWJZ").String()),
			"估算值":  utils.MustFloat64(item.Get("GSZZL").String()),
			"估算涨跌": utils.MustFloat64(item.Get("GSZZL").String()),
			"日期":   item.Get("GZTIME").String(),
		}
		records = append(records, record)
	}

	return records, nil
}

// FundFinancialFundInfoEm 东方财富网-天天基金网-基金数据-理财型基金-历史净值明细
// https://fundf10.eastmoney.com/jjjz_000791.html
// 参数: symbol 理财型基金代码，可以通过 FundFinancialFundDailyEm() 来获取
func FundFinancialFundInfoEm(symbol string) ([]map[string]interface{}, error) {
	if symbol == "" {
		symbol = "000134"
	}

	url := "https://api.fund.eastmoney.com/f10/lsjz"
	params := map[string]string{
		"callback":  "jQuery18307915911837995662_1588249228826",
		"fundCode":  symbol,
		"pageIndex": "1",
		"pageSize":  "10000",
		"startDate": "",
		"endDate":   "",
		"_":         fmt.Sprintf("%d", time.Now().UnixMilli()),
	}

	headers := map[string]string{
		"Referer": fmt.Sprintf("https://fundf10.eastmoney.com/jjjz_%s.html", symbol),
	}

	resp, err := utils.GetWithHeaders(url, params, headers)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	// 提取JSON数据（去掉JSONP包装）
	body := resp.String()
	start := strings.Index(body, "{")
	end := strings.LastIndex(body, "}")
	if start == -1 || end == -1 {
		return []map[string]interface{}{}, nil
	}
	jsonData := body[start : end+1]

	result := gjson.Parse(jsonData)
	items := result.Get("Data.LSJZList").Array()

	records := make([]map[string]interface{}, 0, len(items))
	for _, item := range items {
		record := map[string]interface{}{
			"净值日期": item.Get("FSRQ").String(),
			"单位净值": utils.MustFloat64(item.Get("DWJZ").String()),
			"累计净值": utils.MustFloat64(item.Get("LJJZ").String()),
			"日增长率": utils.MustFloat64(item.Get("JZZZL").String()),
			"申购状态": item.Get("SGZT").String(),
			"赎回状态": item.Get("SHZT").String(),
			"分红送配": item.Get("FHSP").String(),
		}
		records = append(records, record)
	}

	return records, nil
}

// FundGradedFundDailyEm 东方财富网-天天基金网-基金数据-分级基金净值
// https://fund.eastmoney.com/fjjj.html#1_1__0__zdf,desc_1
// 返回: 当前交易日的所有分级基金净值
func FundGradedFundDailyEm() ([]map[string]interface{}, error) {
	url := "https://fund.eastmoney.com/Data/Fund_JJJZ_Data.aspx"
	params := map[string]string{
		"t":        "1",
		"lx":       "9",
		"letter":   "",
		"gsid":     "0",
		"text":     "",
		"sort":     "zdf,desc",
		"page":     "1,10000",
		"dt":       "1580914040623",
		"atfc":     "",
		"onlySale": "0",
	}

	headers := map[string]string{
		"Referer": "https://fund.eastmoney.com/fjjj.html",
	}

	resp, err := utils.GetWithHeaders(url, params, headers)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	// 解析 "var db=" 格式的数据
	body := resp.String()
	body = strings.TrimPrefix(body, "var db=")
	body = strings.TrimSpace(body)

	result := gjson.Parse(body)
	items := result.Get("datas").Array()
	showDay := result.Get("showday").Array()

	day0 := ""
	day1 := ""
	if len(showDay) >= 2 {
		day0 = showDay[0].String()
		day1 = showDay[1].String()
	}

	records := make([]map[string]interface{}, 0, len(items))
	for _, item := range items {
		parts := strings.Split(item.String(), ",")
		if len(parts) < 20 {
			continue
		}

		record := map[string]interface{}{
			"基金代码":                       parts[0],
			"基金简称":                       parts[1],
			fmt.Sprintf("%s-单位净值", day0): utils.MustFloat64(parts[3]),
			fmt.Sprintf("%s-累计净值", day0): utils.MustFloat64(parts[4]),
			fmt.Sprintf("%s-单位净值", day1): utils.MustFloat64(parts[5]),
			fmt.Sprintf("%s-累计净值", day1): utils.MustFloat64(parts[6]),
			"日增长值":                       utils.MustFloat64(parts[7]),
			"日增长率":                       utils.MustFloat64(parts[8]),
			"市价":                         utils.MustFloat64(parts[9]),
			"折价率":                        utils.MustFloat64(parts[10]),
			"手续费":                        utils.MustFloat64(parts[19]),
		}
		records = append(records, record)
	}

	return records, nil
}

// FundGradedFundInfoEm 东方财富网-天天基金网-基金数据-分级基金-历史净值明细
// https://fundf10.eastmoney.com/jjjz_150232.html
// 参数: symbol 分级基金代码，可以通过 FundGradedFundDailyEm() 来获取
func FundGradedFundInfoEm(symbol string) ([]map[string]interface{}, error) {
	if symbol == "" {
		symbol = "150232"
	}

	url := "https://api.fund.eastmoney.com/f10/lsjz"
	params := map[string]string{
		"fundCode":  symbol,
		"pageIndex": "1",
		"pageSize":  "20",
		"startDate": "",
		"endDate":   "",
		"_":         fmt.Sprintf("%d", time.Now().UnixMilli()),
	}

	headers := map[string]string{
		"Referer": fmt.Sprintf("https://fundf10.eastmoney.com/jjjz_%s.html", symbol),
	}

	// 第一次请求获取总页数
	resp, err := utils.GetWithHeaders(url, params, headers)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	result := gjson.ParseBytes(resp.Body())
	totalCount := result.Get("TotalCount").Int()
	totalPage := (totalCount + 19) / 20 // 向上取整

	records := make([]map[string]interface{}, 0)

	// 分页获取所有数据
	for page := int64(1); page <= totalPage; page++ {
		params["pageIndex"] = fmt.Sprintf("%d", page)
		resp, err = utils.GetWithHeaders(url, params, headers)
		if err != nil {
			continue
		}

		result = gjson.ParseBytes(resp.Body())
		items := result.Get("Data.LSJZList").Array()

		for _, item := range items {
			record := map[string]interface{}{
				"净值日期": item.Get("FSRQ").String(),
				"单位净值": utils.MustFloat64(item.Get("DWJZ").String()),
				"累计净值": utils.MustFloat64(item.Get("LJJZ").String()),
				"日增长率": utils.MustFloat64(item.Get("JZZZL").String()),
				"申购状态": item.Get("SGZT").String(),
				"赎回状态": item.Get("SHZT").String(),
			}
			records = append(records, record)
		}
	}

	return records, nil
}

// FundEtfFundDailyEm 东方财富网-天天基金网-基金数据-场内交易基金
// https://fund.eastmoney.com/cnjy_dwjz.html
// 返回: 当前交易日的所有场内交易基金数据
func FundEtfFundDailyEm() ([]map[string]interface{}, error) {
	url := "https://fund.eastmoney.com/cnjy_dwjz.html"
	resp, err := utils.Get(url, nil)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(resp.String()))
	if err != nil {
		return nil, fmt.Errorf("解析HTML失败: %w", err)
	}

	records := make([]map[string]interface{}, 0)

	// 查找表格
	doc.Find("table").Each(func(i int, table *goquery.Selection) {
		// 跳过表头
		table.Find("tbody tr").Each(func(j int, row *goquery.Selection) {
			if j == 0 {
				return // 跳过表头
			}

			cols := row.Find("td")
			if cols.Length() < 11 {
				return
			}

			record := map[string]interface{}{
				"基金代码": strings.TrimSpace(cols.Eq(0).Text()),
				"基金简称": strings.TrimSpace(cols.Eq(1).Text()),
				"类型":   strings.TrimSpace(cols.Eq(2).Text()),
				"单位净值": utils.MustFloat64(strings.TrimSpace(cols.Eq(3).Text())),
				"累计净值": utils.MustFloat64(strings.TrimSpace(cols.Eq(4).Text())),
				"增长值":  utils.MustFloat64(strings.TrimSpace(cols.Eq(7).Text())),
				"增长率":  utils.MustFloat64(strings.TrimSpace(cols.Eq(8).Text())),
				"市价":   utils.MustFloat64(strings.TrimSpace(cols.Eq(9).Text())),
				"折价率":  utils.MustFloat64(strings.TrimSpace(cols.Eq(10).Text())),
			}
			records = append(records, record)
		})
	})

	return records, nil
}

// FundEtfFundInfoEm 东方财富网-天天基金网-基金数据-场内交易基金-历史净值明细
// https://fundf10.eastmoney.com/jjjz_511280.html
// 参数:
//
//	symbol: 场内交易基金代码，可以通过 FundEtfFundDailyEm() 来获取
//	startDate: 开始统计时间，格式: "20000101"
//	endDate: 结束统计时间，格式: "20500101"
func FundEtfFundInfoEm(symbol, startDate, endDate string) ([]map[string]interface{}, error) {
	if symbol == "" {
		symbol = "511280"
	}
	if startDate == "" {
		startDate = "20000101"
	}
	if endDate == "" {
		endDate = "20500101"
	}

	// 格式化日期
	formatDate := func(date string) string {
		if len(date) != 8 {
			return date
		}
		return fmt.Sprintf("%s-%s-%s", date[:4], date[4:6], date[6:])
	}

	url := "https://api.fund.eastmoney.com/f10/lsjz"
	params := map[string]string{
		"fundCode":  symbol,
		"pageIndex": "1",
		"pageSize":  "20",
		"startDate": formatDate(startDate),
		"endDate":   formatDate(endDate),
		"_":         fmt.Sprintf("%d", time.Now().UnixMilli()),
	}

	headers := map[string]string{
		"Referer": fmt.Sprintf("https://fundf10.eastmoney.com/jjjz_%s.html", symbol),
	}

	// 第一次请求获取总页数
	resp, err := utils.GetWithHeaders(url, params, headers)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	result := gjson.ParseBytes(resp.Body())
	totalCount := result.Get("TotalCount").Int()
	totalPage := (totalCount + 19) / 20 // 向上取整

	records := make([]map[string]interface{}, 0)

	// 分页获取所有数据
	for page := int64(1); page <= totalPage; page++ {
		params["pageIndex"] = fmt.Sprintf("%d", page)
		resp, err = utils.GetWithHeaders(url, params, headers)
		if err != nil {
			continue
		}

		result = gjson.ParseBytes(resp.Body())
		items := result.Get("Data.LSJZList").Array()

		for _, item := range items {
			record := map[string]interface{}{
				"净值日期": item.Get("FSRQ").String(),
				"单位净值": utils.MustFloat64(item.Get("DWJZ").String()),
				"累计净值": utils.MustFloat64(item.Get("LJJZ").String()),
				"日增长率": utils.MustFloat64(item.Get("JZZZL").String()),
				"申购状态": item.Get("SGZT").String(),
				"赎回状态": item.Get("SHZT").String(),
			}
			records = append(records, record)
		}
	}

	return records, nil
}

// FundHkFundHistEm 东方财富网-天天基金网-基金数据-香港基金-历史净值明细(分红送配详情)
// https://overseas.1234567.com.cn/f10/FundJz/968092#FHPS
// 参数:
//
//	code: 香港基金代码，通过 FundHkRankEm() 获取
//	indicator: 指标类型，可选值: "历史净值明细", "分红送配详情"
func FundHkFundHistEm(code, indicator string) ([]map[string]interface{}, error) {
	if code == "" {
		code = "1002200683"
	}
	if indicator == "" {
		indicator = "历史净值明细"
	}

	url := "https://overseas.1234567.com.cn/overseasapi/OpenApiHander.ashx"
	params := map[string]string{
		"api":       "HKFDApi",
		"m":         "MethodJZ",
		"hkfcode":   code,
		"pageindex": "0",
		"pagesize":  "1000",
		"date1":     "",
		"date2":     "",
	}

	if indicator == "历史净值明细" {
		params["action"] = "2"
	} else {
		params["action"] = "3"
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	result := gjson.ParseBytes(resp.Body())
	items := result.Get("Data").Array()

	records := make([]map[string]interface{}, 0, len(items))

	if indicator == "历史净值明细" {
		for _, item := range items {
			arr := item.Array()
			if len(arr) < 10 {
				continue
			}

			record := map[string]interface{}{
				"净值日期": arr[3].String(),
				"单位净值": utils.MustFloat64(arr[4].String()),
				"日增长值": utils.MustFloat64(arr[6].String()),
				"日增长率": utils.MustFloat64(arr[7].String()),
				"单位":   arr[9].String(),
			}
			records = append(records, record)
		}
	} else {
		// 分红送配详情
		for _, item := range items {
			arr := item.Array()
			if len(arr) < 7 {
				continue
			}

			record := map[string]interface{}{
				"权益登记日": arr[3].String(),
				"除息日":   arr[4].String(),
				"派息":    arr[5].String(),
				"单位":    arr[6].String(),
			}
			records = append(records, record)
		}
	}

	return records, nil
}

// FundHqTimeEm Python源码中不存在此函数，可能是命名错误或已废弃
func FundHqTimeEm(symbol string) ([]map[string]interface{}, error) {
	return nil, fmt.Errorf("此函数在Python akshare源码中不存在，可能是命名错误或已废弃")
}
