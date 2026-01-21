package futures_derivative

import (
	"testing"
)

func TestFuturesHogCore(t *testing.T) {
	symbols := []string{"外三元", "内三元", "土杂猪"}

	for _, symbol := range symbols {
		t.Run(symbol, func(t *testing.T) {
			records, err := FuturesHogCore(symbol)
			if err != nil {
				t.Logf("FuturesHogCore(%s) warning: %v (may be unavailable)", symbol, err)
				return
			}

			// 允许空数据
			t.Logf("FuturesHogCore(%s) returned %d records", symbol, len(records))

			// 如果有数据，验证结构
			if len(records) > 0 {
				first := records[0]
				requiredFields := []string{"date", "value"}
				for _, field := range requiredFields {
					if _, ok := first[field]; !ok {
						t.Errorf("Missing required field: %s", field)
					}
				}
			}
		})
	}
}

func TestFuturesHogCost(t *testing.T) {
	symbols := []string{"玉米", "豆粕", "二元母猪价格", "仔猪价格"}

	for _, symbol := range symbols {
		t.Run(symbol, func(t *testing.T) {
			records, err := FuturesHogCost(symbol)
			if err != nil {
				t.Logf("FuturesHogCost(%s) warning: %v (may be unavailable)", symbol, err)
				return
			}

			// 允许空数据
			t.Logf("FuturesHogCost(%s) returned %d records", symbol, len(records))

			// 如果有数据，验证结构
			if len(records) > 0 {
				first := records[0]
				requiredFields := []string{"date", "value"}
				for _, field := range requiredFields {
					if _, ok := first[field]; !ok {
						t.Errorf("Missing required field: %s", field)
					}
				}
			}
		})
	}
}

func TestFuturesHogSupply(t *testing.T) {
	symbols := []string{
		"猪肉批发价", "储备冻猪肉", "饲料原料数据", "白条肉",
		"生猪产能", "育肥猪", "肉类价格指数", "猪粮比价",
	}

	for _, symbol := range symbols {
		t.Run(symbol, func(t *testing.T) {
			records, err := FuturesHogSupply(symbol)
			if err != nil {
				t.Logf("FuturesHogSupply(%s) warning: %v (may be unavailable)", symbol, err)
				return
			}

			// 允许空数据
			t.Logf("FuturesHogSupply(%s) returned %d records", symbol, len(records))

			// 如果有数据，验证结构（根据不同symbol有不同字段）
			if len(records) > 0 {
				first := records[0]

				switch symbol {
				case "猪肉批发价", "储备冻猪肉", "肉类价格指数", "猪粮比价":
					// 这些返回 [date, value]
					requiredFields := []string{"date", "value"}
					for _, field := range requiredFields {
						if _, ok := first[field]; !ok {
							t.Errorf("Missing required field: %s", field)
						}
					}
				case "饲料原料数据":
					// 返回 [周期, 大豆进口金额, 大豆播种面积, 玉米进口金额, 玉米播种面积]
					requiredFields := []string{"周期", "大豆进口金额", "大豆播种面积", "玉米进口金额", "玉米播种面积"}
					for _, field := range requiredFields {
						if _, ok := first[field]; !ok {
							t.Errorf("Missing required field: %s", field)
						}
					}
				case "白条肉":
					// 返回 [周期, 白条肉平均出厂价格, 环比, 同比]
					requiredFields := []string{"周期", "白条肉平均出厂价格", "环比", "同比"}
					for _, field := range requiredFields {
						if _, ok := first[field]; !ok {
							t.Errorf("Missing required field: %s", field)
						}
					}
				case "生猪产能":
					// 返回 [周期, 能繁母猪存栏, 猪肉产量, 生猪存栏, 生猪出栏]
					requiredFields := []string{"周期", "能繁母猪存栏", "猪肉产量", "生猪存栏", "生猪出栏"}
					for _, field := range requiredFields {
						if _, ok := first[field]; !ok {
							t.Errorf("Missing required field: %s", field)
						}
					}
				case "育肥猪":
					// 返回 [date, benzhou]
					requiredFields := []string{"date", "benzhou"}
					for _, field := range requiredFields {
						if _, ok := first[field]; !ok {
							t.Errorf("Missing required field: %s", field)
						}
					}
				}
			}
		})
	}
}
