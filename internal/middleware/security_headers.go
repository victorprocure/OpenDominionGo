package middleware

import "net/http"

// SecurityHeaders sets common security-related HTTP headers.
// Note: Adjust CSP to your frontend assets and inline usage.
// DefaultCSP is a conservative baseline. Adjust in production to remove 'unsafe-inline' when possible.
const DefaultCSP = "default-src 'self'; img-src 'self' data:; style-src 'self' 'unsafe-inline'; script-src 'self'; object-src 'none'; frame-ancestors 'self'; base-uri 'self'"

func SecurityHeaders(next http.Handler, opts ...SecurityOption) http.Handler {
	cfg := securityConfig{
		CSP:                 DefaultCSP,
		ReferrerPolicy:      "no-referrer",
		XContentTypeOptions: "nosniff",
		XFrameOptions:       "SAMEORIGIN",
		StrictTransport:     "max-age=31536000; includeSubDomains",
		EnableHSTS:          false,
	}
	for _, o := range opts {
		o(&cfg)
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Security-Policy", cfg.CSP)
		w.Header().Set("Referrer-Policy", cfg.ReferrerPolicy)
		w.Header().Set("X-Content-Type-Options", cfg.XContentTypeOptions)
		w.Header().Set("X-Frame-Options", cfg.XFrameOptions)
		if cfg.EnableHSTS && r.TLS != nil {
			w.Header().Set("Strict-Transport-Security", cfg.StrictTransport)
		}
		next.ServeHTTP(w, r)
	})
}

type securityConfig struct {
	CSP                 string
	ReferrerPolicy      string
	XContentTypeOptions string
	XFrameOptions       string
	StrictTransport     string
	EnableHSTS          bool
}

type SecurityOption func(*securityConfig)

func WithCSP(csp string) SecurityOption { return func(c *securityConfig) { c.CSP = csp } }
func WithReferrerPolicy(p string) SecurityOption {
	return func(c *securityConfig) { c.ReferrerPolicy = p }
}

func WithHSTS(enabled bool, value string) SecurityOption {
	return func(c *securityConfig) {
		c.EnableHSTS = enabled
		if value != "" {
			c.StrictTransport = value
		}
	}
}
