package store

import "time"

type Achievement struct {
	Id          int
	Name        string
	Description string
	Icon        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}