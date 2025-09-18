package achievement

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log/slog"

	"github.com/victorprocure/opendominiongo/internal/repositories"
)

//go:embed sql/insert_achievement.sql
var insertAchievementSQL string

//go:embed sql/get_achievement_by_name.sql
var getAchievementByNameSQL string

//go:embed sql/update_achievement.sql
var updateAchievementSQL string

//go:embed sql/delete_achievement.sql
var deleteAchievementSQL string

//go:embed sql/list_achievements.sql
var listAchievementsSQL string

type Repo struct {
	db  *sql.DB
	log *slog.Logger
}

func NewAchievementsRepo(db *sql.DB, log *slog.Logger) *Repo { return &Repo{db: db, log: log} }

type CreateArgs struct {
	Name        *string
	Description *string
	Icon        *string
}

func (r *Repo) CreateContext(ctx context.Context, tx repositories.DbTx, a CreateArgs) (int, error) {
	var id int
	if err := tx.QueryRowContext(ctx, insertAchievementSQL, a.Name, a.Description, a.Icon).Scan(&id); err != nil {
		return 0, fmt.Errorf("insert achievement: %w", err)
	}
	return id, nil
}

type Row struct {
	ID          int
	Name        *string
	Description *string
	Icon        *string
}

func (r *Repo) GetByNameContext(ctx context.Context, tx repositories.DbTx, name string) (*Row, error) {
	var ac Row
	if err := tx.QueryRowContext(ctx, getAchievementByNameSQL, name).Scan(&ac.ID, &ac.Name, &ac.Description, &ac.Icon); err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, fmt.Errorf("get achievement: %w", err)
	}
	return &ac, nil
}

type UpdateArgs struct {
	ID          int
	Name        *string
	Description *string
	Icon        *string
}

func (r *Repo) UpdateContext(ctx context.Context, tx repositories.DbTx, a UpdateArgs) error {
	if _, err := tx.ExecContext(ctx, updateAchievementSQL, a.ID, a.Name, a.Description, a.Icon); err != nil {
		return fmt.Errorf("update achievement: %w", err)
	}
	return nil
}

func (r *Repo) DeleteByIDContext(ctx context.Context, tx repositories.DbTx, id int) error {
	if _, err := tx.ExecContext(ctx, deleteAchievementSQL, id); err != nil {
		return fmt.Errorf("delete achievement: %w", err)
	}
	return nil
}

func (r *Repo) ListContext(ctx context.Context, tx repositories.DbTx, limit, offset int) ([]Row, error) {
	rows, err := tx.QueryContext(ctx, listAchievementsSQL, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("list achievements: %w", err)
	}
	defer rows.Close()
	var out []Row
	for rows.Next() {
		var ac Row
		if err := rows.Scan(&ac.ID, &ac.Name, &ac.Description, &ac.Icon); err != nil {
			return nil, fmt.Errorf("scan achievement: %w", err)
		}
		out = append(out, ac)
	}
	return out, rows.Err()
}
