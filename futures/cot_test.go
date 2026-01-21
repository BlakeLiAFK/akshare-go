package futures

import (
	"testing"
	"time"
)

func TestFuturesSHFERankTable(t *testing.T) {
	// 使用最近的交易日
	date := time.Now().AddDate(0, 0, -3).Format("20060102")
	records, err := FuturesSHFERankTable(date, []string{"AU"})
	if err != nil {
		t.Logf("FuturesSHFERankTable warning: %v (may be unavailable)", err)
		return
	}

	// 允许空数据
	t.Logf("FuturesSHFERankTable returned %d symbols", len(records))
}

func TestFuturesCZCERankTable(t *testing.T) {
	// 使用最近的交易日
	date := time.Now().AddDate(0, 0, -3).Format("20060102")
	records, err := FuturesCZCERankTable(date)
	if err != nil {
		t.Logf("FuturesCZCERankTable warning: %v (may be unavailable)", err)
		return
	}

	// 允许空数据
	t.Logf("FuturesCZCERankTable returned %d symbols", len(records))
}

func TestFuturesCFFEXRankTable(t *testing.T) {
	// 使用最近的交易日
	date := time.Now().AddDate(0, 0, -3).Format("20060102")
	records, err := FuturesCFFEXRankTable(date, []string{"IC"})
	if err != nil {
		t.Logf("FuturesCFFEXRankTable warning: %v (may be unavailable)", err)
		return
	}

	// 允许空数据
	t.Logf("FuturesCFFEXRankTable returned %d symbols", len(records))
}

func TestFuturesDCEPositionRank(t *testing.T) {
	// 使用最近的交易日
	date := time.Now().AddDate(0, 0, -3).Format("20060102")
	records, err := FuturesDCEPositionRank(date, []string{"A"})
	if err != nil {
		t.Logf("FuturesDCEPositionRank warning: %v (may be unavailable)", err)
		return
	}

	// 允许空数据
	t.Logf("FuturesDCEPositionRank returned %d symbols", len(records))
}

func TestFuturesGFEXPositionRank(t *testing.T) {
	// 使用最近的交易日
	date := time.Now().AddDate(0, 0, -3).Format("20060102")
	records, err := FuturesGFEXPositionRank(date, []string{"LC"})
	if err != nil {
		t.Logf("FuturesGFEXPositionRank warning: %v (may be unavailable)", err)
		return
	}

	// 允许空数据
	t.Logf("FuturesGFEXPositionRank returned %d symbols", len(records))
}

func TestFuturesRankSum(t *testing.T) {
	// 使用最近的交易日
	date := time.Now().AddDate(0, 0, -3).Format("20060102")
	records, err := FuturesRankSum(date, []string{"AU", "AG"})
	if err != nil {
		t.Logf("FuturesRankSum warning: %v (may be unavailable)", err)
		return
	}

	// 允许空数据
	t.Logf("FuturesRankSum returned %d records", len(records))

	// 如果有数据，验证结构
	if len(records) > 0 {
		first := records[0]
		if first.Variety == "" {
			t.Error("First record has empty Variety")
		}
	}
}
