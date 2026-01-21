package bond

import (
	"fmt"
	"strings"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/PuerkitoBio/goquery"
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
	"github.com/tidwall/gjson"
)

// BondSpotQuote 中国外汇交易中心-现券市场做市报价
//
// 获取中国外汇交易中心暨全国银行间同业拆借中心的现券市场做市报价数据
//
// 参数:
//   - 无
//
// 返回:
//   - dataframe.DataFrame: 包含做市报价数据
//   - error: 错误信息
//
// 数据源: https://www.chinamoney.com.cn/chinese/mkdatabond/
func BondSpotQuote() (dataframe.DataFrame, error) {
	// 先调用 bond_china_close_return_map 注册服务（在实际使用中，可能需要先访问一次网站）
	// 这里简化处理，直接请求
	url := "https://www.chinamoney.com.cn/ags/ms/cm-u-md-bond/CbMktMakQuot"
	formData := map[string]string{
		"flag": "1",
		"lang": "cn",
	}

	resp, err := utils.PostForm(url, formData)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	// 解析JSON响应
	records := gjson.Get(resp.String(), "records")
	if !records.Exists() || !records.IsArray() {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到数据")
	}

	// 构建DataFrame记录
	var dfRecords []map[string]interface{}

	records.ForEach(func(_, item gjson.Result) bool {
		itemArray := item.Array()
		if len(itemArray) < 14 {
			return true
		}

		// 报价机构在索引2，债券简称在索引6
		// 买入/卖出收益率在索引11，买入/卖出净价在索引13
		institution := itemArray[2].String()
		bondName := itemArray[6].String()
		yieldStr := itemArray[11].String()
		priceStr := itemArray[13].String()

		// 分割买入/卖出数据
		yieldParts := strings.Split(yieldStr, "/")
		priceParts := strings.Split(priceStr, "/")

		buyYield := 0.0
		sellYield := 0.0
		buyPrice := 0.0
		sellPrice := 0.0

		if len(yieldParts) == 2 {
			buyYield = utils.MustParseFloat(strings.TrimSpace(yieldParts[0]))
			sellYield = utils.MustParseFloat(strings.TrimSpace(yieldParts[1]))
		}
		if len(priceParts) == 2 {
			buyPrice = utils.MustParseFloat(strings.TrimSpace(priceParts[0]))
			sellPrice = utils.MustParseFloat(strings.TrimSpace(priceParts[1]))
		}

		record := map[string]interface{}{
			"报价机构":  institution,
			"债券简称":  bondName,
			"买入净价":  buyPrice,
			"卖出净价":  sellPrice,
			"买入收益率": buyYield,
			"卖出收益率": sellYield,
		}
		dfRecords = append(dfRecords, record)
		return true
	})

	if len(dfRecords) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到有效数据")
	}

	df := dataframe.LoadMaps(dfRecords)
	return df, nil
}

// BondSpotDeal 中国外汇交易中心-现券市场成交行情
//
// 获取中国外汇交易中心暨全国银行间同业拆借中心的现券市场成交行情数据
//
// 参数:
//   - 无
//
// 返回:
//   - dataframe.DataFrame: 包含成交行情数据
//   - error: 错误信息
//
// 数据源: https://www.chinamoney.com.cn/chinese/mkdatabond/
func BondSpotDeal() (dataframe.DataFrame, error) {
	url := "https://www.chinamoney.com.cn/ags/ms/cm-u-md-bond/CbtPri"
	formData := map[string]string{
		"flag":     "1",
		"lang":     "cn",
		"bondName": "",
	}

	resp, err := utils.PostForm(url, formData)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	// 解析JSON响应
	records := gjson.Get(resp.String(), "records")
	if !records.Exists() || !records.IsArray() {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到数据")
	}

	// 构建DataFrame记录
	var dfRecords []map[string]interface{}

	records.ForEach(func(_, item gjson.Result) bool {
		itemArray := item.Array()
		if len(itemArray) < 18 {
			return true
		}

		// 根据Python代码的列映射:
		// 债券简称:2, 涨跌:7, 加权收益率:11, 成交净价:12, 最新收益率:15, 交易量:17
		bondName := itemArray[2].String()
		change := utils.MustParseFloat(itemArray[7].String())
		weightedYield := utils.MustParseFloat(itemArray[11].String())
		dealPrice := utils.MustParseFloat(itemArray[12].String())
		latestYield := utils.MustParseFloat(itemArray[15].String())
		volume := utils.MustParseFloat(itemArray[17].String())

		record := map[string]interface{}{
			"债券简称":  bondName,
			"成交净价":  dealPrice,
			"最新收益率": latestYield,
			"涨跌":    change,
			"加权收益率": weightedYield,
			"交易量":   volume,
		}
		dfRecords = append(dfRecords, record)
		return true
	})

	if len(dfRecords) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到有效数据")
	}

	df := dataframe.LoadMaps(dfRecords)
	return df, nil
}

// BondChinaYield 中国债券信息网-国债及其他债券收益率曲线
//
// 获取中国债券信息网的国债及其他债券收益率曲线数据
// 注意: endDate - startDate 应该小于一年
//
// 参数:
//   - startDate: 开始日期，格式: "20200204"
//   - endDate: 结束日期，格式: "20210124"
//
// 返回:
//   - dataframe.DataFrame: 包含收益率曲线数据
//   - error: 错误信息
//
// 数据源: https://yield.chinabond.com.cn/cbweb-pbc-web/pbc/historyQuery
func BondChinaYield(startDate, endDate string) (dataframe.DataFrame, error) {
	if len(startDate) != 8 || len(endDate) != 8 {
		return dataframe.DataFrame{}, fmt.Errorf("日期格式错误，应为YYYYMMDD")
	}

	url := "https://yield.chinabond.com.cn/cbweb-pbc-web/pbc/historyQuery"
	params := map[string]string{
		"startDate": fmt.Sprintf("%s-%s-%s", startDate[:4], startDate[4:6], startDate[6:]),
		"endDate":   fmt.Sprintf("%s-%s-%s", endDate[:4], endDate[4:6], endDate[6:]),
		"gjqx":      "0",
		"qxId":      "ycqx",
		"locale":    "cn_ZH",
	}

	headers := map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
	}

	resp, err := utils.GetWithHeaders(url, params, headers)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	// 替换&nbsp，Python版本这样做是为了清理HTML
	html := strings.ReplaceAll(resp.String(), "&nbsp", "")

	// 使用goquery解析HTML表格
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("解析HTML失败: %w", err)
	}

	// 查找所有table，Python代码使用了索引[1]取第二个表格
	tables := doc.Find("table")
	if tables.Length() < 2 {
		return dataframe.DataFrame{}, fmt.Errorf("未找到足够的表格")
	}

	// 使用第二个表格（索引1）
	table := tables.Eq(1)

	// 解析表格
	var records []map[string]interface{}
	var tableHeaders []string

	// 提取表头
	table.Find("thead tr th, thead tr td").Each(func(i int, s *goquery.Selection) {
		tableHeaders = append(tableHeaders, strings.TrimSpace(s.Text()))
	})

	// 如果thead中没有找到，尝试从第一行获取
	if len(tableHeaders) == 0 {
		table.Find("tr").First().Find("th, td").Each(func(i int, s *goquery.Selection) {
			tableHeaders = append(tableHeaders, strings.TrimSpace(s.Text()))
		})
	}

	// 提取数据行
	table.Find("tbody tr, tr").Each(func(i int, tr *goquery.Selection) {
		// 跳过表头行
		if tr.Find("th").Length() > 0 {
			return
		}

		var rowData []string
		tr.Find("td").Each(func(j int, td *goquery.Selection) {
			rowData = append(rowData, strings.TrimSpace(td.Text()))
		})

		if len(rowData) > 0 && len(rowData) >= len(tableHeaders) {
			record := make(map[string]interface{})
			for idx, header := range tableHeaders {
				if idx < len(rowData) {
					record[header] = rowData[idx]
				}
			}
			records = append(records, record)
		}
	})

	if len(records) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到有效数据")
	}

	df := dataframe.LoadMaps(records)

	// 转换数据类型（根据Python代码的字段名）
	if df.Names()[0] == "日期" {
		// 转换数值列
		for _, col := range []string{"3月", "6月", "1年", "3年", "5年", "7年", "10年", "30年"} {
			if contains(df.Names(), col) {
				df = df.Mutate(series.New(df.Col(col).Float(), series.Float, col))
			}
		}
	}

	return df, nil
}

// contains 检查字符串切片是否包含指定字符串
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
