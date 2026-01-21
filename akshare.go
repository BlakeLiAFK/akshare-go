// Package akshare 提供中国金融数据接口的 Go 语言实现
//
// akshare-go 是 Python akshare 库的 Go 语言完整复刻版本，
// 提供股票、基金、期货、债券、指数、期权、宏观经济等金融数据接口。
//
// 基本用法:
//
//	import "github.com/BlakeLiAFK/akshare"
//
//	// 获取股票数据
//	data, err := akshare.StockZhASpotEm()
//
// 更多示例请参考 examples 目录。
package akshare

// Version 当前版本号
const Version = "0.1.0"
