package domain

const CurrentTechVersion = 2

type Tech struct {
	Name          string
	Prerequisites string
	Active        bool
	Version       int
	X             int
	Y             int
	Perks         []TechPerk
}

type TechPerk struct {
	Tech    *Tech
	TypeKey string
	Value   string
}