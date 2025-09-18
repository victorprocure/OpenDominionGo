package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/lib/pq"
	"github.com/victorprocure/opendominiongo/handlers"
	"github.com/victorprocure/opendominiongo/internal/app"
	"github.com/victorprocure/opendominiongo/internal/config"
	"github.com/victorprocure/opendominiongo/internal/datasync"
	intdb "github.com/victorprocure/opendominiongo/internal/db"
	"github.com/victorprocure/opendominiongo/internal/middleware"
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

	// Run data syncs at startup
	coord := datasync.NewSyncCoordinator(sqldb, log)

	syncCtx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	if err := coord.RunAll(syncCtx,
		appSvc.NewTechSync(),
		appSvc.NewRacesSync(),
		appSvc.NewSpellsSync(),
		appSvc.NewWondersSync(),
		appSvc.NewHeroUpgradeSync(),
	); err != nil {
		log.Error("initial sync failed", slog.Any("error", err))
		os.Exit(1)
	}

	// Build the HTTP server
	addr := fmt.Sprintf("localhost:%d", cfg.AppPort)
	// Compose middlewares: logging -> security headers -> session -> handlers
	// Build security header options, allowing override via config
	shOpts := []middleware.SecurityOption{
		middleware.WithHSTS(false, ""), // enable in TLS/production only
	}
	if cfg.AppCSP != "" {
		shOpts = append(shOpts, middleware.WithCSP(cfg.AppCSP))
	}
	chain := middleware.WithRequestID(
		middleware.RequestLogger(log,
			middleware.SecurityHeaders(
				session.NewMiddleware(handler, session.WithSecure(cfg.AppSecure)),
				shOpts...,
			),
		),
	)
	server := &http.Server{
		Addr:              addr,
		Handler:           chain,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       120 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
	}

	// Redact sensitive details from logs; avoid logging full DSN
	log.Info("Server is starting",
		slog.String("addr", server.Addr),
		slog.String("db_host", cfg.DBHost),
		slog.Int("db_port", cfg.DBPort),
		slog.String("db_name", cfg.DBName),
		slog.String("sslmode", cfg.DBSSLMode),
	)

	// Start server and handle graceful shutdown on interrupt/terminate
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error("server listen", slog.Any("error", err))
		}
	}()

	// Wait for shutdown signal
	stopCtx, stopCancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stopCancel()
	<-stopCtx.Done()
	log.Info("shutdown signal received")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Error("server shutdown", slog.Any("error", err))
	} else {
		log.Info("server stopped cleanly")
	}
}
