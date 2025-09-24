package server

import (
	"fmt"
	"io/fs"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/victorprocure/opendominiongo/cmd/web"
	"github.com/victorprocure/opendominiongo/cmd/web/about"
	"github.com/victorprocure/opendominiongo/cmd/web/home"
	"github.com/victorprocure/opendominiongo/cmd/web/login"
	"github.com/victorprocure/opendominiongo/cmd/web/rules"
	"github.com/victorprocure/opendominiongo/cmd/web/shared"
	"github.com/victorprocure/opendominiongo/internal/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.New()
	r.Use(middleware.WithRequestID())
	// r.Use(s.deps.Telescope.HTTPMiddleware())
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{fmt.Sprintf("http://localhost:%d", s.cfg.AppPort)},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", middleware.HeaderRequestID},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.GET("/health", s.healthHandler)
	staticFiles, _ := fs.Sub(web.Files, "assets")
	r.StaticFS("/assets", http.FS(staticFiles))

	about.RegisterRoutes(r)
	home.RegisterRoutes(r)
	login.RegisterRoutes(r)
	rules.RegisterRoutes(r)
	shared.RegisterRoutes(r)

	return r
}

func (s *Server) healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, s.deps.DBService.Health())
}
