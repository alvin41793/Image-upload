package routers

import (
	"github.com/alvin41793/Image-upload/internal/handler"
	"github.com/gin-gonic/gin"
)
import "github.com/alvin41793/Image-upload/internal/middleware"

func AvatarRouterInit(engine *gin.RouterGroup, avatarHandler *handler.AvatarHandler) {
	r := engine.Group("/avatar")
	r.Use(middleware.RateLimitMiddleware(), middleware.ImageUploadValidator())
	{
		r.POST("/upload", avatarHandler.Upload)
	}
}
