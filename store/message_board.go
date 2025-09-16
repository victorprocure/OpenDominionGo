package store

import "time"

type MessageBoardCategory struct {
	Id           int
	Name         string
	Slug         string
	RoleRequired string
}

type MessageBoardPost struct {
	Id                 int
	MessageBoardThread *MessageBoardThread
	User               *User
	Body               string
	FlaggedForRemoval  bool
	FlaggedBy          string
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          time.Time
}

type MessageBoardThread struct {
	Id                   int
	MessageBoardCategory *MessageBoardCategory
	User                 *User
	Title                string
	Body                 string
	FlaggedForRemoval    bool
	FlaggedBy            string
	CreatedAt            time.Time
	UpdatedAt            time.Time
	DeletedAt            time.Time
	LastActivity         time.Time
}
