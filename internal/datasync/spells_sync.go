package datasync

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log/slog"

	"github.com/victorprocure/opendominiongo/internal/dto"
	"github.com/victorprocure/opendominiongo/internal/repositories"
	"github.com/victorprocure/opendominiongo/internal/repositories/spells"
	"gopkg.in/yaml.v3"
)

const spellsYamlFile = "import_data/spells.yml"

//go:embed import_data/spells.yml
var spellsFile []byte

type SpellsSync struct {
	db *spells.SpellsRepo
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
		err := s.syncSpell(&spell, ctx, tx)
		if err != nil {
			continue
		}
	}

	return nil
}

func (s *SpellsSync) syncSpell(spell *dto.SpellYaml, ctx context.Context, tx repositories.DbTx) error {
	err := s.db.UpsertSpellForSyncContext(spell, ctx, tx)
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
