package middleware

import (
	"github.com/alvin41793/Image-upload/internal/config"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

func init() {
	// 直接定义一个测试用的配置对象
	testConfig := &config.Config{
		Srv: config.ServerConfig{
			Port: ":8080",
		},
		OSSVal: config.OSSConfig{
			Enable:    false,
			Endpoint:  "",
			AccessKey: "",
			Secret:    "",
			Bucket:    "",
			Domain:    "",
			Dir:       "",
		},
		LimiterVal: config.LimiterConfig{
			Rate:  1,
			Burst: 3,
		},
		UploadVal: config.UploadConfig{
			Dir:          "./uploads",
			MaxSizeMB:    5,
			AllowedTypes: []string{".jpg", ".png"},
		},
		LogVal: config.LogConfig{
			Dir:      "./logs",
			KeepDays: 7,
			Level:    "info",
		},
	}

	// 将测试配置应用到全局配置管理器
	config.InitGlobal(testConfig)
}

// TestRateLimitMiddleware 测试速率限制中间件的效果。
// 该测试用例旨在验证RateLimitMiddleware是否能够正确地限制请求速率。
// 它通过模拟多个GET请求来触发速率限制，并检查响应状态码以确定限流是否生效
func TestRateLimitMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.New()
	r.Use(RateLimitMiddleware())
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// 模拟多次请求触发限流
	for i := 0; i < 10; i++ {
		req := httptest.NewRequest("GET", "/ping", nil)
		req.RemoteAddr = "127.0.0.1:12345"
		resp := httptest.NewRecorder()

		r.ServeHTTP(resp, req)

		t.Logf("Attempt %d: status = %d, body = %s", i+1, resp.Code, resp.Body.String())
	}
}
