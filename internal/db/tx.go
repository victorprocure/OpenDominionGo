package db

import (
	"context"
	"database/sql"
)

// WithTx runs fn within a sql.Tx and commits or rolls back.
// The fn receives the transaction to allow passing it to repo methods that accept DbTx.
func WithTx(ctx context.Context, db *sql.DB, opts *sql.TxOptions, fn func(*sql.Tx) error) error {
	tx, err := db.BeginTx(ctx, opts)
	if err != nil {
		return err
	}
	defer func() {
		// If not committed due to error, rollback
		_ = tx.Rollback()
	}()
	if err := fn(tx); err != nil {
		return err
	}
	return tx.Commit()
}
