package store

import (
	"time"

	"github.com/google/uuid"
)

type TelescopeEntry struct {
	Sequence             int64
	UUID                 uuid.UUID
	BatchId              uuid.UUID
	FamilyHash           string
	ShouldDisplayOnIndex bool
	Type                 string
	Content              string
	CreatedAt            time.Time
}

type TelescopeEntryTag struct {
	Entry *TelescopeEntry
	Tag   string
}

type TelescopeMonitor struct {
	Tag string
}
