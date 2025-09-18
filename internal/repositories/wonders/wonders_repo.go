package wonder

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log/slog"

	"github.com/victorprocure/opendominiongo/internal/helpers"
	"github.com/victorprocure/opendominiongo/internal/repositories"
)

//go:embed sql/upsert_wonder_sync.sql
var upsertWonderSyncSQL string

type Repo struct {
	db  *sql.DB
	log *slog.Logger
}

func NewWondersRepo(db *sql.DB, log *slog.Logger) *Repo {
	return &Repo{db: db, log: log}
}

// UpsertArgs is the normalized contract for upserting a wonder.
type UpsertArgs struct {
	Key    string
	Name   string
	Power  int
	Active bool
	Perks  map[string]string
}

func (r *Repo) UpsertFromSyncContext(ctx context.Context, tx repositories.DbTx, a UpsertArgs) error {
	var perksJSON []byte
	if len(a.Perks) > 0 {
		var err error
		perksJSON, err = helpers.MarshalPerksAsJSONArrayFromMap(a.Perks)
		if err != nil {
			return fmt.Errorf("marshal wonder perks: %w", err)
		}
		r.log.Info("upsert wonder input", "key", a.Key, "perksJSON", string(perksJSON))
	}

	var newID int
	err := tx.QueryRowContext(ctx, upsertWonderSyncSQL,
		a.Key,
		a.Name,
		a.Power,
		a.Active,
		perksJSON,
	).Scan(&newID)
	if err != nil {
		return fmt.Errorf("upsert wonder: %w", err)
	}
	return nil
}
