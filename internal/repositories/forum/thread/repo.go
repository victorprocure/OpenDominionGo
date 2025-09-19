package thread

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log/slog"

	"github.com/victorprocure/opendominiongo/internal/repositories"
)

//go:embed sql/insert_forum_thread.sql
var insertForumThreadSQL string

//go:embed sql/get_forum_thread_by_id.sql
var getForumThreadByIDSQL string

//go:embed sql/list_forum_threads_by_round.sql
var listForumThreadsByRoundSQL string

//go:embed sql/delete_forum_thread.sql
var deleteForumThreadSQL string

type Repo struct {
	db  *sql.DB
	log *slog.Logger
}

func NewRepo(db *sql.DB, log *slog.Logger) *Repo { return &Repo{db: db, log: log} }

type CreateArgs struct {
	RoundID    int
	DominionID int
	Title      string
	Body       string
}

func (r *Repo) CreateContext(ctx context.Context, tx repositories.DbTx, a CreateArgs) (int, error) {
	var id int
	if err := tx.QueryRowContext(ctx, insertForumThreadSQL, a.RoundID, a.DominionID, a.Title, a.Body).Scan(&id); err != nil {
		return 0, fmt.Errorf("insert forum thread: %w", err)
	}
	return id, nil
}

type Row struct {
	ID                int
	RoundID           int
	DominionID        int
	Title             string
	Body              string
	FlaggedForRemoval bool
	FlaggedBy         *string
	LastActivity      sql.NullTime
}

func (r *Repo) GetByIDContext(ctx context.Context, tx repositories.DbTx, id int) (*Row, error) {
	var ft Row
	if err := tx.QueryRowContext(ctx, getForumThreadByIDSQL, id).
		Scan(&ft.ID, &ft.RoundID, &ft.DominionID, &ft.Title, &ft.Body, &ft.FlaggedForRemoval, &ft.FlaggedBy, &ft.LastActivity); err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, fmt.Errorf("get forum thread: %w", err)
	}
	return &ft, nil
}

func (r *Repo) ListByRoundContext(ctx context.Context, tx repositories.DbTx, roundID int, limit, offset int) ([]Row, error) {
	rows, err := tx.QueryContext(ctx, listForumThreadsByRoundSQL, roundID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("list forum threads: %w", err)
	}
	defer rows.Close()
	var out []Row
	for rows.Next() {
		var ft Row
		if err := rows.Scan(&ft.ID, &ft.RoundID, &ft.DominionID, &ft.Title, &ft.FlaggedForRemoval, &ft.FlaggedBy, &ft.LastActivity); err != nil {
			return nil, fmt.Errorf("scan forum thread: %w", err)
		}
		out = append(out, ft)
	}
	return out, rows.Err()
}

func (r *Repo) DeleteByIDContext(ctx context.Context, tx repositories.DbTx, id int) error {
	if _, err := tx.ExecContext(ctx, deleteForumThreadSQL, id); err != nil {
		return fmt.Errorf("delete forum thread: %w", err)
	}
	return nil
}
