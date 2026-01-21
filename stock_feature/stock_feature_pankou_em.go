package stock_feature

import (
	"fmt"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/tidwall/gjson"
)

// StockChangesItem 盘口异动项
type StockChangesItem struct {
	Time  string `json:"time"`  // 时间
	Code  string `json:"code"`  // 代码
	Name  string `json:"name"`  // 名称
	Board string `json:"board"` // 板块
	Info  string `json:"info"`  // 相关信息
}

// StockBoardChangeItem 板块异动项
type StockBoardChangeItem struct {
	BoardName           string                 `json:"board_name"`            // 板块名称
	ChangeRate          float64                `json:"change_rate"`           // 涨跌幅
	MainNetInflow       float64                `json:"main_net_inflow"`       // 主力净流入
	TotalChangeCount    int                    `json:"total_change_count"`    // 板块异动总次数
	MostActiveCode      string                 `json:"most_active_code"`      // 最频繁个股代码
	MostActiveName      string                 `json:"most_active_name"`      // 最频繁个股名称
	MostActiveDirection string                 `json:"most_active_direction"` // 买卖方向
	ChangeTypeList      map[string]interface{} `json:"change_type_list"`      // 异动类型列表
}

// 盘口异动类型映射
var stockChangesSymbolMap = map[string]string{
	"火箭发射":    "8201",
	"快速反弹":    "8202",
	"大笔买入":    "8193",
	"封涨停板":    "4",
	"打开跌停板":   "32",
	"有大买盘":    "64",
	"竞价上涨":    "8207",
	"高开5日线":   "8209",
	"向上缺口":    "8211",
	"60日新高":   "8213",
	"60日大幅上涨": "8215",
	"加速下跌":    "8204",
	"高台跳水":    "8203",
	"大笔卖出":    "8194",
	"封跌停板":    "8",
	"打开涨停板":   "16",
	"有大卖盘":    "128",
	"竞价下跌":    "8208",
	"低开5日线":   "8210",
	"向下缺口":    "8212",
	"60日新低":   "8214",
	"60日大幅下跌": "8216",
}

var stockChangesSymbolReverseMap map[string]string

func init() {
	stockChangesSymbolReverseMap = make(map[string]string)
	for k, v := range stockChangesSymbolMap {
		stockChangesSymbolReverseMap[v] = k
	}
}

// StockChangesEm 东方财富-行情中心-盘口异动
// symbol: 异动类型，可选 "火箭发射", "快速反弹", "大笔买入", "封涨停板", "打开跌停板", "有大买盘",
//
//	"竞价上涨", "高开5日线", "向上缺口", "60日新高", "60日大幅上涨", "加速下跌", "高台跳水",
//	"大笔卖出", "封跌停板", "打开涨停板", "有大卖盘", "竞价下跌", "低开5日线", "向下缺口", "60日新低", "60日大幅下跌"
func StockChangesEm(symbol string) ([]StockChangesItem, error) {
	typeCode, ok := stockChangesSymbolMap[symbol]
	if !ok {
		return nil, fmt.Errorf("不支持的异动类型: %s", symbol)
	}

	url := "https://push2ex.eastmoney.com/getAllStockChanges"
	params := map[string]string{
		"type":      typeCode,
		"pageindex": "0",
		"pagesize":  "5000",
		"ut":        "7eea3edcaed734bea9cbfc24409ed989",
		"dpt":       "wzchanges",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return nil, fmt.Errorf("获取盘口异动失败: %w", err)
	}

	result := gjson.ParseBytes(resp.Body())
	dataArr := result.Get("data.allstock").Array()

	var items []StockChangesItem
	for _, item := range dataArr {
		timeStr := item.Get("tm").String()
		if len(timeStr) == 6 {
			timeStr = timeStr[:2] + ":" + timeStr[2:4] + ":" + timeStr[4:]
		}

		boardCode := item.Get("hy").String()
		boardName := stockChangesSymbolReverseMap[boardCode]
		if boardName == "" {
			boardName = boardCode
		}

		items = append(items, StockChangesItem{
			Time:  timeStr,
			Code:  item.Get("c").String(),
			Name:  item.Get("n").String(),
			Board: boardName,
			Info:  item.Get("i").String(),
		})
	}

	return items, nil
}

// StockBoardChangeEm 东方财富-行情中心-当日板块异动详情
func StockBoardChangeEm() ([]StockBoardChangeItem, error) {
	url := "https://push2ex.eastmoney.com/getAllBKChanges"
	params := map[string]string{
		"ut":        "7eea3edcaed734bea9cbfc24409ed989",
		"dpt":       "wzchanges",
		"pageindex": "0",
		"pagesize":  "5000",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return nil, fmt.Errorf("获取板块异动失败: %w", err)
	}

	result := gjson.ParseBytes(resp.Body())
	dataArr := result.Get("data.allbk").Array()

	var items []StockBoardChangeItem
	for _, item := range dataArr {
		ms := item.Get("ms")
		direction := "大笔买入"
		if ms.Get("m").Int() == 1 {
			direction = "大笔卖出"
		}

		var changeTypeList map[string]interface{}
		if item.Get("cl").Exists() {
			changeTypeList = item.Get("cl").Value().(map[string]interface{})
		}

		items = append(items, StockBoardChangeItem{
			BoardName:           item.Get("n").String(),
			ChangeRate:          item.Get("zdf").Float(),
			MainNetInflow:       item.Get("zlr").Float(),
			TotalChangeCount:    int(item.Get("c").Int()),
			MostActiveCode:      ms.Get("c").String(),
			MostActiveName:      ms.Get("n").String(),
			MostActiveDirection: direction,
			ChangeTypeList:      changeTypeList,
		})
	}

	return items, nil
}
