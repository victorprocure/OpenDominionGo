package domain

import "time"

type HeroTournament struct {
	Round              *Round
	Name               string
	CurrentRoundNumber int
	Finished           bool
	WinnerDominion     *Dominion
	StartDate          time.Time
}

type HeroTournamentBattle struct {
	HeroTournament *HeroTournament
	HeroBattle     *HeroBattle
	RoundNumber    int
}

type HeroTournamentParticipant struct {
	HeroTournament *HeroTournament
	Hero           *Hero
	Wins           int
	Losses         int
	Draws          int
	Standing       int
	Eliminated     bool
}
