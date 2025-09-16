package store

import "time"

type Hero struct {
	Id               int
	Dominion         *Dominion
	Name             string
	Class            string
	Experience       float64
	CreatedAt        time.Time
	UpdatedAt        time.Time
	StatCombatWins   int
	StatCombatLosses int
	StatCombatDraws  int
	CombatRating     int
}
