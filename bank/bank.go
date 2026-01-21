package bank

import (
	"fmt"
	"strings"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/PuerkitoBio/goquery"
	"github.com/go-gota/gota/dataframe"
	"github.com/tidwall/gjson"
)

// itemMap 将中文项目名称映射到对应的itemId
var itemMap = map[string]string{
	"机关":   "4113",
	"本级":   "4114",
	"分局本级": "4115",
}

// BankFjcfTotalNum 获取银保监分局行政处罚的总记录数
//
// 参数:
//   - item: 项目类型，可选值："机关"、"本级"、"分局本级"
//
// 返回:
//   - int: 总记录数
//   - error: 错误信息
//
// 数据源: http://www.cbirc.gov.cn
func BankFjcfTotalNum(item string) (int, error) {
	itemID, ok := itemMap[item]
	if !ok {
		return 0, fmt.Errorf("无效的项目类型: %s，可选值: 机关、本级、分局本级", item)
	}

	url := "https://www.nfra.gov.cn/cbircweb/DocInfo/SelectDocByItemIdAndChild"
	params := map[string]string{
		"itemId":    itemID,
		"pageIndex": "1",
		"pageSize":  "18",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return 0, fmt.Errorf("请求失败: %w", err)
	}

	// 解析JSON响应获取总数
	total := gjson.Get(resp.String(), "data.total").Int()
	return int(total), nil
}

// BankFjcfTotalPage 计算银保监分局行政处罚的总页数
//
// 参数:
//   - item: 项目类型，可选值："机关"、"本级"、"分局本级"
//
// 返回:
//   - int: 总页数
//   - error: 错误信息
//
// 数据源: http://www.cbirc.gov.cn
func BankFjcfTotalPage(item string) (int, error) {
	totalNum, err := BankFjcfTotalNum(item)
	if err != nil {
		return 0, err
	}

	// 每页18条记录，计算总页数
	pageSize := 18
	totalPages := (totalNum + pageSize - 1) / pageSize
	return totalPages, nil
}

// BankFjcfPageUrl 获取银保监分局行政处罚的分页数据列表
//
// 参数:
//   - item: 项目类型，可选值："机关"、"本级"、"分局本级"
//
// 返回:
//   - dataframe.DataFrame: 包含docId、subtitle、publishDate的数据列表
//   - error: 错误信息
//
// 数据源: http://www.cbirc.gov.cn
func BankFjcfPageUrl(item string) (dataframe.DataFrame, error) {
	itemID, ok := itemMap[item]
	if !ok {
		return dataframe.DataFrame{}, fmt.Errorf("无效的项目类型: %s，可选值: 机关、本级、分局本级", item)
	}

	totalPages, err := BankFjcfTotalPage(item)
	if err != nil {
		return dataframe.DataFrame{}, err
	}

	url := "https://www.nfra.gov.cn/cbircweb/DocInfo/SelectDocByItemIdAndChild"

	var allRecords [][]string
	allRecords = append(allRecords, []string{"docId", "subtitle", "publishDate"})

	// 遍历所有页面
	for page := 1; page <= totalPages; page++ {
		params := map[string]string{
			"itemId":    itemID,
			"pageIndex": fmt.Sprintf("%d", page),
			"pageSize":  "18",
		}

		resp, err := utils.Get(url, params)
		if err != nil {
			return dataframe.DataFrame{}, fmt.Errorf("请求第%d页失败: %w", page, err)
		}

		// 解析JSON获取rows数组
		rows := gjson.Get(resp.String(), "data.rows").Array()
		for _, row := range rows {
			docID := row.Get("docId").String()
			subtitle := row.Get("subtitle").String()
			publishDate := row.Get("publishDate").String()

			allRecords = append(allRecords, []string{docID, subtitle, publishDate})
		}
	}

	// 构建DataFrame
	df := dataframe.LoadRecords(allRecords)
	return df, nil
}

// BankFjcfTableDetail 获取银保监分局行政处罚的详细表格数据
//
// 参数:
//   - url: 文档URL，格式如 "http://www.cbirc.gov.cn/cn/view/pages/ItemDetail.html?docId=xxxxx"
//
// 返回:
//   - dataframe.DataFrame: 包含详细处罚信息的DataFrame
//   - error: 错误信息
//
// 数据源: http://www.cbirc.gov.cn
func BankFjcfTableDetail(url string) (dataframe.DataFrame, error) {
	// 从URL中提取docId
	parts := strings.Split(url, "docId=")
	if len(parts) != 2 {
		return dataframe.DataFrame{}, fmt.Errorf("无效的URL格式: %s", url)
	}
	docID := parts[1]

	// 构建新的API URL
	apiURL := fmt.Sprintf("http://www.cbirc.gov.cn/cn/docInfoViewData/%s/%s.json", docID[:6], docID)

	resp, err := utils.Get(apiURL, nil)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	// 解析JSON获取HTML内容
	html := gjson.Get(resp.String(), "html").String()
	if html == "" {
		return dataframe.DataFrame{}, fmt.Errorf("未找到HTML内容")
	}

	// 使用goquery解析HTML
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("解析HTML失败: %w", err)
	}

	// 查找表格
	var allRecords [][]string

	doc.Find("table").Each(func(i int, table *goquery.Selection) {
		table.Find("tr").Each(func(j int, tr *goquery.Selection) {
			var row []string
			tr.Find("td").Each(func(k int, td *goquery.Selection) {
				text := strings.TrimSpace(td.Text())
				row = append(row, text)
			})
			if len(row) > 0 {
				allRecords = append(allRecords, row)
			}
		})
	})

	if len(allRecords) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("未找到表格数据")
	}

	// 标准化列名
	headers := []string{
		"行政处罚决定书文号",
		"姓名",
		"单位",
		"单位名称",
		"主要负责人姓名",
		"主要违法违规事实",
		"行政处罚依据",
		"行政处罚决定",
		"作出处罚决定的机关名称",
		"处罚决定日期",
		"处罚公开日期",
	}

	// 如果记录数大于1，说明有数据，添加标准化的列头
	if len(allRecords) > 1 {
		// 检查第一行是否已经是表头
		if len(allRecords[0]) < len(headers) {
			// 需要填充列
			for i := range allRecords {
				for len(allRecords[i]) < len(headers) {
					allRecords[i] = append(allRecords[i], "")
				}
			}
		}

		// 插入标准列头
		result := [][]string{headers}
		result = append(result, allRecords...)
		df := dataframe.LoadRecords(result)
		return df, nil
	}

	return dataframe.DataFrame{}, fmt.Errorf("数据不足")
}
