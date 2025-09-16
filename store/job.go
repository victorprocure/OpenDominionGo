package store

type Job struct {
	Id int64
	Queue string
	Payload string
	Attempts int16
	ReservedAt int
	AvailableAt int
	CreatedAt int
}