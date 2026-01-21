package economic

import (
	"fmt"
	"strconv"
	"time"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/tidwall/gjson"
)

// macroUsaBaseFunc 美国宏观数据基础函数（金十数据中心）
// 参数:
//   - symbol: 指标名称
//   - attrID: 属性ID，用于区分不同指标
//
// 返回:
//   - []MacroEconomicItem: 宏观经济数据
//   - error: 错误信息
func macroUsaBaseFunc(symbol, attrID string) ([]MacroEconomicItem, error) {
	baseURL := "https://datacenter-api.jin10.com/reports/list_v2"
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)

	// 初始参数
	params := map[string]string{
		"max_date": "",
		"category": "ec",
		"attr_id":  attrID,
		"_":        timestamp,
	}

	var allData []MacroEconomicItem

	// 分页获取数据
	for {
		resp, err := utils.GetWithHeaders(baseURL, params, jin10CommonHeaders)
		if err != nil {
			return nil, fmt.Errorf("获取%s数据失败: %w", symbol, err)
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

			allData = append(allData, MacroEconomicItem{
				Commodity: symbol,
				Date:      date,
				Current:   current,
				Forecast:  forecast,
				Previous:  previous,
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

// MacroUsaGdpMonthly 美国国内生产总值(GDP)报告
// 数据区间从 2008-02-28 至今
// https://datacenter.jin10.com/reportType/dc_usa_gdp
func MacroUsaGdpMonthly() ([]MacroEconomicItem, error) {
	return macroUsaBaseFunc("美国国内生产总值(GDP)", "53")
}

// MacroUsaCpiMonthly 美国CPI月率报告
// 数据区间从 1970-01-01 至今
// https://datacenter.jin10.com/reportType/dc_usa_cpi
func MacroUsaCpiMonthly() ([]MacroEconomicItem, error) {
	return macroUsaBaseFunc("美国CPI月率", "9")
}

// MacroUsaCoreCpiMonthly 美国核心CPI月率报告
// 数据区间从 1970-01-01 至今
// https://datacenter.jin10.com/reportType/dc_usa_core_cpi
func MacroUsaCoreCpiMonthly() ([]MacroEconomicItem, error) {
	return macroUsaBaseFunc("美国核心CPI月率", "6")
}

// MacroUsaPersonalSpending 美国个人支出月率报告
// https://datacenter.jin10.com/reportType/dc_usa_personal_spending
func MacroUsaPersonalSpending() ([]MacroEconomicItem, error) {
	return macroUsaBaseFunc("美国个人支出月率", "35")
}

// MacroUsaRetailSales 美国零售销售月率报告
// https://datacenter.jin10.com/reportType/dc_usa_retail_sales
func MacroUsaRetailSales() ([]MacroEconomicItem, error) {
	return macroUsaBaseFunc("美国零售销售月率", "39")
}

// MacroUsaImportPrice 美国进口物价指数报告
// https://datacenter.jin10.com/reportType/dc_usa_import_price
func MacroUsaImportPrice() ([]MacroEconomicItem, error) {
	return macroUsaBaseFunc("美国进口物价指数", "18")
}

// MacroUsaExportPrice 美国出口价格指数报告
// https://datacenter.jin10.com/reportType/dc_usa_export_price
func MacroUsaExportPrice() ([]MacroEconomicItem, error) {
	return macroUsaBaseFunc("美国出口价格指数", "79")
}

// MacroUsaLmci 美联储劳动力市场状况指数报告
// https://datacenter.jin10.com/reportType/dc_usa_lmci
func MacroUsaLmci() ([]MacroEconomicItem, error) {
	return macroUsaBaseFunc("美联储劳动力市场状况指数", "93")
}

// MacroUsaUnemploymentRate 美国失业率报告
// https://datacenter.jin10.com/reportType/dc_usa_unemployment_rate
func MacroUsaUnemploymentRate() ([]MacroEconomicItem, error) {
	return macroUsaBaseFunc("美国失业率", "47")
}

// MacroUsaJobCuts 美国挑战者企业裁员人数报告
// https://datacenter.jin10.com/reportType/dc_usa_job_cuts
func MacroUsaJobCuts() ([]MacroEconomicItem, error) {
	return macroUsaBaseFunc("美国挑战者企业裁员人数", "78")
}

// MacroUsaNonFarm 美国非农就业人数报告
// https://datacenter.jin10.com/reportType/dc_usa_non_farm
func MacroUsaNonFarm() ([]MacroEconomicItem, error) {
	return macroUsaBaseFunc("美国非农就业人数", "33")
}

// MacroUsaAdpEmployment 美国ADP就业人数报告
// https://datacenter.jin10.com/reportType/dc_usa_adp_employment
func MacroUsaAdpEmployment() ([]MacroEconomicItem, error) {
	return macroUsaBaseFunc("美国ADP就业人数", "1")
}

// MacroUsaCorePcePrice 美国核心PCE物价指数年率报告
// https://datacenter.jin10.com/reportType/dc_usa_core_pce_price
func MacroUsaCorePcePrice() ([]MacroEconomicItem, error) {
	return macroUsaBaseFunc("美国核心PCE物价指数年率", "80")
}

// MacroUsaTradeBalance 美国贸易帐报告
// https://datacenter.jin10.com/reportType/dc_usa_trade_balance
func MacroUsaTradeBalance() ([]MacroEconomicItem, error) {
	return macroUsaBaseFunc("美国贸易帐报告", "42")
}

// MacroUsaCurrentAccount 美国经常账报告
// https://datacenter.jin10.com/reportType/dc_usa_current_account
func MacroUsaCurrentAccount() ([]MacroEconomicItem, error) {
	return macroUsaBaseFunc("美国经常账报告", "12")
}

// MacroUsaPpi 美国生产者物价指数报告
// https://datacenter.jin10.com/reportType/dc_usa_ppi
func MacroUsaPpi() ([]MacroEconomicItem, error) {
	return macroUsaBaseFunc("美国生产者物价指数", "37")
}

// MacroUsaCorePpi 美国核心生产者物价指数报告
// https://datacenter.jin10.com/reportType/dc_usa_core_ppi
func MacroUsaCorePpi() ([]MacroEconomicItem, error) {
	return macroUsaBaseFunc("美国核心生产者物价指数", "7")
}

// MacroUsaApiCrudeStock 美国API原油库存报告
// https://datacenter.jin10.com/reportType/dc_usa_api_crude_stock
func MacroUsaApiCrudeStock() ([]MacroEconomicItem, error) {
	return macroUsaBaseFunc("美国API原油库存", "69")
}

// MacroUsaPmi 美国Markit制造业PMI报告
// https://datacenter.jin10.com/reportType/dc_usa_pmi
func MacroUsaPmi() ([]MacroEconomicItem, error) {
	return macroUsaBaseFunc("美国Markit制造业PMI报告", "74")
}

// MacroUsaIsmPmi 美国ISM制造业PMI报告
// https://datacenter.jin10.com/reportType/dc_usa_ism_pmi
func MacroUsaIsmPmi() ([]MacroEconomicItem, error) {
	return macroUsaBaseFunc("美国ISM制造业PMI报告", "28")
}

// MacroUsaIndustrialProduction 美国工业产出月率报告
// https://datacenter.jin10.com/reportType/dc_usa_industrial_production
func MacroUsaIndustrialProduction() ([]MacroEconomicItem, error) {
	return macroUsaBaseFunc("美国工业产出月率报告", "20")
}

// MacroUsaDurableGoodsOrders 美国耐用品订单月率报告
// https://datacenter.jin10.com/reportType/dc_usa_durable_goods_orders
func MacroUsaDurableGoodsOrders() ([]MacroEconomicItem, error) {
	return macroUsaBaseFunc("美国耐用品订单月率报告", "13")
}

// MacroUsaFactoryOrders 美国工厂订单月率报告
// https://datacenter.jin10.com/reportType/dc_usa_factory_orders
func MacroUsaFactoryOrders() ([]MacroEconomicItem, error) {
	return macroUsaBaseFunc("美国工厂订单月率报告", "16")
}

// MacroUsaServicesPmi 美国Markit服务业PMI初值报告
// https://datacenter.jin10.com/reportType/dc_usa_services_pmi
func MacroUsaServicesPmi() ([]MacroEconomicItem, error) {
	return macroUsaBaseFunc("美国Markit服务业PMI初值报告", "89")
}

// MacroUsaBusinessInventories 美国商业库存月率报告
// https://datacenter.jin10.com/reportType/dc_usa_business_inventories
func MacroUsaBusinessInventories() ([]MacroEconomicItem, error) {
	return macroUsaBaseFunc("美国商业库存月率报告", "4")
}

// MacroUsaIsmNonPmi 美国ISM非制造业PMI报告
// https://datacenter.jin10.com/reportType/dc_usa_ism_non_pmi
func MacroUsaIsmNonPmi() ([]MacroEconomicItem, error) {
	return macroUsaBaseFunc("美国ISM非制造业PMI报告", "29")
}

// MacroUsaNahbHouseMarketIndex 美国NAHB房产市场指数报告
// https://datacenter.jin10.com/reportType/dc_usa_nahb_house_market_index
func MacroUsaNahbHouseMarketIndex() ([]MacroEconomicItem, error) {
	return macroUsaBaseFunc("美国NAHB房产市场指数报告", "31")
}

// MacroUsaHouseStarts 美国新屋开工总数年化报告
// https://datacenter.jin10.com/reportType/dc_usa_house_starts
func MacroUsaHouseStarts() ([]MacroEconomicItem, error) {
	return macroUsaBaseFunc("美国新屋开工总数年化报告", "17")
}

// MacroUsaNewHomeSales 美国新屋销售总数年化报告
// https://datacenter.jin10.com/reportType/dc_usa_new_home_sales
func MacroUsaNewHomeSales() ([]MacroEconomicItem, error) {
	return macroUsaBaseFunc("美国新屋销售总数年化报告", "32")
}

// MacroUsaBuildingPermits 美国营建许可总数报告
// https://datacenter.jin10.com/reportType/dc_usa_building_permits
func MacroUsaBuildingPermits() ([]MacroEconomicItem, error) {
	return macroUsaBaseFunc("美国营建许可总数报告", "3")
}

// MacroUsaExistHomeSales 美国成屋销售总数年化报告
// https://datacenter.jin10.com/reportType/dc_usa_exist_home_sales
func MacroUsaExistHomeSales() ([]MacroEconomicItem, error) {
	return macroUsaBaseFunc("美国成屋销售总数年化报告", "15")
}

// MacroUsaHousePriceIndex 美国FHFA房价指数月率报告
// https://datacenter.jin10.com/reportType/dc_usa_house_price_index
func MacroUsaHousePriceIndex() ([]MacroEconomicItem, error) {
	return macroUsaBaseFunc("美国FHFA房价指数月率报告", "51")
}

// MacroUsaCbConsumerConfidence 美国谘商会消费者信心指数报告
// https://datacenter.jin10.com/reportType/dc_usa_cb_consumer_confidence
func MacroUsaCbConsumerConfidence() ([]MacroEconomicItem, error) {
	return macroUsaBaseFunc("美国谘商会消费者信心指数", "5")
}

// MacroUsaEiaCrudeRate 美国EIA原油库存报告
// https://datacenter.jin10.com/reportType/dc_usa_eia_crude_rate
func MacroUsaEiaCrudeRate() ([]MacroEconomicItem, error) {
	return macroUsaBaseFunc("美国EIA原油库存", "10")
}

// MacroUsaInitialJobless 美国初请失业金人数报告
// https://datacenter.jin10.com/reportType/dc_usa_initial_jobless
func MacroUsaInitialJobless() ([]MacroEconomicItem, error) {
	return macroUsaBaseFunc("美国初请失业金人数", "44")
}

// MacroUsaRealConsumerSpending 美国实际个人消费支出季率初值报告
// https://datacenter.jin10.com/reportType/dc_usa_real_consumer_spending
func MacroUsaRealConsumerSpending() ([]MacroEconomicItem, error) {
	return macroUsaBaseFunc("美国实际个人消费支出季率初值", "81")
}

// MacroUsaSpcs20 美国S&P/CS20座大城市房价指数年率报告
// https://datacenter.jin10.com/reportType/dc_usa_spcs20
func MacroUsaSpcs20() ([]MacroEconomicItem, error) {
	return macroUsaBaseFunc("美国S&P/CS20座大城市房价指数年率", "52")
}

// MacroUsaPendingHomeSales 美国成屋签约销售指数月率报告
// https://datacenter.jin10.com/reportType/dc_usa_pending_home_sales
func MacroUsaPendingHomeSales() ([]MacroEconomicItem, error) {
	return macroUsaBaseFunc("美国成屋签约销售指数月率", "34")
}

// MacroUsaNfibSmallBusiness 美国NFIB小型企业信心指数报告
// https://datacenter.jin10.com/reportType/dc_usa_nfib_small_business
func MacroUsaNfibSmallBusiness() ([]MacroEconomicItem, error) {
	return macroUsaBaseFunc("美国NFIB小型企业信心指数", "63")
}

// MacroUsaMichiganConsumerSentiment 美国密歇根大学消费者信心指数初值报告
// https://datacenter.jin10.com/reportType/dc_usa_michigan_consumer_sentiment
func MacroUsaMichiganConsumerSentiment() ([]MacroEconomicItem, error) {
	return macroUsaBaseFunc("美国密歇根大学消费者信心指数初值", "50")
}

// MacroUsaRigCount 贝克休斯钻井平台报告
// 数据区间从 2008-03-17 至今
// https://datacenter.jin10.com/reportType/dc_rig_count_summary
func MacroUsaRigCount() ([]RigCountItem, error) {
	baseURL := "https://cdn.jin10.com/data_center/reports/baker.json"
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)

	params := map[string]string{
		"_": timestamp,
	}

	resp, err := utils.GetWithHeaders(baseURL, params, jin10CommonHeaders)
	if err != nil {
		return nil, fmt.Errorf("获取贝克休斯钻井平台报告失败: %w", err)
	}

	// 解析 JSON
	json := gjson.ParseBytes(resp.Body())
	values := json.Get("values")
	if !values.Exists() || !values.IsObject() {
		return nil, fmt.Errorf("贝克休斯钻井平台报告数据格式错误")
	}

	var data []RigCountItem

	// 遍历所有日期
	values.ForEach(func(date, value gjson.Result) bool {
		totalRig := value.Get("钻井总数")
		oilRig := value.Get("美国石油钻井")
		mixedRig := value.Get("混合钻井")
		gasRig := value.Get("美国天然气钻井")

		item := RigCountItem{
			Date: date.String(),
		}

		// 解析钻井总数 [钻井数, 变化]
		if totalRig.Exists() && totalRig.IsArray() {
			arr := totalRig.Array()
			if len(arr) >= 2 {
				item.TotalRigCount = arr[0].Float()
				item.TotalRigChange = arr[1].Float()
			}
		}

		// 解析美国石油钻井 [钻井数, 变化]
		if oilRig.Exists() && oilRig.IsArray() {
			arr := oilRig.Array()
			if len(arr) >= 2 {
				item.USAOilRigCount = arr[0].Float()
				item.USAOilRigChange = arr[1].Float()
			}
		}

		// 解析混合钻井 [钻井数, 变化]
		if mixedRig.Exists() && mixedRig.IsArray() {
			arr := mixedRig.Array()
			if len(arr) >= 2 {
				item.MixedRigCount = arr[0].Float()
				item.MixedRigChange = arr[1].Float()
			}
		}

		// 解析美国天然气钻井 [钻井数, 变化]
		if gasRig.Exists() && gasRig.IsArray() {
			arr := gasRig.Array()
			if len(arr) >= 2 {
				item.USANaturalGasRigCount = arr[0].Float()
				item.USANaturalGasRigChange = arr[1].Float()
			}
		}

		data = append(data, item)
		return true // continue iteration
	})

	// 按日期排序（这里简单返回，实际Python代码有排序）
	return data, nil
}

// MacroUsaPhs 美国未决房屋销售月率报告（东方财富）
// https://data.eastmoney.com/cjsj/foreign_0_5.html
func MacroUsaPhs() ([]EastMoneyEconomicItem, error) {
	baseURL := "https://datacenter-web.eastmoney.com/api/data/v1/get"

	params := map[string]string{
		"reportName":  "RPT_ECONOMICVALUE_USA",
		"columns":     "ALL",
		"filter":      `(INDICATOR_ID="EMG00342249")`,
		"pageNumber":  "1",
		"pageSize":    "2000",
		"sortColumns": "REPORT_DATE",
		"sortTypes":   "-1",
		"source":      "WEB",
		"client":      "WEB",
		"p":           "1",
		"pageNo":      "1",
		"pageNum":     "1",
	}

	resp, err := utils.Get(baseURL, params)
	if err != nil {
		return nil, fmt.Errorf("获取美国未决房屋销售月率数据失败: %w", err)
	}

	// 解析 JSON
	json := gjson.ParseBytes(resp.Body())
	dataArray := json.Get("result.data")
	if !dataArray.Exists() || !dataArray.IsArray() {
		return nil, fmt.Errorf("美国未决房屋销售月率数据格式错误")
	}

	var result []EastMoneyEconomicItem

	for _, item := range dataArray.Array() {
		// 字段顺序: REPORT_DATE, PUBLISH_DATE, VALUE, PRE_VALUE
		reportDate := item.Get("REPORT_DATE").String()
		publishDate := item.Get("PUBLISH_DATE").String()
		current := item.Get("VALUE").Float()
		previous := item.Get("PRE_VALUE").Float()

		result = append(result, EastMoneyEconomicItem{
			Time:        reportDate,
			Previous:    previous,
			Current:     current,
			PublishDate: publishDate,
		})
	}

	return result, nil
}

// MacroUsaCpiYoy 美国CPI年率报告（东方财富）
// 数据区间从 2008 至今
// https://data.eastmoney.com/cjsj/foreign_0_12.html
func MacroUsaCpiYoy() ([]EastMoneyEconomicItem, error) {
	baseURL := "https://datacenter-web.eastmoney.com/api/data/v1/get"

	params := map[string]string{
		"reportName":  "RPT_ECONOMICVALUE_USA",
		"columns":     "ALL",
		"filter":      `(INDICATOR_ID="EMG00000733")`,
		"sortColumns": "REPORT_DATE",
		"sortTypes":   "-1",
		"source":      "WEB",
		"client":      "WEB",
	}

	resp, err := utils.Get(baseURL, params)
	if err != nil {
		return nil, fmt.Errorf("获取美国CPI年率数据失败: %w", err)
	}

	// 解析 JSON
	json := gjson.ParseBytes(resp.Body())
	dataArray := json.Get("result.data")
	if !dataArray.Exists() || !dataArray.IsArray() {
		return nil, fmt.Errorf("美国CPI年率数据格式错误")
	}

	var result []EastMoneyEconomicItem

	for _, item := range dataArray.Array() {
		reportDate := item.Get("REPORT_DATE").String()
		publishDate := item.Get("PUBLISH_DATE").String()
		current := item.Get("VALUE").Float()
		previous := item.Get("PRE_VALUE").Float()

		result = append(result, EastMoneyEconomicItem{
			Time:        reportDate,
			Previous:    previous,
			Current:     current,
			PublishDate: publishDate,
		})
	}

	return result, nil
}

// MacroUsaCrudeInner 美国原油产量报告
// 数据区间从 1983-01-07 至今
// https://datacenter.jin10.com/reportType/dc_eia_crude_oil_produce
func MacroUsaCrudeInner() ([]CrudeOilProductionItem, error) {
	baseURL := "https://cdn.jin10.com/data_center/reports/usa_oil.json"
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)

	params := map[string]string{
		"_": timestamp,
	}

	resp, err := utils.GetWithHeaders(baseURL, params, jin10CommonHeaders)
	if err != nil {
		return nil, fmt.Errorf("获取美国原油产量报告失败: %w", err)
	}

	// 解析 JSON
	json := gjson.ParseBytes(resp.Body())
	values := json.Get("values")
	if !values.Exists() || !values.IsObject() {
		return nil, fmt.Errorf("美国原油产量报告数据格式错误")
	}

	var data []CrudeOilProductionItem

	// 遍历所有日期
	values.ForEach(func(date, value gjson.Result) bool {
		totalOil := value.Get("美国国内原油总量")
		mainland48 := value.Get("美国本土48州原油产量")
		alaska := value.Get("美国阿拉斯加州原油产量")

		item := CrudeOilProductionItem{
			Date: date.String(),
		}

		// 解析美国国内原油总量 [产量, 变化]
		if totalOil.Exists() && totalOil.IsArray() {
			arr := totalOil.Array()
			if len(arr) >= 2 {
				item.TotalProduction = arr[0].Float()
				item.TotalChange = arr[1].Float()
			}
		}

		// 解析美国本土48州原油产量 [产量, 变化]
		if mainland48.Exists() && mainland48.IsArray() {
			arr := mainland48.Array()
			if len(arr) >= 2 {
				item.Mainland48Production = arr[0].Float()
				item.Mainland48Change = arr[1].Float()
			}
		}

		// 解析美国阿拉斯加州原油产量 [产量, 变化]
		if alaska.Exists() && alaska.IsArray() {
			arr := alaska.Array()
			if len(arr) >= 2 {
				item.AlaskaProduction = arr[0].Float()
				item.AlaskaChange = arr[1].Float()
			}
		}

		data = append(data, item)
		return true // continue iteration
	})

	return data, nil
}

// cftcBaseFunc CFTC持仓报告基础函数
// 参数:
//   - jsonURL: JSON数据URL
//   - name: 报告名称
//
// 返回:
//   - []CFTCItem: CFTC持仓数据
//   - error: 错误信息
func cftcBaseFunc(jsonURL, name string) ([]CFTCItem, error) {
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)

	params := map[string]string{
		"_": timestamp,
	}

	resp, err := utils.GetWithHeaders(jsonURL, params, jin10CommonHeaders)
	if err != nil {
		return nil, fmt.Errorf("获取%s失败: %w", name, err)
	}

	// 解析 JSON
	json := gjson.ParseBytes(resp.Body())
	values := json.Get("values")
	keys := json.Get("keys")

	if !values.Exists() || !values.IsObject() {
		return nil, fmt.Errorf("%s数据格式错误", name)
	}

	if !keys.Exists() || !keys.IsArray() {
		return nil, fmt.Errorf("%s keys格式错误", name)
	}

	// 获取keys名称
	keyNames := make([]string, 0)
	for _, key := range keys.Array() {
		keyNames = append(keyNames, key.Get("name").String())
	}

	var data []CFTCItem

	// 遍历所有日期
	values.ForEach(func(date, value gjson.Result) bool {
		item := CFTCItem{
			Date: date.String(),
			Data: make(map[string]float64),
		}

		// 遍历所有品种
		value.ForEach(func(variety, varValue gjson.Result) bool {
			varietyName := variety.String()

			// 品种值是数组 [value1, value2, value3]
			if varValue.IsArray() {
				arr := varValue.Array()
				for i, val := range arr {
					if i < len(keyNames) {
						// 构造键名: "品种-指标"
						key := varietyName + "-" + keyNames[i]
						item.Data[key] = val.Float()
					}
				}
			}

			return true // continue iteration
		})

		data = append(data, item)
		return true // continue iteration
	})

	return data, nil
}

// MacroUsaCftcNcHolding 美国商品期货交易委员会CFTC外汇类非商业持仓报告
// 数据区间从 1983-01-07 至今
// https://datacenter.jin10.com/reportType/dc_cftc_nc_report
func MacroUsaCftcNcHolding() ([]CFTCItem, error) {
	return cftcBaseFunc(
		"https://cdn.jin10.com/data_center/reports/cftc_4.json",
		"美国商品期货交易委员会CFTC外汇类非商业持仓报告",
	)
}

// MacroUsaCftcCHolding 美国商品期货交易委员会CFTC商品类非商业持仓报告
// 数据区间从 1983-01-07 至今
// https://datacenter.jin10.com/reportType/dc_cftc_c_report
func MacroUsaCftcCHolding() ([]CFTCItem, error) {
	return cftcBaseFunc(
		"https://cdn.jin10.com/data_center/reports/cftc_2.json",
		"美国商品期货交易委员会CFTC商品类非商业持仓报告",
	)
}

// MacroUsaCftcMerchantCurrencyHolding 美国商品期货交易委员会CFTC外汇类商业持仓报告
// 数据区间从 1986-01-15 至今
// https://datacenter.jin10.com/reportType/dc_cftc_merchant_currency
func MacroUsaCftcMerchantCurrencyHolding() ([]CFTCItem, error) {
	return cftcBaseFunc(
		"https://cdn.jin10.com/data_center/reports/cftc_3.json",
		"美国商品期货交易委员会CFTC外汇类商业持仓报告",
	)
}

// MacroUsaCftcMerchantGoodsHolding 美国商品期货交易委员会CFTC商品类商业持仓报告
// 数据区间从 1986-01-15 至今
// https://datacenter.jin10.com/reportType/dc_cftc_merchant_goods
func MacroUsaCftcMerchantGoodsHolding() ([]CFTCItem, error) {
	return cftcBaseFunc(
		"https://cdn.jin10.com/data_center/reports/cftc_1.json",
		"美国商品期货交易委员会CFTC商品类商业持仓报告",
	)
}

// MacroUsaCmeMerchantGoodsHolding CME贵金属报告
// 数据区间从 2018-04-05 至今
// https://datacenter.jin10.com/org
func MacroUsaCmeMerchantGoodsHolding() ([]CMEItem, error) {
	baseURL := "https://cdn.jin10.com/data_center/reports/cme_3.json"
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)

	params := map[string]string{
		"_": timestamp,
	}

	resp, err := utils.GetWithHeaders(baseURL, params, jin10CommonHeaders)
	if err != nil {
		return nil, fmt.Errorf("获取CME贵金属报告失败: %w", err)
	}

	// 解析 JSON
	json := gjson.ParseBytes(resp.Body())
	values := json.Get("values")
	if !values.Exists() || !values.IsObject() {
		return nil, fmt.Errorf("CME贵金属报告数据格式错误")
	}

	var data []CMEItem

	// 遍历所有日期
	values.ForEach(func(date, value gjson.Result) bool {
		dateStr := date.String()

		// 每个日期下有多条记录
		if value.IsArray() {
			for _, record := range value.Array() {
				if !record.IsArray() || len(record.Array()) < 6 {
					continue
				}

				arr := record.Array()
				// 数据格式: [pz, tc, -, -, -, 成交量, -, -, ...]
				pz := arr[0].String()
				tc := arr[1].String()
				volume := arr[5].Float()

				variety := pz + "-" + tc

				data = append(data, CMEItem{
					Date:    dateStr,
					Variety: variety,
					Volume:  volume,
				})
			}
		}

		return true // continue iteration
	})

	return data, nil
}
