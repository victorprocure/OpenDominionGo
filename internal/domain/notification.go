package domain

import (
	"time"
)

type Notification struct {
	Type           string
	NotifiableType string
	NotifiableID   int64
	Data           string
	ReadAt         time.Time
}
