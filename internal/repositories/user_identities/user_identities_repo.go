package user_identities

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log/slog"

	"github.com/victorprocure/opendominiongo/internal/repositories"
)

//go:embed sql/insert_user_identity.sql
var insertUserIdentitySQL string

//go:embed sql/list_user_identities_by_user.sql
var listUserIdentitiesByUserSQL string

type Repo struct {
	db  *sql.DB
	log *slog.Logger
}

func NewUserIdentitiesRepo(db *sql.DB, log *slog.Logger) *Repo { return &Repo{db: db, log: log} }

type CreateArgs struct {
	UserID      int
	Fingerprint *string
	UserAgent   *string
	Count       int
}

func (r *Repo) CreateContext(ctx context.Context, tx repositories.DbTx, a CreateArgs) (int, error) {
	var id int
	if err := tx.QueryRowContext(ctx, insertUserIdentitySQL, a.UserID, a.Fingerprint, a.UserAgent, a.Count).Scan(&id); err != nil {
		return 0, fmt.Errorf("insert user_identity: %w", err)
	}
	return id, nil
}

type Row struct {
	ID          int
	UserID      int
	Fingerprint sql.NullString
	UserAgent   sql.NullString
	Count       int
}

func (r *Repo) ListByUserContext(ctx context.Context, tx repositories.DbTx, userID int, limit, offset int) ([]Row, error) {
	rows, err := tx.QueryContext(ctx, listUserIdentitiesByUserSQL, userID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("list user_identities: %w", err)
	}
	defer rows.Close()
	var out []Row
	for rows.Next() {
		var x Row
		if err := rows.Scan(&x.ID, &x.UserID, &x.Fingerprint, &x.UserAgent, &x.Count); err != nil {
			return nil, fmt.Errorf("scan user_identity: %w", err)
		}
		out = append(out, x)
	}
	return out, rows.Err()
}
