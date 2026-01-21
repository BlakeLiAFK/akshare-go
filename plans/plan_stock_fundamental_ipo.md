# stock_fundamental 模块 - IPO 接口实现计划

> 任务: 实现 stock_fundamental 模块的 IPO 相关接口
> 创建日期: 2026-01-18
> 完成日期: 2026-01-18
> 状态: 已完成

---

## 一、任务概述

实现 stock_fundamental 模块的第二批 IPO 相关接口，共 4 个函数。

---

## 二、接口清单

### 2.1 东方财富 IPO 接口 (stock_fundamental_ipo_em.go)

#### 1. StockIpoDeclareEm
- **功能**: 东方财富网-数据中心-新股申购-首发申报企业信息
- **API**: `https://datacenter-web.eastmoney.com/api/data/v1/get`
- **参数**: reportName=RPT_IPO_DECORGNEWEST, pageSize=500
- **返回字段**:
  - Index: 序号
  - CompanyName: 企业名称
  - State: 最新状态
  - RegAddress: 注册地
  - RecommendOrg: 保荐机构
  - LawFirm: 律师事务所
  - AccountFirm: 会计师事务所
  - PredictMarket: 拟上市地点
  - EndDate: 更新日期
  - ProspectusURL: 招股说明书链接

#### 2. StockIpoReviewEm
- **功能**: 东方财富网-数据中心-新股申购-过会企业信息
- **API**: `https://datacenter-web.eastmoney.com/api/data/v1/get`
- **参数**: reportName=RPT_IPO_REVIEW, columns=ALL
- **返回字段**:
  - Index: 序号
  - CompanyName: 企业名称
  - StockName: 股票简称
  - StockCode: 股票代码
  - TradeMarket: 上市板块
  - ReviewDate: 上会日期
  - ReviewState: 审核状态
  - ReviewMember: 发审委委员
  - LeadUnderwriter: 主承销商
  - IssueNum: 发行数量(股)
  - FinanceAmt: 拟融资额(元)
  - NoticeDate: 公告日期
  - ListingDate: 上市日期
- **注意**: 响应是 JSONP 格式，需要提取 JSON 部分

#### 3. StockIpoTutorEm
- **功能**: 东方财富网-数据中心-新股申购-辅导备案信息
- **API**: `https://datacenter-web.eastmoney.com/api/data/v1/get`
- **参数**: reportName=RPT_IPO_TUTRECORD
- **返回字段**:
  - Index: 序号
  - CompanyName: 企业名称
  - TutorOrg: 辅导机构
  - TutorProcessState: 辅导状态
  - ReportType: 报告类型
  - DispatchOrg: 派出机构
  - ReportTitle: 报告标题
  - RecordDate: 备案日期
- **注意**: 响应是 JSONP 格式，需要提取 JSON 部分

### 2.2 上交所科创板接口 (stock_fundamental_kcb_sse.go)

#### 4. StockKcbSse
- **功能**: 上交所科创板-股票发行
- **API**: `http://query.sse.com.cn/statusAction.do`
- **参数**: sqlId=SH_XM_LB, 分页请求
- **返回字段**:
  - Index: 序号
  - CompanyName: 企业名称
  - StockCode: 股票代码
  - UpdateTime: 更新日期
  - AuditApplyDate: 受理日期
  - RegisterResult: 注册结果
  - CommitResult: 审议结果
  - CurrentStatus: 当前状态
  - Province: 省份
  - CsrcCode: 证监会行业
  - StockAuditNum: 审核序号

---

## 三、技术要点

### 3.1 JSONP 解析
东方财富的某些接口返回 JSONP 格式，需要提取纯 JSON 部分：
```go
func parseJSONP(text string) (string, error) {
    start := strings.Index(text, "{")
    end := strings.LastIndex(text, "}")
    if start == -1 || end == -1 || end <= start {
        return "", fmt.Errorf("无法找到 JSON 对象")
    }
    return text[start : end+1], nil
}
```

### 3.2 分页处理
所有接口都需要自动分页获取所有数据：
1. 第一次请求获取总页数
2. 循环请求每一页数据
3. 合并所有结果

### 3.3 日期解析
- 东方财富日期格式: "2024-01-18"
- 科创板日期格式: "20240116171217" (需要截取前8位)

### 3.4 状态码转换
科创板接口的状态是数字，需要转换为中文：
- 1: 已受理
- 2: 已问询
- 3: 上市委会议通过
- 4: 提交注册
- 5: 注册结果
- 6: 中止
- 7: 终止

---

## 四、文件结构

```
stock_fundamental/
├── types.go                                # 类型定义 (已更新)
├── stock_fundamental_ipo_em.go             # 东方财富IPO接口 (新建)
├── stock_fundamental_kcb_sse.go            # 上交所科创板接口 (新建)
├── stock_fundamental_ipo_em_test.go        # 测试用例 (新建)
├── stock_fundamental_finance_sina.go       # 新浪财经接口 (已存在)
└── stock_fundamental_finance_sina_test.go  # 新浪财经测试 (已存在)
```

---

## 五、类型定义

### 5.1 StockIpoDeclareEmItem
```go
type StockIpoDeclareEmItem struct {
    Index           int       `json:"index"`
    CompanyName     string    `json:"company_name"`
    State           string    `json:"state"`
    RegAddress      string    `json:"reg_address"`
    RecommendOrg    string    `json:"recommend_org"`
    LawFirm         string    `json:"law_firm"`
    AccountFirm     string    `json:"account_firm"`
    PredictMarket   string    `json:"predict_market"`
    EndDate         time.Time `json:"end_date"`
    ProspectusURL   string    `json:"prospectus_url"`
}
```

### 5.2 StockIpoReviewEmItem
```go
type StockIpoReviewEmItem struct {
    Index           int       `json:"index"`
    CompanyName     string    `json:"company_name"`
    StockName       string    `json:"stock_name"`
    StockCode       string    `json:"stock_code"`
    TradeMarket     string    `json:"trade_market"`
    ReviewDate      time.Time `json:"review_date"`
    ReviewState     string    `json:"review_state"`
    ReviewMember    string    `json:"review_member"`
    LeadUnderwriter string    `json:"lead_underwriter"`
    IssueNum        float64   `json:"issue_num"`
    FinanceAmt      float64   `json:"finance_amt"`
    NoticeDate      time.Time `json:"notice_date"`
    ListingDate     time.Time `json:"listing_date"`
}
```

### 5.3 StockIpoTutorEmItem
```go
type StockIpoTutorEmItem struct {
    Index             int       `json:"index"`
    CompanyName       string    `json:"company_name"`
    TutorOrg          string    `json:"tutor_org"`
    TutorProcessState string    `json:"tutor_process_state"`
    ReportType        string    `json:"report_type"`
    DispatchOrg       string    `json:"dispatch_org"`
    ReportTitle       string    `json:"report_title"`
    RecordDate        time.Time `json:"record_date"`
}
```

### 5.4 StockKcbSseItem
```go
type StockKcbSseItem struct {
    Index          int       `json:"index"`
    CompanyName    string    `json:"company_name"`
    StockCode      string    `json:"stock_code"`
    UpdateTime     time.Time `json:"update_time"`
    AuditApplyDate time.Time `json:"audit_apply_date"`
    RegisterResult string    `json:"register_result"`
    CommitResult   string    `json:"commit_result"`
    CurrentStatus  string    `json:"current_status"`
    Province       string    `json:"province"`
    CsrcCode       string    `json:"csrc_code"`
    StockAuditNum  int       `json:"stock_audit_num"`
}
```

---

## 六、测试结果

### 6.1 StockIpoDeclareEm
- 测试状态: 通过
- 数据量: 3770 条
- 示例数据: 中国科学院沈阳科学仪器股份有限公司, 状态=上市委会议通过

### 6.2 StockIpoReviewEm
- 测试状态: 通过
- 数据量: 5090 条
- 示例数据: 宁波惠康工业科技股份有限公司, 股票简称=惠康科技, 审核状态=未上会

### 6.3 StockIpoTutorEm
- 测试状态: 通过
- 数据量: 5197 条
- 示例数据: 深圳市玩视科技股份有限公司, 辅导机构=招商证券股份有限公司

### 6.4 StockKcbSse
- 测试状态: 通过
- 数据量: 20 条 (单页)
- 示例数据: 深之蓝海洋科技股份有限公司, 当前状态=已问询

---

## 七、完成清单

- [x] 更新 types.go 添加 IPO 相关类型定义
- [x] 创建 stock_fundamental_ipo_em.go 实现 3 个东方财富接口
- [x] 创建 stock_fundamental_kcb_sse.go 实现上交所科创板接口
- [x] 创建 stock_fundamental_ipo_em_test.go 编写测试用例
- [x] 实现 JSONP 解析函数
- [x] 实现分页处理逻辑
- [x] 实现日期解析逻辑
- [x] 实现状态码转换
- [x] 所有测试通过
- [x] 代码格式化
- [x] 静态检查通过

---

## 八、函数签名

```go
// 东方财富首发申报企业信息
func StockIpoDeclareEm() ([]StockIpoDeclareEmItem, error)

// 东方财富过会企业信息
func StockIpoReviewEm() ([]StockIpoReviewEmItem, error)

// 东方财富辅导备案信息
func StockIpoTutorEm() ([]StockIpoTutorEmItem, error)

// 上交所科创板信息
func StockKcbSse() ([]StockKcbSseItem, error)
```

---

*文档版本: v1.0*
*最后更新: 2026-01-18*
