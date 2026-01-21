package economic

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-resty/resty/v2"
	"github.com/tidwall/gjson"
)

// nbsClient 创建忽略 SSL 证书的客户端
var nbsClient = resty.New().
	SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
	SetTimeout(30 * time.Second)

// getNBSTree 获取指标目录树
func getNBSTree(idcode, dbcode string) (gjson.Result, error) {
	url := "https://data.stats.gov.cn/easyquery.htm"

	resp, err := nbsClient.R().
		SetQueryParams(map[string]string{
			"id":     idcode,
			"dbcode": dbcode,
			"wdcode": "zb",
			"m":      "getTree",
		}).
		Post(url)
	if err != nil {
		return gjson.Result{}, fmt.Errorf("请求失败: %w", err)
	}

	return gjson.Parse(resp.String()), nil
}

// getNBSWdsTree 获取地区数据的可选指标目录树
func getNBSWdsTree(idcode, dbcode, rowcode string) (gjson.Result, error) {
	url := "https://data.stats.gov.cn/easyquery.htm"

	resp, err := nbsClient.R().
		SetQueryParams(map[string]string{
			"m":       "getOtherWds",
			"dbcode":  dbcode,
			"rowcode": rowcode,
			"colcode": "sj",
			"wds":     fmt.Sprintf(`[{"wdcode":"zb","valuecode":"%s"}]`, idcode),
			"k1":      fmt.Sprintf("%d", time.Now().UnixMilli()),
		}).
		Post(url)
	if err != nil {
		return gjson.Result{}, fmt.Errorf("请求失败: %w", err)
	}

	return gjson.Get(resp.String(), "returndata.0.nodes"), nil
}

// getCodeFromNBSTree 根据指标名称从目录树中获取编码
func getCodeFromNBSTree(tree gjson.Result, name, target string) (string, error) {
	var code string
	tree.ForEach(func(_, item gjson.Result) bool {
		if item.Get("name").String() == name {
			code = item.Get(target).String()
			return false
		}
		return true
	})
	if code == "" {
		return "", fmt.Errorf("未找到指标: %s", name)
	}
	return code, nil
}

// MacroChinaNBSNation 国家统计局全国数据通用接口
//
// 参数:
//   - kind: 数据类别，可选值: "月度数据", "季度数据", "年度数据"
//   - path: 数据路径，使用 ">" 分隔，如 "工业 > 工业分大类行业出口交货值(2018-至今) > 废弃资源综合利用业"
//   - period: 时间区间，如 "LAST10", "2016-2023", "2016-" 等
//
// 返回:
//   - dataframe.DataFrame: 国家统计局统计数据
//   - error: 错误信息
//
// 数据源: https://data.stats.gov.cn/easyquery.htm
func MacroChinaNBSNation(kind, path, period string) (dataframe.DataFrame, error) {
	// 获取 dbcode
	kindCode := map[string]string{
		"月度数据": "hgyd",
		"季度数据": "hgjd",
		"年度数据": "hgnd",
	}
	dbcode, ok := kindCode[kind]
	if !ok {
		return dataframe.DataFrame{}, fmt.Errorf("无效的数据类别: %s", kind)
	}

	// 获取最终 id
	parentTree, err := getNBSTree("zb", dbcode)
	if err != nil {
		return dataframe.DataFrame{}, err
	}

	pathSplit := strings.Split(strings.ReplaceAll(path, " ", ""), ">")
	indicatorID, err := getCodeFromNBSTree(parentTree, pathSplit[0], "id")
	if err != nil {
		return dataframe.DataFrame{}, err
	}
	pathSplit = pathSplit[1:]

	for len(pathSplit) > 0 {
		tempTree, err := getNBSTree(indicatorID, dbcode)
		if err != nil {
			return dataframe.DataFrame{}, err
		}
		indicatorID, err = getCodeFromNBSTree(tempTree, pathSplit[0], "id")
		if err != nil {
			return dataframe.DataFrame{}, err
		}
		pathSplit = pathSplit[1:]
	}

	// 请求数据
	url := "https://data.stats.gov.cn/easyquery.htm"
	resp, err := nbsClient.R().
		SetQueryParams(map[string]string{
			"m":       "QueryData",
			"dbcode":  dbcode,
			"rowcode": "zb",
			"colcode": "sj",
			"wds":     "[]",
			"dfwds":   fmt.Sprintf(`[{"wdcode":"zb","valuecode":"%s"}, {"wdcode":"sj","valuecode":"%s"}]`, indicatorID, period),
			"k1":      fmt.Sprintf("%d", time.Now().UnixMilli()),
		}).
		Get(url)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求数据失败: %w", err)
	}

	result := gjson.Parse(resp.String())

	// 获取数据节点
	datanodes := result.Get("returndata.datanodes").Array()
	if len(datanodes) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到数据")
	}

	// 获取维度节点
	wdnodes := result.Get("returndata.wdnodes").Array()
	if len(wdnodes) < 2 {
		return dataframe.DataFrame{}, fmt.Errorf("维度数据格式错误")
	}

	// 获取行名和列名
	var rowNames []string
	wdnodes[0].Get("nodes").ForEach(func(_, item gjson.Result) bool {
		name := item.Get("cname").String()
		unit := item.Get("unit").String()
		if unit != "" {
			name = name + "(" + unit + ")"
		}
		rowNames = append(rowNames, name)
		return true
	})

	var colNames []string
	wdnodes[1].Get("nodes").ForEach(func(_, item gjson.Result) bool {
		name := item.Get("cname").String()
		unit := item.Get("unit").String()
		if unit != "" {
			name = name + "(" + unit + ")"
		}
		colNames = append(colNames, name)
		return true
	})

	// 构建数据矩阵
	numRows := len(rowNames)
	numCols := len(colNames)

	// 构建记录 - 第一行是列名，第一列是指标名
	var records [][]string
	header := append([]string{"指标"}, colNames...)
	records = append(records, header)

	// 提取数据值
	var dataValues []string
	for _, node := range datanodes {
		data := node.Get("data")
		if data.Get("hasdata").Bool() {
			dataValues = append(dataValues, data.Get("data").String())
		} else {
			dataValues = append(dataValues, "")
		}
	}

	// 按行填充数据
	for i := 0; i < numRows; i++ {
		record := []string{rowNames[i]}
		for j := 0; j < numCols; j++ {
			idx := i*numCols + j
			if idx < len(dataValues) {
				record = append(record, dataValues[idx])
			} else {
				record = append(record, "")
			}
		}
		records = append(records, record)
	}

	df := dataframe.LoadRecords(records)
	return df, nil
}

// MacroChinaNBSRegion 国家统计局地区数据通用接口
//
// 参数:
//   - kind: 数据类别，可选值: "分省月度数据", "分省季度数据", "分省年度数据",
//     "主要城市月度价格", "主要城市年度数据", "港澳台月度数据", "港澳台年度数据"
//   - path: 数据路径，使用 ">" 分隔
//   - indicator: 指定指标（当 region 不为空时可为空）
//   - region: 指定地区（当 indicator 不为空时可为空）
//   - period: 时间区间
//
// 返回:
//   - dataframe.DataFrame: 国家统计局统计数据
//   - error: 错误信息
//
// 数据源: https://data.stats.gov.cn/easyquery.htm
func MacroChinaNBSRegion(kind, path, indicator, region, period string) (dataframe.DataFrame, error) {
	if indicator == "" && region == "" {
		return dataframe.DataFrame{}, fmt.Errorf("indicator 和 region 参数不能同时为空")
	}

	// 获取 dbcode
	kindDict := map[string]string{
		"分省月度数据":   "fsyd",
		"分省季度数据":   "fsjd",
		"分省年度数据":   "fsnd",
		"主要城市月度价格": "csyd",
		"主要城市年度数据": "csnd",
		"港澳台月度数据":  "gatyd",
		"港澳台年度数据":  "gatnd",
	}
	dbcode, ok := kindDict[kind]
	if !ok {
		return dataframe.DataFrame{}, fmt.Errorf("无效的数据类别: %s", kind)
	}

	// 获取最终 id
	parentTree, err := getNBSTree("zb", dbcode)
	if err != nil {
		return dataframe.DataFrame{}, err
	}

	pathSplit := strings.Split(strings.ReplaceAll(path, " ", ""), ">")
	indicatorID, err := getCodeFromNBSTree(parentTree, pathSplit[0], "id")
	if err != nil {
		return dataframe.DataFrame{}, err
	}
	pathSplit = pathSplit[1:]

	for len(pathSplit) > 0 {
		tempTree, err := getNBSTree(indicatorID, dbcode)
		if err != nil {
			return dataframe.DataFrame{}, err
		}
		indicatorID, err = getCodeFromNBSTree(tempTree, pathSplit[0], "id")
		if err != nil {
			return dataframe.DataFrame{}, err
		}
		pathSplit = pathSplit[1:]
	}

	// 参数设定
	var rowcode, colcode, wds, dfwds string
	if region == "" {
		// 按指标查询所有地区
		indicatorTree, err := getNBSWdsTree(indicatorID, dbcode, "reg")
		if err != nil {
			return dataframe.DataFrame{}, err
		}
		indicatorID, err = getCodeFromNBSTree(indicatorTree, indicator, "code")
		if err != nil {
			return dataframe.DataFrame{}, err
		}
		rowcode = "reg"
		colcode = "sj"
		wds = fmt.Sprintf(`[{"wdcode":"zb","valuecode":"%s"}]`, indicatorID)
		dfwds = fmt.Sprintf(`[{"wdcode":"sj","valuecode":"%s"}]`, period)
	} else {
		// 按地区查询指标
		if indicator != "" {
			indicatorTree, err := getNBSWdsTree(indicatorID, dbcode, "reg")
			if err != nil {
				return dataframe.DataFrame{}, err
			}
			indicatorID, err = getCodeFromNBSTree(indicatorTree, indicator, "code")
			if err != nil {
				return dataframe.DataFrame{}, err
			}
		}
		regionTree, err := getNBSWdsTree(indicatorID, dbcode, "zb")
		if err != nil {
			return dataframe.DataFrame{}, err
		}
		regionID, err := getCodeFromNBSTree(regionTree, region, "code")
		if err != nil {
			return dataframe.DataFrame{}, err
		}
		rowcode = "zb"
		colcode = "sj"
		wds = fmt.Sprintf(`[{"wdcode":"reg","valuecode":"%s"}]`, regionID)
		dfwds = fmt.Sprintf(`[{"wdcode":"zb","valuecode":"%s"}, {"wdcode":"sj","valuecode":"%s"}]`, indicatorID, period)
	}

	// 请求数据
	url := "https://data.stats.gov.cn/easyquery.htm"
	resp, err := nbsClient.R().
		SetQueryParams(map[string]string{
			"m":       "QueryData",
			"dbcode":  dbcode,
			"rowcode": rowcode,
			"colcode": colcode,
			"wds":     wds,
			"dfwds":   dfwds,
			"k1":      fmt.Sprintf("%d", time.Now().UnixMilli()),
		}).
		Get(url)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("请求数据失败: %w", err)
	}

	result := gjson.Parse(resp.String())

	// 获取数据节点
	datanodes := result.Get("returndata.datanodes").Array()
	if len(datanodes) == 0 {
		return dataframe.DataFrame{}, fmt.Errorf("未获取到数据")
	}

	// 获取维度节点
	wdnodes := result.Get("returndata.wdnodes").Array()
	if len(wdnodes) < 3 {
		return dataframe.DataFrame{}, fmt.Errorf("维度数据格式错误")
	}

	// 确定行名和列名的索引
	var rowIdx, colIdx, titleIdx int
	if region == "" {
		rowIdx = 1
		colIdx = 2
		titleIdx = 0
	} else {
		rowIdx = 0
		colIdx = 2
		titleIdx = 1
	}

	// 获取行名和列名
	var rowNames []string
	wdnodes[rowIdx].Get("nodes").ForEach(func(_, item gjson.Result) bool {
		name := item.Get("cname").String()
		unit := item.Get("unit").String()
		if unit != "" {
			name = name + "(" + unit + ")"
		}
		rowNames = append(rowNames, name)
		return true
	})

	var colNames []string
	wdnodes[colIdx].Get("nodes").ForEach(func(_, item gjson.Result) bool {
		name := item.Get("cname").String()
		unit := item.Get("unit").String()
		if unit != "" {
			name = name + "(" + unit + ")"
		}
		colNames = append(colNames, name)
		return true
	})

	// 获取标题名
	titleName := wdnodes[titleIdx].Get("nodes.0.cname").String()

	// 构建数据矩阵
	numRows := len(rowNames)
	numCols := len(colNames)

	// 构建记录
	var records [][]string
	header := append([]string{titleName}, colNames...)
	records = append(records, header)

	// 提取数据值
	var dataValues []string
	for _, node := range datanodes {
		data := node.Get("data")
		if data.Get("hasdata").Bool() {
			dataValues = append(dataValues, data.Get("data").String())
		} else {
			dataValues = append(dataValues, "")
		}
	}

	// 按行填充数据
	for i := 0; i < numRows; i++ {
		record := []string{rowNames[i]}
		for j := 0; j < numCols; j++ {
			idx := i*numCols + j
			if idx < len(dataValues) {
				record = append(record, dataValues[idx])
			} else {
				record = append(record, "")
			}
		}
		records = append(records, record)
	}

	df := dataframe.LoadRecords(records)
	return df, nil
}

// 确保 http.Client 可用
var _ = http.Client{}
