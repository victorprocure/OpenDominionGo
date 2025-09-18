package repositories

import (
	"context"
	"database/sql"
)

type DbTx interface {
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
}

type RowScanner interface {
	Scan(dest ...any) error
}
