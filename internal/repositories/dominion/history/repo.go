package history

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log/slog"

	"github.com/victorprocure/opendominiongo/internal/repositories"
)

//go:embed sql/insert_dominion_history.sql
var insertDominionHistorySQL string

//go:embed sql/list_dominion_history_by_dominion.sql
var listDominionHistoryByDominionSQL string

type Repo struct {
	db  *sql.DB
	log *slog.Logger
}

func NewDominionHistoryRepo(db *sql.DB, log *slog.Logger) *Repo { return &Repo{db: db, log: log} }

type CreateArgs struct {
	DominionID int
	Event      string
	Delta      string
	IP         string
	Device     string
}

func (r *Repo) CreateContext(ctx context.Context, tx repositories.DbTx, a CreateArgs) (int, error) {
	var id int
	if err := tx.QueryRowContext(ctx, insertDominionHistorySQL, a.DominionID, a.Event, a.Delta, a.IP, a.Device).Scan(&id); err != nil {
		return 0, fmt.Errorf("insert dominion_history: %w", err)
	}
	return id, nil
}

type Row struct {
	ID         int
	DominionID int
	Event      string
	Delta      string
	CreatedAt  sql.NullTime
	IP         string
	Device     string
}

func (r *Repo) ListByDominionContext(ctx context.Context, tx repositories.DbTx, dominionID, limit, offset int) ([]Row, error) {
	rows, err := tx.QueryContext(ctx, listDominionHistoryByDominionSQL, dominionID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("list dominion_history: %w", err)
	}
	defer rows.Close()
	var out []Row
	for rows.Next() {
		var h Row
		if err := rows.Scan(&h.ID, &h.DominionID, &h.Event, &h.Delta, &h.CreatedAt, &h.IP, &h.Device); err != nil {
			return nil, fmt.Errorf("scan dominion_history: %w", err)
		}
		out = append(out, h)
	}
	return out, rows.Err()
}
