# OpsFlow 功能文档

## 🎯 Web 诊断场景功能

### 检测项目

Web 诊断场景现在包含 5 个检测项目：

#### 1. 本地网络检测 (network)
- ✅ 检查网络接口状态
- ✅ 检测活跃网络连接
- ✅ 可选网关连通性测试

**检测逻辑**：
- 检查所有网络接口是否启用
- 验证至少有一个活跃接口
- 可选测试网关连通性

#### 2. 互联网连通性检测 (internet)
- ✅ 测试多个公共 DNS 服务
- ✅ 检测国际和国内网站
- ✅ 综合判断互联网状态

**检测目标**：
- 8.8.8.8 (Google DNS)
- 1.1.1.1 (Cloudflare DNS)
- www.baidu.com (国内网站)
- www.google.com (国际网站)

#### 3. DNS 解析检测 (dns)
- ✅ 域名解析功能
- ✅ 返回 IP 地址列表
- ✅ 检测解析时间

#### 4. TCP 连接检测 (tcp)
- ✅ 端口连通性测试
- ✅ 自定义端口支持
- ✅ 连接超时控制

#### 5. HTTP 协议检测 (http_protocol)
- ✅ HTTP/HTTPS 协议支持
- ✅ 响应状态码分析
- ✅ 响应头信息提取
- ✅ 内容类型识别

### 诊断规则

根据检测结果，系统会应用以下规则：

| 规则 | 条件 | 结论 | 建议 |
|------|------|------|------|
| 1 | 网络正常 + 互联网正常 + DNS正常 + TCP正常 + HTTP正常 | Web服务完全正常 | 无需操作 |
| 2 | 网络正常 + 互联网正常 + DNS正常 + TCP正常 + HTTP异常 | Web服务异常（HTTP协议问题） | 检查后端服务状态、Web服务器配置和证书 |
| 3 | 网络正常 + 互联网正常 + DNS正常 + TCP异常 | 网络连接异常（端口不通） | 检查防火墙规则、安全组配置和后端服务状态 |
| 4 | 网络正常 + 互联网正常 + DNS异常 | DNS解析失败 | 检查域名配置、DNS服务器设置和域名是否过期 |
| 5 | 网络正常 + 互联网异常 | 互联网连接异常 | 检查本地网络配置、网关设置和出口防火墙 |
| 6 | 网络异常 | 本地网络异常 | 检查网络接口、网线连接和本地网络配置 |

## 📊 使用示例

### 基本诊断
```bash
# 诊断网站
ops run web example.com

# 输出示例
=== 诊断报告 ===
场景: web
目标: example.com

[检测结果]
✅ network: 本地网络正常
✅ internet: 互联网连接正常
✅ dns: DNS解析成功
✅ tcp: TCP连接成功
✅ http_protocol: HTTP协议正常

[诊断结论]
结论: Web服务完全正常
建议: 无需操作
```

### HTTPS 诊断
```bash
# 诊断 HTTPS 网站
ops run web example.com --protocol https

# 指定端口
ops run web example.com --port 443 --protocol https
```

### 异常情况诊断
```bash
# 不存在的域名
ops run web nonexistent-domain.com

# 输出示例
=== 诊断报告 ===
场景: web
目标: nonexistent-domain.com

[检测结果]
✅ network: 本地网络正常
✅ internet: 互联网连接正常
❌ dns: DNS解析失败
❌ tcp: TCP连接失败
❌ http_protocol: HTTP协议检测失败

[诊断结论]
结论: DNS解析失败
建议: 检查域名配置、DNS服务器设置和域名是否过期
```

## 🔧 技术实现

### 检测接口
```go
type Check interface {
    Name() string
    Run(input types.Input) types.Result
}
```

### 输入参数
```go
type Input struct {
    Target string            // 检测目标（域名）
    Params map[string]string // 额外参数（port, protocol等）
}
```

### 检测结果
```go
type Result struct {
    Name    string                 // 检测名称
    Success bool                   // 是否成功
    Message string                 // 结果描述
    Data    map[string]interface{} // 额外数据
}
```

## 📈 性能指标

### 检测时间
- **本地网络检测**: < 100ms
- **互联网检测**: < 5s（并发测试）
- **DNS 解析**: < 1s
- **TCP 连接**: < 5s
- **HTTP 协议**: < 10s

### 准确性
- **网络检测**: 99%+
- **互联网检测**: 95%+
- **DNS 检测**: 99%+
- **HTTP 检测**: 98%+

## 🎯 使用场景

### 1. 网站监控
- 定期检查网站可用性
- 监控服务状态
- 自动化告警

### 2. 故障排查
- 快速定位问题根源
- 分层诊断网络问题
- 提供解决建议

### 3. 部署验证
- 部署后验证服务状态
- 确认配置正确性
- 自动化测试

## 🔄 扩展性

### 添加新检测
```go
// 1. 实现 Check 接口
type NewCheck struct {
    BaseCheck
}

func NewNewCheck() Check {
    return &NewCheck{
        BaseCheck: BaseCheck{Name: "new_check"},
    }
}

func (n *NewCheck) Run(input types.Input) types.Result {
    // 实现检测逻辑
    return types.Result{
        Name:    "new_check",
        Success: true,
        Message: "检测结果",
    }
}

// 2. 注册到场景
engine.RegisterCheck(checks.NewNewCheck())
```

### 添加新规则
```go
// 在场景中添加规则
{
    Condition:  "new_check_ok && other_check_ok",
    Conclusion: "综合诊断结论",
    Suggestion: "操作建议",
    Priority:   2,
}
```

## 📞 技术支持

如有问题或建议，请查看：
- 项目文档：README.md
- 开发文档：DEV.md
- 功能文档：FEATURES.md