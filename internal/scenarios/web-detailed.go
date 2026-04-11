package scenarios

import (
    "github.com/yourusername/opsflow/internal/checks"
    "github.com/yourusername/opsflow/internal/types"
)

// WebDetailedScenario Web详细诊断场景
func WebDetailedScenario() types.Scenario {
    return types.Scenario{
        Name:        "web-detailed",
        Input:       "domain",
        Checks:      []string{"network_detailed", "internal_network", "dns_detailed", "tcp", "http_protocol"},
        Description: "Web服务详细诊断场景",
        Rules: []types.Rule{
            {
                Condition:  "network_detailed_ok && internal_network_ok && dns_detailed_ok && tcp_ok && http_protocol_ok",
                Conclusion: "Web服务完全正常",
                Suggestion: "无需操作",
                Priority:   1,
            },
            {
                Condition:  "network_detailed_ok && internal_network_ok && dns_detailed_ok && tcp_ok && http_protocol_fail",
                Conclusion: "Web服务异常（HTTP协议问题）",
                Suggestion: "检查后端服务状态、Web服务器配置和证书",
                Priority:   2,
            },
            {
                Condition:  "network_detailed_ok && internal_network_ok && dns_detailed_ok && tcp_fail",
                Conclusion: "网络连接异常（端口不通）",
                Suggestion: "检查防火墙规则、安全组配置和后端服务状态",
                Priority:   3,
            },
            {
                Condition:  "network_detailed_ok && internal_network_ok && dns_detailed_fail",
                Conclusion: "DNS解析失败",
                Suggestion: "检查域名配置、DNS服务器设置和域名是否过期",
                Priority:   4,
            },
            {
                Condition:  "network_detailed_ok && internal_network_fail",
                Conclusion: "内网网络异常",
                Suggestion: "检查本地网络配置、网关设置",
                Priority:   5,
            },
            {
                Condition:  "network_detailed_fail",
                Conclusion: "网络接口异常",
                Suggestion: "检查网络接口、网线连接和本地网络配置",
                Priority:   6,
            },
        },
    }
}

// RegisterWebDetailedScenario 注册详细Web场景
func RegisterWebDetailedScenario(engine interface {
    RegisterScenario(types.Scenario)
    RegisterCheck(checks.Check)
}) {
    // 注册场景
    engine.RegisterScenario(WebDetailedScenario())
    
    // 注册详细检测
    engine.RegisterCheck(checks.NewNetworkDetailedCheck())  // 详细网络检测
    engine.RegisterCheck(checks.NewInternalNetworkCheck())  // 内网检测
    engine.RegisterCheck(checks.NewDNSDetailedCheck())      // 详细DNS检测
    engine.RegisterCheck(checks.NewTCPCheck())              // TCP检测
    engine.RegisterCheck(checks.NewHTTPProtocolCheck())     // HTTP协议检测
}