# OpsFlow 部署文档

## 📋 部署信息

**目标服务器**: root@172.16.195.128  
**部署目录**: /root/ops  
**架构**: Linux ARM64  
**部署时间**: 2026-04-11 17:58

## ✅ 部署验证

### 远程测试结果

```bash
# 列出可用场景
ssh root@172.16.195.128 "/root/ops list"

# Web 诊断测试
ssh root@172.16.195.128 "/root/ops run web example.com"
```

### 测试输出

```
=== 诊断报告 ===
场景: web
目标: example.com

[检测结果]
✅ tcp: TCP连接成功
✅ http: HTTP请求成功
✅ dns: DNS解析成功

[诊断结论]
结论: Web服务正常
建议: 无需操作
```

## 🚀 使用方法

### 基本命令

```bash
# 远程执行诊断
ssh root@172.16.195.128 "/root/ops run web example.com"

# 或者登录服务器后执行
ssh root@172.16.195.128
/root/ops run web example.com
```

### 常用场景

```bash
# Web 诊断
/root/ops run web example.com
/root/ops run web example.com --port 443 --protocol https

# 列出场景
/root/ops list
```

## 📁 部署文件结构

```
/root/
├── ops                          # 主程序
├── ops.backup.*                 # 备份文件
├── opsflow-deploy.tar.gz       # 部署包
└── deploy/                      # 部署目录
    ├── ops                      # 主程序
    ├── deploy.sh               # 部署脚本
    └── opsflow.service         # systemd 服务文件
```

## 🔧 系统服务

### systemd 服务

```bash
# 查看服务状态
systemctl status opsflow

# 启动服务
systemctl start opsflow

# 停止服务
systemctl stop opsflow

# 重启服务
systemctl restart opsflow

# 查看日志
journalctl -u opsflow -f
```

## 🔄 更新部署

### 更新步骤

1. **本地编译新版本**
   ```bash
   cd /tmp/opsflow
   GOOS=linux GOARCH=arm64 go build -o ops-linux-arm64 main.go
   ```

2. **重新打包**
   ```bash
   rm -rf deploy
   mkdir -p deploy
   cp ops-linux-arm64 deploy/ops
   # 复制其他文件...
   tar -czf opsflow-deploy.tar.gz deploy/
   ```

3. **传输并部署**
   ```bash
   scp opsflow-deploy.tar.gz root@172.16.195.128:/root/
   ssh root@172.16.195.128 "cd /root && tar -xzf opsflow-deploy.tar.gz && cd deploy && bash deploy.sh"
   ```

## 📊 监控和维护

### 日志查看

```bash
# 查看程序输出
ssh root@172.16.195.128 "/root/ops run web example.com"

# 查看系统日志
ssh root@172.16.195.128 "journalctl -u opsflow -f"
```

### 备份管理

```bash
# 查看备份文件
ssh root@172.16.195.128 "ls -la /root/ops.backup.*"

# 恢复备份
ssh root@172.16.195.128 "mv /root/ops.backup.20260411_120000 /root/ops && chmod +x /root/ops"
```

## 🔍 故障排除

### 常见问题

1. **无法执行二进制文件**
   - 检查架构是否匹配：`uname -m`
   - 重新编译对应架构版本

2. **权限问题**
   - 确保执行权限：`chmod +x /root/ops`
   - 使用 root 权限执行

3. **网络连接问题**
   - 检查网络连接：`ping example.com`
   - 检查防火墙规则

## 📞 支持

如有问题，请检查：
1. 程序版本是否最新
2. 网络连接是否正常
3. 权限配置是否正确

---

**部署完成时间**: 2026-04-11 17:58  
**部署人员**: OpenClaw Agent  
**架构**: Linux ARM64