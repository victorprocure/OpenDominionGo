package round_league

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log/slog"

	"github.com/victorprocure/opendominiongo/internal/repositories"
)

//go:embed sql/upsert_round_league.sql
var upsertRoundLeagueSQL string

//go:embed sql/get_round_league_by_key.sql
var getRoundLeagueByKeySQL string

//go:embed sql/list_round_leagues.sql
var listRoundLeaguesSQL string

//go:embed sql/delete_round_league.sql
var deleteRoundLeagueSQL string

type Repo struct {
	db  *sql.DB
	log *slog.Logger
}

func NewRoundLeaguesRepo(db *sql.DB, log *slog.Logger) *Repo { return &Repo{db: db, log: log} }

type UpsertArgs struct {
	Key         string
	Description string
}

func (r *Repo) UpsertContext(ctx context.Context, tx repositories.DbTx, a UpsertArgs) (int, error) {
	var id int
	if err := tx.QueryRowContext(ctx, upsertRoundLeagueSQL, a.Key, a.Description).Scan(&id); err != nil {
		return 0, fmt.Errorf("upsert round league: %w", err)
	}
	return id, nil
}

type Row struct {
	ID          int
	Key         string
	Description string
}

func (r *Repo) GetByKeyContext(ctx context.Context, tx repositories.DbTx, key string) (*Row, error) {
	var rl Row
	if err := tx.QueryRowContext(ctx, getRoundLeagueByKeySQL, key).Scan(&rl.ID, &rl.Key, &rl.Description); err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, fmt.Errorf("get round league: %w", err)
	}
	return &rl, nil
}

func (r *Repo) ListContext(ctx context.Context, tx repositories.DbTx) ([]Row, error) {
	rows, err := tx.QueryContext(ctx, listRoundLeaguesSQL)
	if err != nil {
		return nil, fmt.Errorf("list round leagues: %w", err)
	}
	defer rows.Close()
	var out []Row
	for rows.Next() {
		var rl Row
		if err := rows.Scan(&rl.ID, &rl.Key, &rl.Description); err != nil {
			return nil, fmt.Errorf("scan round league: %w", err)
		}
		out = append(out, rl)
	}
	return out, rows.Err()
}

func (r *Repo) DeleteByIDContext(ctx context.Context, tx repositories.DbTx, id int) error {
	if _, err := tx.ExecContext(ctx, deleteRoundLeagueSQL, id); err != nil {
		return fmt.Errorf("delete round league: %w", err)
	}
	return nil
}
