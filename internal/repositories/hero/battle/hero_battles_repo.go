package battle

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log/slog"

	"github.com/victorprocure/opendominiongo/internal/repositories"
)

//go:embed sql/insert_hero_battle.sql
var insertHeroBattleSQL string

//go:embed sql/get_hero_battle_by_id.sql
var getHeroBattleByIDSQL string

//go:embed sql/finish_hero_battle.sql
var finishHeroBattleSQL string

type Repo struct {
	db  *sql.DB
	log *slog.Logger
}

func NewHeroBattlesRepo(db *sql.DB, log *slog.Logger) *Repo { return &Repo{db: db, log: log} }

type CreateArgs struct {
	RoundID     *int
	CurrentTurn int
	PVP         bool
}

func (r *Repo) CreateContext(ctx context.Context, tx repositories.DbTx, a CreateArgs) (int, error) {
	var id int
	if err := tx.QueryRowContext(ctx, insertHeroBattleSQL, a.RoundID, a.CurrentTurn, a.PVP).Scan(&id); err != nil {
		return 0, fmt.Errorf("insert hero_battle: %w", err)
	}
	return id, nil
}

type Row struct {
	ID                int
	RoundID           sql.NullInt64
	CurrentTurn       int
	WinnerCombatantID sql.NullInt64
	Finished          bool
	LastProcessedAt   sql.NullTime
	PVP               bool
}

func (r *Repo) GetByIDContext(ctx context.Context, tx repositories.DbTx, id int) (*Row, error) {
	var b Row
	if err := tx.QueryRowContext(ctx, getHeroBattleByIDSQL, id).
		Scan(&b.ID, &b.RoundID, &b.CurrentTurn, &b.WinnerCombatantID, &b.Finished, &b.LastProcessedAt, &b.PVP); err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, fmt.Errorf("get hero_battle: %w", err)
	}
	return &b, nil
}

func (r *Repo) FinishContext(ctx context.Context, tx repositories.DbTx, id int, winnerCombatantID *int) error {
	if _, err := tx.ExecContext(ctx, finishHeroBattleSQL, id, winnerCombatantID); err != nil {
		return fmt.Errorf("finish hero_battle: %w", err)
	}
	return nil
}
