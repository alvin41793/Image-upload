package handler

import (
	"github.com/alvin41793/Image-upload/internal/logger"
	"github.com/alvin41793/Image-upload/internal/service"
	"github.com/alvin41793/Image-upload/internal/util"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"mime/multipart"
)

type AvatarHandler struct {
	svc *service.AvatarService
}

func NewAvatarHandler(svc *service.AvatarService) *AvatarHandler {
	return &AvatarHandler{svc: svc}
}

func (h *AvatarHandler) Upload(c *gin.Context) {
	fileHeader, _ := c.Get("fileHeader")
	fh := fileHeader.(*multipart.FileHeader)

	// 保存文件
	url, err := h.svc.UploadAvatar(fh, c)
	if err != nil {
		util.Fail(c, 500, "保存失败")
		return
	}
	logger.L().Info("头像上传成功", zap.String("url", url))
	util.Success(c, gin.H{
		"message": "上传成功",
		"url":     url,
	})
}
