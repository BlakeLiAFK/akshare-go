package bond

import (
	"testing"
)

// TestBondCBJSL 测试集思录可转债列表接口
func TestBondCBJSL(t *testing.T) {
	df, err := BondCBJSL("")
	if err != nil {
		t.Fatalf("BondCBJSL 失败: %v", err)
	}

	// 检查 DataFrame 是否为空
	if df.Nrow() == 0 {
		t.Error("BondCBJSL 返回空数据")
	}

	// 检查列数
	expectedCols := 23 // 根据实现,应该有23列
	if df.Ncol() != expectedCols {
		t.Errorf("BondCBJSL 列数不正确，期望 %d，实际 %d", expectedCols, df.Ncol())
	}

	// 检查关键列是否存在
	cols := df.Names()
	keyColumns := []string{"代码", "转债名称", "现价", "正股代码", "转股价", "转股价值", "转股溢价率"}
	for _, col := range keyColumns {
		found := false
		for _, c := range cols {
			if c == col {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("BondCBJSL 缺少关键列: %s", col)
		}
	}

	t.Logf("BondCBJSL 成功获取 %d 行数据, %d 列", df.Nrow(), df.Ncol())
	t.Logf("列名: %v", df.Names())

	// 打印前3行数据作为示例
	if df.Nrow() > 0 {
		t.Logf("前3行数据示例:\n%s", df.Subset([]int{0, 1, 2}))
	}
}

// TestBondCBRedeemJSL 测试集思录可转债强赎数据接口
func TestBondCBRedeemJSL(t *testing.T) {
	df, err := BondCBRedeemJSL()
	if err != nil {
		t.Fatalf("BondCBRedeemJSL 失败: %v", err)
	}

	// 检查 DataFrame 是否为空
	if df.Nrow() == 0 {
		t.Error("BondCBRedeemJSL 返回空数据")
	}

	// 检查列数
	expectedCols := 18 // 根据实现,应该有18列
	if df.Ncol() != expectedCols {
		t.Errorf("BondCBRedeemJSL 列数不正确，期望 %d，实际 %d", expectedCols, df.Ncol())
	}

	// 检查关键列是否存在
	cols := df.Names()
	keyColumns := []string{"代码", "名称", "强赎触发价", "强赎状态", "强赎天计数"}
	for _, col := range keyColumns {
		found := false
		for _, c := range cols {
			if c == col {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("BondCBRedeemJSL 缺少关键列: %s", col)
		}
	}

	t.Logf("BondCBRedeemJSL 成功获取 %d 行数据, %d 列", df.Nrow(), df.Ncol())
	t.Logf("列名: %v", df.Names())

	// 打印前3行数据作为示例
	if df.Nrow() > 0 {
		t.Logf("前3行数据示例:\n%s", df.Subset([]int{0, 1, 2}))
	}
}

// TestBondCBAdjLogsJSL 测试集思录可转债转股价调整记录接口
func TestBondCBAdjLogsJSL(t *testing.T) {
	// 测试有调整记录的可转债
	symbol := "128013" // 洪涛转债
	df, err := BondCBAdjLogsJSL(symbol)
	if err != nil {
		t.Fatalf("BondCBAdjLogsJSL 失败: %v", err)
	}

	// 注意：这个可转债可能没有调整记录，所以 df.Nrow() 可能为 0
	// 这是正常情况，不应该报错
	t.Logf("BondCBAdjLogsJSL(%s) 获取 %d 行数据", symbol, df.Nrow())

	if df.Nrow() > 0 {
		t.Logf("列名: %v", df.Names())
		t.Logf("数据示例:\n%s", df)

		// 检查关键列是否存在
		cols := df.Names()
		keyColumns := []string{"下修前转股价", "下修后转股价", "股东大会日", "新转股价生效日期"}
		for _, col := range keyColumns {
			found := false
			for _, c := range cols {
				if c == col {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("BondCBAdjLogsJSL 缺少关键列: %s", col)
			}
		}
	}
}

// TestBondCBAdjLogsJSL_NoData 测试没有调整记录的情况
func TestBondCBAdjLogsJSL_NoData(t *testing.T) {
	// 测试一个可能没有调整记录的可转债
	symbol := "123456" // 不存在的代码
	df, err := BondCBAdjLogsJSL(symbol)
	if err != nil {
		t.Fatalf("BondCBAdjLogsJSL 失败: %v", err)
	}

	// 应该返回空 DataFrame，不应该报错
	if df.Nrow() != 0 {
		t.Logf("BondCBAdjLogsJSL(%s) 意外获取到 %d 行数据", symbol, df.Nrow())
	} else {
		t.Logf("BondCBAdjLogsJSL(%s) 正确返回空数据", symbol)
	}
}

// TestBondCBAdjLogsJSL_EmptySymbol 测试空代码的情况
func TestBondCBAdjLogsJSL_EmptySymbol(t *testing.T) {
	_, err := BondCBAdjLogsJSL("")
	if err == nil {
		t.Error("BondCBAdjLogsJSL 应该对空代码返回错误")
	} else {
		t.Logf("BondCBAdjLogsJSL 正确处理空代码: %v", err)
	}
}
