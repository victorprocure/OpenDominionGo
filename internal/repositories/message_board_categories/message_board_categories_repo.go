package message_board_category

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log/slog"

	"github.com/victorprocure/opendominiongo/internal/repositories"
)

//go:embed sql/insert_message_board_category.sql
var insertCategorySQL string

//go:embed sql/get_message_board_category_by_slug.sql
var getCategoryBySlugSQL string

//go:embed sql/update_message_board_category.sql
var updateCategorySQL string

//go:embed sql/delete_message_board_category.sql
var deleteCategorySQL string

//go:embed sql/list_message_board_categories.sql
var listCategoriesSQL string

type Repo struct {
	db  *sql.DB
	log *slog.Logger
}

func NewMessageBoardCategoriesRepo(db *sql.DB, log *slog.Logger) *Repo {
	return &Repo{db: db, log: log}
}

type CreateArgs struct {
	Name         string
	Slug         string
	RoleRequired *string
}

func (r *Repo) CreateContext(ctx context.Context, tx repositories.DbTx, a CreateArgs) (int, error) {
	var id int
	if err := tx.QueryRowContext(ctx, insertCategorySQL, a.Name, a.Slug, a.RoleRequired).Scan(&id); err != nil {
		return 0, fmt.Errorf("insert message_board_category: %w", err)
	}
	return id, nil
}

type Row struct {
	ID           int
	Name         string
	Slug         string
	RoleRequired *string
}

func (r *Repo) GetBySlugContext(ctx context.Context, tx repositories.DbTx, slug string) (*Row, error) {
	var c Row
	if err := tx.QueryRowContext(ctx, getCategoryBySlugSQL, slug).Scan(&c.ID, &c.Name, &c.Slug, &c.RoleRequired); err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, fmt.Errorf("get message_board_category: %w", err)
	}
	return &c, nil
}

type UpdateArgs struct {
	ID           int
	Name         string
	Slug         string
	RoleRequired *string
}

func (r *Repo) UpdateContext(ctx context.Context, tx repositories.DbTx, a UpdateArgs) error {
	if _, err := tx.ExecContext(ctx, updateCategorySQL, a.ID, a.Name, a.Slug, a.RoleRequired); err != nil {
		return fmt.Errorf("update message_board_category: %w", err)
	}
	return nil
}

func (r *Repo) DeleteByIDContext(ctx context.Context, tx repositories.DbTx, id int) error {
	if _, err := tx.ExecContext(ctx, deleteCategorySQL, id); err != nil {
		return fmt.Errorf("delete message_board_category: %w", err)
	}
	return nil
}

func (r *Repo) ListContext(ctx context.Context, tx repositories.DbTx) ([]Row, error) {
	rows, err := tx.QueryContext(ctx, listCategoriesSQL)
	if err != nil {
		return nil, fmt.Errorf("list message_board_categories: %w", err)
	}
	defer rows.Close()
	var out []Row
	for rows.Next() {
		var c Row
		if err := rows.Scan(&c.ID, &c.Name, &c.Slug, &c.RoleRequired); err != nil {
			return nil, fmt.Errorf("scan message_board_category: %w", err)
		}
		out = append(out, c)
	}
	return out, rows.Err()
}
