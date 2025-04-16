package storage

import (
	"github.com/gin-gonic/gin"
	"mime/multipart"
)

type Storage interface {
	Upload(fileHeader *multipart.FileHeader, c *gin.Context) (string, error)
}
