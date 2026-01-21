package index

import (
	"fmt"
)

// ExampleIndexZHAHist 示例：获取上证指数历史K线数据
func ExampleIndexZHAHist() {
	klines, err := IndexZHAHist("000001", "daily", "20240101", "20240131")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	for _, k := range klines {
		fmt.Printf("%s: 开盘=%.2f, 收盘=%.2f\n", k.Date.Format("2006-01-02"), k.Open, k.Close)
	}
}

// ExampleStockZHIndexSpotEM 示例：获取沪深重要指数实时行情
func ExampleStockZHIndexSpotEM() {
	quotes, err := StockZHIndexSpotEM("沪深重要指数")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	for _, q := range quotes {
		fmt.Printf("%s(%s): %.2f, %.2f%%\n", q.Name, q.Code, q.Price, q.ChangePct)
	}
}
