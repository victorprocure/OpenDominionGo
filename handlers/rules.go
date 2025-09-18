package handlers

import (
	"net/http"

	"github.com/victorprocure/opendominiongo/components"
)

func (h *Handler) HandleRules(w http.ResponseWriter, r *http.Request) {
	err := components.Rules().Render(r.Context(), w)
	if err != nil {
		h.Log.Error("render rules", "error", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
