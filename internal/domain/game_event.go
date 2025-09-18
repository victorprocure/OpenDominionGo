package domain

type GameEvent struct {
	Round      *Round
	SourceType string
	SourceId   int64
	TargetType string
	TargetId   int64
	Type       string
	Data       string
}