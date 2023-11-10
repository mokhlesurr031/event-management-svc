package usecase

import (
	"context"

	"github.com/mokhlesurr031/event-management-svc/domain"
)

func New(repo domain.ReservationRepository) domain.ReservationUseCase {
	return &ReservationsUseCase{
		repo: repo,
	}
}

type ReservationsUseCase struct {
	repo domain.ReservationRepository
}

func (rsvn *ReservationsUseCase) ReservationsCreate(ctx context.Context, ctr *domain.Reservation) (*domain.Reservation, error) {
	return rsvn.repo.ReservationsCreate(ctx, ctr)
}
