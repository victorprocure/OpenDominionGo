package store

import "time"

type DominionTech struct {
	Id        int
	Dominion  *Dominion
	Tech      *Tech
	CreatedAt time.Time
	UpdatedAt time.Time
}