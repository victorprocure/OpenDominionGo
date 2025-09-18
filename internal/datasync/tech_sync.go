package datasync

import (
	"context"
	"database/sql"
	"embed"
	"fmt"
	"io/fs"
	"log/slog"
	"path"
	"sort"

	"github.com/victorprocure/opendominiongo/internal/domain"
	"github.com/victorprocure/opendominiongo/internal/dto"
	"github.com/victorprocure/opendominiongo/internal/repositories"
	"github.com/victorprocure/opendominiongo/internal/repositories/tech"
	"gopkg.in/yaml.v3"
)

const techsDir = "import_data/techs"

//go:embed import_data/techs
var techsFS embed.FS

type TechSync struct {
	db *tech.TechRepo
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

		for key, t := range tp.Techs {
			// Determine effective active flag once per tech
			effActive := active
			if t.Active != nil {
				effActive = *t.Active
			}
			// Upsert tech and its perks via normalized wrapper
			_, err := s.db.UpsertTechFromSyncContext(ctx, tx, tech.TechUpsertArgs{
				Key:           key,
				Name:          t.Name,
				Prerequisites: t.Prerequisites,
				Active:        effActive,
				Version:       currentVersion,
				X:             t.X,
				Y:             t.Y,
				Perks:         t.Perks,
			})
			if err != nil {
				return fmt.Errorf("upsert tech %s: %w", key, err)
			}
		}
	}
	return nil
}

func getTechsFromFS() ([]dto.TechYaml, error) {
	entries, err := getYmlImportFiles()
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

func getYmlImportFiles() ([]string, error) {
	var files []string
	for _, pat := range []string{
		path.Join(techsDir, "*.yml"),
		path.Join(techsDir, "*.yaml"),
	} {
		matches, err := fs.Glob(techsFS, pat)
		if err != nil {
			return nil, fmt.Errorf("glob techs dir: %s, pattern: %s, error: %w", techsDir, pat, err)
		}
		files = append(files, matches...)
	}
	sort.Strings(files)
	return files, nil
}
