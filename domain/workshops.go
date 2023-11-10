package domain

import (
	"context"
	"time"
)

type Workshop struct {
	Id          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	EventId     uint      `json:"event_id"`
	StartAt     time.Time `json:"start_at"`
	EndAt       time.Time `json:"end_at"`
}

type WorkshopCriteria struct {
	Id          *uint      `json:"id"`
	Title       *string    `json:"title"`
	Description *string    `json:"description"`
	EventId     *uint      `json:"event_id"`
	StartAt     *time.Time `json:"start_at"`
	EndAt       *time.Time `json:"end_at"`
}

type WorkshopRepository interface {
	WorkshopList(ctx context.Context, ctr *WorkshopCriteria) ([]*Workshop, error)
}

type WorkshopUseCase interface {
	WorkshopList(ctx context.Context, ctr *WorkshopCriteria) ([]*Workshop, error)
}
