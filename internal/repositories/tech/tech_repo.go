package tech

import (
	"context"
	"database/sql"
	_ "embed"
	"encoding/json"
	"fmt"
	"log/slog"
	"strings"

	"github.com/victorprocure/opendominiongo/internal/dto"
	"github.com/victorprocure/opendominiongo/internal/repositories"
)

type TechRepo struct {
	db  *sql.DB
	log *slog.Logger
}

func NewTechRepo(db *sql.DB, log *slog.Logger) *TechRepo {
	return &TechRepo{db: db, log: log}
}

//go:embed sql/upsert_tech_with_perks.sql
var upsertTechWithPerksSQL string

// TechUpsertArgs captures the normalized inputs needed to upsert a single tech and its perks.
type TechUpsertArgs struct {
	Key           string
	Name          string
	Prerequisites []string
	Active        bool
	Version       int
	X             int
	Y             int
	Perks         dto.KeyValues
}

// CreateOrUpdateTechSyncContext upserts one tech row and its child perks/types using a single SQL statement.
// Returns the tech id.
func (r *TechRepo) CreateOrUpdateTechSyncContext(ctx context.Context, tx repositories.DbTx, a TechUpsertArgs) (int, error) {
	// Marshal perks to JSON (object mapping preserving order via dto.KeyValues MarshalJSON)
	var perksJSON []byte
	if len(a.Perks) > 0 {
		b, err := json.Marshal(a.Perks)
		if err != nil {
			return 0, fmt.Errorf("marshal tech perks: %w", err)
		}
		perksJSON = b
	}

	// Join prerequisites as comma-delimited text (adjust if DB expects text[])
	prereq := strings.Join(a.Prerequisites, ",")

	var techID int
	err := tx.QueryRowContext(
		ctx,
		upsertTechWithPerksSQL,
		a.Key,
		a.Name,
		prereq,
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
