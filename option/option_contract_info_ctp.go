package option

import (
	"encoding/json"
	"fmt"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

// CTPContractInfo CTP期权合约信息
type CTPContractInfo struct {
	ExchangeID               string  `json:"ExchangeID"`               // 交易所ID
	InstrumentID             string  `json:"InstrumentID"`             // 合约ID
	InstrumentName           string  `json:"InstrumentName"`           // 合约名称
	ProductClass             string  `json:"ProductClass"`             // 商品类别
	ProductID                string  `json:"ProductID"`                // 品种ID
	VolumeMultiple           int     `json:"VolumeMultiple"`           // 合约乘数
	PriceTick                float64 `json:"PriceTick"`                // 最小变动价位
	LongMarginRatioByMoney   float64 `json:"LongMarginRatioByMoney"`   // 做多保证金率
	ShortMarginRatioByMoney  float64 `json:"ShortMarginRatioByMoney"`  // 做空保证金率
	LongMarginRatioByVolume  float64 `json:"LongMarginRatioByVolume"`  // 做多保证金/手
	ShortMarginRatioByVolume float64 `json:"ShortMarginRatioByVolume"` // 做空保证金/手
	OpenRatioByMoney         float64 `json:"OpenRatioByMoney"`         // 开仓手续费率
	OpenRatioByVolume        float64 `json:"OpenRatioByVolume"`        // 开仓手续费/手
	CloseRatioByMoney        float64 `json:"CloseRatioByMoney"`        // 平仓手续费率
	CloseRatioByVolume       float64 `json:"CloseRatioByVolume"`       // 平仓手续费/手
	CloseTodayRatioByMoney   float64 `json:"CloseTodayRatioByMoney"`   // 平今手续费率
	CloseTodayRatioByVolume  float64 `json:"CloseTodayRatioByVolume"`  // 平今手续费/手
	DeliveryYear             int     `json:"DeliveryYear"`             // 交割年份
	DeliveryMonth            int     `json:"DeliveryMonth"`            // 交割月份
	OpenDate                 string  `json:"OpenDate"`                 // 上市日期
	ExpireDate               string  `json:"ExpireDate"`               // 最后交易日
	DeliveryDate             string  `json:"DeliveryDate"`             // 交割日
	UnderlyingInstrID        string  `json:"UnderlyingInstrID"`        // 标的合约ID
	UnderlyingMultiple       int     `json:"UnderlyingMultiple"`       // 标的合约乘数
	OptionsType              string  `json:"OptionsType"`              // 期权类型
	StrikePrice              float64 `json:"StrikePrice"`              // 行权价
	InstLifePhase            string  `json:"InstLifePhase"`            // 合约状态
}

// CTPResponse CTP API响应结构
type CTPResponse struct {
	Code int               `json:"code"`
	Msg  string            `json:"msg"`
	Data []CTPContractInfo `json:"data"`
}

// OptionContractInfoCtp openctp-合约信息接口-期权合约
//
// 返回:
//   - dataframe.DataFrame: 期权合约信息
//   - error: 错误信息
//
// 数据源: http://openctp.cn/instruments.html
func OptionContractInfoCtp() (dataframe.DataFrame, error) {
	url := "http://dict.openctp.cn/instruments?types=option"

	// 请求数据
	resp, err := utils.Get(url, nil)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	// 解析JSON
	var ctpResp CTPResponse
	if err := json.Unmarshal([]byte(resp.String()), &ctpResp); err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("解析JSON失败: %w", err)
	}

	if ctpResp.Code != 0 {
		return dataframe.DataFrame{}, fmt.Errorf("API返回错误: %s", ctpResp.Msg)
	}

	if len(ctpResp.Data) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到合约信息")
	}

	// 转换为DataFrame
	df := convertCTPToDataFrame(ctpResp.Data)
	return df, nil
}

// convertCTPToDataFrame 将CTP合约信息转换为DataFrame
func convertCTPToDataFrame(contracts []CTPContractInfo) dataframe.DataFrame {
	// 字段映射：英文字段名 -> 中文字段名
	columnMapping := map[string]string{
		"ExchangeID":               "交易所ID",
		"InstrumentID":             "合约ID",
		"InstrumentName":           "合约名称",
		"ProductClass":             "商品类别",
		"ProductID":                "品种ID",
		"VolumeMultiple":           "合约乘数",
		"PriceTick":                "最小变动价位",
		"LongMarginRatioByMoney":   "做多保证金率",
		"ShortMarginRatioByMoney":  "做空保证金率",
		"LongMarginRatioByVolume":  "做多保证金/手",
		"ShortMarginRatioByVolume": "做空保证金/手",
		"OpenRatioByMoney":         "开仓手续费率",
		"OpenRatioByVolume":        "开仓手续费/手",
		"CloseRatioByMoney":        "平仓手续费率",
		"CloseRatioByVolume":       "平仓手续费/手",
		"CloseTodayRatioByMoney":   "平今手续费率",
		"CloseTodayRatioByVolume":  "平今手续费/手",
		"DeliveryYear":             "交割年份",
		"DeliveryMonth":            "交割月份",
		"OpenDate":                 "上市日期",
		"ExpireDate":               "最后交易日",
		"DeliveryDate":             "交割日",
		"UnderlyingInstrID":        "标的合约ID",
		"UnderlyingMultiple":       "标的合约乘数",
		"OptionsType":              "期权类型",
		"StrikePrice":              "行权价",
		"InstLifePhase":            "合约状态",
	}

	// 按中文列名顺序构建数据
	chineseColumns := []string{
		"交易所ID", "合约ID", "合约名称", "商品类别", "品种ID", "合约乘数", "最小变动价位",
		"做多保证金率", "做空保证金率", "做多保证金/手", "做空保证金/手", "开仓手续费率",
		"开仓手续费/手", "平仓手续费率", "平仓手续费/手", "平今手续费率", "平今手续费/手",
		"交割年份", "交割月份", "上市日期", "最后交易日", "交割日", "标的合约ID",
		"标的合约乘数", "期权类型", "行权价", "合约状态",
	}

	// 创建反向映射
	reverseMapping := make(map[string]string)
	for eng, chi := range columnMapping {
		reverseMapping[chi] = eng
	}

	// 构建series列表
	var seriesList []series.Series

	// 为每个中文列创建series
	for _, chiCol := range chineseColumns {
		var values []interface{}
		engCol := reverseMapping[chiCol]

		for _, contract := range contracts {
			value := getCTPFieldValue(contract, engCol)
			values = append(values, value)
		}

		seriesList = append(seriesList, series.New(values, series.String, chiCol))
	}

	// 创建DataFrame
	df := dataframe.New(seriesList...)
	return df
}

// getCTPFieldValue 获取CTP合约字段值
func getCTPFieldValue(contract CTPContractInfo, field string) interface{} {
	switch field {
	case "ExchangeID":
		return contract.ExchangeID
	case "InstrumentID":
		return contract.InstrumentID
	case "InstrumentName":
		return contract.InstrumentName
	case "ProductClass":
		return contract.ProductClass
	case "ProductID":
		return contract.ProductID
	case "VolumeMultiple":
		return contract.VolumeMultiple
	case "PriceTick":
		return contract.PriceTick
	case "LongMarginRatioByMoney":
		return contract.LongMarginRatioByMoney
	case "ShortMarginRatioByMoney":
		return contract.ShortMarginRatioByMoney
	case "LongMarginRatioByVolume":
		return contract.LongMarginRatioByVolume
	case "ShortMarginRatioByVolume":
		return contract.ShortMarginRatioByVolume
	case "OpenRatioByMoney":
		return contract.OpenRatioByMoney
	case "OpenRatioByVolume":
		return contract.OpenRatioByVolume
	case "CloseRatioByMoney":
		return contract.CloseRatioByMoney
	case "CloseRatioByVolume":
		return contract.CloseRatioByVolume
	case "CloseTodayRatioByMoney":
		return contract.CloseTodayRatioByMoney
	case "CloseTodayRatioByVolume":
		return contract.CloseTodayRatioByVolume
	case "DeliveryYear":
		return contract.DeliveryYear
	case "DeliveryMonth":
		return contract.DeliveryMonth
	case "OpenDate":
		return contract.OpenDate
	case "ExpireDate":
		return contract.ExpireDate
	case "DeliveryDate":
		return contract.DeliveryDate
	case "UnderlyingInstrID":
		return contract.UnderlyingInstrID
	case "UnderlyingMultiple":
		return contract.UnderlyingMultiple
	case "OptionsType":
		return contract.OptionsType
	case "StrikePrice":
		return contract.StrikePrice
	case "InstLifePhase":
		return contract.InstLifePhase
	default:
		return ""
	}
}
