package futures

import (
	"archive/zip"
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"strings"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/xuri/excelize/v2"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// PositionRank 持仓排名数据
type PositionRank struct {
	Rank                 int     `json:"rank"`                    // 排名
	VolPartyName         string  `json:"vol_party_name"`          // 成交量排序的当前名次会员
	Vol                  float64 `json:"vol"`                     // 该会员成交量
	VolChg               float64 `json:"vol_chg"`                 // 该会员成交量变化量
	LongPartyName        string  `json:"long_party_name"`         // 持多单排序的当前名次会员
	LongOpenInterest     float64 `json:"long_open_interest"`      // 该会员持多单
	LongOpenInterestChg  float64 `json:"long_open_interest_chg"`  // 该会员持多单变化量
	ShortPartyName       string  `json:"short_party_name"`        // 持空单排序的当前名次会员
	ShortOpenInterest    float64 `json:"short_open_interest"`     // 该会员持空单
	ShortOpenInterestChg float64 `json:"short_open_interest_chg"` // 该会员持空单变化量
	Symbol               string  `json:"symbol"`                  // 标的合约
	Variety              string  `json:"variety"`                 // 品种
}

// PositionRankSum 持仓排名汇总数据
type PositionRankSum struct {
	Symbol                    string  `json:"symbol"`                        // 标的合约
	Variety                   string  `json:"variety"`                       // 品种
	Date                      string  `json:"date"`                          // 日期
	VolTop5                   float64 `json:"vol_top5"`                      // 成交量前5会员成交量总和
	VolChgTop5                float64 `json:"vol_chg_top5"`                  // 成交量前5会员成交量变化总和
	LongOpenInterestTop5      float64 `json:"long_open_interest_top5"`       // 持多单前5会员持多单总和
	LongOpenInterestChgTop5   float64 `json:"long_open_interest_chg_top5"`   // 持多单前5会员持多单变化总和
	ShortOpenInterestTop5     float64 `json:"short_open_interest_top5"`      // 持空单前5会员持空单总和
	ShortOpenInterestChgTop5  float64 `json:"short_open_interest_chg_top5"`  // 持空单前5会员持空单变化总和
	VolTop10                  float64 `json:"vol_top10"`                     // 成交量前10
	VolChgTop10               float64 `json:"vol_chg_top10"`                 // 成交量变化前10
	LongOpenInterestTop10     float64 `json:"long_open_interest_top10"`      // 持多单前10
	LongOpenInterestChgTop10  float64 `json:"long_open_interest_chg_top10"`  // 持多单变化前10
	ShortOpenInterestTop10    float64 `json:"short_open_interest_top10"`     // 持空单前10
	ShortOpenInterestChgTop10 float64 `json:"short_open_interest_chg_top10"` // 持空单变化前10
	VolTop15                  float64 `json:"vol_top15"`                     // 成交量前15
	VolChgTop15               float64 `json:"vol_chg_top15"`                 // 成交量变化前15
	LongOpenInterestTop15     float64 `json:"long_open_interest_top15"`      // 持多单前15
	LongOpenInterestChgTop15  float64 `json:"long_open_interest_chg_top15"`  // 持多单变化前15
	ShortOpenInterestTop15    float64 `json:"short_open_interest_top15"`     // 持空单前15
	ShortOpenInterestChgTop15 float64 `json:"short_open_interest_chg_top15"` // 持空单变化前15
	VolTop20                  float64 `json:"vol_top20"`                     // 成交量前20
	VolChgTop20               float64 `json:"vol_chg_top20"`                 // 成交量变化前20
	LongOpenInterestTop20     float64 `json:"long_open_interest_top20"`      // 持多单前20
	LongOpenInterestChgTop20  float64 `json:"long_open_interest_chg_top20"`  // 持多单变化前20
	ShortOpenInterestTop20    float64 `json:"short_open_interest_top20"`     // 持空单前20
	ShortOpenInterestChgTop20 float64 `json:"short_open_interest_chg_top20"` // 持空单变化前20
}

// shfeRankResponse SHFE持仓排名API响应
type shfeRankResponse struct {
	OCursor []struct {
		Rank           int     `json:"RANK"`
		Vol            float64 `json:"CJ1"`
		VolChg         float64 `json:"CJ1_CHG"`
		LongOI         float64 `json:"CJ2"`
		LongOIChg      float64 `json:"CJ2_CHG"`
		ShortOI        float64 `json:"CJ3"`
		ShortOIChg     float64 `json:"CJ3_CHG"`
		VolPartyName   string  `json:"PARTICIPANTABBR1"`
		LongPartyName  string  `json:"PARTICIPANTABBR2"`
		ShortPartyName string  `json:"PARTICIPANTABBR3"`
		ProductName    string  `json:"PRODUCTNAME"`
		Symbol         string  `json:"INSTRUMENTID"`
		ProductSortNo  string  `json:"PRODUCTSORTNO"`
	} `json:"o_cursor"`
}

// FuturesSHFERankTable 上海期货交易所会员成交及持仓排名表
//
// 数据源: https://www.shfe.com.cn/
//
// 参数:
//   - date: 交易日，格式 "20240509"
//   - varsList: 合约品种列表，如 []string{"RB", "AL"}，为空则返回所有
//
// 返回:
//   - map[string][]PositionRank: 按合约分组的持仓排名数据
//   - error: 错误信息
func FuturesSHFERankTable(date string, varsList []string) (map[string][]PositionRank, error) {
	url := fmt.Sprintf("https://www.shfe.com.cn/data/dailydata/%svolrankingdetail.dat", date)
	headers := map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
	}

	resp, err := utils.GetWithHeaders(url, nil, headers)
	if err != nil {
		return nil, fmt.Errorf("请求上期所持仓排名失败: %w", err)
	}

	var apiResp shfeRankResponse
	if err := json.Unmarshal(resp.Body(), &apiResp); err != nil {
		return nil, fmt.Errorf("解析响应失败: %w", err)
	}

	result := make(map[string][]PositionRank)
	for _, item := range apiResp.OCursor {
		if item.Rank <= 0 {
			continue
		}

		symbol := strings.TrimSpace(item.Symbol)
		variety := SymbolVarieties(symbol)

		// 检查是否在筛选列表中
		if len(varsList) > 0 {
			found := false
			for _, v := range varsList {
				if v == variety {
					found = true
					break
				}
			}
			if !found {
				continue
			}
		}

		rank := PositionRank{
			Rank:                 item.Rank,
			VolPartyName:         strings.TrimSpace(item.VolPartyName),
			Vol:                  item.Vol,
			VolChg:               item.VolChg,
			LongPartyName:        strings.TrimSpace(item.LongPartyName),
			LongOpenInterest:     item.LongOI,
			LongOpenInterestChg:  item.LongOIChg,
			ShortPartyName:       strings.TrimSpace(item.ShortPartyName),
			ShortOpenInterest:    item.ShortOI,
			ShortOpenInterestChg: item.ShortOIChg,
			Symbol:               strings.ToUpper(symbol),
			Variety:              variety,
		}

		result[symbol] = append(result[symbol], rank)
	}

	return result, nil
}

// FuturesCZCERankTable 郑州商品交易所前20会员持仓排名数据
//
// 数据源: https://www.czce.com.cn/cn/jysj/ccpm/H077003004index_1.htm
//
// 参数:
//   - date: 交易日，格式 "20240509"
//
// 返回:
//   - map[string][]PositionRank: 按合约分组的持仓排名数据
//   - error: 错误信息
func FuturesCZCERankTable(date string) (map[string][]PositionRank, error) {
	dateInt := 0
	fmt.Sscanf(date, "%d", &dateInt)

	var url string
	if dateInt > 20251101 {
		url = fmt.Sprintf("http://www.czce.com.cn/cn/DFSStaticFiles/Future/%s/%s/FutureDataHolding.xlsx", date[:4], date)
	} else {
		url = fmt.Sprintf("http://www.czce.com.cn/cn/DFSStaticFiles/Future/%s/%s/FutureDataHolding.xls", date[:4], date)
	}

	headers := map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
	}

	resp, err := utils.GetWithHeaders(url, nil, headers)
	if err != nil {
		return nil, fmt.Errorf("请求郑商所持仓排名失败: %w", err)
	}

	f, err := excelize.OpenReader(bytes.NewReader(resp.Body()))
	if err != nil {
		return nil, fmt.Errorf("解析Excel失败: %w", err)
	}
	defer f.Close()

	sheets := f.GetSheetList()
	if len(sheets) == 0 {
		return nil, fmt.Errorf("Excel文件没有工作表")
	}

	rows, err := f.GetRows(sheets[0])
	if err != nil {
		return nil, fmt.Errorf("读取工作表失败: %w", err)
	}

	result := make(map[string][]PositionRank)
	symbolRe := regexp.MustCompile(`[0-9a-zA-Z_]+`)
	varietyRe := regexp.MustCompile(`[a-zA-Z_]+`)

	// 查找品种分隔点
	var symbolIndexes []int
	var symbolList []string

	for i, row := range rows {
		if len(row) > 0 && strings.Contains(row[0], "合计") {
			symbolIndexes = append(symbolIndexes, i+1)
		}
	}

	// 插入起始点并移除最后一个
	if len(symbolIndexes) > 0 {
		symbolIndexes = append([]int{0}, symbolIndexes[:len(symbolIndexes)-1]...)
	}

	// 提取品种代码
	for _, idx := range symbolIndexes {
		if idx < len(rows) && len(rows[idx]) > 0 {
			matches := symbolRe.FindStringSubmatch(strings.Split(rows[idx][0], " ")[0])
			if len(matches) > 0 {
				symbolList = append(symbolList, matches[0])
			}
		}
	}

	// 解析每个品种的数据
	for i := 0; i < len(symbolIndexes)-1 && i < len(symbolList); i++ {
		startIdx := symbolIndexes[i] + 2
		endIdx := symbolIndexes[i+1] - 1
		symbol := symbolList[i]

		varietyMatches := varietyRe.FindStringSubmatch(symbol)
		variety := ""
		if len(varietyMatches) > 0 {
			variety = varietyMatches[0]
		}

		for j := startIdx; j < endIdx && j < len(rows); j++ {
			row := rows[j]
			if len(row) < 10 {
				continue
			}

			rank := int(utils.MustParseFloat(row[0]))
			if rank <= 0 {
				continue
			}

			result[symbol] = append(result[symbol], PositionRank{
				Rank:                 rank,
				VolPartyName:         row[1],
				Vol:                  parseCOTFloat(row[2]),
				VolChg:               parseCOTFloat(row[3]),
				LongPartyName:        row[4],
				LongOpenInterest:     parseCOTFloat(row[5]),
				LongOpenInterestChg:  parseCOTFloat(row[6]),
				ShortPartyName:       row[7],
				ShortOpenInterest:    parseCOTFloat(row[8]),
				ShortOpenInterestChg: parseCOTFloat(row[9]),
				Symbol:               symbol,
				Variety:              variety,
			})
		}
	}

	return result, nil
}

// FuturesCFFEXRankTable 中国金融期货交易所前20会员持仓排名数据
//
// 数据源: http://www.cffex.com.cn/ccpm/
//
// 参数:
//   - date: 交易日，格式 "20240509"
//   - varsList: 合约品种列表，如 []string{"IF", "IC"}，为空则返回所有
//
// 返回:
//   - map[string][]PositionRank: 按合约分组的持仓排名数据
//   - error: 错误信息
func FuturesCFFEXRankTable(date string, varsList []string) (map[string][]PositionRank, error) {
	if len(varsList) == 0 {
		varsList = MarketExchangeSymbols["cffex"]
	}

	headers := map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
	}

	result := make(map[string][]PositionRank)

	for _, variety := range varsList {
		// 检查是否是中金所品种
		found := false
		for _, v := range MarketExchangeSymbols["cffex"] {
			if v == variety {
				found = true
				break
			}
		}
		if !found {
			continue
		}

		url := fmt.Sprintf("http://www.cffex.com.cn/sj/ccpm/%s/%s/%s_1.csv", date[:6], date[6:8], variety)

		resp, err := utils.GetWithHeaders(url, nil, headers)
		if err != nil {
			continue
		}

		// 解码GBK
		decoder := simplifiedchinese.GBK.NewDecoder()
		reader := transform.NewReader(bytes.NewReader(resp.Body()), decoder)
		content, err := io.ReadAll(reader)
		if err != nil {
			continue
		}

		// 解析CSV
		csvReader := csv.NewReader(bytes.NewReader(content))
		csvReader.FieldsPerRecord = -1
		records, err := csvReader.ReadAll()
		if err != nil {
			continue
		}

		// 跳过表头，查找数据行
		dataStart := 0
		for i, row := range records {
			if len(row) > 0 && strings.Contains(row[0], "交易日") {
				dataStart = i + 1
				break
			}
		}

		for i := dataStart; i < len(records); i++ {
			row := records[i]
			if len(row) < 11 {
				continue
			}

			symbol := strings.TrimSpace(row[0])
			if symbol == "" || symbol == "合约" {
				continue
			}

			rank := int(utils.MustParseFloat(row[1]))
			if rank <= 0 {
				continue
			}

			result[symbol] = append(result[symbol], PositionRank{
				Rank:                 rank,
				VolPartyName:         strings.TrimSpace(row[2]),
				Vol:                  parseCOTFloat(row[3]),
				VolChg:               parseCOTFloat(row[4]),
				LongPartyName:        strings.TrimSpace(row[5]),
				LongOpenInterest:     parseCOTFloat(row[6]),
				LongOpenInterestChg:  parseCOTFloat(row[7]),
				ShortPartyName:       strings.TrimSpace(row[8]),
				ShortOpenInterest:    parseCOTFloat(row[9]),
				ShortOpenInterestChg: parseCOTFloat(row[10]),
				Symbol:               symbol,
				Variety:              variety,
			})
		}
	}

	return result, nil
}

// FuturesDCEPositionRank 大连商品交易所-每日持仓排名-具体合约
//
// 数据源: http://www.dce.com.cn/dalianshangpin/xqsj/tjsj26/rtj/rcjccpm/index.html
//
// 参数:
//   - date: 交易日，格式 "20240509"
//   - varsList: 品种列表，如 []string{"M", "Y"}，为空则返回所有
//
// 返回:
//   - map[string][]PositionRank: 按合约分组的持仓排名数据
//   - error: 错误信息
func FuturesDCEPositionRank(date string, varsList []string) (map[string][]PositionRank, error) {
	url := "http://www.dce.com.cn/dcereport/publicweb/dailystat/memberDealPosi/batchDownload"

	payload := map[string]any{
		"tradeDate":  date,
		"varietyId":  "a",
		"contractId": "a2601",
		"tradeType":  "1",
		"lang":       "zh",
	}

	headers := map[string]string{
		"User-Agent":   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
		"Content-Type": "application/json",
	}

	resp, err := utils.PostJSONWithHeaders(url, nil, payload, headers)
	if err != nil {
		return nil, fmt.Errorf("请求大商所持仓排名失败: %w", err)
	}

	// 解析ZIP文件
	zipReader, err := zip.NewReader(bytes.NewReader(resp.Body()), int64(len(resp.Body())))
	if err != nil {
		return nil, fmt.Errorf("解析ZIP文件失败: %w", err)
	}

	result := make(map[string][]PositionRank)
	symbolRe := regexp.MustCompile(`\d`)
	varietyRe := regexp.MustCompile(`[a-zA-Z]+`)

	for _, file := range zipReader.File {
		fileName := file.Name
		if !strings.HasPrefix(fileName, date) {
			continue
		}

		// 提取合约代码
		parts := strings.Split(fileName, "_")
		if len(parts) < 2 {
			continue
		}
		symbol := parts[1]

		// 检查是否在筛选列表中
		if len(varsList) > 0 {
			varietyMatches := varietyRe.FindStringSubmatch(symbol)
			if len(varietyMatches) == 0 {
				continue
			}
			variety := strings.ToUpper(varietyMatches[0])
			found := false
			for _, v := range varsList {
				if v == variety {
					found = true
					break
				}
			}
			if !found {
				continue
			}
		}

		// 读取文件内容
		rc, err := file.Open()
		if err != nil {
			continue
		}

		content, err := io.ReadAll(rc)
		rc.Close()
		if err != nil {
			continue
		}

		// 尝试解析
		ranks := parseDCERankData(content, symbol)
		if len(ranks) > 0 {
			result[symbol] = ranks
		}
	}

	// 过滤不在列表中的品种
	if len(varsList) > 0 {
		for key := range result {
			variety := symbolRe.ReplaceAllString(key, "")
			found := false
			for _, v := range varsList {
				if strings.EqualFold(v, variety) {
					found = true
					break
				}
			}
			if !found {
				delete(result, key)
			}
		}
	}

	return result, nil
}

// parseDCERankData 解析大商所持仓排名数据
func parseDCERankData(content []byte, symbol string) []PositionRank {
	// 尝试UTF-8解码
	lines := strings.Split(string(content), "\n")

	// 如果是乱码，尝试GBK解码
	if len(lines) > 0 && strings.Contains(lines[0], "�") {
		decoder := simplifiedchinese.GBK.NewDecoder()
		reader := transform.NewReader(bytes.NewReader(content), decoder)
		decoded, err := io.ReadAll(reader)
		if err == nil {
			lines = strings.Split(string(decoded), "\n")
		}
	}

	// 查找数据开始位置
	var startIndexes []int
	var endIndexes []int

	for i, line := range lines {
		if strings.HasPrefix(strings.TrimSpace(line), "名次") {
			startIndexes = append(startIndexes, i)
		}
		if strings.Contains(line, "总计") || strings.Contains(line, "合计") {
			endIndexes = append(endIndexes, i)
		}
	}

	if len(startIndexes) < 3 || len(endIndexes) < 3 {
		return nil
	}

	// 解析三个部分：成交量、买持仓、卖持仓
	var volData, longData, shortData [][]string

	for i := startIndexes[0] + 1; i < endIndexes[0] && i < len(lines); i++ {
		fields := strings.Fields(lines[i])
		if len(fields) >= 4 {
			volData = append(volData, fields)
		}
	}

	for i := startIndexes[1] + 1; i < endIndexes[1] && i < len(lines); i++ {
		fields := strings.Fields(lines[i])
		if len(fields) >= 4 {
			longData = append(longData, fields)
		}
	}

	for i := startIndexes[2] + 1; i < endIndexes[2] && i < len(lines); i++ {
		fields := strings.Fields(lines[i])
		if len(fields) >= 4 {
			shortData = append(shortData, fields)
		}
	}

	// 合并数据
	maxLen := len(volData)
	if len(longData) < maxLen {
		maxLen = len(longData)
	}
	if len(shortData) < maxLen {
		maxLen = len(shortData)
	}

	variety := regexp.MustCompile(`[a-zA-Z]+`).FindString(symbol)

	var result []PositionRank
	for i := 0; i < maxLen; i++ {
		rank := PositionRank{
			Rank:    i + 1,
			Symbol:  strings.ToUpper(symbol),
			Variety: strings.ToUpper(variety),
		}

		if i < len(volData) && len(volData[i]) >= 4 {
			rank.VolPartyName = volData[i][1]
			rank.Vol = parseCOTFloat(volData[i][2])
			rank.VolChg = parseCOTFloat(volData[i][3])
		}

		if i < len(longData) && len(longData[i]) >= 4 {
			rank.LongPartyName = longData[i][1]
			rank.LongOpenInterest = parseCOTFloat(longData[i][2])
			rank.LongOpenInterestChg = parseCOTFloat(longData[i][3])
		}

		if i < len(shortData) && len(shortData[i]) >= 4 {
			rank.ShortPartyName = shortData[i][1]
			rank.ShortOpenInterest = parseCOTFloat(shortData[i][2])
			rank.ShortOpenInterestChg = parseCOTFloat(shortData[i][3])
		}

		result = append(result, rank)
	}

	return result
}

// gfexVarsResponse GFEX品种列表响应
type gfexVarsResponse struct {
	Data []struct {
		VarietyID string `json:"varietyId"`
	} `json:"data"`
}

// gfexContractResponse GFEX合约列表响应
type gfexContractResponse struct {
	Data []struct {
		ContractID string `json:"contract_id"`
	} `json:"data"`
}

// gfexRankResponse GFEX持仓排名响应
type gfexRankResponse struct {
	Data []struct {
		Abbr        string `json:"abbr"`
		TodayQty    any    `json:"todayQty"`
		TodayQtyChg any    `json:"todayQtyChg"`
		QtySub      any    `json:"qtySub"`
	} `json:"data"`
}

// FuturesGFEXPositionRank 广州期货交易所-日成交持仓排名
//
// 数据源: http://www.gfex.com.cn/gfex/rcjccpm/hqsj_tjsj.shtml
//
// 参数:
//   - date: 交易日，格式 "20240509"
//   - varsList: 品种列表，如 []string{"SI", "LC"}，为空则返回所有
//
// 返回:
//   - map[string][]PositionRank: 按合约分组的持仓排名数据
//   - error: 错误信息
func FuturesGFEXPositionRank(date string, varsList []string) (map[string][]PositionRank, error) {
	headers := map[string]string{
		"User-Agent":   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
		"Content-Type": "application/x-www-form-urlencoded",
	}

	// 获取品种列表
	if len(varsList) == 0 {
		varsURL := "http://www.gfex.com.cn/u/interfacesWebVariety/loadList"
		resp, err := utils.PostFormWithHeaders(varsURL, nil, headers)
		if err == nil {
			var varsResp gfexVarsResponse
			if json.Unmarshal(resp.Body(), &varsResp) == nil {
				for _, item := range varsResp.Data {
					varsList = append(varsList, item.VarietyID)
				}
			}
		}
	}

	result := make(map[string][]PositionRank)

	for _, variety := range varsList {
		varietyLower := strings.ToLower(variety)

		// 获取合约列表
		contractURL := "http://www.gfex.com.cn/u/interfacesWebTiMemberDealPosiQuotes/loadListContract_id"
		contractPayload := map[string]string{
			"variety":    varietyLower,
			"trade_date": date,
		}

		resp, err := utils.PostFormWithHeaders(contractURL, contractPayload, headers)
		if err != nil {
			continue
		}

		var contractResp gfexContractResponse
		if err := json.Unmarshal(resp.Body(), &contractResp); err != nil {
			continue
		}

		// 获取每个合约的持仓排名
		for _, contract := range contractResp.Data {
			contractID := contract.ContractID
			if contractID == "" {
				continue
			}

			ranks, err := getGFEXContractRank(varietyLower, contractID, date, headers)
			if err != nil {
				continue
			}

			if len(ranks) > 0 {
				result[contractID] = ranks
			}
		}
	}

	return result, nil
}

// getGFEXContractRank 获取广期所单个合约的持仓排名
func getGFEXContractRank(variety, contractID, date string, headers map[string]string) ([]PositionRank, error) {
	url := "http://www.gfex.com.cn/u/interfacesWebTiMemberDealPosiQuotes/loadList"

	var volData, longData, shortData []struct {
		Name string
		Qty  float64
		Chg  float64
	}

	// 获取三种类型的数据
	for dataType := 1; dataType <= 3; dataType++ {
		payload := map[string]string{
			"trade_date":  date,
			"trade_type":  "0",
			"variety":     variety,
			"contract_id": contractID,
			"data_type":   fmt.Sprintf("%d", dataType),
		}

		resp, err := utils.PostFormWithHeaders(url, payload, headers)
		if err != nil {
			continue
		}

		var rankResp gfexRankResponse
		if err := json.Unmarshal(resp.Body(), &rankResp); err != nil {
			continue
		}

		for _, item := range rankResp.Data {
			qty, _ := utils.ToFloat64(item.TodayQty)
			chg, _ := utils.ToFloat64(item.TodayQtyChg)
			if chg == 0 {
				chg, _ = utils.ToFloat64(item.QtySub)
			}

			data := struct {
				Name string
				Qty  float64
				Chg  float64
			}{
				Name: item.Abbr,
				Qty:  qty,
				Chg:  chg,
			}

			switch dataType {
			case 1:
				volData = append(volData, data)
			case 2:
				longData = append(longData, data)
			case 3:
				shortData = append(shortData, data)
			}
		}
	}

	// 合并数据
	maxLen := len(volData)
	if len(longData) > maxLen {
		maxLen = len(longData)
	}
	if len(shortData) > maxLen {
		maxLen = len(shortData)
	}

	// 移除最后一行（合计行）
	if maxLen > 0 {
		maxLen--
	}

	var result []PositionRank
	for i := 0; i < maxLen; i++ {
		rank := PositionRank{
			Rank:    i + 1,
			Symbol:  strings.ToUpper(contractID),
			Variety: strings.ToUpper(variety),
		}

		if i < len(volData) {
			rank.VolPartyName = volData[i].Name
			rank.Vol = volData[i].Qty
			rank.VolChg = volData[i].Chg
		}

		if i < len(longData) {
			rank.LongPartyName = longData[i].Name
			rank.LongOpenInterest = longData[i].Qty
			rank.LongOpenInterestChg = longData[i].Chg
		}

		if i < len(shortData) {
			rank.ShortPartyName = shortData[i].Name
			rank.ShortOpenInterest = shortData[i].Qty
			rank.ShortOpenInterestChg = shortData[i].Chg
		}

		result = append(result, rank)
	}

	return result, nil
}

// parseCOTFloat 解析可能包含逗号的浮点数
func parseCOTFloat(s string) float64 {
	s = strings.TrimSpace(s)
	s = strings.ReplaceAll(s, ",", "")
	s = strings.ReplaceAll(s, "-", "0")
	return utils.MustParseFloat(s)
}

// FuturesRankSum 获取指定交易日五个期货交易所前5/10/15/20会员持仓排名汇总数据
//
// 参数:
//   - date: 交易日，格式 "20240509"
//   - varsList: 合约品种列表，为空则返回所有
//
// 返回:
//   - []PositionRankSum: 持仓排名汇总数据
//   - error: 错误信息
func FuturesRankSum(date string, varsList []string) ([]PositionRankSum, error) {
	if len(varsList) == 0 {
		varsList = ContractSymbols
	}

	bigDict := make(map[string][]PositionRank)

	// 按交易所分类品种
	dceVars := filterByExchange(varsList, "dce")
	shfeVars := filterByExchange(varsList, "shfe")
	czceVars := filterByExchange(varsList, "czce")
	cffexVars := filterByExchange(varsList, "cffex")
	gfexVars := filterByExchange(varsList, "gfex")

	// 获取各交易所数据
	if len(dceVars) > 0 {
		data, err := FuturesDCEPositionRank(date, dceVars)
		if err == nil {
			for k, v := range data {
				bigDict[k] = v
			}
		}
	}

	if len(shfeVars) > 0 {
		data, err := FuturesSHFERankTable(date, shfeVars)
		if err == nil {
			for k, v := range data {
				bigDict[k] = v
			}
		}
	}

	if len(czceVars) > 0 {
		data, err := FuturesCZCERankTable(date)
		if err == nil {
			for k, v := range data {
				bigDict[k] = v
			}
		}
	}

	if len(cffexVars) > 0 {
		data, err := FuturesCFFEXRankTable(date, cffexVars)
		if err == nil {
			for k, v := range data {
				bigDict[k] = v
			}
		}
	}

	if len(gfexVars) > 0 {
		data, err := FuturesGFEXPositionRank(date, gfexVars)
		if err == nil {
			for k, v := range data {
				bigDict[k] = v
			}
		}
	}

	// 计算汇总数据
	var result []PositionRankSum

	for symbol, ranks := range bigDict {
		variety := SymbolVarieties(symbol)

		// 检查是否在筛选列表中
		found := false
		for _, v := range varsList {
			if v == variety {
				found = true
				break
			}
		}
		if !found {
			continue
		}

		sum := calculateRankSum(ranks, symbol, variety, date)
		result = append(result, sum)
	}

	return result, nil
}

// filterByExchange 按交易所筛选品种
func filterByExchange(varsList []string, exchange string) []string {
	exchangeSymbols := MarketExchangeSymbols[exchange]
	var result []string

	for _, v := range varsList {
		for _, e := range exchangeSymbols {
			if v == e {
				result = append(result, v)
				break
			}
		}
	}

	return result
}

// calculateRankSum 计算持仓排名汇总
func calculateRankSum(ranks []PositionRank, symbol, variety, date string) PositionRankSum {
	sum := PositionRankSum{
		Symbol:  symbol,
		Variety: variety,
		Date:    date,
	}

	for _, r := range ranks {
		if r.Rank <= 5 {
			sum.VolTop5 += r.Vol
			sum.VolChgTop5 += r.VolChg
			sum.LongOpenInterestTop5 += r.LongOpenInterest
			sum.LongOpenInterestChgTop5 += r.LongOpenInterestChg
			sum.ShortOpenInterestTop5 += r.ShortOpenInterest
			sum.ShortOpenInterestChgTop5 += r.ShortOpenInterestChg
		}
		if r.Rank <= 10 {
			sum.VolTop10 += r.Vol
			sum.VolChgTop10 += r.VolChg
			sum.LongOpenInterestTop10 += r.LongOpenInterest
			sum.LongOpenInterestChgTop10 += r.LongOpenInterestChg
			sum.ShortOpenInterestTop10 += r.ShortOpenInterest
			sum.ShortOpenInterestChgTop10 += r.ShortOpenInterestChg
		}
		if r.Rank <= 15 {
			sum.VolTop15 += r.Vol
			sum.VolChgTop15 += r.VolChg
			sum.LongOpenInterestTop15 += r.LongOpenInterest
			sum.LongOpenInterestChgTop15 += r.LongOpenInterestChg
			sum.ShortOpenInterestTop15 += r.ShortOpenInterest
			sum.ShortOpenInterestChgTop15 += r.ShortOpenInterestChg
		}
		if r.Rank <= 20 {
			sum.VolTop20 += r.Vol
			sum.VolChgTop20 += r.VolChg
			sum.LongOpenInterestTop20 += r.LongOpenInterest
			sum.LongOpenInterestChgTop20 += r.LongOpenInterestChg
			sum.ShortOpenInterestTop20 += r.ShortOpenInterest
			sum.ShortOpenInterestChgTop20 += r.ShortOpenInterestChg
		}
	}

	return sum
}
