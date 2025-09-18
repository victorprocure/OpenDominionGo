package job

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log/slog"

	"github.com/victorprocure/opendominiongo/internal/repositories"
)

//go:embed sql/insert_job.sql
var insertJobSQL string

//go:embed sql/list_failed_jobs.sql
var listFailedJobsSQL string

type Repo struct {
	db  *sql.DB
	log *slog.Logger
}

func NewJobsRepo(db *sql.DB, log *slog.Logger) *Repo { return &Repo{db: db, log: log} }

type CreateArgs struct {
	Queue   string
	Payload string
}

func (r *Repo) CreateContext(ctx context.Context, tx repositories.DbTx, a CreateArgs) (int64, error) {
	var id int64
	if err := tx.QueryRowContext(ctx, insertJobSQL, a.Queue, a.Payload).Scan(&id); err != nil {
		return 0, fmt.Errorf("insert job: %w", err)
	}
	return id, nil
}

type FailedRow struct {
	ID         int64
	Connection string
	Queue      string
	Payload    string
	Exception  string
}

func (r *Repo) ListFailedContext(ctx context.Context, tx repositories.DbTx, limit, offset int) ([]FailedRow, error) {
	rows, err := tx.QueryContext(ctx, listFailedJobsSQL, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("list failed_jobs: %w", err)
	}
	defer rows.Close()
	var out []FailedRow
	for rows.Next() {
		var f FailedRow
		if err := rows.Scan(&f.ID, &f.Connection, &f.Queue, &f.Payload, &f.Exception); err != nil {
			return nil, fmt.Errorf("scan failed_job: %w", err)
		}
		out = append(out, f)
	}
	return out, rows.Err()
}
