package main

import (
	"flag"
	"fmt"
)

func main() {
	// 解析配置文件路径参数
	configPath := flag.String("config", "./config.default.yaml", "Path to configuration file")
	flag.Parse()

	fmt.Println("Loading config from:", *configPath)

	// 初始化应用
	application, err := NewApp(*configPath)
	if err != nil {
		fmt.Printf("❌ App initialization failed: %v\n", err)
		return
	}

	// 启动服务
	if err = application.Run(); err != nil {
		fmt.Printf("❌ Server run failed: %v\n", err)
	}
}
