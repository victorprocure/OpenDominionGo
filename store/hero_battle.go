package store

import "time"

type HeroBattle struct {
	Id              int
	Round           *Round
	CurrentTurn     int
	WinnerCombatant *HeroCombatant
	Finished        bool
	LastProcessedAt time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time
	PVP             bool
}

type HeroBattleQueue struct {
	Id        int
	Hero      *Hero
	Level     int
	Rating    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type HeroBattleAction struct {
	Id              int
	HeroBattle      *HeroBattle
	Combatant       *HeroCombatant
	TargetCombatant *HeroCombatant
	Turn            int
	Action          string
	Damage          int
	Health          int
	Description     string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
