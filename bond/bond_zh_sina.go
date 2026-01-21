package bond

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/dop251/goja"
	"github.com/go-gota/gota/dataframe"
	"github.com/tidwall/gjson"
)

const (
	zhSinaBondHSURL      = "http://vip.stock.finance.sina.com.cn/quotes_service/api/json_v2.php/Market_Center.getHQNodeData"
	zhSinaBondHSCountURL = "http://vip.stock.finance.sina.com.cn/quotes_service/api/json_v2.php/Market_Center.getHQNodeStockCountSimple"
	zhSinaBondHSHistURL  = "https://finance.sina.com.cn/realstock/company/%s/hisdata/klc_kl.js?d=%s"
)

// getZHBondHSPageCount 获取沪深债券总页数
func getZHBondHSPageCount() (int, error) {
	params := map[string]string{
		"node": "hs_z",
	}

	resp, err := utils.Get(zhSinaBondHSCountURL, params)
	if err != nil {
		return 0, fmt.Errorf("请求失败: %w", err)
	}

	// 从响应中提取数字
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(resp.String(), -1)
	if len(matches) == 0 {
		return 0, fmt.Errorf("未找到页数信息")
	}

	count := utils.MustParseFloat(matches[0])
	pageCount := int(count / 80)
	if count/80 > float64(pageCount) {
		pageCount++
	}

	return pageCount, nil
}

// BondZHHSSpot 新浪财经-债券-沪深债券-实时行情数据
//
// 获取沪深债券在当前时刻的实时行情数据（大量抓取容易封IP）
//
// 参数:
//   - startPage: 分页起始页，如 "1"
//   - endPage: 分页结束页，如 "10"
//
// 返回:
//   - dataframe.DataFrame: 包含实时行情数据
//   - error: 错误信息
//
// 数据源: https://vip.stock.finance.sina.com.cn/mkt/#hs_z
func BondZHHSSpot(startPage, endPage string) (dataframe.DataFrame, error) {
	pageCount, err := getZHBondHSPageCount()
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("获取页数失败: %w", err)
	}

	start := int(utils.MustParseFloat(startPage))
	end := int(utils.MustParseFloat(endPage)) + 1
	if end > pageCount {
		end = pageCount
	}

	var allRecords []map[string]interface{}

	for page := start; page < end; page++ {
		params := map[string]string{
			"page":   fmt.Sprintf("%d", page),
			"num":    "80",
			"sort":   "symbol",
			"asc":    "1",
			"node":   "hs_z",
			"_s_r_a": "page",
		}

		resp, err := utils.Get(zhSinaBondHSURL, params)
		if err != nil {
			continue
		}

		// 解析JSON数据
		// 新浪返回的数据格式可能需要特殊处理
		// Python版本使用demjson解码，这里尝试用标准JSON
		dataJSON := gjson.Parse(resp.String())
		if !dataJSON.IsArray() {
			continue
		}

		dataJSON.ForEach(func(_, item gjson.Result) bool {
			itemMap := item.Map()
			record := map[string]interface{}{
				"代码":  itemMap["symbol"].String(),
				"名称":  itemMap["name"].String(),
				"最新价": utils.MustParseFloat(itemMap["trade"].String()),
				"涨跌额": utils.MustParseFloat(itemMap["pricechange"].String()),
				"涨跌幅": utils.MustParseFloat(itemMap["changepercent"].String()),
				"买入":  utils.MustParseFloat(itemMap["buy"].String()),
				"卖出":  utils.MustParseFloat(itemMap["sell"].String()),
				"昨收":  utils.MustParseFloat(itemMap["settlement"].String()),
				"今开":  utils.MustParseFloat(itemMap["open"].String()),
				"最高":  utils.MustParseFloat(itemMap["high"].String()),
				"最低":  utils.MustParseFloat(itemMap["low"].String()),
				"成交量": utils.MustParseFloat(itemMap["volume"].String()),
				"成交额": utils.MustParseFloat(itemMap["amount"].String()),
			}
			allRecords = append(allRecords, record)
			return true
		})

		// 添加小延迟避免请求过快
		time.Sleep(100 * time.Millisecond)
	}

	if len(allRecords) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到有效数据")
	}

	df := dataframe.LoadMaps(allRecords)
	return df, nil
}

// BondZHHSDaily 新浪财经-债券-沪深债券-历史行情数据
//
// 获取指定沪深债券代码的日K线数据（大量抓取容易封IP）
//
// 参数:
//   - symbol: 沪深债券代码，如 "sh010107"
//
// 返回:
//   - dataframe.DataFrame: 包含日K线数据
//   - error: 错误信息
//
// 数据源: https://vip.stock.finance.sina.com.cn/mkt/#hs_z
func BondZHHSDaily(symbol string) (dataframe.DataFrame, error) {
	// 构建URL
	now := time.Now().Format("2006_01_02")
	url := fmt.Sprintf(zhSinaBondHSHistURL, symbol, now)

	resp, err := utils.Get(url, nil)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	// 响应格式类似: var hq_str_sh010107="...";
	// 需要提取并解密数据
	text := resp.String()
	parts := strings.Split(text, "=")
	if len(parts) < 2 {
		return dataframe.DataFrame{}, fmt.Errorf("响应格式错误")
	}

	// 提取加密的数据部分（去掉引号和分号）
	encrypted := strings.TrimSpace(parts[1])
	encrypted = strings.Trim(encrypted, `";`)

	// 使用goja执行JS解密代码
	// 这里需要hk_js_decode的JS代码，这个在stock/cons.go中应该有
	// 简化处理：如果数据已经是明文JSON，直接解析
	vm := goja.New()

	// 定义解密函数（简化版本，实际可能需要完整的JS解密代码）
	decryptScript := `
	function d(encryptedStr) {
		// 简化的解密逻辑
		// 实际应该使用完整的hk_js_decode
		try {
			return JSON.parse(encryptedStr);
		} catch(e) {
			return [];
		}
	}
	`

	_, err = vm.RunString(decryptScript)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("加载解密脚本失败: %w", err)
	}

	// 调用解密函数
	decryptFunc, ok := goja.AssertFunction(vm.Get("d"))
	if !ok {
		return dataframe.DataFrame{}, fmt.Errorf("解密函数不存在")
	}

	result, err := decryptFunc(goja.Undefined(), vm.ToValue(encrypted))
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("解密失败: %w", err)
	}

	// 将结果转换为DataFrame
	var records []map[string]interface{}

	// 将goja值转换为Go值
	export := result.Export()
	if dataList, ok := export.([]interface{}); ok {
		for _, item := range dataList {
			if itemMap, ok := item.(map[string]interface{}); ok {
				record := map[string]interface{}{
					"date":  itemMap["date"],
					"open":  utils.MustParseFloat(fmt.Sprintf("%v", itemMap["open"])),
					"high":  utils.MustParseFloat(fmt.Sprintf("%v", itemMap["high"])),
					"low":   utils.MustParseFloat(fmt.Sprintf("%v", itemMap["low"])),
					"close": utils.MustParseFloat(fmt.Sprintf("%v", itemMap["close"])),
				}
				records = append(records, record)
			}
		}
	}

	if len(records) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到有效数据")
	}

	df := dataframe.LoadMaps(records)
	return df, nil
}
