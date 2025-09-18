package handlers

import (
	"net/http"

	"github.com/victorprocure/opendominiongo/components"
)

func (h *Handler) HandleRules(w http.ResponseWriter, r *http.Request) {
	components.Rules().Render(r.Context(), w)
}
