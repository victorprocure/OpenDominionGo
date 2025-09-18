package message_board_posts

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log/slog"

	"github.com/victorprocure/opendominiongo/internal/repositories"
)

//go:embed sql/insert_message_board_post.sql
var insertPostSQL string

//go:embed sql/list_message_board_posts_by_thread.sql
var listPostsByThreadSQL string

type Repo struct {
	db  *sql.DB
	log *slog.Logger
}

func NewMessageBoardPostsRepo(db *sql.DB, log *slog.Logger) *Repo { return &Repo{db: db, log: log} }

type CreateArgs struct {
	ThreadID int
	UserID   int
	Body     string
}

func (r *Repo) CreateContext(ctx context.Context, tx repositories.DbTx, a CreateArgs) (int, error) {
	var id int
	if err := tx.QueryRowContext(ctx, insertPostSQL, a.ThreadID, a.UserID, a.Body).Scan(&id); err != nil {
		return 0, fmt.Errorf("insert message_board_post: %w", err)
	}
	return id, nil
}

type Row struct {
	ID                int
	ThreadID          int
	UserID            int
	Body              string
	FlaggedForRemoval bool
	FlaggedBy         *string
}

func (r *Repo) ListByThreadContext(ctx context.Context, tx repositories.DbTx, threadID int, limit, offset int) ([]Row, error) {
	rows, err := tx.QueryContext(ctx, listPostsByThreadSQL, threadID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("list message_board_posts: %w", err)
	}
	defer rows.Close()
	var out []Row
	for rows.Next() {
		var p Row
		if err := rows.Scan(&p.ID, &p.ThreadID, &p.UserID, &p.Body, &p.FlaggedForRemoval, &p.FlaggedBy); err != nil {
			return nil, fmt.Errorf("scan message_board_post: %w", err)
		}
		out = append(out, p)
	}
	return out, rows.Err()
}
