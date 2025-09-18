package domain

type HeroCombatant struct {
	HeroBattle    *HeroBattle
	Hero          *Hero
	Dominion      *Dominion
	Name          string
	Health        int
	Attack        int
	Defense       int
	Evasion       int
	Focus         int
	Counter       int
	Recover       int
	CurrentHealth int
	HasFocus      bool
	Actions       string
	LastAction    string
	TimeBank      int
	Automated     bool
	Strategy      string
	Level         int
}