package shared

import (
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

var path = "/loadNavBar"

func registerNavBarRoutes(r *gin.Engine) {
	r.GET(path, func(c *gin.Context) {
		navBarWebHandler(c.Writer, c.Request)
	})
}

func navBarWebHandler(w http.ResponseWriter, r *http.Request) {
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

	ni := getNavItems(path)
	component := NavItems(ni)
	err := component.Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Failed to render navigation bar", http.StatusInternalServerError)
		return
	}
}

func getNavItems(currentPath string) []NavItem {
	items := []NavItem{
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
