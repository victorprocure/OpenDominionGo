package dto

type TechYaml struct {
	Version int                      `yaml:"version"`
	Techs   map[string]TechYamlEntry `yaml:"techs"`
}

// TechYamlEntry represents a single tech item within the YAML file.
type TechYamlEntry struct {
	Active        *bool     `yaml:"active,omitempty"`
	Name          string    `yaml:"name"`
	Perks         KeyValues `yaml:"perks"`
	Prerequisites []string  `yaml:"requires"`
	X             int       `yaml:"x,omitempty"`
	Y             int       `yaml:"y,omitempty"`
}
