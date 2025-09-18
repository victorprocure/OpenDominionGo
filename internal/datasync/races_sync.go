package datasync

import (
	"context"
	"database/sql"
	"embed"
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/victorprocure/opendominiongo/internal/domain"
	"github.com/victorprocure/opendominiongo/internal/dto"
	"github.com/victorprocure/opendominiongo/internal/encoding/yamlutil"
	"github.com/victorprocure/opendominiongo/internal/helpers"
	"github.com/victorprocure/opendominiongo/internal/repositories"
	racerepo "github.com/victorprocure/opendominiongo/internal/repositories/races"
	"gopkg.in/yaml.v3"
)

const racesDir = "import_data/races"

//go:embed import_data/races
var racesFS embed.FS

type RacesSync struct {
	db  *racerepo.Repo
	log *slog.Logger
}

func NewRacesSync(db *sql.DB, log *slog.Logger) *RacesSync {
	return &RacesSync{db: racerepo.NewRacesRepository(db, log), log: log}
}

func (s *RacesSync) Name() string {
	return "Races"
}

func (s *RacesSync) PerformDataSync(ctx context.Context, tx repositories.DbTx) error {
	entries, err := yamlutil.GetYmlImportFiles(racesFS, racesDir)
	if err != nil {
		return err
	}

	for _, fn := range entries {
		r, err := getRaceFromFile(fn)
		if err != nil {
			return fmt.Errorf("read race file %s: %w", fn, err)
		}

		// Log the race key we're about to upsert for visibility
		s.log.Info("upserting race", slog.String("file", fn), slog.String("race_key", r.Key))

		if err := s.syncRace(ctx, tx, r); err != nil {
			return fmt.Errorf("sync race %s: %w", r.Key, err)
		}
	}

	return nil
}

func (s *RacesSync) syncRace(ctx context.Context, tx repositories.DbTx, r *dto.RaceYaml) error {
	// map DTO -> repo wrapper
	perks := helpers.PerksToMap(r.Perks)
	units := make([]racerepo.UnitUpsertArg, 0, len(r.Units))
	for _, u := range r.Units {
		units = append(units, racerepo.UnitUpsertArg{
			Slot:         u.Slot,
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
			Perks:        helpers.PerksToMap(u.Perks),
		})
	}
	// optional description handling
	desc := ""
	if r.Description != nil {
		desc = *r.Description
	}
	// Marshal and log the serialized perks/units that will be passed to SQL
	if b, err := json.Marshal(perks); err == nil {
		s.log.Info("race perks json", slog.String("race", r.Key), slog.String("perks", string(b)))
	}
	if b, err := json.Marshal(units); err == nil {
		s.log.Info("race units json", slog.String("race", r.Key), slog.String("units", string(b)))
	}

	_, err := s.db.UpsertFromSyncContext(ctx, tx, racerepo.UpsertArgs{
		Key:                 r.Key,
		Name:                r.Name,
		Alignment:           r.Alignment,
		Description:         desc,
		AttackerDifficulty:  r.AttackerDifficulty,
		ExplorerDifficulty:  r.ExplorerDifficulty,
		ConverterDifficulty: r.ConverterDifficulty,
		OverallDifficulty:   r.OverallDifficulty,
		HomeLandType:        r.HomeLandType,
		Playable:            r.Playable.OrDefault(),
		Perks:               perks,
		Units:               units,
	})
	if err != nil {
		return err
	}

	return nil
}

func getRaceFromFile(f string) (*dto.RaceYaml, error) {
	b, err := racesFS.ReadFile(f)
	if err != nil {
		return nil, err
	}

	var r dto.RaceYaml
	if err = yaml.Unmarshal(b, &r); err != nil {
		return nil, err
	}

	assignTypesToUnits(&r)

	return &r, nil
}

func assignTypesToUnits(r *dto.RaceYaml) {
	for i := range r.Units {
		position := i + 1
		slot := fmt.Sprintf("%d", position)

		r.Units[i].Slot = slot
		if r.Units[i].Type == "" {
			r.Units[i].Type = domain.DefaultUnitTypes[position]
		}
	}
}
