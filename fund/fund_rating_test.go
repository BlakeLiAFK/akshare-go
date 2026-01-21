package fund

import (
	"testing"
)

func TestFundRatingAll(t *testing.T) {
	records, err := FundRatingAll()
	if err != nil {
		t.Fatalf("FundRatingAll failed: %v", err)
	}

	// 允许空数据
	t.Logf("FundRatingAll returned %d records", len(records))

	// 如果有数据，验证字段
	if len(records) > 0 {
		first := records[0]
		requiredFields := []string{"基金代码", "基金简称", "3年评级"}
		for _, field := range requiredFields {
			if _, ok := first[field]; !ok {
				t.Errorf("Missing required field: %s", field)
			}
		}
	}
}

func TestFundRatingSh(t *testing.T) {
	// 测试默认日期
	records, err := FundRatingSh("")
	if err != nil {
		t.Logf("FundRatingSh warning: %v (may be unavailable)", err)
		return
	}

	// 允许空数据
	t.Logf("FundRatingSh returned %d records", len(records))

	// 如果有数据，验证字段
	if len(records) > 0 {
		first := records[0]
		requiredFields := []string{"基金代码", "基金简称", "评级", "评级日期"}
		for _, field := range requiredFields {
			if _, ok := first[field]; !ok {
				t.Errorf("Missing required field: %s", field)
			}
		}
	}
}

func TestFundRatingZs(t *testing.T) {
	// 测试默认日期
	records, err := FundRatingZs("")
	if err != nil {
		t.Logf("FundRatingZs warning: %v (may be unavailable)", err)
		return
	}

	// 允许空数据
	t.Logf("FundRatingZs returned %d records", len(records))

	// 如果有数据，验证字段
	if len(records) > 0 {
		first := records[0]
		requiredFields := []string{"基金代码", "基金简称", "评级", "评级日期"}
		for _, field := range requiredFields {
			if _, ok := first[field]; !ok {
				t.Errorf("Missing required field: %s", field)
			}
		}
	}
}

func TestFundRatingJa(t *testing.T) {
	// 测试默认日期
	records, err := FundRatingJa("")
	if err != nil {
		t.Logf("FundRatingJa warning: %v (may be unavailable)", err)
		return
	}

	// 允许空数据
	t.Logf("FundRatingJa returned %d records", len(records))

	// 如果有数据，验证字段
	if len(records) > 0 {
		first := records[0]
		requiredFields := []string{"基金代码", "基金简称", "评级", "评级日期"}
		for _, field := range requiredFields {
			if _, ok := first[field]; !ok {
				t.Errorf("Missing required field: %s", field)
			}
		}
	}
}
