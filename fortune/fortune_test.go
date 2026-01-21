package fortune

import (
	"testing"
)

// TestFortuneRank 测试财富500强
func TestFortuneRank(t *testing.T) {
	data, err := FortuneRank("2023")
	if err != nil {
		t.Logf("FortuneRank 可能暂时不可用: %v", err)
		return
	}
	if len(data) == 0 {
		t.Log("FortuneRank 返回空数据")
		return
	}
	t.Logf("获取到 %d 条财富500强数据", len(data))
	if len(data) > 0 {
		t.Logf("示例: %s %s %s", data[0].Rank, data[0].Company, data[0].Revenue)
	}
}

// TestForbesRank 测试福布斯富豪榜
func TestForbesRank(t *testing.T) {
	data, err := ForbesRank("billionaires")
	if err != nil {
		t.Logf("ForbesRank 可能暂时不可用: %v", err)
		return
	}
	if len(data) == 0 {
		t.Log("ForbesRank 返回空数据")
		return
	}
	t.Logf("获取到 %d 条福布斯数据", len(data))
}

// TestHurunRank 测试胡润富豪榜
func TestHurunRank(t *testing.T) {
	data, err := HurunRank("hurun", "2023")
	if err != nil {
		t.Logf("HurunRank 可能暂时不可用: %v", err)
		return
	}
	if len(data) == 0 {
		t.Log("HurunRank 返回空数据")
		return
	}
	t.Logf("获取到 %d 条胡润富豪榜数据", len(data))
}

// TestXincaifuRank 测试新财富500富人榜
func TestXincaifuRank(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过耗时测试")
	}
	data, err := XincaifuRank("2023")
	if err != nil {
		t.Logf("XincaifuRank 可能暂时不可用: %v", err)
		return
	}
	if len(data) == 0 {
		t.Log("XincaifuRank 返回空数据")
		return
	}
	t.Logf("获取到 %d 条新财富富人榜数据", len(data))
}

// TestIndexBloombergBillionaires 测试彭博亿万富翁指数
func TestIndexBloombergBillionaires(t *testing.T) {
	data, err := IndexBloombergBillionaires()
	if err != nil {
		t.Logf("IndexBloombergBillionaires 可能暂时不可用: %v", err)
		return
	}
	if len(data) == 0 {
		t.Log("IndexBloombergBillionaires 返回空数据")
		return
	}
	t.Logf("获取到 %d 条彭博亿万富翁数据", len(data))
}

// TestIndexBloombergBillionairesHist 测试彭博亿万富翁历史数据
func TestIndexBloombergBillionairesHist(t *testing.T) {
	data, err := IndexBloombergBillionairesHist("2023")
	if err != nil {
		t.Logf("IndexBloombergBillionairesHist 可能暂时不可用: %v", err)
		return
	}
	if len(data) == 0 {
		t.Log("IndexBloombergBillionairesHist 返回空数据")
		return
	}
	t.Logf("获取到 %d 条彭博亿万富翁历史数据", len(data))
}
