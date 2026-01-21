package option

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

// 大连商品期权代码映射
var dceOptionCodeMap = map[string]string{
	"玉米期权":    "c",
	"豆粕期权":    "m",
	"铁矿石期权":   "i",
	"液化石油气期权": "pg",
	"聚乙烯期权":   "l",
	"聚氯乙烯期权":  "v",
	"聚丙烯期权":   "pp",
	"棕榈油期权":   "p",
	"黄大豆1号期权": "a",
	"黄大豆2号期权": "b",
	"豆油期权":    "y",
	"乙二醇期权":   "eg",
	"苯乙烯期权":   "eb",
	"鸡蛋期权":    "jd",
	"玉米淀粉期权":  "cs",
	"生猪期权":    "lh",
	"原木期权":    "lg",
}

// DCEOptionData 大连商品期权数据结构
type DCEOptionData struct {
	Variety      string  `json:"variety"`      // 品种名称
	ContractID   string  `json:"contractId"`   // 合约
	Open         float64 `json:"open"`         // 开盘价
	High         float64 `json:"high"`         // 最高价
	Low          float64 `json:"low"`          // 最低价
	Close        float64 `json:"close"`        // 收盘价
	LastClear    float64 `json:"lastClear"`    // 前结算价
	ClearPrice   float64 `json:"clearPrice"`   // 结算价
	Diff         float64 `json:"diff"`         // 涨跌
	Diff1        float64 `json:"diff1"`        // 涨跌1
	Delta        float64 `json:"delta"`        // Delta
	Volume       int     `json:"volume"`       // 成交量
	OpenInterest int     `json:"openInterest"` // 持仓量
	DiffI        int     `json:"diffI"`        // 持仓量变化
	Turnover     float64 `json:"turnover"`     // 成交额
}

// DCEOptionResponse 大连商品期权API响应
type DCEOptionResponse struct {
	Code int             `json:"code"`
	Msg  string          `json:"msg"`
	Data []DCEOptionData `json:"data"`
}

// OptionHistDce 大连商品交易所-期权-日频行情数据
//
// 参数:
//   - symbol: 期权品种名称，如 "聚丙烯期权"
//   - tradeDate: 交易日，格式 "20251016"
//
// 返回:
//   - dataframe.DataFrame: 日频行情数据
//   - error: 错误信息
//
// 数据源: http://www.dce.com.cn/
func OptionHistDce(symbol, tradeDate string) (dataframe.DataFrame, error) {
	if symbol == "" {
		return dataframe.DataFrame{}, fmt.Errorf("品种名称不能为空")
	}
	if tradeDate == "" {
		return dataframe.DataFrame{}, fmt.Errorf("交易日期不能为空")
	}

	// 检查品种是否有效
	code, exists := dceOptionCodeMap[symbol]
	if !exists {
		return dataframe.DataFrame{}, fmt.Errorf("不支持的品种: %s", symbol)
	}

	// 验证日期格式
	if len(tradeDate) != 8 {
		return dataframe.DataFrame{}, fmt.Errorf("日期格式错误，应为YYYYMMDD")
	}

	// 检查是否为交易日
	calendar, err := GetCalendar()
	if err == nil {
		day, err := ConvertDate(tradeDate)
		if err == nil {
			dayStr := day.Format("20060102")
			isTradingDay := false
			for _, calDay := range calendar {
				if calDay == dayStr {
					isTradingDay = true
					break
				}
			}
			if !isTradingDay {
				return dataframe.DataFrame{}, fmt.Errorf("%s非交易日", tradeDate)
			}
		}
	}

	// 构建请求
	url := "http://www.dce.com.cn/dcereport/publicweb/dailystat/dayQuotes"
	payload := map[string]interface{}{
		"contractId":     "",
		"lang":           "zh",
		"optionSeries":   "",
		"statisticsType": 0,
		"tradeDate":      tradeDate,
		"tradeType":      2,
		"varietyId":      code,
	}

	// 发送POST请求
	resp, err := utils.PostJSON(url, payload)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	// 解析JSON响应
	var dceResp DCEOptionResponse
	if err := json.Unmarshal([]byte(resp.String()), &dceResp); err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("解析JSON失败: %w", err)
	}

	if dceResp.Code != 0 {
		return dataframe.DataFrame{}, fmt.Errorf("API返回错误: %s", dceResp.Msg)
	}

	if len(dceResp.Data) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到数据")
	}

	// 转换为DataFrame
	df := convertDCEToDataFrame(dceResp.Data)
	return df, nil
}

// convertDCEToDataFrame 转换大连商品期权数据为DataFrame
func convertDCEToDataFrame(data []DCEOptionData) dataframe.DataFrame {
	// 定义中文列名
	columns := []string{
		"品种名称", "合约", "开盘价", "最高价", "最低价", "收盘价", "前结算价", "结算价",
		"涨跌", "涨跌1", "Delta", "成交量", "持仓量", "持仓量变化", "成交额",
	}

	// 创建数据记录
	var records [][]string

	// 添加表头
	records = append(records, columns)

	// 添加数据行
	for _, item := range data {
		record := []string{
			item.Variety,
			item.ContractID,
			formatFloat(item.Open),
			formatFloat(item.High),
			formatFloat(item.Low),
			formatFloat(item.Close),
			formatFloat(item.LastClear),
			formatFloat(item.ClearPrice),
			formatFloat(item.Diff),
			formatFloat(item.Diff1),
			formatFloat(item.Delta),
			strconv.Itoa(item.Volume),
			strconv.Itoa(item.OpenInterest),
			strconv.Itoa(item.DiffI),
			formatFloat(item.Turnover),
		}
		records = append(records, record)
	}

	// 构建DataFrame
	df := dataframe.LoadRecords(records)
	return df
}

// formatFloat 格式化浮点数
func formatFloat(f float64) string {
	return strconv.FormatFloat(f, 'f', 4, 64)
}

// GetDceOptionSymbols 获取大连商品期权品种列表
//
// 返回:
//   - dataframe.DataFrame: 包含品种代码和名称的数据
func GetDceOptionSymbols() (dataframe.DataFrame, error) {
	var symbols []string
	var codes []string

	for symbol, code := range dceOptionCodeMap {
		symbols = append(symbols, symbol)
		codes = append(codes, code)
	}

	df := dataframe.New(
		series.New(symbols, series.String, "symbol"),
		series.New(codes, series.String, "code"),
	)

	return df, nil
}

// OptionHistShfe 上海期货交易所-期权-日频行情数据
//
// 参数:
//   - symbol: 期权品种代码，如 "cu" (铜)
//   - tradeDate: 交易日，格式 "20251016"
//
// 返回:
//   - dataframe.DataFrame: 日频行情数据
//   - error: 错误信息
//
// 数据源: https://tsite.shfe.com.cn/
func OptionHistShfe(symbol, tradeDate string) (dataframe.DataFrame, error) {
	if symbol == "" {
		return dataframe.DataFrame{}, fmt.Errorf("品种代码不能为空")
	}
	if tradeDate == "" {
		return dataframe.DataFrame{}, fmt.Errorf("交易日期不能为空")
	}

	// 验证日期格式
	if len(tradeDate) != 8 {
		return dataframe.DataFrame{}, fmt.Errorf("日期格式错误，应为YYYYMMDD")
	}

	// 构建URL
	url := fmt.Sprintf(SHFEOptionURL, tradeDate)

	// 请求数据
	resp, err := utils.Get(url, SHFEHeaders)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	// 解析文本数据
	text := strings.TrimSpace(resp.String())
	if text == "" {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到数据")
	}

	// 按行分割
	lines := strings.Split(text, "\n")
	if len(lines) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("数据格式错误")
	}

	// 解析数据（假设为制表符分隔）
	var records [][]string
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		// 按制表符或空格分割
		fields := strings.Fields(line)
		if len(fields) > 0 {
			records = append(records, fields)
		}
	}

	if len(records) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("没有有效的数据记录")
	}

	// 转换数值列
	for i := 1; i < len(records); i++ { // 跳过表头
		for j := range records[i] {
			records[i][j] = parseNumericValue(records[i][j])
		}
	}

	// 构建DataFrame
	df := dataframe.LoadRecords(records)
	return df, nil
}
