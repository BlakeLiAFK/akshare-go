package index

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/tidwall/gjson"
)

// DrewryWCI Drewry集装箱指数
type DrewryWCI struct {
	Date time.Time `json:"date"` // 日期
	WCI  float64   `json:"wci"`  // 集装箱指数
}

// DrewryWCIIndex Drewry 集装箱指数
//
// 数据源: https://infogram.com/world-container-index-1h17493095xl4zj
//
// 参数:
//   - symbol: 航线类型，可选:
//     "composite" - 综合指数
//     "shanghai-rotterdam" - 上海-鹿特丹
//     "rotterdam-shanghai" - 鹿特丹-上海
//     "shanghai-los angeles" - 上海-洛杉矶
//     "los angeles-shanghai" - 洛杉矶-上海
//     "shanghai-genoa" - 上海-热那亚
//     "new york-rotterdam" - 纽约-鹿特丹
//     "rotterdam-new york" - 鹿特丹-纽约
//
// 返回:
//   - []DrewryWCI: 集装箱指数数据
//   - error: 错误信息
func DrewryWCIIndex(symbol string) ([]DrewryWCI, error) {
	symbolMap := map[string]int{
		"composite":            0,
		"shanghai-rotterdam":   1,
		"rotterdam-shanghai":   2,
		"shanghai-los angeles": 3,
		"los angeles-shanghai": 4,
		"shanghai-genoa":       5,
		"new york-rotterdam":   6,
		"rotterdam-new york":   7,
	}

	idx, ok := symbolMap[symbol]
	if !ok {
		return nil, fmt.Errorf("无效的 symbol: %s", symbol)
	}

	apiURL := "https://infogram.com/world-container-index-1h17493095xl4zj"

	resp, err := utils.Get(apiURL, nil)
	if err != nil {
		return nil, fmt.Errorf("请求 Drewry 集装箱指数失败: %w", err)
	}

	// 从 HTML 中提取 JavaScript 数据
	htmlContent := resp.String()

	// 查找 window.infographicData= 后的 JSON 数据
	re := regexp.MustCompile(`window\.infographicData\s*=\s*(\{[\s\S]*?\});`)
	matches := re.FindStringSubmatch(htmlContent)
	if len(matches) < 2 {
		return nil, fmt.Errorf("未找到 infographicData")
	}

	jsonData := matches[1]

	// 使用 gjson 解析数据
	// 数据路径: elements.content.content.entities.{uuid}.data[idx]
	entitiesPath := "elements.content.content.entities"
	entities := gjson.Get(jsonData, entitiesPath)

	var dataArray gjson.Result
	entities.ForEach(func(key, value gjson.Result) bool {
		if value.Get("data").Exists() {
			dataArray = value.Get("data")
			return false
		}
		return true
	})

	if !dataArray.Exists() {
		return nil, fmt.Errorf("未找到数据数组")
	}

	// 获取指定索引的数据
	targetData := dataArray.Array()
	if idx >= len(targetData) {
		return nil, fmt.Errorf("索引超出范围: %d >= %d", idx, len(targetData))
	}

	dataRows := targetData[idx].Array()
	if len(dataRows) < 2 {
		return nil, fmt.Errorf("数据行数不足")
	}

	result := make([]DrewryWCI, 0, len(dataRows)-1)
	for i := 1; i < len(dataRows); i++ {
		row := dataRows[i].Array()
		if len(row) < 2 {
			continue
		}

		dateStr := row[0].Get("value").String()
		wciStr := row[1].Get("value").String()

		// 解析日期 (格式: 02-Jan-06)
		date, err := parseDrewryDate(dateStr)
		if err != nil {
			continue
		}

		wci := utils.MustParseFloat(wciStr)

		result = append(result, DrewryWCI{
			Date: date,
			WCI:  wci,
		})
	}

	return result, nil
}

// parseDrewryDate 解析 Drewry 日期格式
func parseDrewryDate(dateStr string) (time.Time, error) {
	// 格式: 02-Jan-06
	layouts := []string{
		"02-Jan-06",
		"2-Jan-06",
		"02-Jan-2006",
		"2006-01-02",
	}

	dateStr = strings.TrimSpace(dateStr)
	for _, layout := range layouts {
		if t, err := time.Parse(layout, dateStr); err == nil {
			return t, nil
		}
	}

	return time.Time{}, fmt.Errorf("无法解析日期: %s", dateStr)
}
