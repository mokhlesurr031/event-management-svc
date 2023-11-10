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

type WorkshopDetails struct {
	Id               *uint      `json:"id"`
	Title            *string    `json:"title"`
	Description      *string    `json:"description"`
	StartAt          *time.Time `json:"start_at"`
	EndAt            *time.Time `json:"end_at"`
	TotaReservations int        `json:"total_reservations"`
}

type WorkshopRepository interface {
	WorkshopList(ctx context.Context, ctr *WorkshopCriteria) (*WorkshopList, error)
	WorkshopDetails(ctx context.Context, ctr *WorkshopCriteria) (*WorkshopDetails, error)
}

type WorkshopUseCase interface {
	WorkshopList(ctx context.Context, ctr *WorkshopCriteria) (*WorkshopList, error)
	WorkshopDetails(ctx context.Context, ctr *WorkshopCriteria) (*WorkshopDetails, error)
}
