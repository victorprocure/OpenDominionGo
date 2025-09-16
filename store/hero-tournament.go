package store

import "time"

type HeroTournament struct {
	Id                 int
	Round              *Round
	Name               string
	CurrentRoundNumber int
	Finished           bool
	WinnerDominion     *Dominion
	CreatedAt          time.Time
	UpdatedAt          time.Time
	StartDate          time.Time
}

type HeroTournamentBattle struct {
	Id             int
	HeroTournament *HeroTournament
	HeroBattle     *HeroBattle
	RoundNumber    int
}

type HeroTournamentParticipant struct {
	Id             int
	HeroTournament *HeroTournament
	Hero           *Hero
	Wins           int
	Losses         int
	Draws          int
	Standing       int
	Eliminated     bool
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
