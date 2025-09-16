package store

import "time"

type Realm struct {
	Id                int
	Round             *Round
	MonarchDominion   *Dominion
	Alignment         string
	Number            int
	Name              string
	CreatedAt         time.Time
	UpdatedAt         time.Time
	Motd              string
	MotdUpdatedAt     time.Time
	DiscordRoleId     string
	Rating            int
	DiscordCategoryId string
	GeneralDominion   *Dominion
	MagisterDominion  *Dominion
	MageDominion      *Dominion
	JesterDominion    *Dominion
	SpymasterDominion *Dominion
	Settings          string
	Valor             int
}

type RealHistory struct {
	Id        int
	Realm     *Realm
	Dominion  *Dominion
	Event     string
	Delta     string
	CreatedAt time.Time
}

type RealmWar struct {
	Id                   int
	SourceRealm          *Realm
	SourceRealmNameStart string
	TargetRealm          *Realm
	TargetRealmNameStart string
	ActiveAt             time.Time
	InactiveAt           time.Time
	CreatedAt            time.Time
	UpdatedAt            time.Time
	SourceRealmNameEnd   string
	TargetRealmNameEnd   string
}
