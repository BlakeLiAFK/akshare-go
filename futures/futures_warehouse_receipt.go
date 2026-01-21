package futures

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/xuri/excelize/v2"
)

// WarehouseReceiptCZCE 郑商所仓单日报
type WarehouseReceiptCZCE struct {
	Variety      string  `json:"variety"`       // 品种
	Warehouse    string  `json:"warehouse"`     // 仓库
	ReceiptQty   float64 `json:"receipt_qty"`   // 仓单数量
	ValidReceipt float64 `json:"valid_receipt"` // 有效预报
	TodayChange  float64 `json:"today_change"`  // 当日增减
}

// WarehouseReceiptDCE 大商所仓单日报
type WarehouseReceiptDCE struct {
	VarietyCode   string `json:"variety_code"`   // 品种代码
	VarietyName   string `json:"variety_name"`   // 品种名称
	Warehouse     string `json:"warehouse"`      // 仓库/分库
	DeliveryPlace string `json:"delivery_place"` // 可选提货地点/分库-数量
	LastQty       int64  `json:"last_qty"`       // 昨日仓单量（手）
	TodayQty      int64  `json:"today_qty"`      // 今日仓单量（手）
	Change        int64  `json:"change"`         // 增减（手）
}

// WarehouseReceiptSHFE 上期所仓单日报
type WarehouseReceiptSHFE struct {
	Variety       string `json:"variety"`        // 品种
	Region        string `json:"region"`         // 地区
	Warehouse     string `json:"warehouse"`      // 仓库
	ReceiptQty    int64  `json:"receipt_qty"`    // 仓单数量
	ReceiptChange int64  `json:"receipt_change"` // 仓单增减
}

// WarehouseReceiptGFEX 广期所仓单日报
type WarehouseReceiptGFEX struct {
	Symbol    string  `json:"symbol"`    // 品种代码
	Variety   string  `json:"variety"`   // 品种名称
	Warehouse string  `json:"warehouse"` // 仓库/分库
	LastQty   float64 `json:"last_qty"`  // 昨日仓单量
	TodayQty  float64 `json:"today_qty"` // 今日仓单量
	Change    float64 `json:"change"`    // 增减
}

// FuturesWarehouseReceiptCZCE 郑州商品交易所-交易数据-仓单日报
//
// 数据源: http://www.czce.com.cn/cn/jysj/cdrb/H770310index_1.htm
//
// 参数:
//   - date: 交易日，格式 "20200702"
//
// 返回:
//   - map[string][]WarehouseReceiptCZCE: 按品种分组的仓单日报数据
//   - error: 错误信息
func FuturesWarehouseReceiptCZCE(date string) (map[string][]WarehouseReceiptCZCE, error) {
	var url string
	dateInt := 0
	fmt.Sscanf(date, "%d", &dateInt)

	if dateInt > 20251101 {
		url = fmt.Sprintf("http://www.czce.com.cn/cn/DFSStaticFiles/Future/%s/%s/FutureDataWhsheet.xlsx", date[:4], date)
	} else {
		url = fmt.Sprintf("http://www.czce.com.cn/cn/DFSStaticFiles/Future/%s/%s/FutureDataWhsheet.xls", date[:4], date)
	}

	headers := map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
	}

	resp, err := utils.GetWithHeaders(url, nil, headers)
	if err != nil {
		return nil, fmt.Errorf("请求郑商所仓单日报失败: %w", err)
	}

	// 解析Excel文件
	f, err := excelize.OpenReader(strings.NewReader(string(resp.Body())))
	if err != nil {
		return nil, fmt.Errorf("解析Excel文件失败: %w", err)
	}
	defer f.Close()

	sheets := f.GetSheetList()
	if len(sheets) == 0 {
		return nil, fmt.Errorf("Excel文件没有工作表")
	}

	rows, err := f.GetRows(sheets[0])
	if err != nil {
		return nil, fmt.Errorf("读取工作表失败: %w", err)
	}

	// 查找品种分隔行
	result := make(map[string][]WarehouseReceiptCZCE)
	re := regexp.MustCompile(`[a-zA-Z]+`)

	var currentVariety string
	var headerRow []string
	inDataSection := false

	for _, row := range rows {
		if len(row) == 0 {
			continue
		}

		// 检查是否是品种行
		if len(row) > 0 && strings.Contains(row[0], "品种") {
			matches := re.FindStringSubmatch(row[0])
			if len(matches) > 0 {
				currentVariety = matches[0]
				inDataSection = false
				continue
			}
		}

		// 跳过空行或标题行
		if currentVariety == "" {
			continue
		}

		// 检查是否是表头行
		if !inDataSection && len(row) > 0 {
			for _, cell := range row {
				if strings.Contains(cell, "仓库") || strings.Contains(cell, "仓单") {
					headerRow = row
					inDataSection = true
					break
				}
			}
			continue
		}

		// 读取数据行
		if inDataSection && len(row) >= 3 {
			item := WarehouseReceiptCZCE{
				Variety: currentVariety,
			}

			for i, cell := range row {
				if i >= len(headerRow) {
					break
				}
				header := headerRow[i]
				switch {
				case strings.Contains(header, "仓库"):
					item.Warehouse = cell
				case strings.Contains(header, "数量") && !strings.Contains(header, "增减"):
					item.ReceiptQty = utils.MustParseFloat(cell)
				case strings.Contains(header, "预报"):
					item.ValidReceipt = utils.MustParseFloat(cell)
				case strings.Contains(header, "增减"):
					item.TodayChange = utils.MustParseFloat(cell)
				}
			}

			if item.Warehouse != "" {
				result[currentVariety] = append(result[currentVariety], item)
			}
		}
	}

	return result, nil
}

// dceWarehouseReceiptResponse DCE仓单日报API响应
type dceWarehouseReceiptResponse struct {
	Data struct {
		EntityList []struct {
			Variety      string `json:"variety"`
			WhAbbr       string `json:"whAbbr"`
			DeliveryAbbr string `json:"deliveryAbbr"`
			LastWbillQty int64  `json:"lastWbillQty"`
			WbillQty     int64  `json:"wbillQty"`
			Diff         int64  `json:"diff"`
			VarietyOrder string `json:"varietyOrder"`
		} `json:"entityList"`
	} `json:"data"`
}

// FuturesWarehouseReceiptDCE 大连商品交易所-行情数据-统计数据-日统计-仓单日报
//
// 数据源: http://www.dce.com.cn/dce/channel/list/187.html
//
// 参数:
//   - date: 交易日，格式 "20200702"
//
// 返回:
//   - []WarehouseReceiptDCE: 仓单日报数据
//   - error: 错误信息
func FuturesWarehouseReceiptDCE(date string) ([]WarehouseReceiptDCE, error) {
	url := "http://www.dce.com.cn/dcereport/publicweb/dailystat/wbillWeeklyQuotes"

	payload := map[string]any{
		"tradeDate": date,
		"varietyId": "all",
	}

	headers := map[string]string{
		"User-Agent":   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
		"Content-Type": "application/json",
	}

	resp, err := utils.PostJSONWithHeaders(url, nil, payload, headers)
	if err != nil {
		return nil, fmt.Errorf("请求大商所仓单日报失败: %w", err)
	}

	var apiResp dceWarehouseReceiptResponse
	if err := json.Unmarshal(resp.Body(), &apiResp); err != nil {
		return nil, fmt.Errorf("解析响应失败: %w", err)
	}

	var result []WarehouseReceiptDCE
	for _, item := range apiResp.Data.EntityList {
		result = append(result, WarehouseReceiptDCE{
			VarietyCode:   item.VarietyOrder,
			VarietyName:   item.Variety,
			Warehouse:     item.WhAbbr,
			DeliveryPlace: item.DeliveryAbbr,
			LastQty:       item.LastWbillQty,
			TodayQty:      item.WbillQty,
			Change:        item.Diff,
		})
	}

	return result, nil
}

// shfeWarehouseReceiptResponse SHFE仓单日报API响应
type shfeWarehouseReceiptResponse struct {
	OCursor []struct {
		VarName     string `json:"VARNAME"`
		RegName     string `json:"REGNAME"`
		WhAbbrName  string `json:"WHABBRNAME"`
		Wrtwghts    int64  `json:"WRTWGHTS"`
		WrtwghtsChg int64  `json:"WRTWGHTSCHG"`
	} `json:"o_cursor"`
}

// FuturesSHFEWarehouseReceipt 上海期货交易所指定交割仓库期货仓单日报
//
// 数据源: https://tsite.shfe.com.cn/statements/dataview.html?paramid=dailystock
//
// 参数:
//   - date: 交易日，格式 "20200702"
//
// 返回:
//   - map[string][]WarehouseReceiptSHFE: 按品种分组的仓单日报数据
//   - error: 错误信息
func FuturesSHFEWarehouseReceipt(date string) (map[string][]WarehouseReceiptSHFE, error) {
	headers := map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
	}

	url := fmt.Sprintf("https://www.shfe.com.cn/data/tradedata/future/dailydata/%sdailystock.dat", date)

	if date >= "20140519" {
		resp, err := utils.GetWithHeaders(url, nil, headers)
		if err != nil {
			return nil, fmt.Errorf("请求上期所仓单日报失败: %w", err)
		}

		var apiResp shfeWarehouseReceiptResponse
		if err := json.Unmarshal(resp.Body(), &apiResp); err != nil {
			return nil, fmt.Errorf("解析响应失败: %w", err)
		}

		result := make(map[string][]WarehouseReceiptSHFE)
		for _, item := range apiResp.OCursor {
			// 清理名称中的$符号
			varName := strings.Split(item.VarName, "$")[0]
			regName := strings.Split(item.RegName, "$")[0]
			whAbbrName := strings.Split(item.WhAbbrName, "$")[0]

			receipt := WarehouseReceiptSHFE{
				Variety:       varName,
				Region:        regName,
				Warehouse:     whAbbrName,
				ReceiptQty:    item.Wrtwghts,
				ReceiptChange: item.WrtwghtsChg,
			}

			result[varName] = append(result[varName], receipt)
		}

		return result, nil
	}

	// 早期日期使用HTML格式
	htmlURL := fmt.Sprintf("https://www.shfe.com.cn/data/tradedata/future/dailydata/%sdailystock.html", date)
	resp, err := utils.GetWithHeaders(htmlURL, nil, headers)
	if err != nil {
		return nil, fmt.Errorf("请求上期所仓单日报HTML失败: %w", err)
	}

	// 简化处理：返回空结果
	// 实际需要解析HTML表格
	_ = resp
	return make(map[string][]WarehouseReceiptSHFE), nil
}

// gfexWarehouseReceiptResponse GFEX仓单日报API响应
type gfexWarehouseReceiptResponse struct {
	Data []struct {
		VarietyOrder string `json:"varietyOrder"`
		Variety      string `json:"variety"`
		WhAbbr       string `json:"whAbbr"`
		WhType       string `json:"whType"`
		LastWbillQty any    `json:"lastWbillQty"`
		WbillQty     any    `json:"wbillQty"`
		RegWbillQty  any    `json:"regWbillQty"`
	} `json:"data"`
}

// FuturesGFEXWarehouseReceipt 广州期货交易所-行情数据-仓单日报
//
// 数据源: http://www.gfex.com.cn/gfex/cdrb/hqsj_tjsj.shtml
//
// 参数:
//   - date: 交易日，格式 "20240122"
//
// 返回:
//   - map[string][]WarehouseReceiptGFEX: 按品种分组的仓单日报数据
//   - error: 错误信息
func FuturesGFEXWarehouseReceipt(date string) (map[string][]WarehouseReceiptGFEX, error) {
	url := "http://www.gfex.com.cn/u/interfacesWebTdWbillWeeklyQuotes/loadList"

	headers := map[string]string{
		"User-Agent":   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
		"Content-Type": "application/x-www-form-urlencoded",
	}

	payload := map[string]string{
		"gen_date": date,
	}

	resp, err := utils.PostFormWithHeaders(url, payload, headers)
	if err != nil {
		return nil, fmt.Errorf("请求广期所仓单日报失败: %w", err)
	}

	var apiResp gfexWarehouseReceiptResponse
	if err := json.Unmarshal(resp.Body(), &apiResp); err != nil {
		return nil, fmt.Errorf("解析响应失败: %w", err)
	}

	result := make(map[string][]WarehouseReceiptGFEX)
	for _, item := range apiResp.Data {
		// 跳过无效数据
		if item.WhType == "" {
			continue
		}

		symbol := strings.ToUpper(item.VarietyOrder)
		if symbol == "" {
			continue
		}

		lastQty, _ := utils.ToFloat64(item.LastWbillQty)
		todayQty, _ := utils.ToFloat64(item.WbillQty)
		change, _ := utils.ToFloat64(item.RegWbillQty)

		receipt := WarehouseReceiptGFEX{
			Symbol:    symbol,
			Variety:   item.Variety,
			Warehouse: item.WhAbbr,
			LastQty:   lastQty,
			TodayQty:  todayQty,
			Change:    change,
		}

		result[symbol] = append(result[symbol], receipt)
	}

	return result, nil
}
