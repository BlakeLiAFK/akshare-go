package stock_fundamental

import (
	"fmt"
	"testing"
)

// TestStockProfitForecastEm 测试东方财富盈利预测接口
func TestStockProfitForecastEm(t *testing.T) {
	// 测试获取全部数据
	forecast, err := StockProfitForecastEm("")
	if err != nil {
		t.Fatalf("获取盈利预测失败: %v", err)
	}

	if len(forecast) == 0 {
		t.Skip("没有盈利预测数据")
	}

	// 验证第一条数据
	first := forecast[0]
	if first.Index != 1 {
		t.Errorf("期望序号为1，实际为%d", first.Index)
	}
	if first.Code == "" {
		t.Error("代码不能为空")
	}
	if first.Name == "" {
		t.Error("名称不能为空")
	}

	// 打印示例
	fmt.Printf("盈利预测示例: %s (%s), 研报数: %d, 买入评级: %d\n",
		first.Name, first.Code, first.ReportCount, first.RatingBuy)
}

// TestStockRegisterAllEm 测试IPO审核信息-全部接口
func TestStockRegisterAllEm(t *testing.T) {
	register, err := StockRegisterAllEm()
	if err != nil {
		t.Fatalf("获取IPO审核信息失败: %v", err)
	}

	if len(register) == 0 {
		t.Skip("没有IPO审核数据")
	}

	// 验证第一条数据
	first := register[0]
	if first.Index != 1 {
		t.Errorf("期望序号为1，实际为%d", first.Index)
	}
	if first.CompanyName == "" {
		t.Error("企业名称不能为空")
	}
	if first.State == "" {
		t.Error("最新状态不能为空")
	}

	// 打印示例
	fmt.Printf("IPO审核示例: %s, 状态: %s, 保荐机构: %s\n",
		first.CompanyName, first.State, first.RecommendOrg)
}

// TestStockRegisterKcb 测试IPO审核信息-科创板接口
func TestStockRegisterKcb(t *testing.T) {
	register, err := StockRegisterKcb()
	if err != nil {
		t.Fatalf("获取科创板IPO审核信息失败: %v", err)
	}

	if len(register) == 0 {
		t.Skip("没有科创板IPO审核数据")
	}

	// 验证拟上市地点
	first := register[0]
	if first.PredictMarket != "科创板" {
		t.Errorf("期望拟上市地点为科创板，实际为%s", first.PredictMarket)
	}

	fmt.Printf("科创板IPO审核示例: %s, 状态: %s\n", first.CompanyName, first.State)
}

// TestStockRegisterCyb 测试IPO审核信息-创业板接口
func TestStockRegisterCyb(t *testing.T) {
	register, err := StockRegisterCyb()
	if err != nil {
		t.Fatalf("获取创业板IPO审核信息失败: %v", err)
	}

	if len(register) == 0 {
		t.Skip("没有创业板IPO审核数据")
	}

	// 验证拟上市地点
	first := register[0]
	if first.PredictMarket != "创业板" {
		t.Errorf("期望拟上市地点为创业板，实际为%s", first.PredictMarket)
	}

	fmt.Printf("创业板IPO审核示例: %s, 状态: %s\n", first.CompanyName, first.State)
}

// TestStockRestrictedReleaseSummaryEm 测试限售股解禁汇总接口
func TestStockRestrictedReleaseSummaryEm(t *testing.T) {
	summary, err := StockRestrictedReleaseSummaryEm("全部股票", "20240101", "20240131")
	if err != nil {
		t.Fatalf("获取限售股解禁汇总失败: %v", err)
	}

	if len(summary) == 0 {
		t.Skip("该时间段没有限售股解禁数据")
	}

	// 验证第一条数据
	first := summary[0]
	if first.Index != 1 {
		t.Errorf("期望序号为1，实际为%d", first.Index)
	}

	// 打印示例
	fmt.Printf("限售股解禁汇总示例: %s, 解禁家数: %d, 实际解禁市值: %.2f\n",
		first.LiftDate.Format("2006-01-02"), first.StockCount, first.ActualLiftMarketCap)
}

// TestStockRestrictedReleaseDetailEm 测试限售股解禁详情接口
func TestStockRestrictedReleaseDetailEm(t *testing.T) {
	detail, err := StockRestrictedReleaseDetailEm("20240101", "20240110")
	if err != nil {
		t.Fatalf("获取限售股解禁详情失败: %v", err)
	}

	if len(detail) == 0 {
		t.Skip("该时间段没有限售股解禁详情数据")
	}

	// 验证第一条数据
	first := detail[0]
	if first.Index != 1 {
		t.Errorf("期望序号为1，实际为%d", first.Index)
	}
	if first.Code == "" {
		t.Error("股票代码不能为空")
	}

	// 打印示例
	fmt.Printf("限售股解禁详情示例: %s (%s), 解禁时间: %s, 解禁数量: %.2f\n",
		first.Name, first.Code, first.LiftDate.Format("2006-01-02"), first.LiftShares)
}

// TestStockRestrictedReleaseQueueEm 测试个股限售解禁批次接口
func TestStockRestrictedReleaseQueueEm(t *testing.T) {
	queue, err := StockRestrictedReleaseQueueEm("600000")
	if err != nil {
		t.Fatalf("获取个股限售解禁批次失败: %v", err)
	}

	if len(queue) == 0 {
		t.Skip("该股票没有限售解禁数据")
	}

	// 验证第一条数据
	first := queue[0]
	if first.Index != 1 {
		t.Errorf("期望序号为1，实际为%d", first.Index)
	}

	// 打印示例
	fmt.Printf("个股限售解禁批次示例: %s, 解禁股东数: %d, 实际解禁市值: %.2f\n",
		first.LiftDate.Format("2006-01-02"), first.HolderNum, first.ActualLiftMarketCap)
}

// TestGetSupportedSymbolList 测试获取支持的symbol列表
func TestGetSupportedSymbolList(t *testing.T) {
	symbols := GetSupportedSymbolList()

	if len(symbols) == 0 {
		t.Error("支持的symbol列表不能为空")
	}

	// 验证包含基本的symbol
	expectedSymbols := []string{"全部股票", "沪市A股", "科创板"}
	for _, expected := range expectedSymbols {
		found := false
		for _, s := range symbols {
			if s == expected {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("期望的symbol %s 不在列表中", expected)
		}
	}

	fmt.Printf("支持的symbol列表: %v\n", symbols)
}

// TestStockRestrictedReleaseSummaryEmInvalidSymbol 测试无效的symbol参数
func TestStockRestrictedReleaseSummaryEmInvalidSymbol(t *testing.T) {
	_, err := StockRestrictedReleaseSummaryEm("invalid_symbol", "20240101", "20240131")
	if err == nil {
		t.Error("期望返回错误，但实际没有")
	}
}
