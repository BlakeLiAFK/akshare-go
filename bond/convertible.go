package bond

import (
	"fmt"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/go-gota/gota/dataframe"
	"github.com/tidwall/gjson"
)

// BondCBIndexJSL 集思录可转债等权指数
//
// 参数:
//
//	无
//
// 返回:
//   - dataframe.DataFrame: 包含集思录可转债等权指数历史数据
//   - error: 错误信息
//
// 数据源: https://www.jisilu.cn/web/data/cb/index
func BondCBIndexJSL() (dataframe.DataFrame, error) {
	url := "https://www.jisilu.cn/webapi/cb/index_history/"

	resp, err := utils.Get(url, nil)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	// 解析JSON响应 - 数据是并行数组格式
	dataObj := gjson.Get(resp.String(), "data")
	if !dataObj.Exists() {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到数据")
	}

	// 获取日期数组和指数值数组
	priceDt := gjson.Get(resp.String(), "data.price_dt").Array()
	idxPrice := gjson.Get(resp.String(), "data.idx_price").Array()
	idxIncreaseRt := gjson.Get(resp.String(), "data.idx_increase_rt").Array()

	if len(priceDt) == 0 || len(idxPrice) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到有效数据")
	}

	// 构建DataFrame
	var records [][]string
	// 添加列头
	records = append(records, []string{"date", "idx_price", "idx_increase_rt"})

	// 合并并行数组
	for i := 0; i < len(priceDt); i++ {
		date := priceDt[i].String()
		price := idxPrice[i].String()
		increase := ""
		if i < len(idxIncreaseRt) {
			increase = idxIncreaseRt[i].String()
		}
		records = append(records, []string{date, price, increase})
	}

	df := dataframe.LoadRecords(records)
	return df, nil
}
