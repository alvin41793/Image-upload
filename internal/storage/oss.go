package storage

import (
	"fmt"
	"github.com/alvin41793/Image-upload/internal/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"mime/multipart"
	"path/filepath"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/alvin41793/Image-upload/internal/config"
)

type AliyunOSS struct {
	bucket *oss.Bucket
	domain string
	dir    string
}

func NewAliyunOSS() (*AliyunOSS, error) {
	logger.L().Info("init OSS")
	cfg := config.G().OSS()
	client, err := oss.New(cfg.Endpoint, cfg.AccessKey, cfg.Secret)
	if err != nil {
		logger.L().Error("failed to init OSS", zap.Error(err))
		return nil, err
	}
	bucket, err := client.Bucket(cfg.Bucket)
	if err != nil {
		logger.L().Error("Get OSS bucket failed", zap.Error(err))
		return nil, err
	}
	logger.L().Info("init OSS success")
	return &AliyunOSS{
		bucket: bucket,
		domain: cfg.Domain,
		dir:    cfg.Dir,
	}, nil
}

func (a *AliyunOSS) Upload(fileHeader *multipart.FileHeader, _ *gin.Context) (string, error) {
	filename := fmt.Sprintf("%d_%s", time.Now().UnixNano(), filepath.Base(fileHeader.Filename))

	file, err := fileHeader.Open()
	if err != nil {
		logger.L().Error("Failed to open file", zap.Error(err))
		return "", err
	}
	defer file.Close()

	err = a.bucket.PutObject(filename, file)
	if err != nil {
		logger.L().Error("Failed to upload file to OSS", zap.Error(err))
		return "", err
	}
	return fmt.Sprintf("%s/%s", a.domain, filename), nil
}
