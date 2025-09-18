package realm_history

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log/slog"

	"github.com/victorprocure/opendominiongo/internal/repositories"
)

//go:embed sql/insert_realm_history.sql
var insertRealmHistorySQL string

//go:embed sql/list_realm_history_by_realm.sql
var listRealmHistoryByRealmSQL string

type Repo struct {
	db  *sql.DB
	log *slog.Logger
}

func NewRealmHistoryRepo(db *sql.DB, log *slog.Logger) *Repo { return &Repo{db: db, log: log} }

type CreateArgs struct {
	RealmID    int
	DominionID int
	Event      string
	Delta      string
}

func (r *Repo) CreateContext(ctx context.Context, tx repositories.DbTx, a CreateArgs) (int, error) {
	var id int
	if err := tx.QueryRowContext(ctx, insertRealmHistorySQL, a.RealmID, a.DominionID, a.Event, a.Delta).Scan(&id); err != nil {
		return 0, fmt.Errorf("insert realm_history: %w", err)
	}
	return id, nil
}

type Row struct {
	ID         int
	RealmID    int
	DominionID int
	Event      string
	Delta      string
	CreatedAt  sql.NullTime
}

func (r *Repo) ListByRealmContext(ctx context.Context, tx repositories.DbTx, realmID int, limit, offset int) ([]Row, error) {
	rows, err := tx.QueryContext(ctx, listRealmHistoryByRealmSQL, realmID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("list realm_history: %w", err)
	}
	defer rows.Close()
	var out []Row
	for rows.Next() {
		var h Row
		if err := rows.Scan(&h.ID, &h.RealmID, &h.DominionID, &h.Event, &h.Delta, &h.CreatedAt); err != nil {
			return nil, fmt.Errorf("scan realm_history: %w", err)
		}
		out = append(out, h)
	}
	return out, rows.Err()
}
