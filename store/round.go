package store

import "time"

type Round struct {
	Id                           int
	RoundLeague                  *RoundLeague
	Number                       int
	Name                         string
	StartDate                    time.Time
	EndDate                      time.Time
	CreatedAt                    time.Time
	UpdatedAt                    time.Time
	RealmSize                    int
	PackSize                     int
	PlayersPerRace               int
	MixedAlignment               bool
	OffensiveActionsProhibitedAt time.Time
	DiscordGuildId               string
	TechVersion                  int
	LargestHit                   int
	AssignmentComplete           bool
}

type RoundLeague struct {
	Id          int
	Key         string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type RoundWonder struct {
	Id        int
	Round     *Round
	Realm     *Realm
	Wonder    *Wonder
	Power     int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type RoundWonderDamage struct {
	Id          int
	RoundWonder *RoundWonder
	Realm       *Realm
	Dominion    *Dominion
	Damage      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Source      string
}
