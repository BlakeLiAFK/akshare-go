package futures_derivative

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"strings"

	"github.com/BlakeLiAFK/akshare/utils"
)

// FuturesContractInfoCffex 中国金融期货交易所-数据-交易参数
// 参数: date 查询日期，格式 "20240228"
// 返回: 交易参数汇总数据
func FuturesContractInfoCffex(date string) ([]map[string]interface{}, error) {
	if date == "" {
		date = "20240228"
	}

	// 构建URL: http://www.cffex.com.cn/sj/jycs/202402/28/index.xml
	url := fmt.Sprintf("http://www.cffex.com.cn/sj/jycs/%s/%s/index.xml",
		date[:6], date[6:])

	headers := map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
	}

	resp, err := utils.GetWithHeaders(url, nil, headers)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	// 解析XML
	type INDEX struct {
		TRADING_DAY      string `xml:"TRADING_DAY"`
		PRODUCT_ID       string `xml:"PRODUCT_ID"`
		INSTRUMENT_ID    string `xml:"INSTRUMENT_ID"`
		INSTRUMENT_MONTH string `xml:"INSTRUMENT_MONTH"`
		BASIS_PRICE      string `xml:"BASIS_PRICE"`
		OPEN_DATE        string `xml:"OPEN_DATE"`
		END_TRADING_DAY  string `xml:"END_TRADING_DAY"`
		UPPER_VALUE      string `xml:"UPPER_VALUE"`
		LOWER_VALUE      string `xml:"LOWER_VALUE"`
		UPPERLIMITPRICE  string `xml:"UPPERLIMITPRICE"`
		LOWERLIMITPRICE  string `xml:"LOWERLIMITPRICE"`
		LONG_LIMIT       string `xml:"LONG_LIMIT"`
	}

	type Data struct {
		XMLName xml.Name `xml:"data"`
		Records []INDEX  `xml:"INDEX"`
	}

	var data Data
	if err := xml.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("解析XML失败: %w", err)
	}

	records := make([]map[string]interface{}, 0)
	for _, item := range data.Records {
		record := map[string]interface{}{
			"合约代码":  item.INSTRUMENT_ID,
			"合约月份":  item.INSTRUMENT_MONTH,
			"挂盘基准价": utils.MustFloat64(item.BASIS_PRICE),
			"上市日":   item.OPEN_DATE,
			"最后交易日": item.END_TRADING_DAY,
			"涨停板幅度": item.UPPER_VALUE,
			"跌停板幅度": item.LOWER_VALUE,
			"涨停板价位": utils.MustFloat64(item.UPPERLIMITPRICE),
			"跌停板价位": utils.MustFloat64(item.LOWERLIMITPRICE),
			"持仓限额":  utils.MustInt64(item.LONG_LIMIT),
			"品种":    item.PRODUCT_ID,
			"查询交易日": item.TRADING_DAY,
		}
		records = append(records, record)
	}

	return records, nil
}

// FuturesContractInfoCzce 郑州商品交易所-交易数据-参考数据
// 参数: date 查询日期，格式 "20240228"
// 返回: 交易参数汇总数据
func FuturesContractInfoCzce(date string) ([]map[string]interface{}, error) {
	if date == "" {
		date = "20240228"
	}

	// 构建URL: http://www.czce.com.cn/cn/DFSStaticFiles/Future/2024/20240228/FutureDataReferenceData.xml
	url := fmt.Sprintf("http://www.czce.com.cn/cn/DFSStaticFiles/Future/%s/%s/FutureDataReferenceData.xml",
		date[:4], date)

	headers := map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
		"Host":       "www.czce.com.cn",
	}

	resp, err := utils.GetWithHeaders(url, nil, headers)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	// 解析XML
	type Contract struct {
		Name              string `xml:"Name"`
		CtrCd             string `xml:"CtrCd"`
		PrdCd             string `xml:"PrdCd"`
		TckSz             string `xml:"TckSz"`
		CtrSz             string `xml:"CtrSz"`
		FrstTrdDt         string `xml:"FrstTrdDt"`
		LstTrdDt          string `xml:"LstTrdDt"`
		DlvryMnth         string `xml:"DlvryMnth"`
		Margin            string `xml:"Margin"`
		PxLim             string `xml:"PxLim"`
		TrdFee            string `xml:"TrdFee"`
		IntraDayTrdFee    string `xml:"IntraDayTrdFee"`
		DlvryFee          string `xml:"DlvryFee"`
		LstDlvryDt        string `xml:"LstDlvryDt"`
		LstDlvryDtBoard   string `xml:"LstDlvryDtBoard"`
		FeeCollectionType string `xml:"FeeCollectionType"`
		MnthPosLmt        string `xml:"MnthPosLmt"`
	}

	type Data struct {
		XMLName   xml.Name   `xml:"ContractBaseInfo"`
		Contracts []Contract `xml:"Contract"`
	}

	var data Data
	if err := xml.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("解析XML失败: %w", err)
	}

	records := make([]map[string]interface{}, 0)
	for _, item := range data.Contracts {
		record := map[string]interface{}{
			"产品名称":           item.Name,
			"合约代码":           item.CtrCd,
			"产品代码":           item.PrdCd,
			"最小变动价位":         item.TckSz,
			"交易单位":           item.CtrSz,
			"第一交易日":          item.FrstTrdDt,
			"最后交易日":          item.LstTrdDt,
			"合约交割月份":         item.DlvryMnth,
			"交易保证金率":         item.Margin,
			"涨跌停板":           item.PxLim,
			"交易手续费":          utils.MustFloat64(item.TrdFee),
			"平今仓手续费":         utils.MustFloat64(item.IntraDayTrdFee),
			"交割手续费":          utils.MustFloat64(item.DlvryFee),
			"最后交割日":          item.LstDlvryDt,
			"车（船）板最后交割日":     item.LstDlvryDtBoard,
			"手续费收取方式":        item.FeeCollectionType,
			"日持仓限额期货公司会员不限仓": item.MnthPosLmt,
		}
		records = append(records, record)
	}

	return records, nil
}

// FuturesContractInfoDce 大连商品交易所-业务/服务-业务参数-交易参数-合约信息查询
// 返回: 交易参数汇总数据
func FuturesContractInfoDce() ([]map[string]interface{}, error) {
	url := "http://www.dce.com.cn/dcereport/publicweb/tradepara/contractInfo"
	payload := map[string]interface{}{
		"lang":      "zh",
		"tradeType": "1",
		"varietyId": "all",
	}

	resp, err := utils.PostJSON(url, payload)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	records := make([]map[string]interface{}, 0)
	// 使用 gjson 解析
	dataArray := resp.Get("data").Array()
	for _, item := range dataArray {
		record := map[string]interface{}{
			"品种名称":   item.Get("variety").String(),
			"合约":     item.Get("contractId").String(),
			"交易单位":   utils.MustInt64(item.Get("unit").String()),
			"最小变动价位": utils.MustFloat64(item.Get("tick").String()),
			"开始交易日":  item.Get("startTradeDate").String(),
			"最后交易日":  item.Get("endTradeDate").String(),
			"最后交割日":  item.Get("endDeliveryDate").String(),
		}
		records = append(records, record)
	}

	return records, nil
}

// FuturesContractInfoGfex 广州期货交易所-业务/服务-合约信息
// 返回: 交易参数汇总数据
func FuturesContractInfoGfex() ([]map[string]interface{}, error) {
	url := "http://www.gfex.com.cn/u/interfacesWebTtQueryContractInfo/loadList"
	params := map[string]interface{}{
		"variety":    "",
		"trade_type": "0",
	}

	headers := map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
	}

	resp, err := utils.PostWithHeaders(url, params, headers)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	var result struct {
		Data []struct {
			TradeType        string `json:"tradeType"`
			Variety          string `json:"variety"`
			ContractID       string `json:"contractId"`
			Unit             string `json:"unit"`
			Tick             string `json:"tick"`
			StartTradeDate   string `json:"startTradeDate"`
			EndTradeDate     string `json:"endTradeDate"`
			EndDeliveryDate0 string `json:"endDeliveryDate0"`
		} `json:"data"`
	}

	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return nil, fmt.Errorf("解析JSON失败: %w", err)
	}

	records := make([]map[string]interface{}, 0)
	for _, item := range result.Data {
		record := map[string]interface{}{
			"品种":     item.Variety,
			"合约代码":   item.ContractID,
			"交易单位":   utils.MustInt64(item.Unit),
			"最小变动单位": utils.MustFloat64(item.Tick),
			"开始交易日":  item.StartTradeDate,
			"最后交易日":  item.EndTradeDate,
			"最后交割日":  item.EndDeliveryDate0,
		}
		records = append(records, record)
	}

	return records, nil
}

// FuturesContractInfoIne 上海国际能源交易中心-业务指南-交易参数汇总(期货)
// 参数: date 查询日期，格式 "20241129"
// 返回: 交易参数汇总数据
func FuturesContractInfoIne(date string) ([]map[string]interface{}, error) {
	if date == "" {
		date = "20241129"
	}

	url := fmt.Sprintf("https://www.ine.cn/data/busiparamdata/future/ContractBaseInfo%s.dat", date)
	params := map[string]string{
		"rnd": "0.8312696798757147",
	}

	headers := map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
	}

	resp, err := utils.GetWithHeaders(url, params, headers)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	var result struct {
		ContractBaseInfo []struct {
			BASISPRICE     string `json:"BASISPRICE"`
			ENDDELIVDATE   string `json:"ENDDELIVDATE"`
			EXPIREDATE     string `json:"EXPIREDATE"`
			INSTRUMENTID   string `json:"INSTRUMENTID"`
			OPENDATE       string `json:"OPENDATE"`
			STARTDELIVDATE string `json:"STARTDELIVDATE"`
			TRADINGDAY     string `json:"TRADINGDAY"`
		} `json:"ContractBaseInfo"`
	}

	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return nil, fmt.Errorf("解析JSON失败: %w", err)
	}

	records := make([]map[string]interface{}, 0)
	for _, item := range result.ContractBaseInfo {
		record := map[string]interface{}{
			"合约代码":  item.INSTRUMENTID,
			"上市日":   item.OPENDATE,
			"到期日":   item.EXPIREDATE,
			"开始交割日": item.STARTDELIVDATE,
			"最后交割日": item.ENDDELIVDATE,
			"挂牌基准价": utils.MustFloat64(item.BASISPRICE),
			"交易日":   item.TRADINGDAY,
		}
		records = append(records, record)
	}

	return records, nil
}

// FuturesContractInfoShfe 上海期货交易所-交易所服务-业务数据-交易参数汇总查询
// 参数: date 查询日期，格式 "20240513"
// 返回: 交易参数汇总数据
func FuturesContractInfoShfe(date string) ([]map[string]interface{}, error) {
	if date == "" {
		date = "20240513"
	}

	url := fmt.Sprintf("https://www.shfe.com.cn/data/busiparamdata/future/ContractBaseInfo%s.dat", date)

	headers := map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
	}

	resp, err := utils.GetWithHeaders(url, nil, headers)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	var result struct {
		ContractBaseInfo []struct {
			BASISPRICE     string `json:"BASISPRICE"`
			ENDDELIVDATE   string `json:"ENDDELIVDATE"`
			EXPIREDATE     string `json:"EXPIREDATE"`
			INSTRUMENTID   string `json:"INSTRUMENTID"`
			OPENDATE       string `json:"OPENDATE"`
			STARTDELIVDATE string `json:"STARTDELIVDATE"`
			TRADINGDAY     string `json:"TRADINGDAY"`
		} `json:"ContractBaseInfo"`
		UpdateDate string `json:"update_date"`
	}

	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		// 尝试处理可能的编码问题
		text := strings.TrimSpace(string(resp.Body()))
		if err := json.Unmarshal([]byte(text), &result); err != nil {
			return nil, fmt.Errorf("解析JSON失败: %w", err)
		}
	}

	records := make([]map[string]interface{}, 0)
	for _, item := range result.ContractBaseInfo {
		record := map[string]interface{}{
			"合约代码":  item.INSTRUMENTID,
			"上市日":   item.OPENDATE,
			"到期日":   item.EXPIREDATE,
			"开始交割日": item.STARTDELIVDATE,
			"最后交割日": item.ENDDELIVDATE,
			"挂牌基准价": utils.MustFloat64(item.BASISPRICE),
			"交易日":   item.TRADINGDAY,
			"更新时间":  result.UpdateDate,
		}
		records = append(records, record)
	}

	return records, nil
}
