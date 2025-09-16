package store

import (
	"context"
	_ "embed"
	"fmt"

	"github.com/victorprocure/opendominiongo/helpers"
	"gopkg.in/yaml.v3"
)

const WondersDir = "data/wonders.yml"

//go:embed data/wonders.yml
var wonderImportFile []byte

type wondersYaml struct {
	Name   string                  `yaml:"name"`
	Perks  map[string]string       `yaml:"perks"`
	Power  int                     `yaml:"power"`
	Active helpers.BoolDefaultTrue `yaml:"active,omitempty"`
}

type WondersSync struct {
	storage *Storage
}

func NewWondersSync(s *Storage) *WondersSync {
	return &WondersSync{storage: s}
}

func (s *WondersSync) Name() string {
	return "Wonders"
}

func (s *WondersSync) PerformDataSync(ctx context.Context, tx DbTx) error {
	w, err := s.getWondersFromYaml()
	if err != nil {
		return fmt.Errorf("unable to get wonders from yaml: %w", err)
	}

	if len(w) == 0 {
		return fmt.Errorf("no wonders to add: %v", w)
	}

	err = s.syncWonders(w, ctx, tx)
	if err != nil {
		return fmt.Errorf("unable to sync wonders: %w", err)
	}

	return nil
}

func (s *WondersSync) syncWonders(wl []*Wonder, ctx context.Context, tx DbTx) error {
	for _, w := range wl {
		err := s.storage.CreateOrUpdateWonderContext(w, ctx, tx)
		if err != nil {
			return fmt.Errorf("unable to create or update wonder: %s, error: %w", w.Key, err)
		}

		for _, p := range w.Perks {
			p.Wonder = w
			err := s.storage.CreateOrUpdateWonderPerkTypeContext(p.WonderPerkType, ctx, tx)
			if err != nil {
				return fmt.Errorf("unable to create or update wonder perk type: %s for wonder: %s, error: %w", p.WonderPerkType.Key, w.Key, err)
			}

			err = s.storage.CreateOrUpdateWonderPerkContext(p, ctx, tx)
			if err != nil {
				return fmt.Errorf("unable to create or update wonder perk: %s for wonder: %s, error: %w", p.WonderPerkType.Key, w.Key, err)
			}
		}
	}

	return nil
}

func (s *WondersSync) getWondersFromYaml() ([]*Wonder, error) {
	var byKey map[string]wondersYaml
	if err := yaml.Unmarshal(wonderImportFile, &byKey); err != nil {
		return nil, fmt.Errorf("unable to unmarshal: %v. error: %w", wonderImportFile, err)
	}

	wl := make([]*Wonder, 0, len(byKey))
	for key, wonder := range byKey {

		w := &Wonder{
			Key:    key,
			Name:   wonder.Name,
			Power:  wonder.Power,
			Active: wonder.Active.OrDefault(),
		}

		w.Perks = make([]*WonderPerk, 0, len(wonder.Perks))
		for perkKey, perkValue := range wonder.Perks {
			w.Perks = append(w.Perks,
				&WonderPerk{
					WonderPerkType: &WonderPerkType{Key: perkKey},
					Value:          perkValue,
				})
		}

		wl = append(wl, w)
	}

	return wl, nil
}
