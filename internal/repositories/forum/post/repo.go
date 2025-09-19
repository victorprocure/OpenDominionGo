package post

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log/slog"

	"github.com/victorprocure/opendominiongo/internal/repositories"
)

//go:embed sql/insert_forum_post.sql
var insertForumPostSQL string

//go:embed sql/list_forum_posts_by_thread.sql
var listForumPostsByThreadSQL string

type Repo struct {
	db  *sql.DB
	log *slog.Logger
}

func NewRepo(db *sql.DB, log *slog.Logger) *Repo { return &Repo{db: db, log: log} }

type CreateArgs struct {
	ForumThreadID int
	DominionID    int
	Body          string
}

func (r *Repo) CreateContext(ctx context.Context, tx repositories.DbTx, a CreateArgs) (int, error) {
	var id int
	if err := tx.QueryRowContext(ctx, insertForumPostSQL, a.ForumThreadID, a.DominionID, a.Body).Scan(&id); err != nil {
		return 0, fmt.Errorf("insert forum post: %w", err)
	}
	return id, nil
}

type Row struct {
	ID                int
	ForumThreadID     int
	DominionID        int
	Body              string
	FlaggedForRemoval bool
	FlaggedBy         *string
}

func (r *Repo) ListByThreadContext(ctx context.Context, tx repositories.DbTx, threadID, limit, offset int) ([]Row, error) {
	rows, err := tx.QueryContext(ctx, listForumPostsByThreadSQL, threadID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("list forum posts: %w", err)
	}
	defer rows.Close()
	var out []Row
	for rows.Next() {
		var fp Row
		if err := rows.Scan(&fp.ID, &fp.ForumThreadID, &fp.DominionID, &fp.Body, &fp.FlaggedForRemoval, &fp.FlaggedBy); err != nil {
			return nil, fmt.Errorf("scan forum post: %w", err)
		}
		out = append(out, fp)
	}
	return out, rows.Err()
}
