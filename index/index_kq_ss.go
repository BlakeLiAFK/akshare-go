package index

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/BlakeLiAFK/akshare/utils"
)

// KqFashionIndex 柯桥时尚指数
type KqFashionIndex struct {
	Date        time.Time `json:"date"`         // 日期
	Index       float64   `json:"index"`        // 指数
	ChangeValue float64   `json:"change_value"` // 涨跌值
	ChangeRate  float64   `json:"change_rate"`  // 涨跌幅
}

// kqFashionResponse API 响应结构
type kqFashionResponse struct {
	Data []struct {
		PublishTime string  `json:"publishTime"` // 日期
		IndexValue  float64 `json:"indexValue"`  // 指数值
	} `json:"data"`
}

// KqFashionSymbolMap 柯桥时尚指数映射
var KqFashionSymbolMap = map[string]string{
	"柯桥时尚指数":   "root",
	"时尚创意指数":   "01",
	"时尚设计人才数":  "0101",
	"新花型推出数":   "0102",
	"创意产品成交数":  "0103",
	"创意企业数量":   "0104",
	"时尚活跃度指数":  "02",
	"电商运行数":    "0201",
	"时尚平台拓展数":  "0201",
	"新产品销售额占比": "0201",
	"企业合作占比":   "0201",
	"品牌传播费用":   "0201",
	"时尚推广度指数":  "03",
	"国际交流合作次数": "0301",
	"企业参展次数":   "0302",
	"外商驻点数量变化": "0302",
	"时尚评价指数":   "04",
}

// IndexKqFashion 柯桥时尚指数
//
// 数据源: http://ss.kqindex.cn:9559/rinder_web_kqsszs/index/index_page.do
//
// 参数:
//   - symbol: 指数类型，可选:
//     "柯桥时尚指数", "时尚创意指数", "时尚设计人才数", "新花型推出数",
//     "创意产品成交数", "创意企业数量", "时尚活跃度指数", "电商运行数",
//     "时尚平台拓展数", "新产品销售额占比", "企业合作占比", "品牌传播费用",
//     "时尚推广度指数", "国际交流合作次数", "企业参展次数", "外商驻点数量变化",
//     "时尚评价指数"
//
// 返回:
//   - []KqFashionIndex: 柯桥时尚指数数据
//   - error: 错误信息
func IndexKqFashion(symbol string) ([]KqFashionIndex, error) {
	structCode, ok := KqFashionSymbolMap[symbol]
	if !ok {
		return nil, fmt.Errorf("无效的 symbol: %s", symbol)
	}

	apiURL := "http://api.idx365.com/index/project/34/data"

	params := map[string]string{
		"structCode": structCode,
	}

	resp, err := utils.Get(apiURL, params)
	if err != nil {
		return nil, fmt.Errorf("请求柯桥时尚指数失败: %w", err)
	}

	var apiResp kqFashionResponse
	if err := json.Unmarshal(resp.Body(), &apiResp); err != nil {
		return nil, fmt.Errorf("解析柯桥时尚指数响应失败: %w", err)
	}

	if len(apiResp.Data) == 0 {
		return []KqFashionIndex{}, nil
	}

	// 先收集所有数据并排序
	result := make([]KqFashionIndex, 0, len(apiResp.Data))
	for _, item := range apiResp.Data {
		date, err := time.Parse("2006-01-02", item.PublishTime[:10])
		if err != nil {
			continue
		}

		result = append(result, KqFashionIndex{
			Date:  date,
			Index: item.IndexValue,
		})
	}

	// 计算涨跌值和涨跌幅
	for i := 1; i < len(result); i++ {
		result[i].ChangeValue = result[i].Index - result[i-1].Index
		if result[i-1].Index != 0 {
			result[i].ChangeRate = (result[i].Index - result[i-1].Index) / result[i-1].Index
		}
	}

	return result, nil
}
