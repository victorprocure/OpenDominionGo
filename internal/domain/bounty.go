package domain

type Bounty struct {
	Round               *Round
	SourceRealm         *Realm
	SourceDominion      *Dominion
	TargetDominion      *Dominion
	CollectedByDominion *Dominion
	Type                string
	Reward              bool
}

func NewBounty(round *Round, sourceRealm *Realm, sourceDominion *Dominion, targetDominion *Dominion, collectedByDominion *Dominion, bountyType string, reward bool) *Bounty {
	return &Bounty{
		Round:               round,
		SourceRealm:         sourceRealm,
		SourceDominion:      sourceDominion,
		TargetDominion:      targetDominion,
		CollectedByDominion: collectedByDominion,
		Type:                bountyType,
		Reward:              reward,
	}
}
