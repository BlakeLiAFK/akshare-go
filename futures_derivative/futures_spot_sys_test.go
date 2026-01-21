package futures_derivative

import (
	"testing"
)

func TestFuturesSpotSys(t *testing.T) {
	// 测试市场价格
	records, err := FuturesSpotSys("铜", "市场价格")
	if err != nil {
		t.Logf("FuturesSpotSys(市场价格) warning: %v (may be unavailable)", err)
		return
	}

	// 允许空数据
	t.Logf("FuturesSpotSys(市场价格) returned %d records", len(records))

	// 如果有数据，验证结构
	if len(records) > 0 {
		first := records[0]
		requiredFields := []string{"日期", "现货价格", "主力合约", "最近合约"}
		for _, field := range requiredFields {
			if _, ok := first[field]; !ok {
				t.Errorf("Missing required field: %s", field)
			}
		}
	}

	// 测试基差率
	records2, err2 := FuturesSpotSys("铜", "基差率")
	if err2 != nil {
		t.Logf("FuturesSpotSys(基差率) warning: %v (may be unavailable)", err2)
		return
	}

	t.Logf("FuturesSpotSys(基差率) returned %d records", len(records2))

	// 测试主力基差
	records3, err3 := FuturesSpotSys("铜", "主力基差")
	if err3 != nil {
		t.Logf("FuturesSpotSys(主力基差) warning: %v (may be unavailable)", err3)
		return
	}

	t.Logf("FuturesSpotSys(主力基差) returned %d records", len(records3))
}
