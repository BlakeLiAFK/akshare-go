package stock

import (
	"encoding/json"
	"fmt"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

// StockInfoEmResponse 东方财富股票信息响应
type StockInfoEmResponse struct {
	Rc   int `json:"rc"`
	Rt   int `json:"rt"`
	Data struct {
		Total int                      `json:"total"`
		Diff  []map[string]interface{} `json:"diff"`
	} `json:"data"`
}

// StockInfoEm 东方财富-股票信息
func StockInfoEm(symbol string) (dataframe.DataFrame, error) {
	if symbol == "" {
		return dataframe.DataFrame{}, fmt.Errorf("股票代码不能为空")
	}

	// 解析股票代码
	code, market := ParseStockCode(symbol)
	secid := fmt.Sprintf("%s.%s", MarketCodeMap[market], code)

	url := "http://push2.eastmoney.com/api/qt/stock/get"
	params := map[string]string{
		"secid":  secid,
		"fields": "f57,f58,f59,f60,f61,f62,f63,f64,f65,f66,f67,f68,f69,f70,f71,f72,f73,f74,f75,f76,f77,f78,f79,f80,f81,f82,f83,f84,f85,f86,f87,f88,f89,f90,f91,f92,f93,f94,f95,f96,f97,f98,f99,f100,f101,f102,f103,f104,f105,f106,f107,f108,f109,f110,f111,f112,f113,f114,f115,f116,f117,f118,f119,f120,f121,f122,f123,f124,f125,f126,f127,f128,f129,f130,f131,f132,f133,f134,f135,f136,f137,f138,f139,f140,f141,f142,f143,f144,f145,f146,f147,f148,f149,f150,f151,f152,f153,f154,f155,f156,f157,f158,f159,f160,f161,f162,f163,f164,f165,f166,f167,f168,f169,f170,f171,f172,f173,f174,f175,f176,f177,f178,f179,f180,f181,f182,f183,f184,f185,f186,f187,f188,f189,f190,f191,f192,f193",
		"ut":     "b2884a393a59ad64002292a3e90d46a5",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("解析JSON失败: %w", err)
	}

	data, ok := result["data"].(map[string]interface{})
	if !ok {
		return dataframe.DataFrame{}, fmt.Errorf("数据格式错误")
	}

	// 构建DataFrame
	var keys []string
	var values []string
	for k, v := range data {
		keys = append(keys, k)
		values = append(values, fmt.Sprintf("%v", v))
	}

	df := dataframe.New(
		series.New(keys, series.String, "field"),
		series.New(values, series.String, "value"),
	)

	return df, nil
}

// StockInfoEmList 东方财富-股票列表
func StockInfoEmList(market string) (dataframe.DataFrame, error) {
	if market == "" {
		market = "沪深A股"
	}

	fsMap := map[string]string{
		"沪深A股": "m:0+t:6,m:0+t:80,m:1+t:2,m:1+t:23",
		"上证A股": "m:1+t:2,m:1+t:23",
		"深证A股": "m:0+t:6,m:0+t:80",
		"创业板":  "m:0+t:80",
		"科创板":  "m:1+t:23",
		"沪深B股": "m:0+t:7,m:1+t:3",
	}

	fs, exists := fsMap[market]
	if !exists {
		return dataframe.DataFrame{}, fmt.Errorf("不支持的市场: %s", market)
	}

	url := "http://push2.eastmoney.com/api/qt/clist/get"
	params := map[string]string{
		"pn":     "1",
		"pz":     "5000",
		"po":     "1",
		"np":     "1",
		"fltt":   "2",
		"invt":   "2",
		"fid":    "f3",
		"fs":     fs,
		"fields": "f12,f14,f2,f3,f4,f5,f6,f7,f15,f16,f17,f18",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	var result StockInfoEmResponse
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("解析JSON失败: %w", err)
	}

	if len(result.Data.Diff) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到数据")
	}

	// 构建DataFrame
	headers := []string{"代码", "名称", "最新价", "涨跌幅", "涨跌额", "成交量", "成交额", "振幅", "最高", "最低", "今开", "昨收"}
	var records [][]string
	records = append(records, headers)

	for _, item := range result.Data.Diff {
		record := []string{
			getString(item, "f12"),
			getString(item, "f14"),
			getString(item, "f2"),
			getString(item, "f3"),
			getString(item, "f4"),
			getString(item, "f5"),
			getString(item, "f6"),
			getString(item, "f7"),
			getString(item, "f15"),
			getString(item, "f16"),
			getString(item, "f17"),
			getString(item, "f18"),
		}
		records = append(records, record)
	}

	df := dataframe.LoadRecords(records)
	return df, nil
}

func getString(m map[string]interface{}, key string) string {
	if v, ok := m[key]; ok {
		return fmt.Sprintf("%v", v)
	}
	return ""
}
