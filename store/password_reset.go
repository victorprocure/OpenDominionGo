package store

import "time"

type PasswordReset struct {
	Email     string
	Token     string
	CreatedAt time.Time
}
