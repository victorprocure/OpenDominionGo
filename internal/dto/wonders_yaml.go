package dto

import "github.com/victorprocure/opendominiongo/helpers"

type WondersYaml struct {
	Key    string                  `yaml:"key" json:"key"`
	Name   string                  `yaml:"name" json:"name"`
	Perks  KeyValues               `yaml:"perks" json:"perks"`
	Power  int                     `yaml:"power" json:"power"`
	Active helpers.DefaultTrueBool `yaml:"active" json:"active"`
}
