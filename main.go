package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
	"github.com/victorprocure/opendominiongo/handlers"
	"github.com/victorprocure/opendominiongo/internal/app"
	"github.com/victorprocure/opendominiongo/internal/config"
	intdb "github.com/victorprocure/opendominiongo/internal/db"
	"github.com/victorprocure/opendominiongo/session"
)

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stderr, nil))

	// Load configuration from .env and environment
	cfg, err := config.Load()
	if err != nil {
		log.Error("config load", slog.Any("error", err))
		os.Exit(1)
	}

	// Open DB with pooling
	sqldb, err := intdb.OpenPostgres(context.Background(), cfg)
	if err != nil {
		log.Error("db open", slog.Any("error", err))
		os.Exit(1)
	}
	defer sqldb.Close()

	// Build application service and handlers
	appSvc := app.New(sqldb, log)
	handler := handlers.New(appSvc, log)

	// Build the HTTP server
	addr := fmt.Sprintf("localhost:%d", cfg.AppPort)
	sessionHandler := session.NewMiddleware(handler, session.WithSecure(cfg.AppSecure))
	server := &http.Server{
		Addr:         addr,
		Handler:      sessionHandler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Info("Server is running", slog.String("addr", server.Addr), slog.String("dsn", cfg.BuildPostgresDSN()))
	if err := server.ListenAndServe(); err != nil {
		log.Error("server", slog.Any("error", err))
	}
}
