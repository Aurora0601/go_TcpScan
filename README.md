# go_TcpScan

一个基于Go语言开发的高性能TCP端口扫描工具，支持并发扫描和代理池功能，帮助安全测试人员快速发现目标主机开放的端口。

## 功能特点

- 高性能并发端口扫描
- 支持代理池功能，提高扫描匿名性
- 轻量级设计，仅依赖Go官方库
- 简单易用的命令行接口
- 跨平台支持（可编译为Windows、Linux、macOS等系统的可执行文件）

## 安装使用

### 直接使用编译好的可执行文件

从项目Release页面下载对应系统的可执行文件，直接运行即可。

### 源码编译

```bash
# 克隆仓库
git clone https://github.com/Aurora0601/go_TcpScan.git

# 进入项目目录
cd go_TcpScan

# 编译
go build -o tcpscan
```

## 使用示例

基本扫描：
```bash
# 扫描单个主机的指定端口
tcpscan -u example.com -p 80,443,8080

# 扫描单个主机的端口范围
tcpscan -u example.com -p 1-1000

# 扫描多个主机
tcpscan -u 192.168.1.1,192.168.1.2 -p 1-100
```

使用代理池扫描：
```bash
tcpscan -u example.com -p 1-1000 -proxy proxy.txt
```

## 命令参数说明

| 参数 | 说明 | 示例 |
|------|------|------|
| -u | 目标主机/IP地址（支持多个，用逗号分隔） | -u example.com,192.168.1.1 |
| -p | 端口（支持单个、多个或范围，用逗号分隔） | -p 80 或 -p 80,443 或 -p 1-1000 |
| -proxy | 代理池文件路径 | -proxy proxies.txt |
| -t | 并发数（默认100） | -t 200 |
| -timeout | 超时时间（毫秒，默认5000） | -timeout 3000 |
| -h | 显示帮助信息 | -h |

## 项目结构

```
go_TcpScan/
├── cmd/
│   └── cli/              # 命令行接口模块
├── internal/
│   ├── scanner/          # 核心扫描模块
│   └── proxy/            # 代理池模块
├── pkg/
│   └── utils/            # 工具函数
└── main.go               # 程序入口
```

## 贡献指南

欢迎任何形式的贡献！如果你发现bug或有新功能建议，请在GitHub上提交issue。

如果你想提交代码贡献：
1. Fork本仓库
2. 创建你的特性分支 (`git checkout -b feature/amazing-feature`)
3. 提交你的修改 (`git commit -m 'Add some amazing feature'`)
4. 推送到分支 (`git push origin feature/amazing-feature`)
5. 打开Pull Request

## 许可证

本项目采用Apache License 2.0许可证 - 详见[LICENSE](LICENSE)文件。

## 致谢

感谢Zxc对本项目的支持与贡献。

## 作者

AuroraSEC
