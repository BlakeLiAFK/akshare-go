package option

import (
	"fmt"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/tidwall/gjson"
)

// OptionPremiumAnalysisEm 东方财富网-数据中心-特色数据-期权折溢价
// https://data.eastmoney.com/other/premium.html
func OptionPremiumAnalysisEm() ([]OptionPremiumAnalysisEmItem, error) {
	url := "https://push2.eastmoney.com/api/qt/clist/get"

	var allData []OptionPremiumAnalysisEmItem
	page := 1
	pageSize := 100

	for {
		params := map[string]string{
			"fid":    "f250",
			"po":     "1",
			"pz":     fmt.Sprintf("%d", pageSize),
			"pn":     fmt.Sprintf("%d", page),
			"np":     "1",
			"fltt":   "2",
			"invt":   "2",
			"ut":     "b2884a393a59ad64002292a3e90d46a5",
			"fields": "f1,f2,f3,f12,f13,f14,f161,f250,f330,f331,f332,f333,f334,f335,f337,f301,f152",
			"fs":     "m:10",
		}

		resp, err := utils.Get(url, params)
		if err != nil {
			return nil, fmt.Errorf("请求失败: %w", err)
		}

		text := resp.String()
		data := gjson.Get(text, "data.diff")

		if !data.Exists() || len(data.Array()) == 0 {
			break
		}

		data.ForEach(func(key, value gjson.Result) bool {
			allData = append(allData, OptionPremiumAnalysisEmItem{
				OptionCode:       value.Get("f12").String(),
				OptionName:       value.Get("f14").String(),
				LatestPrice:      value.Get("f2").Float(),
				ChangeRate:       value.Get("f3").Float(),
				StrikePrice:      value.Get("f161").Float(),
				PremiumRate:      value.Get("f250").Float(),
				UnderlyingName:   value.Get("f335").String(),
				UnderlyingPrice:  value.Get("f337").Float(),
				UnderlyingChange: value.Get("f301").Float(),
				BreakEvenPrice:   value.Get("f334").Float(),
				ExpireDate:       value.Get("f330").String(),
			})
			return true
		})

		total := gjson.Get(text, "data.total").Int()
		if int64(len(allData)) >= total {
			break
		}
		page++
	}

	return allData, nil
}

// OptionRiskAnalysisEm 东方财富网-数据中心-特色数据-期权风险分析
// https://data.eastmoney.com/other/riskanal.html
func OptionRiskAnalysisEm() ([]OptionRiskAnalysisEmItem, error) {
	url := "https://push2.eastmoney.com/api/qt/clist/get"

	var allData []OptionRiskAnalysisEmItem
	page := 1
	pageSize := 100

	for {
		params := map[string]string{
			"fid":    "f12",
			"po":     "1",
			"pz":     fmt.Sprintf("%d", pageSize),
			"pn":     fmt.Sprintf("%d", page),
			"np":     "1",
			"fltt":   "2",
			"invt":   "2",
			"ut":     "b2884a393a59ad64002292a3e90d46a5",
			"fields": "f1,f2,f3,f12,f13,f14,f302,f303,f325,f326,f327,f329,f328,f301,f152,f154",
			"fs":     "m:10",
		}

		resp, err := utils.Get(url, params)
		if err != nil {
			return nil, fmt.Errorf("请求失败: %w", err)
		}

		text := resp.String()
		data := gjson.Get(text, "data.diff")

		if !data.Exists() || len(data.Array()) == 0 {
			break
		}

		data.ForEach(func(key, value gjson.Result) bool {
			allData = append(allData, OptionRiskAnalysisEmItem{
				OptionCode:   value.Get("f12").String(),
				OptionName:   value.Get("f14").String(),
				LatestPrice:  value.Get("f2").Float(),
				ChangeRate:   value.Get("f3").Float(),
				Leverage:     value.Get("f325").Float(),
				RealLeverage: value.Get("f326").Float(),
				Delta:        value.Get("f327").Float(),
				Gamma:        value.Get("f329").Float(),
				Vega:         value.Get("f328").Float(),
				Rho:          value.Get("f301").Float(),
				Theta:        value.Get("f154").Float(),
				ExpireDate:   value.Get("f302").String(),
			})
			return true
		})

		total := gjson.Get(text, "data.total").Int()
		if int64(len(allData)) >= total {
			break
		}
		page++
	}

	return allData, nil
}

// OptionValueAnalysisEm 东方财富网-数据中心-特色数据-期权价值分析
// https://data.eastmoney.com/other/valueAnal.html
func OptionValueAnalysisEm() ([]OptionValueAnalysisEmItem, error) {
	url := "https://push2.eastmoney.com/api/qt/clist/get"

	var allData []OptionValueAnalysisEmItem
	page := 1
	pageSize := 100

	for {
		params := map[string]string{
			"fid":    "f301",
			"po":     "1",
			"pz":     fmt.Sprintf("%d", pageSize),
			"pn":     fmt.Sprintf("%d", page),
			"np":     "1",
			"fltt":   "2",
			"invt":   "2",
			"ut":     "b2884a393a59ad64002292a3e90d46a5",
			"fields": "f1,f2,f3,f12,f13,f14,f298,f299,f249,f300,f330,f331,f332,f333,f334,f335,f336,f301,f152",
			"fs":     "m:10",
		}

		resp, err := utils.Get(url, params)
		if err != nil {
			return nil, fmt.Errorf("请求失败: %w", err)
		}

		text := resp.String()
		data := gjson.Get(text, "data.diff")

		if !data.Exists() || len(data.Array()) == 0 {
			break
		}

		data.ForEach(func(key, value gjson.Result) bool {
			allData = append(allData, OptionValueAnalysisEmItem{
				OptionCode:        value.Get("f12").String(),
				OptionName:        value.Get("f14").String(),
				LatestPrice:       value.Get("f2").Float(),
				TimeValue:         value.Get("f299").Float(),
				IntrinsicValue:    value.Get("f249").Float(),
				ImpliedVol:        value.Get("f298").Float(),
				TheoreticalPrice:  value.Get("f300").Float(),
				UnderlyingName:    value.Get("f335").String(),
				UnderlyingPrice:   value.Get("f337").Float(),
				UnderlyingYearVol: value.Get("f336").Float(),
				ExpireDate:        value.Get("f330").String(),
			})
			return true
		})

		total := gjson.Get(text, "data.total").Int()
		if int64(len(allData)) >= total {
			break
		}
		page++
	}

	return allData, nil
}
