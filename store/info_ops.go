package store

import "time"

type InfoOps struct {
	Id             int
	SourceRealm    *Realm
	SourceDominion *Dominion
	TargetDominion *Dominion
	Type           string
	Data           string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	TargetRealm    *Realm
	Latest         bool
}
