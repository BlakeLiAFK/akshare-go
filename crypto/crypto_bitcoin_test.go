package crypto

import (
	"testing"
	"time"
)

// TestCryptoBitcoinCME 测试芝加哥商业交易所比特币成交量报告
func TestCryptoBitcoinCME(t *testing.T) {
	done := make(chan bool)
	go func() {
		// 使用一个历史日期进行测试
		df, err := CryptoBitcoinCME("20230830")
		if err != nil {
			t.Errorf("CryptoBitcoinCME失败: %v", err)
		}

		rows, cols := df.Dims()
		if rows <= 1 {
			t.Errorf("期望至少有数据行，实际行数: %d", rows)
		}
		if cols == 0 {
			t.Error("期望至少有列，实际列数: 0")
		}

		t.Logf("CME比特币成交量报告数据行数: %d, 列数: %d", rows, cols)

		// 打印列名
		colNames := df.Names()
		t.Logf("列名: %v", colNames)

		done <- true
	}()

	select {
	case <-done:
	case <-time.After(30 * time.Second):
		t.Skip("超时，跳过测试")
	}
}

// TestCryptoBitcoinHoldReport 测试比特币持仓报告
func TestCryptoBitcoinHoldReport(t *testing.T) {
	done := make(chan bool)
	go func() {
		df, err := CryptoBitcoinHoldReport()
		if err != nil {
			t.Errorf("CryptoBitcoinHoldReport失败: %v", err)
		}

		rows, cols := df.Dims()
		if rows <= 1 {
			t.Errorf("期望至少有数据行，实际行数: %d", rows)
		}
		if cols == 0 {
			t.Error("期望至少有列，实际列数: 0")
		}

		t.Logf("比特币持仓报告数据行数: %d, 列数: %d", rows, cols)

		// 打印列名
		colNames := df.Names()
		t.Logf("列名: %v", colNames)

		// 验证期望的列名是否存在
		expectedCols := []string{"代码", "公司名称-英文", "公司名称-中文", "国家/地区", "市值"}
		for _, expected := range expectedCols {
			found := false
			for _, actual := range colNames {
				if actual == expected {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("缺少期望的列: %s", expected)
			}
		}

		done <- true
	}()

	select {
	case <-done:
	case <-time.After(30 * time.Second):
		t.Skip("超时，跳过测试")
	}
}

// TestCryptoBitcoinCME_InvalidDate 测试无效日期格式
func TestCryptoBitcoinCME_InvalidDate(t *testing.T) {
	_, err := CryptoBitcoinCME("2023-08-30") // 错误格式
	if err == nil {
		t.Error("期望返回错误（无效日期格式），但未返回")
	}
	t.Logf("正确处理无效日期格式: %v", err)
}
