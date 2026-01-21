package bond

import (
	"fmt"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/go-gota/gota/dataframe"
	"github.com/tidwall/gjson"
)

// BondZHCovInfoTHS 同花顺可转债数据
//
// 获取同花顺数据中心的可转债信息
//
// 参数:
//   - 无
//
// 返回:
//   - dataframe.DataFrame: 包含可转债数据
//   - error: 错误信息
//
// 数据源: https://data.10jqka.com.cn/ipo/bond/
func BondZHCovInfoTHS() (dataframe.DataFrame, error) {
	url := "https://data.10jqka.com.cn/ipo/kzz/"
	headers := map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Safari/537.36",
	}

	resp, err := utils.GetWithHeaders(url, nil, headers)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	// 解析JSON响应
	list := gjson.Get(resp.String(), "list")
	if !list.Exists() || !list.IsArray() {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到可转债数据")
	}

	// 构建DataFrame记录
	var records []map[string]interface{}

	list.ForEach(func(_, item gjson.Result) bool {
		record := map[string]interface{}{
			"债券代码":   item.Get("bond_code").String(),
			"债券简称":   item.Get("bond_name").String(),
			"申购日期":   item.Get("sub_date").String(),
			"申购代码":   item.Get("sub_code").String(),
			"原股东配售码": item.Get("share_code").String(),
			"每股获配额":  utils.MustParseFloat(item.Get("quota").String()),
			"计划发行量":  utils.MustParseFloat(item.Get("plan_total").String()),
			"实际发行量":  utils.MustParseFloat(item.Get("issue_total").String()),
			"中签公布日":  item.Get("sign_date").String(),
			"中签号":    item.Get("number").String(),
			"上市日期":   item.Get("listing_date").String(),
			"正股代码":   item.Get("code").String(),
			"正股简称":   item.Get("name").String(),
			"转股价格":   utils.MustParseFloat(item.Get("price").String()),
			"到期时间":   item.Get("expire_date").String(),
			"中签率":    item.Get("success_rate").String(),
		}
		records = append(records, record)
		return true
	})

	if len(records) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到有效数据")
	}

	df := dataframe.LoadMaps(records)
	return df, nil
}
