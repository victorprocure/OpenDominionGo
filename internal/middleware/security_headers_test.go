package middleware

import (
	"crypto/tls"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSecurityHeaders_Defaults(t *testing.T) {
	next := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) { w.WriteHeader(200) })
	h := SecurityHeaders(next)
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)
	res := rec.Result()

	if got := res.Header.Get("Content-Security-Policy"); got != DefaultCSP {
		t.Fatalf("CSP mismatch: %q", got)
	}
	if got := res.Header.Get("X-Content-Type-Options"); got != "nosniff" {
		t.Fatalf("X-Content-Type-Options mismatch: %q", got)
	}
	if got := res.Header.Get("X-Frame-Options"); got != "SAMEORIGIN" {
		t.Fatalf("X-Frame-Options mismatch: %q", got)
	}
	if got := res.Header.Get("Referrer-Policy"); got != "no-referrer" {
		t.Fatalf("Referrer-Policy mismatch: %q", got)
	}
	if st := res.Header.Get("Strict-Transport-Security"); st != "" {
		t.Fatalf("unexpected HSTS header without TLS: %q", st)
	}
}

func TestSecurityHeaders_OverrideCSP_And_HSTS(t *testing.T) {
	next := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) { w.WriteHeader(200) })
	h := SecurityHeaders(next, WithCSP("default-src 'self'"), WithHSTS(true, "max-age=123"))
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	// simulate TLS by setting a non-nil TLS ConnectionState
	req.TLS = &tls.ConnectionState{}
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)
	res := rec.Result()

	if got := res.Header.Get("Content-Security-Policy"); got != "default-src 'self'" {
		t.Fatalf("CSP override failed: %q", got)
	}
	if got := res.Header.Get("Strict-Transport-Security"); got != "max-age=123" {
		t.Fatalf("HSTS mismatch: %q", got)
	}
}
