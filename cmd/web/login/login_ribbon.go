package login

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var path = "/loadLoginRibbon"

func RegisterRoutes(r *gin.Engine) {
	r.GET(path, func(c *gin.Context) {
		loginRibbonWebHandler(c.Writer, c.Request)
	})
}

func loginRibbonWebHandler(w http.ResponseWriter, r *http.Request) {
	component := LoginRibbon(LoginRibbonOpts{
		LoggedIn: false, // This would be dynamic based on session/auth status
	})
	err := component.Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Failed to render login ribbon", http.StatusInternalServerError)
		return
	}
}
