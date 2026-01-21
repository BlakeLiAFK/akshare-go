package other

import (
	"fmt"
)

// ExampleCarMarketTotalCPCA 示例：获取乘联会-总体市场数据
func ExampleCarMarketTotalCPCA() {
	data, err := CarMarketTotalCPCA("狭义乘用车", "产量")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	for _, item := range data {
		fmt.Printf("%s: 当年=%.2f, 去年=%.2f\n", item.Month, item.CurrentYear, item.PreviousYear)
	}
}

// ExampleCarMarketManRankCPCA 示例：获取乘联会-厂商排名数据
func ExampleCarMarketManRankCPCA() {
	data, err := CarMarketManRankCPCA("狭义乘用车-单月", "批发")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	for _, item := range data {
		fmt.Printf("%s: 当年=%.2f, 去年=%.2f\n", item.Manufacturer, item.CurrentYear, item.PreviousYear)
	}
}

// ExampleCarSaleRankGasgoo 示例：获取盖世汽车-销量排名
func ExampleCarSaleRankGasgoo() {
	data, err := CarSaleRankGasgoo("车企榜", "202311")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	for i, item := range data {
		if i >= 5 {
			break
		}
		fmt.Printf("第%d名 %s: 销量=%.0f\n", item.Rank, item.Name, item.Sales)
	}
}
