package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

// responseWriter wraps http.ResponseWriter to capture status code and bytes written.
type responseWriter struct {
	w      http.ResponseWriter
	status int
	size   int
}

func (rw *responseWriter) Header() http.Header { return rw.w.Header() }
func (rw *responseWriter) Write(b []byte) (int, error) {
	if rw.status == 0 {
		rw.status = http.StatusOK
	}
	n, err := rw.w.Write(b)
	rw.size += n
	return n, err
}
func (rw *responseWriter) WriteHeader(statusCode int) {
	rw.status = statusCode
	rw.w.WriteHeader(statusCode)
}

// RequestLogger logs request details and duration using slog.
func RequestLogger(log *slog.Logger, next http.Handler) http.Handler {
	if log == nil {
		return next
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rw := &responseWriter{w: w}
		next.ServeHTTP(rw, r)
		dur := time.Since(start)

		attrs := []slog.Attr{
			slog.String("method", r.Method),
			slog.String("path", r.URL.Path),
			slog.Int("status", rw.status),
			slog.Int("bytes", rw.size),
			slog.String("remote", r.RemoteAddr),
			slog.String("ua", r.UserAgent()),
			slog.Duration("duration", dur),
		}
		if rid := GetRequestID(r.Context()); rid != "" {
			attrs = append(attrs, slog.String("request_id", rid))
		}
		log.LogAttrs(r.Context(), slog.LevelInfo, "http_request", attrs...)
	})
}
