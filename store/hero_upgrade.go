package store

import (
	"context"
	"fmt"
	"time"
)

type HeroUpgrade struct {
	Id        int
	Key       string
	Name      string
	Level     int
	Type      string
	Icon      string
	Classes   *string
	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
	Perks     []*HeroUpgradePerk
}

func (s *Storage) CreateOrUpdateHeroUpgradeContext(h *HeroUpgrade, ctx context.Context, tx DbTx) error {
	err := tx.QueryRowContext(ctx, `
		INSERT INTO hero_upgrades (key, name, level, type, icon, classes, active)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		ON CONFLICT (key) DO UPDATE SET
		    name = $2,
			level = $3,
			type = $4,
			icon = $5,
			classes = $6,
			active = $7
			RETURNING id`,
		h.Key,
		h.Name,
		h.Level,
		h.Type,
		h.Icon,
		h.Classes,
		h.Active,
	).Scan(&h.Id)
	if err != nil {
		return fmt.Errorf("unable to create or update hero upgrade: %s, error: %w", h.Key, err)
	}

	return nil
}

type HeroUpgradePerk struct {
	Id          int
	HeroUpgrade *HeroUpgrade
	Key         string
	Value       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (s *Storage) CreateOrUpdateHeroUpgradePerkContext(h *HeroUpgradePerk, ctx context.Context, tx DbTx) error {
	err := tx.QueryRowContext(ctx, `
		INSERT INTO hero_upgrade_perks (hero_upgrade_id, key, value)
		VALUES ($1, $2, $3)
		ON CONFLICT (hero_upgrade_id, key) DO UPDATE SET
		    value = $3
		RETURNING id`,
		h.HeroUpgrade.Id,
		h.Key,
		h.Value,
	).Scan(&h.Id)
	
	if err != nil {
		return fmt.Errorf("unable to create or update hero upgrade perk: %s for hero upgrade: %s, error: %w", h.Key, h.HeroUpgrade.Key, err)
	}
	return nil
}
