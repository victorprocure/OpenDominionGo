package dominion

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log/slog"

	"github.com/victorprocure/opendominiongo/internal/repositories"
)

//go:embed sql/get_dominion_by_round_and_name.sql
var getDominionByRoundAndNameSQL string

//go:embed sql/list_dominions_by_realm.sql
var listDominionsByRealmSQL string

//go:embed sql/update_dominion_name.sql
var updateDominionNameSQL string

type Repo struct {
	db  *sql.DB
	log *slog.Logger
}

func NewDominionRepo(db *sql.DB, log *slog.Logger) *Repo { return &Repo{db: db, log: log} }

type Row struct {
	ID      int
	RoundID int
	RealmID int
	RaceID  int
	Name    string
}

func (r *Repo) GetByRoundAndNameContext(ctx context.Context, tx repositories.DbTx, roundID int, name string) (*Row, error) {
	var d Row
	if err := tx.QueryRowContext(ctx, getDominionByRoundAndNameSQL, roundID, name).
		Scan(&d.ID, &d.RoundID, &d.RealmID, &d.RaceID, &d.Name); err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, fmt.Errorf("get dominion by round+name: %w", err)
	}
	return &d, nil
}

func (r *Repo) ListByRealmContext(ctx context.Context, tx repositories.DbTx, realmID int) ([]Row, error) {
	rows, err := tx.QueryContext(ctx, listDominionsByRealmSQL, realmID)
	if err != nil {
		return nil, fmt.Errorf("list dominions by realm: %w", err)
	}
	defer rows.Close()
	var out []Row
	for rows.Next() {
		var d Row
		if err := rows.Scan(&d.ID, &d.RoundID, &d.RealmID, &d.RaceID, &d.Name); err != nil {
			return nil, fmt.Errorf("scan dominion: %w", err)
		}
		out = append(out, d)
	}
	return out, rows.Err()
}

func (r *Repo) UpdateNameByIDContext(ctx context.Context, tx repositories.DbTx, id int, name string) error {
	if _, err := tx.ExecContext(ctx, updateDominionNameSQL, id, name); err != nil {
		return fmt.Errorf("update dominion name: %w", err)
	}
	return nil
}
