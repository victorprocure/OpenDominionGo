package store

import (
	"context"
	"fmt"
	"time"
)

type Tech struct {
	Id            int
	Key           string
	Name          string
	Prerequisites string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Active        bool
	Version       int
	X             int
	Y             int
	Perks         []*TechPerk
}

func (s *Storage) CreateOrUpdateTechContext(t *Tech, ctx context.Context, tx DbTx) error {
	err := tx.QueryRowContext(ctx, `
		INSERT INTO techs (key, name, prerequisites, active, version, x, y)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		ON CONFLICT (key) DO UPDATE SET
		    name = $2,
			prerequisites = $3,
			active = $4,
			version = $5,
			x = $6,
			y = $7
			RETURNING id`,
		t.Key,
		t.Name,
		t.Prerequisites,
		t.Active,
		t.Version,
		t.X,
		t.Y,
	).Scan(&t.Id)
	if err != nil {
		return fmt.Errorf("unable to create or update tech: %s, error: %w", t.Key, err)
	}

	return nil
}

type TechPerk struct {
	Id           int
	Tech         *Tech
	TechPerkType *TechPerkType
	Value        string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (s *Storage) CreateOrUpdateTechPerkContext(t *TechPerk, ctx context.Context, tx DbTx) error {
	err := tx.QueryRowContext(ctx, `
		INSERT INTO tech_perks (tech_id, tech_perk_type_id, value)
		VALUES ($1, $2, $3)
		ON CONFLICT (tech_id, tech_perk_type_id) DO UPDATE SET
			value = $3
		RETURNING id`,
		t.Tech.Id,
		t.TechPerkType.Id,
		t.Value,
	).Scan(&t.Id)
	if err != nil {
		return fmt.Errorf("unable to create or update tech perk: %s for tech: %s, error: %w", t.TechPerkType.Key, t.Tech.Key, err)
	}
	return nil
}

type TechPerkType struct {
	Id        int
	Key       string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (s *Storage) CreateOrUpdateTechPerkTypeContext(t *TechPerkType, ctx context.Context, tx DbTx) error {
	err := tx.QueryRowContext(ctx, `
		INSERT INTO tech_perk_types (key)
		VALUES ($1)
		ON CONFLICT (key) DO UPDATE SET
			key = $1
		RETURNING id`,
		t.Key,
	).Scan(&t.Id)

	if err != nil {
		return fmt.Errorf("unable to create or update tech perk type: %s, error: %w", t.Key, err)
	}
	return nil
}
