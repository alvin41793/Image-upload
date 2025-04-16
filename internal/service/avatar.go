package service

import (
	"github.com/alvin41793/Image-upload/internal/storage"
	"github.com/gin-gonic/gin"
	"mime/multipart"
)

type AvatarService struct {
	store storage.Storage
}

func NewAvatarService(store storage.Storage) *AvatarService {
	return &AvatarService{store: store}
}

func (s *AvatarService) UploadAvatar(fileHeader *multipart.FileHeader, c *gin.Context) (string, error) {
	return s.store.Upload(fileHeader, c)
}
