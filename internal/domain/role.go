package domain

type Role struct {
	Name        string
	GuardName   string
	Permissions []Permission
}
