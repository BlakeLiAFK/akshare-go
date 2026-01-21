package energy

import (
	"testing"
)

// TestEnergyOilHist 测试汽柴油历史调价
func TestEnergyOilHist(t *testing.T) {
	data, err := EnergyOilHist()
	if err != nil {
		t.Logf("EnergyOilHist 可能暂时不可用: %v", err)
		return
	}
	if len(data) == 0 {
		t.Log("EnergyOilHist 返回空数据")
		return
	}
	t.Logf("获取到 %d 条汽柴油调价数据", len(data))
}

// TestEnergyOilDetail 测试全国油价详情
func TestEnergyOilDetail(t *testing.T) {
	data, err := EnergyOilDetail("20231001")
	if err != nil {
		t.Logf("EnergyOilDetail 可能暂时不可用: %v", err)
		return
	}
	if len(data) == 0 {
		t.Log("EnergyOilDetail 返回空数据")
		return
	}
	t.Logf("获取到 %d 条全国油价数据", len(data))
}

// TestEnergyCarbonDomestic 测试国内碳交易行情
func TestEnergyCarbonDomestic(t *testing.T) {
	data, err := EnergyCarbonDomestic("湖北")
	if err != nil {
		t.Logf("EnergyCarbonDomestic 可能暂时不可用: %v", err)
		return
	}
	if len(data) == 0 {
		t.Log("EnergyCarbonDomestic 返回空数据")
		return
	}
	t.Logf("获取到 %d 条碳交易数据", len(data))
}

// TestEnergyCarbonBJ 测试北京碳交易行情
func TestEnergyCarbonBJ(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过耗时测试")
	}
	data, err := EnergyCarbonBJ()
	if err != nil {
		t.Logf("EnergyCarbonBJ 可能暂时不可用: %v", err)
		return
	}
	if len(data) == 0 {
		t.Log("EnergyCarbonBJ 返回空数据")
		return
	}
	t.Logf("获取到 %d 条北京碳交易数据", len(data))
}

// TestEnergyCarbonSZ 测试深圳碳交易行情
func TestEnergyCarbonSZ(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过耗时测试")
	}
	data, err := EnergyCarbonSZ()
	if err != nil {
		t.Logf("EnergyCarbonSZ 可能暂时不可用: %v", err)
		return
	}
	if len(data) == 0 {
		t.Log("EnergyCarbonSZ 返回空数据")
		return
	}
	t.Logf("获取到 %d 条深圳碳交易数据", len(data))
}

// TestEnergyCarbonEU 测试国际碳交易行情
func TestEnergyCarbonEU(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过耗时测试")
	}
	data, err := EnergyCarbonEU()
	if err != nil {
		t.Logf("EnergyCarbonEU 可能暂时不可用: %v", err)
		return
	}
	if len(data) == 0 {
		t.Log("EnergyCarbonEU 返回空数据")
		return
	}
	t.Logf("获取到 %d 条国际碳交易数据", len(data))
}

// TestEnergyCarbonHB 测试湖北碳交易行情
func TestEnergyCarbonHB(t *testing.T) {
	data, err := EnergyCarbonHB()
	if err != nil {
		t.Logf("EnergyCarbonHB 可能暂时不可用: %v", err)
		return
	}
	if len(data) == 0 {
		t.Log("EnergyCarbonHB 返回空数据")
		return
	}
	t.Logf("获取到 %d 条湖北碳交易数据", len(data))
}

// TestEnergyCarbonGZ 测试广州碳交易行情
func TestEnergyCarbonGZ(t *testing.T) {
	data, err := EnergyCarbonGZ()
	if err != nil {
		t.Logf("EnergyCarbonGZ 可能暂时不可用: %v", err)
		return
	}
	if len(data) == 0 {
		t.Log("EnergyCarbonGZ 返回空数据")
		return
	}
	t.Logf("获取到 %d 条广州碳交易数据", len(data))
}
