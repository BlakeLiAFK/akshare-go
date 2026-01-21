package fund

import (
	"testing"
)

func TestFundFinancialFundInfoEm(t *testing.T) {
	testCases := []struct {
		symbol string
	}{
		{"000134"},
		{"000791"},
	}

	for _, tc := range testCases {
		t.Run(tc.symbol, func(t *testing.T) {
			result, err := FundFinancialFundInfoEm(tc.symbol)
			if err != nil {
				t.Logf("FundFinancialFundInfoEm(%s) warning: %v (may be unavailable)", tc.symbol, err)
				return
			}

			t.Logf("FundFinancialFundInfoEm(%s) returned %d records", tc.symbol, len(result))

			if len(result) > 0 {
				t.Logf("First record: %+v", result[0])
				// 验证必需字段
				first := result[0]
				requiredFields := []string{"净值日期", "单位净值", "累计净值", "日增长率", "申购状态", "赎回状态", "分红送配"}
				for _, field := range requiredFields {
					if _, ok := first[field]; !ok {
						t.Errorf("Missing required field: %s", field)
					}
				}
			}
		})
	}
}

func TestFundGradedFundDailyEm(t *testing.T) {
	result, err := FundGradedFundDailyEm()
	if err != nil {
		t.Logf("FundGradedFundDailyEm() warning: %v (may be unavailable)", err)
		return
	}

	t.Logf("FundGradedFundDailyEm() returned %d records", len(result))

	if len(result) > 0 {
		t.Logf("First record: %+v", result[0])
		// 验证基础字段（日期字段是动态的）
		first := result[0]
		requiredFields := []string{"基金代码", "基金简称", "日增长值", "日增长率", "市价", "折价率", "手续费"}
		for _, field := range requiredFields {
			if _, ok := first[field]; !ok {
				t.Errorf("Missing required field: %s", field)
			}
		}
	}
}

func TestFundGradedFundInfoEm(t *testing.T) {
	testCases := []struct {
		symbol string
	}{
		{"150232"},
	}

	for _, tc := range testCases {
		t.Run(tc.symbol, func(t *testing.T) {
			result, err := FundGradedFundInfoEm(tc.symbol)
			if err != nil {
				t.Logf("FundGradedFundInfoEm(%s) warning: %v (may be unavailable)", tc.symbol, err)
				return
			}

			t.Logf("FundGradedFundInfoEm(%s) returned %d records", tc.symbol, len(result))

			if len(result) > 0 {
				t.Logf("First record: %+v", result[0])
				// 验证必需字段
				first := result[0]
				requiredFields := []string{"净值日期", "单位净值", "累计净值", "日增长率", "申购状态", "赎回状态"}
				for _, field := range requiredFields {
					if _, ok := first[field]; !ok {
						t.Errorf("Missing required field: %s", field)
					}
				}
			}
		})
	}
}

func TestFundEtfFundDailyEm(t *testing.T) {
	result, err := FundEtfFundDailyEm()
	if err != nil {
		t.Logf("FundEtfFundDailyEm() warning: %v (may be unavailable)", err)
		return
	}

	t.Logf("FundEtfFundDailyEm() returned %d records", len(result))

	if len(result) > 0 {
		t.Logf("First record: %+v", result[0])
		// 验证必需字段
		first := result[0]
		requiredFields := []string{"基金代码", "基金简称", "类型", "单位净值", "累计净值", "增长值", "增长率", "市价", "折价率"}
		for _, field := range requiredFields {
			if _, ok := first[field]; !ok {
				t.Errorf("Missing required field: %s", field)
			}
		}
	}
}

func TestFundEtfFundInfoEm(t *testing.T) {
	testCases := []struct {
		symbol    string
		startDate string
		endDate   string
	}{
		{"511280", "20230101", "20231231"},
		{"511280", "", ""},
	}

	for _, tc := range testCases {
		t.Run(tc.symbol, func(t *testing.T) {
			result, err := FundEtfFundInfoEm(tc.symbol, tc.startDate, tc.endDate)
			if err != nil {
				t.Logf("FundEtfFundInfoEm(%s, %s, %s) warning: %v (may be unavailable)", tc.symbol, tc.startDate, tc.endDate, err)
				return
			}

			t.Logf("FundEtfFundInfoEm(%s, %s, %s) returned %d records", tc.symbol, tc.startDate, tc.endDate, len(result))

			if len(result) > 0 {
				t.Logf("First record: %+v", result[0])
				// 验证必需字段
				first := result[0]
				requiredFields := []string{"净值日期", "单位净值", "累计净值", "日增长率", "申购状态", "赎回状态"}
				for _, field := range requiredFields {
					if _, ok := first[field]; !ok {
						t.Errorf("Missing required field: %s", field)
					}
				}
			}
		})
	}
}

func TestFundHkFundHistEm(t *testing.T) {
	testCases := []struct {
		code      string
		indicator string
	}{
		{"1002200683", "历史净值明细"},
		{"1002200683", "分红送配详情"},
	}

	for _, tc := range testCases {
		t.Run(tc.code+"_"+tc.indicator, func(t *testing.T) {
			result, err := FundHkFundHistEm(tc.code, tc.indicator)
			if err != nil {
				t.Logf("FundHkFundHistEm(%s, %s) warning: %v (may be unavailable)", tc.code, tc.indicator, err)
				return
			}

			t.Logf("FundHkFundHistEm(%s, %s) returned %d records", tc.code, tc.indicator, len(result))

			if len(result) > 0 {
				t.Logf("First record: %+v", result[0])

				// 验证必需字段
				first := result[0]
				var requiredFields []string
				if tc.indicator == "历史净值明细" {
					requiredFields = []string{"净值日期", "单位净值", "日增长值", "日增长率", "单位"}
				} else {
					requiredFields = []string{"权益登记日", "除息日", "派息", "单位"}
				}

				for _, field := range requiredFields {
					if _, ok := first[field]; !ok {
						t.Errorf("Missing required field: %s", field)
					}
				}
			}
		})
	}
}
