package store

import "time"

type Wonder struct {
	Id        int
	Key       string
	Name      string
	Power     int
	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type WonderPerk struct {
	Id             int
	Wonder         *Wonder
	WonderPerkType *WonderPerkType
	Value          string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type WonderPerkType struct {
	Id        int
	Key       string
	CreatedAt time.Time
	UpdatedAt time.Time
}
