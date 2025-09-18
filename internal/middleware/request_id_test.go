package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWithRequestID_GeneratesAndSetsHeader(t *testing.T) {
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if id := GetRequestID(r.Context()); id == "" {
			t.Fatalf("expected generated request id in context")
		}
		w.WriteHeader(200)
	})
	h := WithRequestID(next)
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	h.ServeHTTP(rec, req)

	if got := rec.Header().Get(HeaderRequestID); got == "" {
		t.Fatalf("expected %s header on response", HeaderRequestID)
	}
}

func TestWithRequestID_RespectsIncomingHeader(t *testing.T) {
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if id := GetRequestID(r.Context()); id != "abc123" {
			t.Fatalf("expected request id propagated, got %q", id)
		}
		w.WriteHeader(200)
	})
	h := WithRequestID(next)
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(HeaderRequestID, "abc123")

	h.ServeHTTP(rec, req)

	if got := rec.Header().Get(HeaderRequestID); got != "abc123" {
		t.Fatalf("expected header propagated, got %q", got)
	}
}
