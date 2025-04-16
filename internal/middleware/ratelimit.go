package middleware

import (
	"github.com/alvin41793/Image-upload/internal/config"
	"github.com/alvin41793/Image-upload/internal/logger"
	"github.com/alvin41793/Image-upload/internal/util"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
	"sync"
	"time"
)

var (
	ipLimiters  = sync.Map{} // thread-safe map
	cleanupOnce sync.Once
)

func getLimiter(ip string) *rate.Limiter {
	cfg := config.G().Limiter()
	limiter, exists := ipLimiters.Load(ip)
	if exists {
		return limiter.(*rate.Limiter)
	}
	newLimiter := rate.NewLimiter(rate.Limit(cfg.Rate), cfg.Burst)
	ipLimiters.Store(ip, newLimiter)
	return newLimiter
}

// 清理 map 中过期 IP（可选高级优化）
func startLimiterCleanup() {
	ticker := time.NewTicker(10 * time.Minute)
	go func() {
		for range ticker.C {
			// 在实际项目中结合 TTL 缓存或 IP 活跃时间判断是否需要清除
			// 此处留作拓展
		}
	}()
}

// 限流中间件
func RateLimitMiddleware() gin.HandlerFunc {
	// 启动一次性清理协程（可选）
	cleanupOnce.Do(startLimiterCleanup)

	return func(c *gin.Context) {
		ip := c.ClientIP()
		limiter := getLimiter(ip)
		if !limiter.Allow() {
			c.Header("Retry-After", "1")
			util.Fail(c, 429, "Requests are too frequent")
			logger.L().Warn("Too many requests", zap.String("ip", ip))
			c.Abort()
			return
		}
		c.Next()
	}
}
