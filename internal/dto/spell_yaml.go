package dto

import "github.com/victorprocure/opendominiongo/helpers"

type SpellYaml struct {
	Name         string                  `yaml:"name"`
	Category     string                  `yaml:"category"`
	Key          string                  `yaml:"key"`
	ManaCost     float64                 `yaml:"cost_mana"`
	StrengthCost float64                 `yaml:"cost_strength"`
	Duration     int                     `yaml:"duration,omitempty"`
	Cooldown     int                     `yaml:"cooldown,omitempty"`
	Races        []string                `yaml:"races,omitempty"`
	Active       helpers.DefaultTrueBool `yaml:"active"`
	Perks        KeyValues               `yaml:"perks"`
}
