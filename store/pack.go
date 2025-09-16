package store

import "time"

type Pack struct {
	Id              int
	Round           *Round
	Realm           *Realm
	Name            string
	Password        string
	Size            int
	CreatedAt       time.Time
	UpdatedAt       time.Time
	ClosedAt        time.Time
	CreatorDominion *Dominion
	Rating          int
}
