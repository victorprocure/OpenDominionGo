package middleware

import (
	"context"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"
)

type testHandler struct{}

func (h testHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(204)
	_, _ = w.Write([]byte{})
}

type captureHandler struct {
	mu   sync.Mutex
	recs []slog.Record
}

func (h *captureHandler) Enabled(context.Context, slog.Level) bool { return true }
func (h *captureHandler) Handle(_ context.Context, rec slog.Record) error {
	h.mu.Lock()
	h.recs = append(h.recs, rec.Clone())
	h.mu.Unlock()
	return nil
}
func (h *captureHandler) WithAttrs([]slog.Attr) slog.Handler { return h }
func (h *captureHandler) WithGroup(string) slog.Handler      { return h }

func TestRequestLogger_Basic(t *testing.T) {
	cap := &captureHandler{}
	logger := slog.New(cap)

	next := testHandler{}
	h := WithRequestID(RequestLogger(logger, next))
	req := httptest.NewRequest(http.MethodGet, "/x", nil)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	if rec.Code != 204 {
		t.Fatalf("status code mismatch: %d", rec.Code)
	}
	if len(cap.recs) == 0 {
		t.Fatalf("expected at least one log record")
	}
	// sanity check duration present and non-negative
	foundDur := false
	foundRID := false
	cap.recs[0].Attrs(func(a slog.Attr) bool {
		if a.Key == "duration" {
			if d, ok := a.Value.Any().(time.Duration); !ok || d < 0 {
				t.Fatalf("invalid duration attr: %#v", a.Value)
			}
			foundDur = true
		}
		if a.Key == "request_id" {
			if s, ok := a.Value.Any().(string); !ok || s == "" {
				t.Fatalf("invalid request_id attr: %#v", a.Value)
			}
			foundRID = true
		}
		return true
	})
	if !foundDur {
		t.Fatalf("duration attribute not found in log record")
	}
	if !foundRID {
		t.Fatalf("request_id attribute not found in log record")
	}
}
