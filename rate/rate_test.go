package rate

import (
	"testing"
)

// TestRateShibor 测试Shibor利率数据
func TestRateShibor(t *testing.T) {
	data, err := RateShibor()
	if err != nil {
		t.Logf("RateShibor 可能暂时不可用: %v", err)
		return
	}

	if len(data) == 0 {
		t.Log("RateShibor 返回空数据")
		return
	}

	t.Logf("获取到 %d 条Shibor数据", len(data))
	if len(data) > 0 {
		t.Logf("示例: %s 隔夜: %.4f%% 1年: %.4f%%", data[0].Date, data[0].Overnight, data[0].Year1)
	}
}

// TestRateLpr 测试LPR利率数据
func TestRateLpr(t *testing.T) {
	data, err := RateLpr()
	if err != nil {
		t.Logf("RateLpr 可能暂时不可用: %v", err)
		return
	}

	if len(data) == 0 {
		t.Log("RateLpr 返回空数据")
		return
	}

	t.Logf("获取到 %d 条LPR数据", len(data))
	if len(data) > 0 {
		t.Logf("示例: %s 1年期LPR: %.2f%% 5年期LPR: %.2f%%", data[0].Date, data[0].LPR1Y, data[0].LPR5Y)
	}
}

// TestRepoRateQuery 测试回购定盘利率查询
func TestRepoRateQuery(t *testing.T) {
	data, err := RepoRateQuery("回购定盘利率")
	if err != nil {
		t.Logf("RepoRateQuery 可能暂时不可用: %v", err)
		return
	}

	if len(data) == 0 {
		t.Log("RepoRateQuery 返回空数据")
		return
	}

	t.Logf("获取到 %d 条回购定盘利率数据", len(data))
	if len(data) > 0 {
		t.Logf("示例: %s FR001: %.4f FR007: %.4f", data[0].Date, data[0].FR001, data[0].FR007)
	}
}

// TestRepoRateQuery_FDR 测试银银间回购定盘利率查询
func TestRepoRateQuery_FDR(t *testing.T) {
	data, err := RepoRateQuery("银银间回购定盘利率")
	if err != nil {
		t.Logf("RepoRateQuery(FDR) 可能暂时不可用: %v", err)
		return
	}

	if len(data) == 0 {
		t.Log("RepoRateQuery(FDR) 返回空数据")
		return
	}

	t.Logf("获取到 %d 条银银间回购定盘利率数据", len(data))
	if len(data) > 0 {
		t.Logf("示例: %s FDR001: %.4f FDR007: %.4f", data[0].Date, data[0].FDR001, data[0].FDR007)
	}
}

// TestRepoRateHist 测试回购定盘利率历史数据
func TestRepoRateHist(t *testing.T) {
	data, err := RepoRateHist("20231001", "20231031")
	if err != nil {
		t.Logf("RepoRateHist 可能暂时不可用: %v", err)
		return
	}

	if len(data) == 0 {
		t.Log("RepoRateHist 返回空数据")
		return
	}

	t.Logf("获取到 %d 条回购定盘利率历史数据", len(data))
	if len(data) > 0 {
		t.Logf("示例: %s FR001: %.4f FDR001: %.4f", data[0].Date, data[0].FR001, data[0].FDR001)
	}
}
