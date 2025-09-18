package tick

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log/slog"

	"github.com/victorprocure/opendominiongo/internal/repositories"
)

//go:embed sql/get_latest_dominion_tick.sql
var getLatestDominionTickSQL string

type Repo struct {
	db  *sql.DB
	log *slog.Logger
}

func NewDominionTickRepo(db *sql.DB, log *slog.Logger) *Repo { return &Repo{db: db, log: log} }

type Row struct {
	ID         int
	DominionID int
	Prestige   int
	Peasants   int
	Morale     int
	UpdatedAt  sql.NullTime
}

func (r *Repo) GetLatestByDominionContext(ctx context.Context, tx repositories.DbTx, dominionID int) (*Row, error) {
	var t Row
	if err := tx.QueryRowContext(ctx, getLatestDominionTickSQL, dominionID).Scan(&t.ID, &t.DominionID, &t.Prestige, &t.Peasants, &t.Morale, &t.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, fmt.Errorf("get latest dominion_tick: %w", err)
	}
	return &t, nil
}
