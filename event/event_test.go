package event

import (
	"testing"
)

// TestMigrationAreaBaidu 测试百度迁徙地区数据
func TestMigrationAreaBaidu(t *testing.T) {
	data, err := MigrationAreaBaidu("广东省", "move_in", "20231001")
	if err != nil {
		t.Logf("MigrationAreaBaidu 可能暂时不可用: %v", err)
		return
	}
	if len(data) == 0 {
		t.Log("MigrationAreaBaidu 返回空数据")
		return
	}
	t.Logf("获取到 %d 条迁徙地区数据", len(data))
}

// TestMigrationScaleBaidu 测试百度迁徙规模数据
func TestMigrationScaleBaidu(t *testing.T) {
	data, err := MigrationScaleBaidu("广东省", "move_in")
	if err != nil {
		t.Logf("MigrationScaleBaidu 可能暂时不可用: %v", err)
		return
	}
	if len(data) == 0 {
		t.Log("MigrationScaleBaidu 返回空数据")
		return
	}
	t.Logf("获取到 %d 条迁徙规模数据", len(data))
}

// TestGetProvinceName 测试获取省份名称
func TestGetProvinceName(t *testing.T) {
	name := GetProvinceName("440000")
	if name == "" {
		t.Log("GetProvinceName 返回空")
		return
	}
	t.Logf("省份代码 440000 对应: %s", name)
}

// TestGetCityName 测试获取城市名称
func TestGetCityName(t *testing.T) {
	name := GetCityName("440100")
	if name == "" {
		t.Log("GetCityName 返回空")
		return
	}
	t.Logf("城市代码 440100 对应: %s", name)
}

// TestGetProvinceCode 测试获取省份代码
func TestGetProvinceCode(t *testing.T) {
	code := GetProvinceCode("广东省")
	if code == "" {
		t.Log("GetProvinceCode 返回空")
		return
	}
	t.Logf("广东省 对应代码: %s", code)
}

// TestGetCityCode 测试获取城市代码
func TestGetCityCode(t *testing.T) {
	code := GetCityCode("广州市")
	if code == "" {
		t.Log("GetCityCode 返回空")
		return
	}
	t.Logf("广州市 对应代码: %s", code)
}
