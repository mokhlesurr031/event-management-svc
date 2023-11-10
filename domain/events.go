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

type EventCriteriaPagination struct {
	Events     []*Event   `json:"events"`
	Pagination Pagination `json:"pagination"`
}

type EventRepository interface {
	EventList(ctx context.Context, ctr *EventCriteria) (*EventCriteriaPagination, error)
}

type EventUseCase interface {
	EventList(ctx context.Context, ctr *EventCriteria) (*EventCriteriaPagination, error)
}
