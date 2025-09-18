package handlers

import (
	"net/http"

	"github.com/victorprocure/opendominiongo/components"
)

func (h *Handler) HandleHome(w http.ResponseWriter, r *http.Request) {
	err := components.Home().Render(r.Context(), w)
	if err != nil {
		h.Log.Error("render home", "error", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
