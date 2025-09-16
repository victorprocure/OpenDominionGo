package store

import "time"

type DominionQueue struct {
	Dominion  *Dominion
	Source    string
	Resource  string
	Hours     int
	Amount    int
	CreatedAt time.Time
	UpdatedAt time.Time
}