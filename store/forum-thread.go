package store

import "time"

type ForumThread struct {
	Id                int
	Round             *Round
	Dominion          *Dominion
	Title             string
	Body              string
	FlaggedForRemoval bool
	FlaggedBy         string
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         time.Time
	LastActivityAt    time.Time
}