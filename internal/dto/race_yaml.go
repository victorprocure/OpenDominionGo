package dto

import "github.com/victorprocure/opendominiongo/helpers"

type RaceYaml struct {
	Key                 string                  `yaml:"key" json:"key"`
	Name                string                  `yaml:"name" json:"name"`
	Alignment           string                  `yaml:"alignment" json:"alignment"`
	Description         *string                 `yaml:"description,omitempty" json:"description,omitempty"`
	AttackerDifficulty  int                     `yaml:"attacker_difficulty" json:"attacker_difficulty"`
	ExplorerDifficulty  int                     `yaml:"explorer_difficulty" json:"explorer_difficulty"`
	ConverterDifficulty int                     `yaml:"converter_difficulty" json:"converter_difficulty"`
	OverallDifficulty   int                     `yaml:"overall_difficulty" json:"overall_difficulty"`
	HomeLandType        string                  `yaml:"home_land_type" json:"home_land_type"`
	Perks               KeyValues               `yaml:"perks,omitempty" json:"perks,omitempty"`
	Playable            helpers.BoolDefaultTrue `yaml:"playable" json:"playable"`
	Units               []struct {
		Name     string    `yaml:"name" json:"name"`
		Perks    KeyValues `yaml:"perks,omitempty" json:"perks,omitempty"`
		NeedBoat bool      `yaml:"need_boat,omitempty" json:"need_boat,omitempty"`
		Type     string    `yaml:"type" json:"type"`
		Cost     struct {
			Platinum int `yaml:"platinum,omitempty" json:"platinum,omitempty"`
			Ore      int `yaml:"ore,omitempty" json:"ore,omitempty"`
			Lumber   int `yaml:"lumber,omitempty" json:"lumber,omitempty"`
			Gems     int `yaml:"gems,omitempty" json:"gems,omitempty"`
			Mana     int `yaml:"mana,omitempty" json:"mana,omitempty"`
		} `yaml:"cost" json:"cost"`
		Power struct {
			Offense int `yaml:"offense,omitempty" json:"offense,omitempty"`
			Defense int `yaml:"defense,omitempty" json:"defense,omitempty"`
		} `yaml:"power" json:"power"`
	} `yaml:"units" json:"units"`
}
