package handlers

import (
	"net/http"
	"net/url"

	"github.com/victorprocure/opendominiongo/components"
)

func (h *Handler) HandleNavBar(w http.ResponseWriter, r *http.Request) {
	// HTMX sends HX-Current-URL with the browser's current page URL.
	// Fall back to Referer, then to the request path.
	current := r.Header.Get("HX-Current-URL")
	path := r.URL.Path
	if current != "" {
		if u, err := url.Parse(current); err == nil {
			path = u.Path
		}
	} else if ref := r.Referer(); ref != "" {
		if u, err := url.Parse(ref); err == nil {
			path = u.Path
		}
	}
	
	err := components.NavItems(getNavItems(path, h)).Render(r.Context(), w)
	if err != nil {
		h.Log.Error("render nav bar", "error", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func getNavItems(currentPath string, h *Handler) []components.NavItem {
	h.Log.Info("Generating nav items", "currentPath", currentPath)

	items := []components.NavItem{
		{Title: "About", Link: "/about"},
		{Title: "Valhalla", Link: "/valhalla"},
		{Title: "Scribes", Link: "/scribes"},
		{Title: "Wiki", Link: "https://wiki.opendominion.net/", NewWindow: true},
		{Title: "Rules", Link: "/rules"},
	}

	// Mark the item selected if the current path equals its Link
	for i := range items {
		if items[i].Link == currentPath {
			items[i].Selected = true
		}
	}
	return items
}
