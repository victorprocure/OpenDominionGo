package domain

import "time"

type MessageBoardCategory struct {
	Name         string
	Slug         string
	RoleRequired string
}

type MessageBoardPost struct {
	MessageBoardThread *MessageBoardThread
	User               *User
	Body               string
	FlaggedForRemoval  bool
	FlaggedBy          string
	IsDeleted          bool
}

type MessageBoardThread struct {
	MessageBoardCategory *MessageBoardCategory
	User                 *User
	Title                string
	Body                 string
	FlaggedForRemoval    bool
	FlaggedBy            string
	IsDeleted            bool
	LastActivity         time.Time
}
