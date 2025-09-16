package store

import "time"

type DominionHistory struct {
	Id        int
	Dominion  *Dominion
	Event     string
	Delta     string
	CreatedAt time.Time
	IP        string
	Device    string
}