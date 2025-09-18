package domain

import "time"

type ForumPost struct {
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
	Round             *Round
	Dominion          *Dominion
	Title             string
	Body              string
	FlaggedForRemoval bool
	FlaggedBy         string
}
