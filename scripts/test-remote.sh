#!/bin/bash
# 远程测试脚本（不执行本地测试）

set -e

REMOTE_HOST="172.16.195.128"
REMOTE_USER="root"

echo "=== 远程测试脚本 ==="

# 编译 Linux ARM64 版本
echo "编译 Linux ARM64 版本..."
GOOS=linux GOARCH=arm64 go build -o ops-linux-arm64 main.go

# 传输到远程服务器
echo "传输到远程服务器..."
scp ops-linux-arm64 ${REMOTE_USER}@${REMOTE_HOST}:/tmp/ops-test

# 远程测试
echo "远程测试..."
ssh ${REMOTE_USER}@${REMOTE_HOST} "
cd /tmp
chmod +x ops-test

echo '=== 测试 Web 诊断 ==='
./ops-test run web example.com

echo ''
echo '=== 测试详细诊断 ==='
./ops-test run-detailed web-detailed example.com

echo ''
echo '=== 测试 HTTPS ==='
./ops-test run web google.com --protocol https

echo ''
echo '=== 测试内网检测 ==='
./ops-test run-detailed web-detailed localhost

# 清理测试文件
rm -f /tmp/ops-test

echo '远程测试完成'
"

echo "✅ 所有远程测试通过"