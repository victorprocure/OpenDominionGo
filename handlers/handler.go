package handlers

import (
	"log/slog"
	"net/http"

	"github.com/victorprocure/opendominiongo/components"
	"github.com/victorprocure/opendominiongo/internal/app"
)

func New(a *app.App, log *slog.Logger) *Handler {
	h := &Handler{
		app: a,
		Log: log,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/assets/", h.HandleAssets)
	mux.HandleFunc("/favicon.ico", h.HandleFavicon)
	mux.HandleFunc("/healthz", h.HandleHealth)
	mux.HandleFunc("/loadNavBar", h.HandleNavBar)
	mux.HandleFunc("/loadLoginRibbon", h.HandleLoginRibbon)
	mux.HandleFunc("/about", h.HandleAbout)
	mux.HandleFunc("/rules", h.HandleRules)
	// Catch-all for dynamic pages (decides between Home and 404)
	mux.HandleFunc("/", h.View)

	h.mux = mux
	return h
}

type Handler struct {
	app *app.App
	Log *slog.Logger
	mux *http.ServeMux
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet, http.MethodHead:
		h.Get(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	// Delegate routing to the internal mux: assets, favicon, nav, or View("/")
	h.mux.ServeHTTP(w, r)
}

// View handles dynamic pages behind the catch-all "/" route.
// It sends unknown paths to a friendly 404 page.
func (h *Handler) View(w http.ResponseWriter, r *http.Request) {
	h.Log.Debug("Handling view for path", slog.String("path", r.URL.Path))
	switch r.URL.Path {
	case "/":
		h.HandleHome(w, r)
	default:
		h.NotFound(w, r)
	}
}

// NotFound renders a friendly 404 page
func (h *Handler) NotFound(w http.ResponseWriter, r *http.Request) {
	h.renderErrorPage(w, r, http.StatusNotFound, "Page not found", "The page you're looking for doesn't exist or may have been moved.")
}

// Unauthorized renders a 401 page
func (h *Handler) Unauthorized(w http.ResponseWriter, r *http.Request) {
	h.renderErrorPage(w, r, http.StatusUnauthorized, "Unauthorized", "You must be signed in to access this page.")
}

// Forbidden renders a 403 page
func (h *Handler) Forbidden(w http.ResponseWriter, r *http.Request) {
	h.renderErrorPage(w, r, http.StatusForbidden, "Forbidden", "You don't have permission to access this resource.")
}

// renderErrorPage writes a consistent styled error page using your existing assets
func (h *Handler) renderErrorPage(w http.ResponseWriter, r *http.Request, status int, title, message string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(status)

	err := components.Error(title, message, status).Render(r.Context(), w)
	if err != nil {
		h.Log.Error("render error page", "error", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

// HandleHealth provides a minimal liveness/readiness check.
// It pings the database connection to ensure dependencies are OK.
func (h *Handler) HandleHealth(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	if err := h.app.DB.PingContext(ctx); err != nil {
		http.Error(w, "unhealthy", http.StatusServiceUnavailable)
		return
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("ok"))
}
