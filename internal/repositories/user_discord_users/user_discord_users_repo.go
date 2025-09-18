package user_discord_users

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log/slog"

	"github.com/victorprocure/opendominiongo/internal/repositories"
)

//go:embed sql/insert_user_discord_user.sql
var insertUserDiscordUserSQL string

//go:embed sql/get_user_discord_user_by_user.sql
var getUserDiscordUserByUserSQL string

type Repo struct {
	db  *sql.DB
	log *slog.Logger
}

func NewUserDiscordUsersRepo(db *sql.DB, log *slog.Logger) *Repo { return &Repo{db: db, log: log} }

type CreateArgs struct {
	UserID        int
	DiscordUserID int64
	Username      string
	Discriminator int
	Email         *string
	RefreshToken  string
	ExpiresAt     *string
}

func (r *Repo) CreateContext(ctx context.Context, tx repositories.DbTx, a CreateArgs) (int, error) {
	var id int
	if err := tx.QueryRowContext(ctx, insertUserDiscordUserSQL, a.UserID, a.DiscordUserID, a.Username, a.Discriminator, a.Email, a.RefreshToken, a.ExpiresAt).Scan(&id); err != nil {
		return 0, fmt.Errorf("insert user_discord_user: %w", err)
	}
	return id, nil
}

type Row struct {
	ID            int
	UserID        int
	DiscordUserID int64
	Username      string
	Discriminator int
	Email         sql.NullString
	RefreshToken  string
}

func (r *Repo) GetByUserContext(ctx context.Context, tx repositories.DbTx, userID int) (*Row, error) {
	var x Row
	if err := tx.QueryRowContext(ctx, getUserDiscordUserByUserSQL, userID).
		Scan(&x.ID, &x.UserID, &x.DiscordUserID, &x.Username, &x.Discriminator, &x.Email, &x.RefreshToken); err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, fmt.Errorf("get user_discord_user: %w", err)
	}
	return &x, nil
}
