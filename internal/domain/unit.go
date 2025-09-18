package domain

var DefaultUnitTypes = map[int]string{
	1: "offensive_specialist",
	2: "defensive_specialist",
	3: "defensive_elite",
	4: "offensive_elite",
}

type Unit struct {
	Race         *Race
	Slot         string
	Name         string
	CostPlatinum int
	CostOre      int
	PowerOffense float64
	PowerDefense float64
	NeedBoat     bool
	Type         string
	CostMana     int
	CostLumber   int
	CostGems     int
	Perks        []UnitPerk
}

type UnitPerk struct {
	TypeKey string
	Value   string
}
