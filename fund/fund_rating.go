package fund

import (
	"fmt"
	"strings"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/PuerkitoBio/goquery"
)

// FundRatingAll 获取全部基金评级数据
func FundRatingAll() ([]map[string]interface{}, error) {
	url := "http://fund.eastmoney.com/data/fundrating.html"

	resp, err := utils.Get(url, nil)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(resp.String()))
	if err != nil {
		return nil, fmt.Errorf("解析HTML失败: %w", err)
	}

	records := make([]map[string]interface{}, 0)

	doc.Find("#dbtable tbody tr").Each(func(i int, s *goquery.Selection) {
		cells := s.Find("td")
		if cells.Length() < 8 {
			return
		}

		record := map[string]interface{}{
			"基金代码":   cells.Eq(1).Text(),
			"基金简称":   cells.Eq(2).Text(),
			"3年评级":   cells.Eq(3).Text(),
			"5年评级":   cells.Eq(4).Text(),
			"招商评级":   cells.Eq(5).Text(),
			"济安金信评级": cells.Eq(6).Text(),
			"上海证券评级": cells.Eq(7).Text(),
		}
		records = append(records, record)
	})

	return records, nil
}

// FundRatingSh 获取上海证券评级
// 参数: date 评级日期，如 "20230630"
func FundRatingSh(date string) ([]map[string]interface{}, error) {
	if date == "" {
		date = "20230630"
	}

	url := fmt.Sprintf("http://www.shzq.com/fund/rating/%s.html", date)

	resp, err := utils.Get(url, nil)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(resp.String()))
	if err != nil {
		return nil, fmt.Errorf("解析HTML失败: %w", err)
	}

	records := make([]map[string]interface{}, 0)

	doc.Find("table tr").Each(func(i int, s *goquery.Selection) {
		if i == 0 {
			return
		}

		cells := s.Find("td")
		if cells.Length() < 5 {
			return
		}

		record := map[string]interface{}{
			"基金代码": cells.Eq(0).Text(),
			"基金简称": cells.Eq(1).Text(),
			"评级":   cells.Eq(2).Text(),
			"评级日期": date,
		}
		records = append(records, record)
	})

	return records, nil
}

// FundRatingZs 获取招商证券评级
// 参数: date 评级日期，如 "20230331"
func FundRatingZs(date string) ([]map[string]interface{}, error) {
	if date == "" {
		date = "20230331"
	}

	url := fmt.Sprintf("http://www.cmschina.com/fund/rating/%s.html", date)

	resp, err := utils.Get(url, nil)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(resp.String()))
	if err != nil {
		return nil, fmt.Errorf("解析HTML失败: %w", err)
	}

	records := make([]map[string]interface{}, 0)

	doc.Find("table tr").Each(func(i int, s *goquery.Selection) {
		if i == 0 {
			return
		}

		cells := s.Find("td")
		if cells.Length() < 5 {
			return
		}

		record := map[string]interface{}{
			"基金代码": cells.Eq(0).Text(),
			"基金简称": cells.Eq(1).Text(),
			"评级":   cells.Eq(2).Text(),
			"评级日期": date,
		}
		records = append(records, record)
	})

	return records, nil
}

// FundRatingJa 获取济安金信评级
// 参数: date 评级日期，如 "20230331"
func FundRatingJa(date string) ([]map[string]interface{}, error) {
	if date == "" {
		date = "20230331"
	}

	url := fmt.Sprintf("http://www.jajxfund.com/rating/%s.html", date)

	resp, err := utils.Get(url, nil)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(resp.String()))
	if err != nil {
		return nil, fmt.Errorf("解析HTML失败: %w", err)
	}

	records := make([]map[string]interface{}, 0)

	doc.Find("table tr").Each(func(i int, s *goquery.Selection) {
		if i == 0 {
			return
		}

		cells := s.Find("td")
		if cells.Length() < 5 {
			return
		}

		record := map[string]interface{}{
			"基金代码": cells.Eq(0).Text(),
			"基金简称": cells.Eq(1).Text(),
			"评级":   cells.Eq(2).Text(),
			"评级日期": date,
		}
		records = append(records, record)
	})

	return records, nil
}
