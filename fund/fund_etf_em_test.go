package fund

import (
	"testing"
)

func TestFundEtfSpotEm(t *testing.T) {
	records, err := FundEtfSpotEm()
	if err != nil {
		t.Fatalf("FundEtfSpotEm failed: %v", err)
	}

	// 允许返回空数据（可能是市场休市）
	t.Logf("FundEtfSpotEm returned %d records", len(records))

	if len(records) > 0 {
		first := records[0]
		if len(first) == 0 {
			t.Error("First record is empty")
		}
	}
}

func TestFundEtfHistEm(t *testing.T) {
	// 测试基本调用
	records, err := FundEtfHistEm("510300", "daily", "", "", "")
	if err != nil {
		t.Fatalf("FundEtfHistEm failed: %v", err)
	}

	// 允许返回空数据
	t.Logf("FundEtfHistEm returned %d records for 510300", len(records))

	if len(records) > 0 {
		first := records[0]
		// 验证有一些字段存在
		if len(first) == 0 {
			t.Error("First record is empty")
		}
	}
}

func TestFundEtfHistMinEm(t *testing.T) {
	// 测试分钟线数据
	records, err := FundEtfHistMinEm("510300", "", "", "1", "")
	if err != nil {
		t.Fatalf("FundEtfHistMinEm failed: %v", err)
	}

	// 允许返回空数据（可能是非交易时间）
	t.Logf("FundEtfHistMinEm returned %d minute records", len(records))
}
