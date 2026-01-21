package fund

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"strconv"
	"time"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/tidwall/gjson"
)

// getResCode1 生成巨潮资讯API所需的Accept-Enckey
// 等同于JS中的getResCode1函数，使用AES-CBC加密当前时间戳
func getResCode1() (string, error) {
	// 获取当前时间戳（秒）
	timestamp := time.Now().Unix()
	timestampStr := strconv.FormatInt(timestamp, 10)

	// AES密钥和IV都是'1234567887654321'
	key := []byte("1234567887654321")
	iv := []byte("1234567887654321")

	// Padding
	plaintext := []byte(timestampStr)
	padding := aes.BlockSize - len(plaintext)%aes.BlockSize
	padtext := make([]byte, len(plaintext)+padding)
	copy(padtext, plaintext)
	for i := len(plaintext); i < len(padtext); i++ {
		padtext[i] = byte(padding)
	}

	// 创建AES加密器
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("创建AES加密器失败: %w", err)
	}

	// 使用CBC模式
	mode := cipher.NewCBCEncrypter(block, iv)
	ciphertext := make([]byte, len(padtext))
	mode.CryptBlocks(ciphertext, padtext)

	// Base64编码
	encoded := base64.StdEncoding.EncodeToString(ciphertext)
	return encoded, nil
}

// FundReportStockCninfo 巨潮资讯-数据中心-专题统计-基金报表-基金重仓股
// https://webapi.cninfo.com.cn/#/thematicStatistics
// 参数: date 报告时间，格式: "YYYYMMDD"，如"20210630"
//
//	选项: "XXXX0331", "XXXX0630", "XXXX0930", "XXXX1231"
func FundReportStockCninfo(date string) ([]map[string]interface{}, error) {
	if date == "" {
		date = "20210630"
	}

	url := "https://webapi.cninfo.com.cn/api/sysapi/p_sysapi1112"

	// 生成Accept-Enckey
	mcode, err := getResCode1()
	if err != nil {
		return nil, fmt.Errorf("生成加密key失败: %w", err)
	}

	headers := map[string]string{
		"Accept":           "*/*",
		"Accept-Enckey":    mcode,
		"Accept-Encoding":  "gzip, deflate",
		"Accept-Language":  "zh-CN,zh;q=0.9,en;q=0.8",
		"Cache-Control":    "no-cache",
		"Content-Length":   "0",
		"Host":             "webapi.cninfo.com.cn",
		"Origin":           "https://webapi.cninfo.com.cn",
		"Pragma":           "no-cache",
		"Proxy-Connection": "keep-alive",
		"Referer":          "https://webapi.cninfo.com.cn/",
		"User-Agent":       "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/93.0.4577.63 Safari/537.36",
		"X-Requested-With": "XMLHttpRequest",
	}

	// 格式化日期: "20210630" -> "2021-06-30"
	rdate := fmt.Sprintf("%s-%s-%s", date[:4], date[4:6], date[6:])

	params := map[string]string{
		"rdate": rdate,
	}

	resp, err := utils.PostWithHeaders(url, params, headers)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	result := gjson.ParseBytes(resp.Body())
	items := result.Get("records").Array()

	records := make([]map[string]interface{}, 0, len(items))
	for i, item := range items {
		record := map[string]interface{}{
			"序号":     i + 1,
			"股票代码":   item.Get("SECCODE").String(),
			"股票简称":   item.Get("SECNAME").String(),
			"报告期":    item.Get("ENDDATE").String(),
			"基金覆盖家数": utils.MustFloat64(item.Get("F001N").String()),
			"持股总数":   utils.MustFloat64(item.Get("F002N").String()),
			"持股总市值":  utils.MustFloat64(item.Get("F003N").String()),
		}
		records = append(records, record)
	}

	return records, nil
}

// FundReportIndustryAllocationCninfo 巨潮资讯-数据中心-专题统计-基金报表-基金行业配置
// https://webapi.cninfo.com.cn/#/thematicStatistics
// 参数: date 报告时间，格式: "YYYYMMDD"，如"20210630"
//
//	从2017年开始，选项: "XXXX0331", "XXXX0630", "XXXX0930", "XXXX1231"
func FundReportIndustryAllocationCninfo(date string) ([]map[string]interface{}, error) {
	if date == "" {
		date = "20210630"
	}

	url := "https://webapi.cninfo.com.cn/api/sysapi/p_sysapi1113"

	mcode, err := getResCode1()
	if err != nil {
		return nil, fmt.Errorf("生成加密key失败: %w", err)
	}

	headers := map[string]string{
		"Accept":           "*/*",
		"Accept-Enckey":    mcode,
		"Accept-Encoding":  "gzip, deflate",
		"Accept-Language":  "zh-CN,zh;q=0.9,en;q=0.8",
		"Cache-Control":    "no-cache",
		"Content-Length":   "0",
		"Host":             "webapi.cninfo.com.cn",
		"Origin":           "https://webapi.cninfo.com.cn",
		"Pragma":           "no-cache",
		"Proxy-Connection": "keep-alive",
		"Referer":          "https://webapi.cninfo.com.cn/",
		"User-Agent":       "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/93.0.4577.63 Safari/537.36",
		"X-Requested-With": "XMLHttpRequest",
	}

	rdate := fmt.Sprintf("%s-%s-%s", date[:4], date[4:6], date[6:])

	params := map[string]string{
		"rdate": rdate,
	}

	resp, err := utils.PostWithHeaders(url, params, headers)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	result := gjson.ParseBytes(resp.Body())
	items := result.Get("records").Array()

	records := make([]map[string]interface{}, 0, len(items))
	for _, item := range items {
		record := map[string]interface{}{
			"行业编码":    item.Get("F001V").String(),
			"证监会行业名称": item.Get("F002V").String(),
			"报告期":     item.Get("ENDDATE").String(),
			"基金覆盖家数":  utils.MustFloat64(item.Get("F003N").String()),
			"行业规模":    utils.MustFloat64(item.Get("F004N").String()),
			"占净资产比例":  utils.MustFloat64(item.Get("F005N").String()),
		}
		records = append(records, record)
	}

	return records, nil
}

// FundReportAssetAllocationCninfo 巨潮资讯-数据中心-专题统计-基金报表-基金资产配置
// https://webapi.cninfo.com.cn/#/thematicStatistics
func FundReportAssetAllocationCninfo() ([]map[string]interface{}, error) {
	url := "https://webapi.cninfo.com.cn/api/sysapi/p_sysapi1114"

	mcode, err := getResCode1()
	if err != nil {
		return nil, fmt.Errorf("生成加密key失败: %w", err)
	}

	headers := map[string]string{
		"Accept":           "*/*",
		"Accept-Enckey":    mcode,
		"Accept-Encoding":  "gzip, deflate",
		"Accept-Language":  "zh-CN,zh;q=0.9,en;q=0.8",
		"Cache-Control":    "no-cache",
		"Content-Length":   "0",
		"Host":             "webapi.cninfo.com.cn",
		"Origin":           "https://webapi.cninfo.com.cn",
		"Pragma":           "no-cache",
		"Proxy-Connection": "keep-alive",
		"Referer":          "https://webapi.cninfo.com.cn/",
		"User-Agent":       "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/93.0.4577.63 Safari/537.36",
		"X-Requested-With": "XMLHttpRequest",
	}

	resp, err := utils.PostWithHeaders(url, nil, headers)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	result := gjson.ParseBytes(resp.Body())
	items := result.Get("records").Array()

	records := make([]map[string]interface{}, 0, len(items))
	for _, item := range items {
		record := map[string]interface{}{
			"报告期":           item.Get("ENDDATE").String(),
			"基金覆盖家数":        utils.MustFloat64(item.Get("F001N").String()),
			"股票权益类占净资产比例":   utils.MustFloat64(item.Get("F006N").String()),
			"债券固定收益类占净资产比例": utils.MustFloat64(item.Get("F007N").String()),
			"现金货币类占净资产比例":   utils.MustFloat64(item.Get("F008N").String()),
			"基金市场净资产规模":     utils.MustFloat64(item.Get("F005N").String()),
		}
		records = append(records, record)
	}

	return records, nil
}

// FundReportCombineCninfo Python源码中不存在此函数，可能是命名错误或已废弃
func FundReportCombineCninfo(symbol string) (map[string]interface{}, error) {
	_ = symbol
	return nil, fmt.Errorf("此函数在Python akshare源码中不存在，可能是命名错误或已废弃")
}
