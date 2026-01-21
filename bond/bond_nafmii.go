package bond

import (
	"fmt"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/go-gota/gota/dataframe"
	"github.com/tidwall/gjson"
)

// BondDebtNAFMII 银行间市场债务融资工具
//
// 获取中国银行间市场交易商协会的债务融资工具注册信息
//
// 参数:
//   - page: 页码，如 "1"
//
// 返回:
//   - dataframe.DataFrame: 包含债务融资工具注册信息
//   - error: 错误信息
//
// 数据源: http://zhuce.nafmii.org.cn/fans/publicQuery/manager
func BondDebtNAFMII(page string) (dataframe.DataFrame, error) {
	if page == "" {
		page = "1"
	}

	url := "http://zhuce.nafmii.org.cn/fans/publicQuery/releFileProjDataGrid"
	formData := map[string]string{
		"regFileName": "",
		"itemType":    "",
		"startTime":   "",
		"endTime":     "",
		"entityName":  "",
		"leadManager": "",
		"regPrdtType": "",
		"page":        page,
		"rows":        "50",
	}

	resp, err := utils.PostForm(url, formData)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	// 解析JSON响应
	rows := gjson.Get(resp.String(), "rows")
	if !rows.Exists() || !rows.IsArray() {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到数据")
	}

	// 构建DataFrame记录
	var records []map[string]any

	rows.ForEach(func(_, item gjson.Result) bool {
		record := map[string]any{
			"债券名称":    item.Get("regFileName").String(),
			"品种":      item.Get("regPrdtType").String(),
			"注册或备案":   item.Get("isReg").String(),
			"金额":      utils.MustParseFloat(item.Get("firstIssueAmount").String()),
			"注册通知书文号": item.Get("regNoticeNo").String(),
			"更新日期":    item.Get("releaseTime").String(),
			"项目状态":    item.Get("projPhase").String(),
		}
		records = append(records, record)
		return true
	})

	if len(records) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到有效数据")
	}

	df := dataframe.LoadMaps(records)
	return df, nil
}
