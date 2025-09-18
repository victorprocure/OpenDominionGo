package handlers

import (
	"net/http"

	"github.com/victorprocure/opendominiongo/components"
)

func (h *Handler) HandleLoginRibbon(w http.ResponseWriter, r *http.Request) {
	components.LoginRibbon(components.LoginRibbonOpts{LoggedIn: false}).Render(r.Context(), w)
}
