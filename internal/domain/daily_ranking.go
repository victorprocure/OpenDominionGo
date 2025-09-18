package domain

type DailyRanking struct {
	Round        *Round
	Dominion     *Dominion
	Race         *Race
	Realm        *Realm
	Key          string
	Value        int
	Rank         int
	PreviousRank int
}

func NewDailyRanking(round *Round, dominion *Dominion, race *Race, realm *Realm, key string, value int, rank int, previousRank int) *DailyRanking {
	return &DailyRanking{
		Round:        round,
		Dominion:     dominion,
		Race:         race,
		Realm:        realm,
		Key:          key,
		Value:        value,
		Rank:         rank,
		PreviousRank: previousRank,
	}
}
