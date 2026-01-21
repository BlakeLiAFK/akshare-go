package hf

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/BlakeLiAFK/akshare/utils"
)

// HfSp500Item S&P500高频数据结构
type HfSp500Item struct {
	Date  string  `json:"date"`  // 日期
	Open  float64 `json:"open"`  // 开盘价
	High  float64 `json:"high"`  // 最高价
	Low   float64 `json:"low"`   // 最低价
	Close float64 `json:"close"` // 收盘价
	Price float64 `json:"price"` // 价格
}

// HfSp500 获取S&P500高频分钟数据
//
// 数据源: https://github.com/FutureSharks/financial-data
//
// 参数:
//   - year: 年份，从2012到2018
//
// 返回:
//   - []HfSp500Item: S&P500分钟数据列表
//   - error: 错误信息
//
// 示例:
//
//	data, err := hf.HfSp500("2017")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, item := range data[:10] {
//	    fmt.Printf("%s %.2f %.2f %.2f %.2f\n", item.Date, item.Open, item.High, item.Low, item.Close)
//	}
func HfSp500(year string) ([]HfSp500Item, error) {
	url := fmt.Sprintf("https://github.com/FutureSharks/financial-data/raw/master/pyfinancialdata/data/stocks/histdata/SPXUSD/DAT_ASCII_SPXUSD_M1_%s.csv", year)

	headers := map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
	}

	resp, err := utils.GetWithHeaders(url, nil, headers)
	if err != nil {
		return nil, fmt.Errorf("获取S&P500数据失败: %w", err)
	}

	var items []HfSp500Item
	scanner := bufio.NewScanner(strings.NewReader(resp.String()))
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ";")
		if len(parts) >= 6 {
			// 日期格式处理
			dateStr := parts[0]
			if len(dateStr) >= 10 {
				dateStr = dateStr[:10]
			}

			item := HfSp500Item{
				Date:  dateStr,
				Open:  utils.MustFloat64(parts[1]),
				High:  utils.MustFloat64(parts[2]),
				Low:   utils.MustFloat64(parts[3]),
				Close: utils.MustFloat64(parts[4]),
				Price: utils.MustFloat64(parts[5]),
			}
			items = append(items, item)
		}
	}

	return items, nil
}
