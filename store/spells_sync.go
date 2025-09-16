package store

import (
	"context"
	"fmt"
	_ "embed"
	"github.com/victorprocure/opendominiongo/helpers"
	"gopkg.in/yaml.v3"
)

const spellsYamlFile = "data/spells.yml"
//go:embed data/spells.yml
var spellsFile []byte

type SpellsSync struct {
	storage        *Storage
}

func NewSpellsSync(s *Storage) *SpellsSync {
	return &SpellsSync{storage: s}
}

func (s *SpellsSync) Name() string {
	return "Spells"
}

func (s *SpellsSync) PerformDataSync(ctx context.Context, tx DbTx) error {
	spells, err := s.getSpellsFromYaml()
	if err != nil {
		return err
	}

	for _, spell := range spells {
		err := s.syncSpell(spell, ctx, tx)
		if err != nil {
			continue
		}

		err = s.syncSpellPerks(spell.Perks, spell, ctx, tx)
		if err != nil {
			return fmt.Errorf("unable to sync spell perks for spell: %s, error: %w", spell.Key, err)
		}
	}

	return nil
}

func (s *SpellsSync) syncSpell(spell *Spell, ctx context.Context, tx DbTx) error {
	err := s.storage.CreateOrUpdateSpellContext(spell, ctx, tx)
	if err != nil {
		return fmt.Errorf("unable to create or update spell: %s, error: %w", spell.Key, err)
	}

	return nil
}

func (s *SpellsSync) syncSpellPerks(perks []*SpellPerk, spell *Spell, ctx context.Context, tx DbTx) error {
	if len(perks) == 0 {
		return fmt.Errorf("no perks to sync")
	}

	for _, perk := range perks {
		if perk.SpellPerkType == nil {
			return fmt.Errorf("spell perk type is nil for spell: %s", perk.Spell.Key)
		}

		perk.Spell = spell
		err := s.storage.CreateOrUpdateSpellPerkTypeContext(perk.SpellPerkType, ctx, tx)
		if err != nil {
			return fmt.Errorf("unable to create or update spell perk type: %s, error: %w", perk.SpellPerkType.Key, err)
		}

		err = s.storage.CreateOrUpdateSpellPerkContext(perk, ctx, tx)
		if err != nil {
			return fmt.Errorf("unable to create or update spell perk for spell: %s, spell perk type: %s, error: %w", perk.Spell.Key, perk.SpellPerkType.Key, err)
		}
	}

	return nil
}

func (s *SpellsSync) getSpellsFromYaml() ([]*Spell, error) {
	type spell struct {
		Name         string                  `yaml:"name"`
		Category     string                  `yaml:"category"`
		ManaCost     float64                 `yaml:"cost_mana"`
		StrengthCost float64                 `yaml:"cost_strength"`
		Duration     int                     `yaml:"duration,omitempty"`
		Cooldown     int                     `yaml:"cooldown,omitempty"`
		Races        []string                `yaml:"races,omitempty"`
		Active       helpers.BoolDefaultTrue `yaml:"active"`
		Perks        map[string]string       `yaml:"perks"`
	}
	spells := make(map[string]spell)
	if err := yaml.Unmarshal(spellsFile, &spells); err != nil {
		return nil, fmt.Errorf("unable to umarshall spells file: %s, error: %w", spellsYamlFile, err)
	}

	spellLen := len(spells)
	keys := make([]string, 0, spellLen)
	for k := range spells {
		keys = append(keys, k)
	}

	dbSpells := make([]*Spell, 0, spellLen)
	for _, k := range keys {
		v := spells[k]
		dbSpell := &Spell{
			Name:         v.Name,
			Key:          k,
			Category:     v.Category,
			CostMana:     v.ManaCost,
			CostStrength: v.StrengthCost,
			Duration:     v.Duration,
			Cooldown:     v.Cooldown,
			Active:       v.Active.OrDefault(),
		}

		raceLen := len(v.Races)
		var races []*Race
		if raceLen > 0 {
			races = make([]*Race, 0, raceLen)
			for _, raceKey := range v.Races {
				race := &Race{Key: raceKey}
				races = append(races, race)
			}
		}
		dbSpell.Races = races

		perkLen := len(v.Perks)
		var perks []*SpellPerk
		if perkLen > 0 {
			perks = make([]*SpellPerk, 0, perkLen)
			for perkKey, perkValue := range v.Perks {
				spellPerkType := &SpellPerkType{Key: perkKey}
				spellPerk := &SpellPerk{
					SpellPerkType: spellPerkType,
					Value:         perkValue,
				}
				perks = append(perks, spellPerk)
			}
		}
		dbSpell.Perks = perks

		dbSpells = append(dbSpells, dbSpell)
	}

	return dbSpells, nil
}
