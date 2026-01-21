package fund

import (
	"fmt"
	"strconv"
	"time"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/tidwall/gjson"
)

var amacHeaders = map[string]string{
	"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
}

// AmacMemberInfo 获取中国证券投资基金业协会-信息公示-会员信息-会员机构综合查询
// https://gs.amac.org.cn/amac-infodisc/res/pof/member/index.html
func AmacMemberInfo() ([]map[string]interface{}, error) {
	url := "https://gs.amac.org.cn/amac-infodisc/api/pof/pofMember"
	params := map[string]string{
		"rand": "0.7665138514630696",
		"page": "0",
		"size": "100",
	}

	// 获取总页数
	resp, err := utils.PostJSONWithHeaders(url, params, map[string]interface{}{}, amacHeaders)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	result := gjson.ParseBytes(resp.Body())
	totalPage := result.Get("totalPages").Int()

	records := make([]map[string]interface{}, 0)

	// 遍历所有分页
	for page := int64(0); page < totalPage; page++ {
		params["page"] = strconv.FormatInt(page, 10)
		resp, err = utils.PostJSONWithHeaders(url, params, map[string]interface{}{}, amacHeaders)
		if err != nil {
			continue
		}

		result := gjson.ParseBytes(resp.Body())
		items := result.Get("content").Array()

		for _, item := range items {
			record := map[string]interface{}{
				"机构（会员）名称": item.Get("managerName").String(),
				"会员代表":     item.Get("memberBehalf").String(),
				"会员类型":     item.Get("memberType").String(),
				"会员编号":     item.Get("memberCode").String(),
				"入会时间":     convertTimestamp(item.Get("memberDate").Int()),
				"公司类型":     item.Get("primaryInvestType").String(),
			}
			records = append(records, record)
		}
	}

	return records, nil
}

// AmacPersonFundOrg 获取中国证券投资基金业协会-信息公示-从业人员信息-基金从业人员资格注册信息
// https://gs.amac.org.cn/amac-infodisc/res/person/fund/index.html
func AmacPersonFundOrg(symbol string) ([]map[string]interface{}, error) {
	if symbol == "" {
		symbol = "公募基金管理公司"
	}

	url := "https://gs.amac.org.cn/amac-infodisc/api/person/fund"
	params := map[string]string{
		"rand": "0.7665138514630696",
		"page": "0",
		"size": "100",
	}

	payload := map[string]interface{}{
		"aoiName": symbol,
	}

	// 获取总页数
	resp, err := utils.PostJSONWithHeaders(url, params, payload, amacHeaders)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	result := gjson.ParseBytes(resp.Body())
	totalPage := result.Get("totalPages").Int()

	records := make([]map[string]interface{}, 0)

	// 遍历所有分页
	for page := int64(0); page < totalPage; page++ {
		params["page"] = strconv.FormatInt(page, 10)
		resp, err = utils.PostJSONWithHeaders(url, params, payload, amacHeaders)
		if err != nil {
			continue
		}

		result := gjson.ParseBytes(resp.Body())
		items := result.Get("content").Array()

		for _, item := range items {
			record := map[string]interface{}{
				"姓名":     item.Get("prcName").String(),
				"性别":     item.Get("prcGender").String(),
				"资格证书编号": item.Get("prcNo").String(),
				"执业岗位":   item.Get("prcPost").String(),
				"注册日期":   convertTimestamp(item.Get("registerDate").Int()),
				"变更日期":   convertTimestamp(item.Get("updateDate").Int()),
				"所在机构":   item.Get("aoiName").String(),
			}
			records = append(records, record)
		}
	}

	return records, nil
}

// AmacPersonFund 获取中国证券投资基金业协会-信息公示-从业人员信息-债券投资交易相关人员公示
// https://gs.amac.org.cn/amac-infodisc/res/person/bond/index.html
func AmacPersonFund() ([]map[string]interface{}, error) {
	url := "https://gs.amac.org.cn/amac-infodisc/api/person/bond"
	params := map[string]string{
		"rand": "0.7665138514630696",
		"page": "0",
		"size": "100",
	}

	// 获取总页数
	resp, err := utils.PostJSONWithHeaders(url, params, map[string]interface{}{}, amacHeaders)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	result := gjson.ParseBytes(resp.Body())
	totalPage := result.Get("totalPages").Int()

	records := make([]map[string]interface{}, 0)

	// 遍历所有分页
	for page := int64(0); page < totalPage; page++ {
		params["page"] = strconv.FormatInt(page, 10)
		resp, err = utils.PostJSONWithHeaders(url, params, map[string]interface{}{}, amacHeaders)
		if err != nil {
			continue
		}

		result := gjson.ParseBytes(resp.Body())
		items := result.Get("content").Array()

		for _, item := range items {
			record := map[string]interface{}{
				"姓名":     item.Get("prcName").String(),
				"性别":     item.Get("prcGender").String(),
				"证书编号":   item.Get("prcNo").String(),
				"执业岗位":   item.Get("prcPost").String(),
				"注册日期":   convertTimestamp(item.Get("registerDate").Int()),
				"所在机构":   item.Get("aoiName").String(),
				"机构所在省市": item.Get("orgProvinceName").String(),
			}
			records = append(records, record)
		}
	}

	return records, nil
}

// AmacManagerInfo 获取中国证券投资基金业协会-信息公示-私募基金管理人公示-私募基金管理人综合查询
// https://gs.amac.org.cn/amac-infodisc/res/pof/manager/index.html
func AmacManagerInfo() ([]map[string]interface{}, error) {
	url := "https://gs.amac.org.cn/amac-infodisc/api/pof/manager"
	params := map[string]string{
		"rand": "0.7665138514630696",
		"page": "0",
		"size": "100",
	}

	// 获取总页数
	resp, err := utils.PostJSONWithHeaders(url, params, map[string]interface{}{}, amacHeaders)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	result := gjson.ParseBytes(resp.Body())
	totalPage := result.Get("totalPages").Int()

	records := make([]map[string]interface{}, 0)

	// 遍历所有分页
	for page := int64(0); page < totalPage; page++ {
		params["page"] = strconv.FormatInt(page, 10)
		resp, err = utils.PostJSONWithHeaders(url, params, map[string]interface{}{}, amacHeaders)
		if err != nil {
			continue
		}

		result := gjson.ParseBytes(resp.Body())
		items := result.Get("content").Array()

		for _, item := range items {
			record := map[string]interface{}{
				"私募基金管理人名称":             item.Get("managerName").String(),
				"法定代表人/执行事务合伙人(委派代表)姓名": item.Get("artificialPersonName").String(),
				"机构类型":    item.Get("primaryInvestType").String(),
				"登记编号":    item.Get("registerNo").String(),
				"注册地":     item.Get("registerProvince").String(),
				"办公地":     item.Get("officeAdrAgg").String(),
				"成立时间":    convertTimestamp(item.Get("establishDate").Int()),
				"登记时间":    convertTimestamp(item.Get("registerDate").Int()),
				"在管基金数量":  item.Get("fundCount").Int(),
				"会员类型":    item.Get("memberType").String(),
				"是否有提示信息": boolToString(item.Get("hasSpecialTips").Bool()),
				"是否有诚信信息": boolToString(item.Get("hasCreditTips").Bool()),
			}
			records = append(records, record)
		}
	}

	return records, nil
}

// AmacManagerClassifyInfo 获取中国证券投资基金业协会-信息公示-私募基金管理人公示-私募基金管理人分类公示
// https://gs.amac.org.cn/amac-infodisc/res/pof/manager/index.html
func AmacManagerClassifyInfo() ([]map[string]interface{}, error) {
	url := "https://gs.amac.org.cn/amac-infodisc/api/pof/manager"
	params := map[string]string{
		"rand": "0.7665138514630696",
		"page": "0",
		"size": "100",
	}

	payload := map[string]interface{}{
		"primaryInvestType": "私募证券投资基金管理人",
	}

	// 获取总页数
	resp, err := utils.PostJSONWithHeaders(url, params, payload, amacHeaders)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	result := gjson.ParseBytes(resp.Body())
	totalPage := result.Get("totalPages").Int()

	records := make([]map[string]interface{}, 0)

	// 遍历所有分页
	for page := int64(0); page < totalPage; page++ {
		params["page"] = strconv.FormatInt(page, 10)
		resp, err = utils.PostJSONWithHeaders(url, params, payload, amacHeaders)
		if err != nil {
			continue
		}

		result := gjson.ParseBytes(resp.Body())
		items := result.Get("content").Array()

		for _, item := range items {
			record := map[string]interface{}{
				"私募基金管理人名称":             item.Get("managerName").String(),
				"法定代表人/执行事务合伙人(委派代表)姓名": item.Get("artificialPersonName").String(),
				"机构类型":    item.Get("primaryInvestType").String(),
				"登记编号":    item.Get("registerNo").String(),
				"注册地":     item.Get("registerProvince").String(),
				"办公地":     item.Get("officeAdrAgg").String(),
				"成立时间":    convertTimestamp(item.Get("establishDate").Int()),
				"登记时间":    convertTimestamp(item.Get("registerDate").Int()),
				"在管基金数量":  item.Get("fundCount").Int(),
				"会员类型":    item.Get("memberType").String(),
				"是否有提示信息": boolToString(item.Get("hasSpecialTips").Bool()),
				"是否有诚信信息": boolToString(item.Get("hasCreditTips").Bool()),
			}
			records = append(records, record)
		}
	}

	return records, nil
}

// AmacMemberSubInfo 获取中国证券投资基金业协会-信息公示-私募基金管理人公示-证券公司私募基金子公司管理人信息公示
// https://gs.amac.org.cn/amac-infodisc/res/pof/member/index.html?primaryInvestType=private
func AmacMemberSubInfo() ([]map[string]interface{}, error) {
	url := "https://gs.amac.org.cn/amac-infodisc/api/pof/pofMember"
	params := map[string]string{
		"rand": "0.7665138514630696",
		"page": "0",
		"size": "100",
	}

	// 获取总页数
	resp, err := utils.PostJSONWithHeaders(url, params, map[string]interface{}{}, amacHeaders)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	result := gjson.ParseBytes(resp.Body())
	totalPage := result.Get("totalPages").Int()

	records := make([]map[string]interface{}, 0)

	// 遍历所有分页
	for page := int64(0); page < totalPage; page++ {
		params["page"] = strconv.FormatInt(page, 10)
		resp, err = utils.PostJSONWithHeaders(url, params, map[string]interface{}{}, amacHeaders)
		if err != nil {
			continue
		}

		result := gjson.ParseBytes(resp.Body())
		items := result.Get("content").Array()

		for _, item := range items {
			record := map[string]interface{}{
				"机构（会员）名称": item.Get("managerName").String(),
				"会员代表":     item.Get("memberBehalf").String(),
				"会员类型":     item.Get("memberType").String(),
				"会员编号":     item.Get("memberCode").String(),
				"入会时间":     convertTimestamp(item.Get("memberDate").Int()),
				"公司类型":     item.Get("primaryInvestType").String(),
			}
			records = append(records, record)
		}
	}

	return records, nil
}

// AmacFundInfo 获取中国证券投资基金业协会-信息公示-基金产品-私募基金管理人基金产品
// https://gs.amac.org.cn/amac-infodisc/res/pof/fund/index.html
func AmacFundInfo(startPage, endPage string) ([]map[string]interface{}, error) {
	if startPage == "" {
		startPage = "1"
	}
	if endPage == "" {
		endPage = "2000"
	}

	url := "https://gs.amac.org.cn/amac-infodisc/api/pof/fund"
	params := map[string]string{
		"rand": "0.7665138514630696",
		"page": "0",
		"size": "100",
	}

	// 获取总页数
	resp, err := utils.PostJSONWithHeaders(url, params, map[string]interface{}{}, amacHeaders)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	result := gjson.ParseBytes(resp.Body())
	totalPage := result.Get("totalPages").Int()

	// 确定实际结束页
	startPageInt, _ := strconv.ParseInt(startPage, 10, 64)
	endPageInt, _ := strconv.ParseInt(endPage, 10, 64)
	if totalPage < endPageInt {
		endPageInt = totalPage
	}

	records := make([]map[string]interface{}, 0)

	// 遍历指定范围分页
	for page := startPageInt - 1; page < endPageInt; page++ {
		params["page"] = strconv.FormatInt(page, 10)
		resp, err = utils.PostJSONWithHeaders(url, params, map[string]interface{}{}, amacHeaders)
		if err != nil {
			continue
		}

		result := gjson.ParseBytes(resp.Body())
		items := result.Get("content").Array()

		for _, item := range items {
			record := map[string]interface{}{
				"基金名称":      item.Get("fundName").String(),
				"私募基金管理人名称": item.Get("managerName").String(),
				"私募基金管理人类型": item.Get("managerType").String(),
				"运行状态":      item.Get("workingState").String(),
				"备案时间":      convertTimestamp(item.Get("putOnRecordDate").Int()),
				"建立时间":      convertTimestamp(item.Get("establishDate").Int()),
				"托管人名称":     item.Get("mandatorName").String(),
			}
			records = append(records, record)
		}
	}

	return records, nil
}

// AmacSecuritiesInfo 获取中国证券投资基金业协会-信息公示-基金产品-证券公司集合资管产品公示
// https://gs.amac.org.cn/amac-infodisc/res/pof/securities/index.html
func AmacSecuritiesInfo() ([]map[string]interface{}, error) {
	url := "https://gs.amac.org.cn/amac-infodisc/api/pof/securities"
	params := map[string]string{
		"rand": "0.7665138514630696",
		"page": "0",
		"size": "100",
	}

	// 获取总页数
	resp, err := utils.PostJSONWithHeaders(url, params, map[string]interface{}{}, amacHeaders)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	result := gjson.ParseBytes(resp.Body())
	totalPage := result.Get("totalPages").Int()

	records := make([]map[string]interface{}, 0)

	// 遍历所有分页
	for page := int64(0); page < totalPage; page++ {
		params["page"] = strconv.FormatInt(page, 10)
		resp, err = utils.PostJSONWithHeaders(url, params, map[string]interface{}{}, amacHeaders)
		if err != nil {
			continue
		}

		result := gjson.ParseBytes(resp.Body())
		items := result.Get("content").Array()

		for _, item := range items {
			record := map[string]interface{}{
				"产品名称":  item.Get("cpmc").String(),
				"产品编码":  item.Get("cpbm").String(),
				"管理人名称": item.Get("gljg").String(),
				"成立日期":  item.Get("slrq").String(),
				"到期时间":  item.Get("dqr").String(),
				"投资类型":  item.Get("tzlx").String(),
				"是否分级":  item.Get("sffj").String(),
				"托管人名称": item.Get("tgjg").String(),
				"备案日期":  item.Get("barq").String(),
				"运作状态":  item.Get("yzzt").String(),
			}
			records = append(records, record)
		}
	}

	return records, nil
}

// AmacAicInfo 获取中国证券投资基金业协会-信息公示-基金产品-证券公司直投基金
// https://gs.amac.org.cn/amac-infodisc/res/aoin/product/index.html
func AmacAicInfo() ([]map[string]interface{}, error) {
	url := "https://gs.amac.org.cn/amac-infodisc/api/aoin/product"
	params := map[string]string{
		"rand": "0.7665138514630696",
		"page": "0",
		"size": "100",
	}

	// 获取总页数
	resp, err := utils.PostJSONWithHeaders(url, params, map[string]interface{}{}, amacHeaders)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	result := gjson.ParseBytes(resp.Body())
	totalPage := result.Get("totalPages").Int()

	records := make([]map[string]interface{}, 0)

	// 遍历所有分页
	for page := int64(0); page < totalPage; page++ {
		params["page"] = strconv.FormatInt(page, 10)
		resp, err = utils.PostJSONWithHeaders(url, params, map[string]interface{}{}, amacHeaders)
		if err != nil {
			continue
		}

		result := gjson.ParseBytes(resp.Body())
		items := result.Get("content").Array()

		for _, item := range items {
			record := map[string]interface{}{
				"产品编码":  item.Get("code").String(),
				"产品名称":  item.Get("name").String(),
				"直投子公司": item.Get("aoinName").String(),
				"管理机构":  item.Get("managerName").String(),
				"设立日期":  convertTimestamp(item.Get("createDate").Int()),
			}
			records = append(records, record)
		}
	}

	return records, nil
}

// AmacFundSub 获取中国证券投资基金业协会-信息公示-基金产品-证券公司私募投资基金
// https://gs.amac.org.cn/amac-infodisc/res/pof/subfund/index.html
func AmacFundSub() ([]map[string]interface{}, error) {
	url := "https://gs.amac.org.cn/amac-infodisc/api/pof/subfund"
	params := map[string]string{
		"rand": "0.7665138514630696",
		"page": "0",
		"size": "100",
	}

	// 获取总页数
	resp, err := utils.PostJSONWithHeaders(url, params, map[string]interface{}{}, amacHeaders)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	result := gjson.ParseBytes(resp.Body())
	totalPage := result.Get("totalPages").Int()

	records := make([]map[string]interface{}, 0)

	// 遍历所有分页
	for page := int64(0); page < totalPage; page++ {
		params["page"] = strconv.FormatInt(page, 10)
		resp, err = utils.PostJSONWithHeaders(url, params, map[string]interface{}{}, amacHeaders)
		if err != nil {
			continue
		}

		result := gjson.ParseBytes(resp.Body())
		items := result.Get("content").Array()

		for _, item := range items {
			record := map[string]interface{}{
				"产品编码":      item.Get("productCode").String(),
				"产品名称":      item.Get("productName").String(),
				"私募基金管理人名称": item.Get("mgrName").String(),
				"托管人名称":     item.Get("trustee").String(),
				"成立日期":      convertTimestamp(item.Get("foundDate").Int()),
				"备案日期":      convertTimestamp(item.Get("registeredDate").Int()),
			}
			records = append(records, record)
		}
	}

	return records, nil
}

// AmacFundAccount 获取中国证券投资基金业协会-信息公示-基金产品-基金公司及子公司集合资管产品公示
// https://gs.amac.org.cn/amac-infodisc/res/fund/account/index.html
func AmacFundAccount() ([]map[string]interface{}, error) {
	url := "https://gs.amac.org.cn/amac-infodisc/api/fund/account"
	params := map[string]string{
		"rand": "0.7665138514630696",
		"page": "0",
		"size": "100",
	}

	// 获取总页数
	resp, err := utils.PostJSONWithHeaders(url, params, map[string]interface{}{}, amacHeaders)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	result := gjson.ParseBytes(resp.Body())
	totalPage := result.Get("totalPages").Int()

	records := make([]map[string]interface{}, 0)

	// 遍历所有分页
	for page := int64(0); page < totalPage; page++ {
		params["page"] = strconv.FormatInt(page, 10)
		resp, err = utils.PostJSONWithHeaders(url, params, map[string]interface{}{}, amacHeaders)
		if err != nil {
			continue
		}

		result := gjson.ParseBytes(resp.Body())
		items := result.Get("content").Array()

		for _, item := range items {
			record := map[string]interface{}{
				"成立日期":  convertTimestamp(item.Get("registerDate").Int()),
				"产品编码":  item.Get("registerCode").String(),
				"产品名称":  item.Get("name").String(),
				"管理人名称": item.Get("manager").String(),
			}
			records = append(records, record)
		}
	}

	return records, nil
}

// AmacFundAccountSubFund 获取中国证券投资基金业协会-信息公示-基金产品-资产支持专项计划
// https://gs.amac.org.cn/amac-infodisc/res/fund/abs/index.html
func AmacFundAccountSubFund() ([]map[string]interface{}, error) {
	url := "https://gs.amac.org.cn/amac-infodisc/api/fund/abs"
	params := map[string]string{
		"rand": "0.7665138514630696",
		"page": "0",
		"size": "100",
	}

	// 获取总页数
	resp, err := utils.PostJSONWithHeaders(url, params, map[string]interface{}{}, amacHeaders)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	result := gjson.ParseBytes(resp.Body())
	totalPage := result.Get("totalPages").Int()

	records := make([]map[string]interface{}, 0)
	idx := 1

	// 遍历所有分页
	for page := int64(0); page < totalPage; page++ {
		params["page"] = strconv.FormatInt(page, 10)
		resp, err = utils.PostJSONWithHeaders(url, params, map[string]interface{}{}, amacHeaders)
		if err != nil {
			continue
		}

		result := gjson.ParseBytes(resp.Body())
		items := result.Get("content").Array()

		for _, item := range items {
			values := item.Array()
			if len(values) < 6 {
				continue
			}

			record := map[string]interface{}{
				"编号":     idx,
				"备案编号":   values[1].String(),
				"专项计划全称": values[3].String(),
				"管理人":    values[4].String(),
				"托管人":    values[5].String(),
				"成立日期":   convertTimestamp(values[7].Int()),
				"预期到期时间": convertTimestamp(values[8].Int()),
				"备案通过时间": convertTimestamp(values[6].Int()),
			}
			records = append(records, record)
			idx++
		}
	}

	return records, nil
}

// AmacFuturesInfo 获取中国证券投资基金业协会-信息公示-基金产品-期货公司集合资管产品公示
// https://gs.amac.org.cn/amac-infodisc/res/pof/futures/index.html
func AmacFuturesInfo() ([]map[string]interface{}, error) {
	url := "https://gs.amac.org.cn/amac-infodisc/api/pof/futures"
	params := map[string]string{
		"rand": "0.7665138514630696",
		"page": "0",
		"size": "100",
	}

	// 获取总页数
	resp, err := utils.PostJSONWithHeaders(url, params, map[string]interface{}{}, amacHeaders)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	result := gjson.ParseBytes(resp.Body())
	totalPage := result.Get("totalPages").Int()

	records := make([]map[string]interface{}, 0)

	// 遍历所有分页
	for page := int64(0); page < totalPage; page++ {
		params["page"] = strconv.FormatInt(page, 10)
		resp, err = utils.PostJSONWithHeaders(url, params, map[string]interface{}{}, amacHeaders)
		if err != nil {
			continue
		}

		result := gjson.ParseBytes(resp.Body())
		items := result.Get("content").Array()

		for _, item := range items {
			record := map[string]interface{}{
				"产品名称":  item.Get("mpiName").String(),
				"产品编码":  item.Get("mpiProductCode").String(),
				"管理人名称": item.Get("aoiName").String(),
				"托管人名称": item.Get("mpiTrustee").String(),
				"成立日期":  item.Get("mpiCreateDate").String(),
				"投资类型":  item.Get("tzlx").String(),
				"是否分级":  item.Get("sfjgh").String(),
				"备案日期":  item.Get("registeredDate").String(),
				"到期日":   item.Get("dueDate").String(),
				"运作状态":  item.Get("fundStatus").String(),
			}
			records = append(records, record)
		}
	}

	return records, nil
}

// AmacManagerSecuritiesInfo 获取中国证券投资基金业协会-信息公示-诚信信息-已注销私募基金管理人名单
// https://gs.amac.org.cn/amac-infodisc/res/cancelled/manager/index.html
func AmacManagerSecuritiesInfo() ([]map[string]interface{}, error) {
	url := "https://gs.amac.org.cn/amac-infodisc/api/cancelled/manager"
	params := map[string]string{
		"rand": "0.7665138514630696",
		"page": "0",
		"size": "100",
	}

	// 获取总页数
	resp, err := utils.PostJSONWithHeaders(url, params, map[string]interface{}{}, amacHeaders)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	result := gjson.ParseBytes(resp.Body())
	totalPage := result.Get("totalPages").Int()

	records := make([]map[string]interface{}, 0)

	// 遍历所有分页
	for page := int64(0); page < totalPage; page++ {
		params["page"] = strconv.FormatInt(page, 10)
		resp, err = utils.PostJSONWithHeaders(url, params, map[string]interface{}{}, amacHeaders)
		if err != nil {
			continue
		}

		result := gjson.ParseBytes(resp.Body())
		items := result.Get("content").Array()

		for _, item := range items {
			record := map[string]interface{}{
				"管理人名称":    item.Get("orgName").String(),
				"统一社会信用代码": item.Get("orgCode").String(),
				"登记时间":     convertTimestamp(item.Get("orgSignDate").Int()),
				"注销时间":     convertTimestamp(item.Get("cancelDate").Int()),
				"注销类型":     item.Get("status").String(),
			}
			records = append(records, record)
		}
	}

	return records, nil
}

// AmacMemberRosterInfo 此函数在Python源码中未找到对应实现
func AmacMemberRosterInfo() ([]map[string]interface{}, error) {
	return nil, fmt.Errorf("此接口暂无对应Python源实现")
}

// convertTimestamp 将毫秒时间戳转换为日期字符串
func convertTimestamp(timestamp int64) string {
	if timestamp == 0 {
		return ""
	}
	t := time.Unix(timestamp/1000, 0)
	return t.Format("2006-01-02")
}

// boolToString 将布尔值转换为"是"/"否"
func boolToString(b bool) string {
	if b {
		return "是"
	}
	return "否"
}
