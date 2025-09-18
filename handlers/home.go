package handlers

import (
	"net/http"

	"github.com/victorprocure/opendominiongo/components"
)

func (h *Handler) HandleHome(w http.ResponseWriter, r *http.Request) {
	components.Home().Render(r.Context(), w)
}
