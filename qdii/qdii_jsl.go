package qdii

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/PuerkitoBio/goquery"
)

const (
	// 集思录QDII数据URL
	jslQdiiURL = "https://www.jisilu.cn/data/qdii/"
)

// QdiiItem QDII基金数据结构
type QdiiItem struct {
	Code        string  `json:"code"`         // 代码
	Name        string  `json:"name"`         // 名称
	Price       float64 `json:"price"`        // 现价
	ChangePct   string  `json:"change_pct"`   // 涨幅
	Volume      string  `json:"volume"`       // 成交额(万)
	Premium     string  `json:"premium"`      // 溢价率
	EstNav      float64 `json:"est_nav"`      // 估算净值
	NavDate     string  `json:"nav_date"`     // 净值日期
	ApplyFee    string  `json:"apply_fee"`    // 申购费
	RedeemFee   string  `json:"redeem_fee"`   // 赎回费
	TrusteeFee  string  `json:"trustee_fee"`  // 托管费
	FundCompany string  `json:"fund_company"` // 基金公司
}

// QdiiEIndexJsl 获取集思录-T+0 QDII-欧美市场-欧美指数数据
//
// 目标地址: https://www.jisilu.cn/data/qdii/#qdiia
//
// 返回:
//   - []QdiiItem: QDII基金数据列表
//   - error: 错误信息
//
// 示例:
//
//	data, err := qdii.QdiiEIndexJsl()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, item := range data {
//	    fmt.Printf("%s %s %.3f %s\n", item.Code, item.Name, item.Price, item.ChangePct)
//	}
func QdiiEIndexJsl() ([]QdiiItem, error) {
	return fetchJslQdiiData("qdiie")
}

// QdiiECommJsl 获取集思录-T+0 QDII-欧美市场-欧美商品数据
//
// 目标地址: https://www.jisilu.cn/data/qdii/#qdiia
//
// 返回:
//   - []QdiiItem: QDII基金数据列表
//   - error: 错误信息
//
// 示例:
//
//	data, err := qdii.QdiiECommJsl()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, item := range data {
//	    fmt.Printf("%s %s %.3f %s\n", item.Code, item.Name, item.Price, item.ChangePct)
//	}
func QdiiECommJsl() ([]QdiiItem, error) {
	return fetchJslQdiiData("qdii_comm")
}

// QdiiAIndexJsl 获取集思录-T+0 QDII-亚洲市场-亚洲指数数据
//
// 目标地址: https://www.jisilu.cn/data/qdii/#qdiia
//
// 返回:
//   - []QdiiItem: QDII基金数据列表
//   - error: 错误信息
//
// 示例:
//
//	data, err := qdii.QdiiAIndexJsl()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, item := range data {
//	    fmt.Printf("%s %s %.3f %s\n", item.Code, item.Name, item.Price, item.ChangePct)
//	}
func QdiiAIndexJsl() ([]QdiiItem, error) {
	return fetchJslQdiiData("qdiia")
}

// fetchJslQdiiData 从集思录获取QDII数据
func fetchJslQdiiData(tabType string) ([]QdiiItem, error) {
	headers := map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
		"Referer":    "https://www.jisilu.cn/",
	}

	resp, err := utils.GetWithHeaders(jslQdiiURL, nil, headers)
	if err != nil {
		return nil, fmt.Errorf("获取QDII数据失败: %w", err)
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(resp.String()))
	if err != nil {
		return nil, fmt.Errorf("解析HTML失败: %w", err)
	}

	var items []QdiiItem

	// 根据tabType选择对应的表格
	tableSelector := fmt.Sprintf("#%s table tbody tr", tabType)
	doc.Find(tableSelector).Each(func(i int, s *goquery.Selection) {
		tds := s.Find("td")
		if tds.Length() < 10 {
			return
		}

		item := QdiiItem{
			Code:        strings.TrimSpace(tds.Eq(0).Text()),
			Name:        strings.TrimSpace(tds.Eq(1).Text()),
			Price:       utils.MustFloat64(strings.TrimSpace(tds.Eq(2).Text())),
			ChangePct:   strings.TrimSpace(tds.Eq(3).Text()),
			Volume:      strings.TrimSpace(tds.Eq(4).Text()),
			Premium:     strings.TrimSpace(tds.Eq(5).Text()),
			EstNav:      utils.MustFloat64(strings.TrimSpace(tds.Eq(6).Text())),
			NavDate:     strings.TrimSpace(tds.Eq(7).Text()),
			ApplyFee:    strings.TrimSpace(tds.Eq(8).Text()),
			RedeemFee:   strings.TrimSpace(tds.Eq(9).Text()),
			TrusteeFee:  strings.TrimSpace(tds.Eq(10).Text()),
			FundCompany: strings.TrimSpace(tds.Eq(11).Text()),
		}
		items = append(items, item)
	})

	// 如果表格解析失败，尝试使用API接口
	if len(items) == 0 {
		return fetchJslQdiiAPI(tabType)
	}

	return items, nil
}

// fetchJslQdiiAPI 通过API获取QDII数据
func fetchJslQdiiAPI(tabType string) ([]QdiiItem, error) {
	apiURL := "https://www.jisilu.cn/data/qdii/qdii_list/"

	params := map[string]string{
		"___jsl": "LST___t=1",
		"rp":     "25",
		"page":   "1",
	}

	headers := map[string]string{
		"User-Agent":       "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
		"Referer":          "https://www.jisilu.cn/data/qdii/",
		"X-Requested-With": "XMLHttpRequest",
	}

	resp, err := utils.GetWithHeaders(apiURL, params, headers)
	if err != nil {
		return nil, fmt.Errorf("获取QDII API数据失败: %w", err)
	}

	// 解析JSON响应
	var result struct {
		Rows []struct {
			Cell struct {
				FundID     string `json:"fund_id"`
				FundNm     string `json:"fund_nm"`
				Price      string `json:"price"`
				IncreaseRt string `json:"increase_rt"`
				Volume     string `json:"volume"`
				Premium    string `json:"discount_rt"`
				EstNav     string `json:"estimate_value"`
				NavDt      string `json:"nav_dt"`
				ApplyFee   string `json:"apply_fee"`
				RedeemFee  string `json:"redeem_fee"`
				TrusteeFee string `json:"manage_fee"`
				FundComp   string `json:"fund_company"`
			} `json:"cell"`
		} `json:"rows"`
	}

	if err := json.Unmarshal([]byte(resp.String()), &result); err != nil {
		// 返回空列表而不是错误，因为网站可能需要登录
		return []QdiiItem{}, nil
	}

	var items []QdiiItem
	for _, row := range result.Rows {
		item := QdiiItem{
			Code:        row.Cell.FundID,
			Name:        row.Cell.FundNm,
			Price:       utils.MustFloat64(row.Cell.Price),
			ChangePct:   row.Cell.IncreaseRt,
			Volume:      row.Cell.Volume,
			Premium:     row.Cell.Premium,
			EstNav:      utils.MustFloat64(row.Cell.EstNav),
			NavDate:     row.Cell.NavDt,
			ApplyFee:    row.Cell.ApplyFee,
			RedeemFee:   row.Cell.RedeemFee,
			TrusteeFee:  row.Cell.TrusteeFee,
			FundCompany: row.Cell.FundComp,
		}
		items = append(items, item)
	}

	return items, nil
}
