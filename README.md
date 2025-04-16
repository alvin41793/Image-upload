# Image Upload Service

一个基于 Golang 的图片上传服务，支持本地存储与阿里云 OSS，具备限流、日志、配置模块、结构化响应等功能。

## ✨ 功能特性

- 支持图片上传接口（支持格式校验）
- 本地存储 & 阿里云 OSS 支持
- 用户请求频率限制（基于 IP 限流）
- 统一结构化响应体
- 配置文件支持多环境
- 模块化架构，便于扩展与维护
- 统一日志记录，支持结构化输出

## 📦 技术栈

- Go 1.22+
- Gin Web Framework
- 阿里云 OSS SDK
- Uber Zap 日志库
- 单元测试 & 中间件支持

## 🛠️ 项目结构



## 🚀 快速开始

### 1. 克隆项目

```bash
git clone https://github.com/alvin41793/Image-upload.git
cd Image-upload


2 配置文件
修改 config.default.yaml 来配置服务端口、本地路径、阿里云 OSS 等：

server:
  port: ":8080" # 服务端口
oss:
  enable: false # 是否启用 OSS
  endpoint: ""
  access_key: ""
  secret: ""
  bucket: ""
  domain: ""
  dir: ""
log:
  dir: "./logs"  # 日志文件默认保存路径
  keep_days: 7    # 保留最近多少天的日志
  level: "info"  # 日志级别

limiter:
  rate: 1     # 每秒请求数
  burst: 3    # 最多积攒3个

upload:
  dir: "./uploads"
  max_size_mb: 5     # 限制上传文件大小（MB）
  allowed_types: # 允许上传的文件类型
    - .jpg
    - .png

3. 运行项目
go run cmd/main.go -config ./config.default.yaml

默认访问接口：

# 头像上传接口

## 接口信息
- **接口地址**：`http://localhost:8080/api/v1/avatar/upload`
- **请求方法**：`POST`

## 请求参数
| 参数名 | 参数类型 | 是否必填 | 描述 |
| ---- | ---- | ---- | ---- |
| avatar | 文件 | 是 | 要上传的头像文件，支持的文件格式为图片（如 JPEG、PNG 等） |

## 请求示例（使用 curl）
```bash
curl -X POST http://localhost:8080/api/v1/avatar/upload \
  -F "avatar=@/path/to/your/image.jpg"



🧹 后续规划
✅ 添加用户身份验证

✅ 添加断点续传 & 进度反馈

✅ 丰富测试用例

✅ 引入 Swagger 文档生成

✅ 支持多存储后端（七牛云、S3 等）





