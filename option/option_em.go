package option

import (
	"fmt"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/tidwall/gjson"
)

// OptionCurrentEm 东方财富网-行情中心-期权市场
// https://quote.eastmoney.com/center/qqsc.html
func OptionCurrentEm() ([]OptionCurrentEmItem, error) {
	// 获取普通期权数据
	normalOptions, err := fetchNormalOptions()
	if err != nil {
		return nil, err
	}

	// 获取中金所期权数据
	cffexOptions, err := OptionCurrentCffexEm()
	if err != nil {
		return nil, err
	}

	// 合并结果
	result := append(normalOptions, cffexOptions...)

	// 重新编号
	for i := range result {
		result[i].Index = i + 1
	}

	return result, nil
}

// fetchNormalOptions 获取普通期权数据
func fetchNormalOptions() ([]OptionCurrentEmItem, error) {
	url := "https://23.push2.eastmoney.com/api/qt/clist/get"

	var allData []OptionCurrentEmItem
	page := 1
	pageSize := 100

	for {
		params := map[string]string{
			"pn":     fmt.Sprintf("%d", page),
			"pz":     fmt.Sprintf("%d", pageSize),
			"po":     "1",
			"np":     "1",
			"ut":     "bd1d9ddb04089700cf9c27f6f7426281",
			"fltt":   "2",
			"invt":   "2",
			"fid":    "f3",
			"fs":     "m:10,m:12,m:140,m:141,m:151,m:163,m:226",
			"fields": "f1,f2,f3,f4,f5,f6,f7,f8,f9,f10,f12,f13,f14,f15,f16,f17,f18,f20,f21,f23,f24,f25,f22,f28,f11,f62,f128,f136,f115,f152,f133,f108,f163,f161,f162",
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
			allData = append(allData, OptionCurrentEmItem{
				Index:        len(allData) + 1,
				Code:         value.Get("f12").String(),
				Name:         value.Get("f14").String(),
				LatestPrice:  value.Get("f2").Float(),
				Change:       value.Get("f4").Float(),
				ChangeRate:   value.Get("f3").Float(),
				Volume:       value.Get("f5").Int(),
				Amount:       value.Get("f6").Float(),
				Position:     value.Get("f133").Int(),
				StrikePrice:  value.Get("f161").Float(),
				RemainingDay: int(value.Get("f162").Int()),
				DayIncrease:  value.Get("f163").Int(),
				PreSettle:    value.Get("f28").Float(),
				Open:         value.Get("f17").Float(),
				MarketId:     value.Get("f13").String(),
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

// OptionCurrentCffexEm 东方财富网-中金所期权行情
func OptionCurrentCffexEm() ([]OptionCurrentEmItem, error) {
	url := "https://futsseapi.eastmoney.com/list/option/221"

	params := map[string]string{
		"orderBy":   "zdf",
		"sort":      "desc",
		"pageSize":  "20000",
		"pageIndex": "0",
		"token":     "58b2fa8f54638b60b87d69b31969089c",
		"field":     "dm,sc,name,p,zsjd,zde,zdf,f152,vol,cje,ccl,xqj,syr,rz,zjsj,o",
		"blockName": "callback",
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	text := resp.String()
	list := gjson.Get(text, "list")

	var result []OptionCurrentEmItem
	list.ForEach(func(key, value gjson.Result) bool {
		result = append(result, OptionCurrentEmItem{
			Index:        len(result) + 1,
			Code:         value.Get("dm").String(),
			Name:         value.Get("name").String(),
			LatestPrice:  value.Get("p").Float(),
			Change:       value.Get("zde").Float(),
			ChangeRate:   value.Get("zdf").Float(),
			Volume:       value.Get("vol").Int(),
			Amount:       value.Get("cje").Float(),
			Position:     value.Get("ccl").Int(),
			StrikePrice:  value.Get("xqj").Float(),
			RemainingDay: int(value.Get("syr").Int()),
			DayIncrease:  value.Get("rz").Int(),
			PreSettle:    value.Get("zjsj").Float(),
			Open:         value.Get("o").Float(),
			MarketId:     value.Get("sc").String(),
		})
		return true
	})

	return result, nil
}
