package fund

import (
	"fmt"
	"strings"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/PuerkitoBio/goquery"
	"github.com/tidwall/gjson"
)

// FundEtfCategorySina 获取新浪财经基金列表
// 参数: symbol 基金类型 "LOF基金"、"ETF基金"等
func FundEtfCategorySina(symbol string) ([]map[string]interface{}, error) {
	if symbol == "" {
		symbol = "LOF基金"
	}

	symbolMap := map[string]string{
		"LOF基金": "http://vip.stock.finance.sina.com.cn/quotes_service/api/json_v2.php/Market_Center.getHQNodeData?page=1&num=5000&sort=symbol&asc=1&node=lof_hq&_s_r_a=page",
		"ETF基金": "http://vip.stock.finance.sina.com.cn/quotes_service/api/json_v2.php/Market_Center.getHQNodeData?page=1&num=5000&sort=symbol&asc=1&node=etf_hq&_s_r_a=page",
		"分级基金":  "http://vip.stock.finance.sina.com.cn/quotes_service/api/json_v2.php/Market_Center.getHQNodeData?page=1&num=5000&sort=symbol&asc=1&node=fs_hq&_s_r_a=page",
	}

	url, ok := symbolMap[symbol]
	if !ok {
		return nil, fmt.Errorf("不支持的基金类型: %s", symbol)
	}

	resp, err := utils.Get(url, nil)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	result := gjson.Parse(resp.String())
	items := result.Array()

	records := make([]map[string]interface{}, 0, len(items))
	for _, item := range items {
		record := map[string]interface{}{
			"代码":  item.Get("symbol").String(),
			"名称":  item.Get("name").String(),
			"最新价": utils.MustFloat64(item.Get("trade").String()),
			"涨跌额": utils.MustFloat64(item.Get("pricechange").String()),
			"涨跌幅": utils.MustFloat64(item.Get("changepercent").String()),
			"买入":  utils.MustFloat64(item.Get("buy").String()),
			"卖出":  utils.MustFloat64(item.Get("sell").String()),
			"昨收":  utils.MustFloat64(item.Get("settlement").String()),
			"今开":  utils.MustFloat64(item.Get("open").String()),
			"最高":  utils.MustFloat64(item.Get("high").String()),
			"最低":  utils.MustFloat64(item.Get("low").String()),
			"成交量": utils.MustFloat64(item.Get("volume").String()),
			"成交额": utils.MustFloat64(item.Get("amount").String()),
		}
		records = append(records, record)
	}

	return records, nil
}

// FundEtfHistSina 获取ETF基金的日行情数据
// 参数: symbol 基金代码，如 "sh510050"
func FundEtfHistSina(symbol string) ([]map[string]interface{}, error) {
	if symbol == "" {
		symbol = "sh510050"
	}

	url := fmt.Sprintf("https://finance.sina.com.cn/realstock/company/%s/hisdata/klc_kl.js", symbol)

	resp, err := utils.Get(url, nil)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	// 解析HTML表格
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(resp.String()))
	if err != nil {
		return nil, fmt.Errorf("解析HTML失败: %w", err)
	}

	records := make([]map[string]interface{}, 0)

	// 查找表格数据
	doc.Find("table tr").Each(func(i int, s *goquery.Selection) {
		if i == 0 {
			return // 跳过表头
		}

		cells := s.Find("td")
		if cells.Length() < 7 {
			return
		}

		record := map[string]interface{}{
			"日期":  cells.Eq(0).Text(),
			"开盘价": utils.MustFloat64(cells.Eq(1).Text()),
			"最高价": utils.MustFloat64(cells.Eq(2).Text()),
			"收盘价": utils.MustFloat64(cells.Eq(3).Text()),
			"最低价": utils.MustFloat64(cells.Eq(4).Text()),
			"成交量": utils.MustFloat64(cells.Eq(5).Text()),
			"成交额": utils.MustFloat64(cells.Eq(6).Text()),
		}
		records = append(records, record)
	})

	return records, nil
}

// FundEtfDividendSina 获取ETF基金的累计分红信息
// 参数: symbol 基金代码，如 "sh510050"
func FundEtfDividendSina(symbol string) ([]map[string]interface{}, error) {
	if symbol == "" {
		symbol = "sh510050"
	}

	url := fmt.Sprintf("https://finance.sina.com.cn/fund/quotes/%s/bc.shtml", symbol)

	resp, err := utils.Get(url, nil)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	// 解析HTML表格
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(resp.String()))
	if err != nil {
		return nil, fmt.Errorf("解析HTML失败: %w", err)
	}

	records := make([]map[string]interface{}, 0)

	// 查找分红表格
	doc.Find("table.list_table tr").Each(func(i int, s *goquery.Selection) {
		if i == 0 {
			return // 跳过表头
		}

		cells := s.Find("td")
		if cells.Length() < 4 {
			return
		}

		record := map[string]interface{}{
			"权益登记日": cells.Eq(0).Text(),
			"除息日":   cells.Eq(1).Text(),
			"每份分红":  utils.MustFloat64(cells.Eq(2).Text()),
			"分红发放日": cells.Eq(3).Text(),
		}
		records = append(records, record)
	})

	return records, nil
}
