package futures

import (
	"testing"
)

func TestFuturesSymbolMarkFunc(t *testing.T) {
	records, err := FuturesSymbolMarkFunc()
	if err != nil {
		t.Fatalf("FuturesSymbolMarkFunc failed: %v", err)
	}

	// 允许空数据
	t.Logf("FuturesSymbolMarkFunc returned %d records", len(records))

	// 如果有数据，验证结构
	if len(records) > 0 {
		first := records[0]
		if first.Symbol == "" {
			t.Error("First record has empty Symbol")
		}
	}
}

func TestFuturesZhRealtimeFunc(t *testing.T) {
	// 测试黄金主力合约
	records, err := FuturesZhRealtimeFunc("AU0")
	if err != nil {
		t.Logf("FuturesZhRealtimeFunc warning: %v (may be unavailable)", err)
		return
	}

	// 允许空数据
	t.Logf("FuturesZhRealtimeFunc returned %d records", len(records))

	// 如果有数据，验证结构
	if len(records) > 0 {
		first := records[0]
		if first.Symbol == "" {
			t.Error("First record has empty Symbol")
		}
	}
}

func TestFuturesZhSpotFunc(t *testing.T) {
	// 测试上期所
	records, err := FuturesZhSpotFunc("AU0", "SHFE")
	if err != nil {
		t.Logf("FuturesZhSpotFunc warning: %v (may be unavailable)", err)
		return
	}

	// 允许空数据
	t.Logf("FuturesZhSpotFunc returned %d records", len(records))
}

func TestFuturesZhMinuteSina(t *testing.T) {
	// 测试黄金主力1分钟数据
	records, err := FuturesZhMinuteSina("AU0", "1")
	if err != nil {
		t.Logf("FuturesZhMinuteSina warning: %v (may be unavailable)", err)
		return
	}

	// 允许空数据
	t.Logf("FuturesZhMinuteSina returned %d records", len(records))
}

func TestFuturesZhDailySina(t *testing.T) {
	// 测试黄金主力日线数据
	records, err := FuturesZhDailySina("AU0")
	if err != nil {
		t.Logf("FuturesZhDailySina warning: %v (may be unavailable)", err)
		return
	}

	// 允许空数据
	t.Logf("FuturesZhDailySina returned %d records", len(records))

	// 如果有数据，验证结构
	if len(records) > 0 {
		first := records[0]
		if first.Date.IsZero() {
			t.Error("First record has zero Date")
		}
	}
}
