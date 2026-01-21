package bank

// PenaltyRecord 银保监行政处罚记录
type PenaltyRecord struct {
	DocNumber       string `json:"行政处罚决定书文号"`   // 行政处罚决定书文号
	Name            string `json:"姓名"`          // 姓名
	Organization    string `json:"单位"`          // 单位
	OrgName         string `json:"单位名称"`        // 单位名称
	ResponsibleName string `json:"主要负责人姓名"`     // 主要负责人姓名
	Violations      string `json:"主要违法违规事实"`    // 主要违法违规事实（案由）
	LegalBasis      string `json:"行政处罚依据"`      // 行政处罚依据
	Decision        string `json:"行政处罚决定"`      // 行政处罚决定
	Authority       string `json:"作出处罚决定的机关名称"` // 作出处罚决定的机关名称
	DecisionDate    string `json:"处罚决定日期"`      // 处罚决定日期
	PublicationDate string `json:"处罚公开日期"`      // 处罚公开日期
}

// PageInfo 分页信息
type PageInfo struct {
	DocID       string `json:"docId"`       // 文档ID
	Subtitle    string `json:"subtitle"`    // 副标题
	PublishDate string `json:"publishDate"` // 发布日期
}
