package stock

import (
	"testing"
	"time"
)

func TestStockZhAHist(t *testing.T) {
	tests := []struct {
		name      string
		code      string
		period    string
		startDate string
		endDate   string
		adjust    string
		wantErr   bool
	}{
		{
			name:      "获取平安银行日K线数据-不复权",
			code:      "000001",
			period:    "daily",
			startDate: "20240101",
			endDate:   "20240115",
			adjust:    "",
			wantErr:   false,
		},
		{
			name:      "获取平安银行日K线数据-前复权",
			code:      "000001",
			period:    "daily",
			startDate: "20240101",
			endDate:   "20240115",
			adjust:    "qfq",
			wantErr:   false,
		},
		{
			name:      "获取贵州茅台周K线数据-后复权",
			code:      "600519",
			period:    "weekly",
			startDate: "20240101",
			endDate:   "20240331",
			adjust:    "hfq",
			wantErr:   false,
		},
		{
			name:      "获取宁德时代月K线数据",
			code:      "300750",
			period:    "monthly",
			startDate: "20230101",
			endDate:   "20231231",
			adjust:    "",
			wantErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			klines, err := StockZhAHist(tt.code, tt.period, tt.startDate, tt.endDate, tt.adjust)
			if (err != nil) != tt.wantErr {
				t.Errorf("StockZhAHist() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				if len(klines) == 0 {
					t.Errorf("StockZhAHist() 返回数据为空")
					return
				}

				// 验证第一条数据的基本字段
				first := klines[0]
				if first.Date.IsZero() {
					t.Errorf("日期字段为空")
				}
				if first.Open <= 0 || first.Close <= 0 || first.High <= 0 || first.Low <= 0 {
					t.Errorf("价格字段异常: open=%.2f, close=%.2f, high=%.2f, low=%.2f",
						first.Open, first.Close, first.High, first.Low)
				}
				if first.Volume < 0 {
					t.Errorf("成交量异常: %d", first.Volume)
				}
				if first.Adjust != tt.adjust {
					t.Errorf("复权类型不匹配: got %s, want %s", first.Adjust, tt.adjust)
				}

				t.Logf("获取到 %d 条K线数据", len(klines))
				t.Logf("第一条数据: 日期=%s, 开盘=%.2f, 收盘=%.2f, 最高=%.2f, 最低=%.2f, 成交量=%d, 涨跌幅=%.2f%%",
					first.Date.Format("2006-01-02"),
					first.Open, first.Close, first.High, first.Low,
					first.Volume, first.ChangePct)
			}
		})
	}
}

func TestGetSecID(t *testing.T) {
	tests := []struct {
		name string
		code string
		want string
	}{
		{
			name: "沪市主板-6开头",
			code: "600519",
			want: "1.600519",
		},
		{
			name: "科创板-688开头",
			code: "688001",
			want: "1.688001",
		},
		{
			name: "深市主板-000开头",
			code: "000001",
			want: "0.000001",
		},
		{
			name: "创业板-300开头",
			code: "300750",
			want: "0.300750",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getSecID(tt.code)
			if got != tt.want {
				t.Errorf("getSecID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStockZhAHistDateRange(t *testing.T) {
	// 测试日期范围
	endDate := time.Now().Format("20060102")
	startDate := time.Now().AddDate(0, 0, -7).Format("20060102") // 最近7天

	klines, err := StockZhAHist("000001", "daily", startDate, endDate, "")
	if err != nil {
		t.Fatalf("获取最近7天数据失败: %v", err)
	}

	if len(klines) == 0 {
		t.Logf("最近7天无交易数据（可能是节假日）")
		return
	}

	// 验证日期是否在指定范围内
	startTime, _ := time.Parse("20060102", startDate)
	endTime, _ := time.Parse("20060102", endDate)

	for _, kline := range klines {
		if kline.Date.Before(startTime) || kline.Date.After(endTime) {
			t.Errorf("日期超出范围: %s (范围: %s ~ %s)",
				kline.Date.Format("2006-01-02"),
				startDate, endDate)
		}
	}

	t.Logf("日期范围测试通过，共 %d 条数据", len(klines))
}

// BenchmarkStockZhAHist 性能基准测试
func BenchmarkStockZhAHist(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := StockZhAHist("000001", "daily", "20240101", "20240115", "")
		if err != nil {
			b.Fatal(err)
		}
	}
}
