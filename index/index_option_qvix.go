package index

import (
	"encoding/csv"
	"fmt"
	"strings"
	"time"

	"github.com/BlakeLiAFK/akshare/utils"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// OptionQvixDaily 期权波动率指数日线数据
type OptionQvixDaily struct {
	Date  time.Time `json:"date"`  // 日期
	Open  float64   `json:"open"`  // 开盘
	High  float64   `json:"high"`  // 最高
	Low   float64   `json:"low"`   // 最低
	Close float64   `json:"close"` // 收盘
}

// OptionQvixMin 期权波动率指数分时数据
type OptionQvixMin struct {
	Time string  `json:"time"` // 时间
	Qvix float64 `json:"qvix"` // QVIX
}

// 期权波动率指数列索引映射
var qvixColumnMap = map[string][]int{
	"50ETF":     {0, 1, 2, 3, 4},
	"300ETF":    {0, 9, 10, 11, 12},
	"500ETF":    {0, 67, 68, 69, 70},
	"CYB":       {0, 71, 72, 73, 74},
	"KCB":       {0, 83, 84, 85, 86},
	"100ETF":    {0, 75, 76, 77, 78},
	"300Index":  {0, 17, 18, 19, 20},
	"1000Index": {0, 25, 26, 27, 28},
	"50Index":   {0, 79, 80, 81, 82},
}

// 期权波动率指数分时URL映射
var qvixMinURLMap = map[string]string{
	"50ETF":     "http://1.optbbs.com/d/csv/d/vix50.csv",
	"300ETF":    "http://1.optbbs.com/d/csv/d/vix300.csv",
	"500ETF":    "http://1.optbbs.com/d/csv/d/vix500.csv",
	"CYB":       "http://1.optbbs.com/d/csv/d/vixcyb.csv",
	"KCB":       "http://1.optbbs.com/d/csv/d/vixkcb.csv",
	"100ETF":    "http://1.optbbs.com/d/csv/d/vix100.csv",
	"300Index":  "http://1.optbbs.com/d/csv/d/vixindex.csv",
	"1000Index": "http://1.optbbs.com/d/csv/d/vixindex1000.csv",
	"50Index":   "http://1.optbbs.com/d/csv/d/vix50index.csv",
}

// fetchQvixDaily 获取期权波动率指数日线数据
func fetchQvixDaily(symbol string) ([]OptionQvixDaily, error) {
	columns, ok := qvixColumnMap[symbol]
	if !ok {
		return nil, fmt.Errorf("无效的 symbol: %s", symbol)
	}

	apiURL := "http://1.optbbs.com/d/csv/d/k.csv"

	resp, err := utils.Get(apiURL, nil)
	if err != nil {
		return nil, fmt.Errorf("请求期权波动率指数失败: %w", err)
	}

	// 转换 GBK 到 UTF-8
	reader := transform.NewReader(strings.NewReader(resp.String()), simplifiedchinese.GBK.NewDecoder())
	csvReader := csv.NewReader(reader)
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("解析CSV失败: %w", err)
	}

	if len(records) < 2 {
		return []OptionQvixDaily{}, nil
	}

	result := make([]OptionQvixDaily, 0, len(records)-1)
	for i := 1; i < len(records); i++ {
		row := records[i]
		if len(row) <= columns[4] {
			continue
		}

		date, err := time.Parse("2006/1/2", row[columns[0]])
		if err != nil {
			date, _ = time.Parse("2006-01-02", row[columns[0]])
		}

		result = append(result, OptionQvixDaily{
			Date:  date,
			Open:  utils.MustParseFloat(row[columns[1]]),
			High:  utils.MustParseFloat(row[columns[2]]),
			Low:   utils.MustParseFloat(row[columns[3]]),
			Close: utils.MustParseFloat(row[columns[4]]),
		})
	}

	return result, nil
}

// fetchQvixMin 获取期权波动率指数分时数据
func fetchQvixMin(symbol string) ([]OptionQvixMin, error) {
	apiURL, ok := qvixMinURLMap[symbol]
	if !ok {
		return nil, fmt.Errorf("无效的 symbol: %s", symbol)
	}

	resp, err := utils.Get(apiURL, nil)
	if err != nil {
		return nil, fmt.Errorf("请求期权波动率指数分时数据失败: %w", err)
	}

	csvReader := csv.NewReader(strings.NewReader(resp.String()))
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("解析CSV失败: %w", err)
	}

	if len(records) < 2 {
		return []OptionQvixMin{}, nil
	}

	result := make([]OptionQvixMin, 0, len(records)-1)
	for i := 1; i < len(records); i++ {
		row := records[i]
		if len(row) < 2 {
			continue
		}

		result = append(result, OptionQvixMin{
			Time: row[0],
			Qvix: utils.MustParseFloat(row[1]),
		})
	}

	return result, nil
}

// IndexOption50etfQvix 50ETF 期权波动率指数 QVIX
func IndexOption50etfQvix() ([]OptionQvixDaily, error) {
	return fetchQvixDaily("50ETF")
}

// IndexOption50etfMinQvix 50ETF 期权波动率指数 QVIX 分时
func IndexOption50etfMinQvix() ([]OptionQvixMin, error) {
	return fetchQvixMin("50ETF")
}

// IndexOption300etfQvix 300ETF 期权波动率指数 QVIX
func IndexOption300etfQvix() ([]OptionQvixDaily, error) {
	return fetchQvixDaily("300ETF")
}

// IndexOption300etfMinQvix 300ETF 期权波动率指数 QVIX 分时
func IndexOption300etfMinQvix() ([]OptionQvixMin, error) {
	return fetchQvixMin("300ETF")
}

// IndexOption500etfQvix 500ETF 期权波动率指数 QVIX
func IndexOption500etfQvix() ([]OptionQvixDaily, error) {
	return fetchQvixDaily("500ETF")
}

// IndexOption500etfMinQvix 500ETF 期权波动率指数 QVIX 分时
func IndexOption500etfMinQvix() ([]OptionQvixMin, error) {
	return fetchQvixMin("500ETF")
}

// IndexOptionCybQvix 创业板 期权波动率指数 QVIX
func IndexOptionCybQvix() ([]OptionQvixDaily, error) {
	return fetchQvixDaily("CYB")
}

// IndexOptionCybMinQvix 创业板 期权波动率指数 QVIX 分时
func IndexOptionCybMinQvix() ([]OptionQvixMin, error) {
	return fetchQvixMin("CYB")
}

// IndexOptionKcbQvix 科创板 期权波动率指数 QVIX
func IndexOptionKcbQvix() ([]OptionQvixDaily, error) {
	return fetchQvixDaily("KCB")
}

// IndexOptionKcbMinQvix 科创板 期权波动率指数 QVIX 分时
func IndexOptionKcbMinQvix() ([]OptionQvixMin, error) {
	return fetchQvixMin("KCB")
}

// IndexOption100etfQvix 深证100ETF 期权波动率指数 QVIX
func IndexOption100etfQvix() ([]OptionQvixDaily, error) {
	return fetchQvixDaily("100ETF")
}

// IndexOption100etfMinQvix 深证100ETF 期权波动率指数 QVIX 分时
func IndexOption100etfMinQvix() ([]OptionQvixMin, error) {
	return fetchQvixMin("100ETF")
}

// IndexOption300indexQvix 中证300股指 期权波动率指数 QVIX
func IndexOption300indexQvix() ([]OptionQvixDaily, error) {
	return fetchQvixDaily("300Index")
}

// IndexOption300indexMinQvix 中证300股指 期权波动率指数 QVIX 分时
func IndexOption300indexMinQvix() ([]OptionQvixMin, error) {
	return fetchQvixMin("300Index")
}

// IndexOption1000indexQvix 中证1000股指 期权波动率指数 QVIX
func IndexOption1000indexQvix() ([]OptionQvixDaily, error) {
	return fetchQvixDaily("1000Index")
}

// IndexOption1000indexMinQvix 中证1000股指 期权波动率指数 QVIX 分时
func IndexOption1000indexMinQvix() ([]OptionQvixMin, error) {
	return fetchQvixMin("1000Index")
}

// IndexOption50indexQvix 上证50股指 期权波动率指数 QVIX
func IndexOption50indexQvix() ([]OptionQvixDaily, error) {
	return fetchQvixDaily("50Index")
}

// IndexOption50indexMinQvix 上证50股指 期权波动率指数 QVIX 分时
func IndexOption50indexMinQvix() ([]OptionQvixMin, error) {
	return fetchQvixMin("50Index")
}
