package combatant

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log/slog"

	"github.com/victorprocure/opendominiongo/internal/repositories"
)

//go:embed sql/insert_hero_combatant.sql
var insertHeroCombatantSQL string

//go:embed sql/list_hero_combatants_by_battle.sql
var listHeroCombatantsByBattleSQL string

type Repo struct {
	db  *sql.DB
	log *slog.Logger
}

func NewHeroCombatantsRepo(db *sql.DB, log *slog.Logger) *Repo { return &Repo{db: db, log: log} }

type CreateArgs struct {
	HeroBattleID int
	HeroID       *int
	DominionID   *int
	Name         string
	Health       int
	Attack       int
	Defense      int
	Evasion      int
	Focus        int
	Counter      int
	Recover      int
	Level        int
}

func (r *Repo) CreateContext(ctx context.Context, tx repositories.DbTx, a CreateArgs) (int, error) {
	var id int
	if err := tx.QueryRowContext(ctx, insertHeroCombatantSQL, a.HeroBattleID, a.HeroID, a.DominionID, a.Name, a.Health, a.Attack, a.Defense, a.Evasion, a.Focus, a.Counter, a.Recover, a.Level).Scan(&id); err != nil {
		return 0, fmt.Errorf("insert hero_combatant: %w", err)
	}
	return id, nil
}

type Row struct {
	ID            int
	HeroBattleID  int
	HeroID        sql.NullInt64
	DominionID    sql.NullInt64
	Name          string
	Health        int
	Attack        int
	Defense       int
	Evasion       int
	Focus         int
	Counter       int
	Recover       int
	CurrentHealth int
	Level         int
}

func (r *Repo) ListByBattleContext(ctx context.Context, tx repositories.DbTx, heroBattleID int) ([]Row, error) {
	rows, err := tx.QueryContext(ctx, listHeroCombatantsByBattleSQL, heroBattleID)
	if err != nil {
		return nil, fmt.Errorf("list hero_combatants: %w", err)
	}
	defer rows.Close()
	var out []Row
	for rows.Next() {
		var c Row
		if err := rows.Scan(&c.ID, &c.HeroBattleID, &c.HeroID, &c.DominionID, &c.Name, &c.Health, &c.Attack, &c.Defense, &c.Evasion, &c.Focus, &c.Counter, &c.Recover, &c.CurrentHealth, &c.Level); err != nil {
			return nil, fmt.Errorf("scan hero_combatant: %w", err)
		}
		out = append(out, c)
	}
	return out, rows.Err()
}
