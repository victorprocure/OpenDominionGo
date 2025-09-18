package domain

type Wonder struct {
	Key    string
	Name   string
	Power  int
	Active bool
	Perks  []WonderPerk
}

type WonderPerk struct {
	TypeKey string
	Value   string
}
