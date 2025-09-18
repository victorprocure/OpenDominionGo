package domain

import "time"

type Pack struct {
	Round           *Round
	Realm           *Realm
	Name            string
	Password        string
	Size            int
	ClosedAt        time.Time
	CreatorDominion *Dominion
	Rating          int
}
