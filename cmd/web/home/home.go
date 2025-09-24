package home

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var path = "/"

func RegisterRoutes(r *gin.Engine) {
	r.GET(path, func(c *gin.Context){
		homeWebHandler(c.Writer, c.Request)
	})
}

func homeWebHandler(w http.ResponseWriter, r *http.Request) {
	component := Home()
	err := component.Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Failed to render home page", http.StatusInternalServerError)
		return
	}
}