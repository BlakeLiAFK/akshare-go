package index

import (
	"fmt"
)

// IndexAllCNI 国证指数-最近交易日的所有指数
//
// 返回:
//   - []CNIIndexInfo: 国证指数信息列表
//   - error: 错误信息
func IndexAllCNI() ([]CNIIndexInfo, error) {
	return []CNIIndexInfo{}, fmt.Errorf("国证指数数据解析较复杂，暂未实现")
}

// IndexHistCNI 国证指数历史行情数据
//
// 参数:
//   - symbol: 指数代码，如 "399001"
//   - startDate: 开始时间，格式 "20230114"
//   - endDate: 结束时间，格式 "20240114"
//
// 返回:
//   - []GlobalIndexKLine: 国证指数K线数据
//   - error: 错误信息
func IndexHistCNI(symbol, startDate, endDate string) ([]GlobalIndexKLine, error) {
	return []GlobalIndexKLine{}, fmt.Errorf("国证指数历史数据解析较复杂，暂未实现")
}

// IndexDetailCNI 国证指数-样本详情-指定日期的样本成份
//
// 参数:
//   - symbol: 指数代码，如 "399001"
//
// 返回:
//   - []CNIIndexDetail: 国证指数成份股详情
//   - error: 错误信息
func IndexDetailCNI(symbol string) ([]CNIIndexDetail, error) {
	return []CNIIndexDetail{}, fmt.Errorf("国证指数成份股需要Excel解析，暂未实现")
}

// IndexDetailHistCNI 国证指数-样本详情-历史样本
//
// 参数:
//   - symbol: 指数代码，如 "399101"
//
// 返回:
//   - []CNIIndexDetail: 国证指数历史样本
//   - error: 错误信息
func IndexDetailHistCNI(symbol string) ([]CNIIndexDetail, error) {
	return []CNIIndexDetail{}, fmt.Errorf("国证指数历史样本需要Excel解析，暂未实现")
}

// IndexDetailHistAdjustCNI 国证指数-样本详情-历史调样
//
// 参数:
//   - symbol: 指数代码，如 "399005"
//
// 返回:
//   - []CNIIndexDetail: 国证指数历史调样
//   - error: 错误信息
func IndexDetailHistAdjustCNI(symbol string) ([]CNIIndexDetail, error) {
	return []CNIIndexDetail{}, fmt.Errorf("国证指数历史调样需要Excel解析，暂未实现")
}
