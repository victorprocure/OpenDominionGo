package domain

import "time"

type HeroBattle struct {
	Round           *Round
	CurrentTurn     int
	WinnerCombatant *HeroCombatant
	Finished        bool
	LastProcessedAt time.Time
	PVP             bool
}

type HeroBattleQueue struct {
	ID     int
	Hero   *Hero
	Level  int
	Rating int
}

type HeroBattleAction struct {
	ID              int
	HeroBattle      *HeroBattle
	Combatant       *HeroCombatant
	TargetCombatant *HeroCombatant
	Turn            int
	Action          string
	Damage          int
	Health          int
	Description     string
}
