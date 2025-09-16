package store

import "time"

type DominionSpell struct {
	Dominion       *Dominion
	Duration       int
	CastByDominion *Dominion
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Spell          *Spell
}
