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

func (e *EventUseCase) EventList(ctx context.Context, ctr *domain.EventCriteria) ([]*domain.Event, error) {
	return e.repo.EventList(ctx, ctr)
}
