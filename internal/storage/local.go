package storage

import (
	"fmt"
	"github.com/alvin41793/Image-upload/internal/config"
	"github.com/alvin41793/Image-upload/internal/logger"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

type LocalStorage struct {
	dir string
}

func NewLocalStorage() *LocalStorage {
	cfg := config.G().Upload()
	return &LocalStorage{dir: cfg.Dir}
}

func (s *LocalStorage) Upload(fileHeader *multipart.FileHeader, c *gin.Context) (string, error) {
	// 创建日期目录
	dateDir := time.Now().Format("2006-01-02")
	saveDir := filepath.Join(s.dir, dateDir)
	if err := os.MkdirAll(saveDir, os.ModePerm); err != nil {
		return "", fmt.Errorf("create dir failed: %w", err)
	}

	// 生成保存文件名
	ext := filepath.Ext(fileHeader.Filename)
	saveName := uuid.New().String() + ext
	savePath := filepath.Join(saveDir, saveName)

	// 保存上传文件
	if err := saveUploadedFile(fileHeader, savePath); err != nil {
		logger.L().Error("Failed to save file",
			zap.String("filename", fileHeader.Filename),
			zap.String("path", savePath),
			zap.Error(err),
		)
		return "", err
	}

	// 获取相对路径构建 URL
	relPath, err := filepath.Rel(s.dir, savePath)
	if err != nil {
		return "", fmt.Errorf("relative path error: %w", err)
	}

	// 用 filepath.Base(s.dir) + relPath 构建最终路径（如 uploads/2025-04-16/xxx.jpg）
	urlPath := filepath.ToSlash(filepath.Join(filepath.Base(s.dir), relPath))
	return buildUrlPath(c, urlPath), nil
}

func saveUploadedFile(fileHeader *multipart.FileHeader, path string) error {
	src, err := fileHeader.Open()
	if err != nil {
		return fmt.Errorf("open file failed: %w", err)
	}
	defer src.Close()

	dst, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("create file failed: %w", err)
	}
	defer dst.Close()

	if _, err := dst.ReadFrom(src); err != nil {
		return fmt.Errorf("write file failed: %w", err)
	}
	return nil
}

// 构建文件路径
func buildUrlPath(c *gin.Context, path string) string {
	// 获取请求的协议
	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	}
	// 尝试从 X-Forwarded-Host 头获取域名或 IP
	host := c.Request.Header.Get("X-Forwarded-Host")
	if host == "" {
		// 如果 X-Forwarded-Host 头不存在，使用 Host 头
		host = c.Request.Host
	}
	// 返回可访问的图片 URL
	return fmt.Sprintf("%s://%s/%s", scheme, host, path)
}
