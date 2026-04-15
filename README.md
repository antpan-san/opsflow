# OpsFlow - 运维诊断与自动化执行引擎

> ❌ 把运维经验变成一条可执行命令 (Turn ops experience into executable commands)

## 🎯 产品定位

OpsFlow 是一个**场景驱动的运维诊断与自动化执行引擎**，支持：

- ✅ 故障诊断（web / k8s / CI/CD）
- ✅ 自动建议（rule-based）
- ✅ 自动生成（YAML / CI）
- ✅ 可扩展能力（plugin system）
- ✅ AI 增强（解释 + 生成）

## 🚀 快速开始

### 安装

```bash
# 克隆项目
git clone https://github.com/antpan-san/opsflow.git
cd opsflow

# 构建
go build -o ops main.go
```

### 使用示例

```bash
# Web 诊断
./ops run web example.com

# 指定端口和协议
./ops run web example.com --port 443 --protocol https

# 列出可用场景
./ops list

# 运行场景文件
./ops run -f scenarios/check_host.yaml -v
```

## 📋 执行模型

```text
Input → Checks → Results → Rules → Conclusion → Actions
```

## 🏗️ 项目结构

```bash
opsflow/
├── main.go                 # 程序入口
├── cmd/                    # CLI 命令
│   ├── root.go            # 根命令
│   ├── run.go             # 运行诊断
│   └── list.go            # 列出场景
├── internal/
│   ├── engine/            # 引擎核心
│   ├── types/             # 类型定义
│   ├── checks/            # 检测实现
│   ├── rules/             # 规则引擎
│   └── scenarios/         # 场景定义
├── go.mod                 # Go 模块定义
└── build.sh               # 构建脚本
```

## 🔧 核心组件

### 场景 (Scenario)

定义诊断流程：检测哪些项目、应用什么规则、执行什么动作。

### 检测 (Check)

具体的检测逻辑，如 DNS 解析、TCP 连接、HTTP 请求。

### 规则 (Rule)

基于检测结果的诊断逻辑，生成结论和建议。

### 动作 (Action)

根据诊断结果执行的具体操作。

## 📦 已支持场景

### Web 诊断

检测域名的 DNS 解析、TCP 连接、HTTP 服务状态。

```bash
ops run web example.com
```

**检测项目：**
- ✅ DNS 解析
- ✅ TCP 连接
- ✅ HTTP 请求

**诊断规则：**
- DNS + TCP + HTTP 正常 → 服务正常
- DNS + TCP 正常 + HTTP 异常 → 后端服务异常
- DNS + TCP 异常 → 网络连接异常
- DNS 异常 → 域名配置问题

## 🤖 AI 增强

| 能力 | 是否使用AI |
|------|-----------|
| 规则判断 | ❌ |
| 结果解释 | ✅ |
| YAML生成 | ✅ |
| 日志分析 | ✅ |

## 📝 开发计划

- [x] Phase 1: CLI + Engine + Web 诊断
- [ ] Phase 2: 规则引擎优化
- [ ] Phase 3: 插件系统
- [ ] Phase 4: 模板系统
- [ ] Phase 5: AI 增强
- [ ] Phase 6: CI/CD 集成

## 🧪 测试

```bash
# 运行诊断
./ops run web example.com

# 预期输出
=== 诊断报告 ===
场景: web
目标: example.com

[检测结果]
✅ dns: DNS解析成功
✅ tcp: TCP连接成功
✅ http: HTTP请求成功

[诊断结论]
结论: Web服务正常
建议: 无需操作
```

## 📖 API 文档

### 类型定义

```go
type Input struct {
    Target string
    Params map[string]string
}

type Result struct {
    Name    string
    Success bool
    Message string
    Data    map[string]interface{}
}

type Rule struct {
    Condition  string
    Conclusion string
    Suggestion string
    Priority   int
}
```

### 接口定义

```go
type Check interface {
    Name() string
    Run(input Input) Result
}

type Action interface {
    Name() string
    Execute(ctx Context) error
}
```

## 🛠️ 开发指南

### 添加新检测

1. 在 `internal/checks/` 创建检测文件
2. 实现 `Check` 接口
3. 在场景中注册检测

### 添加新场景

1. 在 `internal/scenarios/` 创建场景文件
2. 定义检测列表和规则
3. 在引擎中注册场景

## 📄 许可证

MIT License

## 🤝 贡献

欢迎提交 Issue 和 Pull Request！

## 📞 联系方式

- GitHub: https://github.com/yourusername/opsflow# Test update
# test change
# 测试 Git 钩子
# 再次测试 Git 钩子
# 第三次测试 Git 钩子
# 第四次测试 Git 钩子
