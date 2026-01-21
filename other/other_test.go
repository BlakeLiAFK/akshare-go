package other

import (
	"testing"
)

// TestCarMarketTotalCPCA 测试乘联会-统计数据-总体市场
func TestCarMarketTotalCPCA(t *testing.T) {
	tests := []struct {
		name      string
		symbol    string
		indicator string
	}{
		{"狭义乘用车-产量", CarTypeNarrow, IndicatorProduction},
		{"狭义乘用车-批发", CarTypeNarrow, IndicatorWholesale},
		{"狭义乘用车-零售", CarTypeNarrow, IndicatorRetail},
		{"狭义乘用车-出口", CarTypeNarrow, IndicatorExport},
		{"广义乘用车-产量", CarTypeWide, IndicatorProduction},
		{"广义乘用车-批发", CarTypeWide, IndicatorWholesale},
		{"广义乘用车-零售", CarTypeWide, IndicatorRetail},
		{"广义乘用车-出口", CarTypeWide, IndicatorExport},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if testing.Short() {
				t.Skip("跳过外部网络测试")
			}
			data, err := CarMarketTotalCPCA(tt.symbol, tt.indicator)
			if err != nil {
				t.Logf("CarMarketTotalCPCA(%s, %s) failed: %v", tt.symbol, tt.indicator, err)
				return
			}
			if len(data) > 0 {
				t.Logf("CarMarketTotalCPCA(%s, %s): got %d items", tt.symbol, tt.indicator, len(data))
			} else {
				t.Log("CarMarketTotalCPCA: no data returned")
			}
		})
	}
}

// TestCarMarketManRankCPCA 测试乘联会-统计数据-厂商排名
func TestCarMarketManRankCPCA(t *testing.T) {
	tests := []struct {
		name      string
		symbol    string
		indicator string
	}{
		{"狭义乘用车-单月-批发", CarTypeNarrow + "-" + StatTypeMonthly, IndicatorWholesale},
		{"狭义乘用车-累计-批发", CarTypeNarrow + "-" + StatTypeCumulative, IndicatorWholesale},
		{"广义乘用车-单月-批发", CarTypeWide + "-" + StatTypeMonthly, IndicatorWholesale},
		{"广义乘用车-累计-批发", CarTypeWide + "-" + StatTypeCumulative, IndicatorWholesale},
		{"狭义乘用车-单月-零售", CarTypeNarrow + "-" + StatTypeMonthly, IndicatorRetail},
		{"狭义乘用车-累计-零售", CarTypeNarrow + "-" + StatTypeCumulative, IndicatorRetail},
		{"广义乘用车-单月-零售", CarTypeWide + "-" + StatTypeMonthly, IndicatorRetail},
		{"广义乘用车-累计-零售", CarTypeWide + "-" + StatTypeCumulative, IndicatorRetail},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if testing.Short() {
				t.Skip("跳过外部网络测试")
			}
			data, err := CarMarketManRankCPCA(tt.symbol, tt.indicator)
			if err != nil {
				t.Logf("CarMarketManRankCPCA(%s, %s) failed: %v", tt.symbol, tt.indicator, err)
				return
			}
			if len(data) > 0 {
				t.Logf("CarMarketManRankCPCA(%s, %s): got %d items", tt.symbol, tt.indicator, len(data))
			} else {
				t.Log("CarMarketManRankCPCA: no data returned")
			}
		})
	}
}

// TestCarMarketCateCPCA 测试乘联会-统计数据-车型大类
func TestCarMarketCateCPCA(t *testing.T) {
	tests := []struct {
		name      string
		symbol    string
		indicator string
	}{
		{"MPV-批发", CategoryMPV, IndicatorWholesale},
		{"SUV-批发", CategorySUV, IndicatorWholesale},
		{"轿车-批发", CategorySedan, IndicatorWholesale},
		{"MPV-零售", CategoryMPV, IndicatorRetail},
		{"SUV-零售", CategorySUV, IndicatorRetail},
		{"轿车-零售", CategorySedan, IndicatorRetail},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if testing.Short() {
				t.Skip("跳过外部网络测试")
			}
			data, err := CarMarketCateCPCA(tt.symbol, tt.indicator)
			if err != nil {
				t.Logf("CarMarketCateCPCA(%s, %s) failed: %v", tt.symbol, tt.indicator, err)
				return
			}
			if len(data) > 0 {
				t.Logf("CarMarketCateCPCA(%s, %s): got %d items", tt.symbol, tt.indicator, len(data))
			} else {
				t.Log("CarMarketCateCPCA: no data returned")
			}
		})
	}
}

// TestCarMarketCountryCPCA 测试乘联会-统计数据-国别细分市场
func TestCarMarketCountryCPCA(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过外部网络测试")
	}
	data, err := CarMarketCountryCPCA()
	if err != nil {
		t.Logf("CarMarketCountryCPCA failed: %v", err)
		return
	}
	if len(data) > 0 {
		t.Logf("CarMarketCountryCPCA: got %d items", len(data))
	} else {
		t.Log("CarMarketCountryCPCA: no data returned")
	}
}

// TestCarMarketSegmentCPCA 测试乘联会-统计数据-级别细分市场
func TestCarMarketSegmentCPCA(t *testing.T) {
	tests := []struct {
		name   string
		symbol string
	}{
		{"MPV", CategoryMPV},
		{"SUV", CategorySUV},
		{"轿车", CategorySedan},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if testing.Short() {
				t.Skip("跳过外部网络测试")
			}
			data, err := CarMarketSegmentCPCA(tt.symbol)
			if err != nil {
				t.Logf("CarMarketSegmentCPCA(%s) failed: %v", tt.symbol, err)
				return
			}
			if len(data) > 0 {
				t.Logf("CarMarketSegmentCPCA(%s): got %d items", tt.symbol, len(data))
			} else {
				t.Log("CarMarketSegmentCPCA: no data returned")
			}
		})
	}
}

// TestCarMarketFuelCPCA 测试乘联会-统计数据-新能源细分市场
func TestCarMarketFuelCPCA(t *testing.T) {
	tests := []struct {
		name   string
		symbol string
	}{
		{"整体市场", FuelOverallMarket},
		{"销量占比-PHEV-BEV", FuelRatioPHEVBEV},
		{"销量占比-ICE-NEV", FuelRatioICENEV},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if testing.Short() {
				t.Skip("跳过外部网络测试")
			}
			data, err := CarMarketFuelCPCA(tt.symbol)
			if err != nil {
				t.Logf("CarMarketFuelCPCA(%s) failed: %v", tt.symbol, err)
				return
			}
			if len(data) > 0 {
				t.Logf("CarMarketFuelCPCA(%s): got %d items", tt.symbol, len(data))
			} else {
				t.Log("CarMarketFuelCPCA: no data returned")
			}
		})
	}
}

// TestCarSaleRankGasgoo 测试盖世汽车-销量排名
func TestCarSaleRankGasgoo(t *testing.T) {
	tests := []struct {
		name   string
		symbol string
		date   string
	}{
		{"车企榜", RankTypeFirm, "202311"},
		{"品牌榜", RankTypeBrand, "202311"},
		{"车型榜", RankTypeModel, "202311"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if testing.Short() {
				t.Skip("跳过外部网络测试")
			}
			data, err := CarSaleRankGasgoo(tt.symbol, tt.date)
			if err != nil {
				t.Logf("CarSaleRankGasgoo(%s, %s) failed: %v", tt.symbol, tt.date, err)
				return
			}
			if len(data) > 0 {
				t.Logf("CarSaleRankGasgoo(%s, %s): got %d items", tt.symbol, tt.date, len(data))
			} else {
				t.Log("CarSaleRankGasgoo: no data returned")
			}
		})
	}
}

// BenchmarkCarMarketTotalCPCA 基准测试-总体市场
func BenchmarkCarMarketTotalCPCA(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		_, _ = CarMarketTotalCPCA(CarTypeNarrow, IndicatorProduction)
	}
}

// BenchmarkCarSaleRankGasgoo 基准测试-盖世汽车销量排名
func BenchmarkCarSaleRankGasgoo(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		_, _ = CarSaleRankGasgoo(RankTypeFirm, "202311")
	}
}
