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

type WorkshopList struct {
	Id       *uint       `json:"id"`
	Title    *string     `json:"title"`
	StartAt  *time.Time  `json:"start_at"`
	EndAt    *time.Time  `json:"end_at"`
	Workshop []*Workshop `json:"workshops"`
}

type WorkshopRepository interface {
	WorkshopList(ctx context.Context, ctr *WorkshopCriteria) (*WorkshopList, error)
}

type WorkshopUseCase interface {
	WorkshopList(ctx context.Context, ctr *WorkshopCriteria) (*WorkshopList, error)
}
