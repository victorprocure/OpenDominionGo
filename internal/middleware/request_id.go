package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/ksuid"
)

// ctxKey is an unexported type for context keys defined in this package.
type ctxKey string

const (
	requestIDKey    ctxKey = "request_id"
	HeaderRequestID string = "X-Request-ID"
)

func WithRequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqID := c.GetHeader(HeaderRequestID)
		reqID = strings.TrimSpace(reqID)
		if reqID == "" {
			reqID = ksuid.New().String()
		}

		c.Set(string(requestIDKey), reqID)
		c.Writer.Header().Set(HeaderRequestID, reqID)
		c.Next()
	}
}

// GetRequestID returns the request id from context or empty string.
func GetRequestID(ctx *gin.Context) string {
	v := ctx.Value(requestIDKey)
	if s, ok := v.(string); ok {
		return s
	}
	return ""
}
