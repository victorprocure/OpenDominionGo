package store

import "time"

type CouncilThread struct {
	Id             int
	Realm          *Realm
	Dominion       *Dominion
	Title          string
	Body           string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      time.Time
	LastActivityAt time.Time
}
