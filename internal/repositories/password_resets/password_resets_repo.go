package password_resets

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log/slog"

	"github.com/victorprocure/opendominiongo/internal/repositories"
)

//go:embed sql/upsert_password_reset.sql
var upsertPasswordResetSQL string

//go:embed sql/get_password_reset_by_email.sql
var getPasswordResetByEmailSQL string

//go:embed sql/delete_password_reset.sql
var deletePasswordResetSQL string

type Repo struct {
	db  *sql.DB
	log *slog.Logger
}

func NewPasswordResetsRepo(db *sql.DB, log *slog.Logger) *Repo { return &Repo{db: db, log: log} }

type UpsertArgs struct {
	Email string
	Token string
}

func (r *Repo) UpsertContext(ctx context.Context, tx repositories.DbTx, a UpsertArgs) error {
	if _, err := tx.ExecContext(ctx, upsertPasswordResetSQL, a.Email, a.Token); err != nil {
		return fmt.Errorf("upsert password_reset: %w", err)
	}
	return nil
}

type Row struct {
	Email     string
	Token     string
	CreatedAt sql.NullTime
}

func (r *Repo) GetByEmailContext(ctx context.Context, tx repositories.DbTx, email string) (*Row, error) {
	var pr Row
	if err := tx.QueryRowContext(ctx, getPasswordResetByEmailSQL, email).Scan(&pr.Email, &pr.Token, &pr.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, fmt.Errorf("get password_reset: %w", err)
	}
	return &pr, nil
}

func (r *Repo) DeleteByEmailContext(ctx context.Context, tx repositories.DbTx, email string) error {
	if _, err := tx.ExecContext(ctx, deletePasswordResetSQL, email); err != nil {
		return fmt.Errorf("delete password_reset: %w", err)
	}
	return nil
}
