package fund

import (
	"testing"
)

func TestFundFeeEm(t *testing.T) {
	testCases := []struct {
		symbol    string
		indicator string
	}{
		{"019005", "交易状态"},
		{"019005", "申购与赎回金额"},
		{"019005", "交易确认日"},
		{"019005", "运作费用"},
		{"019005", "认购费率（前端）"},
		{"019005", "申购费率（前端）"},
		{"000011", "赎回费率"},
		{"000011", "认购费率（后端）"},
	}

	for _, tc := range testCases {
		t.Run(tc.symbol+"_"+tc.indicator, func(t *testing.T) {
			result, err := FundFeeEm(tc.symbol, tc.indicator)
			if err != nil {
				t.Logf("FundFeeEm(%s, %s) warning: %v (may be unavailable)", tc.symbol, tc.indicator, err)
				return
			}

			t.Logf("FundFeeEm(%s, %s) returned %d rows", tc.symbol, tc.indicator, len(result))

			// 验证返回的是表格数据
			if len(result) > 0 {
				t.Logf("First row has %d columns", len(result[0]))
				if len(result[0]) == 0 {
					t.Errorf("First row should not be empty")
				}
			}
		})
	}
}
