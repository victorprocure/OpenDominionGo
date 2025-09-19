package user

import (
	"context"
	"database/sql"
	_ "embed"
	"log/slog"
	"time"

	"github.com/victorprocure/opendominiongo/internal/domain"
	"github.com/victorprocure/opendominiongo/internal/repositories"
)

//go:embed sql/get_users.sql
var getUsersSQL string

//go:embed sql/upsert_user.sql
var upsertUserSQL string

type Repo struct {
	db  *sql.DB
	log *slog.Logger
}

func NewRepo(db *sql.DB, log *slog.Logger) *Repo {
	return &Repo{db: db, log: log}
}

func (r *Repo) UpsertContext(ctx context.Context, tx repositories.DbTx, u *domain.User) error {
	args := fromDomain(u)
	err := tx.QueryRowContext(ctx, upsertUserSQL, args...).Scan()
	if err != nil {
		return err
	}

	return nil
}

func (r *Repo) GetAllUsersContext(ctx context.Context, tx repositories.DbTx) ([]*domain.User, error) {
	rows, err := tx.QueryContext(ctx, getUsersSQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*domain.User
	for rows.Next() {
		u, err := scanUserRow(rows)
		if err != nil {
			return nil, err
		}
		users = append(users, toDomain(u))
	}

	return users, rows.Err()
}

func (r *Repo) GetAllUsers() ([]*domain.User, error) {
	return r.GetAllUsersContext(context.Background(), r.db)
}

func (r *Repo) GetUserByIDContext(ctx context.Context, tx repositories.DbTx, i int) (*domain.User, error) {
	row := tx.QueryRowContext(ctx, getUsersSQL+"WHERE id = $1", i)
	u, err := scanUserRow(row)
	if err != nil {
		return nil, err
	}
	return toDomain(u), nil
}

func (r *Repo) GetUserByID(id int) (*domain.User, error) {
	return r.GetUserByIDContext(context.Background(), r.db, id)
}

func (r *Repo) GetUserByEmailContext(ctx context.Context, tx repositories.DbTx, e string) (*domain.User, error) {
	row := tx.QueryRowContext(ctx, getUsersSQL+"WHERE email = $1", e)
	u, err := scanUserRow(row)
	if err != nil {
		return nil, err
	}
	return toDomain(u), nil
}

func (r *Repo) GetUserByEmail(e string) (*domain.User, error) {
	return r.GetUserByEmailContext(context.Background(), r.db, e)
}

func scanUserRow(s repositories.RowScanner) (*userRow, error) {
	var u userRow
	if err := s.Scan(&u.ID,
		&u.Activated, &u.ActivationCode, &u.Avatar,
		&u.CreatedAt, &u.DisplayName,
		&u.Email, &u.IsDeleted, &u.LastOnline,
		&u.MessageBoardLastRead, &u.PasswordHash,
		&u.Rating, &u.RememberToken, &u.Settings,
		&u.Skin, &u.UpdatedAt,
	); err != nil {
		return nil, err
	}

	return &u, nil
}

func toDomain(r *userRow) *domain.User {
	return &domain.User{
		Activated:            r.Activated,
		Avatar:               r.Avatar,
		DisplayName:          r.DisplayName,
		Email:                r.Email,
		IsDeleted:            r.IsDeleted,
		LastOnline:           r.LastOnline,
		MessageBoardLastRead: r.MessageBoardLastRead,
		PasswordHash:         r.PasswordHash,
		Rating:               r.Rating,
		Settings:             r.Settings,
		Skin:                 r.Skin,
	}
}

func fromDomain(u *domain.User) []any {
	return []any{
		u.Activated,
		u.Avatar,
		u.DisplayName,
		u.Email,
		u.LastOnline,
		u.MessageBoardLastRead,
		u.PasswordHash,
		u.Rating,
		u.Settings,
		u.Skin,
	}
}

type userRow struct {
	ID                   int       `db:"id"`
	Activated            bool      `db:"activated"`
	ActivationCode       string    `db:"activation_code"`
	Avatar               string    `db:"avatar"`
	CreatedAt            time.Time `db:"created_at"`
	DisplayName          string    `db:"display_name"`
	Email                string    `db:"email"`
	IsDeleted            bool      `db:"is_deleted"`
	LastOnline           time.Time `db:"last_online"`
	MessageBoardLastRead time.Time `db:"message_board_last_read"`
	PasswordHash         string    `db:"password"`
	Rating               int       `db:"rating"`
	RememberToken        string    `db:"remember_token"`
	Settings             string    `db:"settings"`
	Skin                 string    `db:"skin"`
	UpdatedAt            time.Time `db:"updated_at"`
}
