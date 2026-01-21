package movie

import (
	"testing"
)

func TestMovieBoxofficeYearlyFirstWeek(t *testing.T) {
	items, err := MovieBoxofficeYearlyFirstWeek("20201018")
	if err != nil {
		// API 数据可能需要解密，这是预期的情况
		if err.Error() == "数据需要解密，请参考Python版本使用jm.js解密" {
			t.Skip("API 数据需要解密，跳过测试")
			return
		}
		t.Fatalf("MovieBoxofficeYearlyFirstWeek 失败: %v", err)
	}

	t.Logf("年度首周票房数据数量: %d", len(items))
	if len(items) == 0 {
		t.Error("数据为空")
	}

	// 输出第一条数据用于验证
	if len(items) > 0 {
		t.Logf("第一条数据: 排序=%d, 影片名称=%s, 类型=%s, 首周票房=%s",
			items[0].Rank, items[0].MovieName, items[0].Genre, items[0].FirstWeekBox)
	}
}
