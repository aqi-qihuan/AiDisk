package core

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.GetHeader("Origin")
		if origin == "" {
			c.Next()
			return
		}

		cfg := GetConfig()
		allowed := []string{cfg.FrontendBaseURL}
		if cfg.Debug {
			allowed = append(allowed,
				"http://localhost:8081",
				"http://127.0.0.1:8081",
				"http://localhost:5173",
				"http://127.0.0.1:5173",
				"http://localhost:3000",
				"http://127.0.0.1:3000",
			)
		}

		allow := "*"
		for _, o := range allowed {
			if o == origin {
				allow = origin
				break
			}
		}

		c.Header("Access-Control-Allow-Origin", allow)
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, token, Authorization, share-token")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Content-Disposition")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}
