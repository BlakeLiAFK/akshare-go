package stock_fundamental

import (
	"fmt"
	"testing"
)

// TestStockIpoDeclareEm 测试东方财富首发申报企业信息
func TestStockIpoDeclareEm(t *testing.T) {
	result, err := StockIpoDeclareEm()
	if err != nil {
		t.Fatalf("StockIpoDeclareEm() error = %v", err)
	}

	if len(result) == 0 {
		t.Logf("警告: 未获取到首发申报企业信息数据")
		return
	}

	t.Logf("获取到 %d 条首发申报企业信息", len(result))

	// 验证第一条数据
	first := result[0]
	t.Logf("第一条数据: 序号=%d, 企业名称=%s, 状态=%s, 保荐机构=%s",
		first.Index, first.CompanyName, first.State, first.RecommendOrg)

	// 验证必填字段
	if first.Index == 0 {
		t.Error("序号不能为0")
	}
	if first.CompanyName == "" {
		t.Error("企业名称不能为空")
	}
}

// ExampleStockIpoDeclareEm 示例：获取首发申报企业信息
func ExampleStockIpoDeclareEm() {
	declare, err := StockIpoDeclareEm()
	if err != nil {
		fmt.Printf("获取失败: %v\n", err)
		return
	}
	if len(declare) > 0 {
		fmt.Printf("企业名称: %s, 状态: %s\n", declare[0].CompanyName, declare[0].State)
	}
}

// TestStockIpoReviewEm 测试东方财富过会企业信息
func TestStockIpoReviewEm(t *testing.T) {
	result, err := StockIpoReviewEm()
	if err != nil {
		t.Fatalf("StockIpoReviewEm() error = %v", err)
	}

	if len(result) == 0 {
		t.Logf("警告: 未获取到过会企业信息数据")
		return
	}

	t.Logf("获取到 %d 条过会企业信息", len(result))

	// 验证第一条数据
	first := result[0]
	t.Logf("第一条数据: 序号=%d, 企业名称=%s, 股票简称=%s, 审核状态=%s",
		first.Index, first.CompanyName, first.StockName, first.ReviewState)

	// 验证必填字段
	if first.Index == 0 {
		t.Error("序号不能为0")
	}
	if first.CompanyName == "" {
		t.Error("企业名称不能为空")
	}
}

// ExampleStockIpoReviewEm 示例：获取过会企业信息
func ExampleStockIpoReviewEm() {
	review, err := StockIpoReviewEm()
	if err != nil {
		fmt.Printf("获取失败: %v\n", err)
		return
	}
	if len(review) > 0 {
		fmt.Printf("企业名称: %s, 股票简称: %s, 审核状态: %s\n",
			review[0].CompanyName, review[0].StockName, review[0].ReviewState)
	}
}

// TestStockIpoTutorEm 测试东方财富辅导备案信息
func TestStockIpoTutorEm(t *testing.T) {
	result, err := StockIpoTutorEm()
	if err != nil {
		t.Fatalf("StockIpoTutorEm() error = %v", err)
	}

	if len(result) == 0 {
		t.Logf("警告: 未获取到辅导备案信息数据")
		return
	}

	t.Logf("获取到 %d 条辅导备案信息", len(result))

	// 验证第一条数据
	first := result[0]
	t.Logf("第一条数据: 序号=%d, 企业名称=%s, 辅导机构=%s, 辅导状态=%s",
		first.Index, first.CompanyName, first.TutorOrg, first.TutorProcessState)

	// 验证必填字段
	if first.Index == 0 {
		t.Error("序号不能为0")
	}
	if first.CompanyName == "" {
		t.Error("企业名称不能为空")
	}
}

// ExampleStockIpoTutorEm 示例：获取辅导备案信息
func ExampleStockIpoTutorEm() {
	tutor, err := StockIpoTutorEm()
	if err != nil {
		fmt.Printf("获取失败: %v\n", err)
		return
	}
	if len(tutor) > 0 {
		fmt.Printf("企业名称: %s, 辅导机构: %s, 辅导状态: %s\n",
			tutor[0].CompanyName, tutor[0].TutorOrg, tutor[0].TutorProcessState)
	}
}

// TestStockKcbSse 测试上交所科创板信息
func TestStockKcbSse(t *testing.T) {
	result, err := StockKcbSse()
	if err != nil {
		t.Fatalf("StockKcbSse() error = %v", err)
	}

	if len(result) == 0 {
		t.Logf("警告: 未获取到科创板信息数据")
		return
	}

	t.Logf("获取到 %d 条科创板信息", len(result))

	// 验证第一条数据
	first := result[0]
	t.Logf("第一条数据: 序号=%d, 企业名称=%s, 股票代码=%s, 当前状态=%s",
		first.Index, first.CompanyName, first.StockCode, first.CurrentStatus)

	// 验证必填字段
	if first.Index == 0 {
		t.Error("序号不能为0")
	}
	if first.CompanyName == "" {
		t.Error("企业名称不能为空")
	}
}

// ExampleStockKcbSse 示例：获取科创板信息
func ExampleStockKcbSse() {
	kcb, err := StockKcbSse()
	if err != nil {
		fmt.Printf("获取失败: %v\n", err)
		return
	}
	if len(kcb) > 0 {
		fmt.Printf("企业名称: %s, 股票代码: %s, 当前状态: %s\n",
			kcb[0].CompanyName, kcb[0].StockCode, kcb[0].CurrentStatus)
	}
}
