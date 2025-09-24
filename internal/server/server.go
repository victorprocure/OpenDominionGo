package server

import (
	"fmt"
	"net/http"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/victorprocure/opendominiongo/internal/config"
	"github.com/victorprocure/opendominiongo/internal/datasync"
	"github.com/victorprocure/opendominiongo/internal/db"
	"github.com/victorprocure/opendominiongo/internal/observability/telescope"
)

type Dependencies struct {
	Telescope       telescope.Service
	SyncCoordinator datasync.SyncCoordinator
	DBService       db.Service
}

type Server struct {
	cfg  *config.AppConfig
	deps Dependencies
}

func NewServer(cfg *config.AppConfig, deps Dependencies) *Server {
	NewServer := &Server{
		cfg:  cfg,
		deps: deps,
	}

	return NewServer
}

func (s *Server) HTTPServer() *http.Server {
	server := &http.Server{
		Addr:              fmt.Sprintf("localhost:%d", s.cfg.AppPort),
		Handler:           s.RegisterRoutes(),
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       120 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
	}

	return server
}
