package fund

import (
	"fmt"
	"strings"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/tidwall/gjson"
)

// FundInfoXq 获取雪球基金基本信息
// https://danjuanfunds.com/djapi/fund/000001
func FundInfoXq(symbol string) (map[string]interface{}, error) {
	if symbol == "" {
		symbol = "000001"
	}

	url := fmt.Sprintf("https://danjuanfunds.com/djapi/fund/%s", symbol)
	resp, err := utils.Get(url, nil)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	result := gjson.ParseBytes(resp.Body())
	data := result.Get("data")

	info := map[string]interface{}{
		"基金代码":   data.Get("fd_code").String(),
		"基金名称":   data.Get("fd_name").String(),
		"基金全称":   data.Get("fd_full_name").String(),
		"成立时间":   data.Get("found_date").String(),
		"最新规模":   utils.MustFloat64(data.Get("totshare").String()),
		"基金公司":   data.Get("keeper_name").String(),
		"基金经理":   data.Get("manager_name").String(),
		"托管银行":   data.Get("trup_name").String(),
		"基金类型":   data.Get("type_desc").String(),
		"评级机构":   data.Get("rating_source").String(),
		"基金评级":   data.Get("rating_desc").String(),
		"投资策略":   data.Get("invest_orientation").String(),
		"投资目标":   data.Get("invest_target").String(),
		"业绩比较基准": data.Get("performance_bench_mark").String(),
	}

	return info, nil
}

// FundHistXq 获取雪球基金历史业绩数据
// https://danjuanfunds.com/djapi/fundx/base/fund/achievement/000001
func FundHistXq(symbol string) ([]map[string]interface{}, error) {
	if symbol == "" {
		symbol = "000001"
	}

	url := fmt.Sprintf("https://danjuanfunds.com/djapi/fundx/base/fund/achievement/%s", symbol)
	resp, err := utils.Get(url, nil)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	result := gjson.ParseBytes(resp.Body())
	data := result.Get("data")

	records := make([]map[string]interface{}, 0)

	// 年度业绩
	annualList := data.Get("annual_performance_list").Array()
	for _, item := range annualList {
		record := map[string]interface{}{
			"业绩类型":     "年度业绩",
			"周期":       item.Get("period_time").String(),
			"本产品区间收益":  parsePercent(item.Get("self_nav").String()),
			"本产品最大回撤":  parsePercent(item.Get("self_max_draw_down").String()),
			"周期收益同类排名": item.Get("self_nav_rank").String(),
		}
		records = append(records, record)
	}

	// 阶段业绩
	stageList := data.Get("stage_performance_list").Array()
	for _, item := range stageList {
		record := map[string]interface{}{
			"业绩类型":     "阶段业绩",
			"周期":       item.Get("period_time").String(),
			"本产品区间收益":  parsePercent(item.Get("self_nav").String()),
			"本产品最大回撤":  parsePercent(item.Get("self_max_draw_down").String()),
			"周期收益同类排名": item.Get("self_nav_rank").String(),
		}
		records = append(records, record)
	}

	return records, nil
}

// FundMarketXq 获取雪球基金数据分析
// https://danjuanfunds.com/djapi/fund/base/quote/data/index/analysis/000001
func FundMarketXq(symbol string) (map[string]interface{}, error) {
	if symbol == "" {
		symbol = "000001"
	}

	url := fmt.Sprintf("https://danjuanfunds.com/djapi/fund/base/quote/data/index/analysis/%s", symbol)
	resp, err := utils.Get(url, nil)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	result := gjson.ParseBytes(resp.Body())
	dataList := result.Get("data.index_data_list").Array()

	if len(dataList) == 0 {
		return nil, fmt.Errorf("无数据")
	}

	// 返回第一个周期的数据作为当前行情
	item := dataList[0]
	market := map[string]interface{}{
		"周期":       item.Get("index_time_period").String(),
		"较同类风险收益比": utils.MustFloat64(item.Get("investment_cost_performance").String()),
		"较同类抗风险波动": utils.MustFloat64(item.Get("risk_control").String()),
		"年化波动率":    utils.MustFloat64(item.Get("self_index.volatility_rank").String()) * 100,
		"年化夏普比率":   utils.MustFloat64(item.Get("self_index.sharpe_rank").String()),
		"最大回撤":     utils.MustFloat64(item.Get("self_index.max_draw_down").String()) * 100,
	}

	return market, nil
}

// FundNetValueXq 获取雪球基金盈利概率数据
// https://danjuanfunds.com/djapi/fundx/base/fund/profit/ratio/000001
func FundNetValueXq(symbol string) ([]map[string]interface{}, error) {
	if symbol == "" {
		symbol = "000001"
	}

	url := fmt.Sprintf("https://danjuanfunds.com/djapi/fundx/base/fund/profit/ratio/%s", symbol)
	resp, err := utils.Get(url, nil)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	result := gjson.ParseBytes(resp.Body())
	dataList := result.Get("data.data_list").Array()

	records := make([]map[string]interface{}, 0, len(dataList))
	for _, item := range dataList {
		record := map[string]interface{}{
			"持有时长": item.Get("holding_time").String(),
			"盈利概率": parsePercent(item.Get("profit_ratio").String()),
			"平均收益": parsePercent(item.Get("average_income").String()),
		}
		records = append(records, record)
	}

	return records, nil
}

// FundPortfolioXq 获取雪球基金持仓数据
// https://danjuanfunds.com/djapi/fundx/base/fund/record/asset/percent
func FundPortfolioXq(symbol string) ([]map[string]interface{}, error) {
	if symbol == "" {
		symbol = "002804"
	}

	url := "https://danjuanfunds.com/djapi/fundx/base/fund/record/asset/percent"
	params := map[string]string{
		"fund_code":   symbol,
		"report_date": "", // 默认使用最新报告期
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	result := gjson.ParseBytes(resp.Body())
	chartList := result.Get("data.chart_list").Array()

	records := make([]map[string]interface{}, 0, len(chartList))
	for _, item := range chartList {
		record := map[string]interface{}{
			"资产类型": item.Get("type_desc").String(),
			"仓位占比": utils.MustFloat64(item.Get("percent").String()),
		}
		records = append(records, record)
	}

	return records, nil
}

// FundSearchXq 获取雪球基金交易规则
// https://danjuanfunds.com/djapi/fund/detail/000001
func FundSearchXq(keyword string) ([]map[string]interface{}, error) {
	if keyword == "" {
		keyword = "000001"
	}

	url := fmt.Sprintf("https://danjuanfunds.com/djapi/fund/detail/%s", keyword)
	resp, err := utils.Get(url, nil)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	result := gjson.ParseBytes(resp.Body())
	fundRates := result.Get("data.fund_rates")

	records := make([]map[string]interface{}, 0)

	// 买入规则
	declareRates := fundRates.Get("declare_rate_table").Array()
	for _, item := range declareRates {
		record := map[string]interface{}{
			"费用类型":  "买入规则",
			"条件或名称": item.Get("name").String(),
			"费用":    utils.MustFloat64(item.Get("value").String()),
		}
		records = append(records, record)
	}

	// 卖出规则
	withdrawRates := fundRates.Get("withdraw_rate_table").Array()
	for _, item := range withdrawRates {
		record := map[string]interface{}{
			"费用类型":  "卖出规则",
			"条件或名称": item.Get("name").String(),
			"费用":    utils.MustFloat64(item.Get("value").String()),
		}
		records = append(records, record)
	}

	// 其他费用
	otherRates := fundRates.Get("other_rate_table").Array()
	for _, item := range otherRates {
		record := map[string]interface{}{
			"费用类型":  "其他费用",
			"条件或名称": item.Get("name").String(),
			"费用":    utils.MustFloat64(item.Get("value").String()),
		}
		records = append(records, record)
	}

	return records, nil
}

// parsePercent 解析百分比字符串，去掉%符号并转换为float64
func parsePercent(s string) float64 {
	s = strings.TrimSuffix(s, "%")
	return utils.MustFloat64(s)
}
