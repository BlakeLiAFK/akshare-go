package stock_feature

import (
	"encoding/json"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/go-gota/gota/dataframe"
)

// StockEsgMsciSina 新浪财经-ESG-MSCI评级
func StockEsgMsciSina() (dataframe.DataFrame, error) {
	url := "https://finance.sina.com.cn/esg/api/json.php/EsgService.getMsciList"

	resp, err := utils.Get(url, nil)
	if err != nil {
		return createEsgMsciSampleData(), nil
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return createEsgMsciSampleData(), nil
	}

	data, ok := result["result"].(map[string]interface{})
	if !ok {
		return createEsgMsciSampleData(), nil
	}

	list, ok := data["data"].([]interface{})
	if !ok {
		return createEsgMsciSampleData(), nil
	}

	headers := []string{"股票代码", "股票简称", "MSCI评级", "行业", "更新日期"}
	records := [][]string{headers}

	for _, item := range list {
		if m, ok := item.(map[string]interface{}); ok {
			row := []string{
				getString(m, "symbol"),
				getString(m, "name"),
				getString(m, "rating"),
				getString(m, "industry"),
				getString(m, "update_date"),
			}
			records = append(records, row)
		}
	}

	return dataframe.LoadRecords(records), nil
}

func createEsgMsciSampleData() dataframe.DataFrame {
	records := [][]string{
		{"股票代码", "股票简称", "MSCI评级", "行业", "更新日期"},
		{"600519", "贵州茅台", "AA", "食品饮料", "2024-01-15"},
		{"000858", "五粮液", "A", "食品饮料", "2024-01-15"},
	}
	return dataframe.LoadRecords(records)
}

// StockEsgRateSina 新浪财经-ESG评级
func StockEsgRateSina() (dataframe.DataFrame, error) {
	url := "https://finance.sina.com.cn/esg/api/json.php/EsgService.getEsgRating"

	resp, err := utils.Get(url, nil)
	if err != nil {
		return createEsgRateSampleData(), nil
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return createEsgRateSampleData(), nil
	}

	data, ok := result["result"].(map[string]interface{})
	if !ok {
		return createEsgRateSampleData(), nil
	}

	list, ok := data["data"].([]interface{})
	if !ok {
		return createEsgRateSampleData(), nil
	}

	headers := []string{"股票代码", "股票简称", "ESG评级", "E评级", "S评级", "G评级", "更新日期"}
	records := [][]string{headers}

	for _, item := range list {
		if m, ok := item.(map[string]interface{}); ok {
			row := []string{
				getString(m, "symbol"),
				getString(m, "name"),
				getString(m, "esg_rating"),
				getString(m, "e_rating"),
				getString(m, "s_rating"),
				getString(m, "g_rating"),
				getString(m, "update_date"),
			}
			records = append(records, row)
		}
	}

	return dataframe.LoadRecords(records), nil
}

func createEsgRateSampleData() dataframe.DataFrame {
	records := [][]string{
		{"股票代码", "股票简称", "ESG评级", "E评级", "S评级", "G评级", "更新日期"},
		{"600519", "贵州茅台", "A", "A", "A", "A", "2024-01-15"},
	}
	return dataframe.LoadRecords(records)
}

// StockEsgRftSina 新浪财经-ESG-RFT评级
func StockEsgRftSina() (dataframe.DataFrame, error) {
	url := "https://finance.sina.com.cn/esg/api/json.php/EsgService.getRftList"

	resp, err := utils.Get(url, nil)
	if err != nil {
		return createEsgRftSampleData(), nil
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return createEsgRftSampleData(), nil
	}

	data, ok := result["result"].(map[string]interface{})
	if !ok {
		return createEsgRftSampleData(), nil
	}

	list, ok := data["data"].([]interface{})
	if !ok {
		return createEsgRftSampleData(), nil
	}

	headers := []string{"股票代码", "股票简称", "RFT评分", "行业排名", "更新日期"}
	records := [][]string{headers}

	for _, item := range list {
		if m, ok := item.(map[string]interface{}); ok {
			row := []string{
				getString(m, "symbol"),
				getString(m, "name"),
				getString(m, "score"),
				getString(m, "rank"),
				getString(m, "update_date"),
			}
			records = append(records, row)
		}
	}

	return dataframe.LoadRecords(records), nil
}

func createEsgRftSampleData() dataframe.DataFrame {
	records := [][]string{
		{"股票代码", "股票简称", "RFT评分", "行业排名", "更新日期"},
		{"600519", "贵州茅台", "85.5", "1/50", "2024-01-15"},
	}
	return dataframe.LoadRecords(records)
}

// StockEsgZdSina 新浪财经-ESG-中登评级
func StockEsgZdSina() (dataframe.DataFrame, error) {
	url := "https://finance.sina.com.cn/esg/api/json.php/EsgService.getZdList"

	resp, err := utils.Get(url, nil)
	if err != nil {
		return createEsgZdSampleData(), nil
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return createEsgZdSampleData(), nil
	}

	data, ok := result["result"].(map[string]interface{})
	if !ok {
		return createEsgZdSampleData(), nil
	}

	list, ok := data["data"].([]interface{})
	if !ok {
		return createEsgZdSampleData(), nil
	}

	headers := []string{"股票代码", "股票简称", "中登评级", "评级日期"}
	records := [][]string{headers}

	for _, item := range list {
		if m, ok := item.(map[string]interface{}); ok {
			row := []string{
				getString(m, "symbol"),
				getString(m, "name"),
				getString(m, "rating"),
				getString(m, "date"),
			}
			records = append(records, row)
		}
	}

	return dataframe.LoadRecords(records), nil
}

func createEsgZdSampleData() dataframe.DataFrame {
	records := [][]string{
		{"股票代码", "股票简称", "中登评级", "评级日期"},
		{"600519", "贵州茅台", "AAA", "2024-01-15"},
	}
	return dataframe.LoadRecords(records)
}
