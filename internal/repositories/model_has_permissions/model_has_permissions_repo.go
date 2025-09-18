package model_has_permissions

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log/slog"

	"github.com/victorprocure/opendominiongo/internal/repositories"
)

//go:embed sql/assign_permission.sql
var assignPermissionSQL string

//go:embed sql/revoke_permission.sql
var revokePermissionSQL string

//go:embed sql/list_permissions_for_model.sql
var listPermissionsForModelSQL string

type Repo struct {
	db  *sql.DB
	log *slog.Logger
}

func NewModelHasPermissionsRepo(db *sql.DB, log *slog.Logger) *Repo { return &Repo{db: db, log: log} }

type AssignArgs struct {
	PermissionID int
	ModelType    string
	ModelID      int64
}

func (r *Repo) AssignContext(ctx context.Context, tx repositories.DbTx, a AssignArgs) error {
	if _, err := tx.ExecContext(ctx, assignPermissionSQL, a.PermissionID, a.ModelType, a.ModelID); err != nil {
		return fmt.Errorf("assign permission: %w", err)
	}
	return nil
}

func (r *Repo) RevokeContext(ctx context.Context, tx repositories.DbTx, a AssignArgs) error {
	if _, err := tx.ExecContext(ctx, revokePermissionSQL, a.PermissionID, a.ModelType, a.ModelID); err != nil {
		return fmt.Errorf("revoke permission: %w", err)
	}
	return nil
}

func (r *Repo) ListForModelContext(ctx context.Context, tx repositories.DbTx, modelType string, modelID int64, limit, offset int) ([]int, error) {
	rows, err := tx.QueryContext(ctx, listPermissionsForModelSQL, modelType, modelID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("list permissions for model: %w", err)
	}
	defer rows.Close()
	var out []int
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			return nil, fmt.Errorf("scan permission id: %w", err)
		}
		out = append(out, id)
	}
	return out, rows.Err()
}
