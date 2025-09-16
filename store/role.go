package store

import "time"

type Role struct {
	Id        int
	Name      string
	GuardName string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type RoleHasPermission struct {
	Permission *Permission
	Role       *Role
}
