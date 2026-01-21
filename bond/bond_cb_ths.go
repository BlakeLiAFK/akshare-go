package bond

import (
	"fmt"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/tidwall/gjson"
)

// BondZhCovInfoThsItem 同花顺-可转债行情项
type BondZhCovInfoThsItem struct {
	BondCode     string  `json:"bond_code"`     // 债券代码
	BondName     string  `json:"bond_name"`     // 债券简称
	SubDate      string  `json:"sub_date"`      // 申购日期
	SubCode      string  `json:"sub_code"`      // 申购代码
	ShareCode    string  `json:"share_code"`    // 原股东配售码
	Quota        float64 `json:"quota"`         // 每股获配额
	PlanTotal    float64 `json:"plan_total"`    // 计划发行量
	IssueTotal   float64 `json:"issue_total"`   // 实际发行量
	SignDate     string  `json:"sign_date"`     // 中签公布日
	SignNumber   string  `json:"sign_number"`   // 中签号
	ListingDate  string  `json:"listing_date"`  // 上市日期
	StockCode    string  `json:"stock_code"`    // 正股代码
	StockName    string  `json:"stock_name"`    // 正股简称
	ConvertPrice float64 `json:"convert_price"` // 转股价格
	ExpireDate   string  `json:"expire_date"`   // 到期时间
	SuccessRate  float64 `json:"success_rate"`  // 中签率
}

// BondZhCovInfoThs 同花顺-数据中心-可转债
// https://data.10jqka.com.cn/ipo/bond/
func BondZhCovInfoThs() ([]BondZhCovInfoThsItem, error) {
	url := "https://data.10jqka.com.cn/ipo/kzz/"

	resp, err := utils.GetWithHeaders(url, nil, map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Safari/537.36",
	})
	if err != nil {
		return nil, fmt.Errorf("获取同花顺可转债数据失败: %w", err)
	}

	result := gjson.ParseBytes(resp.Body())
	listArr := result.Get("list").Array()

	var items []BondZhCovInfoThsItem
	for _, item := range listArr {
		items = append(items, BondZhCovInfoThsItem{
			BondCode:     item.Get("bond_code").String(),
			BondName:     item.Get("bond_name").String(),
			SubDate:      item.Get("sub_date").String(),
			SubCode:      item.Get("sub_code").String(),
			ShareCode:    item.Get("share_code").String(),
			Quota:        item.Get("quota").Float(),
			PlanTotal:    item.Get("plan_total").Float(),
			IssueTotal:   item.Get("issue_total").Float(),
			SignDate:     item.Get("sign_date").String(),
			SignNumber:   item.Get("number").String(),
			ListingDate:  item.Get("listing_date").String(),
			StockCode:    item.Get("code").String(),
			StockName:    item.Get("name").String(),
			ConvertPrice: item.Get("price").Float(),
			ExpireDate:   item.Get("expire_date").String(),
			SuccessRate:  item.Get("success_rate").Float(),
		})
	}

	return items, nil
}
