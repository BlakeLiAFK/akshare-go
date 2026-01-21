package air

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

// XML数据结构定义
type xmlRoot struct {
	XMLName xml.Name  `xml:"root"`
	Cities  []xmlCity `xml:"City"`
}

type xmlCity struct {
	Name     string       `xml:"Name"`
	Pointers []xmlPointer `xml:"Pointer"`
}

type xmlPointer struct {
	Region   string    `xml:"Region"`
	Name     string    `xml:"Name"`
	DataTime string    `xml:"DataTime"`
	AQI      string    `xml:"AQI"`
	Level    string    `xml:"Level"`
	MaxPoll  string    `xml:"MaxPoll"`
	CLng     string    `xml:"CLng"`
	CLat     string    `xml:"CLat"`
	Polls    []xmlPoll `xml:"Poll"`
}

type xmlPoll struct {
	Name  string `xml:"Name"`
	Value string `xml:"Value"`
	IAQI  string `xml:"IAQI"`
}

// AirQualityHebei 获取河北省空气质量数据
//
// 数据源: 河北省环境监测
// URL: http://218.11.10.130:8080/api/hour/130000.xml
//
// 返回:
//   - dataframe.DataFrame: 包含城市、区域、监测点、AQI及各污染物数据
//   - error: 错误信息
//
// 示例:
//
//	df, err := air.AirQualityHebei()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(df)
func AirQualityHebei() (dataframe.DataFrame, error) {
	url := "http://218.11.10.130:8080/api/hour/130000.xml"
	
	// 发起HTTP请求
	resp, err := utils.Get(url, nil)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求失败: %w", err)
	}
	
	// 解析XML
	var root xmlRoot
	if err := xml.Unmarshal([]byte(resp.String()), &root); err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("解析XML失败: %w", err)
	}
	
	// 提取数据
	var records []map[string]interface{}
	for _, city := range root.Cities {
		for _, pointer := range city.Pointers {
			record := map[string]interface{}{
				"城市":    city.Name,
				"区域":    pointer.Region,
				"监测点":   pointer.Name,
				"时间":    pointer.DataTime,
				"AQI":   parseFloat(pointer.AQI),
				"空气质量等级": pointer.Level,
				"首要污染物": pointer.MaxPoll,
				"经度":    parseFloat(pointer.CLng),
				"纬度":    parseFloat(pointer.CLat),
			}
			
			// 处理污染物数据
			for _, poll := range pointer.Polls {
				valueName := fmt.Sprintf("%s_浓度", poll.Name)
				iaqiName := fmt.Sprintf("%s_IAQI", poll.Name)
				record[valueName] = parseFloat(poll.Value)
				record[iaqiName] = parseFloat(poll.IAQI)
			}
			
			records = append(records, record)
		}
	}
	
	if len(records) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到数据")
	}
	
	// 转换为DataFrame
	return recordsToDataFrame(records), nil
}

// parseFloat 将字符串转换为float64，转换失败返回0
func parseFloat(s string) float64 {
	if s == "" {
		return 0
	}
	v, _ := strconv.ParseFloat(s, 64)
	return v
}

// recordsToDataFrame 将map记录转换为DataFrame
func recordsToDataFrame(records []map[string]interface{}) dataframe.DataFrame {
	if len(records) == 0 {
		return dataframe.DataFrame{}
	}
	
	// 获取所有列名
	var columns []string
	columnSet := make(map[string]bool)
	for _, record := range records {
		for key := range record {
			if !columnSet[key] {
				columns = append(columns, key)
				columnSet[key] = true
			}
		}
	}
	
	// 按照固定顺序排列列
	orderedColumns := []string{
		"城市", "区域", "监测点", "时间", "AQI", 
		"空气质量等级", "首要污染物", "经度", "纬度",
	}
	
	// 添加其他列（污染物数据）
	for _, col := range columns {
		found := false
		for _, ordered := range orderedColumns {
			if col == ordered {
				found = true
				break
			}
		}
		if !found {
			orderedColumns = append(orderedColumns, col)
		}
	}
	
	// 构建series
	var seriesList []series.Series
	for _, colName := range orderedColumns {
		var values []interface{}
		for _, record := range records {
			if val, ok := record[colName]; ok {
				values = append(values, val)
			} else {
				values = append(values, nil)
			}
		}
		seriesList = append(seriesList, series.New(values, series.String, colName))
	}
	
	return dataframe.New(seriesList...)
}
