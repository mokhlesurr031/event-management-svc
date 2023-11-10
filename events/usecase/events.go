package usecase

import (
	"context"

	"github.com/mokhlesurr031/event-management-svc/domain"
)

func New(repo domain.EventRepository) domain.EventUseCase {
	return &EventUseCase{
		repo: repo,
	}
}

type EventUseCase struct {
	repo domain.EventRepository
}

func (e *EventUseCase) EventList(ctx context.Context, ctr *domain.EventCriteria) (*domain.EventCriteriaPagination, error) {
	return e.repo.EventList(ctx, ctr)
}

func (c *EventUseCase) EventDetails(ctx context.Context, ctr *domain.EventDetails) (*domain.EventDetails, error) {
	return c.repo.EventDetails(ctx, ctr)
}
