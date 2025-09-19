package feedback

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log/slog"

	"github.com/victorprocure/opendominiongo/internal/repositories"
)

//go:embed sql/insert_user_feedback.sql
var insertUserFeedbackSQL string

//go:embed sql/list_user_feedback_for_target.sql
var listUserFeedbackForTargetSQL string

type Repo struct {
	db  *sql.DB
	log *slog.Logger
}

func NewRepo(db *sql.DB, log *slog.Logger) *Repo { return &Repo{db: db, log: log} }

type CreateArgs struct {
	SourceID int
	TargetID int
	Endorsed bool
}

func (r *Repo) CreateContext(ctx context.Context, tx repositories.DbTx, a CreateArgs) (int, error) {
	var id int
	if err := tx.QueryRowContext(ctx, insertUserFeedbackSQL, a.SourceID, a.TargetID, a.Endorsed).Scan(&id); err != nil {
		return 0, fmt.Errorf("insert user_feedback: %w", err)
	}
	return id, nil
}

type Row struct {
	ID       int
	SourceID int
	TargetID int
	Endorsed bool
}

func (r *Repo) ListForTargetContext(ctx context.Context, tx repositories.DbTx, targetID, limit, offset int) ([]Row, error) {
	rows, err := tx.QueryContext(ctx, listUserFeedbackForTargetSQL, targetID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("list user_feedback: %w", err)
	}
	defer rows.Close()
	var out []Row
	for rows.Next() {
		var x Row
		if err := rows.Scan(&x.ID, &x.SourceID, &x.TargetID, &x.Endorsed); err != nil {
			return nil, fmt.Errorf("scan user_feedback: %w", err)
		}
		out = append(out, x)
	}
	return out, rows.Err()
}
