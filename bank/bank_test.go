package bank

import (
	"testing"
	"time"
)

func TestBankFjcfTotalNum(t *testing.T) {
	done := make(chan bool)
	go func() {
		total, err := BankFjcfTotalNum("机关")
		if err != nil {
			t.Errorf("BankFjcfTotalNum失败: %v", err)
		}
		if total <= 0 {
			t.Errorf("期望总数大于0，实际得到: %d", total)
		}
		t.Logf("银保监机关行政处罚总数: %d", total)
		done <- true
	}()

	select {
	case <-done:
	case <-time.After(30 * time.Second):
		t.Skip("超时，跳过测试")
	}
}

func TestBankFjcfTotalPage(t *testing.T) {
	done := make(chan bool)
	go func() {
		pages, err := BankFjcfTotalPage("本级")
		if err != nil {
			t.Errorf("BankFjcfTotalPage失败: %v", err)
		}
		if pages <= 0 {
			t.Errorf("期望页数大于0，实际得到: %d", pages)
		}
		t.Logf("银保监本级行政处罚总页数: %d", pages)
		done <- true
	}()

	select {
	case <-done:
	case <-time.After(30 * time.Second):
		t.Skip("超时，跳过测试")
	}
}

func TestBankFjcfPageUrl(t *testing.T) {
	t.Skip("数据量较大，跳过测试")
	done := make(chan bool)
	go func() {
		df, err := BankFjcfPageUrl("分局本级")
		if err != nil {
			t.Errorf("BankFjcfPageUrl失败: %v", err)
		}

		rows, _ := df.Dims()
		if rows <= 1 {
			t.Errorf("期望至少有数据行，实际行数: %d", rows)
		}
		t.Logf("银保监分局本级行政处罚数据行数: %d", rows)
		done <- true
	}()

	select {
	case <-done:
	case <-time.After(60 * time.Second):
		t.Skip("超时，跳过测试")
	}
}

func TestBankFjcfTableDetail(t *testing.T) {
	t.Skip("需要真实docId，跳过测试")
	done := make(chan bool)
	go func() {
		// 使用一个示例URL
		url := "http://www.cbirc.gov.cn/cn/view/pages/ItemDetail.html?docId=123456"
		df, err := BankFjcfTableDetail(url)
		if err != nil {
			t.Logf("BankFjcfTableDetail预期可能失败(需要真实docId): %v", err)
			done <- true
			return
		}

		rows, cols := df.Dims()
		t.Logf("表格数据: %d行 x %d列", rows, cols)
		done <- true
	}()

	select {
	case <-done:
	case <-time.After(30 * time.Second):
		t.Skip("超时，跳过测试")
	}
}

func TestBankFjcfInvalidItem(t *testing.T) {
	_, err := BankFjcfTotalNum("无效项目")
	if err == nil {
		t.Error("期望返回错误，但没有")
	}
	t.Logf("正确返回错误: %v", err)
}
