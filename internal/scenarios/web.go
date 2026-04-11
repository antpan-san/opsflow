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
        Checks:      []string{"dns", "tcp", "http"},
        Description: "Web服务诊断场景",
        Rules: []types.Rule{
            {
                Condition:  "dns_ok && tcp_ok && http_ok",
                Conclusion: "Web服务正常",
                Suggestion: "无需操作",
                Priority:   1,
            },
            {
                Condition:  "dns_ok && tcp_ok && http_fail",
                Conclusion: "Web服务异常",
                Suggestion: "检查后端服务状态和配置",
                Priority:   2,
            },
            {
                Condition:  "dns_ok && tcp_fail",
                Conclusion: "网络连接异常",
                Suggestion: "检查防火墙规则和网络配置",
                Priority:   3,
            },
            {
                Condition:  "dns_fail",
                Conclusion: "DNS解析失败",
                Suggestion: "检查域名配置和DNS服务器",
                Priority:   4,
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
    engine.RegisterCheck(checks.NewDNSCheck())
    engine.RegisterCheck(checks.NewTCPCheck())
    engine.RegisterCheck(checks.NewHTTPCheck())
}