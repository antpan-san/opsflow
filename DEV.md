# OpsFlow 开发文档

## 🎯 开发目标

构建一个场景驱动的运维诊断与自动化执行引擎。

## 📅 开发时间表

### Phase 1: MVP（已完成 ✅）
- ✅ Go 项目结构搭建
- ✅ Cobra CLI 框架
- ✅ 核心类型定义
- ✅ 引擎框架实现
- ✅ Web诊断场景
- ✅ 基础测试验证

**时间**：1天
**状态**：已完成，可运行

### Phase 2: 规则引擎优化（待开发）
- [ ] 完善条件表达式解析
- [ ] 规则优先级优化
- [ ] 规则热加载
- [ ] 结构化输出优化

**时间**：2-3天
**依赖**：Phase 1

### Phase 3: 插件系统（待开发）
- [ ] 插件接口定义
- [ ] 动态加载机制
- [ ] 插件管理命令
- [ ] 沙箱安全隔离

**时间**：3-4天
**依赖**：Phase 2

### Phase 4: 模板系统（待开发）
- [ ] 模板引擎
- [ ] 参数替换
- [ ] 文件生成
- [ ] 模板库

**时间**：2-3天
**依赖**：Phase 3

### Phase 5: AI增强（待开发）
- [ ] AI接口设计
- [ ] OpenAI集成
- [ ] 智能建议
- [ ] 日志分析

**时间**：2-3天
**依赖**：Phase 4

### Phase 6: CI/CD集成（待开发）
- [ ] GitHub Actions支持
- [ ] GitOps集成
- [ ] 生产就绪

**时间**：2-3天
**依赖**：Phase 5

## 🏗️ 当前架构

### 执行流程
```
用户输入 → CLI解析 → 引擎调度 → 检测执行 → 规则判断 → 结果输出
```

### 核心组件
- **Engine**: 引擎核心，协调各组件
- **Check**: 检测接口，具体检测逻辑
- **Rule**: 规则引擎，诊断逻辑
- **Scenario**: 场景定义，流程编排

## 🔧 技术实现

### 1. 类型系统
```go
// 输入参数
type Input struct {
    Target string
    Params map[string]string
}

// 检测结果
type Result struct {
    Name    string
    Success bool
    Message string
    Data    map[string]interface{}
}

// 诊断规则
type Rule struct {
    Condition  string
    Conclusion string
    Suggestion string
    Priority   int
}
```

### 2. 检测实现
- DNS检测：`net.LookupHost`
- TCP检测：`net.DialTimeout`
- HTTP检测：`http.Client.Get`

### 3. 规则引擎
- 条件表达式解析
- 逻辑运算符支持（&&、||）
- 优先级排序

## 📋 开发规范

### 代码风格
- 遵循 Go 语言规范
- 使用有意义的变量名
- 添加必要的注释
- 保持函数简洁

### 测试要求
- 每个检测必须有测试
- 核心逻辑必须有单元测试
- 集成测试验证完整流程

### 提交规范
- 功能开发：`feat: add web scenario`
- Bug修复：`fix: resolve DNS check issue`
- 文档更新：`docs: update README`

## 🎯 下一步开发

### 立即开始
1. **完善规则引擎**
   - 支持更复杂的条件表达式
   - 添加规则热加载功能

2. **扩展检测类型**
   - SSL证书检测
   - 性能指标检测
   - 安全检测

3. **优化输出格式**
   - JSON格式输出
   - 彩色终端输出
   - 详细报告生成

### 开发命令
```bash
# 构建项目
go build -o ops main.go

# 运行测试
go test ./...

# 运行诊断
./ops run web example.com
```

## 🤝 贡献指南

1. Fork 项目
2. 创建功能分支
3. 提交代码
4. 创建 Pull Request

## 📞 联系方式

- GitHub: https://github.com/yourusername/opsflow