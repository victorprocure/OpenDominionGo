package shared

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine) {
	registerNavBarRoutes(r)
}