package races

import (
	"context"
	"database/sql"
	_ "embed"
	"encoding/json"
	"fmt"
	"log/slog"

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

type UnitUpsertArg struct {
	Name         string            `json:"name"`
	Type         string            `json:"type"`
	NeedBoat     bool              `json:"need_boat"`
	CostPlatinum int               `json:"cost_platinum"`
	CostOre      int               `json:"cost_ore"`
	CostLumber   int               `json:"cost_lumber"`
	CostGems     int               `json:"cost_gems"`
	CostMana     int               `json:"cost_mana"`
	PowerOffense int               `json:"power_offense"`
	PowerDefense int               `json:"power_defense"`
	Perks        map[string]string `json:"perks"`
}

type RaceUpsertArgs struct {
	Key                 string
	Name                string
	Alignment           string
	Description         string
	AttackerDifficulty  int
	ExplorerDifficulty  int
	ConverterDifficulty int
	OverallDifficulty   int
	HomeLandType        string
	Playable            bool
	Perks               map[string]string
	Units               []UnitUpsertArg
}

func (r *RacesRepo) UpsertRaceFromSyncContext(ctx context.Context, tx repositories.DbTx, a RaceUpsertArgs) (int, error) {
	var rpJSON []byte
	if len(a.Perks) > 0 {
		b, err := json.Marshal(a.Perks)
		if err != nil {
			return 0, fmt.Errorf("marshal race perks: %w", err)
		}

		rpJSON = b
	}

	us := a.Units

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
		a.Key, a.Name, a.Alignment, a.Description,
		a.AttackerDifficulty, a.ExplorerDifficulty, a.ConverterDifficulty,
		a.OverallDifficulty, a.HomeLandType, a.Playable,
		rpJSON, usJSON,
	).Scan(&raceID)
	if err != nil {
		return 0, fmt.Errorf("unable to upsert race with units: %w", err)
	}
	return raceID, nil
}
