package tech

import (
	"context"
	"database/sql"
	_ "embed"
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/victorprocure/opendominiongo/internal/repositories"
)

type Repo struct {
	db  *sql.DB
	log *slog.Logger
}

func NewRepo(db *sql.DB, log *slog.Logger) *Repo {
	return &Repo{db: db, log: log}
}

//go:embed sql/upsert_tech_with_perks.sql
var upsertTechWithPerksSQL string

// UpsertArgs captures the normalized inputs needed to upsert a single tech and its perks.
type UpsertArgs struct {
	Key           string
	Name          string
	Prerequisites string
	Active        bool
	Version       int
	X             int
	Y             int
	Perks         map[string]string
}

// UpsertFromSyncContext upserts one tech row and its child perks/types using a single SQL statement.
// Returns the tech id.
func (r *Repo) UpsertFromSyncContext(ctx context.Context, tx repositories.DbTx, a UpsertArgs) (int, error) {
	var perksJSON []byte
	if len(a.Perks) > 0 {
		b, err := json.Marshal(a.Perks)
		if err != nil {
			return 0, fmt.Errorf("marshal tech perks: %w", err)
		}
		perksJSON = b
	}

	var techID int
	err := tx.QueryRowContext(
		ctx,
		upsertTechWithPerksSQL,
		a.Key,
		a.Name,
		a.Prerequisites,
		a.Active,
		a.Version,
		a.X,
		a.Y,
		perksJSON,
	).Scan(&techID)
	if err != nil {
		return 0, fmt.Errorf("upsert tech %q: %w", a.Key, err)
	}
	return techID, nil
}
