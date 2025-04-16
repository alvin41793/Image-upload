package config

type AppConfig interface {
	Server() ServerConfig
	OSS() OSSConfig
	Log() LogConfig
	Limiter() LimiterConfig
	Upload() UploadConfig
}
