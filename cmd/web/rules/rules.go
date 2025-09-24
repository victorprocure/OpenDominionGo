package rules

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var path = "/rules"

func RegisterRoutes(r *gin.Engine) {
	r.GET(path, func(c *gin.Context){
		rulesWebHandler(c.Writer, c.Request)
	})
}

func rulesWebHandler(w http.ResponseWriter, r *http.Request) {
	component := Rules()
	err := component.Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Failed to render about page", http.StatusInternalServerError)
		return
	}
}