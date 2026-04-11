package checks

import (
    "net"
    "strings"
    "time"

    "github.com/yourusername/opsflow/internal/types"
)

// DNSDetailedCheck DNS详细检测
type DNSDetailedCheck struct {
    BaseCheck
}

func NewDNSDetailedCheck() Check {
    return &DNSDetailedCheck{
        BaseCheck: BaseCheck{Name: "dns_detailed"},
    }
}

func (d *DNSDetailedCheck) Name() string {
    return d.BaseCheck.Name
}

func (d *DNSDetailedCheck) Run(input types.Input) types.Result {
    domain := input.Target
    
    // DNS 解析详细信息
    ips, err := net.LookupHost(domain)
    if err != nil {
        return types.Result{
            Name:    "dns_detailed",
            Success: false,
            Message: "DNS解析失败: " + err.Error(),
            Data: map[string]interface{}{
                "domain": domain,
                "error":  err.Error(),
            },
        }
    }

    // 获取 IPv4 和 IPv6 地址
    var ipv4Addresses []string
    var ipv6Addresses []string

    for _, ip := range ips {
        if strings.Contains(ip, ":") {
            ipv6Addresses = append(ipv6Addresses, ip)
        } else {
            ipv4Addresses = append(ipv4Addresses, ip)
        }
    }

    // 检查 CNAME 记录
    cname, err := net.LookupCNAME(domain)
    if err != nil {
        cname = "无 CNAME 记录"
    }

    // 获取 DNS 服务器信息
    dnsServers := []string{}
    // 这里可以添加 DNS 服务器检测逻辑

    return types.Result{
        Name:    "dns_detailed",
        Success: true,
        Message: "DNS解析成功",
        Data: map[string]interface{}{
            "domain":         domain,
            "ips":            ips,
            "ipv4_addresses": ipv4Addresses,
            "ipv6_addresses": ipv6Addresses,
            "cname":          cname,
            "dns_servers":    dnsServers,
            "record_count":   len(ips),
        },
    }
}

// NetworkDetailedCheck 网络详细检测
type NetworkDetailedCheck struct {
    BaseCheck
}

func NewNetworkDetailedCheck() Check {
    return &NetworkDetailedCheck{
        BaseCheck: BaseCheck{Name: "network_detailed"},
    }
}

func (n *NetworkDetailedCheck) Name() string {
    return n.BaseCheck.Name
}

func (n *NetworkDetailedCheck) Run(input types.Input) types.Result {
    // 获取所有网络接口
    interfaces, err := net.Interfaces()
    if err != nil {
        return types.Result{
            Name:    "network_detailed",
            Success: false,
            Message: "无法获取网络接口: " + err.Error(),
            Data:    map[string]interface{}{},
        }
    }

    var activeInterfaces []map[string]interface{}
    var ipv4Addresses []string
    var ipv6Addresses []string
    var hasInternet bool

    for _, iface := range interfaces {
        // 检查接口状态
        isUp := iface.Flags&net.FlagUp != 0
        if !isUp {
            continue
        }

        // 获取接口地址
        addrs, err := iface.Addrs()
        if err != nil {
            continue
        }

        interfaceInfo := map[string]interface{}{
            "name":      iface.Name,
            "flags":     iface.Flags.String(),
            "mtu":       iface.MTU,
            "addresses": []string{},
        }

        for _, addr := range addrs {
            addrStr := addr.String()
            interfaceInfo["addresses"] = append(interfaceInfo["addresses"].([]string), addrStr)

            // 分析 IP 地址类型
            if strings.Contains(addrStr, ":") {
                ipv6Addresses = append(ipv6Addresses, addrStr)
            } else {
                ipv4Addresses = append(ipv4Addresses, addrStr)
            }
        }

        activeInterfaces = append(activeInterfaces, interfaceInfo)
    }

    // 检查互联网连接
    testHosts := []string{"8.8.8.8", "1.1.1.1"}
    for _, host := range testHosts {
        conn, err := net.DialTimeout("tcp", host+":80", 3*time.Second)
        if err == nil {
            conn.Close()
            hasInternet = true
            break
        }
    }

    return types.Result{
        Name:    "network_detailed",
        Success: true,
        Message: "网络检测完成",
        Data: map[string]interface{}{
            "active_interfaces": activeInterfaces,
            "ipv4_addresses":    ipv4Addresses,
            "ipv6_addresses":    ipv6Addresses,
            "has_internet":      hasInternet,
            "interface_count":   len(activeInterfaces),
        },
    }
}

// InternalNetworkCheck 内网检测
type InternalNetworkCheck struct {
    BaseCheck
}

func NewInternalNetworkCheck() Check {
    return &InternalNetworkCheck{
        BaseCheck: BaseCheck{Name: "internal_network"},
    }
}

func (i *InternalNetworkCheck) Name() string {
    return i.BaseCheck.Name
}

func (i *InternalNetworkCheck) Run(input types.Input) types.Result {
    // 检测内网特征
    var isInternal bool
    var internalIPs []string
    var networkType string

    // 获取本地 IP 地址
    interfaces, err := net.Interfaces()
    if err != nil {
        return types.Result{
            Name:    "internal_network",
            Success: false,
            Message: "无法获取网络接口",
        }
    }

    for _, iface := range interfaces {
        if iface.Flags&net.FlagUp == 0 {
            continue
        }

        addrs, err := iface.Addrs()
        if err != nil {
            continue
        }

        for _, addr := range addrs {
            addrStr := addr.String()
            ip, _, _ := net.ParseCIDR(addrStr)

            if ip != nil {
                ipStr := ip.String()

                // 检查是否为内网 IP
                if isPrivateIP(ip) {
                    isInternal = true
                    internalIPs = append(internalIPs, ipStr)

                    // 判断网络类型
                    if strings.HasPrefix(ipStr, "192.168.") {
                        networkType = "家庭/办公网络"
                    } else if strings.HasPrefix(ipStr, "10.") {
                        networkType = "企业内网"
                    } else if strings.HasPrefix(ipStr, "172.16.") || strings.HasPrefix(ipStr, "172.17.") ||
                        strings.HasPrefix(ipStr, "172.18.") || strings.HasPrefix(ipStr, "172.19.") ||
                        strings.HasPrefix(ipStr, "172.20.") || strings.HasPrefix(ipStr, "172.21.") ||
                        strings.HasPrefix(ipStr, "172.22.") || strings.HasPrefix(ipStr, "172.23.") ||
                        strings.HasPrefix(ipStr, "172.24.") || strings.HasPrefix(ipStr, "172.25.") ||
                        strings.HasPrefix(ipStr, "172.26.") || strings.HasPrefix(ipStr, "172.27.") ||
                        strings.HasPrefix(ipStr, "172.28.") || strings.HasPrefix(ipStr, "172.29.") ||
                        strings.HasPrefix(ipStr, "172.30.") || strings.HasPrefix(ipStr, "172.31.") {
                        networkType = "企业内网"
                    }
                }
            }
        }
    }

    return types.Result{
        Name:    "internal_network",
        Success: true,
        Message: "内网检测完成",
        Data: map[string]interface{}{
            "is_internal":   isInternal,
            "internal_ips":  internalIPs,
            "network_type":  networkType,
            "detection_time": time.Now().Format("2006-01-02 15:04:05"),
        },
    }
}

// isPrivateIP 检查是否为私有 IP 地址
func isPrivateIP(ip net.IP) bool {
    if ip.IsLoopback() || ip.IsLinkLocalUnicast() || ip.IsLinkLocalMulticast() {
        return true
    }

    ipv4 := ip.To4()
    if ipv4 != nil {
        return ipv4[0] == 10 ||
            (ipv4[0] == 172 && ipv4[1] >= 16 && ipv4[1] <= 31) ||
            (ipv4[0] == 192 && ipv4[1] == 168)
    }

    return false
}