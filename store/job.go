package store

import "time"

type FailedJob struct {
	Id         int64
	Connection string
	Queue      string
	Payload    string
	Exception  string
	FailedAt   time.Time
}

type Job struct {
	Id          int64
	Queue       string
	Payload     string
	Attempts    int16
	ReservedAt  int
	AvailableAt int
	CreatedAt   int
}