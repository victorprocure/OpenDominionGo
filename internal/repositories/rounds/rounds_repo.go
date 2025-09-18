package round

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log/slog"

	"github.com/victorprocure/opendominiongo/internal/repositories"
)

//go:embed sql/insert_round.sql
var insertRoundSQL string

//go:embed sql/get_round_by_number.sql
var getRoundByNumberSQL string

//go:embed sql/update_round.sql
var updateRoundSQL string

//go:embed sql/delete_round.sql
var deleteRoundSQL string

//go:embed sql/list_rounds.sql
var listRoundsSQL string

type Repo struct {
	db  *sql.DB
	log *slog.Logger
}

func NewRoundsRepo(db *sql.DB, log *slog.Logger) *Repo {
	return &Repo{db: db, log: log}
}

type CreateArgs struct {
	RoundLeagueID int
	Number        int
	Name          string
	StartDate     string // ISO8601 string to avoid time zone issues at this layer
	EndDate       string
}

func (r *Repo) CreateContext(ctx context.Context, tx repositories.DbTx, a CreateArgs) (int, error) {
	var id int
	if err := tx.QueryRowContext(ctx, insertRoundSQL,
		a.RoundLeagueID, a.Number, a.Name, a.StartDate, a.EndDate,
	).Scan(&id); err != nil {
		return 0, fmt.Errorf("insert round: %w", err)
	}
	return id, nil
}

type RoundRow struct {
	ID     int
	Number int
	Name   string
}

func (r *Repo) GetByNumberContext(ctx context.Context, tx repositories.DbTx, number int) (*RoundRow, error) {
	var row RoundRow
	if err := tx.QueryRowContext(ctx, getRoundByNumberSQL, number).Scan(&row.ID, &row.Number, &row.Name); err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, fmt.Errorf("get round by number: %w", err)
	}
	return &row, nil
}

func (r *Repo) UpdateNameContext(ctx context.Context, tx repositories.DbTx, id int, name string) error {
	if _, err := tx.ExecContext(ctx, updateRoundSQL, id, name); err != nil {
		return fmt.Errorf("update round name: %w", err)
	}
	return nil
}

func (r *Repo) DeleteByIDContext(ctx context.Context, tx repositories.DbTx, id int) error {
	if _, err := tx.ExecContext(ctx, deleteRoundSQL, id); err != nil {
		return fmt.Errorf("delete round: %w", err)
	}
	return nil
}

type ListArgs struct {
	Limit  int
	Offset int
}

func (r *Repo) ListContext(ctx context.Context, tx repositories.DbTx, a ListArgs) ([]RoundRow, error) {
	rows, err := tx.QueryContext(ctx, listRoundsSQL, a.Limit, a.Offset)
	if err != nil {
		return nil, fmt.Errorf("list rounds: %w", err)
	}
	defer rows.Close()
	var out []RoundRow
	for rows.Next() {
		var rr RoundRow
		if err := rows.Scan(&rr.ID, &rr.Number, &rr.Name); err != nil {
			return nil, fmt.Errorf("scan round row: %w", err)
		}
		out = append(out, rr)
	}
	return out, rows.Err()
}
