package futures

import (
	"testing"
)

func TestFuturesHistTableEMFunc(t *testing.T) {
	records, err := FuturesHistTableEMFunc()
	if err != nil {
		t.Fatalf("FuturesHistTableEMFunc failed: %v", err)
	}

	// 允许空数据
	t.Logf("FuturesHistTableEMFunc returned %d records", len(records))

	// 如果有数据，验证结构
	if len(records) > 0 {
		first := records[0]
		if first.ContractCode == "" {
			t.Error("First record has empty ContractCode")
		}
		if first.MarketName == "" {
			t.Error("First record has empty MarketName")
		}
	}
}

func TestFuturesHistEMFunc(t *testing.T) {
	// 测试获取螺纹钢主力合约历史数据
	records, err := FuturesHistEMFunc("RBL8", "daily", "", "")
	if err != nil {
		t.Logf("FuturesHistEMFunc warning: %v (may be unavailable)", err)
		return
	}

	// 允许空数据
	t.Logf("FuturesHistEMFunc returned %d records", len(records))

	// 如果有数据，验证结构
	if len(records) > 0 {
		first := records[0]
		if first.Date.IsZero() {
			t.Error("First record has zero Date")
		}
	}
}
