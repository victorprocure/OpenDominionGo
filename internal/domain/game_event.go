package domain

type GameEvent struct {
	Round      *Round
	SourceType string
	SourceID   int64
	TargetType string
	TargetID   int64
	Type       string
	Data       string
}
