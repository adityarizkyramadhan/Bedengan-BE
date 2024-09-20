package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func CheckToken(redis *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		// Jika kosong maka izinkan request namun jika ada dan sudah logout maka tolak
		if token == "" {
			c.Next()
			return
		}

		token = strings.Replace(token, "Bearer ", "", 1)

		// Jika ada token pada redis, maka token invalid karena sudah logout
		if _, err := redis.Get(c, token).Result(); err == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is invalid"})
			c.Abort()
			return
		}

		c.Next()
	}
}
