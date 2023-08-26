package domain

import "time"

type LogType int

const (
	Milk LogType = iota
	Poo
	Pee
)

type Log struct {
	ID        string
	LogType   LogType
	CheckinAt time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
