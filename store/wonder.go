package store

import (
	"context"
	"fmt"
	"time"
)

type Wonder struct {
	Id        int
	Key       string
	Name      string
	Power     int
	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
	Perks     []*WonderPerk
}

func (s *Storage) CreateOrUpdateWonderContext(w *Wonder, ctx context.Context, tx DbTx) error {
	err := tx.QueryRowContext(ctx, `
		INSERT INTO wonders (key, name, power, active)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (key) DO UPDATE SET
		    name = $2,
			power = $3,
			active = $4
			RETURNING id`,
		w.Key,
		w.Name,
		w.Power,
		w.Active,
	).Scan(&w.Id)

	if err != nil {
		return fmt.Errorf("unable to create or update wonder: %s, error: %w", w.Key, err)
	}
	return nil
}

type WonderPerk struct {
	Id             int
	Wonder         *Wonder
	WonderPerkType *WonderPerkType
	Value          string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (s *Storage) CreateOrUpdateWonderPerkContext(w *WonderPerk, ctx context.Context, tx DbTx) error {
	err := tx.QueryRowContext(ctx, `
		INSERT INTO wonder_perks (wonder_id, wonder_perk_type_id, value)
		VALUES ($1, $2, $3)
		ON CONFLICT (wonder_id, wonder_perk_type_id) DO UPDATE SET
		    value = $3
		RETURNING id`,
		w.Wonder.Id,
		w.WonderPerkType.Id,
		w.Value,
	).Scan(&w.Id)
	if err != nil {
		return fmt.Errorf("unable to create or update wonder perk: %s for wonder: %s, error: %w", w.WonderPerkType.Key, w.Wonder.Key, err)
	}
	return nil
}

type WonderPerkType struct {
	Id        int
	Key       string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (s *Storage) CreateOrUpdateWonderPerkTypeContext(w *WonderPerkType, ctx context.Context, tx DbTx) error {
	err := tx.QueryRowContext(ctx, `
		INSERT INTO wonder_perk_types (key) VALUES ($1)
		ON CONFLICT (key) DO UPDATE SET key = EXCLUDED.key
		RETURNING id`,
		w.Key,
	).Scan(&w.Id)
	if err != nil {
		return fmt.Errorf("unable to create or update wonder perk type: %s, error: %w", w.Key, err)
	}
	return nil
}
