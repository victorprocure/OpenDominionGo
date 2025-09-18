package action

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log/slog"

	"github.com/victorprocure/opendominiongo/internal/repositories"
)

//go:embed sql/insert_hero_battle_action.sql
var insertHeroBattleActionSQL string

//go:embed sql/list_hero_battle_actions_by_battle.sql
var listHeroBattleActionsByBattleSQL string

type Repo struct {
	db  *sql.DB
	log *slog.Logger
}

func NewHeroBattleActionsRepo(db *sql.DB, log *slog.Logger) *Repo { return &Repo{db: db, log: log} }

type CreateArgs struct {
	HeroBattleID      int
	CombatantID       int
	TargetCombatantID *int
	Turn              int
	Action            string
	Damage            int
	Health            int
	Description       string
}

func (r *Repo) CreateContext(ctx context.Context, tx repositories.DbTx, a CreateArgs) (int, error) {
	var id int
	if err := tx.QueryRowContext(ctx, insertHeroBattleActionSQL, a.HeroBattleID, a.CombatantID, a.TargetCombatantID, a.Turn, a.Action, a.Damage, a.Health, a.Description).Scan(&id); err != nil {
		return 0, fmt.Errorf("insert hero_battle_action: %w", err)
	}
	return id, nil
}

type Row struct {
	ID                int
	HeroBattleID      int
	CombatantID       int
	TargetCombatantID sql.NullInt64
	Turn              int
	Action            string
	Damage            int
	Health            int
	Description       string
}

func (r *Repo) ListByBattleContext(ctx context.Context, tx repositories.DbTx, heroBattleID int) ([]Row, error) {
	rows, err := tx.QueryContext(ctx, listHeroBattleActionsByBattleSQL, heroBattleID)
	if err != nil {
		return nil, fmt.Errorf("list hero_battle_actions: %w", err)
	}
	defer rows.Close()
	var out []Row
	for rows.Next() {
		var a Row
		if err := rows.Scan(&a.ID, &a.HeroBattleID, &a.CombatantID, &a.TargetCombatantID, &a.Turn, &a.Action, &a.Damage, &a.Health, &a.Description); err != nil {
			return nil, fmt.Errorf("scan hero_battle_action: %w", err)
		}
		out = append(out, a)
	}
	return out, rows.Err()
}
