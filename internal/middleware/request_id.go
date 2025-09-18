package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/segmentio/ksuid"
)

// ctxKey is an unexported type for context keys defined in this package.
type ctxKey string

const (
	requestIDKey    ctxKey = "request_id"
	HeaderRequestID        = "X-Request-Id"
)

// WithRequestID ensures each request has a stable ID, taken from header or generated.
// It injects the ID into the context and sets the response header as well.
func WithRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.Header.Get(HeaderRequestID)
		id = strings.TrimSpace(id)
		if id == "" {
			id = ksuid.New().String()
		}

		ctx := context.WithValue(r.Context(), requestIDKey, id)
		w.Header().Set(HeaderRequestID, id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GetRequestID returns the request id from context or empty string.
func GetRequestID(ctx context.Context) string {
	v := ctx.Value(requestIDKey)
	if s, ok := v.(string); ok {
		return s
	}
	return ""
}
