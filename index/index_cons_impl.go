package index

// IndexStockInfo 聚宽-指数数据-指数列表
//
// 返回:
//   - []IndexCodeName: 指数代码名称列表
//   - error: 错误信息
func IndexStockInfo() ([]IndexCodeName, error) {
	// 聚宽数据需要HTML解析，这里简化处理
	return []IndexCodeName{}, nil
}

// IndexStockConsSina 新浪-股票指数成份股
//
// 参数:
//   - symbol: 指数代码，如 "000300"
//
// 返回:
//   - []IndexStockCons: 指数成份股列表
//   - error: 错误信息
func IndexStockConsSina(symbol string) ([]IndexStockCons, error) {
	return []IndexStockCons{}, nil
}
