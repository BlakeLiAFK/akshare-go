package bond

import (
	"fmt"
	"strings"
	"time"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/PuerkitoBio/goquery"
	"github.com/go-gota/gota/dataframe"
	"github.com/tidwall/gjson"
)

// BondCBJSL 集思录可转债列表
//
// 获取集思录网站的可转债列表数据，包括现价、涨跌幅、转股价值等信息
//
// 参数:
//   - cookie: 可选的浏览器 cookie，默认为空
//
// 返回:
//   - dataframe.DataFrame: 包含可转债列表数据
//   - error: 错误信息
//
// 数据源: https://www.jisilu.cn/data/cbnew/#cb
func BondCBJSL(cookie string) (dataframe.DataFrame, error) {
	url := "https://www.jisilu.cn/data/cbnew/cb_list_new/"

	headers := map[string]string{
		"Accept":           "application/json, text/javascript, */*; q=0.01",
		"Accept-Encoding":  "gzip, deflate, br",
		"Accept-Language":  "zh-CN,zh;q=0.9,en;q=0.8",
		"Cache-Control":    "no-cache",
		"Content-Type":     "application/x-www-form-urlencoded; charset=UTF-8",
		"Origin":           "https://www.jisilu.cn",
		"Pragma":           "no-cache",
		"Referer":          "https://www.jisilu.cn/data/cbnew/",
		"User-Agent":       "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.164 Safari/537.36",
		"X-Requested-With": "XMLHttpRequest",
	}

	if cookie != "" {
		headers["Cookie"] = cookie
	}

	// 添加时间戳参数
	params := map[string]string{
		"___jsl": fmt.Sprintf("LST___t=%d", time.Now().UnixMilli()),
	}

	// 构建 payload - 使用表单数据格式
	formData := map[string]string{
		"fprice":       "",
		"tprice":       "",
		"curr_iss_amt": "",
		"volume":       "",
		"svolume":      "",
		"premium_rt":   "",
		"ytm_rt":       "",
		"market":       "",
		"rating_cd":    "",
		"is_search":    "N",
		"btype":        "",
		"listed":       "Y",
		"qflag":        "N",
		"sw_cd":        "",
		"bond_ids":     "",
		"rp":           "50",
	}

	resp, err := utils.PostFormWithParams(url, params, formData, headers, map[string][]string{
		"market_cd[]": {"shmb", "shkc", "szmb", "szcy"},
	})
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求集思录可转债列表失败: %w", err)
	}

	// 解析 JSON 响应
	json := gjson.ParseBytes(resp.Body())

	rows := json.Get("rows")
	if !rows.Exists() || !rows.IsArray() {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到可转债数据")
	}

	// 构建 DataFrame 记录
	var records []map[string]interface{}

	rows.ForEach(func(_, row gjson.Result) bool {
		cell := row.Get("cell")
		if !cell.Exists() {
			return true
		}

		record := map[string]interface{}{
			"代码":     cell.Get("bond_id").String(),
			"转债名称":   cell.Get("bond_nm").String(),
			"现价":     utils.MustParseFloat(cell.Get("price").String()),
			"涨跌幅":    utils.MustParseFloat(cell.Get("increase_rt").String()),
			"正股代码":   cell.Get("stock_id").String(),
			"正股名称":   cell.Get("stock_nm").String(),
			"正股价":    utils.MustParseFloat(cell.Get("sprice").String()),
			"正股涨跌":   utils.MustParseFloat(cell.Get("sincrease_rt").String()),
			"正股PB":   utils.MustParseFloat(cell.Get("pb").String()),
			"转股价":    utils.MustParseFloat(cell.Get("convert_price").String()),
			"转股价值":   utils.MustParseFloat(cell.Get("convert_value").String()),
			"转股溢价率":  utils.MustParseFloat(cell.Get("premium_rt").String()),
			"债券评级":   cell.Get("rating_cd").String(),
			"回售触发价":  utils.MustParseFloat(cell.Get("put_convert_price").String()),
			"强赎触发价":  utils.MustParseFloat(cell.Get("force_redeem_price").String()),
			"转债占比":   utils.MustParseFloat(cell.Get("convert_amt_ratio").String()),
			"到期时间":   parseDate(cell.Get("maturity_dt").String()),
			"剩余年限":   utils.MustParseFloat(cell.Get("year_left").String()),
			"剩余规模":   utils.MustParseFloat(cell.Get("curr_iss_amt").String()),
			"成交额":    utils.MustParseFloat(cell.Get("volume").String()),
			"换手率":    utils.MustParseFloat(cell.Get("turnover_rt").String()),
			"到期税前收益": utils.MustParseFloat(cell.Get("ytm_rt").String()),
			"双低":     utils.MustParseFloat(cell.Get("dblow").String()),
		}

		records = append(records, record)
		return true
	})

	if len(records) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到可转债数据")
	}

	df := dataframe.LoadMaps(records)
	return df, nil
}

// BondCBRedeemJSL 集思录可转债强赎数据
//
// 获取集思录网站的可转债强赎数据，包括强赎触发价、强赎状态等信息
//
// 参数:
//   - 无
//
// 返回:
//   - dataframe.DataFrame: 包含可转债强赎数据
//   - error: 错误信息
//
// 数据源: https://www.jisilu.cn/data/cbnew/#redeem
func BondCBRedeemJSL() (dataframe.DataFrame, error) {
	url := "https://www.jisilu.cn/data/cbnew/redeem_list/"

	headers := map[string]string{
		"Accept":           "application/json, text/javascript, */*; q=0.01",
		"Accept-Encoding":  "gzip, deflate, br",
		"Accept-Language":  "zh-CN,zh;q=0.9,en;q=0.8",
		"Cache-Control":    "no-cache",
		"Content-Type":     "application/x-www-form-urlencoded; charset=UTF-8",
		"Origin":           "https://www.jisilu.cn",
		"Pragma":           "no-cache",
		"Referer":          "https://www.jisilu.cn/data/cbnew/",
		"User-Agent":       "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.67 Safari/537.36",
		"X-Requested-With": "XMLHttpRequest",
	}

	params := map[string]string{
		"___jsl": fmt.Sprintf("LST___t=%d", time.Now().UnixMilli()),
	}

	formData := map[string]string{
		"rp": "50",
	}

	resp, err := utils.PostFormWithParams(url, params, formData, headers, nil)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求集思录可转债强赎数据失败: %w", err)
	}

	// 解析 JSON 响应
	json := gjson.ParseBytes(resp.Body())

	rows := json.Get("rows")
	if !rows.Exists() || !rows.IsArray() {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到强赎数据")
	}

	// 构建 DataFrame 记录
	var records []map[string]interface{}

	rows.ForEach(func(_, row gjson.Result) bool {
		cell := row.Get("cell")
		if !cell.Exists() {
			return true
		}

		// 处理强赎触发比（去除百分号）
		redeemRatio := cell.Get("redeem_price_ratio").String()
		redeemRatio = strings.TrimSuffix(strings.TrimSpace(redeemRatio), "%")

		// 处理强赎状态映射
		redeemIcon := cell.Get("redeem_icon").String()
		redeemStatus := mapRedeemStatus(redeemIcon)

		// 处理强赎天计数（提取 x/y | z 格式）
		redeemCount := cell.Get("redeem_count").String()
		redeemCount = extractRedeemCount(redeemCount)

		record := map[string]interface{}{
			"代码":    cell.Get("bond_id").String(),
			"名称":    cell.Get("bond_nm").String(),
			"现价":    utils.MustParseFloat(cell.Get("price").String()),
			"正股代码":  cell.Get("stock_id").String(),
			"正股名称":  cell.Get("stock_nm").String(),
			"规模":    utils.MustParseFloat(cell.Get("orig_iss_amt").String()),
			"剩余规模":  utils.MustParseFloat(cell.Get("curr_iss_amt").String()),
			"转股起始日": parseDate(cell.Get("convert_dt").String()),
			"最后交易日": parseDate(cell.Get("delist_dt").String()),
			"到期日":   parseDate(cell.Get("maturity_dt").String()),
			"转股价":   utils.MustParseFloat(cell.Get("convert_price").String()),
			"强赎触发比": utils.MustParseFloat(redeemRatio),
			"强赎触发价": utils.MustParseFloat(cell.Get("force_redeem_price").String()),
			"正股价":   utils.MustParseFloat(cell.Get("sprice").String()),
			"强赎价":   utils.MustParseFloat(cell.Get("real_force_redeem_price").String()),
			"强赎天计数": redeemCount,
			"强赎条款":  cell.Get("redeem_tc").String(),
			"强赎状态":  redeemStatus,
		}

		records = append(records, record)
		return true
	})

	if len(records) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到强赎数据")
	}

	df := dataframe.LoadMaps(records)
	return df, nil
}

// BondCBAdjLogsJSL 集思录可转债转股价调整记录
//
// 获取集思录网站的可转债转股价调整记录
//
// 参数:
//   - symbol: 可转债代码，如 "128013"
//
// 返回:
//   - dataframe.DataFrame: 包含转股价调整记录，如果没有记录则返回空 DataFrame
//   - error: 错误信息
//
// 数据源: https://www.jisilu.cn/data/cbnew/#cb
func BondCBAdjLogsJSL(symbol string) (dataframe.DataFrame, error) {
	if symbol == "" {
		return dataframe.DataFrame{}, fmt.Errorf("可转债代码不能为空")
	}

	url := fmt.Sprintf("https://www.jisilu.cn/data/cbnew/adj_logs/?bond_id=%s", symbol)

	resp, err := utils.Get(url, nil)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求集思录转股价调整记录失败: %w", err)
	}

	htmlContent := resp.String()

	// 检查是否包含表格数据
	if !strings.Contains(htmlContent, "</table>") {
		// 没有表格数据，返回空 DataFrame
		// 可能的情况：
		// 1. 该可转债没有转股价调整记录，返回文本 '暂无数据'
		// 2. 无效可转债代码，返回 {"timestamp":xxx,"isError":1,"msg":"无效代码格式"}
		return dataframe.DataFrame{}, nil
	}

	// 使用 goquery 解析 HTML
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("解析HTML失败: %w", err)
	}

	// 查找表格
	table := doc.Find("table").First()
	if table.Length() == 0 {
		return dataframe.DataFrame{}, nil
	}

	// 提取表头
	var headers []string
	table.Find("thead tr th").Each(func(i int, th *goquery.Selection) {
		header := strings.TrimSpace(th.Text())
		header = strings.ReplaceAll(header, " ", "")
		headers = append(headers, header)
	})

	if len(headers) == 0 {
		// 尝试从第一行 td 提取表头
		table.Find("tr").First().Find("td").Each(func(i int, td *goquery.Selection) {
			header := strings.TrimSpace(td.Text())
			header = strings.ReplaceAll(header, " ", "")
			headers = append(headers, header)
		})
	}

	// 提取数据行
	var records []map[string]interface{}
	table.Find("tbody tr").Each(func(i int, tr *goquery.Selection) {
		record := make(map[string]interface{})
		tr.Find("td").Each(func(j int, td *goquery.Selection) {
			if j < len(headers) {
				value := strings.TrimSpace(td.Text())
				header := headers[j]

				// 根据列名进行类型转换
				switch header {
				case "下修前转股价", "下修后转股价", "下修底价":
					record[header] = utils.MustParseFloat(value)
				case "股东大会日", "新转股价生效日期":
					record[header] = parseDate(value)
				default:
					record[header] = value
				}
			}
		})
		if len(record) > 0 {
			records = append(records, record)
		}
	})

	if len(records) == 0 {
		return dataframe.DataFrame{}, nil
	}

	df := dataframe.LoadMaps(records)
	return df, nil
}

// parseDate 解析日期字符串，支持多种格式
func parseDate(dateStr string) string {
	dateStr = strings.TrimSpace(dateStr)
	if dateStr == "" || dateStr == "-" {
		return ""
	}

	// 尝试解析多种日期格式
	layouts := []string{
		"2006-01-02",
		"2006/01/02",
		"20060102",
	}

	for _, layout := range layouts {
		if t, err := time.Parse(layout, dateStr); err == nil {
			return t.Format("2006-01-02")
		}
	}

	// 如果无法解析，返回原始字符串
	return dateStr
}

// mapRedeemStatus 映射强赎状态
func mapRedeemStatus(icon string) string {
	statusMap := map[string]string{
		"R": "已公告强赎",
		"O": "公告要强赎",
		"G": "公告不强赎",
		"B": "已满足强赎条件",
	}

	if status, ok := statusMap[icon]; ok {
		return status
	}
	return ""
}

// extractRedeemCount 提取强赎天计数（格式: x/y | z）
func extractRedeemCount(count string) string {
	count = strings.TrimSpace(count)
	if count == "" {
		return ""
	}

	// 使用正则表达式提取 x/y | z 格式
	// 简化处理：直接返回包含数字和 / | 的部分
	// 更精确的实现可以使用 regexp 包
	return count
}
