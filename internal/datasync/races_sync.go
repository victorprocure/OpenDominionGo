package datasync

import (
	"context"
	"database/sql"
	"embed"
	"io/fs"
	"log/slog"

	"github.com/victorprocure/opendominiongo/helpers"
	"github.com/victorprocure/opendominiongo/internal/domain"
	"github.com/victorprocure/opendominiongo/internal/dto"
	"github.com/victorprocure/opendominiongo/internal/repositories"
	"github.com/victorprocure/opendominiongo/internal/repositories/races"
	"gopkg.in/yaml.v3"
)

const racesDir = "import_data/races"

//go:embed import_data/races
var racesFS embed.FS

type RacesSync struct{ db *races.RacesRepo }

func NewRacesSync(db *sql.DB, log *slog.Logger) *RacesSync {
	return &RacesSync{db: races.NewRacesRepository(db, log)}
}

func (s *RacesSync) Name() string {
	return "Races"
}

func (s *RacesSync) PerformDataSync(ctx context.Context, tx repositories.DbTx) error {
	entries, err := fs.ReadDir(racesFS, racesDir)
	if err != nil {
		return err
	}

	for _, e := range entries {
		name, valid, _ := helpers.IsValidYamlFileName(e)
		if !valid {
			continue
		}

		r, err := getRaceFromFile(name)
		if err != nil {
			continue
		}

		err = s.syncRace(r, ctx, tx)
		if err != nil {
			continue
		}
	}

	return nil
}

func (s *RacesSync) syncRace(r *dto.RaceYaml, ctx context.Context, tx repositories.DbTx) error {
	// map DTO -> repo wrapper
	perks := make(map[string]string, len(r.Perks))
	for _, kv := range r.Perks {
		perks[kv.Key] = kv.Value
	}
	units := make([]races.UnitUpsertArg, 0, len(r.Units))
	for _, u := range r.Units {
		units = append(units, races.UnitUpsertArg{
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
			Perks:        toPerkMap(u.Perks),
		})
	}
	// optional description handling
	desc := ""
	if r.Description != nil {
		desc = *r.Description
	}
	_, err := s.db.UpsertRaceFromSyncContext(ctx, tx, races.RaceUpsertArgs{
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

func toPerkMap(kv dto.KeyValues) map[string]string {
	if len(kv) == 0 {
		return nil
	}
	m := make(map[string]string, len(kv))
	for _, p := range kv {
		m[p.Key] = p.Value
	}
	return m
}

func getRaceFromFile(n string) (*dto.RaceYaml, error) {
	b, err := racesFS.ReadFile(racesDir + "/" + n)
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
	for i, v := range r.Units {
		position := i + 1
		if v.Type == "" {
			v.Type = domain.DefaultUnitTypes[position]
		}
	}
}
