package main

import (
	"fmt"
	"github.com/alvin41793/Image-upload/internal/config"
	"github.com/alvin41793/Image-upload/internal/handler"
	"github.com/alvin41793/Image-upload/internal/logger"
	"github.com/alvin41793/Image-upload/internal/middleware"
	"github.com/alvin41793/Image-upload/internal/routers"
	"github.com/alvin41793/Image-upload/internal/service"
	"github.com/alvin41793/Image-upload/internal/storage"
	"github.com/gin-gonic/gin"
)

type App struct {
	Config  config.AppConfig
	Logger  logger.Logger
	Storage storage.Storage
	Engine  *gin.Engine
}

func NewApp(configPath string) (*App, error) {
	// 加载配置

	if err := config.Load(configPath); err != nil {
		fmt.Println("failed to load config file")
		return nil, err
	}

	// 替换原来的 logger.InitGlobal
	if err := logger.Init(); err != nil {
		fmt.Println("failed to init logger")
		return nil, err
	}

	cfg := config.G().OSS()
	// 初始化存储
	var store storage.Storage
	var err error
	if cfg.Enable {
		store, err = storage.NewAliyunOSS()
		if err != nil {
			logger.L().Error("oss init failed")
			return nil, err
		}
		logger.L().Info("oss storage init success")
	} else {
		store = storage.NewLocalStorage()
		logger.L().Info("local storage init success")
	}

	// 初始化服务/路由
	avatarSvc := service.NewAvatarService(store)
	avatarHandler := handler.NewAvatarHandler(avatarSvc)

	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	r.Static("/uploads", "./uploads")

	apiV1 := r.Group("/api/v1")
	routers.AvatarRouterInit(apiV1, avatarHandler)

	return &App{
		Config:  config.G(),
		Logger:  logger.L(),
		Storage: store,
		Engine:  r,
	}, nil
}

func (a *App) Run() error {
	port := config.G().Server().Port
	a.Logger.Info("server starting port " + port)
	return a.Engine.Run(port)
}
