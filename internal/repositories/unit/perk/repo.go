package perk

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log/slog"

	"github.com/victorprocure/opendominiongo/internal/repositories"
)

//go:embed sql/insert_unit_perk.sql
var insertUnitPerkSQL string

//go:embed sql/list_unit_perks_by_race_and_unit.sql
var listUnitPerksByRaceAndUnitSQL string

type Repo struct {
	db  *sql.DB
	log *slog.Logger
}

func NewUnitPerkRepo(db *sql.DB, log *slog.Logger) *Repo { return &Repo{db: db, log: log} }

type CreateArgs struct {
	RaceID         int
	UnitSlot       string
	UnitPerkTypeID int
	Value          string
}

func (r *Repo) CreateContext(ctx context.Context, tx repositories.DbTx, a CreateArgs) (int, error) {
	var id int
	if err := tx.QueryRowContext(ctx, insertUnitPerkSQL, a.RaceID, a.UnitSlot, a.UnitPerkTypeID, a.Value).Scan(&id); err != nil {
		return 0, fmt.Errorf("insert unit_perk: %w", err)
	}
	return id, nil
}

type Row struct {
	ID             int
	RaceID         int
	UnitSlot       string
	UnitPerkTypeID int
	Value          string
}

func (r *Repo) ListByRaceAndUnitContext(ctx context.Context, tx repositories.DbTx, raceID int, unitSlot string, limit, offset int) ([]Row, error) {
	rows, err := tx.QueryContext(ctx, listUnitPerksByRaceAndUnitSQL, raceID, unitSlot, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("list unit_perks: %w", err)
	}
	defer rows.Close()
	var out []Row
	for rows.Next() {
		var x Row
		if err := rows.Scan(&x.ID, &x.RaceID, &x.UnitSlot, &x.UnitPerkTypeID, &x.Value); err != nil {
			return nil, fmt.Errorf("scan unit_perk: %w", err)
		}
		out = append(out, x)
	}
	return out, rows.Err()
}
