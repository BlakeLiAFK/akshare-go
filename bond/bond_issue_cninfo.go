package bond

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"strconv"
	"time"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/go-gota/gota/dataframe"
	"github.com/tidwall/gjson"
)

// getCninfoEncKey 生成巨潮资讯API所需的Accept-Enckey
// 等同于JS中的getResCode1函数，使用AES-CBC加密当前时间戳
func getCninfoEncKey() (string, error) {
	timestamp := time.Now().Unix()
	timestampStr := strconv.FormatInt(timestamp, 10)

	key := []byte("1234567887654321")
	iv := []byte("1234567887654321")

	plaintext := []byte(timestampStr)
	padding := aes.BlockSize - len(plaintext)%aes.BlockSize
	padtext := make([]byte, len(plaintext)+padding)
	copy(padtext, plaintext)
	for i := len(plaintext); i < len(padtext); i++ {
		padtext[i] = byte(padding)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("创建AES加密器失败: %w", err)
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	ciphertext := make([]byte, len(padtext))
	mode.CryptBlocks(ciphertext, padtext)

	encoded := base64.StdEncoding.EncodeToString(ciphertext)
	return encoded, nil
}

// getCninfoHeaders 获取巨潮资讯API请求头
func getCninfoHeaders() (map[string]string, error) {
	mcode, err := getCninfoEncKey()
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"Accept":           "*/*",
		"Accept-Enckey":    mcode,
		"Accept-Encoding":  "gzip, deflate",
		"Accept-Language":  "zh-CN,zh;q=0.9,en;q=0.8",
		"Cache-Control":    "no-cache",
		"Content-Length":   "0",
		"Host":             "webapi.cninfo.com.cn",
		"Origin":           "http://webapi.cninfo.com.cn",
		"Pragma":           "no-cache",
		"Proxy-Connection": "keep-alive",
		"Referer":          "http://webapi.cninfo.com.cn/",
		"User-Agent":       "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/93.0.4577.63 Safari/537.36",
		"X-Requested-With": "XMLHttpRequest",
	}, nil
}

// formatDate 格式化日期 "20210910" -> "2021-09-10"
func formatDate(date string) string {
	if len(date) != 8 {
		return date
	}
	return fmt.Sprintf("%s-%s-%s", date[:4], date[4:6], date[6:])
}

// BondTreasureIssueCninfo 巨潮资讯-数据中心-专题统计-债券报表-债券发行-国债发行
//
// 参数:
//   - startDate: 开始统计时间，格式 "20210910"
//   - endDate: 结束统计时间，格式 "20211109"
//
// 返回:
//   - dataframe.DataFrame: 国债发行数据
//   - error: 错误信息
//
// 数据源: http://webapi.cninfo.com.cn/#/thematicStatistics
func BondTreasureIssueCninfo(startDate, endDate string) (dataframe.DataFrame, error) {
	url := "http://webapi.cninfo.com.cn/api/sysapi/p_sysapi1120"

	headers, err := getCninfoHeaders()
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("生成请求头失败: %w", err)
	}

	params := map[string]string{
		"sdate": formatDate(startDate),
		"edate": formatDate(endDate),
	}

	resp, err := utils.PostWithHeaders(url, params, headers)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	records := gjson.Get(resp.String(), "records")
	if !records.Exists() || !records.IsArray() {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到数据")
	}

	var dfRecords []map[string]interface{}
	records.ForEach(func(_, item gjson.Result) bool {
		record := map[string]interface{}{
			"债券代码":   item.Get("SECCODE").String(),
			"债券简称":   item.Get("SECNAME").String(),
			"发行起始日":  item.Get("F004D").String(),
			"发行终止日":  item.Get("F003D").String(),
			"计划发行总量": utils.MustParseFloat(item.Get("F006N").String()),
			"实际发行总量": utils.MustParseFloat(item.Get("F005N").String()),
			"发行价格":   utils.MustParseFloat(item.Get("F007N").String()),
			"单位面值":   utils.MustParseFloat(item.Get("F008N").String()),
			"缴款日":    item.Get("F009D").String(),
			"增发次数":   utils.MustParseFloat(item.Get("F028N").String()),
			"交易市场":   item.Get("F002V").String(),
			"发行方式":   item.Get("F013V").String(),
			"发行对象":   item.Get("F014V").String(),
			"公告日期":   item.Get("DECLAREDATE").String(),
			"债券名称":   item.Get("BONDNAME").String(),
		}
		dfRecords = append(dfRecords, record)
		return true
	})

	if len(dfRecords) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到有效数据")
	}

	df := dataframe.LoadMaps(dfRecords)
	df = df.Select([]string{"债券代码", "债券简称", "发行起始日", "发行终止日", "计划发行总量", "实际发行总量", "发行价格", "单位面值", "缴款日", "增发次数", "交易市场", "发行方式", "发行对象", "公告日期", "债券名称"})
	return df, nil
}

// BondLocalGovernmentIssueCninfo 巨潮资讯-数据中心-专题统计-债券报表-债券发行-地方债发行
//
// 参数:
//   - startDate: 开始统计时间，格式 "20210911"
//   - endDate: 结束统计时间，格式 "20211110"
//
// 返回:
//   - dataframe.DataFrame: 地方债发行数据
//   - error: 错误信息
//
// 数据源: http://webapi.cninfo.com.cn/#/thematicStatistics
func BondLocalGovernmentIssueCninfo(startDate, endDate string) (dataframe.DataFrame, error) {
	url := "http://webapi.cninfo.com.cn/api/sysapi/p_sysapi1121"

	headers, err := getCninfoHeaders()
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("生成请求头失败: %w", err)
	}

	params := map[string]string{
		"sdate": formatDate(startDate),
		"edate": formatDate(endDate),
	}

	resp, err := utils.PostWithHeaders(url, params, headers)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	records := gjson.Get(resp.String(), "records")
	if !records.Exists() || !records.IsArray() {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到数据")
	}

	var dfRecords []map[string]interface{}
	records.ForEach(func(_, item gjson.Result) bool {
		record := map[string]interface{}{
			"债券代码":   item.Get("SECCODE").String(),
			"债券简称":   item.Get("SECNAME").String(),
			"发行起始日":  item.Get("F004D").String(),
			"发行终止日":  item.Get("F003D").String(),
			"计划发行总量": utils.MustParseFloat(item.Get("F006N").String()),
			"实际发行总量": utils.MustParseFloat(item.Get("F005N").String()),
			"发行价格":   utils.MustParseFloat(item.Get("F007N").String()),
			"单位面值":   utils.MustParseFloat(item.Get("F008N").String()),
			"缴款日":    item.Get("F009D").String(),
			"增发次数":   utils.MustParseFloat(item.Get("F028N").String()),
			"交易市场":   item.Get("F002V").String(),
			"发行方式":   item.Get("F013V").String(),
			"发行对象":   item.Get("F014V").String(),
			"公告日期":   item.Get("DECLAREDATE").String(),
			"债券名称":   item.Get("BONDNAME").String(),
		}
		dfRecords = append(dfRecords, record)
		return true
	})

	if len(dfRecords) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到有效数据")
	}

	df := dataframe.LoadMaps(dfRecords)
	df = df.Select([]string{"债券代码", "债券简称", "发行起始日", "发行终止日", "计划发行总量", "实际发行总量", "发行价格", "单位面值", "缴款日", "增发次数", "交易市场", "发行方式", "发行对象", "公告日期", "债券名称"})
	return df, nil
}

// BondCorporateIssueCninfo 巨潮资讯-数据中心-专题统计-债券报表-债券发行-企业债发行
//
// 参数:
//   - startDate: 开始统计时间，格式 "20210911"
//   - endDate: 结束统计时间，格式 "20211110"
//
// 返回:
//   - dataframe.DataFrame: 企业债发行数据
//   - error: 错误信息
//
// 数据源: http://webapi.cninfo.com.cn/#/thematicStatistics
func BondCorporateIssueCninfo(startDate, endDate string) (dataframe.DataFrame, error) {
	url := "http://webapi.cninfo.com.cn/api/sysapi/p_sysapi1122"

	headers, err := getCninfoHeaders()
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("生成请求头失败: %w", err)
	}

	params := map[string]string{
		"sdate": formatDate(startDate),
		"edate": formatDate(endDate),
	}

	resp, err := utils.PostWithHeaders(url, params, headers)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	records := gjson.Get(resp.String(), "records")
	if !records.Exists() || !records.IsArray() {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到数据")
	}

	var dfRecords []map[string]interface{}
	records.ForEach(func(_, item gjson.Result) bool {
		record := map[string]interface{}{
			"债券代码":       item.Get("SECCODE").String(),
			"债券简称":       item.Get("SECNAME").String(),
			"公告日期":       item.Get("DECLAREDATE").String(),
			"交易所网上发行起始日": item.Get("F003D").String(),
			"交易所网上发行终止日": item.Get("F004D").String(),
			"计划发行总量":     utils.MustParseFloat(item.Get("F005N").String()),
			"实际发行总量":     utils.MustParseFloat(item.Get("F006N").String()),
			"发行面值":       utils.MustParseFloat(item.Get("F008N").String()),
			"发行价格":       utils.MustParseFloat(item.Get("F007N").String()),
			"发行方式":       item.Get("F013V").String(),
			"发行对象":       item.Get("F014V").String(),
			"发行范围":       item.Get("F015V").String(),
			"承销方式":       item.Get("F017V").String(),
			"最小认购单位":     utils.MustParseFloat(item.Get("F022N").String()),
			"募资用途说明":     item.Get("F023V").String(),
			"最低认购额":      utils.MustParseFloat(item.Get("F052N").String()),
			"债券名称":       item.Get("BONDNAME").String(),
		}
		dfRecords = append(dfRecords, record)
		return true
	})

	if len(dfRecords) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到有效数据")
	}

	df := dataframe.LoadMaps(dfRecords)
	df = df.Select([]string{"债券代码", "债券简称", "公告日期", "交易所网上发行起始日", "交易所网上发行终止日", "计划发行总量", "实际发行总量", "发行面值", "发行价格", "发行方式", "发行对象", "发行范围", "承销方式", "最小认购单位", "募资用途说明", "最低认购额", "债券名称"})
	return df, nil
}

// BondCovIssueCninfo 巨潮资讯-数据中心-专题统计-债券报表-债券发行-可转债发行
//
// 参数:
//   - startDate: 开始统计时间，格式 "20210913"
//   - endDate: 结束统计时间，格式 "20211112"
//
// 返回:
//   - dataframe.DataFrame: 可转债发行数据
//   - error: 错误信息
//
// 数据源: http://webapi.cninfo.com.cn/#/thematicStatistics
func BondCovIssueCninfo(startDate, endDate string) (dataframe.DataFrame, error) {
	url := "http://webapi.cninfo.com.cn/api/sysapi/p_sysapi1123"

	headers, err := getCninfoHeaders()
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("生成请求头失败: %w", err)
	}

	params := map[string]string{
		"sdate": formatDate(startDate),
		"edate": formatDate(endDate),
	}

	resp, err := utils.PostWithHeaders(url, params, headers)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	records := gjson.Get(resp.String(), "records")
	if !records.Exists() || !records.IsArray() {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到数据")
	}

	var dfRecords []map[string]interface{}
	records.ForEach(func(_, item gjson.Result) bool {
		record := map[string]interface{}{
			"债券代码":     item.Get("SECCODE").String(),
			"债券简称":     item.Get("SECNAME").String(),
			"公告日期":     item.Get("DECLAREDATE").String(),
			"发行起始日":    item.Get("F029D").String(),
			"发行终止日":    item.Get("F003D").String(),
			"计划发行总量":   utils.MustParseFloat(item.Get("F005N").String()),
			"实际发行总量":   utils.MustParseFloat(item.Get("F006N").String()),
			"发行面值":     utils.MustParseFloat(item.Get("F007N").String()),
			"发行价格":     utils.MustParseFloat(item.Get("F052N").String()),
			"发行方式":     item.Get("F013V").String(),
			"发行对象":     item.Get("F014V").String(),
			"发行范围":     item.Get("F015V").String(),
			"承销方式":     item.Get("F017V").String(),
			"募资用途说明":   item.Get("F021V").String(),
			"初始转股价格":   utils.MustParseFloat(item.Get("F026N").String()),
			"转股开始日期":   item.Get("F027D").String(),
			"转股终止日期":   item.Get("F053D").String(),
			"网上申购日期":   item.Get("F051D").String(),
			"网上申购代码":   item.Get("F031V").String(),
			"网上申购简称":   item.Get("F032V").String(),
			"网上申购数量上限": utils.MustParseFloat(item.Get("F008N").String()),
			"网上申购数量下限": utils.MustParseFloat(item.Get("F066N").String()),
			"网上申购单位":   utils.MustParseFloat(item.Get("F067N").String()),
			"网上申购中签结果公告日及退款日": item.Get("F068D").String(),
			"优先申购日":   item.Get("F004D").String(),
			"配售价格":    utils.MustParseFloat(item.Get("F065N").String()),
			"债权登记日":   item.Get("F028D").String(),
			"优先申购缴款日": item.Get("F054D").String(),
			"转股代码":    item.Get("F086V").String(),
			"交易市场":    item.Get("F002V").String(),
			"债券名称":    item.Get("BONDNAME").String(),
		}
		dfRecords = append(dfRecords, record)
		return true
	})

	if len(dfRecords) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到有效数据")
	}

	df := dataframe.LoadMaps(dfRecords)
	df = df.Select([]string{"债券代码", "债券简称", "公告日期", "发行起始日", "发行终止日", "计划发行总量", "实际发行总量", "发行面值", "发行价格", "发行方式", "发行对象", "发行范围", "承销方式", "募资用途说明", "初始转股价格", "转股开始日期", "转股终止日期", "网上申购日期", "网上申购代码", "网上申购简称", "网上申购数量上限", "网上申购数量下限", "网上申购单位", "网上申购中签结果公告日及退款日", "优先申购日", "配售价格", "债权登记日", "优先申购缴款日", "转股代码", "交易市场", "债券名称"})
	return df, nil
}

// BondCovStockIssueCninfo 巨潮资讯-数据中心-专题统计-债券报表-债券发行-可转债转股
//
// 返回:
//   - dataframe.DataFrame: 可转债转股数据
//   - error: 错误信息
//
// 数据源: http://webapi.cninfo.com.cn/#/thematicStatistics
func BondCovStockIssueCninfo() (dataframe.DataFrame, error) {
	url := "http://webapi.cninfo.com.cn/api/sysapi/p_sysapi1124"

	headers, err := getCninfoHeaders()
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("生成请求头失败: %w", err)
	}

	resp, err := utils.PostWithHeaders(url, nil, headers)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}

	records := gjson.Get(resp.String(), "records")
	if !records.Exists() || !records.IsArray() {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到数据")
	}

	var dfRecords []map[string]interface{}
	records.ForEach(func(_, item gjson.Result) bool {
		record := map[string]interface{}{
			"债券代码":     item.Get("SECCODE").String(),
			"债券简称":     item.Get("SECNAME").String(),
			"公告日期":     item.Get("DECLAREDATE").String(),
			"转股代码":     item.Get("F001V").String(),
			"转股简称":     item.Get("F002V").String(),
			"转股价格":     utils.MustParseFloat(item.Get("F003N").String()),
			"自愿转换期起始日": item.Get("F004D").String(),
			"自愿转换期终止日": item.Get("F005D").String(),
			"标的股票":     item.Get("F017V").String(),
			"债券名称":     item.Get("BONDNAME").String(),
		}
		dfRecords = append(dfRecords, record)
		return true
	})

	if len(dfRecords) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到有效数据")
	}

	df := dataframe.LoadMaps(dfRecords)
	df = df.Select([]string{"债券代码", "债券简称", "公告日期", "转股代码", "转股简称", "转股价格", "自愿转换期起始日", "自愿转换期终止日", "标的股票", "债券名称"})
	return df, nil
}
