package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// Logger is a middleware that logs request details
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method

		c.Next()

		latency := time.Since(start)
		statusCode := c.Writer.Status()

		// Using log package instead of c.Logger() which doesn't exist
		log.Printf("%s %s %d %s", method, path, statusCode, latency)
	}
}

// ErrorHandler is a middleware that catches panics and returns 500 errors
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.JSON(500, gin.H{
					"success": false,
					"message": "Internal server error",
				})
			}
		}()

		c.Next()
	}
}
