package engine

import (
    "fmt"
    "strings"

    "github.com/yourusername/opsflow/internal/types"
)

// FormatOutputTiming 格式化时延输出
func (e *Engine) FormatOutputTiming(ctx *types.Context) string {
    var output strings.Builder

    output.WriteString(fmt.Sprintf("=== Web 诊断报告 ===\n"))
    output.WriteString(fmt.Sprintf("场景: %s\n", ctx.Scenario.Name))
    output.WriteString(fmt.Sprintf("目标: %s\n", ctx.Input.Target))
    output.WriteString(fmt.Sprintf("时间: %s\n", ctx.Input.Params["timestamp"]))
    output.WriteString("\n")

    // 检测结果摘要
    output.WriteString("[检测结果]\n")
    for checkName, result := range ctx.Results {
        if checkName == "rule" {
            continue
        }
        status := "❌"
        if result.Success {
            status = "✅"
        }
        output.WriteString(fmt.Sprintf("%s %s: %s\n", status, checkName, result.Message))
    }
    output.WriteString("\n")

    // 时延详细信息
    if timingResult, exists := ctx.Results["timing"]; exists && timingResult.Success {
        output.WriteString("[时延信息]\n")
        
        if dnsTime, ok := timingResult.Data["dns_lookup"].(string); ok {
            output.WriteString(fmt.Sprintf("  DNS解析: %s\n", dnsTime))
        }
        
        if tcpTime, ok := timingResult.Data["tcp_connect"].(string); ok {
            output.WriteString(fmt.Sprintf("  TCP连接: %s\n", tcpTime))
        }
        
        if firstByte, ok := timingResult.Data["first_byte"].(string); ok {
            output.WriteString(fmt.Sprintf("  首字节: %s\n", firstByte))
        }
        
        if serverConnect, ok := timingResult.Data["server_connect"].(string); ok {
            output.WriteString(fmt.Sprintf("  服务连接: %s\n", serverConnect))
        }
        
        if totalTime, ok := timingResult.Data["total"].(string); ok {
            output.WriteString(fmt.Sprintf("  总时长: %s\n", totalTime))
        }
        
        // HTTP 详细信息
        if statusCode, ok := timingResult.Data["status_code"].(int); ok {
            output.WriteString(fmt.Sprintf("  状态码: %d\n", statusCode))
        }
        
        if contentType, ok := timingResult.Data["content_type"].(string); ok && contentType != "" {
            output.WriteString(fmt.Sprintf("  内容类型: %s\n", contentType))
        }
        
        if server, ok := timingResult.Data["server"].(string); ok && server != "" {
            output.WriteString(fmt.Sprintf("  服务器: %s\n", server))
        }
        
        output.WriteString("\n")
    }

    // 诊断结论
    if ruleResult, exists := ctx.Results["rule"]; exists {
        output.WriteString("[诊断结论]\n")
        output.WriteString(fmt.Sprintf("结论: %s\n", ruleResult.Data["conclusion"]))
        output.WriteString(fmt.Sprintf("建议: %s\n", ruleResult.Data["suggestion"]))
    }

    return output.String()
}