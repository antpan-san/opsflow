package types

import "time"

// Input 定义检测输入参数
type Input struct {
    Target string            // 检测目标（如域名、IP、Pod名称）
    Params map[string]string // 额外参数
}

// Result 定义检测结果
type Result struct {
    Name    string                 // 检测名称
    Success bool                   // 是否成功
    Message string                 // 结果描述
    Data    map[string]interface{} // 额外数据
}

// Scenario 定义诊断场景
type Scenario struct {
    Name        string   // 场景名称
    Input       string   // 输入类型（如 domain、pod）
    Checks      []string // 检测列表
    Rules       []Rule   // 规则列表
    Actions     []Action // 动作列表
    Description string   // 场景描述
}

// Rule 定义诊断规则
type Rule struct {
    Condition   string // 条件表达式（如 "tcp_ok && http_fail")
    Conclusion  string // 诊断结论
    Suggestion  string // 建议措施
    Priority    int    // 优先级（数字越小优先级越高）
}

// Action 定义动作接口
type Action interface {
    Name() string
    Execute(ctx *Context) error
}

// Context 定义执行上下文
type Context struct {
    Input    Input
    Results  map[string]Result
    Scenario *Scenario
    Output   string // 输出格式（text/json）
    Duration time.Duration // 执行耗时
}