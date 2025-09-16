package store

import (
	"context"
	"time"
)

var DefaultUnitTypes = map[int]string {
		1: "offensive_specialist",
		2: "defensive_specialist",
		3: "defensive_elite",
		4: "offensive_elite",
	}

type Unit struct {
	Id           int
	Race         *Race
	Slot         string
	Name         string
	CostPlatinum int
	CostOre      int
	PowerOffense float64
	PowerDefense float64
	NeedBoat     bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Type         string
	CostMana     int
	CostLumber   int
	CostGems     int
	Perks        []*UnitPerk
}

func (s *Storage) CreateOrUpdateUnitContext(u *Unit, ctx context.Context, db DbTx) error {
	err := db.QueryRowContext(ctx, `
		INSERT INTO units (race_id, slot, name, cost_platinum, cost_ore, power_offense, power_defense, need_boat, type, cost_mana, cost_lumber, cost_gems)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
		ON CONFLICT (race_id, slot) DO UPDATE SET
			name = $3,
			cost_platinum = $4,
			cost_ore = $5,
			power_offense = $6,
			power_defense = $7,
			need_boat = $8,
			type = $9,
			cost_mana = $10,
			cost_lumber = $11,
			cost_gems = $12
			RETURNING id`,
			u.Race.Id,
			u.Slot,
			u.Name,
			u.CostPlatinum,
			u.CostOre,
			u.PowerOffense,
			u.PowerDefense,
			u.NeedBoat,
			u.Type,
			u.CostMana,
			u.CostLumber,
			u.CostGems,).Scan(&u.Id,)

	if err != nil {
		return err;
	}
	return nil
}

type UnitPerk struct {
	Id           int
	Unit         *Unit
	UnitPerkType *UnitPerkType
	Value        string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (s *Storage) CreateOrUpdateUnitPerkContext(up *UnitPerk, ctx context.Context, db DbTx) error {
	err := db.QueryRowContext(ctx, `
		INSERT INTO unit_perks (unit_id, unit_perk_type_id, value) VALUES ($1, $2, $3)
		ON CONFLICT (unit_id, unit_perk_type_id) DO UPDATE SET value = $3
		RETURNING id`,
		up.Unit.Id,
		up.UnitPerkType.Id,
		up.Value,
	).Scan(&up.Id)
	
	if err != nil {
		return err
	}
	return nil
}

type UnitPerkType struct {
	Id        int
	Key       string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (s *Storage) CreateOrUpdateUnitPerkTypeContext(upt *UnitPerkType, ctx context.Context, db DbTx) error {
	err := db.QueryRowContext(ctx, `
		INSERT INTO unit_perk_types (key) VALUES ($1)
		ON CONFLICT (key) DO UPDATE SET key = $1
		RETURNING id`,
		upt.Key,
	).Scan(&upt.Id)

	if err != nil {
		return err
	}
	return nil
}
