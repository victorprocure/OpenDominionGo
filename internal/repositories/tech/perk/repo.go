package perk

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log/slog"

	"github.com/victorprocure/opendominiongo/internal/repositories"
)

//go:embed sql/insert_tech_perk.sql
var insertTechPerkSQL string

//go:embed sql/list_tech_perks_by_tech.sql
var listTechPerksByTechSQL string

type Repo struct {
	db  *sql.DB
	log *slog.Logger
}

func NewRepo(db *sql.DB, log *slog.Logger) *Repo { return &Repo{db: db, log: log} }

type CreateArgs struct {
	TechID         int
	TechPerkTypeID int
	Value          string
}

func (r *Repo) CreateContext(ctx context.Context, tx repositories.DbTx, a CreateArgs) (int, error) {
	var id int
	if err := tx.QueryRowContext(ctx, insertTechPerkSQL, a.TechID, a.TechPerkTypeID, a.Value).Scan(&id); err != nil {
		return 0, fmt.Errorf("insert tech_perk: %w", err)
	}
	return id, nil
}

type Row struct {
	ID             int
	TechID         int
	TechPerkTypeID int
	Value          string
}

func (r *Repo) ListByTechContext(ctx context.Context, tx repositories.DbTx, techID, limit, offset int) ([]Row, error) {
	rows, err := tx.QueryContext(ctx, listTechPerksByTechSQL, techID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("list tech_perks: %w", err)
	}
	defer rows.Close()
	var out []Row
	for rows.Next() {
		var x Row
		if err := rows.Scan(&x.ID, &x.TechID, &x.TechPerkTypeID, &x.Value); err != nil {
			return nil, fmt.Errorf("scan tech_perk: %w", err)
		}
		out = append(out, x)
	}
	return out, rows.Err()
}
