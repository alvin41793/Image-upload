package middleware

import (
	"github.com/alvin41793/Image-upload/internal/config"
	"github.com/alvin41793/Image-upload/internal/logger"
	"github.com/alvin41793/Image-upload/internal/util"
	"go.uber.org/zap"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

// ImageUploadValidator 中间件：验证上传的图片是否合法 图片大小限制
func ImageUploadValidator() gin.HandlerFunc {
	return func(c *gin.Context) {
		cfg := config.G().Upload()
		// 限制请求体大小
		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, cfg.MaxSizeMB<<20)

		fileHeader, err := c.FormFile("avatar")
		if err != nil {
			logger.L().Error("fail to read file", zap.Error(err))
			util.Fail(c, http.StatusBadRequest, "fail to read file")
			c.Abort()
			return
		}

		// 检查扩展名
		ext := strings.ToLower(filepath.Ext(fileHeader.Filename))
		if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
			logger.L().Error("The file format is not supported", zap.Error(err))
			util.Fail(c, http.StatusBadRequest, "Only supports JPG/PNG image formats")
			c.Abort()
			return
		}

		file, err := fileHeader.Open()
		if err != nil {
			logger.L().Error("fail to open file", zap.Error(err))
			util.Fail(c, http.StatusInternalServerError, "fail to open file")
			c.Abort()
			return
		}
		defer file.Close()

		// 检查是否是真图片（防止伪造内容）
		if !isValidImage(file) {
			logger.L().Error("The file is not a valid image", zap.Error(err), zap.String("ip", c.ClientIP()))
			util.Fail(c, http.StatusBadRequest, "The file is not a valid image")
			c.Abort()
			return
		}

		// 将验证通过的内容传递给后续处理器
		c.Set("fileHeader", fileHeader)
		/*	c.Set("fileExt", ext)*/

		c.Next()
	}
}

func isValidImage(file multipart.File) bool {
	_, _, err := image.Decode(file)
	return err == nil
}
