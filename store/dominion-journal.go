package store

import "time"

type DominionJournal struct {
	Id        int
	Dominion  *Dominion
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}