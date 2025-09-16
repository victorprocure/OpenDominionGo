package store

import "time"

type Valor struct {
	Id        int
	Round     *Round
	Realm     *Realm
	Dominion  *Dominion
	Source    string
	Amount    float64
	CreatedAt time.Time
	UpdatedAt time.Time
}
