package permission

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log/slog"

	"github.com/victorprocure/opendominiongo/internal/repositories"
)

//go:embed sql/insert_permission.sql
var insertPermissionSQL string

//go:embed sql/get_permission_by_name.sql
var getPermissionByNameSQL string

//go:embed sql/list_permissions.sql
var listPermissionsSQL string

type Repo struct {
	db  *sql.DB
	log *slog.Logger
}

func NewPermissionsRepo(db *sql.DB, log *slog.Logger) *Repo { return &Repo{db: db, log: log} }

type CreateArgs struct {
	Name      string
	GuardName string
}

func (r *Repo) CreateContext(ctx context.Context, tx repositories.DbTx, a CreateArgs) (int, error) {
	var id int
	if err := tx.QueryRowContext(ctx, insertPermissionSQL, a.Name, a.GuardName).Scan(&id); err != nil {
		return 0, fmt.Errorf("insert permission: %w", err)
	}
	return id, nil
}

type Row struct {
	ID        int
	Name      string
	GuardName string
}

func (r *Repo) GetByNameContext(ctx context.Context, tx repositories.DbTx, name string) (*Row, error) {
	var perm Row
	if err := tx.QueryRowContext(ctx, getPermissionByNameSQL, name).Scan(&perm.ID, &perm.Name, &perm.GuardName); err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, fmt.Errorf("get permission: %w", err)
	}
	return &perm, nil
}

func (r *Repo) ListContext(ctx context.Context, tx repositories.DbTx) ([]Row, error) {
	rows, err := tx.QueryContext(ctx, listPermissionsSQL)
	if err != nil {
		return nil, fmt.Errorf("list permissions: %w", err)
	}
	defer rows.Close()
	var out []Row
	for rows.Next() {
		var perm Row
		if err := rows.Scan(&perm.ID, &perm.Name, &perm.GuardName); err != nil {
			return nil, fmt.Errorf("scan permission: %w", err)
		}
		out = append(out, perm)
	}
	return out, rows.Err()
}
