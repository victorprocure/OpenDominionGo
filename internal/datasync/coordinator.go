package datasync

import (
	"context"
	"database/sql"
	"log/slog"

	"github.com/victorprocure/opendominiongo/internal/db"
)

type SyncCoordinator interface {
	// RunAll executes all syncers in sequence, each within its own transaction.
	// If any syncer fails, it stops and returns the error.
	RunAll(ctx context.Context) error
}

// SyncCoordinator runs a set of Syncers inside DB transactions.
type syncCoordinator struct {
	DB      *sql.DB
	Log     *slog.Logger
	syncers []Syncer
}

// NewSyncCoordinator constructs a coordinator with the DB and logger.
func NewSyncCoordinator(db *sql.DB, log *slog.Logger) SyncCoordinator {
	return &syncCoordinator{DB: db, Log: log}
}

// RunAll executes the provided syncers sequentially. Each syncer is executed inside its own transaction.
// If any syncer returns an error, RunAll stops and returns that error.
func (c *syncCoordinator) RunAll(ctx context.Context) error {
	c.createAll()
	for _, s := range c.syncers {
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

func (c *syncCoordinator) createAll() {
	syncers := []Syncer{
		NewTechSync(c.DB, c.Log),
		NewRacesSync(c.DB, c.Log),
		NewSpellsSync(c.DB, c.Log),
		NewWondersSync(c.DB, c.Log),
		NewHeroesSync(c.DB, c.Log),
	}

	c.syncers = syncers
}
