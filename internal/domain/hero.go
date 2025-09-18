package domain

type Hero struct {
	Id               int
	Dominion         *Dominion
	Name             string
	Class            string
	Experience       float64
	StatCombatWins   int
	StatCombatLosses int
	StatCombatDraws  int
	CombatRating     int
	HeroUpgrades    []HeroHeroUpgrade
}

type HeroHeroUpgrade struct {
	Hero        *Hero
	HeroUpgrade *HeroUpgrade
}


