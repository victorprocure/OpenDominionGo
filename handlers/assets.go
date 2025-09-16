package handlers

import "net/http"

func (h *Handler) HandleAssets(w http.ResponseWriter, r *http.Request) {
	http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))).ServeHTTP(w, r)
}

func (h *Handler) HandleFavicon(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "assets/favicon.ico")
}