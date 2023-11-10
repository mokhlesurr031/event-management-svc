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

type EventDetails struct {
	Id             *uint      `json:"id"`
	Title          *string    `json:"title"`
	StartAt        *time.Time `json:"start_at"`
	EndAt          *time.Time `json:"end_at"`
	TotalWorkshops int        `json:"total_workshops"`
}

type EventRepository interface {
	EventList(ctx context.Context, ctr *EventCriteria) (*EventCriteriaPagination, error)
	EventDetails(ctx context.Context, ctr *EventDetails) (*EventDetails, error)
}

type EventUseCase interface {
	EventList(ctx context.Context, ctr *EventCriteria) (*EventCriteriaPagination, error)
	EventDetails(ctx context.Context, ctr *EventDetails) (*EventDetails, error)
}
