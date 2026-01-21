package fund

import (
	"testing"
)

func TestAmacMemberInfo(t *testing.T) {
	records, err := AmacMemberInfo()
	if err != nil {
		t.Logf("AmacMemberInfo warning: %v (may be unavailable)", err)
		return
	}

	t.Logf("AmacMemberInfo returned %d records", len(records))

	if len(records) > 0 {
		first := records[0]
		requiredFields := []string{"机构（会员）名称", "会员代表", "会员类型", "会员编号", "入会时间", "公司类型"}
		for _, field := range requiredFields {
			if _, ok := first[field]; !ok {
				t.Errorf("Missing required field: %s", field)
			}
		}
	}
}

func TestAmacPersonFundOrg(t *testing.T) {
	records, err := AmacPersonFundOrg("公募基金管理公司")
	if err != nil {
		t.Logf("AmacPersonFundOrg warning: %v (may be unavailable)", err)
		return
	}

	t.Logf("AmacPersonFundOrg returned %d records", len(records))

	if len(records) > 0 {
		first := records[0]
		requiredFields := []string{"姓名", "性别", "资格证书编号", "执业岗位", "注册日期", "所在机构"}
		for _, field := range requiredFields {
			if _, ok := first[field]; !ok {
				t.Errorf("Missing required field: %s", field)
			}
		}
	}
}

func TestAmacPersonFund(t *testing.T) {
	records, err := AmacPersonFund()
	if err != nil {
		t.Logf("AmacPersonFund warning: %v (may be unavailable)", err)
		return
	}

	t.Logf("AmacPersonFund returned %d records", len(records))

	if len(records) > 0 {
		first := records[0]
		requiredFields := []string{"姓名", "性别", "证书编号", "执业岗位", "注册日期", "所在机构"}
		for _, field := range requiredFields {
			if _, ok := first[field]; !ok {
				t.Errorf("Missing required field: %s", field)
			}
		}
	}
}

func TestAmacManagerInfo(t *testing.T) {
	records, err := AmacManagerInfo()
	if err != nil {
		t.Logf("AmacManagerInfo warning: %v (may be unavailable)", err)
		return
	}

	t.Logf("AmacManagerInfo returned %d records", len(records))

	if len(records) > 0 {
		first := records[0]
		requiredFields := []string{"私募基金管理人名称", "机构类型", "登记编号", "成立时间", "登记时间"}
		for _, field := range requiredFields {
			if _, ok := first[field]; !ok {
				t.Errorf("Missing required field: %s", field)
			}
		}
	}
}

func TestAmacManagerClassifyInfo(t *testing.T) {
	records, err := AmacManagerClassifyInfo()
	if err != nil {
		t.Logf("AmacManagerClassifyInfo warning: %v (may be unavailable)", err)
		return
	}

	t.Logf("AmacManagerClassifyInfo returned %d records", len(records))

	if len(records) > 0 {
		first := records[0]
		requiredFields := []string{"私募基金管理人名称", "机构类型", "登记编号"}
		for _, field := range requiredFields {
			if _, ok := first[field]; !ok {
				t.Errorf("Missing required field: %s", field)
			}
		}
	}
}

func TestAmacMemberSubInfo(t *testing.T) {
	records, err := AmacMemberSubInfo()
	if err != nil {
		t.Logf("AmacMemberSubInfo warning: %v (may be unavailable)", err)
		return
	}

	t.Logf("AmacMemberSubInfo returned %d records", len(records))

	if len(records) > 0 {
		first := records[0]
		requiredFields := []string{"机构（会员）名称", "会员代表", "会员类型", "会员编号"}
		for _, field := range requiredFields {
			if _, ok := first[field]; !ok {
				t.Errorf("Missing required field: %s", field)
			}
		}
	}
}

func TestAmacFundInfo(t *testing.T) {
	// 测试小范围分页，避免测试时间过长
	records, err := AmacFundInfo("1", "2")
	if err != nil {
		t.Logf("AmacFundInfo warning: %v (may be unavailable)", err)
		return
	}

	t.Logf("AmacFundInfo returned %d records", len(records))

	if len(records) > 0 {
		first := records[0]
		requiredFields := []string{"基金名称", "私募基金管理人名称", "运行状态", "备案时间"}
		for _, field := range requiredFields {
			if _, ok := first[field]; !ok {
				t.Errorf("Missing required field: %s", field)
			}
		}
	}
}

func TestAmacSecuritiesInfo(t *testing.T) {
	records, err := AmacSecuritiesInfo()
	if err != nil {
		t.Logf("AmacSecuritiesInfo warning: %v (may be unavailable)", err)
		return
	}

	t.Logf("AmacSecuritiesInfo returned %d records", len(records))

	if len(records) > 0 {
		first := records[0]
		requiredFields := []string{"产品名称", "产品编码", "管理人名称", "成立日期"}
		for _, field := range requiredFields {
			if _, ok := first[field]; !ok {
				t.Errorf("Missing required field: %s", field)
			}
		}
	}
}

func TestAmacAicInfo(t *testing.T) {
	records, err := AmacAicInfo()
	if err != nil {
		t.Logf("AmacAicInfo warning: %v (may be unavailable)", err)
		return
	}

	t.Logf("AmacAicInfo returned %d records", len(records))

	if len(records) > 0 {
		first := records[0]
		requiredFields := []string{"产品编码", "产品名称", "直投子公司", "管理机构"}
		for _, field := range requiredFields {
			if _, ok := first[field]; !ok {
				t.Errorf("Missing required field: %s", field)
			}
		}
	}
}

func TestAmacFundSub(t *testing.T) {
	records, err := AmacFundSub()
	if err != nil {
		t.Logf("AmacFundSub warning: %v (may be unavailable)", err)
		return
	}

	t.Logf("AmacFundSub returned %d records", len(records))

	if len(records) > 0 {
		first := records[0]
		requiredFields := []string{"产品编码", "产品名称", "私募基金管理人名称", "托管人名称"}
		for _, field := range requiredFields {
			if _, ok := first[field]; !ok {
				t.Errorf("Missing required field: %s", field)
			}
		}
	}
}

func TestAmacFundAccount(t *testing.T) {
	records, err := AmacFundAccount()
	if err != nil {
		t.Logf("AmacFundAccount warning: %v (may be unavailable)", err)
		return
	}

	t.Logf("AmacFundAccount returned %d records", len(records))

	if len(records) > 0 {
		first := records[0]
		requiredFields := []string{"产品编码", "产品名称", "管理人名称", "成立日期"}
		for _, field := range requiredFields {
			if _, ok := first[field]; !ok {
				t.Errorf("Missing required field: %s", field)
			}
		}
	}
}

func TestAmacFundAccountSubFund(t *testing.T) {
	records, err := AmacFundAccountSubFund()
	if err != nil {
		t.Logf("AmacFundAccountSubFund warning: %v (may be unavailable)", err)
		return
	}

	t.Logf("AmacFundAccountSubFund returned %d records", len(records))

	if len(records) > 0 {
		first := records[0]
		requiredFields := []string{"编号", "备案编号", "专项计划全称", "管理人", "托管人"}
		for _, field := range requiredFields {
			if _, ok := first[field]; !ok {
				t.Errorf("Missing required field: %s", field)
			}
		}
	}
}

func TestAmacFuturesInfo(t *testing.T) {
	records, err := AmacFuturesInfo()
	if err != nil {
		t.Logf("AmacFuturesInfo warning: %v (may be unavailable)", err)
		return
	}

	t.Logf("AmacFuturesInfo returned %d records", len(records))

	if len(records) > 0 {
		first := records[0]
		requiredFields := []string{"产品名称", "产品编码", "管理人名称", "托管人名称"}
		for _, field := range requiredFields {
			if _, ok := first[field]; !ok {
				t.Errorf("Missing required field: %s", field)
			}
		}
	}
}

func TestAmacManagerSecuritiesInfo(t *testing.T) {
	records, err := AmacManagerSecuritiesInfo()
	if err != nil {
		t.Logf("AmacManagerSecuritiesInfo warning: %v (may be unavailable)", err)
		return
	}

	t.Logf("AmacManagerSecuritiesInfo returned %d records", len(records))

	if len(records) > 0 {
		first := records[0]
		requiredFields := []string{"管理人名称", "统一社会信用代码", "登记时间", "注销时间"}
		for _, field := range requiredFields {
			if _, ok := first[field]; !ok {
				t.Errorf("Missing required field: %s", field)
			}
		}
	}
}

func TestAmacMemberRosterInfo(t *testing.T) {
	_, err := AmacMemberRosterInfo()
	if err == nil {
		t.Error("AmacMemberRosterInfo should return error (not implemented)")
	}
	t.Logf("AmacMemberRosterInfo correctly returns error: %v", err)
}
