package middleware

import (
    "net/http"
)

// Recoverer captures panics, returns 500, and logs via the provided function.
func Recoverer(logf func(msg string, args ...any)) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            defer func() {
                if rec := recover(); rec != nil {
                    logf("panic recovered: %v", rec)
                    http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
                }
            }()
            next.ServeHTTP(w, r)
        })
    }
}
