package nlp

// AvpItem 属性值对数据结构
// 对应 Python: nlp_ownthink 返回的 avp DataFrame
type AvpItem struct {
	Field string `json:"field"` // 字段
	Value string `json:"value"` // 值
}

// KnowledgeResult 知识图谱查询结果
type KnowledgeResult struct {
	Entity string    // 实体名称
	Desc   string    // 描述
	Avp    []AvpItem // 属性值对列表
	Tag    []string  // 标签列表
}
