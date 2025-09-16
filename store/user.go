package store

import (
	"database/sql"
	"errors"
	"time"
)

type User struct {
	Id                   int
	Email                string
	PasswordHash         string
	DisplayName          string
	RememberToken        string
	Activated            bool
	ActivationCode       string
	LastOnline           time.Time
	CreatedAt            time.Time
	UpdatedAt            time.Time
	Avatar               string
	Settings             string
	Skin                 string
	MessageBoardLastRead time.Time
	Rating               int
	IsDeleted            bool
}

func (s *Storage) GetUserById(id int) (*User, error) {
	var u User
	err := s.db.QueryRow(
		"SELECT id, email, password, display_name, remember_token, activated, activation_code, last_online, created_at, updated_at, avatar, settings, skin, message_board_last_read, rating, isDeleted FROM users WHERE id = $1",
		id,
	).Scan(
		&u.Id,
		&u.Email,
		&u.PasswordHash,
		&u.DisplayName,
		&u.RememberToken,
		&u.Activated,
		&u.ActivationCode,
		&u.LastOnline,
		&u.CreatedAt,
		&u.UpdatedAt,
		&u.Avatar,
		&u.Settings,
		&u.Skin,
		&u.MessageBoardLastRead,
		&u.Rating,
		&u.IsDeleted,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrorNotFound
		}
		return nil, err
	}

	return &u, nil
}

func (s *Storage) GetUserByEmail(email string) (*User, error) {
	var u User
	err := s.db.QueryRow(
		"SELECT id, email, password, display_name, remember_token, activated, activation_code, last_online, created_at, updated_at, avatar, settings, skin, message_board_last_read, rating, isDeleted FROM users WHERE email = $1",
		email,
	).Scan(
		&u.Id,
		&u.Email,
		&u.PasswordHash,
		&u.DisplayName,
		&u.RememberToken,
		&u.Activated,
		&u.ActivationCode,
		&u.LastOnline,
		&u.CreatedAt,
		&u.UpdatedAt,
		&u.Avatar,
		&u.Settings,
		&u.Skin,
		&u.MessageBoardLastRead,
		&u.Rating,
		&u.IsDeleted,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrorNotFound
		}
		return nil, err
	}

	return &u, nil
}

func (s *Storage) CreateOrUpdateUser(u *User) (*User, error) {
	err := s.db.QueryRow(
		`INSERT INTO users (
		 email, password, display_name, remember_token, activated,
		 activation_code, last_online, avatar, settings,
		 skin, message_board_last_read, rating)
		 VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
		 ON CONFLICT (email) DO UPDATE SET
		 	email = $1,
			password = $2,
			display_name = $3,
			remember_token = $4,
			activated = $5,
			activation_code = $6,
			last_online = $7,
			avatar = $8,
			settings = $9,
			skin = $10,
			message_board_last_read = $11,
			rating = $12
		 RETURNING id`,
		u.Email,
		u.PasswordHash,
		u.DisplayName,
		u.RememberToken,
		u.Activated,
		u.ActivationCode,
		u.LastOnline,
		u.Avatar,
		u.Settings,
		u.Skin,
		u.MessageBoardLastRead,
		u.Rating,
	).Scan(&u.Id)

	if err != nil {
		return nil, err
	}

	return u, nil
}

func (s *Storage) DeleteUser(id int) error {
	_, err := s.db.Exec(
		"UPDATE users SET is_deleted = TRUE WHERE id = $2",
		id,
	)

	return err
}

type UserAchievement struct {
	Id          int
	User        *User
	Achievement *Achievement
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type UserActivity struct {
	Id        int
	User      *User
	Ip        string
	Key       string
	Context   string
	CreatedAt time.Time
	Status    string
	Device    string
}

type UserDiscordUser struct {
	Id            int
	User          *User
	DiscordUserId string
	Username      string
	Discriminator int
	Email         string
	RefreshToken  string
	ExpiresAt     time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type UserFeedback struct {
	Id        int
	SourceId  int
	TargetId  int
	Endorsed  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserIdentity struct {
	Id          int
	User        *User
	Fingerprint string
	UserAgent   string
	Count       int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type UserOrigin struct {
	Id        int
	User      *User
	Dominion  *Dominion
	IPAddress string
	Count     int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserOriginLookup struct {
	Id           int
	IPAddress    string
	ISP          string
	Organization string
	Country      string
	Region       string
	City         string
	VPN          bool
	Score        float64
	Data         string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
