package model

import "time"

type Entry struct {
	ID        string
	Name      string
	Data      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
