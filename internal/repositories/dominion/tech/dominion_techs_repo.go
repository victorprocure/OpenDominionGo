package tech

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log/slog"

	"github.com/victorprocure/opendominiongo/internal/repositories"
)

//go:embed sql/insert_dominion_tech.sql
var insertDominionTechSQL string

//go:embed sql/list_dominion_techs_by_dominion.sql
var listDominionTechsByDominionSQL string

type Repo struct {
	db  *sql.DB
	log *slog.Logger
}

func NewDominionTechsRepo(db *sql.DB, log *slog.Logger) *Repo { return &Repo{db: db, log: log} }

type CreateArgs struct {
	DominionID int
	TechID     int
}

func (r *Repo) CreateContext(ctx context.Context, tx repositories.DbTx, a CreateArgs) (int, error) {
	var id int
	if err := tx.QueryRowContext(ctx, insertDominionTechSQL, a.DominionID, a.TechID).Scan(&id); err != nil {
		return 0, fmt.Errorf("insert dominion_tech: %w", err)
	}
	return id, nil
}

type Row struct {
	ID         int
	DominionID int
	TechID     int
}

func (r *Repo) ListByDominionContext(ctx context.Context, tx repositories.DbTx, dominionID int) ([]Row, error) {
	rows, err := tx.QueryContext(ctx, listDominionTechsByDominionSQL, dominionID)
	if err != nil {
		return nil, fmt.Errorf("list dominion_techs: %w", err)
	}
	defer rows.Close()
	var out []Row
	for rows.Next() {
		var t Row
		if err := rows.Scan(&t.ID, &t.DominionID, &t.TechID); err != nil {
			return nil, fmt.Errorf("scan dominion_tech: %w", err)
		}
		out = append(out, t)
	}
	return out, rows.Err()
}
