package spell

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log/slog"

	"github.com/victorprocure/opendominiongo/internal/repositories"
)

//go:embed sql/upsert_dominion_spell.sql
var upsertDominionSpellSQL string

//go:embed sql/list_dominion_spells_by_dominion.sql
var listDominionSpellsByDominionSQL string

type Repo struct {
	db  *sql.DB
	log *slog.Logger
}

func NewDominionSpellsRepo(db *sql.DB, log *slog.Logger) *Repo { return &Repo{db: db, log: log} }

type UpsertArgs struct {
	DominionID       int
	SpellID          int
	Duration         int
	CastByDominionID *int
}

func (r *Repo) UpsertContext(ctx context.Context, tx repositories.DbTx, a UpsertArgs) error {
	if _, err := tx.ExecContext(ctx, upsertDominionSpellSQL, a.DominionID, a.Duration, a.CastByDominionID, a.SpellID); err != nil {
		return fmt.Errorf("upsert dominion_spell: %w", err)
	}
	return nil
}

type Row struct {
	DominionID       int
	Duration         int
	CastByDominionID *int
	SpellID          int
}

func (r *Repo) ListByDominionContext(ctx context.Context, tx repositories.DbTx, dominionID int) ([]Row, error) {
	rows, err := tx.QueryContext(ctx, listDominionSpellsByDominionSQL, dominionID)
	if err != nil {
		return nil, fmt.Errorf("list dominion_spells: %w", err)
	}
	defer rows.Close()
	var out []Row
	for rows.Next() {
		var s Row
		if err := rows.Scan(&s.DominionID, &s.Duration, &s.CastByDominionID, &s.SpellID); err != nil {
			return nil, fmt.Errorf("scan dominion_spell: %w", err)
		}
		out = append(out, s)
	}
	return out, rows.Err()
}
