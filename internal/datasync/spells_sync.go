package datasync

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log/slog"

	"github.com/victorprocure/opendominiongo/internal/dto"
	"github.com/victorprocure/opendominiongo/internal/helpers"
	"github.com/victorprocure/opendominiongo/internal/repositories"
	"github.com/victorprocure/opendominiongo/internal/repositories/spells"
	"gopkg.in/yaml.v3"
)

const spellsYamlFile = "import_data/spells.yml"

//go:embed import_data/spells.yml
var spellsFile []byte

type SpellsSync struct {
	db *spells.Repo
}

func NewSpellsSync(db *sql.DB, log *slog.Logger) *SpellsSync {
	return &SpellsSync{db: spells.NewSpellsRepository(db, log)}
}

func (s *SpellsSync) Name() string {
	return "Spells"
}

func (s *SpellsSync) PerformDataSync(ctx context.Context, tx repositories.DbTx) error {
	spells, err := s.getSpellsFromYaml()
	if err != nil {
		return err
	}

	for _, spell := range spells {
		if err := s.syncSpell(ctx, tx, &spell); err != nil {
			// Stop on first error so we surface the underlying DB error
			// (Postgres marks the transaction as failed after the first error).
			return err
		}
	}

	return nil
}

func (s *SpellsSync) syncSpell(ctx context.Context, tx repositories.DbTx, spell *dto.SpellYaml) error {
	// map dto -> repo args
	perks := helpers.PerksToMap(spell.Perks)
	err := s.db.UpsertFromSyncContext(ctx, tx, spells.UpsertArgs{
		Key:          spell.Key,
		Name:         spell.Name,
		Category:     spell.Category,
		ManaCost:     spell.ManaCost,
		StrengthCost: spell.StrengthCost,
		Duration:     spell.Duration,
		Cooldown:     spell.Cooldown,
		Active:       spell.Active.OrDefault(),
		Races:        spell.Races,
		Perks:        perks,
	})
	if err != nil {
		return fmt.Errorf("unable to create or update spell: %s, error: %w", spell.Key, err)
	}

	return nil
}

func (s *SpellsSync) getSpellsFromYaml() ([]dto.SpellYaml, error) {
	var byKey map[string]dto.SpellYaml
	if err := yaml.Unmarshal(spellsFile, &byKey); err != nil {
		return nil, fmt.Errorf("unable to umarshall spells file: %s, error: %w", spellsYamlFile, err)
	}

	dbSpells := make([]dto.SpellYaml, 0, len(byKey))
	for k, v := range byKey {
		v.Key = k
		dbSpells = append(dbSpells, v)
	}

	return dbSpells, nil
}
