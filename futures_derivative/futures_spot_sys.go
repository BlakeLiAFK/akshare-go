package futures_derivative

import (
	"fmt"
	"strings"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/PuerkitoBio/goquery"
)

// getSysSpotFuturesDict 生意社-商品与期货-现期图: 品种和网址字典
func getSysSpotFuturesDict() (map[string]string, error) {
	url := "https://www.100ppi.com/sf/792.html"
	resp, err := utils.Get(url, nil)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(resp.String()))
	if err != nil {
		return nil, fmt.Errorf("解析HTML失败: %w", err)
	}

	nameURLDict := make(map[string]string)

	// 查找 class="q8" 的 div，然后提取所有 li 中的 a 标签
	doc.Find("div.q8 li a").Each(func(i int, s *goquery.Selection) {
		name := strings.TrimSpace(s.Text())
		href, exists := s.Attr("href")
		if exists && name != "" {
			nameURLDict[name] = href
		}
	})

	return nameURLDict, nil
}

// FuturesSpotSys 生意社-商品与期货-现期图
// 参数: symbol 期货品种，如 "铜"
//
//	indicator 指标类型，可选 "市场价格", "基差率", "主力基差"
//
// 返回: 现期图数据
func FuturesSpotSys(symbol, indicator string) ([]map[string]interface{}, error) {
	if symbol == "" {
		symbol = "铜"
	}
	if indicator == "" {
		indicator = "市场价格"
	}

	// 获取品种和URL映射
	nameURLDict, err := getSysSpotFuturesDict()
	if err != nil {
		return nil, fmt.Errorf("获取品种字典失败: %w", err)
	}

	relativeURL, exists := nameURLDict[symbol]
	if !exists {
		return nil, fmt.Errorf("未找到品种 %s", symbol)
	}

	url := "https://www.100ppi.com" + relativeURL
	resp, err := utils.Get(url, nil)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(resp.String()))
	if err != nil {
		return nil, fmt.Errorf("解析HTML失败: %w", err)
	}

	// 根据 indicator 选择不同的表格和字段
	var tableIndex int
	var headers []string
	switch indicator {
	case "市场价格":
		tableIndex = 1
		headers = []string{"日期", "现货价格", "主力合约", "最近合约"}
	case "基差率":
		tableIndex = 2
		headers = []string{"日期", "基差率"}
	case "主力基差":
		tableIndex = 3
		headers = []string{"日期", "主力基差"}
	default:
		return nil, fmt.Errorf("请输入正确的 indicator 参数: 市场价格, 基差率, 主力基差")
	}

	records := make([]map[string]interface{}, 0)

	// 查找指定索引的表格
	tableCount := 0
	doc.Find("table").Each(func(i int, table *goquery.Selection) {
		if tableCount == tableIndex {
			// 先提取所有列头（第一行）
			var columnHeaders []string
			table.Find("tr").First().Find("th, td").Each(func(j int, th *goquery.Selection) {
				columnHeaders = append(columnHeaders, strings.TrimSpace(th.Text()))
			})

			// 然后处理数据行（从第二行开始）
			table.Find("tr").Each(func(rowIdx int, row *goquery.Selection) {
				if rowIdx == 0 {
					return // 跳过表头
				}

				cells := row.Find("td")
				if cells.Length() < len(headers) {
					return
				}

				// 提取单元格数据
				cellData := make([]string, 0)
				cells.Each(func(cellIdx int, cell *goquery.Selection) {
					cellData = append(cellData, strings.TrimSpace(cell.Text()))
				})

				if len(cellData) < len(headers) {
					return
				}

				// 构建记录
				record := make(map[string]interface{})

				if indicator == "市场价格" {
					record[headers[0]] = cellData[0] // 日期
					record[headers[1]] = utils.MustFloat64(cellData[1])
					record[headers[2]] = utils.MustFloat64(cellData[2])
					record[headers[3]] = utils.MustFloat64(cellData[3])
				} else if indicator == "基差率" {
					record[headers[0]] = cellData[0] // 日期
					// 去除百分号
					rateStr := strings.ReplaceAll(cellData[1], "%", "")
					record[headers[1]] = utils.MustFloat64(rateStr)
				} else if indicator == "主力基差" {
					record[headers[0]] = cellData[0] // 日期
					record[headers[1]] = utils.MustFloat64(cellData[1])
				}

				records = append(records, record)
			})
		}
		tableCount++
	})

	return records, nil
}
