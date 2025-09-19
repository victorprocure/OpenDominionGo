package activity

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log/slog"

	"github.com/victorprocure/opendominiongo/internal/repositories"
)

//go:embed sql/insert_user_activity.sql
var insertUserActivitySQL string

//go:embed sql/list_user_activities_by_user.sql
var listUserActivitiesByUserSQL string

type Repo struct {
	db  *sql.DB
	log *slog.Logger
}

func NewRepo(db *sql.DB, log *slog.Logger) *Repo { return &Repo{db: db, log: log} }

type CreateArgs struct {
	UserID  int
	IP      string
	Key     string
	Context string
	Status  *string
	Device  *string
}

func (r *Repo) CreateContext(ctx context.Context, tx repositories.DbTx, a CreateArgs) (int, error) {
	var id int
	if err := tx.QueryRowContext(ctx, insertUserActivitySQL, a.UserID, a.IP, a.Key, a.Context, a.Status, a.Device).Scan(&id); err != nil {
		return 0, fmt.Errorf("insert user_activity: %w", err)
	}
	return id, nil
}

type Row struct {
	ID      int
	UserID  int
	IP      string
	Key     string
	Context string
	Status  sql.NullString
	Device  sql.NullString
}

func (r *Repo) ListByUserContext(ctx context.Context, tx repositories.DbTx, userID, limit, offset int) ([]Row, error) {
	rows, err := tx.QueryContext(ctx, listUserActivitiesByUserSQL, userID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("list user_activities: %w", err)
	}
	defer rows.Close()
	var out []Row
	for rows.Next() {
		var x Row
		if err := rows.Scan(&x.ID, &x.UserID, &x.IP, &x.Key, &x.Context, &x.Status, &x.Device); err != nil {
			return nil, fmt.Errorf("scan user_activity: %w", err)
		}
		out = append(out, x)
	}
	return out, rows.Err()
}
