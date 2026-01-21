package interest_rate

import (
	"testing"
)

// TestRateInterbank 测试银行间同业拆借利率
func TestRateInterbank(t *testing.T) {
	data, err := RateInterbank("上海银行间同业拆放利率", "Shibor人民币", "利率")
	if err != nil {
		t.Logf("RateInterbank 可能暂时不可用: %v", err)
		return
	}
	if len(data) == 0 {
		t.Log("RateInterbank 返回空数据")
		return
	}
	t.Logf("获取到 %d 条同业拆借利率数据", len(data))
	if len(data) > 0 {
		t.Logf("示例: %s %.4f%% %.4f", data[0].ReportDate, data[0].Rate, data[0].Change)
	}
}

// TestRateInterbank_Libor 测试Libor利率
func TestRateInterbank_Libor(t *testing.T) {
	data, err := RateInterbank("Libor人民币", "", "利率")
	if err != nil {
		t.Logf("RateInterbank Libor 可能暂时不可用: %v", err)
		return
	}
	if len(data) == 0 {
		t.Log("RateInterbank Libor 返回空数据")
		return
	}
	t.Logf("获取到 %d 条Libor利率数据", len(data))
}

// TestRateInterbank_Hibor 测试Hibor利率
func TestRateInterbank_Hibor(t *testing.T) {
	data, err := RateInterbank("香港银行同业拆借利率", "Hibor港币", "利率")
	if err != nil {
		t.Logf("RateInterbank Hibor 可能暂时不可用: %v", err)
		return
	}
	if len(data) == 0 {
		t.Log("RateInterbank Hibor 返回空数据")
		return
	}
	t.Logf("获取到 %d 条Hibor利率数据", len(data))
}
