package races

import (
	"context"
	"database/sql"
	_ "embed"
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/victorprocure/opendominiongo/internal/dto"
	"github.com/victorprocure/opendominiongo/internal/repositories"
)

//go:embed sql/upsert_race_for_sync.sql
var upsertRaceFromYamlSQL string

type RacesRepo struct {
	db  *sql.DB
	log *slog.Logger
}

func NewRacesRepository(db *sql.DB, log *slog.Logger) *RacesRepo {
	return &RacesRepo{db: db, log: log}
}

func (r *RacesRepo) UpsertRaceFromYamlContext(race *dto.RaceYaml, ctx context.Context, tx repositories.DbTx) (int, error) {
	var rpJSON []byte
	if len(race.Perks) > 0 {
		b, err := json.Marshal(race.Perks)
		if err != nil {
			return 0, fmt.Errorf("marshal race perks: %w", err)
		}

		rpJSON = b
	}

	us := make([]dto.UnitSyncJSON, 0, len(race.Units))
	for _, u := range race.Units {
		us = append(us, dto.UnitSyncJSON{
			Name:         u.Name,
			Type:         u.Type,
			NeedBoat:     u.NeedBoat,
			CostPlatinum: u.Cost.Platinum,
			CostOre:      u.Cost.Ore,
			CostLumber:   u.Cost.Lumber,
			CostGems:     u.Cost.Gems,
			CostMana:     u.Cost.Mana,
			PowerOffense: u.Power.Offense,
			PowerDefense: u.Power.Defense,
			Perks:        u.Perks,
		})
	}

	var usJSON []byte
	if len(us) > 0 {
		b, err := json.Marshal(us)
		if err != nil {
			return 0, fmt.Errorf("marshal units: %w", err)
		}
		usJSON = b
	}

	var raceID int
	err := tx.QueryRowContext(ctx, upsertRaceFromYamlSQL,
		race.Key, race.Name, race.Alignment, race.Description,
		race.AttackerDifficulty, race.ExplorerDifficulty, race.ConverterDifficulty,
		race.OverallDifficulty, race.HomeLandType, race.Playable.OrDefault(),
		rpJSON, usJSON,
	).Scan(&raceID)
	if err != nil {
		return 0, fmt.Errorf("unable to upsert race with units: %w", err)
	}
	return raceID, nil
}
