package telescope

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/victorprocure/opendominiongo/internal/middleware"
)

type Service interface {
	HTTPMiddleware() gin.HandlerFunc
}

type responseRecorder struct {
	http.ResponseWriter
	status int
	size   int
}

func (rw *responseRecorder) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}

func (rw *responseRecorder) Write(b []byte) (int, error) {
	if rw.status == 0 {
		rw.status = http.StatusOK
	}
	n, err := rw.ResponseWriter.Write(b)
	rw.size += n
	return n, err
}

// HTTPMiddleware captures request/response metadata as a telescope entry.
func (s *service) HTTPMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		batchID := uuid.New()

		reqID := middleware.GetRequestID(c)
		content := map[string]any{
			"method":      c.Request.Method,
			"path":        c.Request.URL.Path,
			"status":      c.Request.Response.Status,
			"duration_ms": time.Since(start).Milliseconds(),
			"size":        c.Writer.Size(),
			"remote_addr": c.Request.RemoteAddr,
			"user_agent":  c.GetHeader("User-Agent"),
			"request_id":  reqID,
		}
		tags := []string{"http", "method:" + c.Request.Method, "status:" + c.Request.Response.Status}
		if reqID != "" {
			tags = append(tags, "request-id:"+reqID)
		}
		_, _ = s.Capture(c, "request", content, WithBatchID(batchID), WithDisplayOnIndex(true), WithTags(tags...))
	}
}
