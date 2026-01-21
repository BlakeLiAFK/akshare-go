package futures

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/PuerkitoBio/goquery"
)

// FuturesContractDetail 期货合约详情
type FuturesContractDetail struct {
	Item  string `json:"item"`  // 项目
	Value string `json:"value"` // 值
}

// FuturesContractDetailFunc 查询期货合约详情-新浪
//
// 数据源: https://finance.sina.com.cn/futures/quotes/V2101.shtml
//
// 参数:
//   - symbol: 合约代码，如 "AP2101"
//
// 返回:
//   - []FuturesContractDetail: 期货合约详情
//   - error: 错误信息
func FuturesContractDetailFunc(symbol string) ([]FuturesContractDetail, error) {
	url := fmt.Sprintf("https://finance.sina.com.cn/futures/quotes/%s.shtml", symbol)

	resp, err := utils.GetWithHeaders(url, nil, map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
	})
	if err != nil {
		return nil, fmt.Errorf("请求合约详情失败: %w", err)
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(resp.Body())))
	if err != nil {
		return nil, fmt.Errorf("解析HTML失败: %w", err)
	}

	var result []FuturesContractDetail
	// 查找第7个表格（索引6）
	tableIndex := 0
	doc.Find("table").Each(func(i int, table *goquery.Selection) {
		if tableIndex == 6 {
			table.Find("tr").Each(func(_ int, tr *goquery.Selection) {
				tds := tr.Find("td")
				// 每行可能有多对 item-value
				for j := 0; j < tds.Length()-1; j += 2 {
					item := strings.TrimSpace(tds.Eq(j).Text())
					value := strings.TrimSpace(tds.Eq(j + 1).Text())
					if item != "" {
						result = append(result, FuturesContractDetail{
							Item:  item,
							Value: value,
						})
					}
				}
			})
		}
		tableIndex++
	})

	return result, nil
}

// contractDetailEMResponse 东方财富合约详情API响应
type contractDetailEMResponse map[string]string

// FuturesContractDetailEMFunc 查询期货合约详情-东方财富
//
// 数据源: https://quote.eastmoney.com/qihuo/v2602F.html
//
// 参数:
//   - symbol: 合约代码，如 "v2602F"
//
// 返回:
//   - []FuturesContractDetail: 期货合约详情
//   - error: 错误信息
func FuturesContractDetailEMFunc(symbol string) ([]FuturesContractDetail, error) {
	// 首先获取页面，提取内部代码
	pageURL := fmt.Sprintf("https://quote.eastmoney.com/qihuo/%s.html", symbol)
	resp, err := utils.GetWithHeaders(pageURL, nil, map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
	})
	if err != nil {
		return nil, fmt.Errorf("请求合约页面失败: %w", err)
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(resp.Body())))
	if err != nil {
		return nil, fmt.Errorf("解析HTML失败: %w", err)
	}

	// 查找内部代码
	var innerSymbol string
	doc.Find("div.sidertabbox_tsplit div.onet a").Each(func(_ int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists && strings.Contains(href, "#futures_") {
			re := regexp.MustCompile(`#futures_(\w+)`)
			matches := re.FindStringSubmatch(href)
			if len(matches) > 1 {
				innerSymbol = matches[1]
			}
		}
	})

	if innerSymbol == "" {
		// 尝试从symbol中提取
		innerSymbol = strings.ToLower(symbol)
	}

	// 获取合约详情数据
	infoURL := fmt.Sprintf("https://futsse-static.eastmoney.com/redis?msgid=%s_info", innerSymbol)
	resp, err = utils.GetWithHeaders(infoURL, nil, map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
	})
	if err != nil {
		return nil, fmt.Errorf("请求合约详情数据失败: %w", err)
	}

	var dataMap contractDetailEMResponse
	if err := json.Unmarshal(resp.Body(), &dataMap); err != nil {
		return nil, fmt.Errorf("解析合约详情数据失败: %w", err)
	}

	// 字段映射
	columnMapping := map[string]string{
		"vname":   "交易品种",
		"vcode":   "交易代码",
		"jydw":    "交易单位",
		"bjdw":    "报价单位",
		"market":  "上市交易所",
		"zxbddw":  "最小变动价格",
		"zdtbfd":  "跌涨停板幅度",
		"hyjgyf":  "合约交割月份",
		"jysj":    "交易时间",
		"zhjyr":   "最后交易日",
		"zhjgr":   "最后交割日",
		"jgpj":    "交割品级",
		"zcjybzj": "最初交易保证金",
		"jgfs":    "交割方式",
	}

	var result []FuturesContractDetail
	for key, value := range dataMap {
		item := key
		if mappedName, ok := columnMapping[key]; ok {
			item = mappedName
		}
		result = append(result, FuturesContractDetail{
			Item:  item,
			Value: value,
		})
	}

	return result, nil
}
