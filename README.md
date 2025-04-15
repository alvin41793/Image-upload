# Image-upload
图片上传模块文件结构
Image-upload/
├── cmd/main.go
├── config.yaml
├── internal/
│   ├── config/config.go
│   ├── handler/avatar.go
│   ├── service/avatar.go
│   ├── storage/storage.go
│   ├── storage/local.go
│   ├── storage/oss.go
│   ├── middleware/rate_limit.go
│   └── util/{config.go, logger.go, response.go, validator.go}
└── go.mod
