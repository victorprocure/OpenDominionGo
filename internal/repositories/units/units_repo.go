package unit

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log/slog"

	"github.com/victorprocure/opendominiongo/internal/repositories"
)

//go:embed sql/list_units_by_race_id.sql
var listUnitsByRaceIDSQL string

//go:embed sql/get_unit_by_race_and_slot.sql
var getUnitByRaceAndSlotSQL string

type Repo struct {
	db  *sql.DB
	log *slog.Logger
}

func NewUnitsRepo(db *sql.DB, log *slog.Logger) *Repo {
	return &Repo{db: db, log: log}
}

type Row struct {
	ID           int
	RaceID       int
	Name         string
	Slot         string
	Type         string
	NeedBoat     bool
	CostPlatinum int
	CostOre      int
	CostLumber   int
	CostGems     int
	CostMana     int
	PowerOffense float64
	PowerDefense float64
}

func (r *Repo) ListByRaceIDContext(ctx context.Context, tx repositories.DbTx, raceID int) ([]Row, error) {
	rows, err := tx.QueryContext(ctx, listUnitsByRaceIDSQL, raceID)
	if err != nil {
		return nil, fmt.Errorf("list units by race: %w", err)
	}
	defer rows.Close()
	var out []Row
	for rows.Next() {
		var ur Row
		if err := rows.Scan(&ur.ID, &ur.RaceID, &ur.Name, &ur.Slot, &ur.Type, &ur.NeedBoat,
			&ur.CostPlatinum, &ur.CostOre, &ur.CostLumber, &ur.CostGems, &ur.CostMana,
			&ur.PowerOffense, &ur.PowerDefense); err != nil {
			return nil, fmt.Errorf("scan unit row: %w", err)
		}
		out = append(out, ur)
	}
	return out, rows.Err()
}

func (r *Repo) GetByRaceAndSlotContext(ctx context.Context, tx repositories.DbTx, raceID int, slot string) (*Row, error) {
	var ur Row
	if err := tx.QueryRowContext(ctx, getUnitByRaceAndSlotSQL, raceID, slot).Scan(&ur.ID, &ur.RaceID, &ur.Name, &ur.Slot, &ur.Type, &ur.NeedBoat,
		&ur.CostPlatinum, &ur.CostOre, &ur.CostLumber, &ur.CostGems, &ur.CostMana,
		&ur.PowerOffense, &ur.PowerDefense); err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, fmt.Errorf("get unit by race+slot: %w", err)
	}
	return &ur, nil
}
