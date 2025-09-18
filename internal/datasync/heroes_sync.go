package datasync

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log/slog"

	"github.com/victorprocure/opendominiongo/internal/dto"
	"github.com/victorprocure/opendominiongo/internal/repositories"
	"github.com/victorprocure/opendominiongo/internal/repositories/heroes"
	"gopkg.in/yaml.v3"
)

//go:embed import_data/heroes.yml
var heroUpgradeImportFile []byte

type HeroesSync struct {
	db *heroes.HeroesRepo
}

func NewHeroesSync(db *sql.DB, log *slog.Logger) *HeroesSync {
	return &HeroesSync{db: heroes.NewHeroesRepo(db, log)}
}

func (s *HeroesSync) Name() string {
	return "Hero Upgrades"
}

func (s *HeroesSync) PerformDataSync(ctx context.Context, tx repositories.DbTx) error {
	hu, err := s.getHeroUpgradesFromYaml()
	if err != nil {
		return fmt.Errorf("unable to get hero upgrades from yaml: %w", err)
	}

	if len(hu) == 0 {
		return fmt.Errorf("no hero upgrades to add: %v", hu)
	}

	for _, h := range hu {
		// Build normalized repo args
		classes := h.Classes.ToString()
		_, err := s.db.CreateOrUpdateHeroUpgradeSyncContext(ctx, tx, heroes.HeroUpgradeUpsertArgs{
			Key:     h.Key,
			Name:    h.Name,
			Level:   h.Level,
			Type:    h.Type,
			Icon:    h.Icon,
			Classes: &classes,
			Active:  h.Active.OrDefault(),
			Perks:   h.Perks,
		})
		if err != nil {
			return fmt.Errorf("upsert hero upgrade %s: %w", h.Key, err)
		}
	}

	return nil
}

func (s *HeroesSync) getHeroUpgradesFromYaml() ([]dto.HeroUpgradeYaml, error) {
	var byKey map[string]dto.HeroUpgradeYaml
	if err := yaml.Unmarshal(heroUpgradeImportFile, &byKey); err != nil {
		return nil, fmt.Errorf("unable to unmarshal heroes.yml: %w", err)
	}

	hul := make([]dto.HeroUpgradeYaml, 0, len(byKey))
	for k, v := range byKey {
		v.Key = k
		hul = append(hul, v)
	}
	return hul, nil
}
