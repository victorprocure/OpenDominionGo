package domain

type InfoOps struct {
	Id             int
	SourceRealm    *Realm
	SourceDominion *Dominion
	TargetDominion *Dominion
	Type           string
	Data           string
	TargetRealm    *Realm
	Latest         bool
}
