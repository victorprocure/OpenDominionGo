package store

import (
	"context"
	"embed"
	"errors"
	"io/fs"
	"sort"
	"strconv"

	"github.com/victorprocure/opendominiongo/helpers"
	"gopkg.in/yaml.v3"
)

//go:embed data/races
var racesFS embed.FS

type RacesSync struct{ storage *Storage }

func NewRacesSync(s *Storage) *RacesSync {
	return &RacesSync{storage: s}
}

func (s *RacesSync) Name() string {
	return "Races"
}

func (s *RacesSync) PerformDataSync(ctx context.Context, tx DbTx) error {
	entries, err := fs.ReadDir(racesFS, "data/races")
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

		err = s.syncRacePerks(name, r, ctx, tx)
		if err != nil {
			return err
		}

		err = s.syncRaceUnits(name, r, ctx, tx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *RacesSync) syncRaceUnits(fn string, r *Race, ctx context.Context, tx DbTx) error {
	units, err := getUnitsForRaceFromYaml(r, fn)
	if err != nil {
		return err
	}

	if len(units) == 0 {
		return errors.New("no units found")
	}

	for _, u := range units {
		err = s.storage.CreateOrUpdateUnitContext(u, ctx, tx)
		if err != nil {
			return err
		}

		for _, perkValue := range u.Perks {
			err := s.storage.CreateOrUpdateUnitPerkTypeContext(perkValue.UnitPerkType, ctx, tx)
			if err != nil {
				return err
			}

			err = s.storage.CreateOrUpdateUnitPerkContext(perkValue, ctx, tx)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *RacesSync) syncRace(r *Race, ctx context.Context, tx DbTx) error {
	err := s.storage.CreateOrUpdateRaceContext(r, ctx, tx)
	if err != nil {
		return err
	}

	return nil
}

func (s *RacesSync) syncRacePerks(fn string, r *Race, ctx context.Context, tx DbTx) error {
	perks, err := getPerksForRaceFromYaml(r, fn)
	if err != nil {
		return err
	}

	if len(perks) == 0 {
		return errors.New("no perks found")
	}

	for _, p := range perks {
		rpt, err := s.storage.GetRacePerkTypeByKeyContext(p.RacePerkType.Key, ctx, tx)
		if err != nil {
			return err
		}

		if rpt == nil {
			err = s.storage.CreateOrUpdateRacePerkTypeContext(p.RacePerkType, ctx, tx)
			if err != nil {
				return err
			}
			rpt = p.RacePerkType
		}

		p.RacePerkType = rpt
		err = s.storage.CreateOrUpdateRacePerkContext(&p, ctx, tx)
		if err != nil {
			return err
		}
	}

	return nil
}

func getUnitsForRaceFromYaml(r *Race, f string) ([]*Unit, error) {
	rf, err := racesFS.ReadFile("data/races/" + f)
	if err != nil {
		return nil, err
	}

	type unit struct {
		Name     string             `yaml:"name"`
		Type     string             `yaml:"type"`
		Cost     map[string]int     `yaml:"cost"`
		Power    map[string]float64 `yaml:"power"`
		Perks    map[string]string  `yaml:"perks"`
		NeedBoat bool               `yaml:"need_boat,omitempty"`
	}
	var tmp struct {
		Units []unit `yaml:"units"`
	}
	if err := yaml.Unmarshal(rf, &tmp); err != nil {
		return nil, err
	}

	ul := len(tmp.Units)
	if ul == 0 {
		return nil, errors.New("no units found")
	}

	units := make([]*Unit, 0, ul)
	for i, u := range tmp.Units {
		var ut string
		slot := i + 1
		if u.Type == "" {
			ut = DefaultUnitTypes[slot]
		} else {
			ut = u.Type
		}

		unit := &Unit{
			Race:         r,
			Slot:         strconv.Itoa(slot),
			Name:         u.Name,
			Type:         ut,
			CostPlatinum: u.Cost["platinum"],
			CostOre:      u.Cost["ore"],
			CostMana:     u.Cost["mana"],
			CostLumber:   u.Cost["lumber"],
			CostGems:     u.Cost["gems"],
			PowerOffense: u.Power["offense"],
			PowerDefense: u.Power["defense"],
			NeedBoat:     u.NeedBoat,
		}

		perks := make([]*UnitPerk, 0, len(u.Perks))
		for k, v := range u.Perks {
			perk := &UnitPerk{
				Unit:         unit,
				UnitPerkType: &UnitPerkType{Key: k},
				Value:        v,
			}
			perks = append(perks, perk)
		}
		unit.Perks = perks
		units = append(units, unit)
	}

	return units, nil
}

func getPerksForRaceFromYaml(r *Race, f string) ([]RacePerk, error) {
	rf, err := racesFS.ReadFile("data/races/" + f)
	if err != nil {
		return nil, err
	}

	var tmp struct {
		Perks map[string]float64 `yaml:"perks"`
	}

	if err := yaml.Unmarshal(rf, &tmp); err != nil {
		return nil, err
	}

	if len(tmp.Perks) == 0 {
		return []RacePerk{}, nil
	}

	keys := make([]string, 0, len(tmp.Perks))
	for k := range tmp.Perks {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	perks := make([]RacePerk, 0, len(tmp.Perks))
	for _, k := range keys {
		v := tmp.Perks[k]
		perk := RacePerk{
			Race:         r,
			RacePerkType: &RacePerkType{Key: k},
			Value:        v,
		}

		perks = append(perks, perk)
	}

	return perks, nil
}

func getRaceFromFile(n string) (*Race, error) {
	b, err := racesFS.ReadFile("data/races/" + n)
	if err != nil {
		return nil, err
	}

	var r Race
	if err = yaml.Unmarshal(b, &r); err != nil {
		return nil, err
	}

	return &r, nil
}
