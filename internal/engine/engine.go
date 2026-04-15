package engine

import (
    "fmt"
    "strings"
    "time"

    "github.com/yourusername/opsflow/internal/checks"
    "github.com/yourusername/opsflow/internal/rules"
    "github.com/yourusername/opsflow/internal/types"
)

// Engine 诊断引擎
type Engine struct {
    scenarios map[string]types.Scenario
    checkMap  map[string]checks.Check
    ruleEngine *rules.Engine
    colorFormatter *ColorFormatter
    jsonFormatter *JSONFormatter
}

// NewEngine 创建引擎实例
func NewEngine() *Engine {
    return &Engine{
        scenarios: make(map[string]types.Scenario),
        checkMap:  make(map[string]checks.Check),
        ruleEngine: rules.NewEngine(),
        colorFormatter: NewColorFormatter(),
        jsonFormatter: NewJSONFormatter(),
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
        return e.FormatOutputJSON(ctx)
    default:
        // 检查是否有时延检测结果
        if _, exists := ctx.Results["timing"]; exists {
            return e.FormatOutputTiming(ctx)
        }
        return e.FormatOutputText(ctx)
    }
}

// FormatOutputJSON 格式化JSON输出
func (e *Engine) FormatOutputJSON(ctx *types.Context) string {
    // 构建JSON报告
    results := make(map[string]interface{})
    for name, result := range ctx.Results {
        if name == "rule" {
            continue
        }
        results[name] = map[string]interface{}{
            "success": result.Success,
            "message": result.Message,
            "data":    result.Data,
        }
    }

    var conclusion, suggestion string
    if ruleResult, exists := ctx.Results["rule"]; exists {
        conclusion = ruleResult.Data["conclusion"].(string)
        suggestion = ruleResult.Data["suggestion"].(string)
    }

    // 创建JSON报告
    report := e.jsonFormatter.CreateReport(
        ctx.Scenario.Name,
        ctx.Input.Target,
        make(map[string]CheckResult),
        conclusion,
        suggestion,
        ctx.Duration,
    )
    
    // 使用报告数据
    _ = report // 避免未使用变量错误
    
    // 手动构建JSON以避免复杂性

    // 手动构建JSON以避免复杂性
    var output strings.Builder
    output.WriteString("{\n")
    output.WriteString(fmt.Sprintf("  \"scenario\": \"%s\",\n", ctx.Scenario.Name))
    output.WriteString(fmt.Sprintf("  \"target\": \"%s\",\n", ctx.Input.Target))
    output.WriteString(fmt.Sprintf("  \"timestamp\": \"%s\",\n", time.Now().Format(time.RFC3339)))
    output.WriteString("  \"results\": {\n")
    
    first := true
    for name, result := range results {
        if !first {
            output.WriteString(",\n")
        }
        first = false
        
        resultMap := result.(map[string]interface{})
        output.WriteString(fmt.Sprintf("    \"%s\": {\n", name))
        output.WriteString(fmt.Sprintf("      \"success\": %v,\n", resultMap["success"]))
        output.WriteString(fmt.Sprintf("      \"message\": \"%s\"\n", resultMap["message"]))
        output.WriteString("    }")
    }
    
    output.WriteString("\n  },\n")
    output.WriteString(fmt.Sprintf("  \"conclusion\": \"%s\",\n", conclusion))
    output.WriteString(fmt.Sprintf("  \"suggestion\": \"%s\"\n", suggestion))
    output.WriteString("}")
    
    return output.String()
}

// FormatOutputText 格式化文本输出
func (e *Engine) FormatOutputText(ctx *types.Context) string {
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

// FormatOutputColor 格式化彩色输出
func (e *Engine) FormatOutputColor(ctx *types.Context) string {
    var output strings.Builder

    output.WriteString(e.colorFormatter.FormatTitle("=== 诊断报告 ===") + "\n")
    output.WriteString(fmt.Sprintf("场景: %s\n", ctx.Scenario.Name))
    output.WriteString(fmt.Sprintf("目标: %s\n", ctx.Input.Target))
    output.WriteString("\n")

    // 检测结果
    output.WriteString("[检测结果]\n")
    for name, result := range ctx.Results {
        if name == "rule" {
            continue
        }
        
        var line string
        if result.Success {
            line = e.colorFormatter.FormatSuccess(name + ": " + result.Message)
        } else {
            line = e.colorFormatter.FormatError(name + ": " + result.Message)
        }
        output.WriteString(line + "\n")
    }
    output.WriteString("\n")

    // 诊断结论
    if ruleResult, exists := ctx.Results["rule"]; exists {
        output.WriteString("[诊断结论]\n")
        output.WriteString(e.colorFormatter.FormatInfo("结论: " + ruleResult.Data["conclusion"].(string)) + "\n")
        output.WriteString(e.colorFormatter.FormatInfo("建议: " + ruleResult.Data["suggestion"].(string)) + "\n")
    }

    return output.String()
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