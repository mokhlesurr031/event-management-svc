package usecase

import (
	"context"

	"github.com/mokhlesurr031/event-management-svc/domain"
)

func New(repo domain.WorkshopRepository) domain.WorkshopUseCase {
	return &WorkshopUseCase{
		repo: repo,
	}
}

type WorkshopUseCase struct {
	repo domain.WorkshopRepository
}

func (w *WorkshopUseCase) WorkshopList(ctx context.Context, ctr *domain.WorkshopCriteria) (*domain.WorkshopList, error) {
	return w.repo.WorkshopList(ctx, ctr)
}

func (w *WorkshopUseCase) WorkshopDetails(ctx context.Context, ctr *domain.WorkshopCriteria) (*domain.WorkshopDetails, error) {
	return w.repo.WorkshopDetails(ctx, ctr)
}
