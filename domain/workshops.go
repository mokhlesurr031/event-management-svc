package domain

import "time"

type Workshops struct {
	Id          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	EventId     uint      `json:"event_id"`
	StartAt     time.Time `json:"start_at"`
	EndAt       time.Time `json:"end_at"`
}

type WorkshopsCriteria struct {
	Id          *uint      `json:"id"`
	Title       *string    `json:"title"`
	Description *string    `json:"description"`
	EventId     *uint      `json:"event_id"`
	StartAt     *time.Time `json:"start_at"`
	EndAt       *time.Time `json:"end_at"`
}
