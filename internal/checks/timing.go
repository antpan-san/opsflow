package checks

import (
    "net"
    "net/http"
    "time"

    "github.com/yourusername/opsflow/internal/types"
)

// TimingResult 时延检测结果
type TimingResult struct {
    DNSLookup     time.Duration `json:"dns_lookup"`
    TCPConnect    time.Duration `json:"tcp_connect"`
    TLSHandshake  time.Duration `json:"tls_handshake"`
    ServerConnect time.Duration `json:"server_connect"`
    FirstByte     time.Duration `json:"first_byte"`
    Total         time.Duration `json:"total"`
}

// TimingCheck 时延检测
type TimingCheck struct {
    BaseCheck
}

func NewTimingCheck() Check {
    return &TimingCheck{
        BaseCheck: BaseCheck{Name: "timing"},
    }
}

func (t *TimingCheck) Name() string {
    return t.BaseCheck.Name
}

func (t *TimingCheck) Run(input types.Input) types.Result {
    target := input.Target
    port := input.Params["port"]
    protocol := input.Params["protocol"]

    // 默认值
    if port == "" {
        if protocol == "https" {
            port = "443"
        } else {
            port = "80"
        }
    }
    if protocol == "" {
        protocol = "http"
    }

    // 开始计时
    startTime := time.Now()
    var timing TimingResult

    // 1. DNS 解析时延
    dnsStart := time.Now()
    ips, err := net.LookupHost(target)
    timing.DNSLookup = time.Since(dnsStart)
    
    if err != nil {
        return types.Result{
            Name:    "timing",
            Success: false,
            Message: "DNS解析失败: " + err.Error(),
            Data: map[string]interface{}{
                "target":   target,
                "dns_time": timing.DNSLookup.String(),
                "error":    err.Error(),
            },
        }
    }

    if len(ips) == 0 {
        return types.Result{
            Name:    "timing",
            Success: false,
            Message: "DNS解析返回空结果",
            Data: map[string]interface{}{
                "target":   target,
                "dns_time": timing.DNSLookup.String(),
            },
        }
    }

    // 2. TCP 连接时延
    tcpStart := time.Now()
    address := net.JoinHostPort(ips[0], port)
    conn, err := net.DialTimeout("tcp", address, 10*time.Second)
    timing.TCPConnect = time.Since(tcpStart)

    if err != nil {
        timing.Total = time.Since(startTime)
        return types.Result{
            Name:    "timing",
            Success: false,
            Message: "TCP连接失败: " + err.Error(),
            Data: map[string]interface{}{
                "target":     target,
                "dns_lookup": timing.DNSLookup.String(),
                "tcp_connect": timing.TCPConnect.String(),
                "total":      timing.Total.String(),
                "error":      err.Error(),
            },
        }
    }
    defer conn.Close()

    // 3. HTTP 请求时延（如果是 HTTP/HTTPS）
    var firstByte time.Duration
    var serverInfo string
    var statusCode int
    var contentType string

    if protocol == "http" || protocol == "https" {
        httpStart := time.Now()
        
        // 构建请求 URL
        var url string
        if protocol == "https" && port == "443" {
            url = "https://" + target
        } else if protocol == "http" && port == "80" {
            url = "http://" + target
        } else {
            url = protocol + "://" + target + ":" + port
        }

        // 创建 HTTP 客户端
        client := &http.Client{
            Timeout: 10 * time.Second,
            CheckRedirect: func(req *http.Request, via []*http.Request) error {
                return nil
            },
        }

        // 发送请求
        req, err := http.NewRequest("GET", url, nil)
        if err == nil {
            req.Header.Set("User-Agent", "OpsFlow/1.0")
            
            // 计时
            firstByteStart := time.Now()
            resp, err := client.Do(req)
            if err == nil {
                defer resp.Body.Close()
                firstByte = time.Since(firstByteStart)
                
                statusCode = resp.StatusCode
                contentType = resp.Header.Get("Content-Type")
                
                // 获取服务器信息
                if server := resp.Header.Get("Server"); server != "" {
                    serverInfo = server
                }
            }
        }
        
        timing.ServerConnect = time.Since(httpStart) - firstByte
    }

    timing.Total = time.Since(startTime)

    // 构建结果
    data := map[string]interface{}{
        "target":       target,
        "port":         port,
        "protocol":     protocol,
        "dns_lookup":   timing.DNSLookup.String(),
        "tcp_connect":  timing.TCPConnect.String(),
        "total":        timing.Total.String(),
    }

    if firstByte > 0 {
        data["first_byte"] = firstByte.String()
        data["server_connect"] = timing.ServerConnect.String()
    }

    if statusCode > 0 {
        data["status_code"] = statusCode
    }

    if contentType != "" {
        data["content_type"] = contentType
    }

    if serverInfo != "" {
        data["server"] = serverInfo
    }

    return types.Result{
        Name:    "timing",
        Success: true,
        Message: "时延检测完成",
        Data:    data,
    }
}