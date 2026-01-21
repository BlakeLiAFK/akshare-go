package nlp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestNlpOwnthinkEntity 测试知识图谱-实体查询
func TestNlpOwnthinkEntity(t *testing.T) {
	result, err := NlpOwnthinkEntity("人工智能")
	assert.NoError(t, err, "NlpOwnthinkEntity 不应返回错误")
	assert.NotEmpty(t, result, "实体名称不应为空")
	t.Logf("实体: %s", result)
}

// TestNlpOwnthinkDesc 测试知识图谱-描述查询
func TestNlpOwnthinkDesc(t *testing.T) {
	result, err := NlpOwnthinkDesc("人工智能")
	assert.NoError(t, err, "NlpOwnthinkDesc 不应返回错误")
	assert.NotEmpty(t, result, "描述不应为空")
	t.Logf("描述: %s", result)
}

// TestNlpOwnthinkAvp 测试知识图谱-属性值对查询
func TestNlpOwnthinkAvp(t *testing.T) {
	result, err := NlpOwnthinkAvp("人工智能")
	assert.NoError(t, err, "NlpOwnthinkAvp 不应返回错误")
	assert.NotNil(t, result, "属性值对不应为空")

	if len(result) > 0 {
		t.Logf("获取到 %d 个属性值对", len(result))
		t.Logf("第一个: %s = %s", result[0].Field, result[0].Value)
	}
}

// TestNlpOwnthinkTag 测试知识图谱-标签查询
func TestNlpOwnthinkTag(t *testing.T) {
	result, err := NlpOwnthinkTag("人工智能")
	assert.NoError(t, err, "NlpOwnthinkTag 不应返回错误")
	// 标签可能为空，这是正常情况
	t.Logf("获取到 %d 个标签: %v", len(result), result)
}

// TestNlpOwnthink 测试通用知识图谱接口
func TestNlpOwnthink(t *testing.T) {
	// 测试 entity
	result, err := NlpOwnthink("姚明", "entity")
	assert.NoError(t, err)
	assert.NotEmpty(t, result)
	t.Logf("姚明 entity: %v", result)

	// 测试 desc
	result, err = NlpOwnthink("姚明", "desc")
	assert.NoError(t, err)
	t.Logf("姚明 desc: %v", result)
}

// TestNlpAnswer 测试智能问答接口
func TestNlpAnswer(t *testing.T) {
	result, err := NlpAnswer("姚明的身高")
	assert.NoError(t, err, "NlpAnswer 不应返回错误")
	assert.NotEmpty(t, result, "答案不应为空")
	t.Logf("问: 姚明的身高")
	t.Logf("答: %s", result)
}

// TestNlpAnswerGeneral 测试智能问答-通用问题
func TestNlpAnswerGeneral(t *testing.T) {
	result, err := NlpAnswer("人工智能是什么")
	assert.NoError(t, err, "NlpAnswer 不应返回错误")
	assert.NotEmpty(t, result, "答案不应为空")
	t.Logf("问: 人工智能是什么")
	t.Logf("答: %s", result)
}

// TestNlpOwnthinkNotFound 测试知识图谱-不存在的词语
func TestNlpOwnthinkNotFound(t *testing.T) {
	_, err := NlpOwnthinkEntity("这是一个不存在的测试词语12345")
	// 可能返回错误或空结果
	if err != nil {
		t.Logf("预期的错误: %v", err)
	}
}
