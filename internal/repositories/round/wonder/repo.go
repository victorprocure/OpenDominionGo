package wonder

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log/slog"

	"github.com/victorprocure/opendominiongo/internal/repositories"
)

//go:embed sql/insert_round_wonder.sql
var insertRoundWonderSQL string

//go:embed sql/update_round_wonder_power.sql
var updateRoundWonderPowerSQL string

//go:embed sql/list_round_wonders_by_round.sql
var listRoundWondersByRoundSQL string

//go:embed sql/delete_round_wonder.sql
var deleteRoundWonderSQL string

type Repo struct {
	db  *sql.DB
	log *slog.Logger
}

func NewRepo(db *sql.DB, log *slog.Logger) *Repo {
	return &Repo{db: db, log: log}
}

type InsertArgs struct {
	RoundID  int
	RealmID  *int // nullable
	WonderID int
	Power    int
}

func (r *Repo) InsertContext(ctx context.Context, tx repositories.DbTx, a InsertArgs) (int, error) {
	var id int
	if err := tx.QueryRowContext(ctx, insertRoundWonderSQL, a.RoundID, a.RealmID, a.WonderID, a.Power).Scan(&id); err != nil {
		return 0, fmt.Errorf("insert round_wonder: %w", err)
	}
	return id, nil
}

func (r *Repo) UpdatePowerContext(ctx context.Context, tx repositories.DbTx, id, power int) error {
	if _, err := tx.ExecContext(ctx, updateRoundWonderPowerSQL, id, power); err != nil {
		return fmt.Errorf("update round_wonder power: %w", err)
	}
	return nil
}

type Row struct {
	ID       int
	RoundID  int
	RealmID  *int
	WonderID int
	Power    int
}

func (r *Repo) ListByRoundContext(ctx context.Context, tx repositories.DbTx, roundID int) ([]Row, error) {
	rows, err := tx.QueryContext(ctx, listRoundWondersByRoundSQL, roundID)
	if err != nil {
		return nil, fmt.Errorf("list round_wonders: %w", err)
	}
	defer rows.Close()
	var out []Row
	for rows.Next() {
		var rw Row
		if err := rows.Scan(&rw.ID, &rw.RoundID, &rw.RealmID, &rw.WonderID, &rw.Power); err != nil {
			return nil, fmt.Errorf("scan round_wonder: %w", err)
		}
		out = append(out, rw)
	}
	return out, rows.Err()
}

func (r *Repo) DeleteByIDContext(ctx context.Context, tx repositories.DbTx, id int) error {
	if _, err := tx.ExecContext(ctx, deleteRoundWonderSQL, id); err != nil {
		return fmt.Errorf("delete round_wonder: %w", err)
	}
	return nil
}
