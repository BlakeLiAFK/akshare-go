package article

import (
	"testing"
	"time"
)

// TestArticleOmanRVShort 测试Oxford-Man简化版已实现波动率接口
func TestArticleOmanRVShort(t *testing.T) {
	t.Skip("Oxford-Man数据源不稳定，跳过测试")

	df, err := ArticleOmanRVShort("FTSE")
	if err != nil {
		t.Fatalf("ArticleOmanRVShort失败: %v", err)
	}

	if df.Nrow() == 0 {
		t.Error("返回数据为空")
	}

	t.Logf("获取到 %d 行数据", df.Nrow())
	t.Logf("列名: %v", df.Names())
}

// TestArticleOmanRV 测试Oxford-Man已实现波动率接口
func TestArticleOmanRV(t *testing.T) {
	t.Skip("Oxford-Man数据源不稳定，跳过测试")

	df, err := ArticleOmanRV("FTSE", "rk_th2")
	if err != nil {
		t.Fatalf("ArticleOmanRV失败: %v", err)
	}

	if df.Nrow() == 0 {
		t.Error("返回数据为空")
	}

	t.Logf("获取到 %d 行数据", df.Nrow())
	t.Logf("列名: %v", df.Names())
}

// TestArticleRlabRV 测试Risk-Lab已实现波动率接口
func TestArticleRlabRV(t *testing.T) {
	t.Skip("Risk-Lab数据源不稳定，跳过测试")

	df, err := ArticleRlabRV("39693")
	if err != nil {
		t.Fatalf("ArticleRlabRV失败: %v", err)
	}

	if df.Nrow() == 0 {
		t.Error("返回数据为空")
	}

	t.Logf("获取到 %d 行数据", df.Nrow())
	t.Logf("列名: %v", df.Names())
}

// TestArticleFFCRR 测试Fama-French多因子数据接口
func TestArticleFFCRR(t *testing.T) {
	done := make(chan bool)

	go func() {
		defer func() { done <- true }()

		df, err := ArticleFFCRR()
		if err != nil {
			t.Errorf("ArticleFFCRR失败: %v", err)
			return
		}

		if df.Nrow() == 0 {
			t.Error("返回数据为空")
			return
		}

		t.Logf("获取到 %d 行数据", df.Nrow())
		t.Logf("列名: %v", df.Names())
	}()

	select {
	case <-done:
	case <-time.After(30 * time.Second):
		t.Skip("超时，跳过测试")
	}
}

// TestArticleEPUIndex 测试经济政策不确定性指数接口
func TestArticleEPUIndex(t *testing.T) {
	testCases := []struct {
		name   string
		symbol string
	}{
		{"中国", "China"},
		{"美国", "USA"},
		{"欧洲", "Europe"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Skip("EPU数据源不稳定，跳过测试")

			df, err := ArticleEPUIndex(tc.symbol)
			if err != nil {
				t.Errorf("ArticleEPUIndex(%s)失败: %v", tc.symbol, err)
				return
			}

			if df.Nrow() == 0 {
				t.Errorf("%s 返回数据为空", tc.name)
				return
			}

			t.Logf("%s: 获取到 %d 行数据", tc.name, df.Nrow())
			t.Logf("%s: 列名: %v", tc.name, df.Names())
		})
	}
}

// TestFredMD 测试FRED月度数据接口
func TestFredMD(t *testing.T) {
	done := make(chan bool)

	go func() {
		defer func() { done <- true }()

		// 使用较早的日期确保数据存在
		df, err := FredMD("2020-01")
		if err != nil {
			t.Errorf("FredMD失败: %v", err)
			return
		}

		if df.Nrow() == 0 {
			t.Error("返回数据为空")
			return
		}

		t.Logf("获取到 %d 行数据", df.Nrow())
		t.Logf("列名: %v", df.Names())
	}()

	select {
	case <-done:
	case <-time.After(30 * time.Second):
		t.Skip("超时，跳过测试")
	}
}

// TestFredQD 测试FRED季度数据接口
func TestFredQD(t *testing.T) {
	done := make(chan bool)

	go func() {
		defer func() { done <- true }()

		// 使用较早的日期确保数据存在
		df, err := FredQD("2020-01")
		if err != nil {
			t.Errorf("FredQD失败: %v", err)
			return
		}

		if df.Nrow() == 0 {
			t.Error("返回数据为空")
			return
		}

		t.Logf("获取到 %d 行数据", df.Nrow())
		t.Logf("列名: %v", df.Names())
	}()

	select {
	case <-done:
	case <-time.After(30 * time.Second):
		t.Skip("超时，跳过测试")
	}
}
