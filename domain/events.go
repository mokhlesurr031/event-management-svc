package domain

import (
	"context"
	"time"
)

type Event struct {
	Id      uint      `json:"id"`
	Title   string    `json:"title"`
	StartAt time.Time `json:"start_at"`
	EndAt   time.Time `json:"end_at"`
}

type EventCriteria struct {
	Id      *uint      `json:"id"`
	Title   *string    `json:"title"`
	StartAt *time.Time `json:"start_at"`
	EndAt   *time.Time `json:"end_at"`
}

type EventRepository interface {
	EventList(ctx context.Context, ctr *EventCriteria) ([]*Event, error)
}

type EventUseCase interface {
	EventList(ctx context.Context, ctr *EventCriteria) ([]*Event, error)
}
