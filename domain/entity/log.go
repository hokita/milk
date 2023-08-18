package domain

import "time"

type LogType int

const (
	Milk LogType = iota
	Poo
	Pee
)

type Log struct {
	LogType LogType
	Date    time.Time
}
