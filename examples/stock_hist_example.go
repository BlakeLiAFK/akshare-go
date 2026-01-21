package main

import (
	"fmt"
	"log"
	"time"

	"github.com/BlakeLiAFK/akshare/stock"
)

func main() {
	// 示例1: 获取平安银行最近30天的日K线数据（不复权）
	fmt.Println("=== 示例1: 获取平安银行最近30天日K线（不复权） ===")
	endDate := time.Now().Format("20060102")
	startDate := time.Now().AddDate(0, 0, -30).Format("20060102")

	klines, err := stock.StockZhAHist("000001", "daily", startDate, endDate, "")
	if err != nil {
		log.Fatalf("获取K线数据失败: %v", err)
	}

	fmt.Printf("共获取 %d 条数据\n", len(klines))
	if len(klines) > 0 {
		fmt.Println("\n最近5个交易日数据:")
		count := 5
		if len(klines) < 5 {
			count = len(klines)
		}
		for i := len(klines) - count; i < len(klines); i++ {
			k := klines[i]
			fmt.Printf("%s: 开盘=%.2f, 收盘=%.2f, 最高=%.2f, 最低=%.2f, 成交量=%d, 涨跌幅=%.2f%%\n",
				k.Date.Format("2006-01-02"),
				k.Open, k.Close, k.High, k.Low,
				k.Volume, k.ChangePct)
		}
	}

	// 示例2: 获取贵州茅台2024年周K线数据（前复权）
	fmt.Println("\n=== 示例2: 获取贵州茅台2024年周K线（前复权） ===")
	weeklyKlines, err := stock.StockZhAHist("600519", "weekly", "20240101", "20241231", "qfq")
	if err != nil {
		log.Fatalf("获取周K线数据失败: %v", err)
	}

	fmt.Printf("共获取 %d 条周K线数据\n", len(weeklyKlines))
	if len(weeklyKlines) > 0 {
		fmt.Println("\n前3条周K线数据:")
		count := 3
		if len(weeklyKlines) < 3 {
			count = len(weeklyKlines)
		}
		for i := 0; i < count; i++ {
			k := weeklyKlines[i]
			fmt.Printf("%s: 开盘=%.2f, 收盘=%.2f, 涨跌幅=%.2f%%, 成交额=%.2f亿\n",
				k.Date.Format("2006-01-02"),
				k.Open, k.Close, k.ChangePct,
				k.Amount/100000000)
		}
	}

	// 示例3: 获取宁德时代2023年月K线数据（后复权）
	fmt.Println("\n=== 示例3: 获取宁德时代2023年月K线（后复权） ===")
	monthlyKlines, err := stock.StockZhAHist("300750", "monthly", "20230101", "20231231", "hfq")
	if err != nil {
		log.Fatalf("获取月K线数据失败: %v", err)
	}

	fmt.Printf("共获取 %d 条月K线数据\n", len(monthlyKlines))
	if len(monthlyKlines) > 0 {
		// 计算2023年最大涨幅和最大跌幅
		maxChangePct := monthlyKlines[0].ChangePct
		minChangePct := monthlyKlines[0].ChangePct
		maxDate := monthlyKlines[0].Date
		minDate := monthlyKlines[0].Date

		for _, k := range monthlyKlines {
			if k.ChangePct > maxChangePct {
				maxChangePct = k.ChangePct
				maxDate = k.Date
			}
			if k.ChangePct < minChangePct {
				minChangePct = k.ChangePct
				minDate = k.Date
			}
		}

		fmt.Printf("\n2023年月度表现:\n")
		fmt.Printf("最大涨幅: %.2f%% (发生在 %s)\n", maxChangePct, maxDate.Format("2006-01"))
		fmt.Printf("最大跌幅: %.2f%% (发生在 %s)\n", minChangePct, minDate.Format("2006-01"))
	}

	// 示例4: 计算简单移动平均线（SMA）
	fmt.Println("\n=== 示例4: 计算5日均线 ===")
	if len(klines) >= 5 {
		recentKlines := klines[len(klines)-5:]
		sum := 0.0
		for _, k := range recentKlines {
			sum += k.Close
		}
		sma5 := sum / 5.0
		lastClose := klines[len(klines)-1].Close
		fmt.Printf("最近5日收盘均价: %.2f\n", sma5)
		fmt.Printf("最新收盘价: %.2f\n", lastClose)
		if lastClose > sma5 {
			fmt.Printf("当前价格在5日均线上方 (+%.2f%%)\n", (lastClose-sma5)/sma5*100)
		} else {
			fmt.Printf("当前价格在5日均线下方 (%.2f%%)\n", (lastClose-sma5)/sma5*100)
		}
	}
}
