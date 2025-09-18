package roles

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log/slog"

	"github.com/victorprocure/opendominiongo/internal/repositories"
)

//go:embed sql/insert_role.sql
var insertRoleSQL string

//go:embed sql/get_role_by_name.sql
var getRoleByNameSQL string

//go:embed sql/list_roles.sql
var listRolesSQL string

type Repo struct {
	db  *sql.DB
	log *slog.Logger
}

func NewRolesRepo(db *sql.DB, log *slog.Logger) *Repo { return &Repo{db: db, log: log} }

type CreateArgs struct {
	Name      string
	GuardName string
}

func (r *Repo) CreateContext(ctx context.Context, tx repositories.DbTx, a CreateArgs) (int, error) {
	var id int
	if err := tx.QueryRowContext(ctx, insertRoleSQL, a.Name, a.GuardName).Scan(&id); err != nil {
		return 0, fmt.Errorf("insert role: %w", err)
	}
	return id, nil
}

type Row struct {
	ID        int
	Name      string
	GuardName string
}

func (r *Repo) GetByNameContext(ctx context.Context, tx repositories.DbTx, name string) (*Row, error) {
	var role Row
	if err := tx.QueryRowContext(ctx, getRoleByNameSQL, name).Scan(&role.ID, &role.Name, &role.GuardName); err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, fmt.Errorf("get role: %w", err)
	}
	return &role, nil
}

func (r *Repo) ListContext(ctx context.Context, tx repositories.DbTx) ([]Row, error) {
	rows, err := tx.QueryContext(ctx, listRolesSQL)
	if err != nil {
		return nil, fmt.Errorf("list roles: %w", err)
	}
	defer rows.Close()
	var out []Row
	for rows.Next() {
		var role Row
		if err := rows.Scan(&role.ID, &role.Name, &role.GuardName); err != nil {
			return nil, fmt.Errorf("scan role: %w", err)
		}
		out = append(out, role)
	}
	return out, rows.Err()
}
