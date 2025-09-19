package perk

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log/slog"

	"github.com/victorprocure/opendominiongo/internal/repositories"
)

//go:embed sql/insert_spell_perk.sql
var insertSpellPerkSQL string

//go:embed sql/list_spell_perks_by_spell.sql
var listSpellPerksBySpellSQL string

type Repo struct {
	db  *sql.DB
	log *slog.Logger
}

func NewRepo(db *sql.DB, log *slog.Logger) *Repo { return &Repo{db: db, log: log} }

type CreateArgs struct {
	SpellID         int
	SpellPerkTypeID int
	Value           string
}

func (r *Repo) CreateContext(ctx context.Context, tx repositories.DbTx, a CreateArgs) (int, error) {
	var id int
	if err := tx.QueryRowContext(ctx, insertSpellPerkSQL, a.SpellID, a.SpellPerkTypeID, a.Value).Scan(&id); err != nil {
		return 0, fmt.Errorf("insert spell_perk: %w", err)
	}
	return id, nil
}

type Row struct {
	ID              int
	SpellID         int
	SpellPerkTypeID int
	Value           string
}

func (r *Repo) ListBySpellContext(ctx context.Context, tx repositories.DbTx, spellID, limit, offset int) ([]Row, error) {
	rows, err := tx.QueryContext(ctx, listSpellPerksBySpellSQL, spellID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("list spell_perks: %w", err)
	}
	defer rows.Close()
	var out []Row
	for rows.Next() {
		var x Row
		if err := rows.Scan(&x.ID, &x.SpellID, &x.SpellPerkTypeID, &x.Value); err != nil {
			return nil, fmt.Errorf("scan spell_perk: %w", err)
		}
		out = append(out, x)
	}
	return out, rows.Err()
}
