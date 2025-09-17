package store

import (
	"context"
	_ "embed"
	"fmt"

	"github.com/victorprocure/opendominiongo/helpers"
	"gopkg.in/yaml.v3"
)

const heroUpgradesFile = "data/heroes.yml"

//go:embed data/heroes.yml
var heroUpgradeImportFile []byte

type heroUpgradeYaml struct {
	Name    string                      `yaml:"name"`
	Level   int                         `yaml:"level"`
	Type    string                      `yaml:"type"`
	Icon    string                      `yaml:"icon"`
	Perks   map[string]string           `yaml:"perks"`
	Active  helpers.BoolDefaultTrue     `yaml:"active,omitempty"`
	Classes helpers.CommaDelimitedArray `yaml:"classes,omitempty"`
}

type HeroUpgradeSync struct {
	storage *Storage
}

func NewHeroUpgradeSync(s *Storage) *HeroUpgradeSync {
	return &HeroUpgradeSync{storage: s}
}

func (s *HeroUpgradeSync) Name() string {
	return "Hero Upgrades"
}

func (s *HeroUpgradeSync) PerformDataSync(ctx context.Context, tx DbTx) error {
	hu, err := s.getHeroUpgradesFromYaml()
	if err != nil {
		return fmt.Errorf("unable to get hero upgrades from yaml: %w", err)
	}

	if len(hu) == 0 {
		return fmt.Errorf("no hero upgrades to add: %v", hu)
	}

	err = s.syncHeroUpgrades(hu, ctx, tx)
	if err != nil {
		return fmt.Errorf("unable to sync hero upgrades: %w", err)
	}

	return nil
}

func (s *HeroUpgradeSync) syncHeroUpgrades(hu []*HeroUpgrade, ctx context.Context, tx DbTx) error {
	for _, h := range hu {
		err := s.storage.CreateOrUpdateHeroUpgradeContext(h, ctx, tx)
		if err != nil {
			return fmt.Errorf("unable to create or update hero upgrade: %s, error: %w", h.Key, err)
		}

		for _, p := range h.Perks {
			p.HeroUpgrade = h
			err = s.storage.CreateOrUpdateHeroUpgradePerkContext(p, ctx, tx)
			if err != nil {
				return fmt.Errorf("unable to create or update hero upgrade perk: %s for hero upgrade: %s, error: %w", p.Key, h.Key, err)
			}
		}
	}

	return nil
}

func (s *HeroUpgradeSync) getHeroUpgradesFromYaml() ([]*HeroUpgrade, error) {
	var byKey map[string]heroUpgradeYaml
	if err := yaml.Unmarshal(heroUpgradeImportFile, &byKey); err != nil {
		return nil, fmt.Errorf("unable to unmarshal: %s. error: %w", heroUpgradesFile, err)
	}

	hul := make([]*HeroUpgrade, 0, len(byKey))
	for k, v := range byKey {
		classes := v.Classes.ToString()
		hu := &HeroUpgrade{
			Key:     k,
			Name:    v.Name,
			Level:   v.Level,
			Type:    v.Type,
			Icon:    v.Icon,
			Active:  v.Active.OrDefault(),
			Classes: &classes,
		}

		hu.Perks = make([]*HeroUpgradePerk, 0, len(v.Perks))
		for pk, pv := range v.Perks {
			hu.Perks = append(hu.Perks, &HeroUpgradePerk{
				HeroUpgrade: hu,
				Key:         pk,
				Value:       pv,
			})
		}

		hul = append(hul, hu)
	}

	return hul, nil
}
