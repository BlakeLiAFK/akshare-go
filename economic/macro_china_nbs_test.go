package economic

import (
	"testing"
)

func TestMacroChinaNBSNation(t *testing.T) {
	// 测试月度数据 - 使用完整路径
	df, err := MacroChinaNBSNation("月度数据", "工业 > 工业分大类行业出口交货值(2018-至今) > 废弃资源综合利用业", "LAST5")
	if err != nil {
		t.Fatalf("MacroChinaNBSNation 月度数据失败: %v", err)
	}
	t.Logf("月度数据行数: %d, 列数: %d", df.Nrow(), df.Ncol())
}

func TestMacroChinaNBSNationYearly(t *testing.T) {
	// 测试年度数据 - 使用完整路径
	df, err := MacroChinaNBSNation("年度数据", "国民经济核算 > 国内生产总值", "LAST10")
	if err != nil {
		t.Fatalf("MacroChinaNBSNation 年度数据失败: %v", err)
	}
	t.Logf("年度数据行数: %d, 列数: %d", df.Nrow(), df.Ncol())
}

func TestMacroChinaNBSRegion(t *testing.T) {
	// 测试分省数据 - 按指标查询所有地区
	df, err := MacroChinaNBSRegion("分省季度数据", "国民经济核算 > 地区生产总值", "地区生产总值_累计值(亿元)", "", "LAST5")
	if err != nil {
		t.Fatalf("MacroChinaNBSRegion 按指标失败: %v", err)
	}
	t.Logf("分省数据(按指标)行数: %d, 列数: %d", df.Nrow(), df.Ncol())
}

func TestMacroChinaNBSRegionByRegion(t *testing.T) {
	// 测试分省数据 - 按地区查询所有指标
	df, err := MacroChinaNBSRegion("分省季度数据", "人民生活 > 居民人均可支配收入", "", "北京市", "2018-2022")
	if err != nil {
		t.Fatalf("MacroChinaNBSRegion 按地区失败: %v", err)
	}
	t.Logf("分省数据(按地区)行数: %d, 列数: %d", df.Nrow(), df.Ncol())
}
