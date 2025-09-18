package datasync

import (
	"context"
	"database/sql"
	"embed"
	"fmt"
	"io/fs"
	"log/slog"
	"sort"
	"strings"

	"github.com/victorprocure/opendominiongo/internal/domain"
	"github.com/victorprocure/opendominiongo/internal/dto"
	"github.com/victorprocure/opendominiongo/internal/encoding/yamlutil"
	"github.com/victorprocure/opendominiongo/internal/helpers"
	"github.com/victorprocure/opendominiongo/internal/repositories"
	"github.com/victorprocure/opendominiongo/internal/repositories/tech"
	"gopkg.in/yaml.v3"
)

const techsDir = "import_data/techs"

//go:embed import_data/techs
var techsFS embed.FS

type TechSync struct {
	db *tech.Repo
}

func NewTechSync(db *sql.DB, log *slog.Logger) *TechSync {
	return &TechSync{db: tech.NewTechRepo(db, log)}
}

func (s *TechSync) Name() string {
	return "Tech"
}

func (s *TechSync) PerformDataSync(ctx context.Context, tx repositories.DbTx) error {
	techs, err := getTechsFromFS()
	if err != nil {
		return fmt.Errorf("unable to get techs from fs: %w", err)
	}

	for _, tp := range techs {
		currentVersion := tp.Version
		active := false
		if currentVersion == domain.CurrentTechVersion {
			active = true
		}

		// Iterate techs in deterministic key order
		keys := make([]string, 0, len(tp.Techs))
		for k := range tp.Techs {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, key := range keys {
			t := tp.Techs[key]
			// Determine effective active flag once per tech
			effActive := active
			if t.Active != nil {
				effActive = *t.Active
			}
			// Normalize at sync boundary
			perks := helpers.PerksToMap(t.Perks)
			prereq := strings.Join(t.Prerequisites, ",")
			// Upsert tech and its perks via normalized wrapper
			_, err := s.db.UpsertFromSyncContext(ctx, tx, tech.UpsertArgs{
				Key:           key,
				Name:          t.Name,
				Prerequisites: prereq,
				Active:        effActive,
				Version:       currentVersion,
				X:             t.X,
				Y:             t.Y,
				Perks:         perks,
			})
			if err != nil {
				return fmt.Errorf("upsert tech %s: %w", key, err)
			}
		}
	}
	return nil
}

func getTechsFromFS() ([]dto.TechYaml, error) {
	entries, err := yamlutil.GetYmlImportFiles(techsFS, techsDir)
	if err != nil {
		return nil, fmt.Errorf("unable to get tech files from fs: %w", err)
	}

	var techs []dto.TechYaml
	for _, e := range entries {
		b, err := fs.ReadFile(techsFS, e)
		if err != nil {
			return nil, fmt.Errorf("unable to read tech file: %s, error: %w", e, err)
		}

		var tech dto.TechYaml
		if err := yaml.Unmarshal(b, &tech); err != nil {
			return nil, fmt.Errorf("unable to unmarshal tech file: %s, error: %w", e, err)
		}

		techs = append(techs, tech)
	}

	return techs, nil
}
