package article

import (
	"fmt"
	"strings"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/PuerkitoBio/goquery"
	"github.com/go-gota/gota/dataframe"
)

// ArticleFFCRR 获取 Fama-French Current Research Returns 多因子数据
//
// 返回:
//   - dataframe.DataFrame: Fama-French多因子模型数据
//   - error: 错误信息
//
// 数据源: https://mba.tuck.dartmouth.edu/pages/faculty/ken.french/data_library.html
func ArticleFFCRR() (dataframe.DataFrame, error) {
	url := "https://mba.tuck.dartmouth.edu/pages/faculty/ken.french/data_library.html"

	resp, err := utils.Get(url, nil)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	// 解析HTML
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(resp.String()))
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("解析HTML失败: %w", err)
	}

	// 提取第5个表格（索引为4）
	tables := doc.Find("table")
	if tables.Length() < 5 {
		return dataframe.DataFrame{}, fmt.Errorf("页面表格数量不足")
	}

	targetTable := tables.Eq(4)

	// 解析表格数据
	var records [][]string
	var headers []string

	// 提取表头
	targetTable.Find("tr").Each(func(i int, tr *goquery.Selection) {
		var row []string
		tr.Find("td, th").Each(func(j int, td *goquery.Selection) {
			text := strings.TrimSpace(td.Text())
			row = append(row, text)
		})

		if len(row) > 0 {
			if i == 0 {
				headers = row
			} else {
				records = append(records, row)
			}
		}
	})

	if len(headers) == 0 {
		headers = []string{"item"}
		for i := 1; i < len(records[0]); i++ {
			headers = append(headers, fmt.Sprintf("col_%d", i))
		}
	}

	// 构建DataFrame
	allRecords := append([][]string{headers}, records...)
	df := dataframe.LoadRecords(allRecords)

	return df, nil
}
