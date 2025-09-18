package notification

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log/slog"

	"github.com/victorprocure/opendominiongo/internal/repositories"
)

//go:embed sql/insert_notification.sql
var insertNotificationSQL string

//go:embed sql/list_notifications_by_notifiable.sql
var listNotificationsByNotifiableSQL string

type Repo struct {
	db  *sql.DB
	log *slog.Logger
}

func NewNotificationsRepo(db *sql.DB, log *slog.Logger) *Repo { return &Repo{db: db, log: log} }

type CreateArgs struct {
	Type           string
	NotifiableType string
	NotifiableID   int64
	Data           string
}

func (r *Repo) CreateContext(ctx context.Context, tx repositories.DbTx, a CreateArgs) (string, error) {
	var id string
	if err := tx.QueryRowContext(ctx, insertNotificationSQL, a.Type, a.NotifiableType, a.NotifiableID, a.Data).Scan(&id); err != nil {
		return "", fmt.Errorf("insert notification: %w", err)
	}
	return id, nil
}

type Row struct {
	ID             string
	Type           string
	NotifiableType string
	NotifiableID   int64
	Data           string
}

func (r *Repo) ListByNotifiableContext(ctx context.Context, tx repositories.DbTx, notifiableType string, notifiableID int64, limit, offset int) ([]Row, error) {
	rows, err := tx.QueryContext(ctx, listNotificationsByNotifiableSQL, notifiableType, notifiableID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("list notifications: %w", err)
	}
	defer rows.Close()
	var out []Row
	for rows.Next() {
		var n Row
		if err := rows.Scan(&n.ID, &n.Type, &n.NotifiableType, &n.NotifiableID, &n.Data); err != nil {
			return nil, fmt.Errorf("scan notification: %w", err)
		}
		out = append(out, n)
	}
	return out, rows.Err()
}
