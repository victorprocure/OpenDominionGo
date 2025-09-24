package about

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var path = "/about"

func RegisterRoutes(r *gin.Engine) {
	r.GET(path, func(c *gin.Context) {
		aboutWebHandler(c.Writer, c.Request)
	})
}

func aboutWebHandler(w http.ResponseWriter, r *http.Request) {
	component := About()
	err := component.Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Failed to render about page", http.StatusInternalServerError)
		return
	}
}
