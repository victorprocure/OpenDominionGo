package permission

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log/slog"

	"github.com/victorprocure/opendominiongo/internal/repositories"
)

//go:embed sql/assign_permission_to_role.sql
var assignPermissionToRoleSQL string

//go:embed sql/revoke_permission_from_role.sql
var revokePermissionFromRoleSQL string

//go:embed sql/list_permissions_for_role.sql
var listPermissionsForRoleSQL string

type Repo struct {
	db  *sql.DB
	log *slog.Logger
}

func NewRoleHasPermissionRepo(db *sql.DB, log *slog.Logger) *Repo { return &Repo{db: db, log: log} }

type AssignArgs struct {
	RoleID       int
	PermissionID int
}

func (r *Repo) AssignContext(ctx context.Context, tx repositories.DbTx, a AssignArgs) error {
	if _, err := tx.ExecContext(ctx, assignPermissionToRoleSQL, a.RoleID, a.PermissionID); err != nil {
		return fmt.Errorf("assign permission to role: %w", err)
	}
	return nil
}

func (r *Repo) RevokeContext(ctx context.Context, tx repositories.DbTx, a AssignArgs) error {
	if _, err := tx.ExecContext(ctx, revokePermissionFromRoleSQL, a.RoleID, a.PermissionID); err != nil {
		return fmt.Errorf("revoke permission from role: %w", err)
	}
	return nil
}

func (r *Repo) ListForRoleContext(ctx context.Context, tx repositories.DbTx, roleID int, limit, offset int) ([]int, error) {
	rows, err := tx.QueryContext(ctx, listPermissionsForRoleSQL, roleID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("list permissions for role: %w", err)
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
