package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	//"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/victorprocure/opendominiongo/handlers"
	"github.com/victorprocure/opendominiongo/session"
	"github.com/victorprocure/opendominiongo/store"
)

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stderr, nil))
	config := store.Envs
	cfg := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", config.DBUser, config.DBPassword, config.DBAddress, config.DBName)
	db, err := store.NewSqlStorage(cfg, *log)
	if err != nil {
		log.Error("Error connecting to the database: ", slog.Any("error", err))
		os.Exit(1)
	}
	defer db.Close()

	store := store.NewStore(db, log)
	initStore(store, log)

	handler := handlers.New(store, log)
	secureFlag := config.SecureFlag
	sessionHandler := session.NewMiddleware(handler, session.WithSecure(secureFlag))

	server := &http.Server{
		Addr: "localhost:" + config.Port,
		Handler: sessionHandler,
		ReadTimeout: time.Second * 10,
		WriteTimeout: time.Second * 10,
	}

	err = store.PerformDataSync()
	if err != nil {
		log.Error("Unable to perform data sync", slog.Any("error", err))
		os.Exit(1)
	}

	log.Info("Server is running", slog.Group("Config", slog.String("Server Address", server.Addr), slog.String("DB Address", config.DBAddress), slog.String("DB Name", config.DBName)))
	server.ListenAndServe()
}

func initStore(db *store.Storage, log *slog.Logger) {
	err := db.Ping()
	if err != nil {
		log.Error("Error pinging the database: ", slog.Any("error", err))
		os.Exit(1)
	}
	
	log.Debug("Database connection established")
}