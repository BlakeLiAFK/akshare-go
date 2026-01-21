package futures_derivative

import (
	"testing"
	"time"
)

func TestFuturesHoldPosSina(t *testing.T) {
	// 使用最近的交易日
	date := time.Now().AddDate(0, 0, -5).Format("20060102")

	// 测试成交量
	records, err := FuturesHoldPosSina("成交量", "IC2403", date)
	if err != nil {
		t.Logf("FuturesHoldPosSina(成交量) warning: %v (may be unavailable)", err)
		return
	}

	// 允许空数据
	t.Logf("FuturesHoldPosSina(成交量) returned %d records", len(records))

	// 如果有数据，验证结构
	if len(records) > 0 {
		first := records[0]
		requiredFields := []string{"名次", "会员简称", "成交量", "比上交易增减"}
		for _, field := range requiredFields {
			if _, ok := first[field]; !ok {
				t.Errorf("Missing required field: %s", field)
			}
		}
	}

	// 测试多单持仓
	records2, err2 := FuturesHoldPosSina("多单持仓", "OI2501", date)
	if err2 != nil {
		t.Logf("FuturesHoldPosSina(多单持仓) warning: %v (may be unavailable)", err2)
		return
	}

	t.Logf("FuturesHoldPosSina(多单持仓) returned %d records", len(records2))

	// 测试空单持仓
	records3, err3 := FuturesHoldPosSina("空单持仓", "OI2501", date)
	if err3 != nil {
		t.Logf("FuturesHoldPosSina(空单持仓) warning: %v (may be unavailable)", err3)
		return
	}

	t.Logf("FuturesHoldPosSina(空单持仓) returned %d records", len(records3))
}
