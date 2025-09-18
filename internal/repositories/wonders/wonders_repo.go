package wonders

import (
	"context"
	"database/sql"
	_ "embed"
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/victorprocure/opendominiongo/internal/dto"
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

func (r *WondersRepo) UpsertWonderSyncContext(wonder *dto.WondersYaml, ctx context.Context, tx repositories.DbTx) error {
	var perksJSON []byte
	if len(wonder.Perks) > 0 {
		b, err := json.Marshal(wonder.Perks)
		if err != nil {
			return fmt.Errorf("marshal wonder perks: %w", err)
		}
		perksJSON = b
	}

	var newId int
	err := tx.QueryRowContext(ctx, upsertWonderSyncSQL,
		wonder.Key,
		wonder.Name,
		wonder.Power,
		wonder.Active,
		perksJSON,
	).Scan(&newId)
	if err != nil {
		return fmt.Errorf("upsert wonder: %w", err)
	}
	return nil
}
