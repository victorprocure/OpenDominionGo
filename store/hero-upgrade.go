package store

import "time"

type HeroHeroUpgrade struct {
	Id          int
	Hero        *Hero
	HeroUpgrade *HeroUpgrade
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type HeroUpgrade struct {
	Id        int
	Key       string
	Name      string
	Level     int
	Type      string
	Icon      string
	Classes   string
	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type HeroUpgradePerk struct {
	Id          int
	HeroUpgrade *HeroUpgrade
	Key         string
	Value       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
