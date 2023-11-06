package domain

import "time"

type Events struct {
	Id      uint      `json:"id"`
	Title   string    `json:"title"`
	StartAt time.Time `json:"start_at"`
	EndAt   time.Time `json:"end_at"`
}

type EventsCriteria struct {
	Id      *uint      `json:"id"`
	Title   *string    `json:"title"`
	StartAt *time.Time `json:"start_at"`
	EndAt   *time.Time `json:"end_at"`
}
