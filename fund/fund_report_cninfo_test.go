package fund

import (
	"testing"
)

func TestFundReportStockCninfo(t *testing.T) {
	testCases := []struct {
		date string
	}{
		{"20210630"},
		{"20240630"},
	}

	for _, tc := range testCases {
		t.Run(tc.date, func(t *testing.T) {
			result, err := FundReportStockCninfo(tc.date)
			if err != nil {
				t.Logf("FundReportStockCninfo(%s) warning: %v (may be unavailable)", tc.date, err)
				return
			}

			t.Logf("FundReportStockCninfo(%s) returned %d records", tc.date, len(result))

			if len(result) > 0 {
				t.Logf("First record: %+v", result[0])
				// 验证必需字段
				first := result[0]
				requiredFields := []string{"序号", "股票代码", "股票简称", "报告期", "基金覆盖家数", "持股总数", "持股总市值"}
				for _, field := range requiredFields {
					if _, ok := first[field]; !ok {
						t.Errorf("Missing required field: %s", field)
					}
				}
			}
		})
	}
}

func TestFundReportIndustryAllocationCninfo(t *testing.T) {
	testCases := []struct {
		date string
	}{
		{"20210630"},
		{"20240630"},
	}

	for _, tc := range testCases {
		t.Run(tc.date, func(t *testing.T) {
			result, err := FundReportIndustryAllocationCninfo(tc.date)
			if err != nil {
				t.Logf("FundReportIndustryAllocationCninfo(%s) warning: %v (may be unavailable)", tc.date, err)
				return
			}

			t.Logf("FundReportIndustryAllocationCninfo(%s) returned %d records", tc.date, len(result))

			if len(result) > 0 {
				t.Logf("First record: %+v", result[0])
				// 验证必需字段
				first := result[0]
				requiredFields := []string{"行业编码", "证监会行业名称", "报告期", "基金覆盖家数", "行业规模", "占净资产比例"}
				for _, field := range requiredFields {
					if _, ok := first[field]; !ok {
						t.Errorf("Missing required field: %s", field)
					}
				}
			}
		})
	}
}

func TestFundReportAssetAllocationCninfo(t *testing.T) {
	result, err := FundReportAssetAllocationCninfo()
	if err != nil {
		t.Logf("FundReportAssetAllocationCninfo() warning: %v (may be unavailable)", err)
		return
	}

	t.Logf("FundReportAssetAllocationCninfo() returned %d records", len(result))

	if len(result) > 0 {
		t.Logf("First record: %+v", result[0])
		// 验证必需字段
		first := result[0]
		requiredFields := []string{"报告期", "基金覆盖家数", "股票权益类占净资产比例", "债券固定收益类占净资产比例", "现金货币类占净资产比例", "基金市场净资产规模"}
		for _, field := range requiredFields {
			if _, ok := first[field]; !ok {
				t.Errorf("Missing required field: %s", field)
			}
		}
	}
}

func TestFundReportCombineCninfo(t *testing.T) {
	_, err := FundReportCombineCninfo("000001")
	if err == nil {
		t.Error("FundReportCombineCninfo should return error (not implemented)")
	}
	t.Logf("FundReportCombineCninfo correctly returns error: %v", err)
}
