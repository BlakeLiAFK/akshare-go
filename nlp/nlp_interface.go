package nlp

import (
	"fmt"

	"github.com/BlakeLiAFK/akshare/utils"
	"github.com/tidwall/gjson"
)

// NlpOwnthink 知识图谱接口
// https://ownthink.com/
// word: 中文词语
// indicator: 查询类型 (entity/desc/avp/tag)
func NlpOwnthink(word string, indicator string) (interface{}, error) {
	url := "https://api.ownthink.com/kg/knowledge"

	payload := map[string]string{
		"entity": word,
	}

	resp, err := utils.Post(url, payload)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}

	text := resp.String()
	data := gjson.Get(text, "data")

	if !data.Exists() || data.String() == "" || data.String() == "null" {
		return nil, fmt.Errorf("未找到资源，请输入正确的词语")
	}

	switch indicator {
	case "entity":
		return data.Get("entity").String(), nil
	case "desc":
		return data.Get("desc").String(), nil
	case "avp":
		var avpList []AvpItem
		data.Get("avp").ForEach(func(key, value gjson.Result) bool {
			arr := value.Array()
			if len(arr) >= 2 {
				avpList = append(avpList, AvpItem{
					Field: arr[0].String(),
					Value: arr[1].String(),
				})
			}
			return true
		})
		return avpList, nil
	case "tag":
		var tags []string
		data.Get("tag").ForEach(func(key, value gjson.Result) bool {
			tags = append(tags, value.String())
			return true
		})
		return tags, nil
	default:
		return nil, fmt.Errorf("不支持的 indicator: %s", indicator)
	}
}

// NlpOwnthinkEntity 知识图谱-获取实体名称
func NlpOwnthinkEntity(word string) (string, error) {
	result, err := NlpOwnthink(word, "entity")
	if err != nil {
		return "", err
	}
	return result.(string), nil
}

// NlpOwnthinkDesc 知识图谱-获取描述
func NlpOwnthinkDesc(word string) (string, error) {
	result, err := NlpOwnthink(word, "desc")
	if err != nil {
		return "", err
	}
	return result.(string), nil
}

// NlpOwnthinkAvp 知识图谱-获取属性值对
func NlpOwnthinkAvp(word string) ([]AvpItem, error) {
	result, err := NlpOwnthink(word, "avp")
	if err != nil {
		return nil, err
	}
	return result.([]AvpItem), nil
}

// NlpOwnthinkTag 知识图谱-获取标签
func NlpOwnthinkTag(word string) ([]string, error) {
	result, err := NlpOwnthink(word, "tag")
	if err != nil {
		return nil, err
	}
	return result.([]string), nil
}

// NlpAnswer 智能问答
// https://ownthink.com/robot.html
// question: 问题（中文）
func NlpAnswer(question string) (string, error) {
	url := "https://api.ownthink.com/bot"

	params := map[string]string{
		"spoken": question,
	}

	resp, err := utils.Get(url, params)
	if err != nil {
		return "", fmt.Errorf("请求失败: %w", err)
	}

	text := resp.String()
	answer := gjson.Get(text, "data.info.text").String()

	return answer, nil
}
