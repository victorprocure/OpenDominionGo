package domain

import "time"

type Round struct {
	RoundLeague                  *RoundLeague
	Number                       int
	Name                         string
	StartDate                    time.Time
	EndDate                      time.Time
	RealmSize                    int
	PackSize                     int
	PlayersPerRace               int
	MixedAlignment               bool
	OffensiveActionsProhibitedAt time.Time
	DiscordGuildID               string
	TechVersion                  int
	LargestHit                   int
	AssignmentComplete           bool
}

type RoundLeague struct {
	Key         string
	Description string
}

type RoundWonder struct {
	Round  *Round
	Realm  *Realm
	Wonder *Wonder
	Power  int
}

type RoundWonderDamage struct {
	RoundWonder *RoundWonder
	Realm       *Realm
	Dominion    *Dominion
	Damage      int
	Source      string
}
