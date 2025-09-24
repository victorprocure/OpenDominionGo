package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/lib/pq"
	"github.com/victorprocure/opendominiongo/internal/config"
	"github.com/victorprocure/opendominiongo/internal/datasync"
	"github.com/victorprocure/opendominiongo/internal/db"
	tel "github.com/victorprocure/opendominiongo/internal/observability/telescope"
	"github.com/victorprocure/opendominiongo/internal/server"
)

func gracefulShutdown(apiServer *http.Server, done chan bool) {
	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Listen for the interrupt signal.
	<-ctx.Done()

	log.Println("shutting down gracefully, press Ctrl+C again to force")
	stop() // Allow Ctrl+C to force shutdown

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := apiServer.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown with error: %v", err)
	}

	log.Println("Server exiting")

	// Notify the main goroutine that the shutdown is complete
	done <- true
}

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stderr, nil))
	// Load configuration from .env and environment
	cfg, err := config.Load()
	if err != nil {
		log.Error("config load", slog.Any("error", err))
		// ensure any outstanding cancel funcs run before exiting
		// none are in scope here, so exit directly
		os.Exit(1)
	}
	cfg.Log = log

	db, err := db.New(cfg)
	if err != nil {
		log.Error("db open", slog.Any("error", err))
		// ensure any outstanding cancel funcs run before exiting
		// none are in scope here, so exit directly
		os.Exit(1)
	}
	defer db.Close()

	telescope := tel.NewService(db.RawDB(), cfg)
	datasync := datasync.NewSyncCoordinator(db.RawDB(), log)

	deps := server.Dependencies{
		SyncCoordinator: datasync,
		Telescope:       telescope,
		DBService:       db,
	}
	server := server.NewServer(cfg, deps)
	httpServer := server.HTTPServer()

	done := make(chan bool, 1)
	go gracefulShutdown(httpServer, done)

	log.Info("Server is starting",
		slog.String("addr", httpServer.Addr),
		slog.String("db_host", cfg.DBHost),
		slog.Int("db_port", cfg.DBPort),
		slog.String("db_name", cfg.DBName),
		slog.String("sslmode", cfg.DBSSLMode),
	)

	err = httpServer.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		panic(fmt.Sprintf("http server error: %s", err))
	}

	<-done
	log.Info("graceful shutdown complete")
}
