package store

import "time"

type Bounty struct {
	Id                  int
	Round               *Round
	SourceRealm         *Realm
	SourceDominion      *Dominion
	TargetDominion      *Dominion
	CollectedByDominion *Dominion
	Type                string
	CreatedAt           time.Time
	UpdatedAt           time.Time
	Reward              bool
}
