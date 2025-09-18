package queue

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log/slog"

	"github.com/victorprocure/opendominiongo/internal/repositories"
)

//go:embed sql/insert_hero_battle_queue.sql
var insertHeroBattleQueueSQL string

//go:embed sql/list_hero_battle_queue.sql
var listHeroBattleQueueSQL string

type Repo struct {
	db  *sql.DB
	log *slog.Logger
}

func NewHeroBattleQueueRepo(db *sql.DB, log *slog.Logger) *Repo { return &Repo{db: db, log: log} }

type CreateArgs struct {
	HeroID int
	Level  int
	Rating int
}

func (r *Repo) CreateContext(ctx context.Context, tx repositories.DbTx, a CreateArgs) (int, error) {
	var id int
	if err := tx.QueryRowContext(ctx, insertHeroBattleQueueSQL, a.HeroID, a.Level, a.Rating).Scan(&id); err != nil {
		return 0, fmt.Errorf("insert hero_battle_queue: %w", err)
	}
	return id, nil
}

type Row struct {
	ID     int
	HeroID int
	Level  int
	Rating int
}

func (r *Repo) ListContext(ctx context.Context, tx repositories.DbTx, limit, offset int) ([]Row, error) {
	rows, err := tx.QueryContext(ctx, listHeroBattleQueueSQL, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("list hero_battle_queue: %w", err)
	}
	defer rows.Close()
	var out []Row
	for rows.Next() {
		var q Row
		if err := rows.Scan(&q.ID, &q.HeroID, &q.Level, &q.Rating); err != nil {
			return nil, fmt.Errorf("scan hero_battle_queue: %w", err)
		}
		out = append(out, q)
	}
	return out, rows.Err()
}
