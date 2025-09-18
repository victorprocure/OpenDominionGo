package domain

import (
	"time"
)

type Notification struct {
	Type           string
	NotifiableType string
	NotifiableId   int64
	Data           string
	ReadAt         time.Time
}
