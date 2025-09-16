package handlers

import (
	"net/http"

	"github.com/victorprocure/opendominiongo/components"
)

func (h *Handler) HandleAbout(w http.ResponseWriter, r *http.Request) {
	components.About().Render(r.Context(), w)
}