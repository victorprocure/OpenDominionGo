package perk_type

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log/slog"

	"github.com/victorprocure/opendominiongo/internal/repositories"
)

//go:embed sql/list_unit_perk_types.sql
var listUnitPerkTypesSQL string

type Repo struct {
	db  *sql.DB
	log *slog.Logger
}

func NewUnitPerkTypeRepo(db *sql.DB, log *slog.Logger) *Repo { return &Repo{db: db, log: log} }

type Row struct {
	ID   int
	Name string
}

func (r *Repo) ListContext(ctx context.Context, tx repositories.DbTx) ([]Row, error) {
	rows, err := tx.QueryContext(ctx, listUnitPerkTypesSQL)
	if err != nil {
		return nil, fmt.Errorf("list unit_perk_types: %w", err)
	}
	defer rows.Close()
	var out []Row
	for rows.Next() {
		var t Row
		if err := rows.Scan(&t.ID, &t.Name); err != nil {
			return nil, fmt.Errorf("scan unit_perk_type: %w", err)
		}
		out = append(out, t)
	}
	return out, rows.Err()
}
