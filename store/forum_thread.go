package store

import "time"

type ForumPost struct {
	Id                int
	ForumThread       *ForumThread
	Dominion          *Dominion
	Body              string
	FlaggedForRemoval bool
	FlaggedBy         string
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         time.Time
}

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