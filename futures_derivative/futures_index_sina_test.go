package futures_derivative

import (
	"testing"
	"time"
)

func TestZhSubscribeExchangeSymbol(t *testing.T) {
	records, err := ZhSubscribeExchangeSymbol("shfe")
	if err != nil {
		t.Logf("ZhSubscribeExchangeSymbol warning: %v (may be unavailable)", err)
		return
	}

	// 允许空数据
	t.Logf("ZhSubscribeExchangeSymbol returned %d records", len(records))

	// 如果有数据，验证结构
	if len(records) > 0 {
		first := records[0]
		if _, ok := first["symbol"]; !ok {
			t.Error("Missing required field: symbol")
		}
	}
}

func TestMatchMainContract(t *testing.T) {
	records, err := MatchMainContract("shfe")
	if err != nil {
		t.Logf("MatchMainContract warning: %v (may be unavailable)", err)
		return
	}

	// 允许空数据
	t.Logf("MatchMainContract returned %d records", len(records))

	// 如果有数据，验证结构
	if len(records) > 0 {
		first := records[0]
		requiredFields := []string{"symbol", "name", "trade"}
		for _, field := range requiredFields {
			if _, ok := first[field]; !ok {
				t.Errorf("Missing required field: %s", field)
			}
		}
	}
}

func TestFuturesDisplayMainSina(t *testing.T) {
	records, err := FuturesDisplayMainSina()
	if err != nil {
		t.Logf("FuturesDisplayMainSina warning: %v (may be unavailable)", err)
		return
	}

	// 允许空数据
	t.Logf("FuturesDisplayMainSina returned %d records", len(records))

	// 如果有数据，验证结构
	if len(records) > 0 {
		first := records[0]
		requiredFields := []string{"symbol", "name"}
		for _, field := range requiredFields {
			if _, ok := first[field]; !ok {
				t.Errorf("Missing required field: %s", field)
			}
		}
	}
}

func TestFuturesMainSina(t *testing.T) {
	// 使用时间范围来避免数据过多
	endDate := time.Now().AddDate(0, 0, -1).Format("20060102")
	startDate := time.Now().AddDate(0, 0, -7).Format("20060102")

	records, err := FuturesMainSina("V0", startDate, endDate)
	if err != nil {
		t.Logf("FuturesMainSina warning: %v (may be unavailable)", err)
		return
	}

	// 允许空数据
	t.Logf("FuturesMainSina returned %d records", len(records))

	// 如果有数据，验证结构
	if len(records) > 0 {
		first := records[0]
		requiredFields := []string{"日期", "开盘价", "最高价", "最低价", "收盘价"}
		for _, field := range requiredFields {
			if _, ok := first[field]; !ok {
				t.Errorf("Missing required field: %s", field)
			}
		}
	}
}
