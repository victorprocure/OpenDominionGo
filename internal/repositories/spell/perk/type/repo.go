package perktype

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log/slog"

	"github.com/victorprocure/opendominiongo/internal/repositories"
)

//go:embed sql/list_spell_perk_types.sql
var listSpellPerkTypesSQL string

type Repo struct {
	db  *sql.DB
	log *slog.Logger
}

func NewSpellPerkTypeRepo(db *sql.DB, log *slog.Logger) *Repo { return &Repo{db: db, log: log} }

type Row struct {
	ID   int
	Name string
}

func (r *Repo) ListContext(ctx context.Context, tx repositories.DbTx) ([]Row, error) {
	rows, err := tx.QueryContext(ctx, listSpellPerkTypesSQL)
	if err != nil {
		return nil, fmt.Errorf("list spell_perk_types: %w", err)
	}
	defer rows.Close()
	var out []Row
	for rows.Next() {
		var t Row
		if err := rows.Scan(&t.ID, &t.Name); err != nil {
			return nil, fmt.Errorf("scan spell_perk_type: %w", err)
		}
		out = append(out, t)
	}
	return out, rows.Err()
}
