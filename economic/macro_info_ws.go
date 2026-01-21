package economic

import (
	"fmt"
	"time"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/go-gota/gota/dataframe"
	"github.com/tidwall/gjson"
)

// MacroInfoWS 华尔街见闻-日历-宏观
//
// 参数:
//   - date: 日期，格式为 YYYYMMDD，如 "20240514"
//
// 返回:
//   - dataframe.DataFrame: 宏观日历数据
//   - error: 错误信息
//
// 数据源: https://wallstreetcn.com/calendar
func MacroInfoWS(date string) (dataframe.DataFrame, error) {
	// 解析日期
	t, err := time.Parse("20060102", date)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("日期格式错误: %w", err)
	}

	// 计算开始和结束时间戳
	startTime := t.Unix()
	endTime := t.Add(24 * time.Hour).Unix()

	url := "https://api-one-wscn.awtmt.com/apiv1/finance/macrodatas"
	params := map[string]string{
		"start": fmt.Sprintf("%d", startTime),
		"end":   fmt.Sprintf("%d", endTime),
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	// 解析 JSON 响应
	items := gjson.Get(resp.String(), "data.items").Array()
	if len(items) == 0 {
		// 返回空 DataFrame
		columns := []string{"时间", "地区", "事件", "重要性", "今值", "预期", "前值", "链接"}
		records := [][]string{columns}
		return dataframe.LoadRecords(records), nil
	}

	// 定义列名
	columns := []string{"时间", "地区", "事件", "重要性", "今值", "预期", "前值", "链接"}

	// 构建记录
	var records [][]string
	records = append(records, columns)

	// 使用上海时区
	loc, _ := time.LoadLocation("Asia/Shanghai")

	for _, item := range items {
		// 解析时间戳
		publicDate := item.Get("public_date").Int()
		t := time.Unix(publicDate, 0).In(loc)
		timeStr := t.Format("2006-01-02 15:04:05")

		country := item.Get("country").String()
		title := item.Get("title").String()
		importance := item.Get("importance").String()
		actual := item.Get("actual").String()
		forecast := item.Get("forecast").String()
		previous := item.Get("previous").String()
		revised := item.Get("revised").String()
		uri := item.Get("uri").String()

		// 如果有修正值，用修正值替换前值
		if revised != "" && revised != "null" {
			previous = revised
		}

		record := []string{timeStr, country, title, importance, actual, forecast, previous, uri}
		records = append(records, record)
	}

	df := dataframe.LoadRecords(records)
	return df, nil
}
