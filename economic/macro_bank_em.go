package economic

import (
	"fmt"
	"strconv"
	"time"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/tidwall/gjson"
)

// getInterestRateData 获取央行利率决议报告数据（通用函数）
// 参数:
//   - attrID: 内置属性ID，用于区分不同央行
//   - name: 利率报告名称
//
// 返回:
//   - []InterestRateItem: 利率决议报告数据
//   - error: 错误信息
func getInterestRateData(attrID, name string) ([]InterestRateItem, error) {
	baseURL := "https://datacenter-api.jin10.com/reports/list_v2"
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)

	// 请求头
	headers := map[string]string{
		"Accept":          "*/*",
		"Accept-Language": "zh-CN,zh;q=0.9,en;q=0.8",
		"User-Agent":      jin10CommonHeaders["User-Agent"],
		"Origin":          "https://datacenter.jin10.com",
		"Referer":         "https://datacenter.jin10.com/",
		"x-app-id":        jin10CommonHeaders["x-app-id"],
		"x-version":       jin10CommonHeaders["x-version"],
	}

	// 初始参数
	params := map[string]string{
		"max_date": "",
		"category": "ec",
		"attr_id":  attrID,
		"_":        timestamp,
	}

	var allData []InterestRateItem

	// 分页获取数据
	for {
		resp, err := utils.GetWithHeaders(baseURL, params, headers)
		if err != nil {
			return nil, fmt.Errorf("获取%s数据失败: %w", name, err)
		}

		// 解析 JSON
		json := gjson.ParseBytes(resp.Body())

		// 检查是否有数据
		values := json.Get("data.values")
		if !values.Exists() || !values.IsArray() || len(values.Array()) == 0 {
			break
		}

		// 解析数据
		for _, item := range values.Array() {
			arr := item.Array()
			if len(arr) < 4 {
				continue
			}

			// 解析日期
			dateStr := arr[0].String()
			date, err := time.Parse("2006-01-02", dateStr)
			if err != nil {
				continue
			}

			// 解析数值（允许空值）
			current := 0.0
			if arr[1].Exists() && arr[1].String() != "" {
				current = arr[1].Float()
			}

			forecast := 0.0
			if arr[2].Exists() && arr[2].String() != "" {
				forecast = arr[2].Float()
			}

			previous := 0.0
			if arr[3].Exists() && arr[3].String() != "" {
				previous = arr[3].Float()
			}

			allData = append(allData, InterestRateItem{
				Date:     date,
				Current:  current,
				Forecast: forecast,
				Previous: previous,
			})
		}

		// 更新 max_date 用于分页
		lastDate := values.Array()[len(values.Array())-1].Array()[0].String()
		date, err := time.Parse("2006-01-02", lastDate)
		if err != nil {
			break
		}

		// 减去1天作为下一页的 max_date
		nextDate := date.AddDate(0, 0, -1).Format("2006-01-02")
		params["max_date"] = nextDate
	}

	return allData, nil
}

// MacroBankUsaInterestRate 美联储利率决议报告
// 数据区间从 1982-09-27 至今
// https://datacenter.jin10.com/reportType/dc_usa_interest_rate_decision
func MacroBankUsaInterestRate() ([]InterestRateItem, error) {
	return getInterestRateData("24", "美联储利率决议报告")
}

// MacroBankEuroInterestRate 欧洲央行利率决议报告
// 数据区间从 1999-01-01 至今
// https://datacenter.jin10.com/reportType/dc_interest_rate_decision
func MacroBankEuroInterestRate() ([]InterestRateItem, error) {
	return getInterestRateData("21", "欧洲央行利率决议报告")
}

// MacroBankNewzealandInterestRate 新西兰联储利率决议报告
// 数据区间从 1999-04-01 至今
// https://datacenter.jin10.com/reportType/dc_newzealand_interest_rate_decision
func MacroBankNewzealandInterestRate() ([]InterestRateItem, error) {
	return getInterestRateData("23", "新西兰联储利率决议报告")
}

// MacroBankChinaInterestRate 中国央行利率决议报告
// 数据区间从 1999-01-05 至今
// https://datacenter.jin10.com/reportType/dc_china_interest_rate_decision
func MacroBankChinaInterestRate() ([]InterestRateItem, error) {
	return getInterestRateData("91", "中国央行利率决议报告")
}

// MacroBankSwitzerlandInterestRate 瑞士央行利率决议报告
// 数据区间从 2008-03-13 至今
// https://datacenter.jin10.com/reportType/dc_switzerland_interest_rate_decision
func MacroBankSwitzerlandInterestRate() ([]InterestRateItem, error) {
	return getInterestRateData("25", "瑞士央行利率决议报告")
}

// MacroBankEnglishInterestRate 英国央行利率决议报告
// 数据区间从 1970-01-01 至今
// https://datacenter.jin10.com/reportType/dc_english_interest_rate_decision
func MacroBankEnglishInterestRate() ([]InterestRateItem, error) {
	return getInterestRateData("26", "英国央行利率决议报告")
}

// MacroBankAustraliaInterestRate 澳洲联储利率决议报告
// 数据区间从 1980-02-01 至今
// https://datacenter.jin10.com/reportType/dc_australia_interest_rate_decision
func MacroBankAustraliaInterestRate() ([]InterestRateItem, error) {
	return getInterestRateData("27", "澳洲联储利率决议报告")
}

// MacroBankJapanInterestRate 日本央行利率决议报告
// 数据区间从 2008-02-14 至今
// https://datacenter.jin10.com/reportType/dc_japan_interest_rate_decision
func MacroBankJapanInterestRate() ([]InterestRateItem, error) {
	return getInterestRateData("22", "日本央行利率决议报告")
}

// MacroBankRussiaInterestRate 俄罗斯央行利率决议报告
// 数据区间从 2003-06-01 至今
// https://datacenter.jin10.com/reportType/dc_russia_interest_rate_decision
func MacroBankRussiaInterestRate() ([]InterestRateItem, error) {
	return getInterestRateData("64", "俄罗斯央行利率决议报告")
}

// MacroBankIndiaInterestRate 印度央行利率决议报告
// 数据区间从 2000-08-01 至今
// https://datacenter.jin10.com/reportType/dc_india_interest_rate_decision
func MacroBankIndiaInterestRate() ([]InterestRateItem, error) {
	return getInterestRateData("68", "印度央行利率决议报告")
}

// MacroBankBrazilInterestRate 巴西央行利率决议报告
// 数据区间从 2008-02-01 至今
// https://datacenter.jin10.com/reportType/dc_brazil_interest_rate_decision
func MacroBankBrazilInterestRate() ([]InterestRateItem, error) {
	return getInterestRateData("55", "巴西央行利率决议报告")
}
