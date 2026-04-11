#!/bin/bash

# 构建脚本

echo "=== 构建 OpsFlow ==="

# 检查 Go 是否安装
if ! command -v go &> /dev/null; then
    echo "错误: Go 未安装"
    exit 1
fi

# 检查项目目录
if [ ! -f "go.mod" ]; then
    echo "错误: 请在项目根目录运行此脚本"
    exit 1
fi

# 下载依赖
echo "下载依赖..."
go mod tidy

# 构建
echo "构建中..."
go build -o ops main.go

if [ $? -eq 0 ]; then
    echo "✅ 构建成功!"
    echo "可执行文件: ./ops"
    echo ""
    echo "使用示例:"
    echo "  ./ops run web example.com"
    echo "  ./ops list"
else
    echo "❌ 构建失败"
    exit 1
fi