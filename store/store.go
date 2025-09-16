package store

import (
	"database/sql"
	"errors"
	"log/slog"
)

var ErrorNotFound = errors.New("record not found")
type Storage struct {
	db *sql.DB
	Log *slog.Logger
}

func NewStore(db *sql.DB, log *slog.Logger) *Storage {
	return &Storage{db: db, Log: log}
}