package store

import (
	"database/sql"
	"log/slog"
)

var ErrorNotFound = sql.ErrNoRows
type Storage struct {
	db *sql.DB
	Log *slog.Logger
}

func NewStore(db *sql.DB, log *slog.Logger) *Storage {
	return &Storage{db: db, Log: log}
}
