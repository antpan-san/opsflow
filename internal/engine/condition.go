package engine

import (
	"fmt"
	"strings"
)

// ConditionParser 条件表达式解析器
type ConditionParser interface {
	Parse(expr string) (Condition, error)
	Evaluate(condition Condition, results map[string]bool) bool
}

// Condition 条件节点
type Condition struct {
	Type     string      // "and", "or", "not", "variable"
	Children []Condition // 子条件
	Variable string      // 变量名
}

// NewConditionParser 创建条件表达式解析器
func NewConditionParser() ConditionParser {
	return &conditionParser{}
}

type conditionParser struct{}

// Parse 解析条件表达式
func (p *conditionParser) Parse(expr string) (Condition, error) {
	expr = strings.TrimSpace(expr)
	if expr == "" {
		return Condition{}, fmt.Errorf("empty expression")
	}

	// 处理括号
	if strings.HasPrefix(expr, "(") && strings.HasSuffix(expr, ")") {
		expr = expr[1 : len(expr)-1]
	}

	// 检查是否包含逻辑运算符
	// 优先检查 AND (空格分隔)
	if strings.Contains(expr, " && ") {
		parts := splitByLogicalOperator(expr, " && ")
		children := make([]Condition, 0, len(parts))
		for _, part := range parts {
			child, err := p.Parse(strings.TrimSpace(part))
			if err != nil {
				return Condition{}, err
			}
			children = append(children, child)
		}
		return Condition{Type: "and", Children: children}, nil
	}

	// 检查 OR
	if strings.Contains(expr, " || ") {
		parts := splitByLogicalOperator(expr, " || ")
		children := make([]Condition, 0, len(parts))
		for _, part := range parts {
			child, err := p.Parse(strings.TrimSpace(part))
			if err != nil {
				return Condition{}, err
			}
			children = append(children, child)
		}
		return Condition{Type: "or", Children: children}, nil
	}

	// 检查 NOT
	if strings.HasPrefix(expr, "!") {
		child, err := p.Parse(strings.TrimSpace(expr[1:]))
		if err != nil {
			return Condition{}, err
		}
		return Condition{Type: "not", Children: []Condition{child}}, nil
	}

	// 变量
	return Condition{Type: "variable", Variable: expr}, nil
}

// Evaluate 评估条件表达式
func (p *conditionParser) Evaluate(condition Condition, results map[string]bool) bool {
	switch condition.Type {
	case "and":
		for _, child := range condition.Children {
			if !p.Evaluate(child, results) {
				return false
			}
		}
		return true
	case "or":
		for _, child := range condition.Children {
			if p.Evaluate(child, results) {
				return true
			}
		}
		return false
	case "not":
		if len(condition.Children) == 1 {
			return !p.Evaluate(condition.Children[0], results)
		}
		return false
	case "variable":
		return results[condition.Variable]
	default:
		return false
	}
}

// splitByLogicalOperator 按逻辑运算符分割表达式
func splitByLogicalOperator(expr, operator string) []string {
	var parts []string
	var current strings.Builder
	depth := 0

	for i := 0; i < len(expr); i++ {
		if expr[i] == '(' {
			depth++
		} else if expr[i] == ')' {
			depth--
		}

		if depth == 0 && i+len(operator) <= len(expr) && expr[i:i+len(operator)] == operator {
			parts = append(parts, current.String())
			current.Reset()
			i += len(operator) - 1
			continue
		}

		current.WriteByte(expr[i])
	}

	if current.Len() > 0 {
		parts = append(parts, current.String())
	}

	return parts
}