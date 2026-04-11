package engine

import (
    "fmt"
    "strings"

    "github.com/yourusername/opsflow/internal/types"
)

// FormatOutputDetailed 格式化详细输出
func (e *Engine) FormatOutputDetailed(ctx *types.Context) string {
    var output strings.Builder

    output.WriteString(fmt.Sprintf("=== Web 详细诊断报告 ===\n"))
    output.WriteString(fmt.Sprintf("场景: %s\n", ctx.Scenario.Name))
    output.WriteString(fmt.Sprintf("目标: %s\n", ctx.Input.Target))
    output.WriteString(fmt.Sprintf("时间: %s\n", ctx.Input.Params["timestamp"]))
    output.WriteString("\n")

    // 网络详细信息
    if networkResult, exists := ctx.Results["network_detailed"]; exists {
        output.WriteString("[网络接口信息]\n")
        if networkResult.Success {
            if interfaces, ok := networkResult.Data["active_interfaces"].([]map[string]interface{}); ok {
                for _, iface := range interfaces {
                    output.WriteString(fmt.Sprintf("  接口: %s\n", iface["name"]))
                    output.WriteString(fmt.Sprintf("    状态: %s\n", iface["flags"]))
                    output.WriteString(fmt.Sprintf("    MTU: %d\n", iface["mtu"]))
                    if addrs, ok := iface["addresses"].([]string); ok {
                        output.WriteString(fmt.Sprintf("    地址: %s\n", strings.Join(addrs, ", ")))
                    }
                }
            }
        }
        output.WriteString("\n")
    }

    // 内网信息
    if internalResult, exists := ctx.Results["internal_network"]; exists {
        output.WriteString("[内网信息]\n")
        if internalResult.Success {
            if isInternal, ok := internalResult.Data["is_internal"].(bool); ok && isInternal {
                output.WriteString(fmt.Sprintf("  网络类型: %s\n", internalResult.Data["network_type"]))
                if internalIPs, ok := internalResult.Data["internal_ips"].([]string); ok {
                    output.WriteString(fmt.Sprintf("  内网IP: %s\n", strings.Join(internalIPs, ", ")))
                }
            } else {
                output.WriteString("  网络类型: 公网\n")
            }
        }
        output.WriteString("\n")
    }

    // DNS 详细信息
    if dnsResult, exists := ctx.Results["dns_detailed"]; exists {
        output.WriteString("[DNS 解析信息]\n")
        if dnsResult.Success {
            if domain, ok := dnsResult.Data["domain"].(string); ok {
                output.WriteString(fmt.Sprintf("  域名: %s\n", domain))
            }
            if cname, ok := dnsResult.Data["cname"].(string); ok && cname != "无 CNAME 记录" {
                output.WriteString(fmt.Sprintf("  CNAME: %s\n", cname))
            }
            if ipv4Addrs, ok := dnsResult.Data["ipv4_addresses"].([]string); ok && len(ipv4Addrs) > 0 {
                output.WriteString(fmt.Sprintf("  IPv4: %s\n", strings.Join(ipv4Addrs, ", ")))
            }
            if ipv6Addrs, ok := dnsResult.Data["ipv6_addresses"].([]string); ok && len(ipv6Addrs) > 0 {
                output.WriteString(fmt.Sprintf("  IPv6: %s\n", strings.Join(ipv6Addrs, ", ")))
            }
            if recordCount, ok := dnsResult.Data["record_count"].(int); ok {
                output.WriteString(fmt.Sprintf("  记录数: %d\n", recordCount))
            }
        } else {
            output.WriteString(fmt.Sprintf("  错误: %s\n", dnsResult.Message))
        }
        output.WriteString("\n")
    }

    // TCP 连接信息
    if tcpResult, exists := ctx.Results["tcp"]; exists {
        output.WriteString("[TCP 连接信息]\n")
        status := "❌"
        if tcpResult.Success {
            status = "✅"
        }
        output.WriteString(fmt.Sprintf("  状态: %s %s\n", status, tcpResult.Message))
        if port, ok := tcpResult.Data["port"].(string); ok {
            output.WriteString(fmt.Sprintf("  端口: %s\n", port))
        }
        output.WriteString("\n")
    }

    // HTTP 协议信息
    if httpResult, exists := ctx.Results["http_protocol"]; exists {
        output.WriteString("[HTTP 协议信息]\n")
        status := "❌"
        if httpResult.Success {
            status = "✅"
        }
        output.WriteString(fmt.Sprintf("  状态: %s %s\n", status, httpResult.Message))
        
        if statusCode, ok := httpResult.Data["status_code"].(int); ok {
            output.WriteString(fmt.Sprintf("  状态码: %d\n", statusCode))
        }
        if contentType, ok := httpResult.Data["content_type"].(string); ok && contentType != "" {
            output.WriteString(fmt.Sprintf("  内容类型: %s\n", contentType))
        }
        if server, ok := httpResult.Data["server"].(string); ok && server != "" {
            output.WriteString(fmt.Sprintf("  服务器: %s\n", server))
        }
        if url, ok := httpResult.Data["url"].(string); ok {
            output.WriteString(fmt.Sprintf("  URL: %s\n", url))
        }
        output.WriteString("\n")
    }

    // 时延信息
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