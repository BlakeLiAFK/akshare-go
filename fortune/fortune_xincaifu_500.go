package fortune

import (
	"fmt"
	"strings"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/tidwall/gjson"
)

const (
	// 新财富500URL
	xincaifuURL = "http://service.ikuyu.cn/XinCaiFu2/pcremoting/bdListAction.do"
)

// XincaifuRankItem 新财富500富豪榜数据结构
type XincaifuRankItem struct {
	Rank     string `json:"rank"`     // 排名
	Wealth   string `json:"wealth"`   // 财富
	Name     string `json:"name"`     // 姓名
	Company  string `json:"company"`  // 主要公司
	Industry string `json:"industry"` // 相关行业
	Location string `json:"location"` // 公司总部
	Gender   string `json:"gender"`   // 性别
	Age      string `json:"age"`      // 年龄
	Year     string `json:"year"`     // 年份
}

// XincaifuRank 获取新财富500人富豪榜
//
// 目标地址: http://www.xcf.cn/zhuanti/ztzz/hdzt1/500frb/index.html
//
// 参数:
//   - year: 年份，从2003年开始
//
// 返回:
//   - []XincaifuRankItem: 富豪榜数据列表
//   - error: 错误信息
//
// 示例:
//
//	data, err := fortune.XincaifuRank("2022")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, item := range data[:10] {
//	    fmt.Printf("%s %s %s %s\n", item.Rank, item.Name, item.Wealth, item.Company)
//	}
func XincaifuRank(year string) ([]XincaifuRankItem, error) {
	params := map[string]string{
		"method":   "getPage",
		"callback": "jsonpCallback",
		"sortBy":   "",
		"order":    "",
		"type":     "4",
		"keyword":  "",
		"pageSize": "1000",
		"year":     year,
		"pageNo":   "1",
		"from":     "jsonp",
	}

	headers := map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
	}

	resp, err := utils.GetWithHeaders(xincaifuURL, params, headers)
	if err != nil {
		return nil, fmt.Errorf("获取新财富500数据失败: %w", err)
	}

	text := resp.String()
	// 去除JSONP包装
	if strings.HasPrefix(text, "jsonpCallback(") {
		text = strings.TrimPrefix(text, "jsonpCallback(")
		text = strings.TrimSuffix(text, ")")
	}

	rows := gjson.Get(text, "data.rows")
	if !rows.Exists() {
		return nil, fmt.Errorf("解析数据失败")
	}

	var items []XincaifuRankItem
	rows.ForEach(func(_, value gjson.Result) bool {
		item := XincaifuRankItem{
			Rank:     value.Get("rank").String(),
			Wealth:   value.Get("assets").String(),
			Name:     value.Get("name").String(),
			Company:  value.Get("company").String(),
			Industry: value.Get("industry").String(),
			Location: value.Get("addr").String(),
			Gender:   value.Get("sex").String(),
			Age:      value.Get("age").String(),
			Year:     value.Get("year").String(),
		}
		items = append(items, item)
		return true
	})

	return items, nil
}
