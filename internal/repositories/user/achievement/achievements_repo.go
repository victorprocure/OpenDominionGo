package achievement

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log/slog"

	"github.com/victorprocure/opendominiongo/internal/repositories"
)

//go:embed sql/insert_user_achievement.sql
var insertUserAchievementSQL string

//go:embed sql/list_user_achievements_by_user.sql
var listUserAchievementsByUserSQL string

//go:embed sql/delete_user_achievement.sql
var deleteUserAchievementSQL string

type Repo struct {
	db  *sql.DB
	log *slog.Logger
}

func NewRepo(db *sql.DB, log *slog.Logger) *Repo { return &Repo{db: db, log: log} }

type CreateArgs struct {
	UserID        int
	AchievementID int
}

func (r *Repo) CreateContext(ctx context.Context, tx repositories.DbTx, a CreateArgs) (int, error) {
	var id int
	if err := tx.QueryRowContext(ctx, insertUserAchievementSQL, a.UserID, a.AchievementID).Scan(&id); err != nil {
		return 0, fmt.Errorf("insert user_achievement: %w", err)
	}
	return id, nil
}

type Row struct {
	ID            int
	UserID        int
	AchievementID int
}

func (r *Repo) ListByUserContext(ctx context.Context, tx repositories.DbTx, userID int, limit, offset int) ([]Row, error) {
	rows, err := tx.QueryContext(ctx, listUserAchievementsByUserSQL, userID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("list user_achievements: %w", err)
	}
	defer rows.Close()
	var out []Row
	for rows.Next() {
		var x Row
		if err := rows.Scan(&x.ID, &x.UserID, &x.AchievementID); err != nil {
			return nil, fmt.Errorf("scan user_achievement: %w", err)
		}
		out = append(out, x)
	}
	return out, rows.Err()
}

func (r *Repo) DeleteContext(ctx context.Context, tx repositories.DbTx, id int) error {
	if _, err := tx.ExecContext(ctx, deleteUserAchievementSQL, id); err != nil {
		return fmt.Errorf("delete user_achievement: %w", err)
	}
	return nil
}
