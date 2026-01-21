package bond

import (
	"fmt"
	"strconv"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
	"github.com/tidwall/gjson"
)

// BondInfoCmQuery 中国外汇交易中心暨全国银行间同业拆借中心-查询相关指标的参数
//
// 参数:
//   - symbol: 查询类型，可选 {"主承销商", "债券类型", "息票类型", "发行年份", "评级等级"}
//
// 返回:
//   - dataframe.DataFrame: 包含 name 和 code 两列
//   - error: 错误信息
//
// 数据源: https://www.chinamoney.com.cn/chinese/scsjzqxx/
func BondInfoCmQuery(symbol string) (dataframe.DataFrame, error) {
	if symbol == "主承销商" {
		url := "https://www.chinamoney.com.cn/ags/ms/cm-u-bond-md/EntyFullNameSearchCondition"
		resp, err := utils.PostForm(url, nil)
		if err != nil {
			return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
		}

		data := gjson.Get(resp.String(), "data.enty")
		if !data.Exists() || !data.IsArray() {
			return dataframe.DataFrame{}, fmt.Errorf("未获取到数据")
		}

		var names, codes []string
		data.ForEach(func(_, item gjson.Result) bool {
			arr := item.Array()
			if len(arr) >= 2 {
				codes = append(codes, arr[0].String())
				names = append(names, arr[1].String())
			}
			return true
		})

		df := dataframe.New(
			series.New(names, series.String, "name"),
			series.New(codes, series.String, "code"),
		)
		return df, nil
	}

	// 其他类型的查询
	symbolMap := map[string]string{
		"债券类型": "bondType",
		"息票类型": "couponType",
		"发行年份": "issueYear",
		"评级等级": "bondRtngShrt",
	}

	fieldName, ok := symbolMap[symbol]
	if !ok {
		return dataframe.DataFrame{}, fmt.Errorf("不支持的查询类型: %s", symbol)
	}

	url := "https://www.chinamoney.com.cn/ags/ms/cm-u-bond-md/BondBaseInfoSearchCondition"
	resp, err := utils.PostForm(url, nil)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	data := gjson.Get(resp.String(), "data."+fieldName)
	if !data.Exists() || !data.IsArray() {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到数据")
	}

	var names, codes []string
	data.ForEach(func(_, item gjson.Result) bool {
		arr := item.Array()
		if len(arr) >= 2 {
			codes = append(codes, arr[0].String())
			names = append(names, arr[1].String())
		} else if len(arr) == 1 {
			// 只有一个值时，code和name相同
			val := arr[0].String()
			codes = append(codes, val)
			names = append(names, val)
		}
		return true
	})

	df := dataframe.New(
		series.New(names, series.String, "name"),
		series.New(codes, series.String, "code"),
	)
	return df, nil
}

// BondInfoCm 中国外汇交易中心暨全国银行间同业拆借中心-数据-债券信息-信息查询
//
// 参数:
//   - bondName: 债券名称
//   - bondCode: 债券代码
//   - bondIssue: 发行人/受托机构
//   - bondType: 债券类型
//   - couponType: 息票类型
//   - issueYear: 发行年份
//   - underwriter: 主承销商
//   - grade: 评级等级
//
// 返回:
//   - dataframe.DataFrame: 信息查询结果
//   - error: 错误信息
//
// 数据源: https://www.chinamoney.com.cn/chinese/scsjzqxx/
func BondInfoCm(bondName, bondCode, bondIssue, bondType, couponType, issueYear, underwriter, grade string) (dataframe.DataFrame, error) {
	// 处理债券类型
	bondTypeValue := ""
	if bondType != "" {
		df, err := BondInfoCmQuery("债券类型")
		if err == nil {
			for i := 0; i < df.Nrow(); i++ {
				if df.Elem(i, 0).String() == bondType {
					bondTypeValue = df.Elem(i, 1).String()
					break
				}
			}
		}
	}

	// 处理息票类型
	couponTypeValue := ""
	if couponType != "" {
		df, err := BondInfoCmQuery("息票类型")
		if err == nil {
			for i := 0; i < df.Nrow(); i++ {
				if df.Elem(i, 0).String() == couponType {
					couponTypeValue = df.Elem(i, 1).String()
					break
				}
			}
		}
	}

	// 处理主承销商
	underwriterValue := ""
	if underwriter != "" {
		df, err := BondInfoCmQuery("主承销商")
		if err == nil {
			for i := 0; i < df.Nrow(); i++ {
				if df.Elem(i, 0).String() == underwriter {
					underwriterValue = df.Elem(i, 1).String()
					break
				}
			}
		}
	}

	url := "https://www.chinamoney.com.cn/ags/ms/cm-u-bond-md/BondMarketInfoList2"

	// 获取总页数
	formData := map[string]string{
		"pageNo":            "1",
		"pageSize":          "15",
		"bondName":          bondName,
		"bondCode":          bondCode,
		"issueEnty":         bondIssue,
		"bondType":          bondTypeValue,
		"bondSpclPrjctVrty": "",
		"couponType":        couponTypeValue,
		"issueYear":         issueYear,
		"entyDefinedCode":   underwriterValue,
		"rtngShrt":          grade,
	}

	resp, err := utils.PostForm(url, formData)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	totalPage := gjson.Get(resp.String(), "data.pageTotal").Int()

	// 收集所有数据
	var records []map[string]interface{}

	for page := int64(1); page <= totalPage; page++ {
		formData["pageNo"] = strconv.FormatInt(page, 10)
		resp, err := utils.PostForm(url, formData)
		if err != nil {
			continue
		}

		resultList := gjson.Get(resp.String(), "data.resultList")
		if !resultList.Exists() || !resultList.IsArray() {
			continue
		}

		resultList.ForEach(func(_, item gjson.Result) bool {
			record := map[string]interface{}{
				"债券简称":     item.Get("bondName").String(),
				"债券代码":     item.Get("bondCode").String(),
				"发行人/受托机构": item.Get("entyFullName").String(),
				"债券类型":     item.Get("bondType").String(),
				"发行日期":     item.Get("issueStartDate").String(),
				"最新债项评级":   item.Get("debtRtng").String(),
				"查询代码":     item.Get("bondDefinedCode").String(),
			}
			records = append(records, record)
			return true
		})
	}

	if len(records) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到有效数据")
	}

	df := dataframe.LoadMaps(records)
	// 选择列顺序
	df = df.Select([]string{"债券简称", "债券代码", "发行人/受托机构", "债券类型", "发行日期", "最新债项评级", "查询代码"})
	return df, nil
}

// BondInfoDetailCm 中国外汇交易中心暨全国银行间同业拆借中心-数据-债券信息-信息查询-债券详情
//
// 参数:
//   - symbol: 债券简称
//
// 返回:
//   - dataframe.DataFrame: 债券详情，包含 name 和 value 两列
//   - error: 错误信息
//
// 数据源: https://www.chinamoney.com.cn/chinese/zqjc/?bondDefinedCode=xxx
func BondInfoDetailCm(symbol string) (dataframe.DataFrame, error) {
	// 先查询获取债券的查询代码
	bondInfoDf, err := BondInfoCm(symbol, "", "", "", "", "", "", "")
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("查询债券信息失败: %w", err)
	}

	if bondInfoDf.Nrow() == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("未找到债券: %s", symbol)
	}

	bondCode := bondInfoDf.Elem(0, 6).String() // 查询代码列

	url := "https://www.chinamoney.com.cn/ags/ms/cm-u-bond-md/BondDetailInfo"
	formData := map[string]string{
		"bondDefinedCode": bondCode,
	}

	resp, err := utils.PostForm(url, formData)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	bondBaseInfo := gjson.Get(resp.String(), "data.bondBaseInfo")
	if !bondBaseInfo.Exists() {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到债券详情")
	}

	var names, values []string
	bondBaseInfo.ForEach(func(key, value gjson.Result) bool {
		keyStr := key.String()
		// 跳过复杂嵌套字段
		if keyStr == "creditRateEntyList" || keyStr == "exerciseInfoList" {
			return true
		}
		names = append(names, keyStr)
		values = append(values, value.String())
		return true
	})

	df := dataframe.New(
		series.New(names, series.String, "name"),
		series.New(values, series.String, "value"),
	)
	return df, nil
}
