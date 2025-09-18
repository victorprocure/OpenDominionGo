package domain

import (
	"github.com/google/uuid"
)

type TelescopeEntry struct {
	Sequence             int64
	BatchID              uuid.UUID
	FamilyHash           string
	ShouldDisplayOnIndex bool
	Type                 string
	Content              string
}

type TelescopeEntryTag struct {
	Entry *TelescopeEntry
	Tag   string
}

type TelescopeMonitor struct {
	Tag string
}
