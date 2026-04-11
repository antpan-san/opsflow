package checks

import (
    "net"
    "time"

    "github.com/yourusername/opsflow/internal/types"
)

// DNSCheck DNS检测
type DNSCheck struct {
    BaseCheck
}

func NewDNSCheck() Check {
    return &DNSCheck{
        BaseCheck: BaseCheck{Name: "dns"},
    }
}

func (d *DNSCheck) Name() string {
    return d.BaseCheck.Name
}

func (d *DNSCheck) Run(input types.Input) types.Result {
    domain := input.Target
    
    // 解析DNS
    ips, err := net.LookupHost(domain)
    if err != nil {
        return types.Result{
            Name:    "dns",
            Success: false,
            Message: "DNS解析失败: " + err.Error(),
            Data:    map[string]interface{}{"domain": domain},
        }
    }

    return types.Result{
        Name:    "dns",
        Success: true,
        Message: "DNS解析成功",
        Data: map[string]interface{}{
            "domain": domain,
            "ips":    ips,
        },
    }
}

// TCPCheck TCP连接检测
type TCPCheck struct {
    BaseCheck
}

func NewTCPCheck() Check {
    return &TCPCheck{
        BaseCheck: BaseCheck{Name: "tcp"},
    }
}

func (t *TCPCheck) Name() string {
    return t.BaseCheck.Name
}

func (t *TCPCheck) Run(input types.Input) types.Result {
    domain := input.Target
    port := input.Params["port"]
    if port == "" {
        port = "80"
    }

    // 尝试TCP连接
    conn, err := net.DialTimeout("tcp", domain+":"+port, 5*time.Second)
    if err != nil {
        return types.Result{
            Name:    "tcp",
            Success: false,
            Message: "TCP连接失败: " + err.Error(),
            Data:    map[string]interface{}{"domain": domain, "port": port},
        }
    }
    defer conn.Close()

    return types.Result{
        Name:    "tcp",
        Success: true,
        Message: "TCP连接成功",
        Data: map[string]interface{}{
            "domain": domain,
            "port":   port,
        },
    }
}

// HTTPCheck HTTP检测（使用新的 HTTPProtocolCheck）
type HTTPCheck struct {
    BaseCheck
}

func NewHTTPCheck() Check {
    return &HTTPCheck{
        BaseCheck: BaseCheck{Name: "http"},
    }
}

func (h *HTTPCheck) Name() string {
    return h.BaseCheck.Name
}

func (h *HTTPCheck) Run(input types.Input) types.Result {
    // 调用新的 HTTPProtocolCheck
    protocolCheck := NewHTTPProtocolCheck()
    return protocolCheck.Run(input)
}// Test code change
// Test GitHub push
// Final test
// Test main branch
// Test main branch deployment
// Test code push skill
