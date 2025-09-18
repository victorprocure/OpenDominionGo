package game_events

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log/slog"

	"github.com/victorprocure/opendominiongo/internal/repositories"
)

//go:embed sql/insert_game_event.sql
var insertGameEventSQL string

//go:embed sql/list_game_events_by_round.sql
var listGameEventsByRoundSQL string

type Repo struct {
	db  *sql.DB
	log *slog.Logger
}

func NewGameEventsRepo(db *sql.DB, log *slog.Logger) *Repo { return &Repo{db: db, log: log} }

type CreateArgs struct {
	RoundID    int
	SourceType string
	SourceID   int64
	TargetType *string
	TargetID   *int64
	Type       string
	Data       *string
}

func (r *Repo) CreateContext(ctx context.Context, tx repositories.DbTx, a CreateArgs) error {
	var id string
	if err := tx.QueryRowContext(ctx, insertGameEventSQL, a.RoundID, a.SourceType, a.SourceID, a.TargetType, a.TargetID, a.Type, a.Data).Scan(&id); err != nil {
		return fmt.Errorf("insert game_event: %w", err)
	}
	return nil
}

type Row struct {
	ID         string
	RoundID    int
	SourceType string
	SourceID   int64
	TargetType *string
	TargetID   *int64
	Type       string
	Data       *string
}

func (r *Repo) ListByRoundContext(ctx context.Context, tx repositories.DbTx, roundID int, limit, offset int) ([]Row, error) {
	rows, err := tx.QueryContext(ctx, listGameEventsByRoundSQL, roundID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("list game_events: %w", err)
	}
	defer rows.Close()
	var out []Row
	for rows.Next() {
		var e Row
		if err := rows.Scan(&e.ID, &e.RoundID, &e.SourceType, &e.SourceID, &e.TargetType, &e.TargetID, &e.Type, &e.Data); err != nil {
			return nil, fmt.Errorf("scan game_event: %w", err)
		}
		out = append(out, e)
	}
	return out, rows.Err()
}
