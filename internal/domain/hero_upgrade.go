package domain

type HeroUpgrade struct {
	Key     string
	Name    string
	Level   int
	Type    string
	Icon    string
	Classes []string
	Active  bool
	Perks   []HeroUpgradePerk
}

type HeroUpgradePerk struct {
	HeroUpgrade *HeroUpgrade
	Key         string
	Value       string
}
