package domain

type Valor struct {
	Round    *Round
	Realm    *Realm
	Dominion *Dominion
	Source   string
	Amount   float64
}
