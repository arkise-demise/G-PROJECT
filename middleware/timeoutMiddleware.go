package middleware

import (
	"G-PROJECT/utils"
	"context"
	"net/http"
	"time"

	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
)

type contextKey string

const (
	requestIDKey contextKey = "requestID"
)

func TimeoutMiddleware() gin.HandlerFunc {
	return timeout.New(
		timeout.WithTimeout(5*time.Second),
		timeout.WithHandler(func(c *gin.Context) {
			requestID := utils.GenerateRequestID()
			ctx := c.Request.Context()
			ctx = context.WithValue(ctx, requestIDKey, requestID)
			c.Request = c.Request.WithContext(ctx)
			c.Next()
		}),
		timeout.WithResponse(func(c *gin.Context) {
			requestID := c.Request.Context().Value(requestIDKey)
			if requestID != nil {
				c.JSON(http.StatusRequestTimeout, gin.H{"error": "Request timed out"})
			}
		}),
	)
}
