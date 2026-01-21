package stock_fundamental

import (
	"fmt"
	"testing"
)

// TestStockZyjsThs 测试同花顺主营介绍接口
func TestStockZyjsThs(t *testing.T) {
	zyjs, err := StockZyjsThs("000066")
	if err != nil {
		t.Fatalf("获取同花顺主营介绍失败: %v", err)
	}

	if len(zyjs) == 0 {
		t.Skip("没有主营介绍数据")
	}

	// 验证股票代码存在
	if zyjs["股票代码"] != "000066" {
		t.Errorf("期望股票代码为000066，实际为%s", zyjs["股票代码"])
	}

	// 打印示例数据
	fmt.Println("同花顺主营介绍示例:")
	for key, value := range zyjs {
		fmt.Printf("  %s: %s\n", key, value)
	}
}

// TestStockProfitForecastThs 测试同花顺盈利预测接口
func TestStockProfitForecastThs(t *testing.T) {
	tests := []struct {
		name      string
		symbol    string
		indicator string
	}{
		{"预测年报每股收益", "600519", "预测年报每股收益"},
		{"预测年报净利润", "600519", "预测年报净利润"},
		{"业绩预测详表-机构", "600519", "业绩预测详表-机构"},
		{"业绩预测详表-详细指标预测", "600519", "业绩预测详表-详细指标预测"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			forecast, err := StockProfitForecastThs(tt.symbol, tt.indicator)
			if err != nil {
				t.Fatalf("获取%s失败: %v", tt.name, err)
			}

			fmt.Printf("\n%s示例:\n", tt.name)
			if len(forecast) == 0 {
				fmt.Printf("  暂无数据\n")
				return
			}

			// 打印前3条数据
			for i, item := range forecast {
				if i >= 3 {
					break
				}
				fmt.Printf("  行%d: %v\n", i+1, item)
			}
		})
	}
}

// TestStockProfitForecastThsInvalidIndicator 测试无效的indicator参数
func TestStockProfitForecastThsInvalidIndicator(t *testing.T) {
	_, err := StockProfitForecastThs("600519", "invalid_indicator")
	if err == nil {
		t.Error("期望返回错误，但实际没有")
	}
}

// TestStockHkProfitForecastEt 测试经济通港股盈利预测接口
func TestStockHkProfitForecastEt(t *testing.T) {
	tests := []struct {
		name      string
		symbol    string
		indicator string
	}{
		{"评级总览", "09999", "评级总览"},
		{"去年度业绩表现", "09999", "去年度业绩表现"},
		{"综合盈利预测", "09999", "综合盈利预测"},
		{"盈利预测概览", "09999", "盈利预测概览"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			forecast, err := StockHkProfitForecastEt(tt.symbol, tt.indicator)
			if err != nil {
				t.Fatalf("获取%s失败: %v", tt.name, err)
			}

			fmt.Printf("\n%s示例:\n", tt.name)
			if len(forecast) == 0 {
				fmt.Printf("  暂无数据\n")
				return
			}

			// 打印前3条数据
			for i, item := range forecast {
				if i >= 3 {
					break
				}
				fmt.Printf("  行%d: %v\n", i+1, item)
			}
		})
	}
}

// TestStockHkProfitForecastEtInvalidIndicator 测试无效的indicator参数
func TestStockHkProfitForecastEtInvalidIndicator(t *testing.T) {
	_, err := StockHkProfitForecastEt("09999", "invalid_indicator")
	if err == nil {
		t.Error("期望返回错误，但实际没有")
	}
}

// TestStockIndividualBasicInfoXq 测试雪球A股公司简介接口
func TestStockIndividualBasicInfoXq(t *testing.T) {
	info, err := StockIndividualBasicInfoXq("SH601127", "")
	if err != nil {
		t.Fatalf("获取雪球A股公司简介失败: %v", err)
	}

	if len(info) == 0 {
		t.Skip("没有公司简介数据")
	}

	// 打印示例数据
	fmt.Println("\n雪球A股公司简介示例:")
	for i, item := range info {
		if i >= 10 {
			break
		}
		fmt.Printf("  %s: %s\n", item.Item, item.Value)
	}
}

// TestStockIndividualBasicInfoUsXq 测试雪球美股公司简介接口
func TestStockIndividualBasicInfoUsXq(t *testing.T) {
	info, err := StockIndividualBasicInfoUsXq("NVDA", "")
	if err != nil {
		t.Fatalf("获取雪球美股公司简介失败: %v", err)
	}

	if len(info) == 0 {
		t.Skip("没有公司简介数据")
	}

	// 打印示例数据
	fmt.Println("\n雪球美股公司简介示例:")
	for i, item := range info {
		if i >= 10 {
			break
		}
		fmt.Printf("  %s: %s\n", item.Item, item.Value)
	}
}

// TestStockIndividualBasicInfoHkXq 测试雪球港股公司简介接口
func TestStockIndividualBasicInfoHkXq(t *testing.T) {
	info, err := StockIndividualBasicInfoHkXq("02097", "")
	if err != nil {
		t.Fatalf("获取雪球港股公司简介失败: %v", err)
	}

	if len(info) == 0 {
		t.Skip("没有公司简介数据")
	}

	// 打印示例数据
	fmt.Println("\n雪球港股公司简介示例:")
	for i, item := range info {
		if i >= 10 {
			break
		}
		fmt.Printf("  %s: %s\n", item.Item, item.Value)
	}
}

// TestStockIndividualBasicInfoXqEmptySymbol 测试空股票代码
func TestStockIndividualBasicInfoXqEmptySymbol(t *testing.T) {
	_, err := StockIndividualBasicInfoXq("", "")
	if err == nil {
		t.Error("期望返回错误，但实际没有")
	}
}

// TestStockIndividualBasicInfoUsXqEmptySymbol 测试空股票代码
func TestStockIndividualBasicInfoUsXqEmptySymbol(t *testing.T) {
	_, err := StockIndividualBasicInfoUsXq("", "")
	if err == nil {
		t.Error("期望返回错误，但实际没有")
	}
}
