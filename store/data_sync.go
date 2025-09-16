package store

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
)

type DataSync interface {
	Name() string
	PerformDataSync(ctx context.Context, db DbTx) error
}

type SyncCoordinator struct {
	storage *Storage
	syncers []DataSync
}

func NewSyncCoordinator(s *Storage, syncers ...DataSync) *SyncCoordinator {
	return &SyncCoordinator{storage: s, syncers: syncers}
}

func (c *SyncCoordinator) RunDataSync(ctx context.Context) error {
	tx, err := c.storage.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	for _, s := range c.syncers {
		c.storage.Log.Info("Starting Data Sync", slog.String("DataSync", s.Name()))
		if err = s.PerformDataSync(ctx, tx); err != nil {
			return fmt.Errorf("data Sync '%s' failed: %w", s.Name(), err)
		}
		c.storage.Log.Info("Completed Data Sync", slog.String("DataSync", s.Name()))
	}

	return tx.Commit();
}