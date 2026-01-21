package cal

import (
	"strings"
	"testing"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

// TestRVFromStockZhAHistMinEM 测试股票分钟数据接口（依赖未实现）
func TestRVFromStockZhAHistMinEM(t *testing.T) {
	_, err := RVFromStockZhAHistMinEM("000001", "2021-10-20 09:30:00", "2024-11-01 15:00:00", "5", "")
	if err == nil {
		t.Error("期望返回错误，但未返回")
	}
	if !strings.Contains(err.Error(), "stock_feature") {
		t.Errorf("错误信息应包含依赖说明，实际: %v", err)
	}
	t.Logf("正确返回依赖错误: %v", err)
}

// TestRVFromFuturesZhMinuteSina 测试期货分钟数据接口（依赖未实现）
func TestRVFromFuturesZhMinuteSina(t *testing.T) {
	_, err := RVFromFuturesZhMinuteSina("IF2008", "5")
	if err == nil {
		t.Error("期望返回错误，但未返回")
	}
	if !strings.Contains(err.Error(), "futures") {
		t.Errorf("错误信息应包含依赖说明，实际: %v", err)
	}
	t.Logf("正确返回依赖错误: %v", err)
}

// TestVolatilityYZRV 测试 Yang-Zhang 已实现波动率计算
func TestVolatilityYZRV(t *testing.T) {
	// 构造测试数据
	// 模拟连续5天的分钟级数据，每天10个数据点
	dates := []string{
		"2024-01-01 09:30:00", "2024-01-01 10:00:00", "2024-01-01 10:30:00",
		"2024-01-01 11:00:00", "2024-01-01 11:30:00", "2024-01-01 13:00:00",
		"2024-01-01 13:30:00", "2024-01-01 14:00:00", "2024-01-01 14:30:00",
		"2024-01-01 15:00:00",

		"2024-01-02 09:30:00", "2024-01-02 10:00:00", "2024-01-02 10:30:00",
		"2024-01-02 11:00:00", "2024-01-02 11:30:00", "2024-01-02 13:00:00",
		"2024-01-02 13:30:00", "2024-01-02 14:00:00", "2024-01-02 14:30:00",
		"2024-01-02 15:00:00",

		"2024-01-03 09:30:00", "2024-01-03 10:00:00", "2024-01-03 10:30:00",
		"2024-01-03 11:00:00", "2024-01-03 11:30:00", "2024-01-03 13:00:00",
		"2024-01-03 13:30:00", "2024-01-03 14:00:00", "2024-01-03 14:30:00",
		"2024-01-03 15:00:00",
	}

	opens := []float64{
		100.0, 101.0, 102.0, 101.5, 102.5, 103.0, 102.8, 103.5, 104.0, 103.8,
		103.5, 104.0, 105.0, 104.5, 105.5, 106.0, 105.8, 106.5, 107.0, 106.8,
		106.5, 107.0, 108.0, 107.5, 108.5, 109.0, 108.8, 109.5, 110.0, 109.8,
	}

	highs := []float64{
		101.5, 102.5, 103.0, 103.0, 104.0, 104.5, 104.0, 105.0, 105.5, 105.0,
		105.0, 106.0, 106.5, 106.0, 107.0, 107.5, 107.0, 108.0, 108.5, 108.0,
		108.0, 109.0, 109.5, 109.0, 110.0, 110.5, 110.0, 111.0, 111.5, 111.0,
	}

	lows := []float64{
		99.5, 100.5, 101.5, 101.0, 102.0, 102.5, 102.3, 103.0, 103.5, 103.3,
		103.0, 103.5, 104.5, 104.0, 105.0, 105.5, 105.3, 106.0, 106.5, 106.3,
		106.0, 106.5, 107.5, 107.0, 108.0, 108.5, 108.3, 109.0, 109.5, 109.3,
	}

	closes := []float64{
		101.0, 102.0, 101.5, 102.5, 103.0, 102.8, 103.5, 104.0, 103.8, 104.5,
		104.0, 105.0, 104.5, 105.5, 106.0, 105.8, 106.5, 107.0, 106.8, 107.5,
		107.0, 108.0, 107.5, 108.5, 109.0, 108.8, 109.5, 110.0, 109.8, 110.5,
	}

	// 构建 DataFrame
	df := dataframe.New(
		series.New(dates, series.String, "Date"),
		series.New(opens, series.Float, "Open"),
		series.New(highs, series.Float, "High"),
		series.New(lows, series.Float, "Low"),
		series.New(closes, series.Float, "Close"),
	)

	// 调用函数
	result, err := VolatilityYZRV(df)
	if err != nil {
		t.Fatalf("VolatilityYZRV 失败: %v", err)
	}

	// 验证结果
	rows, _ := result.Dims()
	if rows == 0 {
		t.Error("结果为空")
	}

	// 检查列名
	colNames := result.Names()
	if len(colNames) != 2 || colNames[0] != "date" || colNames[1] != "rv" {
		t.Errorf("列名不正确，期望 [date, rv]，实际: %v", colNames)
	}

	// 检查数值是否为正数
	rvCol := result.Col("rv").Float()
	for i, rv := range rvCol {
		if rv < 0 {
			t.Errorf("第 %d 行的 rv 值为负数: %f", i, rv)
		}
	}

	t.Logf("Yang-Zhang RV 计算结果行数: %d", rows)
	t.Logf("前3个结果:")
	for i := 0; i < min(3, rows); i++ {
		date := result.Elem(i, 0).String()
		rv := result.Elem(i, 1).Float()
		t.Logf("  日期: %s, RV: %.6f", date, rv)
	}
}

// TestVolatilityYZRV_EmptyData 测试空数据
func TestVolatilityYZRV_EmptyData(t *testing.T) {
	df := dataframe.New(
		series.New([]string{}, series.String, "Date"),
		series.New([]float64{}, series.Float, "Open"),
		series.New([]float64{}, series.Float, "High"),
		series.New([]float64{}, series.Float, "Low"),
		series.New([]float64{}, series.Float, "Close"),
	)

	_, err := VolatilityYZRV(df)
	if err == nil {
		t.Error("期望返回错误（数据不足），但未返回")
	}
	t.Logf("正确处理空数据: %v", err)
}

// TestVolatilityYZRV_MissingColumn 测试缺少列
func TestVolatilityYZRV_MissingColumn(t *testing.T) {
	df := dataframe.New(
		series.New([]string{"2024-01-01"}, series.String, "Date"),
		series.New([]float64{100.0}, series.Float, "Open"),
		series.New([]float64{101.0}, series.Float, "High"),
		// 缺少 Low 列
		series.New([]float64{100.5}, series.Float, "Close"),
	)

	_, err := VolatilityYZRV(df)
	if err == nil {
		t.Error("期望返回错误（缺少列），但未返回")
	}
	if !strings.Contains(err.Error(), "缺少必需列") {
		t.Errorf("错误信息不正确，期望包含'缺少必需列'，实际: %v", err)
	}
	t.Logf("正确处理缺少列: %v", err)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
