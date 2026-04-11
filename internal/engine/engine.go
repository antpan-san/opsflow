package engine

import (
    "fmt"
    "strings"

    "github.com/yourusername/opsflow/internal/checks"
    "github.com/yourusername/opsflow/internal/rules"
    "github.com/yourusername/opsflow/internal/types"
)

// Engine 诊断引擎
type Engine struct {
    scenarios map[string]types.Scenario
    checkMap  map[string]checks.Check
    ruleEngine *rules.Engine
}

// NewEngine 创建引擎实例
func NewEngine() *Engine {
    return &Engine{
        scenarios: make(map[string]types.Scenario),
        checkMap:  make(map[string]checks.Check),
        ruleEngine: rules.NewEngine(),
    }
}

// RegisterScenario 注册场景
func (e *Engine) RegisterScenario(scenario types.Scenario) {
    e.scenarios[scenario.Name] = scenario
    
    // 注册场景中的规则
    for _, rule := range scenario.Rules {
        e.ruleEngine.AddRule(rule)
    }
}

// RegisterCheck 注册检测
func (e *Engine) RegisterCheck(check checks.Check) {
    e.checkMap[check.Name()] = check
}

// Run 运行诊断
func (e *Engine) Run(scenarioName string, input types.Input) (*types.Context, error) {
    scenario, exists := e.scenarios[scenarioName]
    if !exists {
        return nil, fmt.Errorf("场景不存在: %s", scenarioName)
    }

    ctx := &types.Context{
        Input:    input,
        Results:  make(map[string]types.Result),
        Scenario: &scenario,
    }

    // 执行所有检测
    for _, checkName := range scenario.Checks {
        check, exists := e.checkMap[checkName]
        if !exists {
            return nil, fmt.Errorf("检测不存在: %s", checkName)
        }

        result := check.Run(input)
        ctx.Results[checkName] = result
    }

    // 应用规则
    rule := e.ruleEngine.Evaluate(ctx.Results)
    ctx.Results["rule"] = types.Result{
        Name:    "rule_evaluation",
        Success: true,
        Message: rule.Conclusion,
        Data: map[string]interface{}{
            "conclusion":  rule.Conclusion,
            "suggestion":  rule.Suggestion,
        },
    }

    return ctx, nil
}

// FormatOutput 格式化输出
func (e *Engine) FormatOutput(ctx *types.Context, format string) string {
    switch format {
    case "json":
        return e.formatJSON(ctx)
    default:
        return e.formatText(ctx)
    }
}

// formatText 格式化文本输出
func (e *Engine) formatText(ctx *types.Context) string {
    var output strings.Builder

    output.WriteString(fmt.Sprintf("=== 诊断报告 ===\n"))
    output.WriteString(fmt.Sprintf("场景: %s\n", ctx.Scenario.Name))
    output.WriteString(fmt.Sprintf("目标: %s\n", ctx.Input.Target))
    output.WriteString("\n")

    // 检测结果
    output.WriteString("[检测结果]\n")
    for name, result := range ctx.Results {
        if name == "rule" {
            continue
        }
        status := "❌"
        if result.Success {
            status = "✅"
        }
        output.WriteString(fmt.Sprintf("%s %s: %s\n", status, name, result.Message))
    }
    output.WriteString("\n")

    // 诊断结论
    if ruleResult, exists := ctx.Results["rule"]; exists {
        output.WriteString("[诊断结论]\n")
        output.WriteString(fmt.Sprintf("结论: %s\n", ruleResult.Data["conclusion"]))
        output.WriteString(fmt.Sprintf("建议: %s\n", ruleResult.Data["suggestion"]))
    }

    return output.String()
}

// formatJSON 格式化JSON输出
func (e *Engine) formatJSON(ctx *types.Context) string {
    // 简化的JSON格式化
    var output strings.Builder
    output.WriteString("{\n")
    output.WriteString(fmt.Sprintf("  \"scenario\": \"%s\",\n", ctx.Scenario.Name))
    output.WriteString(fmt.Sprintf("  \"target\": \"%s\",\n", ctx.Input.Target))
    output.WriteString("  \"results\": {\n")
    
    first := true
    for name, result := range ctx.Results {
        if !first {
            output.WriteString(",\n")
        }
        first = false
        
        output.WriteString(fmt.Sprintf("    \"%s\": {\n", name))
        output.WriteString(fmt.Sprintf("      \"success\": %v,\n", result.Success))
        output.WriteString(fmt.Sprintf("      \"message\": \"%s\"\n", result.Message))
        output.WriteString("    }")
    }
    
    output.WriteString("\n  }\n")
    output.WriteString("}")
    
    return output.String()
}