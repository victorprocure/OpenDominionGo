package user_origins

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log/slog"

	"github.com/victorprocure/opendominiongo/internal/repositories"
)

//go:embed sql/insert_user_origin.sql
var insertUserOriginSQL string

//go:embed sql/list_user_origins_by_user.sql
var listUserOriginsByUserSQL string

type Repo struct {
	db  *sql.DB
	log *slog.Logger
}

func NewUserOriginsRepo(db *sql.DB, log *slog.Logger) *Repo { return &Repo{db: db, log: log} }

type CreateArgs struct {
	UserID     int
	DominionID *int
	IPAddress  string
	Count      int
}

func (r *Repo) CreateContext(ctx context.Context, tx repositories.DbTx, a CreateArgs) (int, error) {
	var id int
	if err := tx.QueryRowContext(ctx, insertUserOriginSQL, a.UserID, a.DominionID, a.IPAddress, a.Count).Scan(&id); err != nil {
		return 0, fmt.Errorf("insert user_origin: %w", err)
	}
	return id, nil
}

type Row struct {
	ID         int
	UserID     int
	DominionID sql.NullInt64
	IPAddress  string
	Count      int
}

func (r *Repo) ListByUserContext(ctx context.Context, tx repositories.DbTx, userID int, limit, offset int) ([]Row, error) {
	rows, err := tx.QueryContext(ctx, listUserOriginsByUserSQL, userID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("list user_origins: %w", err)
	}
	defer rows.Close()
	var out []Row
	for rows.Next() {
		var x Row
		if err := rows.Scan(&x.ID, &x.UserID, &x.DominionID, &x.IPAddress, &x.Count); err != nil {
			return nil, fmt.Errorf("scan user_origin: %w", err)
		}
		out = append(out, x)
	}
	return out, rows.Err()
}
