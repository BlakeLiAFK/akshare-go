package index

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/BlakeLiAFK/akshare/utils"
)

// CxIndex 财新指数数据
type CxIndex struct {
	Date      time.Time `json:"date"`       // 日期
	Value     float64   `json:"value"`      // 指数值
	Change    float64   `json:"change"`     // 变化值/变化幅度
	IndexName string    `json:"index_name"` // 指数名称
}

// cxIndexResponse 财新指数 API 响应
type cxIndexResponse struct {
	Data [][]interface{} `json:"data"`
}

// fetchCxIndex 通用财新指数获取函数
func fetchCxIndex(indexType, indexName string, extraParams map[string]string) ([]CxIndex, error) {
	apiURL := "https://yun.ccxe.com.cn/api/index/pro/cxIndexTrendInfo"

	params := map[string]string{
		"type": indexType,
	}
	for k, v := range extraParams {
		params[k] = v
	}

	respBody, err := utils.Get(apiURL, params)
	if err != nil {
		return nil, fmt.Errorf("请求财新%s失败: %w", indexName, err)
	}

	var resp cxIndexResponse
	if err := json.Unmarshal(respBody.Body(), &resp); err != nil {
		return nil, fmt.Errorf("解析财新%s响应失败: %w", indexName, err)
	}

	result := make([]CxIndex, 0, len(resp.Data))
	for _, item := range resp.Data {
		if len(item) < 3 {
			continue
		}

		// 日期是毫秒时间戳
		var date time.Time
		if ts, ok := item[2].(float64); ok {
			date = time.UnixMilli(int64(ts))
		}

		var value, change float64
		if v, ok := item[1].(float64); ok {
			value = v
		}
		if v, ok := item[0].(float64); ok {
			change = v
		}

		result = append(result, CxIndex{
			Date:      date,
			Value:     value,
			Change:    change,
			IndexName: indexName,
		})
	}

	return result, nil
}

// IndexPmiComCx 财新数据-指数报告-综合PMI
func IndexPmiComCx() ([]CxIndex, error) {
	return fetchCxIndex("com", "综合PMI", nil)
}

// IndexPmiManCx 财新数据-指数报告-制造业PMI
func IndexPmiManCx() ([]CxIndex, error) {
	return fetchCxIndex("man", "制造业PMI", nil)
}

// IndexPmiSerCx 财新数据-指数报告-服务业PMI
func IndexPmiSerCx() ([]CxIndex, error) {
	return fetchCxIndex("ser", "服务业PMI", nil)
}

// IndexDeiCx 财新数据-指数报告-数字经济指数
func IndexDeiCx() ([]CxIndex, error) {
	return fetchCxIndex("dei", "数字经济指数", nil)
}

// IndexIiCx 财新数据-指数报告-产业指数
func IndexIiCx() ([]CxIndex, error) {
	return fetchCxIndex("ii", "产业指数", nil)
}

// IndexSiCx 财新数据-指数报告-溢出指数
func IndexSiCx() ([]CxIndex, error) {
	return fetchCxIndex("si", "溢出指数", nil)
}

// IndexFiCx 财新数据-指数报告-融合指数
func IndexFiCx() ([]CxIndex, error) {
	return fetchCxIndex("fi", "融合指数", nil)
}

// IndexBiCx 财新数据-指数报告-基础指数
func IndexBiCx() ([]CxIndex, error) {
	return fetchCxIndex("bi", "基础指数", nil)
}

// IndexNeiCx 财新数据-指数报告-中国新经济指数
func IndexNeiCx() ([]CxIndex, error) {
	return fetchCxIndex("nei", "中国新经济指数", nil)
}

// IndexLiCx 财新数据-指数报告-劳动力投入指数
func IndexLiCx() ([]CxIndex, error) {
	return fetchCxIndex("li", "劳动力投入指数", nil)
}

// IndexCiCx 财新数据-指数报告-资本投入指数
func IndexCiCx() ([]CxIndex, error) {
	return fetchCxIndex("ci", "资本投入指数", nil)
}

// IndexTiCx 财新数据-指数报告-科技投入指数
func IndexTiCx() ([]CxIndex, error) {
	return fetchCxIndex("ti", "科技投入指数", nil)
}

// IndexNeawCx 财新数据-指数报告-新经济行业入职平均工资水平
func IndexNeawCx() ([]CxIndex, error) {
	return fetchCxIndex("neaw", "新经济行业入职平均工资水平", nil)
}

// IndexAwprCx 财新数据-指数报告-新经济入职工资溢价水平
func IndexAwprCx() ([]CxIndex, error) {
	return fetchCxIndex("awpr", "新经济入职工资溢价水平", nil)
}

// IndexCciCx 财新数据-指数报告-大宗商品指数
func IndexCciCx() ([]CxIndex, error) {
	return fetchCxIndex("cci", "大宗商品指数", map[string]string{
		"code":  "1000050",
		"month": "-1",
	})
}

// IndexQliCx 财新数据-指数报告-高质量因子指数
func IndexQliCx() ([]CxIndex, error) {
	return fetchCxIndex("qli", "高质量因子指数", map[string]string{
		"code":  "1000050",
		"month": "-1",
	})
}

// IndexAiCx 财新数据-指数报告-AI策略指数
func IndexAiCx() ([]CxIndex, error) {
	return fetchCxIndex("ai", "AI策略指数", map[string]string{
		"code":  "1000050",
		"month": "-1",
	})
}

// IndexBeiCx 财新数据-指数报告-基石经济指数
func IndexBeiCx() ([]CxIndex, error) {
	return fetchCxIndex("ind", "基石经济指数", map[string]string{
		"code":  "930927",
		"month": "-1",
	})
}

// IndexNeeiCx 财新数据-指数报告-新动能指数
func IndexNeeiCx() ([]CxIndex, error) {
	return fetchCxIndex("ind", "新动能指数", map[string]string{
		"code":  "930928",
		"month": "1",
	})
}
