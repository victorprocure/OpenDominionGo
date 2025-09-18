package hero

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log/slog"

	"github.com/victorprocure/opendominiongo/internal/repositories"
)

//go:embed sql/insert_hero.sql
var insertHeroSQL string

//go:embed sql/list_heroes_by_dominion.sql
var listHeroesByDominionSQL string

//go:embed sql/update_hero.sql
var updateHeroSQL string

type Repo struct {
	db  *sql.DB
	log *slog.Logger
}

func NewHeroRepo(db *sql.DB, log *slog.Logger) *Repo { return &Repo{db: db, log: log} }

type CreateArgs struct {
	DominionID int
	Name       *string
	Class      *string
}

func (r *Repo) CreateContext(ctx context.Context, tx repositories.DbTx, a CreateArgs) (int, error) {
	var id int
	if err := tx.QueryRowContext(ctx, insertHeroSQL, a.DominionID, a.Name, a.Class).Scan(&id); err != nil {
		return 0, fmt.Errorf("insert hero: %w", err)
	}
	return id, nil
}

type Row struct {
	ID           int
	DominionID   int
	Name         *string
	Class        *string
	Experience   float64
	CombatRating int
}

func (r *Repo) ListByDominionContext(ctx context.Context, tx repositories.DbTx, dominionID int, limit, offset int) ([]Row, error) {
	rows, err := tx.QueryContext(ctx, listHeroesByDominionSQL, dominionID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("list heroes: %w", err)
	}
	defer rows.Close()
	var out []Row
	for rows.Next() {
		var h Row
		if err := rows.Scan(&h.ID, &h.DominionID, &h.Name, &h.Class, &h.Experience, &h.CombatRating); err != nil {
			return nil, fmt.Errorf("scan hero: %w", err)
		}
		out = append(out, h)
	}
	return out, rows.Err()
}

type UpdateArgs struct {
	ID         int
	Name       *string
	Class      *string
	Experience *float64
}

func (r *Repo) UpdateContext(ctx context.Context, tx repositories.DbTx, a UpdateArgs) error {
	if _, err := tx.ExecContext(ctx, updateHeroSQL, a.ID, a.Name, a.Class, a.Experience); err != nil {
		return fmt.Errorf("update hero: %w", err)
	}
	return nil
}
