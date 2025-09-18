package telescope_monitoring

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log/slog"

	"github.com/victorprocure/opendominiongo/internal/repositories"
)

//go:embed sql/insert_telescope_monitoring_tag.sql
var insertTelescopeMonitoringTagSQL string

//go:embed sql/delete_telescope_monitoring_tag.sql
var deleteTelescopeMonitoringTagSQL string

//go:embed sql/list_telescope_monitoring_tags.sql
var listTelescopeMonitoringTagsSQL string

type Repo struct {
	db  *sql.DB
	log *slog.Logger
}

func NewTelescopeMonitoringRepo(db *sql.DB, log *slog.Logger) *Repo { return &Repo{db: db, log: log} }

func (r *Repo) AddTagContext(ctx context.Context, tx repositories.DbTx, tag string) error {
	if _, err := tx.ExecContext(ctx, insertTelescopeMonitoringTagSQL, tag); err != nil {
		return fmt.Errorf("insert telescope_monitoring: %w", err)
	}
	return nil
}

func (r *Repo) RemoveTagContext(ctx context.Context, tx repositories.DbTx, tag string) error {
	if _, err := tx.ExecContext(ctx, deleteTelescopeMonitoringTagSQL, tag); err != nil {
		return fmt.Errorf("delete telescope_monitoring: %w", err)
	}
	return nil
}

func (r *Repo) ListContext(ctx context.Context, tx repositories.DbTx, limit, offset int) ([]string, error) {
	rows, err := tx.QueryContext(ctx, listTelescopeMonitoringTagsSQL, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("list telescope_monitoring: %w", err)
	}
	defer rows.Close()
	var out []string
	for rows.Next() {
		var tag string
		if err := rows.Scan(&tag); err != nil {
			return nil, fmt.Errorf("scan telescope_monitoring: %w", err)
		}
		out = append(out, tag)
	}
	return out, rows.Err()
}
