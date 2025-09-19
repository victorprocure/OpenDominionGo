package queue

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log/slog"

	"github.com/victorprocure/opendominiongo/internal/repositories"
)

//go:embed sql/insert_dominion_queue.sql
var insertDominionQueueSQL string

//go:embed sql/list_dominion_queue_by_dominion.sql
var listDominionQueueByDominionSQL string

type Repo struct {
	db  *sql.DB
	log *slog.Logger
}

func NewRepo(db *sql.DB, log *slog.Logger) *Repo { return &Repo{db: db, log: log} }

type CreateArgs struct {
	DominionID int
	Source     string
	Resource   string
	Hours      int
	Amount     int
}

func (r *Repo) CreateContext(ctx context.Context, tx repositories.DbTx, a CreateArgs) error {
	if _, err := tx.ExecContext(ctx, insertDominionQueueSQL, a.DominionID, a.Source, a.Resource, a.Hours, a.Amount); err != nil {
		return fmt.Errorf("insert dominion_queue: %w", err)
	}
	return nil
}

type Row struct {
	DominionID int
	Source     string
	Resource   string
	Hours      int
	Amount     int
}

func (r *Repo) ListByDominionContext(ctx context.Context, tx repositories.DbTx, dominionID int) ([]Row, error) {
	rows, err := tx.QueryContext(ctx, listDominionQueueByDominionSQL, dominionID)
	if err != nil {
		return nil, fmt.Errorf("list dominion_queue: %w", err)
	}
	defer rows.Close()
	var out []Row
	for rows.Next() {
		var q Row
		if err := rows.Scan(&q.DominionID, &q.Source, &q.Resource, &q.Hours, &q.Amount); err != nil {
			return nil, fmt.Errorf("scan dominion_queue: %w", err)
		}
		out = append(out, q)
	}
	return out, rows.Err()
}
