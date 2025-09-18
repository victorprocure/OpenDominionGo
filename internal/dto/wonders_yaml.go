package dto

import "github.com/victorprocure/opendominiongo/internal/encoding/yamlutil"

type WondersYaml struct {
	Key    string                   `yaml:"key" json:"key"`
	Name   string                   `yaml:"name" json:"name"`
	Perks  KeyValues                `yaml:"perks" json:"perks"`
	Power  int                      `yaml:"power" json:"power"`
	Active yamlutil.DefaultTrueBool `yaml:"active" json:"active"`
}
