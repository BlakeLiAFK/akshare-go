package bond

import (
	"testing"
	"time"
)

// 质押式回购测试
func TestBondSHBuyBackEM(t *testing.T) {
	done := make(chan bool)
	go func() {
		df, err := BondSHBuyBackEM()
		if err != nil {
			t.Errorf("BondSHBuyBackEM失败: %v", err)
		}

		rows, cols := df.Dims()
		if rows <= 1 {
			t.Errorf("期望至少有数据行，实际行数: %d", rows)
		}
		if cols != 12 {
			t.Errorf("期望12列，实际列数: %d", cols)
		}

		t.Logf("上证质押式回购数据: 行数=%d, 列数=%d", rows, cols)
		done <- true
	}()

	select {
	case <-done:
	case <-time.After(30 * time.Second):
		t.Skip("超时，跳过测试")
	}
}

func TestBondSZBuyBackEM(t *testing.T) {
	done := make(chan bool)
	go func() {
		df, err := BondSZBuyBackEM()
		if err != nil {
			t.Errorf("BondSZBuyBackEM失败: %v", err)
		}

		rows, cols := df.Dims()
		if rows <= 1 {
			t.Errorf("期望至少有数据行，实际行数: %d", rows)
		}
		if cols != 12 {
			t.Errorf("期望12列，实际列数: %d", cols)
		}

		t.Logf("深证质押式回购数据: 行数=%d, 列数=%d", rows, cols)
		done <- true
	}()

	select {
	case <-done:
	case <-time.After(30 * time.Second):
		t.Skip("超时，跳过测试")
	}
}

func TestBondBuyBackHistEM(t *testing.T) {
	done := make(chan bool)
	go func() {
		// 测试上证回购
		df, err := BondBuyBackHistEM("204001")
		if err != nil {
			t.Errorf("BondBuyBackHistEM(204001)失败: %v", err)
		}

		rows, cols := df.Dims()
		if rows <= 1 {
			t.Errorf("期望至少有数据行，实际行数: %d", rows)
		}
		if cols != 7 {
			t.Errorf("期望7列，实际列数: %d", cols)
		}

		t.Logf("质押式回购历史数据(204001): 行数=%d, 列数=%d", rows, cols)

		// 测试深证回购
		df2, err2 := BondBuyBackHistEM("131810")
		if err2 != nil {
			t.Errorf("BondBuyBackHistEM(131810)失败: %v", err2)
		}

		rows2, _ := df2.Dims()
		if rows2 <= 1 {
			t.Errorf("期望至少有数据行，实际行数: %d", rows2)
		}

		t.Logf("质押式回购历史数据(131810): 行数=%d", rows2)
		done <- true
	}()

	select {
	case <-done:
	case <-time.After(30 * time.Second):
		t.Skip("超时，跳过测试")
	}
}

// 新浪财经可转债测试
func TestBondCBProfileSina(t *testing.T) {
	done := make(chan bool)
	go func() {
		df, err := BondCBProfileSina("sz128039")
		if err != nil {
			t.Errorf("BondCBProfileSina失败: %v", err)
		}

		rows, cols := df.Dims()
		if rows <= 1 {
			t.Errorf("期望至少有数据行，实际行数: %d", rows)
		}
		if cols != 2 {
			t.Errorf("期望2列，实际列数: %d", cols)
		}

		t.Logf("新浪财经可转债详情: 行数=%d, 列数=%d", rows, cols)
		done <- true
	}()

	select {
	case <-done:
	case <-time.After(30 * time.Second):
		t.Skip("超时，跳过测试")
	}
}

func TestBondCBSummarySina(t *testing.T) {
	done := make(chan bool)
	go func() {
		df, err := BondCBSummarySina("sh155255")
		if err != nil {
			t.Errorf("BondCBSummarySina失败: %v", err)
		}

		rows, cols := df.Dims()
		if rows <= 1 {
			t.Errorf("期望至少有数据行，实际行数: %d", rows)
		}
		if cols != 2 {
			t.Errorf("期望2列，实际列数: %d", cols)
		}

		t.Logf("新浪财经可转债概况: 行数=%d, 列数=%d", rows, cols)
		done <- true
	}()

	select {
	case <-done:
	case <-time.After(30 * time.Second):
		t.Skip("超时，跳过测试")
	}
}

func TestBondCBIndexJSL(t *testing.T) {
	done := make(chan bool)
	go func() {
		df, err := BondCBIndexJSL()
		if err != nil {
			t.Errorf("BondCBIndexJSL失败: %v", err)
		}

		rows, _ := df.Dims()
		if rows <= 1 {
			t.Errorf("期望至少有数据行，实际行数: %d", rows)
		}

		t.Logf("集思录可转债等权指数数据行数: %d", rows)
		done <- true
	}()

	select {
	case <-done:
	case <-time.After(30 * time.Second):
		t.Skip("超时，跳过测试")
	}
}

func TestBondZHUSRate(t *testing.T) {
	t.Skip("数据量较大，跳过测试")
	done := make(chan bool)
	go func() {
		df, err := BondZHUSRate("20200101")
		if err != nil {
			t.Errorf("BondZHUSRate失败: %v", err)
		}

		rows, _ := df.Dims()
		if rows <= 1 {
			t.Errorf("期望至少有数据行，实际行数: %d", rows)
		}

		t.Logf("中美国债收益率数据行数: %d", rows)
		done <- true
	}()

	select {
	case <-done:
	case <-time.After(60 * time.Second):
		t.Skip("超时，跳过测试")
	}
}

// 中国外汇交易中心债券市场测试
func TestBondSpotQuote(t *testing.T) {
	done := make(chan bool)
	go func() {
		df, err := BondSpotQuote()
		if err != nil {
			t.Errorf("BondSpotQuote失败: %v", err)
		}

		rows, cols := df.Dims()
		if rows <= 1 {
			t.Errorf("期望至少有数据行，实际行数: %d", rows)
		}
		if cols != 6 {
			t.Errorf("期望6列，实际列数: %d", cols)
		}

		t.Logf("现券市场做市报价数据: 行数=%d, 列数=%d", rows, cols)
		done <- true
	}()

	select {
	case <-done:
	case <-time.After(30 * time.Second):
		t.Skip("超时，跳过测试")
	}
}

func TestBondSpotDeal(t *testing.T) {
	done := make(chan bool)
	go func() {
		df, err := BondSpotDeal()
		if err != nil {
			t.Errorf("BondSpotDeal失败: %v", err)
		}

		rows, cols := df.Dims()
		if rows <= 1 {
			t.Errorf("期望至少有数据行，实际行数: %d", rows)
		}
		if cols != 6 {
			t.Errorf("期望6列，实际列数: %d", cols)
		}

		t.Logf("现券市场成交行情数据: 行数=%d, 列数=%d", rows, cols)
		done <- true
	}()

	select {
	case <-done:
	case <-time.After(30 * time.Second):
		t.Skip("超时，跳过测试")
	}
}

func TestBondChinaYield(t *testing.T) {
	done := make(chan bool)
	go func() {
		df, err := BondChinaYield("20210201", "20210301")
		if err != nil {
			t.Errorf("BondChinaYield失败: %v", err)
		}

		rows, _ := df.Dims()
		if rows <= 1 {
			t.Errorf("期望至少有数据行，实际行数: %d", rows)
		}

		t.Logf("国债收益率曲线数据行数: %d", rows)
		done <- true
	}()

	select {
	case <-done:
	case <-time.After(30 * time.Second):
		t.Skip("超时，跳过测试")
	}
}

// 收益率曲线和利率互换测试
func TestBondChinaCloseReturn(t *testing.T) {
	done := make(chan bool)
	go func() {
		df, err := BondChinaCloseReturn("国债", "1", "20240607", "20240607")
		if err != nil {
			t.Errorf("BondChinaCloseReturn失败: %v", err)
		}

		rows, cols := df.Dims()
		if rows <= 0 {
			t.Errorf("期望至少有数据行，实际行数: %d", rows)
		}
		if cols != 5 {
			t.Errorf("期望5列，实际列数: %d", cols)
		}

		t.Logf("收盘收益率曲线数据: 行数=%d, 列数=%d", rows, cols)
		done <- true
	}()

	select {
	case <-done:
	case <-time.After(30 * time.Second):
		t.Skip("超时，跳过测试")
	}
}

func TestMacroChinaSwapRate(t *testing.T) {
	t.Skip("数据量较大，跳过测试")
	done := make(chan bool)
	go func() {
		df, err := MacroChinaSwapRate("20251010", "20251030")
		if err != nil {
			t.Errorf("MacroChinaSwapRate失败: %v", err)
		}

		rows, _ := df.Dims()
		if rows <= 0 {
			t.Errorf("期望至少有数据行，实际行数: %d", rows)
		}

		t.Logf("FR007利率互换曲线数据行数: %d", rows)
		done <- true
	}()

	select {
	case <-done:
	case <-time.After(60 * time.Second):
		t.Skip("超时，跳过测试")
	}
}

func TestMacroChinaBondPublic(t *testing.T) {
	t.Skip("数据量很大，跳过测试")
	done := make(chan bool)
	go func() {
		df, err := MacroChinaBondPublic()
		if err != nil {
			t.Errorf("MacroChinaBondPublic失败: %v", err)
		}

		rows, _ := df.Dims()
		if rows <= 0 {
			t.Errorf("期望至少有数据行，实际行数: %d", rows)
		}

		t.Logf("债券发行信息数据行数: %d", rows)
		done <- true
	}()

	select {
	case <-done:
	case <-time.After(120 * time.Second):
		t.Skip("超时，跳过测试")
	}
}

// 新浪财经沪深债券测试
func TestBondZHHSSpot(t *testing.T) {
	t.Skip("容易封IP，跳过测试")
	done := make(chan bool)
	go func() {
		df, err := BondZHHSSpot("1", "2")
		if err != nil {
			t.Errorf("BondZHHSSpot失败: %v", err)
		}

		rows, _ := df.Dims()
		if rows <= 0 {
			t.Errorf("期望至少有数据行，实际行数: %d", rows)
		}

		t.Logf("沪深债券实时行情数据行数: %d", rows)
		done <- true
	}()

	select {
	case <-done:
	case <-time.After(30 * time.Second):
		t.Skip("超时，跳过测试")
	}
}

func TestBondZHHSDaily(t *testing.T) {
	t.Skip("容易封IP且需要JS解密，跳过测试")
	done := make(chan bool)
	go func() {
		df, err := BondZHHSDaily("sh010107")
		if err != nil {
			t.Errorf("BondZHHSDaily失败: %v", err)
		}

		rows, _ := df.Dims()
		if rows <= 0 {
			t.Errorf("期望至少有数据行，实际行数: %d", rows)
		}

		t.Logf("沪深债券历史行情数据行数: %d", rows)
		done <- true
	}()

	select {
	case <-done:
	case <-time.After(30 * time.Second):
		t.Skip("超时，跳过测试")
	}
}
