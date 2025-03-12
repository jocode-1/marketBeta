package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
	"time"
)

var (
	requests = make(map[string]int)
	mu       sync.Mutex
)

// RateLimitMiddleware limits requests to 1 per second
func RateLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()

		mu.Lock()
		count := requests[ip]
		if count >= 5 { // Allow burst of 5 requests
			mu.Unlock()
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Too many requests, slow down"})
			c.Abort()
			return
		}

		requests[ip]++
		mu.Unlock()

		time.AfterFunc(time.Second, func() {
			mu.Lock()
			requests[ip]--
			mu.Unlock()
		})

		c.Next()
	}
}
