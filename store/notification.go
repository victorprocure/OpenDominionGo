package store

import (
	"time"

	"github.com/google/uuid"
)

type Notification struct {
	Id             uuid.UUID
	Type           string
	NotifiableType string
	NotifiableId   int64
	Data           string
	ReadAt         time.Time
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
