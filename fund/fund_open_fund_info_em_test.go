package fund

import (
	"testing"
)

func TestFundOpenFundInfoEm(t *testing.T) {
	testCases := []struct {
		symbol    string
		indicator string
		period    string
	}{
		{"710001", "单位净值走势", ""},
		{"710001", "累计净值走势", ""},
		{"710001", "累计收益率走势", "成立来"},
		{"710001", "同类排名走势", ""},
		{"710001", "同类排名百分比", ""},
		{"710001", "分红送配详情", ""},
		{"710001", "拆分详情", ""},
		{"000001", "单位净值走势", ""},
		{"000001", "累计收益率走势", "1年"},
	}

	for _, tc := range testCases {
		t.Run(tc.symbol+"_"+tc.indicator, func(t *testing.T) {
			result, err := FundOpenFundInfoEm(tc.symbol, tc.indicator, tc.period)
			if err != nil {
				t.Logf("FundOpenFundInfoEm(%s, %s, %s) warning: %v (may be unavailable)", tc.symbol, tc.indicator, tc.period, err)
				return
			}

			t.Logf("FundOpenFundInfoEm(%s, %s, %s) returned %d records", tc.symbol, tc.indicator, tc.period, len(result))

			if len(result) > 0 {
				t.Logf("First record: %+v", result[0])
			}
		})
	}
}
