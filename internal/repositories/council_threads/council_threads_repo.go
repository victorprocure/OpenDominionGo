package council_threads

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log/slog"

	"github.com/victorprocure/opendominiongo/internal/repositories"
)

//go:embed sql/insert_council_thread.sql
var insertCouncilThreadSQL string

//go:embed sql/get_council_thread_by_id.sql
var getCouncilThreadByIDSQL string

//go:embed sql/list_council_threads_by_realm.sql
var listCouncilThreadsByRealmSQL string

//go:embed sql/delete_council_thread.sql
var deleteCouncilThreadSQL string

type Repo struct {
	db  *sql.DB
	log *slog.Logger
}

func NewCouncilThreadsRepo(db *sql.DB, log *slog.Logger) *Repo { return &Repo{db: db, log: log} }

type CreateArgs struct {
	RealmID    int
	DominionID int
	Title      string
	Body       string
}

func (r *Repo) CreateContext(ctx context.Context, tx repositories.DbTx, a CreateArgs) (int, error) {
	var id int
	if err := tx.QueryRowContext(ctx, insertCouncilThreadSQL, a.RealmID, a.DominionID, a.Title, a.Body).Scan(&id); err != nil {
		return 0, fmt.Errorf("insert council_thread: %w", err)
	}
	return id, nil
}

type Row struct {
	ID           int
	RealmID      int
	DominionID   int
	Title        string
	Body         string
	DeletedAt    sql.NullTime
	LastActivity sql.NullTime
}

func (r *Repo) GetByIDContext(ctx context.Context, tx repositories.DbTx, id int) (*Row, error) {
	var t Row
	if err := tx.QueryRowContext(ctx, getCouncilThreadByIDSQL, id).
		Scan(&t.ID, &t.RealmID, &t.DominionID, &t.Title, &t.Body, &t.DeletedAt, &t.LastActivity); err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, fmt.Errorf("get council_thread: %w", err)
	}
	return &t, nil
}

func (r *Repo) ListByRealmContext(ctx context.Context, tx repositories.DbTx, realmID int, limit, offset int) ([]Row, error) {
	rows, err := tx.QueryContext(ctx, listCouncilThreadsByRealmSQL, realmID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("list council_threads: %w", err)
	}
	defer rows.Close()
	var out []Row
	for rows.Next() {
		var t Row
		if err := rows.Scan(&t.ID, &t.RealmID, &t.DominionID, &t.Title, &t.DeletedAt, &t.LastActivity); err != nil {
			return nil, fmt.Errorf("scan council_thread: %w", err)
		}
		out = append(out, t)
	}
	return out, rows.Err()
}

func (r *Repo) DeleteByIDContext(ctx context.Context, tx repositories.DbTx, id int) error {
	if _, err := tx.ExecContext(ctx, deleteCouncilThreadSQL, id); err != nil {
		return fmt.Errorf("delete council_thread: %w", err)
	}
	return nil
}
