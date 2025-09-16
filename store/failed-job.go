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