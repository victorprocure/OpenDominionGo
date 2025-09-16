package store

import "time"

type Spell struct {
	Id           int
	Key          string
	Name         string
	Category     string
	CostMana     float64
	CostStrength float64
	Duration     int
	Cooldown     int
	Races        string
	Active       bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type SpellPerk struct {
	Id            int
	Spell         *Spell
	SpellPerkType *SpellPerkType
	Value         string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type SpellPerkType struct {
	Id        int
	Key       string
	CreatedAt time.Time
	UpdatedAt time.Time
}
