#!/bin/bash
echo "=== OpsFlow 部署脚本 ==="

if [ "$EUID" -ne 0 ]; then
    echo "请使用 root 权限运行此脚本"
    exit 1
fi

# 备份旧版本
if [ -f "/root/ops" ]; then
    echo "备份旧版本..."
    mv /root/ops /root/ops.backup.$(date +%Y%m%d_%H%M%S)
fi

# 部署新版本
echo "部署新版本..."
cp ops /root/ops
chmod +x /root/ops
ln -sf /root/ops /usr/local/bin/ops 2>/dev/null || true

# 验证部署
echo "验证部署..."
/root/ops list

echo "✅ 部署完成!"
