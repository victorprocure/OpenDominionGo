package realm

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log/slog"

	"github.com/victorprocure/opendominiongo/internal/repositories"
)

//go:embed sql/insert_realm.sql
var insertRealmSQL string

//go:embed sql/get_realm_by_round_and_number.sql
var getRealmByRoundAndNumberSQL string

//go:embed sql/update_realm_name.sql
var updateRealmNameSQL string

//go:embed sql/delete_realm.sql
var deleteRealmSQL string

//go:embed sql/list_realms_by_round.sql
var listRealmsByRoundSQL string

type Repo struct {
	db  *sql.DB
	log *slog.Logger
}

func NewRealmsRepo(db *sql.DB, log *slog.Logger) *Repo {
	return &Repo{db: db, log: log}
}

type CreateArgs struct {
	RoundID   int
	Number    int
	Name      *string // nullable
	Alignment string
}

func (r *Repo) CreateContext(ctx context.Context, tx repositories.DbTx, a CreateArgs) (int, error) {
	var id int
	if err := tx.QueryRowContext(ctx, insertRealmSQL, a.RoundID, a.Number, a.Name, a.Alignment).Scan(&id); err != nil {
		return 0, fmt.Errorf("insert realm: %w", err)
	}
	return id, nil
}

type RealmRow struct {
	ID      int
	RoundID int
	Number  int
	Name    *string
}

func (r *Repo) GetByRoundAndNumberContext(ctx context.Context, tx repositories.DbTx, roundID, number int) (*RealmRow, error) {
	var row RealmRow
	if err := tx.QueryRowContext(ctx, getRealmByRoundAndNumberSQL, roundID, number).Scan(&row.ID, &row.RoundID, &row.Number, &row.Name); err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, fmt.Errorf("get realm by round/number: %w", err)
	}
	return &row, nil
}

func (r *Repo) UpdateNameContext(ctx context.Context, tx repositories.DbTx, id int, name *string) error {
	if _, err := tx.ExecContext(ctx, updateRealmNameSQL, id, name); err != nil {
		return fmt.Errorf("update realm name: %w", err)
	}
	return nil
}

func (r *Repo) DeleteByIDContext(ctx context.Context, tx repositories.DbTx, id int) error {
	if _, err := tx.ExecContext(ctx, deleteRealmSQL, id); err != nil {
		return fmt.Errorf("delete realm: %w", err)
	}
	return nil
}

type ListArgs struct {
	RoundID int
	Limit   int
	Offset  int
}

func (r *Repo) ListByRoundContext(ctx context.Context, tx repositories.DbTx, a ListArgs) ([]RealmRow, error) {
	rows, err := tx.QueryContext(ctx, listRealmsByRoundSQL, a.RoundID, a.Limit, a.Offset)
	if err != nil {
		return nil, fmt.Errorf("list realms: %w", err)
	}
	defer rows.Close()
	var out []RealmRow
	for rows.Next() {
		var rr RealmRow
		if err := rows.Scan(&rr.ID, &rr.RoundID, &rr.Number, &rr.Name); err != nil {
			return nil, fmt.Errorf("scan realm row: %w", err)
		}
		out = append(out, rr)
	}
	return out, rows.Err()
}
