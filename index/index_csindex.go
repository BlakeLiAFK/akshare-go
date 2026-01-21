package index

import (
	"fmt"

	"github.com/BlakeLiAFK/akshare/utils"
)

// IndexCSIndexAll 中证指数网站-指数列表
//
// 返回:
//   - []CSIndexInfo: 中证指数信息列表
//   - error: 错误信息
func IndexCSIndexAll() ([]CSIndexInfo, error) {
	// 中证指数需要POST请求
	payload := `{
		"sorter": {"sortField": "null", "sortOrder": null},
		"pager": {"pageNum": 1, "pageSize": 5000},
		"indexFilter": {
			"ifCustomized": null,
			"ifTracked": null,
			"ifWeightCapped": null,
			"indexCompliance": null,
			"hotSpot": null,
			"indexClassify": null,
			"currency": null,
			"region": null,
			"indexSeries": ["1"],
			"undefined": null
		}
	}`

	headers := map[string]string{
		"Content-Type": "application/json;charset=UTF-8",
		"Referer":      "https://www.csindex.com.cn/",
	}

	_, err := utils.PostWithHeaders(CSIndexExportURL, []byte(payload), headers)
	if err != nil {
		return nil, fmt.Errorf("请求中证指数列表失败: %w", err)
	}

	// 中证指数返回的是Excel文件，这里简化处理
	// 实际实现需要使用Excel解析库
	return []CSIndexInfo{}, fmt.Errorf("中证指数需要Excel解析，暂未实现")
}

// IndexStockConsCSIndex 中证指数网站-成份股目录
//
// 参数:
//   - symbol: 指数代码，如 "000300"
//
// 返回:
//   - []IndexStockConsWeightCSIndex: 成份股列表
//   - error: 错误信息
func IndexStockConsCSIndex(symbol string) ([]IndexStockConsWeightCSIndex, error) {
	url := fmt.Sprintf("https://oss-ch.csindex.com.cn/static/html/csindex/public/uploads/file/autofile/cons/%scons.xls", symbol)

	headers := map[string]string{
		"Referer": "https://www.csindex.com.cn/",
	}
	_, err := utils.GetWithHeaders(url, nil, headers)
	if err != nil {
		return nil, fmt.Errorf("请求中证指数成份股失败: %w", err)
	}

	// 中证指数返回的是Excel文件，这里简化处理
	return []IndexStockConsWeightCSIndex{}, fmt.Errorf("中证指数成份股需要Excel解析，暂未实现")
}

// IndexStockConsWeightCSI 中证指数网站-样本权重
//
// 参数:
//   - symbol: 指数代码，如 "000300"
//
// 返回:
//   - []IndexStockConsWeightCSIndex: 成份股权重列表
//   - error: 错误信息
func IndexStockConsWeightCSI(symbol string) ([]IndexStockConsWeightCSIndex, error) {
	url := fmt.Sprintf("https://oss-ch.csindex.com.cn/static/html/csindex/public/uploads/file/autofile/closeweight/%scloseweight.xls", symbol)

	headers := map[string]string{
		"Referer": "https://www.csindex.com.cn/",
	}
	_, err := utils.GetWithHeaders(url, nil, headers)
	if err != nil {
		return nil, fmt.Errorf("请求中证指数成份股权重失败: %w", err)
	}

	// 中证指数返回的是Excel文件，这里简化处理
	return []IndexStockConsWeightCSIndex{}, fmt.Errorf("中证指数成份股权重需要Excel解析，暂未实现")
}
