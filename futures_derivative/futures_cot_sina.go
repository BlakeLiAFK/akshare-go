package futures_derivative

import (
	"fmt"
	"strings"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/PuerkitoBio/goquery"
)

// FuturesHoldPosSina 新浪财经-期货-成交持仓
// 参数: symbolType 数据类型，可选 "成交量", "多单持仓", "空单持仓"
//
//	contract 期货合约，如 "IC2403", "OI2501"
//	date 查询日期，格式 "20240223"
//
// 返回: 成交持仓数据
func FuturesHoldPosSina(symbolType, contract, date string) ([]map[string]interface{}, error) {
	if symbolType == "" {
		symbolType = "成交量"
	}
	if contract == "" {
		contract = "OI2501"
	}
	if date == "" {
		date = "20240223"
	}

	// 转换日期格式: 20240223 -> 2024-02-23
	if len(date) == 8 {
		date = date[:4] + "-" + date[4:6] + "-" + date[6:]
	}

	url := "https://vip.stock.finance.sina.com.cn/q/view/vFutures_Positions_cjcc.php"
	params := map[string]string{
		"t_breed": contract,
		"t_date":  date,
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(resp.String()))
	if err != nil {
		return nil, fmt.Errorf("解析HTML失败: %w", err)
	}

	// 根据 symbolType 选择不同的表格
	var tableIndex int
	var headers []string
	switch symbolType {
	case "成交量":
		tableIndex = 2
		headers = []string{"名次", "会员简称", "成交量", "比上交易增减"}
	case "多单持仓":
		tableIndex = 3
		headers = []string{"名次", "会员简称", "多单持仓", "比上交易增减"}
	case "空单持仓":
		tableIndex = 4
		headers = []string{"名次", "会员简称", "空单持仓", "比上交易增减"}
	default:
		return nil, fmt.Errorf("请输入正确的 symbolType 参数: 成交量, 多单持仓, 空单持仓")
	}

	records := make([]map[string]interface{}, 0)

	// 查找指定索引的表格
	tableCount := 0
	doc.Find("table").Each(func(i int, table *goquery.Selection) {
		if tableCount == tableIndex {
			// 跳过表头，处理数据行
			table.Find("tr").Each(func(rowIdx int, row *goquery.Selection) {
				if rowIdx == 0 {
					return // 跳过表头
				}

				cells := row.Find("td")
				if cells.Length() < 4 {
					return
				}

				// 提取单元格数据
				cellData := make([]string, 0)
				cells.Each(func(cellIdx int, cell *goquery.Selection) {
					cellData = append(cellData, strings.TrimSpace(cell.Text()))
				})

				// 检查是否为合计行（最后一行），跳过
				if len(cellData) > 0 && strings.Contains(cellData[0], "合计") {
					return
				}

				// 构建记录
				if len(cellData) >= 4 {
					record := map[string]interface{}{
						headers[0]: utils.MustInt64(cellData[0]),
						headers[1]: cellData[1],
						headers[2]: utils.MustInt64(cellData[2]),
						headers[3]: utils.MustInt64(cellData[3]),
					}
					records = append(records, record)
				}
			})
		}
		tableCount++
	})

	return records, nil
}
