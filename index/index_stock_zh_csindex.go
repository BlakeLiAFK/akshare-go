package index

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/BlakeLiAFK/akshare/utils"
)

// CSIndexHist 中证指数历史行情
type CSIndexHist struct {
	Date        time.Time `json:"date"`          // 日期
	IndexCode   string    `json:"index_code"`    // 指数代码
	IndexNameCN string    `json:"index_name_cn"` // 指数中文全称
	IndexName   string    `json:"index_name"`    // 指数中文简称
	IndexNameEN string    `json:"index_name_en"` // 指数英文全称
	IndexShort  string    `json:"index_short"`   // 指数英文简称
	Open        float64   `json:"open"`          // 开盘
	High        float64   `json:"high"`          // 最高
	Low         float64   `json:"low"`           // 最低
	Close       float64   `json:"close"`         // 收盘
	Change      float64   `json:"change"`        // 涨跌
	ChangePct   float64   `json:"change_pct"`    // 涨跌幅
	Volume      int64     `json:"volume"`        // 成交量
	Amount      float64   `json:"amount"`        // 成交金额
	SampleNum   int       `json:"sample_num"`    // 样本数量
	PETTM       float64   `json:"pe_ttm"`        // 滚动市盈率
}

// CSIndexValue 中证指数估值数据
type CSIndexValue struct {
	Date        time.Time `json:"date"`          // 日期
	IndexCode   string    `json:"index_code"`    // 指数代码
	IndexNameCN string    `json:"index_name_cn"` // 指数中文全称
	IndexName   string    `json:"index_name"`    // 指数中文简称
	IndexNameEN string    `json:"index_name_en"` // 指数英文全称
	IndexShort  string    `json:"index_short"`   // 指数英文简称
	PE1         float64   `json:"pe1"`           // 市盈率1
	PE2         float64   `json:"pe2"`           // 市盈率2
	DY1         float64   `json:"dy1"`           // 股息率1
	DY2         float64   `json:"dy2"`           // 股息率2
}

// csIndexHistResponse 中证指数历史行情API响应
type csIndexHistResponse struct {
	Data [][]interface{} `json:"data"`
}

// StockZhIndexHistCSIndex 中证指数历史行情
//
// 数据源: https://www.csindex.com.cn/zh-CN/indices/index-detail/H30374
//
// 参数:
//   - symbol: 指数代码，如 "000928"
//   - startDate: 开始日期，格式 "20180526"
//   - endDate: 结束日期，格式 "20240604"
//
// 返回:
//   - []CSIndexHist: 指数历史行情数据
//   - error: 错误信息
func StockZhIndexHistCSIndex(symbol, startDate, endDate string) ([]CSIndexHist, error) {
	apiURL := "https://www.csindex.com.cn/csindex-home/perf/index-perf"

	params := map[string]string{
		"indexCode": symbol,
		"startDate": startDate,
		"endDate":   endDate,
	}

	resp, err := utils.Get(apiURL, params)
	if err != nil {
		return nil, fmt.Errorf("请求中证指数历史行情失败: %w", err)
	}

	var apiResp csIndexHistResponse
	if err := json.Unmarshal(resp.Body(), &apiResp); err != nil {
		return nil, fmt.Errorf("解析中证指数历史行情响应失败: %w", err)
	}

	result := make([]CSIndexHist, 0, len(apiResp.Data))
	for _, row := range apiResp.Data {
		if len(row) < 16 {
			continue
		}

		dateStr := toString(row[0])
		date, err := time.Parse("2006-01-02", dateStr[:10])
		if err != nil {
			continue
		}

		result = append(result, CSIndexHist{
			Date:        date,
			IndexCode:   toString(row[1]),
			IndexNameCN: toString(row[2]),
			IndexName:   toString(row[3]),
			IndexNameEN: toString(row[4]),
			IndexShort:  toString(row[5]),
			Open:        toFloat64(row[6]),
			High:        toFloat64(row[7]),
			Low:         toFloat64(row[8]),
			Close:       toFloat64(row[9]),
			Change:      toFloat64(row[10]),
			ChangePct:   toFloat64(row[11]),
			Volume:      int64(toFloat64(row[12])),
			Amount:      toFloat64(row[13]),
			SampleNum:   int(toFloat64(row[14])),
			PETTM:       toFloat64(row[15]),
		})
	}

	return result, nil
}

// StockZhIndexValueCSIndex 中证指数估值数据
//
// 数据源: https://www.csindex.com.cn/zh-CN/indices/index-detail/H30374
//
// 参数:
//   - symbol: 指数代码，如 "H30374"
//
// 返回:
//   - []CSIndexValue: 指数估值数据
//   - error: 错误信息
//
// 注意: 该接口需要下载Excel文件，目前暂不支持
func StockZhIndexValueCSIndex(symbol string) ([]CSIndexValue, error) {
	// 此接口需要下载Excel文件解析，暂时返回空
	// 原Python实现使用 pd.read_excel(url) 读取Excel文件
	// Go实现需要使用excelize等库来解析
	return nil, fmt.Errorf("该接口暂不支持，需要下载Excel文件: https://oss-ch.csindex.com.cn/static/html/csindex/public/uploads/file/autofile/indicator/%sindicator.xls", symbol)
}
