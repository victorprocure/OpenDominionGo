package heroes

import (
	"context"
	"database/sql"
	_ "embed"
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/victorprocure/opendominiongo/internal/dto"
	"github.com/victorprocure/opendominiongo/internal/repositories"
)

//go:embed sql/upsert_hero_upgrade_with_perks.sql
var upsertHeroUpgradeWithPerksSQL string

type HeroesRepo struct {
	db  *sql.DB
	log *slog.Logger
}

func NewHeroesRepo(db *sql.DB, log *slog.Logger) *HeroesRepo {
	return &HeroesRepo{db: db, log: log}
}

// HeroUpgradeUpsertArgs is the repo-boundary contract for upserting a hero upgrade.
// It decouples the repository from YAML-specific DTOs and normalizes inputs.
type HeroUpgradeUpsertArgs struct {
	Key     string
	Name    string
	Level   int
	Type    string
	Icon    string
	Classes *string
	Active  bool
	Perks   dto.KeyValues
}

func (r *HeroesRepo) CreateOrUpdateHeroUpgradeSyncContext(ctx context.Context, tx repositories.DbTx, a HeroUpgradeUpsertArgs) (int, error) {
	var perksJSON []byte
	if len(a.Perks) > 0 {
		b, err := json.Marshal(a.Perks)
		if err != nil {
			return 0, fmt.Errorf("marshal hero upgrade perks: %w", err)
		}
		perksJSON = b
	}

	var id int
	err := tx.QueryRowContext(ctx, upsertHeroUpgradeWithPerksSQL,
		a.Key,
		a.Name,
		a.Level,
		a.Type,
		a.Icon,
		a.Classes,
		a.Active,
		perksJSON,
	).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("upsert hero upgrade %q: %w", a.Key, err)
	}
	return id, nil
}
