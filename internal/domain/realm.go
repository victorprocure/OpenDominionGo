package domain

import "time"

type Realm struct {
	Round             *Round
	MonarchDominion   *Dominion
	Alignment         string
	Number            int
	Name              string
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
	Settings          []string
	Valor             int
}

type RealHistory struct {
	Realm    *Realm
	Dominion *Dominion
	Event    string
	Delta    string
}

type RealmWar struct {
	SourceRealm          *Realm
	SourceRealmNameStart string
	TargetRealm          *Realm
	TargetRealmNameStart string
	ActiveAt             time.Time
	InactiveAt           time.Time
	SourceRealmNameEnd   string
	TargetRealmNameEnd   string
}
