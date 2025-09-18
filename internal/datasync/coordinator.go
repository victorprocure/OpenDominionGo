package datasync

import (
	"context"
	"database/sql"
	"log/slog"

	"github.com/victorprocure/opendominiongo/internal/db"
)

// SyncCoordinator runs a set of Syncers inside DB transactions.
type SyncCoordinator struct {
	DB  *sql.DB
	Log *slog.Logger
}

// NewSyncCoordinator constructs a coordinator with the DB and logger.
func NewSyncCoordinator(db *sql.DB, log *slog.Logger) *SyncCoordinator {
	return &SyncCoordinator{DB: db, Log: log}
}

// RunAll executes the provided syncers sequentially. Each syncer is executed inside its own transaction.
// If any syncer returns an error, RunAll stops and returns that error.
func (c *SyncCoordinator) RunAll(ctx context.Context, syncers ...Syncer) error {
	for _, s := range syncers {
		name := s.Name()
		c.Log.Info("starting sync", slog.String("sync", name))
		err := db.WithTx(ctx, c.DB, nil, func(tx *sql.Tx) error {
			// *sql.Tx implements the subset of methods in repositories.DbTx so we can pass it directly
			return s.PerformDataSync(ctx, tx)
		})
		if err != nil {
			c.Log.Error("sync failed", slog.String("sync", name), slog.Any("error", err))
			return err
		}
		c.Log.Info("sync complete", slog.String("sync", name))
	}
	return nil
}
