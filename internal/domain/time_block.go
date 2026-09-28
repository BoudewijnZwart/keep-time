package domain

import "time"

type TimeBlock struct {
	Id int
	Started_at time.Time
	Stopped_at time.Time
	Notes string
	Project *Project
}