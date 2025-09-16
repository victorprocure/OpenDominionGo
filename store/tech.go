package store

import "time"

type Tech struct {
	Id            int
	Key           string
	Name          string
	Prerequisites string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Active        bool
	Version       int
	X             int
	Y             int
}

type TechPerk struct {
	Id           int
	Tech         *Tech
	TechPerkType *TechPerkType
	Value        string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type TechPerkType struct {
	Id        int
	Key       string
	CreatedAt time.Time
	UpdatedAt time.Time
}
