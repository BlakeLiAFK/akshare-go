package fx

import (
	"fmt"
	"strings"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/tidwall/gjson"
)

const (
	// 百度外汇行情URL
	baiduFxURL = "https://finance.pae.baidu.com/api/getforeignrank"
)

// FxQuoteBaiduItem 百度外汇行情数据结构
type FxQuoteBaiduItem struct {
	Code      string  `json:"code"`       // 代码
	Name      string  `json:"name"`       // 名称
	Price     float64 `json:"price"`      // 最新价
	Change    float64 `json:"change"`     // 涨跌额
	ChangePct float64 `json:"change_pct"` // 涨跌幅
}

// FxQuoteBaidu 获取百度股市通-外汇-行情榜单
//
// 目标地址: https://gushitong.baidu.com/top/foreign-rmb
//
// 参数:
//   - symbol: 货币类型，可选 "人民币" 或 "美元"
//
// 返回:
//   - []FxQuoteBaiduItem: 外汇行情数据列表
//   - error: 错误信息
//
// 示例:
//
//	data, err := fx.FxQuoteBaidu("人民币")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, item := range data {
//	    fmt.Printf("%s %s %.4f %.4f %.4f%%\n", item.Code, item.Name, item.Price, item.Change, item.ChangePct*100)
//	}
func FxQuoteBaidu(symbol string) ([]FxQuoteBaiduItem, error) {
	symbolMap := map[string]string{
		"人民币": "rmb",
		"美元":  "dollar",
	}

	typeVal, ok := symbolMap[symbol]
	if !ok {
		return nil, fmt.Errorf("无效的symbol参数，可选: 人民币, 美元")
	}

	headers := map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
	}

	var allItems []FxQuoteBaiduItem
	pn := 0

	for {
		params := map[string]string{
			"type":          typeVal,
			"pn":            fmt.Sprintf("%d", pn),
			"rn":            "20",
			"finClientType": "pc",
		}

		resp, err := utils.GetWithHeaders(baiduFxURL, params, headers)
		if err != nil {
			break
		}

		text := resp.String()
		result := gjson.Get(text, "Result")
		if !result.Exists() || !result.IsArray() || len(result.Array()) == 0 {
			break
		}

		result.ForEach(func(_, value gjson.Result) bool {
			code := value.Get("code").String()
			name := value.Get("name").String()
			list := value.Get("list")

			var price, change, changePct float64
			if list.Exists() && list.IsArray() {
				list.ForEach(func(_, item gjson.Result) bool {
					key := ""
					val := ""
					item.ForEach(func(k, v gjson.Result) bool {
						if key == "" {
							key = v.String()
						} else {
							val = v.String()
						}
						return true
					})
					switch key {
					case "最新价":
						price = utils.MustFloat64(val)
					case "涨跌额":
						change = utils.MustFloat64(val)
					case "涨跌幅":
						changePct = utils.MustFloat64(strings.TrimSuffix(val, "%")) / 100
					}
					return true
				})
			}

			allItems = append(allItems, FxQuoteBaiduItem{
				Code:      code,
				Name:      name,
				Price:     price,
				Change:    change,
				ChangePct: changePct,
			})
			return true
		})

		pn += 20
		if len(result.Array()) < 20 {
			break
		}
	}

	return allItems, nil
}
