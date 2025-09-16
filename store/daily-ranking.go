package store

import "time"

type DailyRanking struct {
	Id           int
	Round        *Round
	Dominion     *Dominion
	Race         *Race
	Realm        *Realm
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Key          string
	Value        int
	Rank         int
	PreviousRank int
}
