package store

import "time"

type Permission struct {
	Id        int
	Name      string
	GuardName string
	CreatedAt time.Time
	UpdatedAt time.Time
}