package dto

import "time"

type RaceJSON struct {
	ID                  int       `json:"id"`
	Key                 string    `json:"key"`
	Name                string    `json:"name"`
	Alignment           string    `json:"alignment"`
	HomeLandType        string    `json:"home_land_type"`
	Description         string    `json:"description"`
	Playable            bool      `json:"playable"`
	AttackerDifficulty  int       `json:"attacker_difficulty"`
	ExplorerDifficulty  int       `json:"explorer_difficulty"`
	ConverterDifficulty int       `json:"converter_difficulty"`
	OverallDifficulty   int       `json:"overall_difficulty"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}
