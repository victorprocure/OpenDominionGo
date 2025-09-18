package wonders

import (
	"context"
	"database/sql"
	_ "embed"
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/victorprocure/opendominiongo/internal/repositories"
)

//go:embed sql/upsert_wonder_sync.sql
var upsertWonderSyncSQL string

type WondersRepo struct {
	db  *sql.DB
	log *slog.Logger
}

func NewWondersRepo(db *sql.DB, log *slog.Logger) *WondersRepo {
	return &WondersRepo{db: db, log: log}
}

// WonderUpsertArgs is the normalized contract for upserting a wonder.
type WonderUpsertArgs struct {
	Key    string
	Name   string
	Power  int
	Active bool
	Perks  map[string]string
}

func (r *WondersRepo) UpsertWonderFromSyncContext(ctx context.Context, tx repositories.DbTx, a WonderUpsertArgs) error {
	var perksJSON []byte
	if len(a.Perks) > 0 {
		b, err := json.Marshal(a.Perks)
		if err != nil {
			return fmt.Errorf("marshal wonder perks: %w", err)
		}
		perksJSON = b
	}

	var newId int
	err := tx.QueryRowContext(ctx, upsertWonderSyncSQL,
		a.Key,
		a.Name,
		a.Power,
		a.Active,
		perksJSON,
	).Scan(&newId)
	if err != nil {
		return fmt.Errorf("upsert wonder: %w", err)
	}
	return nil
}
