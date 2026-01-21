package currency

import (
	"bytes"
	"fmt"
	"io"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// 货币代码映射缓存
var (
	currencyMapCache     map[string]string
	currencyMapCacheLock sync.Mutex
)

// getCurrencyBocSinaMap 获取货币名称到代码的映射
//
// 从新浪财经页面解析货币代码映射关系
//
// 参数:
//   - startDate: 开始日期，格式 "20230304"
//   - endDate: 结束日期，格式 "20231110"
//
// 返回:
//   - map[string]string: 货币名称到代码的映射
//   - error: 错误信息
func getCurrencyBocSinaMap(startDate, endDate string) (map[string]string, error) {
	// 检查缓存
	currencyMapCacheLock.Lock()
	if currencyMapCache != nil {
		defer currencyMapCacheLock.Unlock()
		return currencyMapCache, nil
	}
	currencyMapCacheLock.Unlock()

	url := "http://biz.finance.sina.com.cn/forex/forex.php"

	// 日期格式转换：20230304 -> 2023-03-04
	startFormatted := fmt.Sprintf("%s-%s-%s", startDate[:4], startDate[4:6], startDate[6:])
	endFormatted := fmt.Sprintf("%s-%s-%s", endDate[:4], endDate[4:6], endDate[6:])

	params := map[string]string{
		"startdate":  startFormatted,
		"enddate":    endFormatted,
		"money_code": "EUR",
		"type":       "0",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return nil, fmt.Errorf("请求货币映射失败: %w", err)
	}

	// GBK转UTF-8
	reader := transform.NewReader(bytes.NewReader(resp.Body()), simplifiedchinese.GBK.NewDecoder())
	utf8Body, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("编码转换失败: %w", err)
	}

	// 解析HTML
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(utf8Body))
	if err != nil {
		return nil, fmt.Errorf("解析HTML失败: %w", err)
	}

	// 提取货币代码映射
	result := make(map[string]string)
	doc.Find("#money_code option").Each(func(i int, s *goquery.Selection) {
		name := strings.TrimSpace(s.Text())
		code, exists := s.Attr("value")
		if exists && name != "" && code != "" {
			result[name] = code
		}
	})

	if len(result) == 0 {
		return nil, fmt.Errorf("未能获取货币代码映射")
	}

	// 缓存结果
	currencyMapCacheLock.Lock()
	currencyMapCache = result
	currencyMapCacheLock.Unlock()

	return result, nil
}

// CurrencyBocSina 中国银行人民币牌价历史数据查询
//
// 从新浪财经获取中国银行人民币外汇牌价历史数据
//
// 参数:
//   - symbol: 货币名称，如 "美元"、"欧元"、"英镑" 等
//     支持的货币: 美元、英镑、欧元、澳门元、泰国铢、菲律宾比索、港币、瑞士法郎、新加坡元、
//     瑞典克朗、丹麦克朗、挪威克朗、日元、加拿大元、澳大利亚元、新西兰元、韩国元
//   - startDate: 开始日期，格式 "20230304"
//   - endDate: 结束日期，格式 "20231110"
//
// 返回:
//   - []CurrencyBocSinaItem: 中行人民币牌价历史数据列表
//   - error: 错误信息
//
// 示例:
//
//	rates, err := currency.CurrencyBocSina("美元", "20230304", "20231110")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, item := range rates {
//	    fmt.Printf("日期=%s, 汇买价=%.4f, 中间价=%.4f\n",
//	        item.Date.Format("2006-01-02"), item.BuyingRate, item.MiddleRate)
//	}
func CurrencyBocSina(symbol, startDate, endDate string) ([]CurrencyBocSinaItem, error) {
	// 获取货币代码映射
	currencyMap, err := getCurrencyBocSinaMap(startDate, endDate)
	if err != nil {
		return nil, err
	}

	// 查找货币代码
	moneyCode, exists := currencyMap[symbol]
	if !exists {
		return nil, fmt.Errorf("不支持的货币: %s", symbol)
	}

	url := "http://biz.finance.sina.com.cn/forex/forex.php"

	// 日期格式转换
	startFormatted := fmt.Sprintf("%s-%s-%s", startDate[:4], startDate[4:6], startDate[6:])
	endFormatted := fmt.Sprintf("%s-%s-%s", endDate[:4], endDate[4:6], endDate[6:])

	params := map[string]string{
		"money_code": moneyCode,
		"type":       "0",
		"startdate":  startFormatted,
		"enddate":    endFormatted,
		"page":       "1",
		"call_type":  "ajax",
	}

	// 首次请求获取总页数
	resp, err := utils.Get(url, params)
	if err != nil {
		return nil, fmt.Errorf("请求数据失败: %w", err)
	}

	// GBK转UTF-8
	reader := transform.NewReader(bytes.NewReader(resp.Body()), simplifiedchinese.GBK.NewDecoder())
	utf8Body, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("编码转换失败: %w", err)
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(utf8Body))
	if err != nil {
		return nil, fmt.Errorf("解析HTML失败: %w", err)
	}

	// 获取总页数
	pageNum := 1
	pageElements := doc.Find("a.page")
	if pageElements.Length() > 0 {
		lastPageText := pageElements.Eq(pageElements.Length() - 2).Text()
		if lastPageNum, err := strconv.Atoi(strings.TrimSpace(lastPageText)); err == nil {
			pageNum = lastPageNum
		}
	}

	// 分页获取所有数据
	var allItems []CurrencyBocSinaItem

	for page := 1; page <= pageNum; page++ {
		params["page"] = strconv.Itoa(page)

		resp, err := utils.Get(url, params)
		if err != nil {
			return nil, fmt.Errorf("请求第%d页失败: %w", page, err)
		}

		// GBK转UTF-8
		reader := transform.NewReader(bytes.NewReader(resp.Body()), simplifiedchinese.GBK.NewDecoder())
		utf8Body, err := io.ReadAll(reader)
		if err != nil {
			return nil, fmt.Errorf("编码转换失败: %w", err)
		}

		doc, err := goquery.NewDocumentFromReader(bytes.NewReader(utf8Body))
		if err != nil {
			return nil, fmt.Errorf("解析第%d页HTML失败: %w", page, err)
		}

		// 解析表格数据
		// 注意：第一行是表头，从第二行开始是数据
		doc.Find("table tr").Each(func(i int, s *goquery.Selection) {
			// 跳过表头行
			if i == 0 {
				return
			}

			cells := s.Find("td")
			if cells.Length() < 6 {
				return
			}

			// 解析日期
			dateStr := strings.TrimSpace(cells.Eq(0).Text())
			date, _ := time.ParseInLocation("2006-01-02", dateStr, time.Local)

			// 解析各个价格
			buyingRate := utils.MustFloat64(strings.TrimSpace(cells.Eq(1).Text()))
			cashBuying := utils.MustFloat64(strings.TrimSpace(cells.Eq(2).Text()))
			sellingRate := utils.MustFloat64(strings.TrimSpace(cells.Eq(3).Text()))
			middleRate := utils.MustFloat64(strings.TrimSpace(cells.Eq(4).Text()))
			convertPrice := utils.MustFloat64(strings.TrimSpace(cells.Eq(5).Text()))

			allItems = append(allItems, CurrencyBocSinaItem{
				Date:         date,
				BuyingRate:   buyingRate,
				CashBuying:   cashBuying,
				SellingRate:  sellingRate,
				MiddleRate:   middleRate,
				ConvertPrice: convertPrice,
			})
		})
	}

	// 按日期排序（升序）
	// 数据已经是按日期升序，无需额外排序

	return allItems, nil
}
