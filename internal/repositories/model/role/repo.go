package role

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log/slog"

	"github.com/victorprocure/opendominiongo/internal/repositories"
)

//go:embed sql/assign_role.sql
var assignRoleSQL string

//go:embed sql/revoke_role.sql
var revokeRoleSQL string

//go:embed sql/list_roles_for_model.sql
var listRolesForModelSQL string

type Repo struct {
	db  *sql.DB
	log *slog.Logger
}

func NewRepo(db *sql.DB, log *slog.Logger) *Repo { return &Repo{db: db, log: log} }

type AssignArgs struct {
	RoleID    int
	ModelType string
	ModelID   int64
}

func (r *Repo) AssignContext(ctx context.Context, tx repositories.DbTx, a AssignArgs) error {
	if _, err := tx.ExecContext(ctx, assignRoleSQL, a.RoleID, a.ModelType, a.ModelID); err != nil {
		return fmt.Errorf("assign role: %w", err)
	}
	return nil
}

func (r *Repo) RevokeContext(ctx context.Context, tx repositories.DbTx, a AssignArgs) error {
	if _, err := tx.ExecContext(ctx, revokeRoleSQL, a.RoleID, a.ModelType, a.ModelID); err != nil {
		return fmt.Errorf("revoke role: %w", err)
	}
	return nil
}

func (r *Repo) ListForModelContext(ctx context.Context, tx repositories.DbTx, modelType string, modelID int64, limit, offset int) ([]int, error) {
	rows, err := tx.QueryContext(ctx, listRolesForModelSQL, modelType, modelID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("list roles for model: %w", err)
	}
	defer rows.Close()
	var out []int
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			return nil, fmt.Errorf("scan role id: %w", err)
		}
		out = append(out, id)
	}
	return out, rows.Err()
}
