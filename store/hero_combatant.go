package store

import "time"

type HeroCombatant struct {
	Id            int
	HeroBattle    *HeroBattle
	Hero          *Hero
	Dominion      *Dominion
	Name          string
	Health        int
	Attack        int
	Defense       int
	Evasion       int
	Focus         int
	Counter       int
	Recover       int
	CurrentHealth int
	HasFocus      bool
	Actions       string
	LastAction    string
	TimeBank      int
	Automated     bool
	Strategy      string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Level         int
}