package currency

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/PuerkitoBio/goquery"
	"github.com/tidwall/gjson"
	"github.com/xuri/excelize/v2"
)

// CurrencyBocSafe 人民币汇率中间价
//
// 从国家外汇管理局获取人民币汇率中间价数据
// 该接口会合并历史Excel数据和最新POST请求数据
//
// 返回:
//   - []map[string]any: 汇率数据列表，每个元素包含日期和各货币汇率
//   - error: 错误信息
//
// 示例:
//
//	rates, err := currency.CurrencyBocSafe()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, item := range rates {
//	    fmt.Printf("日期=%v\n", item["日期"])
//	    for k, v := range item {
//	        if k != "日期" {
//	            fmt.Printf("  %s: %v\n", k, v)
//	        }
//	    }
//	}
func CurrencyBocSafe() ([]map[string]any, error) {
	// 第一步：获取Excel文件URL
	pageURL := "https://www.safe.gov.cn/safe/2020/1218/17833.html"

	resp, err := http.Get(pageURL)
	if err != nil {
		return nil, fmt.Errorf("请求页面失败: %w", err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("解析页面失败: %w", err)
	}

	// 查找包含"人民币汇率"的链接
	var excelURL string
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		text := strings.TrimSpace(s.Text())
		if matched, _ := regexp.MatchString("人民币汇率", text); matched {
			if href, exists := s.Attr("href"); exists {
				excelURL = href
			}
		}
	})

	if excelURL == "" {
		return nil, fmt.Errorf("未找到Excel文件链接")
	}

	// 构建完整URL
	if !strings.HasPrefix(excelURL, "http") {
		excelURL = "https://www.safe.gov.cn" + excelURL
	}

	// 第二步：下载Excel文件
	excelResp, err := http.Get(excelURL)
	if err != nil {
		return nil, fmt.Errorf("下载Excel文件失败: %w", err)
	}
	defer excelResp.Body.Close()

	excelData, err := io.ReadAll(excelResp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取Excel文件失败: %w", err)
	}

	// 第三步：解析Excel文件
	f, err := excelize.OpenReader(bytes.NewReader(excelData))
	if err != nil {
		return nil, fmt.Errorf("打开Excel文件失败: %w", err)
	}
	defer f.Close()

	// 获取第一个工作表
	sheets := f.GetSheetList()
	if len(sheets) == 0 {
		return nil, fmt.Errorf("Excel文件无工作表")
	}

	rows, err := f.GetRows(sheets[0])
	if err != nil {
		return nil, fmt.Errorf("读取Excel数据失败: %w", err)
	}

	if len(rows) < 2 {
		return nil, fmt.Errorf("Excel数据不足")
	}

	// 解析表头
	headers := rows[0]

	// 解析数据行
	var historyData []map[string]any
	for i := 1; i < len(rows); i++ {
		row := rows[i]
		if len(row) == 0 {
			continue
		}

		item := make(map[string]any)
		for j, header := range headers {
			if j < len(row) {
				cellValue := strings.TrimSpace(row[j])
				if cellValue != "" {
					// 尝试解析为浮点数
					if floatVal := utils.MustFloat64(cellValue); floatVal != 0 || cellValue == "0" {
						item[header] = floatVal
					} else {
						// 尝试解析为日期
						if dateVal, err := time.Parse("2006-01-02", cellValue); err == nil {
							item[header] = dateVal
						} else {
							item[header] = cellValue
						}
					}
				}
			}
		}

		if len(item) > 0 {
			historyData = append(historyData, item)
		}
	}

	// 第四步：POST请求获取最新数据
	// 计算日期范围（从Excel最后日期到今天）
	now := time.Now()
	var startDate time.Time

	// 尝试从最后一条历史数据获取日期
	if len(historyData) > 0 {
		lastItem := historyData[len(historyData)-1]
		if dateVal, ok := lastItem["日期"].(time.Time); ok {
			startDate = dateVal.AddDate(0, 0, 1) // 从下一天开始
		}
	}

	// 如果没有历史数据，使用30天前
	if startDate.IsZero() {
		startDate = now.AddDate(0, 0, -30)
	}

	// POST请求最新数据
	postURL := "https://www.safe.gov.cn/AppStructured/hlw/RMBQuery.do"

	formData := url.Values{}
	formData.Set("startDate", startDate.Format("2006-01-02"))
	formData.Set("endDate", now.Format("2006-01-02"))
	formData.Set("queryYN", "true")

	client := &http.Client{}
	req, err := http.NewRequest("POST", postURL, strings.NewReader(formData.Encode()))
	if err != nil {
		// 如果POST失败，只返回历史数据
		return historyData, nil
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Referer", "https://www.safe.gov.cn/safe/2020/1218/17833.html")

	postResp, err := client.Do(req)
	if err != nil {
		// 如果POST失败，只返回历史数据
		return historyData, nil
	}
	defer postResp.Body.Close()

	postBody, err := io.ReadAll(postResp.Body)
	if err != nil {
		// 如果读取失败，只返回历史数据
		return historyData, nil
	}

	// 解析POST返回的JSON数据
	json := gjson.ParseBytes(postBody)
	dataArray := json.Get("data").Array()

	var recentData []map[string]any
	for _, data := range dataArray {
		item := make(map[string]any)

		data.ForEach(func(key, value gjson.Result) bool {
			keyStr := key.String()
			valueStr := value.String()

			// 解析日期
			if keyStr == "date" || keyStr == "日期" {
				if dateVal, err := time.Parse("2006-01-02", valueStr); err == nil {
					item["日期"] = dateVal
				} else {
					item["日期"] = valueStr
				}
			} else {
				// 尝试解析为浮点数
				if floatVal := utils.MustFloat64(valueStr); floatVal != 0 || valueStr == "0" {
					item[keyStr] = floatVal
				} else {
					item[keyStr] = valueStr
				}
			}

			return true
		})

		if len(item) > 0 {
			recentData = append(recentData, item)
		}
	}

	// 第五步：合并历史数据和最新数据
	allData := append(historyData, recentData...)

	return allData, nil
}
