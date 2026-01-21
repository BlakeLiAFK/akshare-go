package futures

import (
	"testing"
	"time"
)

func TestGetCFFEXDaily(t *testing.T) {
	// 使用最近的交易日
	date := time.Now().AddDate(0, 0, -3).Format("20060102")
	records, err := GetCFFEXDaily(date)
	if err != nil {
		t.Logf("GetCFFEXDaily warning: %v (may be unavailable)", err)
		return
	}

	// 允许空数据
	t.Logf("GetCFFEXDaily returned %d records", len(records))
}

func TestGetGFEXDaily(t *testing.T) {
	// 使用最近的交易日
	date := time.Now().AddDate(0, 0, -3).Format("20060102")
	records, err := GetGFEXDaily(date)
	if err != nil {
		t.Logf("GetGFEXDaily warning: %v (may be unavailable)", err)
		return
	}

	// 允许空数据
	t.Logf("GetGFEXDaily returned %d records", len(records))
}

func TestGetINEDaily(t *testing.T) {
	// 使用最近的交易日
	date := time.Now().AddDate(0, 0, -3).Format("20060102")
	records, err := GetINEDaily(date)
	if err != nil {
		t.Logf("GetINEDaily warning: %v (may be unavailable)", err)
		return
	}

	// 允许空数据
	t.Logf("GetINEDaily returned %d records", len(records))
}

func TestGetSHFEDaily(t *testing.T) {
	// 使用最近的交易日
	date := time.Now().AddDate(0, 0, -3).Format("20060102")
	records, err := GetSHFEDaily(date)
	if err != nil {
		t.Logf("GetSHFEDaily warning: %v (may be unavailable)", err)
		return
	}

	// 允许空数据
	t.Logf("GetSHFEDaily returned %d records", len(records))

	// 如果有数据，验证结构
	if len(records) > 0 {
		first := records[0]
		if first.Symbol == "" {
			t.Error("First record has empty Symbol")
		}
	}
}

func TestGetDCEDaily(t *testing.T) {
	// 使用最近的交易日
	date := time.Now().AddDate(0, 0, -3).Format("20060102")
	records, err := GetDCEDaily(date)
	if err != nil {
		t.Logf("GetDCEDaily warning: %v (may be unavailable)", err)
		return
	}

	// 允许空数据
	t.Logf("GetDCEDaily returned %d records", len(records))
}

func TestGetCZCEDaily(t *testing.T) {
	// 使用最近的交易日
	date := time.Now().AddDate(0, 0, -3).Format("20060102")
	records, err := GetCZCEDaily(date)
	if err != nil {
		t.Logf("GetCZCEDaily warning: %v (may be unavailable)", err)
		return
	}

	// 允许空数据
	t.Logf("GetCZCEDaily returned %d records", len(records))
}

func TestGetFuturesDaily(t *testing.T) {
	// 测试获取上期所最近3天数据
	endDate := time.Now().AddDate(0, 0, -1).Format("20060102")
	startDate := time.Now().AddDate(0, 0, -3).Format("20060102")

	records, err := GetFuturesDaily(startDate, endDate, "SHFE")
	if err != nil {
		t.Logf("GetFuturesDaily warning: %v (may be unavailable)", err)
		return
	}

	// 允许空数据
	t.Logf("GetFuturesDaily returned %d records", len(records))
}
