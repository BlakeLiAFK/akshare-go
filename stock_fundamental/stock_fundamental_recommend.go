package stock_fundamental

import (
	"fmt"
	"strings"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/PuerkitoBio/goquery"
)

// StockInstituteRecommendItem 机构推荐项
type StockInstituteRecommendItem struct {
	Code         string `json:"code"`          // 股票代码
	Name         string `json:"name"`          // 股票名称
	RatingDate   string `json:"rating_date"`   // 评级日期
	Rating       string `json:"rating"`        // 评级
	RatingChange string `json:"rating_change"` // 评级变动
	OrgName      string `json:"org_name"`      // 机构名称
	Researcher   string `json:"researcher"`    // 研究员
	Industry     string `json:"industry"`      // 所属行业
}

// StockInstituteRecommendSummaryItem 股票综合评级项
type StockInstituteRecommendSummaryItem struct {
	Code          string  `json:"code"`           // 股票代码
	Name          string  `json:"name"`           // 股票名称
	OrgCount      int     `json:"org_count"`      // 机构数
	BuyCount      int     `json:"buy_count"`      // 买入
	AddCount      int     `json:"add_count"`      // 增持
	NeutralCount  int     `json:"neutral_count"`  // 中性
	ReduceCount   int     `json:"reduce_count"`   // 减持
	SellCount     int     `json:"sell_count"`     // 卖出
	SummaryRating float64 `json:"summary_rating"` // 综合评级
	LatestPrice   float64 `json:"latest_price"`   // 最新价
}

// StockInstituteRecommendTargetItem 目标涨幅排名项
type StockInstituteRecommendTargetItem struct {
	Code           string  `json:"code"`             // 股票代码
	Name           string  `json:"name"`             // 股票名称
	TargetIncrease float64 `json:"target_increase"`  // 平均目标涨幅
	LatestPrice    float64 `json:"latest_price"`     // 最新价
	AvgTargetPrice float64 `json:"avg_target_price"` // 平均目标价
	MaxTargetPrice float64 `json:"max_target_price"` // 最高目标价
	MinTargetPrice float64 `json:"min_target_price"` // 最低目标价
}

// StockInstituteRecommendAttentionItem 机构关注度项
type StockInstituteRecommendAttentionItem struct {
	Code              string  `json:"code"`                // 股票代码
	Name              string  `json:"name"`                // 股票名称
	Attention         int     `json:"attention"`           // 关注度
	BuyCount          int     `json:"buy_count"`           // 买入
	AddCount          int     `json:"add_count"`           // 增持
	NeutralCount      int     `json:"neutral_count"`       // 中性
	ReduceCount       int     `json:"reduce_count"`        // 减持
	SellCount         int     `json:"sell_count"`          // 卖出
	AvgTargetIncrease float64 `json:"avg_target_increase"` // 平均目标涨幅
	Industry          string  `json:"industry"`            // 所属行业
}

// StockInstituteRecommend 新浪财经-机构推荐池
// symbol: 推荐类型，可选 "最新投资评级", "上调评级股票", "下调评级股票", "股票综合评级", "首次评级股票", "目标涨幅排名", "机构关注度", "行业关注度", "投资评级选股"
func StockInstituteRecommend(symbol string) ([]map[string]string, error) {
	// URL映射
	indicatorMap := map[string]string{
		"最新投资评级": "http://stock.finance.sina.com.cn/stock/go.php/vIR_RatingNewest/index.phtml",
		"上调评级股票": "http://stock.finance.sina.com.cn/stock/go.php/vIR_RatingUp/index.phtml",
		"下调评级股票": "http://stock.finance.sina.com.cn/stock/go.php/vIR_RatingDown/index.phtml",
		"股票综合评级": "http://stock.finance.sina.com.cn/stock/go.php/vIR_StockRating/index.phtml",
		"首次评级股票": "http://stock.finance.sina.com.cn/stock/go.php/vIR_RatingFirst/index.phtml",
		"目标涨幅排名": "http://stock.finance.sina.com.cn/stock/go.php/vIR_TargetRise/index.phtml",
		"机构关注度":  "http://stock.finance.sina.com.cn/stock/go.php/vIR_StockHeat/index.phtml",
		"行业关注度":  "http://stock.finance.sina.com.cn/stock/go.php/vIR_IndustryHeat/index.phtml",
		"投资评级选股": "http://stock.finance.sina.com.cn/stock/go.php/vIR_SuggestSelect/index.phtml",
	}

	targetURL, ok := indicatorMap[symbol]
	if !ok {
		return nil, fmt.Errorf("不支持的推荐类型: %s", symbol)
	}

	// 获取数据
	params := map[string]string{
		"num": "10000",
		"p":   "1",
	}

	resp, err := utils.Get(targetURL, params)
	if err != nil {
		return nil, fmt.Errorf("获取数据失败: %w", err)
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(resp.String()))
	if err != nil {
		return nil, fmt.Errorf("解析HTML失败: %w", err)
	}

	// 解析表格
	var items []map[string]string

	// 获取表头
	var headers []string
	doc.Find("table thead tr th").Each(func(i int, th *goquery.Selection) {
		text := strings.TrimSpace(th.Text())
		// 清理排序标记
		text = strings.ReplaceAll(text, "↓", "")
		text = strings.ReplaceAll(text, "↑", "")
		headers = append(headers, text)
	})

	// 解析数据行
	doc.Find("table tbody tr").Each(func(i int, tr *goquery.Selection) {
		item := make(map[string]string)
		tr.Find("td").Each(func(j int, td *goquery.Selection) {
			if j < len(headers) {
				text := strings.TrimSpace(td.Text())
				// 股票代码补零
				if headers[j] == "股票代码" && len(text) < 6 {
					text = fmt.Sprintf("%06s", text)
				}
				item[headers[j]] = text
			}
		})
		if len(item) > 0 {
			items = append(items, item)
		}
	})

	return items, nil
}

// StockInstituteRecommendDetail 新浪财经-机构推荐池-股票评级记录
// symbol: 股票代码
func StockInstituteRecommendDetail(symbol string) ([]map[string]string, error) {
	url := fmt.Sprintf("http://stock.finance.sina.com.cn/stock/go.php/vIR_StockSearch/key/%s.phtml", symbol)
	params := map[string]string{
		"num": "5000",
		"p":   "1",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return nil, fmt.Errorf("获取股票评级记录失败: %w", err)
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(resp.String()))
	if err != nil {
		return nil, fmt.Errorf("解析HTML失败: %w", err)
	}

	var items []map[string]string

	// 获取表头
	var headers []string
	doc.Find("table thead tr th").Each(func(i int, th *goquery.Selection) {
		text := strings.TrimSpace(th.Text())
		text = strings.ReplaceAll(text, "↓", "")
		text = strings.ReplaceAll(text, "↑", "")
		headers = append(headers, text)
	})

	// 解析数据行
	doc.Find("table tbody tr").Each(func(i int, tr *goquery.Selection) {
		item := make(map[string]string)
		tr.Find("td").Each(func(j int, td *goquery.Selection) {
			if j < len(headers) {
				text := strings.TrimSpace(td.Text())
				if headers[j] == "股票代码" && len(text) < 6 {
					text = fmt.Sprintf("%06s", text)
				}
				item[headers[j]] = text
			}
		})
		if len(item) > 0 {
			items = append(items, item)
		}
	})

	return items, nil
}
