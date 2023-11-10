package domain

import (
	"context"
)

type Reservation struct {
	Id         uint   `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	WorkshopId uint   `json:"workshop_id"`
}

type ReservationCriteria struct {
	Id         *uint   `json:"id"`
	Name       *string `json:"name"`
	Email      *string `json:"email"`
	WorkshopId *uint   `json:"workshop_id"`
}

type ReservationRepository interface {
	ReservationsCreate(ctx context.Context, ctr *Reservation) (*Reservation, error)
}

type ReservationUseCase interface {
	ReservationsCreate(ctx context.Context, ctr *Reservation) (*Reservation, error)
}
