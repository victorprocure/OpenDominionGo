package store

import (
	"context"
	"embed"
	"fmt"
	"io/fs"

	"github.com/victorprocure/opendominiongo/helpers"
	"gopkg.in/yaml.v3"
)

const techsDir = "data/techs"

//go:embed data/techs
var techFS embed.FS

type techParentYaml struct {
	Version int                 `yaml:"version"`
	Tech    map[string]techYaml `yaml:"techs"`
}

type techYaml struct {
	X             int                         `yaml:"x,omitempty"`
	Y             int                         `yaml:"y,omitempty"`
	Name          string                      `yaml:"name"`
	Perks         map[string]string           `yaml:"perks"`
	Prerequisites helpers.CommaDelimitedArray `yaml:"requires"`
	Active        *bool                       `yaml:"active,omitempty"`
}

type TechSync struct {
	storage            *Storage
	currentTechVersion int
}

func NewTechSync(s *Storage) *TechSync {
	return &TechSync{storage: s, currentTechVersion: 2}
}

func (s *TechSync) Name() string {
	return "Technologies"
}

func (s *TechSync) PerformDataSync(ctx context.Context, tx DbTx) error {
	entries, err := fs.ReadDir(techFS, techsDir)
	if err != nil {
		return err
	}

	for _, e := range entries {
		name, valid, _ := helpers.IsValidYamlFileName(e)
		if !valid {
			continue
		}

		r, err := s.getTechFromFile(name)
		if err != nil {
			continue
		}

		err = s.syncTech(r, ctx, tx)
		if err != nil {
			continue
		}
	}

	return nil
}

func (s *TechSync) syncTech(tl []*Tech, ctx context.Context, tx DbTx) error {
	tlen := len(tl)
	if tlen == 0 {
		return fmt.Errorf("no tech to add: %v", tl)
	}

	for _, t := range tl {
		err := s.storage.CreateOrUpdateTechContext(t, ctx, tx)
		if err != nil {
			return fmt.Errorf("unable to create or update tech: %s, error: %w", t.Key, err)
		}

		for _, p := range t.Perks {
			p.Tech = t
			err = s.storage.CreateOrUpdateTechPerkTypeContext(p.TechPerkType, ctx, tx)
			if err != nil {
				return fmt.Errorf("unable to create or update tech perk type: %s for tech: %s, error: %w", p.TechPerkType.Key, t.Key, err)
			}

			err = s.storage.CreateOrUpdateTechPerkContext(p, ctx, tx)
			if err != nil {
				return fmt.Errorf("unable to create or update tech perk: %s for tech: %s, error: %w", p.TechPerkType.Key, t.Key, err)
			}
		}
	}

	return nil
}

func (s *TechSync) getTechFromFile(n string) ([]*Tech, error) {
	b, err := techFS.ReadFile(fmt.Sprintf("%s/%s", techsDir, n))
	if err != nil {
		return nil, fmt.Errorf("unable to read from tech file: %s. error: %w", n, err)
	}

	var r techParentYaml
	if err = yaml.Unmarshal(b, &r); err != nil {
		return nil, fmt.Errorf("unable to unmarshal: %v. error: %w", b, err)
	}

	if len(r.Tech) == 0 {
		return nil, fmt.Errorf("no techs found in %v", r)
	}

	tl := make([]*Tech, 0, len(r.Tech))
	tvm := r.Version == s.currentTechVersion
	for key, tech := range r.Tech {
		if tech.Active == nil || !tvm {
			tech.Active = &tvm
		}

		t := &Tech{
			Key:           key,
			Name:          tech.Name,
			X:             tech.X,
			Y:             tech.Y,
			Prerequisites: tech.Prerequisites.ToString(),
			Active:        *tech.Active,
			Version:       r.Version,
		}

		t.Perks = make([]*TechPerk, 0, len(tech.Perks))
		for perkKey, perkValue := range tech.Perks {
			t.Perks = append(t.Perks,
				&TechPerk{
					TechPerkType: &TechPerkType{Key: perkKey},
					Value:        perkValue,
				})
		}

		tl = append(tl, t)
	}

	return tl, nil
}
