# Image Upload Service

一个基于 golong + gin 的图片上传服务，支持本地｜阿里云 OSS存储，扩展能力强



## ✨ 功能特性

- 图片上传api（支持图片格式校验、支持跨域请求、基于ip的限流访问）
- 存储方式动态配置 ，目前支持本地 & 阿里云 OSS（可扩展七牛云、S3）
- 统一结构化响应体（可优化）
- 模块化架构，便于扩展与维护
- 统一日志记录，支持结构化输出（可以升级成支持分布式）


## 关于go环境

要构建 Lotus，您需要安装[Go 1.22.7 或更高版本](https://golang.org/dl/)：

```
wget -c https://golang.org/dl/go1.22.7.linux-amd64.tar.gz -O - | sudo tar -xz -C /usr/local
```

**提示：** 你需要将其添加`/usr/local/go/bin`到你的路径中。对于大多数 Linux 发行版，你可以运行以下命令：

```
echo "export PATH=$PATH:/usr/local/go/bin" >> ~/.bashrc && source ~/.bashrc
```

如果遇到困难，请参阅[官方 Golang 安装说明。](https://golang.org/doc/install)


## 🛠️ 项目结构

```
├── app/
│   ├── app.go  // 配置文件、日志初始化等
│   └── main.go  // 程序主入口，启动应用
├── internal/
│   ├── config/  // 配置读取、解析与模型定义
│   ├── handler/  // 处理HTTP请求，接收请求并返回响应
│   ├── logger/  // 日志模块，负责全局日志记录与切割
│   ├── middleware/  // Gin相关中间件，如流量控制、日志、跨域处理
│   ├── routers/  // 路由注册，映射URL到请求处理函数
│   ├── service/  // 实现核心业务逻辑
│   ├── storage/  // 存储操作，包括本地和OSS接口定义与实现
│   └── util/  // 提供工具类功能，如返回值整理等
├── config.default.yaml  // 项目默认配置文件
├── go.mod  // Go语言模块管理，管理项目依赖
└── README.md  // 项目说明文档，介绍功能、使用和安装等
```


## 🚀 快速开始

### 1. 克隆项目

```bash
git clone https://github.com/alvin41793/Image-upload.git
cd Image-upload
```

### 2. 配置文件

修改 `config.default.yaml` 来配置服务端口、本地路径、阿里云 OSS 等：

```yaml
server:
  port: ":8080"          # 服务端口

oss:
  enable: false          # 是否启用 OSS
  endpoint: ""           # OSS endpoint
  access_key: ""         # OSS access key
  secret: ""             # OSS secret key
  bucket: ""             # OSS bucket 名称
  domain: ""             # OSS 域名
  dir: ""                # 上传目录

log:
  dir: "./logs"          # 日志文件默认保存路径
  keep_days: 7           # 保留最近多少天的日志
  level: "info"          # 日志级别

limiter:
  rate: 1                # 每秒请求数
  burst: 3               # 最多积攒3个请求

upload:
  dir: "./uploads"       # 上传文件存储路径
  max_size_mb: 5         # 限制上传文件大小（MB）
  allowed_types:        # 允许上传的文件类型
    - .jpg
    - .png

```

### 3. 运行项目

运行以下命令启动服务：

```bash
go run cmd/main.go --config ./config.default.yaml
```


### 4. 头像上传接口

#### 接口信息

- **接口地址**：`http://localhost:8080/api/v1/avatar/upload`
- **请求方法**：`POST`

#### 请求参数

| 参数名 | 参数类型 | 是否必填 | 描述                                                      |
| ------ | -------- | -------- | --------------------------------------------------------- |
| avatar | 文件     | 是       | 要上传的头像文件，支持的文件格式为图片（如 JPEG、PNG 等） |

#### 请求示例（使用 curl）

```bash
curl -X POST http://localhost:8080/api/v1/avatar/upload \
  -F "avatar=@/path/to/your/image.jpg"
```


## 🧹 后续规划

- ✅ 添加用户身份验证
- ✅ 添加断点续传 & 进度反馈
- ✅ 丰富测试用例
- ✅ 引入 Swagger 文档生成
- ✅ 支持多存储后端（七牛云、S3 等）
- ✅ 微服务架构改造（架构拆分、接口设计、服务治理、安全、性能和容器化部署）
- ...
