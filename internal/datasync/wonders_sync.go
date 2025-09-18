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
	"github.com/victorprocure/opendominiongo/internal/repositories/wonders"
	"gopkg.in/yaml.v3"
)

//go:embed import_data/wonders.yml
var wonderImportFile []byte

type WondersSync struct {
	db *wonders.Repo
}

func NewWondersSync(db *sql.DB, log *slog.Logger) *WondersSync {
	return &WondersSync{db: wonders.NewWondersRepo(db, log)}
}

func (s *WondersSync) Name() string {
	return "Wonders"
}

func (s *WondersSync) PerformDataSync(ctx context.Context, tx repositories.DbTx) error {
	w, err := s.getWondersFromYaml()
	if err != nil {
		return fmt.Errorf("unable to get wonders from yaml: %w", err)
	}

	if len(w) == 0 {
		return fmt.Errorf("no wonders to add: %v", w)
	}

	err = s.syncWonders(w, ctx, tx)
	if err != nil {
		return fmt.Errorf("unable to sync wonders: %w", err)
	}

	return nil
}

func (s *WondersSync) syncWonders(wl []dto.WondersYaml, ctx context.Context, tx repositories.DbTx) error {
	for _, w := range wl {
		perks := helpers.PerksToMap(w.Perks)
		err := s.db.UpsertFromSyncContext(ctx, tx, wonders.UpsertArgs{
			Key:    w.Key,
			Name:   w.Name,
			Power:  w.Power,
			Active: w.Active.OrDefault(),
			Perks:  perks,
		})
		if err != nil {
			return fmt.Errorf("unable to create or update wonder: %s, error: %w", w.Key, err)
		}
	}
	return nil
}

func (s *WondersSync) getWondersFromYaml() ([]dto.WondersYaml, error) {
	var byKey map[string]dto.WondersYaml
	if err := yaml.Unmarshal(wonderImportFile, &byKey); err != nil {
		return nil, fmt.Errorf("unable to unmarshal: %v. error: %w", wonderImportFile, err)
	}

	wl := make([]dto.WondersYaml, 0, len(byKey))
	for key, wonder := range byKey {
		wonder.Key = key
		wl = append(wl, wonder)
	}

	return wl, nil
}
