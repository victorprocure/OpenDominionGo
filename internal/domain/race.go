package domain

type Race struct {
	Name                string
	Alignment           string
	HomeLandType        string
	Description         string
	Playable            bool
	AttackerDifficulty  int
	ExplorerDifficulty  int
	ConverterDifficulty int
	OverallDifficulty   int
	Perks               []RacePerk
	Units               []Unit
}

type RacePerk struct {
	Race    *Race
	TypeKey string
	Value   float64
}
