package middleware

import (
	"net/http"
	"time"

	"github.com/juju/ratelimit"

	"github.com/gin-gonic/gin"
)

func RateLimitMiddleware(fillInterval time.Duration, cap int64) func(c *gin.Context) {
	bucket := ratelimit.NewBucket(fillInterval, cap) // 令牌发送的速率和令牌桶的容量
	return func(c *gin.Context) {
		if bucket.TakeAvailable(1) == 0 { // 每次取一个
			c.String(http.StatusOK, "rate limit...")
			c.Abort()
			return
		}
		c.Next()
	}
}
