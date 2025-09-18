package telescope

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/victorprocure/opendominiongo/internal/middleware"
)

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
func (s *Service) HTTPMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rr := &responseRecorder{ResponseWriter: w}
		batchID := uuid.New()
		next.ServeHTTP(rr, r)

		reqID := middleware.GetRequestID(r.Context())
		content := map[string]any{
			"method":      r.Method,
			"path":        r.URL.Path,
			"status":      rr.status,
			"duration_ms": time.Since(start).Milliseconds(),
			"size":        rr.size,
			"remote_addr": r.RemoteAddr,
			"user_agent":  r.Header.Get("User-Agent"),
			"request_id":  reqID,
		}
		tags := []string{"http", "method:" + r.Method, "status:" + http.StatusText(rr.status)}
		if reqID != "" {
			tags = append(tags, "request-id:"+reqID)
		}
		_, _ = s.Capture(r.Context(), "request", content, WithBatchID(batchID), WithDisplayOnIndex(true), WithTags(tags...))
	})
}
