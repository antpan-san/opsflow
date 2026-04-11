package checks

import (
    "net"
    "net/http"
    "time"

    "github.com/yourusername/opsflow/internal/types"
)

// NetworkCheck 本地网络检测
type NetworkCheck struct {
    BaseCheck
}

func NewNetworkCheck() Check {
    return &NetworkCheck{
        BaseCheck: BaseCheck{Name: "network"},
    }
}

func (n *NetworkCheck) Name() string {
    return n.BaseCheck.Name
}

func (n *NetworkCheck) Run(input types.Input) types.Result {
    // 检测本地网络接口
    interfaces, err := net.Interfaces()
    if err != nil {
        return types.Result{
            Name:    "network",
            Success: false,
            Message: "无法获取网络接口: " + err.Error(),
            Data:    map[string]interface{}{},
        }
    }

    // 检查是否有活跃的网络接口
    var hasActiveInterface bool
    var activeInterfaces []string

    for _, iface := range interfaces {
        // 检查接口是否启用
        if iface.Flags&net.FlagUp != 0 {
            hasActiveInterface = true
            activeInterfaces = append(activeInterfaces, iface.Name)
        }
    }

    if !hasActiveInterface {
        return types.Result{
            Name:    "network",
            Success: false,
            Message: "没有活跃的网络接口",
            Data:    map[string]interface{}{},
        }
    }

    // 尝试连接本地网关（如果有）
    var canReachGateway bool
    gateway := input.Params["gateway"]
    if gateway != "" {
        conn, err := net.DialTimeout("tcp", gateway+":80", 3*time.Second)
        if err == nil {
            conn.Close()
            canReachGateway = true
        }
    }

    return types.Result{
        Name:    "network",
        Success: true,
        Message: "本地网络正常",
        Data: map[string]interface{}{
            "interfaces":       activeInterfaces,
            "can_reach_gateway": canReachGateway,
        },
    }
}

// InternetCheck 互联网连通性检测
type InternetCheck struct {
    BaseCheck
}

func NewInternetCheck() Check {
    return &InternetCheck{
        BaseCheck: BaseCheck{Name: "internet"},
    }
}

func (i *InternetCheck) Name() string {
    return i.BaseCheck.Name
}

func (i *InternetCheck) Run(input types.Input) types.Result {
    // 测试多个公共 DNS 服务
    testHosts := []string{
        "8.8.8.8",           // Google DNS
        "1.1.1.1",           // Cloudflare DNS
        "www.baidu.com",     // 国内常用网站
        "www.google.com",    // 国际常用网站
    }

    var successCount int
    var results []map[string]interface{}

    for _, host := range testHosts {
        conn, err := net.DialTimeout("tcp", host+":80", 5*time.Second)
        if err == nil {
            conn.Close()
            successCount++
            results = append(results, map[string]interface{}{
                "host":   host,
                "status": "success",
            })
        } else {
            results = append(results, map[string]interface{}{
                "host":   host,
                "status": "failed",
                "error":  err.Error(),
            })
        }
    }

    // 如果超过一半的测试成功，认为互联网正常
    success := successCount > len(testHosts)/2

    message := "互联网连接正常"
    if !success {
        message = "互联网连接异常"
    }

    return types.Result{
        Name:    "internet",
        Success: success,
        Message: message,
        Data: map[string]interface{}{
            "test_results": results,
            "success_count": successCount,
            "total_tests":   len(testHosts),
        },
    }
}

// HTTPProtocolCheck HTTP协议检测
type HTTPProtocolCheck struct {
    BaseCheck
}

func NewHTTPProtocolCheck() Check {
    return &HTTPProtocolCheck{
        BaseCheck: BaseCheck{Name: "http_protocol"},
    }
}

func (h *HTTPProtocolCheck) Name() string {
    return h.BaseCheck.Name
}

func (h *HTTPProtocolCheck) Run(input types.Input) types.Result {
    domain := input.Target
    protocol := input.Params["protocol"]
    if protocol == "" {
        protocol = "http"
    }

    url := protocol + "://" + domain
    
    // 创建 HTTP 客户端
    client := &http.Client{
        Timeout: 10 * time.Second,
        CheckRedirect: func(req *http.Request, via []*http.Request) error {
            return nil // 允许重定向
        },
    }
    
    // 发送 HTTP 请求
    resp, err := client.Get(url)
    if err != nil {
        return types.Result{
            Name:    "http_protocol",
            Success: false,
            Message: "HTTP协议检测失败: " + err.Error(),
            Data: map[string]interface{}{
                "url":      url,
                "protocol": protocol,
            },
        }
    }
    defer resp.Body.Close()

    // 分析响应
    statusCode := resp.StatusCode
    statusCategory := ""
    
    switch {
    case statusCode >= 200 && statusCode < 300:
        statusCategory = "成功"
    case statusCode >= 300 && statusCode < 400:
        statusCategory = "重定向"
    case statusCode >= 400 && statusCode < 500:
        statusCategory = "客户端错误"
    case statusCode >= 500:
        statusCategory = "服务器错误"
    }

    // 检查响应头
    contentType := resp.Header.Get("Content-Type")
    server := resp.Header.Get("Server")
    
    success := statusCode < 400

    message := "HTTP协议正常"
    if !success {
        message = "HTTP协议异常 (" + resp.Status + ")"
    }

    return types.Result{
        Name:    "http_protocol",
        Success: success,
        Message: message,
        Data: map[string]interface{}{
            "url":           url,
            "status_code":   statusCode,
            "status_text":   resp.Status,
            "status_category": statusCategory,
            "content_type":  contentType,
            "server":        server,
            "protocol":      protocol,
        },
    }
}