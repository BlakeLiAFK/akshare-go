package fund

import (
	"testing"
)

func TestFundInfoXq(t *testing.T) {
	info, err := FundInfoXq("000001")
	if err != nil {
		t.Logf("FundInfoXq warning: %v (may be unavailable)", err)
		return
	}

	// 验证必要字段
	requiredFields := []string{"基金代码", "基金名称", "基金全称", "成立时间", "基金公司", "基金经理"}
	for _, field := range requiredFields {
		if _, ok := info[field]; !ok {
			t.Errorf("Missing required field: %s", field)
		}
	}

	t.Logf("FundInfoXq returned info for fund: %v", info["基金名称"])
}

func TestFundHistXq(t *testing.T) {
	records, err := FundHistXq("000001")
	if err != nil {
		t.Logf("FundHistXq warning: %v (may be unavailable)", err)
		return
	}

	// 允许空数据
	t.Logf("FundHistXq returned %d records", len(records))

	// 如果有数据，验证结构
	if len(records) > 0 {
		first := records[0]
		requiredFields := []string{"业绩类型", "周期", "本产品区间收益", "本产品最大回撤", "周期收益同类排名"}
		for _, field := range requiredFields {
			if _, ok := first[field]; !ok {
				t.Errorf("Missing required field: %s", field)
			}
		}
	}
}

func TestFundMarketXq(t *testing.T) {
	market, err := FundMarketXq("000001")
	if err != nil {
		t.Logf("FundMarketXq warning: %v (may be unavailable)", err)
		return
	}

	// 验证必要字段
	requiredFields := []string{"周期", "较同类风险收益比", "较同类抗风险波动", "年化波动率", "年化夏普比率", "最大回撤"}
	for _, field := range requiredFields {
		if _, ok := market[field]; !ok {
			t.Errorf("Missing required field: %s", field)
		}
	}

	t.Logf("FundMarketXq returned market data for period: %v", market["周期"])
}

func TestFundNetValueXq(t *testing.T) {
	records, err := FundNetValueXq("000001")
	if err != nil {
		t.Logf("FundNetValueXq warning: %v (may be unavailable)", err)
		return
	}

	// 允许空数据
	t.Logf("FundNetValueXq returned %d records", len(records))

	// 如果有数据，验证结构
	if len(records) > 0 {
		first := records[0]
		requiredFields := []string{"持有时长", "盈利概率", "平均收益"}
		for _, field := range requiredFields {
			if _, ok := first[field]; !ok {
				t.Errorf("Missing required field: %s", field)
			}
		}
	}
}

func TestFundPortfolioXq(t *testing.T) {
	records, err := FundPortfolioXq("002804")
	if err != nil {
		t.Logf("FundPortfolioXq warning: %v (may be unavailable)", err)
		return
	}

	// 允许空数据
	t.Logf("FundPortfolioXq returned %d records", len(records))

	// 如果有数据，验证结构
	if len(records) > 0 {
		first := records[0]
		requiredFields := []string{"资产类型", "仓位占比"}
		for _, field := range requiredFields {
			if _, ok := first[field]; !ok {
				t.Errorf("Missing required field: %s", field)
			}
		}
	}
}

func TestFundSearchXq(t *testing.T) {
	records, err := FundSearchXq("000001")
	if err != nil {
		t.Logf("FundSearchXq warning: %v (may be unavailable)", err)
		return
	}

	// 允许空数据
	t.Logf("FundSearchXq returned %d records", len(records))

	// 如果有数据，验证结构
	if len(records) > 0 {
		first := records[0]
		requiredFields := []string{"费用类型", "条件或名称", "费用"}
		for _, field := range requiredFields {
			if _, ok := first[field]; !ok {
				t.Errorf("Missing required field: %s", field)
			}
		}
	}
}
