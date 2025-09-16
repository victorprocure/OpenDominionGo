package store

import "time"

type CouncilPost struct {
	Id            int
	CouncilThread *CouncilThread
	Dominion      *Dominion
	Body          string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     time.Time
}