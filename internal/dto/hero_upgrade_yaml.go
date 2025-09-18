package dto

import "github.com/victorprocure/opendominiongo/internal/encoding/yamlutil"

// HeroUpgradeYaml represents a single hero upgrade item from YAML.
// It is expected to be loaded from a mapping of key -> HeroUpgradeYaml.
type HeroUpgradeYaml struct {
	Key     string                      `yaml:"key" json:"key"`
	Name    string                      `yaml:"name" json:"name"`
	Level   int                         `yaml:"level" json:"level"`
	Type    string                      `yaml:"type" json:"type"`
	Icon    string                      `yaml:"icon" json:"icon"`
	Perks   KeyValues                   `yaml:"perks" json:"perks"`
	Active  yamlutil.DefaultTrueBool     `yaml:"active,omitempty" json:"active"`
	Classes yamlutil.CommaDelimitedArray `yaml:"classes,omitempty" json:"classes"`
}
