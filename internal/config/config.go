package config

import (
	"github.com/spf13/viper"
)

type OSSConfig struct {
	Enable    bool   `mapstructure:"enable"`
	Endpoint  string `mapstructure:"endpoint"`
	AccessKey string `mapstructure:"access_key"`
	Secret    string `mapstructure:"secret"`
	Bucket    string `mapstructure:"bucket"`
	Domain    string `mapstructure:"domain"`
	Dir       string `mapstructure:"dir"`
}

type LogConfig struct {
	Dir      string `mapstructure:"dir"`
	KeepDays int    `mapstructure:"keep_days"`
	Level    string `mapstructure:"level"`
}

type LimiterConfig struct {
	Rate  float64 `mapstructure:"rate"`
	Burst int     `mapstructure:"burst"`
}

type UploadConfig struct {
	Dir          string   `mapstructure:"dir"`
	MaxSizeMB    int64    `mapstructure:"max_size_mb"`
	AllowedTypes []string `mapstructure:"allowed_types"`
}

type ServerConfig struct {
	Port string `mapstructure:"port"`
}

type Config struct {
	Srv        ServerConfig  `mapstructure:"server"`
	OSSVal     OSSConfig     `mapstructure:"oss"`
	LogVal     LogConfig     `mapstructure:"log"`
	LimiterVal LimiterConfig `mapstructure:"limiter"`
	UploadVal  UploadConfig  `mapstructure:"upload"`
}

func Load(path string) error {
	v := viper.New()
	v.SetConfigFile(path)

	if err := v.ReadInConfig(); err != nil {
		return err
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return err
	}

	InitGlobal(&cfg) // ✅ 注册为实现了 AppConfig 的接口实例
	return nil
}

func (c *Config) Server() ServerConfig   { return c.Srv }
func (c *Config) OSS() OSSConfig         { return c.OSSVal }
func (c *Config) Log() LogConfig         { return c.LogVal }
func (c *Config) Limiter() LimiterConfig { return c.LimiterVal }
func (c *Config) Upload() UploadConfig   { return c.UploadVal }
