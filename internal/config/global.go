// config/global.go

package config

var global AppConfig

func InitGlobal(cfg AppConfig) {
	global = cfg
}

func G() AppConfig {
	return global
}
