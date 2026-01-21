package fx

import (
	"fmt"
	"time"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/tidwall/gjson"
)

// FxCSwapItem 外汇掉期C-Swap定盘曲线数据结构
type FxCSwapItem struct {
	CurveTime  string  `json:"curve_time"`   // 日期时间
	Tenor      string  `json:"tenor"`        // 期限品种
	SwapPnt    float64 `json:"swap_pnt"`     // 掉期点(Pips)
	DataSource string  `json:"data_source"`  // 掉期点数据源
	SwapAllPrc float64 `json:"swap_all_prc"` // 全价汇率
}

// FxCSwapCm 获取外汇掉期C-Swap定盘曲线
//
// 目标地址: https://www.chinamoney.org.cn/chinese/bkcurvfsw
//
// 返回:
//   - []FxCSwapItem: 外汇掉期C-Swap定盘曲线数据列表
//   - error: 错误信息
//
// 示例:
//
//	data, err := fx.FxCSwapCm()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, item := range data {
//	    fmt.Printf("%s %s %.2f %.4f\n", item.CurveTime, item.Tenor, item.SwapPnt, item.SwapAllPrc)
//	}
func FxCSwapCm() ([]FxCSwapItem, error) {
	params := map[string]string{
		"t": fmt.Sprintf("%d", time.Now().UnixMilli()),
	}

	headers := map[string]string{
		"User-Agent": ShortUserAgent,
	}

	resp, err := utils.PostFormWithHeaders(FxCSwapURL, params, headers)
	if err != nil {
		return nil, fmt.Errorf("获取外汇掉期C-Swap定盘曲线失败: %w", err)
	}

	text := resp.String()
	records := gjson.Get(text, "records")
	if !records.Exists() {
		return nil, fmt.Errorf("解析数据失败")
	}

	var items []FxCSwapItem
	records.ForEach(func(_, value gjson.Result) bool {
		item := FxCSwapItem{
			CurveTime:  value.Get("curveTime").String(),
			Tenor:      value.Get("tenor").String(),
			SwapPnt:    value.Get("swapPnt").Float(),
			DataSource: value.Get("dataSource").String(),
			SwapAllPrc: value.Get("swapAllPrc").Float(),
		}
		items = append(items, item)
		return true
	})

	return items, nil
}
