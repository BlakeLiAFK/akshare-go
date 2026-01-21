package fund

import (
	"fmt"
	"strings"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/tidwall/gjson"
)

// FundScaleOpenSina 获取新浪开放式基金规模数据
// https://vip.stock.finance.sina.com.cn/fund_center/index.html#jjhqetf
// 参数: symbol choice of {"股票型基金", "混合型基金", "债券型基金", "货币型基金", "QDII基金"}
func FundScaleOpenSina(symbol string) ([]map[string]interface{}, error) {
	if symbol == "" {
		symbol = "股票型基金"
	}

	fundMap := map[string]string{
		"股票型基金":  "2",
		"混合型基金":  "1",
		"债券型基金":  "3",
		"货币型基金":  "5",
		"QDII基金": "6",
	}

	fundType, ok := fundMap[symbol]
	if !ok {
		return nil, fmt.Errorf("请输入正确的基金类型参数: 股票型基金, 混合型基金, 债券型基金, 货币型基金, QDII基金")
	}

	url := "http://vip.stock.finance.sina.com.cn/fund_center/data/jsonp.php/IO.XSRV2.CallbackList['J2cW8KXheoWKdSHc']/NetValueReturn_Service.NetValueReturnOpen"
	params := map[string]string{
		"page":  "1",
		"num":   "10000",
		"sort":  "zmjgm",
		"asc":   "0",
		"ccode": "",
		"type2": fundType,
		"type3": "",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	// 解析JSONP格式: IO.XSRV2.CallbackList['xxx']({...})
	text := resp.String()
	startIdx := strings.Index(text, "({")
	endIdx := strings.LastIndex(text, "})")
	if startIdx == -1 || endIdx == -1 {
		return nil, fmt.Errorf("解析JSONP失败")
	}
	jsonText := text[startIdx+1 : endIdx+1]

	result := gjson.Parse(jsonText)
	items := result.Get("data").Array()

	records := make([]map[string]interface{}, 0, len(items))
	for i, item := range items {
		record := map[string]interface{}{
			"序号":    i + 1,
			"基金代码":  item.Get("symbol").String(),
			"基金简称":  item.Get("sname").String(),
			"单位净值":  utils.MustFloat64(item.Get("dwjz").String()),
			"总募集规模": utils.MustFloat64(item.Get("zmjgm").String()),
			"最近总份额": utils.MustFloat64(item.Get("zjzfe").String()),
			"成立日期":  item.Get("clrq").String(),
			"基金经理":  item.Get("jjjl").String(),
			"更新日期":  item.Get("jzrq").String(),
		}
		records = append(records, record)
	}

	return records, nil
}

// FundScaleCloseSina 获取新浪封闭式基金规模数据
// https://vip.stock.finance.sina.com.cn/fund_center/index.html#jjhqetf
func FundScaleCloseSina() ([]map[string]interface{}, error) {
	url := "http://vip.stock.finance.sina.com.cn/fund_center/data/jsonp.php/IO.XSRV2.CallbackList['_bjN6KvXOkfPy2Bu']/NetValueReturn_Service.NetValueReturnClose"
	params := map[string]string{
		"page":  "1",
		"num":   "1000",
		"sort":  "zmjgm",
		"asc":   "0",
		"ccode": "",
		"type2": "",
		"type3": "",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	// 解析JSONP格式
	text := resp.String()
	startIdx := strings.Index(text, "({")
	endIdx := strings.LastIndex(text, "})")
	if startIdx == -1 || endIdx == -1 {
		return nil, fmt.Errorf("解析JSONP失败")
	}
	jsonText := text[startIdx+1 : endIdx+1]

	result := gjson.Parse(jsonText)
	items := result.Get("data").Array()

	records := make([]map[string]interface{}, 0, len(items))
	for i, item := range items {
		record := map[string]interface{}{
			"序号":    i + 1,
			"基金代码":  item.Get("symbol").String(),
			"基金简称":  item.Get("sname").String(),
			"单位净值":  utils.MustFloat64(item.Get("dwjz").String()),
			"总募集规模": utils.MustFloat64(item.Get("zmjgm").String()),
			"最近总份额": utils.MustFloat64(item.Get("zjzfe").String()),
			"成立日期":  item.Get("clrq").String(),
			"基金经理":  item.Get("jjjl").String(),
			"更新日期":  item.Get("jzrq").String(),
		}
		records = append(records, record)
	}

	return records, nil
}

// FundScaleStructuredSina 获取新浪分级基金规模数据
// https://vip.stock.finance.sina.com.cn/fund_center/index.html#jjgmfjall
func FundScaleStructuredSina() ([]map[string]interface{}, error) {
	url := "http://vip.stock.finance.sina.com.cn/fund_center/data/jsonp.php/IO.XSRV2.CallbackList['cRrwseM7NWX68rDa']/NetValueReturn_Service.NetValueReturnCX"
	params := map[string]string{
		"page":  "1",
		"num":   "1000",
		"sort":  "zmjgm",
		"asc":   "0",
		"ccode": "",
		"type2": "",
		"type3": "",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	// 解析JSONP格式
	text := resp.String()
	startIdx := strings.Index(text, "({")
	endIdx := strings.LastIndex(text, "})")
	if startIdx == -1 || endIdx == -1 {
		return nil, fmt.Errorf("解析JSONP失败")
	}
	jsonText := text[startIdx+1 : endIdx+1]

	result := gjson.Parse(jsonText)
	items := result.Get("data").Array()

	records := make([]map[string]interface{}, 0, len(items))
	for i, item := range items {
		record := map[string]interface{}{
			"序号":    i + 1,
			"基金代码":  item.Get("symbol").String(),
			"基金简称":  item.Get("sname").String(),
			"单位净值":  utils.MustFloat64(item.Get("dwjz").String()),
			"总募集规模": utils.MustFloat64(item.Get("zmjgm").String()),
			"最近总份额": utils.MustFloat64(item.Get("zjzfe").String()),
			"成立日期":  item.Get("clrq").String(),
			"基金经理":  item.Get("jjjl").String(),
			"更新日期":  item.Get("jzrq").String(),
		}
		records = append(records, record)
	}

	return records, nil
}
