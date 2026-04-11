# OpsFlow 诊断场景全面测试命令

## 🎯 场景概览

### 1. Web 诊断场景 (web)
- DNS 解析检测
- TCP 端口连通性
- HTTP 协议检测

### 2. Web 详细诊断场景 (web-detailed)
- DNS 详细解析（CNAME、IPv4/IPv6）
- 内网网络识别
- 详细诊断报告

## 📋 完整测试命令

### 基础 Web 诊断

#### 1. 基本域名诊断
```bash
# 测试 HTTP 网站
./ops run web example.com

# 测试 HTTPS 网站
./ops run web google.com --protocol https

# 测试指定端口
./ops run web example.com --port 443 --protocol https
```

#### 2. 不同类型网站诊断
```bash
# 测试搜索引擎
./ops run web google.com
./ops run web baidu.com
./ops run web bing.com

# 测试社交媒体
./ops run web github.com
./ops run web twitter.com
./ops run web facebook.com

# 测试国内网站
./ops run web www.baidu.com
./ops run web www.qq.com
./ops run web www.taobao.com

# 测试企业网站
./ops run web www.example.com
./ops run web www.company.com
```

#### 3. 网络连通性测试
```bash
# 测试本地网络
./ops run web localhost

# 测试内网服务
./ops run web 192.168.1.1
./ops run web 10.0.0.1

# 测试端口连通性
./ops run web example.com --port 80
./ops run web example.com --port 443
./ops run web example.com --port 8080
```

### 详细诊断场景

#### 1. DNS 详细解析
```bash
# 详细 DNS 诊断
./ops run-detailed web-detailed example.com

# 测试不同域名
./ops run-detailed web-detailed google.com
./ops run-detailed web-detailed github.com
./ops run-detailed web-detailed baidu.com
```

#### 2. 网络类型识别
```bash
# 测试公网网站
./ops run-detailed web-detailed example.com

# 测试本地服务
./ops run-detailed web-detailed localhost

# 测试内网 IP
./ops run-detailed web-detailed 192.168.1.1
```

#### 3. 详细协议检测
```bash
# HTTP 详细诊断
./ops run-detailed web-detailed example.com --protocol http

# HTTPS 详细诊断
./ops run-detailed web-detailed example.com --protocol https

# 指定端口详细诊断
./ops run-detailed web-detailed example.com --port 443 --protocol https
```

## 🔧 场景详解

### 场景 1: Web 诊断 (web)

**功能**：
- DNS 解析检测
- TCP 端口连通性
- HTTP 协议检测

**命令格式**：
```bash
./ops run web <目标> [flags]
```

**参数**：
- `--port` : 指定端口（默认 80）
- `--protocol` : 指定协议（http/https）

**输出示例**：
```
=== Web 诊断报告 ===
场景: web
目标: example.com

[DNS 解析]
✅ DNS解析成功: 93.184.216.34

[TCP 连接]
✅ TCP连接成功 (端口: 80)

[HTTP 协议]
✅ HTTP协议正常 (状态码: 200)

[诊断结论]
结论: Web服务完全正常
建议: 无需操作
```

### 场景 2: Web 详细诊断 (web-detailed)

**功能**：
- DNS 详细解析（CNAME、IPv4/IPv6）
- 内网网络识别
- 详细网络接口信息
- 详细诊断报告

**命令格式**：
```bash
./ops run-detailed web-detailed <目标> [flags]
```

**参数**：
- `--port` : 指定端口（默认 80）
- `--protocol` : 指定协议（http/https）

**输出示例**：
```
=== Web 详细诊断报告 ===
场景: web-detailed
目标: example.com
时间: 2026-04-11 19:29:00

[网络接口信息]
  接口: en0
    状态: up
    MTU: 1500
    地址: 192.168.1.100/24

[内网信息]
  网络类型: 家庭/办公网络
  内网IP: 192.168.1.100

[DNS 解析信息]
  域名: example.com
  CNAME: cdn.example.com
  IPv4: 93.184.216.34
  IPv6: 2606:2800:220:1:248:1893:25c8:1946
  记录数: 2

[TCP 连接信息]
  ✅ 状态: TCP连接成功
  端口: 80

[HTTP 协议信息]
  ✅ 状态: HTTP协议正常
  状态码: 200
  内容类型: text/html
  服务器: Apache
  URL: http://example.com

[诊断结论]
结论: Web服务完全正常
建议: 无需操作
```

## 🧪 完整测试套件

### 测试套件 1: 基础功能测试

```bash
#!/bin/bash
echo "=== OpsFlow 基础功能测试 ==="

# 测试场景列表
echo "1. 测试场景列表"
./ops list

# 测试基本诊断
echo ""
echo "2. 测试基本诊断"
./ops run web example.com

# 测试 HTTPS
echo ""
echo "3. 测试 HTTPS"
./ops run web google.com --protocol https

# 测试详细诊断
echo ""
echo "4. 测试详细诊断"
./ops run-detailed web-detailed example.com

echo ""
echo "=== 测试完成 ==="
```

### 测试套件 2: 不同网站测试

```bash
#!/bin/bash
echo "=== 不同网站诊断测试 ==="

# 定义测试网站列表
websites=(
    "example.com"
    "google.com"
    "github.com"
    "baidu.com"
    "www.qq.com"
)

for site in "${websites[@]}"; do
    echo ""
    echo "测试网站: $site"
    ./ops run web $site
    echo "---"
done

echo ""
echo "=== 所有网站测试完成 ==="
```

### 测试套件 3: 网络类型测试

```bash
#!/bin/bash
echo "=== 网络类型诊断测试 ==="

# 测试公网网站
echo "1. 公网网站测试"
./ops run-detailed web-detailed example.com

# 测试本地服务
echo ""
echo "2. 本地服务测试"
./ops run-detailed web-detailed localhost

# 测试内网 IP（根据实际网络调整）
echo ""
echo "3. 内网测试"
./ops run-detailed web-detailed 192.168.1.1 2>/dev/null || echo "内网设备不可达"

echo ""
echo "=== 网络类型测试完成 ==="
```

### 测试套件 4: 协议和端口测试

```bash
#!/bin/bash
echo "=== 协议和端口测试 ==="

# 测试不同协议
echo "1. HTTP 协议测试"
./ops run web example.com --protocol http

echo ""
echo "2. HTTPS 协议测试"
./ops run web example.com --protocol https

# 测试不同端口
echo ""
echo "3. 端口测试"
./ops run web example.com --port 80
./ops run web example.com --port 443 --protocol https

echo ""
echo "=== 协议和端口测试完成 ==="
```

### 测试套件 5: 详细诊断全面测试

```bash
#!/bin/bash
echo "=== 详细诊断全面测试 ==="

# 测试不同域名的详细诊断
domains=(
    "example.com"
    "google.com"
    "github.com"
    "baidu.com"
)

for domain in "${domains[@]}"; do
    echo ""
    echo "详细诊断: $domain"
    ./ops run-detailed web-detailed $domain
    echo "---"
done

echo ""
echo "=== 详细诊断测试完成 ==="
```

## 📊 预期结果

### 成功场景
- ✅ DNS 解析正常
- ✅ TCP 连接成功
- ✅ HTTP/HTTPS 协议正常
- ✅ 返回状态码 200

### 失败场景
- ❌ DNS 解析失败（域名不存在）
- ❌ TCP 连接失败（端口关闭）
- ❌ HTTP 协议错误（状态码 4xx/5xx）

## 🔍 故障诊断

### DNS 解析失败
```bash
# 检查 DNS 配置
nslookup example.com

# 使用详细诊断查看更多信息
./ops run-detailed web-detailed example.com
```

### TCP 连接失败
```bash
# 检查端口是否开放
telnet example.com 80

# 使用详细诊断查看网络信息
./ops run-detailed web-detailed example.com
```

### HTTP 协议错误
```bash
# 检查 HTTP 响应
curl -I http://example.com

# 使用详细诊断查看详细信息
./ops run-detailed web-detailed example.com
```

## 📝 测试报告模板

```markdown
# OpsFlow 诊断测试报告

**测试时间**: 2026-04-11 19:29:00  
**测试环境**: macOS / Linux

## 测试场景

### 1. 基础 Web 诊断
- [ ] example.com - HTTP
- [ ] google.com - HTTPS
- [ ] github.com - HTTPS

### 2. 详细诊断
- [ ] example.com - DNS 详细解析
- [ ] google.com - 网络类型识别
- [ ] localhost - 本地服务诊断

## 测试结果

### 成功率
- 基础诊断: 100%
- 详细诊断: 100%
- 协议测试: 100%

### 性能指标
- 诊断时间: < 1s
- DNS 解析: < 100ms
- TCP 连接: < 200ms
- HTTP 检测: < 500ms

## 结论

所有诊断场景测试通过，功能正常。
```

---

**准备就绪！现在可以使用上述命令全面测试 OpsFlow 的所有诊断场景。** 🚀