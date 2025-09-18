package council_post

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log/slog"

	"github.com/victorprocure/opendominiongo/internal/repositories"
)

//go:embed sql/insert_council_post.sql
var insertCouncilPostSQL string

//go:embed sql/list_council_posts_by_thread.sql
var listCouncilPostsByThreadSQL string

type Repo struct {
	db  *sql.DB
	log *slog.Logger
}

func NewCouncilPostsRepo(db *sql.DB, log *slog.Logger) *Repo { return &Repo{db: db, log: log} }

type CreateArgs struct {
	CouncilThreadID int
	DominionID      int
	Body            string
}

func (r *Repo) CreateContext(ctx context.Context, tx repositories.DbTx, a CreateArgs) (int, error) {
	var id int
	if err := tx.QueryRowContext(ctx, insertCouncilPostSQL, a.CouncilThreadID, a.DominionID, a.Body).Scan(&id); err != nil {
		return 0, fmt.Errorf("insert council_post: %w", err)
	}
	return id, nil
}

type Row struct {
	ID              int
	CouncilThreadID int
	DominionID      int
	Body            string
	DeletedAt       sql.NullTime
}

func (r *Repo) ListByThreadContext(ctx context.Context, tx repositories.DbTx, threadID int, limit, offset int) ([]Row, error) {
	rows, err := tx.QueryContext(ctx, listCouncilPostsByThreadSQL, threadID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("list council_posts: %w", err)
	}
	defer rows.Close()
	var out []Row
	for rows.Next() {
		var p Row
		if err := rows.Scan(&p.ID, &p.CouncilThreadID, &p.DominionID, &p.Body, &p.DeletedAt); err != nil {
			return nil, fmt.Errorf("scan council_post: %w", err)
		}
		out = append(out, p)
	}
	return out, rows.Err()
}
