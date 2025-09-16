package store

import (
	"context"
	"database/sql"
	"log/slog"
)

type DbTx interface {
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
}

type RowScanner interface {
    Scan(dest ...any) error
}

func NewSqlStorage(connectionString string, log slog.Logger) (*sql.DB, error) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Error("Error opening database connection: ", slog.Any("error", err))
		return nil, err
	}

	return db, nil
}

func (s *Storage) InitDataSync(syncers ...DataSync) *SyncCoordinator {
	return NewSyncCoordinator(s, syncers...)
}

func (s *Storage) Ping() error {
	return s.db.Ping()
}