package thread

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log/slog"

	"github.com/victorprocure/opendominiongo/internal/repositories"
)

//go:embed sql/insert_message_board_thread.sql
var insertThreadSQL string

//go:embed sql/get_message_board_thread_by_id.sql
var getThreadByIDSQL string

//go:embed sql/list_message_board_threads_by_category.sql
var listThreadsByCategorySQL string

//go:embed sql/delete_message_board_thread.sql
var deleteThreadSQL string

type Repo struct {
	db  *sql.DB
	log *slog.Logger
}

func NewRepo(db *sql.DB, log *slog.Logger) *Repo { return &Repo{db: db, log: log} }

type CreateArgs struct {
	CategoryID int
	UserID     int
	Title      string
	Body       string
}

func (r *Repo) CreateContext(ctx context.Context, tx repositories.DbTx, a CreateArgs) (int, error) {
	var id int
	if err := tx.QueryRowContext(ctx, insertThreadSQL, a.CategoryID, a.UserID, a.Title, a.Body).Scan(&id); err != nil {
		return 0, fmt.Errorf("insert message_board_thread: %w", err)
	}
	return id, nil
}

type Row struct {
	ID                int
	CategoryID        int
	UserID            int
	Title             string
	Body              string
	FlaggedForRemoval bool
	FlaggedBy         *string
	LastActivity      sql.NullTime
}

func (r *Repo) GetByIDContext(ctx context.Context, tx repositories.DbTx, id int) (*Row, error) {
	var t Row
	if err := tx.QueryRowContext(ctx, getThreadByIDSQL, id).Scan(&t.ID, &t.CategoryID, &t.UserID, &t.Title, &t.Body, &t.FlaggedForRemoval, &t.FlaggedBy, &t.LastActivity); err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, fmt.Errorf("get message_board_thread: %w", err)
	}
	return &t, nil
}

func (r *Repo) ListByCategoryContext(ctx context.Context, tx repositories.DbTx, categoryID int) ([]Row, error) {
	rows, err := tx.QueryContext(ctx, listThreadsByCategorySQL, categoryID)
	if err != nil {
		return nil, fmt.Errorf("list message_board_threads: %w", err)
	}
	defer rows.Close()
	var out []Row
	for rows.Next() {
		var t Row
		if err := rows.Scan(&t.ID, &t.CategoryID, &t.UserID, &t.Title, &t.FlaggedForRemoval, &t.FlaggedBy, &t.LastActivity); err != nil {
			return nil, fmt.Errorf("scan message_board_thread: %w", err)
		}
		out = append(out, t)
	}
	return out, rows.Err()
}

func (r *Repo) DeleteByIDContext(ctx context.Context, tx repositories.DbTx, id int) error {
	if _, err := tx.ExecContext(ctx, deleteThreadSQL, id); err != nil {
		return fmt.Errorf("delete message_board_thread: %w", err)
	}
	return nil
}
