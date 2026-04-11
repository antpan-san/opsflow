package rules

import (
    "strings"

    "github.com/yourusername/opsflow/internal/types"
)

// Engine 规则引擎
type Engine struct {
    rules []types.Rule
}

// NewEngine 创建规则引擎
func NewEngine() *Engine {
    return &Engine{
        rules: make([]types.Rule, 0),
    }
}

// AddRule 添加规则
func (e *Engine) AddRule(rule types.Rule) {
    e.rules = append(e.rules, rule)
}

// Evaluate 评估规则
func (e *Engine) Evaluate(results map[string]types.Result) types.Rule {
    // 按优先级排序规则
    sortedRules := e.sortRules()

    for _, rule := range sortedRules {
        if e.evaluateCondition(rule.Condition, results) {
            return rule
        }
    }

    // 默认规则
    return types.Rule{
        Conclusion: "无法确定问题原因",
        Suggestion: "请检查系统日志或联系管理员",
    }
}

// evaluateCondition 评估条件表达式
func (e *Engine) evaluateCondition(condition string, results map[string]types.Result) bool {
    // 简单的条件表达式解析
    // 支持: tcp_ok && http_fail, dns_ok || tcp_ok 等

    // 替换条件中的检查结果
    expr := condition
    for name, result := range results {
        placeholder := name + "_ok"
        value := "false"
        if result.Success {
            value = "true"
        }
        expr = strings.ReplaceAll(expr, placeholder, value)
    }

    // 简单的逻辑表达式求值
    return e.evaluateLogicalExpression(expr)
}

// evaluateLogicalExpression 简单的逻辑表达式求值
func (e *Engine) evaluateLogicalExpression(expr string) bool {
    // 处理 AND 操作
    if strings.Contains(expr, "&&") {
        parts := strings.Split(expr, "&&")
        for _, part := range parts {
            if !e.evaluateSimpleExpression(strings.TrimSpace(part)) {
                return false
            }
        }
        return true
    }

    // 处理 OR 操作
    if strings.Contains(expr, "||") {
        parts := strings.Split(expr, "||")
        for _, part := range parts {
            if e.evaluateSimpleExpression(strings.TrimSpace(part)) {
                return true
            }
        }
        return false
    }

    // 单个表达式
    return e.evaluateSimpleExpression(expr)
}

// evaluateSimpleExpression 简单表达式求值
func (e *Engine) evaluateSimpleExpression(expr string) bool {
    return expr == "true"
}

// sortRules 按优先级排序规则
func (e *Engine) sortRules() []types.Rule {
    // 简单的冒泡排序
    sorted := make([]types.Rule, len(e.rules))
    copy(sorted, e.rules)

    for i := 0; i < len(sorted)-1; i++ {
        for j := 0; j < len(sorted)-i-1; j++ {
            if sorted[j].Priority > sorted[j+1].Priority {
                sorted[j], sorted[j+1] = sorted[j+1], sorted[j]
            }
        }
    }

    return sorted
}