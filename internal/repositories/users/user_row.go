package users

import "time"

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
