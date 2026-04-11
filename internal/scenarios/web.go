package scenarios

import (
    "github.com/yourusername/opsflow/internal/checks"
    "github.com/yourusername/opsflow/internal/types"
)

// WebScenario Web诊断场景
func WebScenario() types.Scenario {
    return types.Scenario{
        Name:        "web",
        Input:       "domain",
        Checks:      []string{"network", "internet", "dns", "tcp", "http_protocol"},
        Description: "Web服务诊断场景",
        Rules: []types.Rule{
            {
                Condition:  "network_ok && internet_ok && dns_ok && tcp_ok && http_protocol_ok",
                Conclusion: "Web服务完全正常",
                Suggestion: "无需操作",
                Priority:   1,
            },
            {
                Condition:  "network_ok && internet_ok && dns_ok && tcp_ok && http_protocol_fail",
                Conclusion: "Web服务异常（HTTP协议问题）",
                Suggestion: "检查后端服务状态、Web服务器配置和证书",
                Priority:   2,
            },
            {
                Condition:  "network_ok && internet_ok && dns_ok && tcp_fail",
                Conclusion: "网络连接异常（端口不通）",
                Suggestion: "检查防火墙规则、安全组配置和后端服务状态",
                Priority:   3,
            },
            {
                Condition:  "network_ok && internet_ok && dns_fail",
                Conclusion: "DNS解析失败",
                Suggestion: "检查域名配置、DNS服务器设置和域名是否过期",
                Priority:   4,
            },
            {
                Condition:  "network_ok && internet_fail",
                Conclusion: "互联网连接异常",
                Suggestion: "检查本地网络配置、网关设置和出口防火墙",
                Priority:   5,
            },
            {
                Condition:  "network_fail",
                Conclusion: "本地网络异常",
                Suggestion: "检查网络接口、网线连接和本地网络配置",
                Priority:   6,
            },
        },
    }
}

// RegisterWebScenario 注册Web场景
func RegisterWebScenario(engine interface {
    RegisterScenario(types.Scenario)
    RegisterCheck(checks.Check)
}) {
    // 注册场景
    engine.RegisterScenario(WebScenario())
    
    // 注册检测
    engine.RegisterCheck(checks.NewNetworkCheck())    // 新增：本地网络检测
    engine.RegisterCheck(checks.NewInternetCheck())   // 新增：互联网检测
    engine.RegisterCheck(checks.NewDNSCheck())        // DNS检测
    engine.RegisterCheck(checks.NewTCPCheck())        // TCP检测
    engine.RegisterCheck(checks.NewHTTPProtocolCheck()) // 新增：HTTP协议检测
}