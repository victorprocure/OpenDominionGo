package realm_war

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log/slog"

	"github.com/victorprocure/opendominiongo/internal/repositories"
)

//go:embed sql/insert_realm_war.sql
var insertRealmWarSQL string

//go:embed sql/list_active_realm_wars_for_realm.sql
var listActiveRealmWarsForRealmSQL string

//go:embed sql/end_realm_war.sql
var endRealmWarSQL string

type Repo struct {
	db  *sql.DB
	log *slog.Logger
}

func NewRealmWarsRepo(db *sql.DB, log *slog.Logger) *Repo { return &Repo{db: db, log: log} }

type CreateArgs struct {
	SourceRealmID int
	TargetRealmID int
	ActiveAt      *string
}

func (r *Repo) CreateContext(ctx context.Context, tx repositories.DbTx, a CreateArgs) (int, error) {
	var id int
	if err := tx.QueryRowContext(ctx, insertRealmWarSQL, a.SourceRealmID, a.TargetRealmID, a.ActiveAt).Scan(&id); err != nil {
		return 0, fmt.Errorf("insert realm_war: %w", err)
	}
	return id, nil
}

type Row struct {
	ID            int
	SourceRealmID int
	TargetRealmID int
	ActiveAt      sql.NullTime
	InactiveAt    sql.NullTime
}

func (r *Repo) ListActiveForRealmContext(ctx context.Context, tx repositories.DbTx, realmID int) ([]Row, error) {
	rows, err := tx.QueryContext(ctx, listActiveRealmWarsForRealmSQL, realmID)
	if err != nil {
		return nil, fmt.Errorf("list realm_wars: %w", err)
	}
	defer rows.Close()
	var out []Row
	for rows.Next() {
		var w Row
		if err := rows.Scan(&w.ID, &w.SourceRealmID, &w.TargetRealmID, &w.ActiveAt, &w.InactiveAt); err != nil {
			return nil, fmt.Errorf("scan realm_war: %w", err)
		}
		out = append(out, w)
	}
	return out, rows.Err()
}

func (r *Repo) EndWarContext(ctx context.Context, tx repositories.DbTx, id int) error {
	if _, err := tx.ExecContext(ctx, endRealmWarSQL, id); err != nil {
		return fmt.Errorf("end realm_war: %w", err)
	}
	return nil
}
