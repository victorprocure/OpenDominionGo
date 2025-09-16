package store

import (
	"time"

	"github.com/google/uuid"
)

type GameEvent struct {
	Id         uuid.UUID
	Round      *Round
	SourceType string
	SourceId   int64
	TargetType string
	TargetId   int64
	Type       string
	Data       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}