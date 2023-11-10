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

func (e *WorkshopUseCase) WorkshopList(ctx context.Context, ctr *domain.WorkshopCriteria) ([]*domain.Workshop, error) {
	return e.repo.WorkshopList(ctx, ctr)
}
