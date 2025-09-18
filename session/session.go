package session

import (
	"net/http"

	"github.com/segmentio/ksuid"
)

type (
	MiddlewareOpts func(*Middleware)
	Middleware     struct {
		Next     http.Handler
		Secure   bool
		HTTPOnly bool
		MaxAge   int // seconds; 0 means session cookie
	}
)

func NewMiddleware(next http.Handler, opts ...MiddlewareOpts) http.Handler {
	mw := Middleware{
		Next:     next,
		Secure:   true,
		HTTPOnly: true,
		MaxAge:   0,
	}

	for _, opt := range opts {
		opt(&mw)
	}

	return mw
}

func WithSecure(secure bool) MiddlewareOpts {
	return func(m *Middleware) {
		m.Secure = secure
	}
}

func WithHTTPOnly(httpOnly bool) MiddlewareOpts {
	return func(m *Middleware) {
		m.HTTPOnly = httpOnly
	}
}

func WithMaxAge(seconds int) MiddlewareOpts {
	return func(m *Middleware) {
		m.MaxAge = seconds
	}
}

func ID(r *http.Request) (id string) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		return id
	}

	return cookie.Value
}

func (mw Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := ID(r)
	if id == "" {
		id = ksuid.New().String()
		http.SetCookie(w, &http.Cookie{
			Name:     "session_id",
			Value:    id,
			Path:     "/",
			Secure:   mw.Secure,
			HttpOnly: mw.HTTPOnly,
			SameSite: http.SameSiteLaxMode,
			MaxAge:   mw.MaxAge,
		})
	}
	mw.Next.ServeHTTP(w, r)
}
